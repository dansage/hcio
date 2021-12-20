# HCio

HCio is a straightforward way to ping [Healthchecks.io checks][2] directly from a Go application.

## Getting Started

Create a simple `Check`:

```go
check := hcio.NewCheck("460a896b-e5e7-4d4a-996b-5ea0f3533a16")
```

To update the status of the check use the relevant status method:

```go
// success
check.Ping()
check.Success()

// failure
check.Fail()

// or a specific error code (0-255 only)
check.FailCode(51)
```

You can also time a longer operation and indicate success or failure as needed:

```go
// indicate the operation is starting
check.Start()

// do the real work
time.Sleep(30 * time.Second)

// indicate everything worked out just fine
check.Success()
```

You can also check a few options from the defaults

```go
check := hcio.NewCheck("460a896b-e5e7-4d4a-996b-5ea0f3533a16", hcio.Options{
	// the URL of the pinging API server, default is "https://hc-ping.com/"
	BaseUrl: "https://my-private-instance/",
	
	// the number of retries in the event of an HTTP failure, default is 3, must be over 0
	MaxRetries: 5,

	// the user agent to send with the requests, defaults to the Go standard user agent
	UserAgent: "mytool/0.1.0",
})
```

## License

HCio is open-source software released under the [MIT License][1].

[1]: https://choosealicense.com/licenses/mit/
[2]: https://github.com/healthchecks/healthchecks
