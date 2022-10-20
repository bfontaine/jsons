# Jsons Changelog

## 1.1.0 (2022/10/20)

* Add `Writer.AddAll`
* `NewWriter` now return a `Writer` instead of a `*Writer`. Methods are updated to take a `Writer` as well.
* `FileWriter` and `FileReader` no longer reference the underlying file after `Close()`

## 1.0.1 (2022/10/18)

* Use `go mod`
* Fix a typo in the `NewFileWriter` documentation

## 1.0.0 (2015/07/12)

Initial release.
