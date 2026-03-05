package crawler

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCrawlFiles(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("hello"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.go"), []byte("package main"), 0644)

	files, errs := CrawlFiles(tmpDir, "", false)

	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}

	if len(files) != 2 {
		t.Errorf("expected 2 files, got %d", len(files))
	}
}

func TestCrawlFilesWithExtensionFilter(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("hello"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file3.go"), []byte("package main"), 0644)

	files, errs := CrawlFiles(tmpDir, ".go", false)

	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}

	if len(files) != 2 {
		t.Errorf("expected 2 .go files, got %d", len(files))
	}

	for _, f := range files {
		if filepath.Ext(f.Name) != ".go" {
			t.Errorf("expected .go extension, got %s", f.Name)
		}
	}
}

func TestCrawlFilesNoMatch(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("hello"), 0644)

	files, errs := CrawlFiles(tmpDir, ".xyz", false)

	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}

	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}
}

func TestCrawlFilesNested(t *testing.T) {
	tmpDir := t.TempDir()

	subDir := filepath.Join(tmpDir, "subdir")
	os.MkdirAll(subDir, 0755)

	os.WriteFile(filepath.Join(tmpDir, "root.txt"), []byte("root"), 0644)
	os.WriteFile(filepath.Join(subDir, "nested.go"), []byte("nested"), 0644)

	files, errs := CrawlFiles(tmpDir, "", false)

	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}

	if len(files) != 2 {
		t.Errorf("expected 2 files, got %d", len(files))
	}
}
