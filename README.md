# gofind

A CLI tool that finds the largest files in a directory.

## Usage

```bash
go run ./cmd/gofind -path <directory> -count <number>
```

### Flags

- `-path`: Directory to search (default: current directory)
- `-count`: Number of top files to display (default: 10)

## Project Structure

```
cmd/gofind/main.go         # Entry point
internal/
  models/file.go           # FileResult type and utilities
  crawler/crawler.go       # Directory crawling logic
  sorter/sorter.go         # File sorting logic
```

## Building

```bash
go build ./cmd/gofind
```
