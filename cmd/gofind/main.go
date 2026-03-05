package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/kai-xlr/gofind/internal/crawler"
	"github.com/kai-xlr/gofind/internal/models"
	"github.com/kai-xlr/gofind/internal/sorter"
)

func printTopFiles(files []models.FileResult, count int) {
	if len(files) == 0 {
		fmt.Println("No files found.")
		return
	}

	limit := count
	if limit > len(files) {
		limit = len(files)
	}

	fmt.Printf("\nRANK\tSIZE\t\tNAME\t\tPATH\n")
	fmt.Printf("----\t----\t\t----\t\t----\n")
	for i, f := range files[:limit] {
		fmt.Printf("%d\t%-10s\t%-15s\t%s\n", i+1, f.PrettySize(), f.Name, f.Path)
	}
}

func printSummary(fileCount int, errorCount int, duration float64) {
	fmt.Printf("\nScanned %d files in %.2f seconds. Found %d errors.\n",
		fileCount, duration, errorCount)
}

func main() {
	start := time.Now()

	pathFlag := flag.String("path", "/", "Directory to search")
	countFlag := flag.Int("count", 10, "Number of top files to display")
	extFlag := flag.String("ext", "", "Filter by file extension (e.g., .go, .txt)")
	parallelFlag := flag.Bool("parallel", false, "Use parallel crawling for large directories")
	flag.Parse()

	files, errors := crawler.CrawlFiles(*pathFlag, *extFlag, *parallelFlag)

	sorter.SortFiles(files)

	printTopFiles(files, *countFlag)
	printSummary(len(files), len(errors), time.Since(start).Seconds())

	if len(errors) > 0 {
		fmt.Println("\nErrors encountered during crawl:")
		for _, e := range errors {
			fmt.Println("-", e)
		}
	}
}
