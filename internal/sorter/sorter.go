package sorter

import (
	"slices"

	"github.com/kai-xlr/gofind/internal/models"
)

func SortFiles(files []models.FileResult) {
	slices.SortFunc(files, func(a, b models.FileResult) int {
		switch {
		case b.Size > a.Size:
			return 1
		case b.Size < a.Size:
			return -1
		default:
			return 0
		}
	})
}
