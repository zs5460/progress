# progress

Console progress bar for Golang

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
