# jsons

**jsons** is a Go library to work with JSONS files.

JSONS files contain JSON objects, one per line. This storage format is more
efficient than storing an array of objects since it doesn’t need to be loaded
in memory when reading/writing it.

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

            log.Printf("got %+v", m)
        }
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
