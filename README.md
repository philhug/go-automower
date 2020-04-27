# Control your Husqvarna AutoMower

This project is not affiliated with Husqvarna and their APIs could change at any time and break this implementation.
Use at your own risk.

## Quick start
```go
c, err := client.NewClientWithUserAndPassword(MOWER_USER, MOWER_PW)
if err != nil {
        log.Println(err)
}
mowers, err := c.Mowers()
if err != nil {
        log.Println(err)
}
fmt.Println(mowers)
```

## Documentation
[https://pkg.go.dev/github.com/philhug/go-automower/pkg/automower?tab=doc](https://pkg.go.dev/github.com/philhug/go-automower/pkg/automower?tab=doc)

## TODO
Switch to official API
