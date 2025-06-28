package commands

import (
	"fmt"

	"github.com/micromdm/plist"
)

// ResponseStatus is the status of a device's response to a command
type ResponseStatus string

const (
	// The device processed the command successfully.
	ResponseStatusAcknowledged ResponseStatus = "Acknowledged"
	// An error occurred. See the ErrorChain for more details.
	ResponseStatusError ResponseStatus = "Error"
	// A protocol error occurred, which can result from a malformed command.
	ResponseStatusCommandFormatError ResponseStatus = "CommandFormatError"
	// The device is idle; there's no status.
	ResponseStatusIdle ResponseStatus = "Idle"
	// The device received the command, but couldn’t execute it.
	ResponseStatusNotNow ResponseStatus = "NotNow"
)

type ResponsePayload interface {
	ResponseRequestType() string
}

type ResponseError struct {
	// The error code.
	ErrorCode int `plist:"ErrorCode"`
	// The error domain.
	ErrorDomain string `plist:"ErrorDomain"`
	// A description of the error in the device’s localized language.
	LocalizedDescription string `plist:"LocalizedDescription"`
	// A description of the error in U.S. English.
	USEnglishDescription string `plist:"USEnglishDescription"`
}

type Response[T ResponsePayload] struct {
	// The unique identifier of the command for this response.
	CommandUUID string `plist:"CommandUUID"`
	// If present, an array of dictionaries that describes any errors that occurred.
	// The first item represents the top-level error.
	// Subsequent items represent the underlying errors that led up to the top-level error.
	ErrorChain []*ResponseError `plist:"ErrorChain,omitempty"`
	// The status of the response.
	Status ResponseStatus `plist:"Status"`
	// The unique identifier of the device.
	UDID string `plist:"UDID"`

	// The response payload
	ResponsePayload T `plist:"-"`
}

func UnmarshalResponse[T ResponsePayload](buf []byte) (*Response[T], error) {
	resp := new(Response[T])
	if err := plist.Unmarshal(buf, resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	rp := new(T)
	if err := plist.Unmarshal(buf, rp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response payload: %w", err)
	}
	resp.ResponsePayload = *rp

	return resp, nil
}
