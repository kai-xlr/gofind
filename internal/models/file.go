package models

import "fmt"

type FileResult struct {
	Name string
	Path string
	Size int64
}

func (f FileResult) PrettySize() string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	size := float64(f.Size)

	switch {
	case f.Size < KB:
		return fmt.Sprintf("%d B", f.Size)
	case f.Size < MB:
		return fmt.Sprintf("%.2f KB", size/KB)
	case f.Size < GB:
		return fmt.Sprintf("%.2f MB", size/MB)
	default:
		return fmt.Sprintf("%.2f GB", size/GB)
	}
}
