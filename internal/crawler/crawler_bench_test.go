package crawler

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkCrawlFiles(b *testing.B) {
	tmpDir := b.TempDir()

	for i := 0; i < 1000; i++ {
		os.WriteFile(filepath.Join(tmpDir, "file"+string(rune(i))+"txt"), []byte("hello"), 0644)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CrawlFiles(tmpDir, "", false)
	}
}

func BenchmarkCrawlFilesParallel(b *testing.B) {
	tmpDir := b.TempDir()

	for i := 0; i < 1000; i++ {
		os.WriteFile(filepath.Join(tmpDir, "file"+string(rune(i))+"txt"), []byte("hello"), 0644)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CrawlFiles(tmpDir, "", true)
	}
}

func BenchmarkCrawlFilesWithExtFilter(b *testing.B) {
	tmpDir := b.TempDir()

	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			os.WriteFile(filepath.Join(tmpDir, "file"+string(rune(i))+".go"), []byte("hello"), 0644)
		} else {
			os.WriteFile(filepath.Join(tmpDir, "file"+string(rune(i))+".txt"), []byte("hello"), 0644)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CrawlFiles(tmpDir, ".go", false)
	}
}
