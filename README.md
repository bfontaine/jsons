# jsons

[![Build Status](https://travis-ci.org/bfontaine/jsons.svg?branch=master)](https://travis-ci.org/bfontaine/jsons)
[![GoDoc](https://godoc.org/github.com/bfontaine/jsons?status.svg)](https://godoc.org/github.com/bfontaine/jsons)

**jsons** is a Go library to work with JSONS/[NDJSON][] files.

JSONS files contain JSON objects, one per line. This storage format is more
efficient than storing an array of objects since it doesn’t need to be loaded
in memory when reading/writing it.

[NDJSON]: http://ndjson.org/

## Install

    go get github.com/bfontaine/jsons

## Usage

### Reading

```go
func read() {
    j := jsons.NewFileReader("foo.jsons")

    if err := j.Open(); err != nil {
        log.Fatal(err)
    }

    defer j.Close()

    for {
        var m map[string]string
        if err := j.Next(&m); err != nil {
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }

        log.Printf("got %+v", m)
    }
}
```

### Writing

```go
func write() {
    j := jsons.NewFileWriter("foo.jsons")

    if err := j.Open(); err != nil {
        log.Fatal(err)
    }

    defer j.Close()

    j.Add(map[string]string{
        "foo": "bar",
        "qux": "abc",
    })

    j.Add(map[string]string{
        "foo": "def",
        "ghi": "jkl",
        "mno": "qpr",
    })
}
```
