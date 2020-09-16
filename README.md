# progress

Console progress bar for Golang

[![Build Status](https://travis-ci.org/zs5460/progress.svg?branch=master)](https://travis-ci.org/zs5460/progress)
[![Go Report Card](https://goreportcard.com/badge/github.com/zs5460/progress)](https://goreportcard.com/report/github.com/zs5460/progress)
[![codecov](https://codecov.io/gh/zs5460/progress/branch/master/graph/badge.svg)](https://codecov.io/gh/zs5460/progress)


## Install

```bash
go get github.com/zs5460/progress
```

## Example

```go
    count := 5460
    bar := progress.New(count)
    //bar.SetGraph("▫")

    for i := 0; i <= count; i++ {
        time.Sleep(1 * time.Millisecond)
        bar.Incr()
    }
```

```bash
[>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>           ]  80%     4341/5460
```

## License

Released under MIT license, see [LICENSE](LICENSE) for details.
