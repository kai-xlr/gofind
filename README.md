# gofind

A fast CLI tool that finds the largest files in a directory.

## Usage

```bash
gofind -path <directory> -count <number>
```

### Flags

- `-path`: Directory to search (default: "/")
- `-count`: Number of top files to display (default: 10)
- `-ext`: Filter by file extension (e.g., `.go`, `.txt`)
- `-parallel`: Use parallel crawling for large directories

### Examples

```bash
# Find top 10 largest files in current directory
gofind

# Find top 20 largest files in /var/log
gofind -path /var/log -count 20

# Find largest .go files
gofind -path ./src -ext .go

# Use parallel mode for very large directories
gofind -path / -count 50 -parallel
```

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

## Testing

```bash
go test ./...
```

## Benchmarking

```bash
go test -bench=. -benchtime=3s ./...
```
