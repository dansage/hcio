package hcio

import "fmt"

// Options encapsulates all configurable aspects of the check pinging process
type Options struct {
	// BaseUrl specifies the pinging API URL of the Healthchecks instance being used. Default is `https://hc-ping.com/`.
	BaseUrl string

	// MaxRetries specifies the number of attempts to send the ping request before failing. Default is 3, must be over 0.
	MaxRetries uint8

	// UserAgent specifies which user agent string to send in all HTTP requests. Default is blank, using the Go standard.
	UserAgent string
}

// The DefaultOptions used if any are left unspecified
var DefaultOptions = Options{
	BaseUrl:    "https://hc-ping.com/",
	MaxRetries: 3,
}

// defaultOptions merges the user specified values with the defaults to ensure all options are set
func defaultOptions(options ...Options) Options {
	// use the default options if none were specified
	if len(options) == 0 {
		return DefaultOptions
	}

	// pull the first set of options
	opts := options[0]

	// use the default values if any are missing
	if opts.BaseUrl == "" {
		opts.BaseUrl = DefaultOptions.BaseUrl
	}
	if opts.MaxRetries == 0 {
		opts.MaxRetries = DefaultOptions.MaxRetries
	}

	// verify the base URL ends with a forward slash
	if opts.BaseUrl[len(opts.BaseUrl)-1:] != "/" {
		opts.BaseUrl = fmt.Sprintf("%s/", opts.BaseUrl)
	}

	return opts
}
