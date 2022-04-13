package main

import (
	"links"
	"log"
	"os"
)

func crawl(url string) []string {
	println(url)
	lists, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return lists

}
func main() {
	breadthFirstSearch(crawl, os.Args[1:])
}

func breadthFirstSearch(c func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, c(item)...)
			}
		}
	}
}
