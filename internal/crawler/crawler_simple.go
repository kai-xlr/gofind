package crawler

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/kai-xlr/gofind/internal/models"
)

func CrawlFilesSimple(path string, ext string) ([]models.FileResult, []error) {
	files := []models.FileResult{}
	errors := []error{}

	err := filepath.WalkDir(path, func(p string, f fs.DirEntry, err error) error {
		if err != nil {
			if p == path {
				errors = append(errors, fmt.Errorf("error accessing %s: %w", path, err))
			}
			return nil
		}

		if f.IsDir() {
			return nil
		}

		info, infoErr := f.Info()
		if infoErr != nil {
			errors = append(errors, fmt.Errorf("error getting info for %s: %w", p, infoErr))
			return nil
		}

		name := f.Name()
		if ext == "" || strings.HasSuffix(name, ext) {
			files = append(files, models.FileResult{Name: name, Path: p, Size: info.Size()})
		}
		return nil
	})

	if err != nil {
		errors = append(errors, err)
	}

	return files, errors
}
