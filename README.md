# Bolt.cm Documentation Linkchecker

Checks for broken links in the Bolt Documentation.

## Installation

```
go get github.com/PuerkitoBio/gocrawl
go get github.com/PhillippOhlandt/bolt-docs-linkchecker
```

## Usage

```
go run main.go
```

### Custom Hostname (for local testing)

```
go run main.go http://docs.bolt.dev
```

### Custom Versions (default is 2.2 and 3.0)

```
go run main.go http://docs.bolt.dev 3.0,3.1
```
