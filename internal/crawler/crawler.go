package crawler

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kai-xlr/gofind/internal/models"
)

func CrawlFiles(path string, ext string, parallel bool) ([]models.FileResult, []error) {
	if parallel {
		return crawlFilesParallel(path, ext)
	}
	return CrawlFilesSimple(path, ext)
}

func crawlFilesParallel(root string, ext string) ([]models.FileResult, []error) {
	fileChan := make(chan models.FileResult)
	errChan := make(chan error)
	var wg sync.WaitGroup

	wg.Add(1)
	go findFiles(root, fileChan, errChan, &wg)

	go func() {
		wg.Wait()
		close(fileChan)
		close(errChan)
	}()

	files := []models.FileResult{}
	errors := []error{}

	done := make(chan struct{})
	go func() {
		for f := range fileChan {
			if ext == "" || strings.HasSuffix(f.Name, ext) {
				files = append(files, f)
			}
		}
		done <- struct{}{}
	}()

	for e := range errChan {
		errors = append(errors, e)
	}

	<-done
	return files, errors
}

func findFiles(root string, fileChan chan<- models.FileResult, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	_ = filepath.WalkDir(root, func(path string, f fs.DirEntry, err error) error {
		if err != nil {
			if path == root {
				errChan <- fmt.Errorf("error accessing %s: %w", path, err)
			}
			return nil
		}

		if f.IsDir() {
			return nil
		}

		info, infoErr := f.Info()
		if infoErr != nil {
			errChan <- fmt.Errorf("error getting info for %s: %w", path, infoErr)
			return nil
		}

		fileChan <- models.FileResult{Name: f.Name(), Path: path, Size: info.Size()}
		return nil
	})
}
