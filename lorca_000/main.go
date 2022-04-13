package main

import (
	"embed"
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

var fs embed.FS

type counter struct {
	sync.Mutex
	count int
	//New interface{}
}

func (c *counter) New(n int) {
	c.Lock()
	defer c.Unlock()
}

func main(){
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	ui, err := lorca.New("", "", 1024, 512, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("start", func() {
		log.Printf("UI is ready")
	})

	c := &counter{}
	ui.Bind("New", c.New)


	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprint("http://%s/www", ln.Addr()))

	ui.Eval(`console.log("hello!")`)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
	case <-ui.Done():
	}

	fmt.Print("exiting...")
}
