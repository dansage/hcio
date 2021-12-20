package hcio_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dansage/hcio"
)

var id string

// TestMain initializes the state by loading the check ID to use for each test
func TestMain(m *testing.M) {
	var ok bool

	// attempt to find the ID from the environment
	if id, ok = os.LookupEnv("CHECK_ID"); !ok || id == "" {
		// a valid check ID must be specified for the tests to run
		panic(fmt.Errorf("a valid check ID was not specified in env:CHECK_ID"))
	}

	// start the tests
	m.Run()
}

// TestNewCheck verifies that creating a check passes the ID correctly
func TestNewCheck(t *testing.T) {
	// create a check using the ID from the environment
	check := hcio.NewCheck(id)

	// verify the ID was set correctly
	if check.ID != id {
		t.Errorf("check ID not set to user specified value, got: %q, want: %q", check.ID, id)
	}
}

// TestPingCheck verifies that pinging a check to indicate success operates correctly
func TestPingCheck(t *testing.T) {
	// create a check using the ID from the environment
	check := hcio.NewCheck(id, hcio.Options{
		UserAgent: "HCio/TestPingCheck",
	})

	// send a success ping for the check
	err := check.Ping()
	if err != nil {
		t.Errorf("failed to ping check: %v", err)
	}
}

// TestPingFail verifies that pinging a check to indicate failure operates correctly
func TestPingFail(t *testing.T) {
	// create a check using the ID from the environment
	check := hcio.NewCheck(id, hcio.Options{
		UserAgent: "HCio/TestPingFail",
	})

	// send a failure ping for the check
	err := check.Fail()
	if err != nil {
		t.Errorf("failed to fail check: %v", err)
	}
}

// TestPingFailCode verifies that pinging a check to indicate failure code operates correctly
func TestPingFailCode(t *testing.T) {
	// create a check using the ID from the environment
	check := hcio.NewCheck(id, hcio.Options{
		UserAgent: "HCio/TestPingFailCode",
	})

	// send a failure ping for the check
	err := check.FailCode(3)
	if err != nil {
		t.Errorf("failed to fail check with code: %v", err)
	}
}

// TestPingTimed verifies that pinging a check to indicate process duration operates correctly
func TestPingTimed(t *testing.T) {
	// create a check using the ID from the environment
	check := hcio.NewCheck(id, hcio.Options{
		UserAgent: "HCio/TestPingTimed",
	})

	// send a start ping for the check
	err := check.Start()
	if err != nil {
		t.Errorf("failed to start check: %v", err)
	}

	// wait for 10 seconds
	time.Sleep(10 * time.Second)

	// send a success ping for the check
	err = check.Success()
	if err != nil {
		t.Errorf("failed to ping check: %v", err)
	}
}
