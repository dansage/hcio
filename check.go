package hcio

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	responseOK          = "OK"
	responseNotFound    = "OK (not found)"
	responseRateLimited = "OK (rate limited)"
)

// A Check to monitor the status of a recurring task.
type Check struct {
	// ID is the UUID of the given check, obtained from the management API or interface of a Healthchecks instance
	ID string

	// the Options to use when pinging the check
	Options Options
}

// NewCheck creates a new check with the specified ID and options, if required
func NewCheck(id string, options ...Options) *Check {
	// merge the user specified options (if any) with the defaults
	opts := defaultOptions(options...)

	// create the check with the specified ID and merged options
	check := Check{
		ID:      id,
		Options: opts,
	}

	return &check
}

// sendPing sends a ping to the assembled URL
func (c *Check) sendPing(url string) error {
	// create the HTTP request itself
	r, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	// set the user agent in the request, if applicable
	if c.Options.UserAgent != "" {
		r.Header.Set("User-Agent", c.Options.UserAgent)
	}

	var attempts uint8
	var response *http.Response

	// iterate until there are no more attempts remaining
	for attempts < c.Options.MaxRetries {
		// check if this is the first attempt
		if attempts != 0 {
			// delay for an exponential amount of time: 1s, 2s, 4s by default
			time.Sleep(time.Duration(math.Pow(2, float64(attempts-1))) * time.Second)
		}

		// send the request to the server
		response, err = http.DefaultClient.Do(r)
		if err != nil {
			// increment the number of attempts
			attempts += 1
			continue
		}

		// break the loop, the request succeeded
		break
	}

	// verify a response was retrieved
	if response != nil {
		// ensure the body is closed
		defer response.Body.Close()

		// read the entire response body
		b, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		// convert the body into a string
		body := string(b)

		// verify the response is acceptable
		switch body {
		case responseOK:
			return nil

		case responseNotFound:
			return fmt.Errorf("the server could not find a check with ID: %q", c.ID)

		case responseRateLimited:
			return fmt.Errorf("the server indicates the check was pinged too frequently (5+ times in one minute)")
		}

		return fmt.Errorf("the server returned an unknown response: %v", body)
	}

	return err
}

// FailCode sends a signal to indicate the check has finished due to a failure with the specified code
func (c *Check) FailCode(code uint8) error {
	// send a simple ping to the fail URL
	return c.sendPing(fmt.Sprintf("%s%s/%d", c.Options.BaseUrl, c.ID, code))
}

// Fail sends a signal to indicate the check has finished due to a failure
func (c *Check) Fail() error {
	// send a simple ping to the fail URL
	return c.sendPing(fmt.Sprintf("%s%s/fail", c.Options.BaseUrl, c.ID))
}

// Ping sends a signal to indicate the check has finished successfully, or is still alive
func (c *Check) Ping() error {
	// send a simple ping to the check URL
	return c.sendPing(fmt.Sprintf("%s%s", c.Options.BaseUrl, c.ID))
}

// Start sends a signal to indicate the check is beginning execution
func (c *Check) Start() error {
	// send a simple ping to the start URL
	return c.sendPing(fmt.Sprintf("%s%s/start", c.Options.BaseUrl, c.ID))
}

// Success sends a signal to indicate the check has finished successfully
func (c *Check) Success() error {
	// send a simple ping to the check URL
	return c.sendPing(fmt.Sprintf("%s%s", c.Options.BaseUrl, c.ID))
}
