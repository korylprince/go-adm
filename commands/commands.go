//go:generate cmdgen -repo "https://github.com/apple/device-management.git" -commit "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

package commands

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/korylprince/go-adm/commands/commands"
	"github.com/korylprince/go-adm/tagutil"
	"github.com/micromdm/plist"
)

var ErrUnknownRequestType = errors.New("unknown request type")

type CommandPayload interface {
	RequestType() string
}

type Command struct {
	// The unique identifier of the command
	CommandUUID string

	// If true, the device needs to be network-tethered to run the command.
	RequestRequiresNetworkTether bool

	// The command dictionary
	Command CommandPayload
}

func (c *Command) Type() string {
	if c == nil || c.Command == nil {
		return ""
	}
	return c.Command.RequestType()
}

type command struct {
	CommandUUID string `plist:"CommandUUID"`
	Command     any    `plist:"Command"`
}

// PlistValue returns a composite type suitable for marshalling to plist
func (c *Command) PlistValue() (any, error) {
	// verify Payload is correct type
	if c.Command == nil {
		return nil, errors.New("Payload is nil")
	}
	typ := reflect.TypeOf(c.Command)
	// deref pointer
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("Payload is not a struct or pointer to a struct")
	}

	// fill struct defaults
	if err := tagutil.SetDefaults(c.Command); err != nil {
		return nil, fmt.Errorf("could not set struct defaults: %w", err)
	}

	cmd := &command{CommandUUID: c.CommandUUID}

	// create a new struct with the RequestType field
	flds := []reflect.StructField{{
		Name: "RequestType",
		Type: reflect.TypeFor[string](),
		Tag:  reflect.StructTag(`plist:"RequestType"`),
	}, {
		Name: "RequestRequiresNetworkTether",
		Type: reflect.TypeFor[bool](),
		Tag:  reflect.StructTag(`plist:"RequestRequiresNetworkTether,omitempty"`),
	}}

	// append the existing fields
	for idx := 0; idx < typ.NumField(); idx++ {
		fld := typ.Field(idx)
		if fld.Name != "RequestType" && fld.Name != "RequestRequiresNetworkTether" {
			flds = append(flds, fld)
		}
	}

	// create a new struct and copy the values
	payload := reflect.New(reflect.StructOf(flds))
	payload.Elem().Field(0).Set(reflect.ValueOf(c.Type()))

	old := reflect.ValueOf(c.Command).Elem()
	for idx := 0; idx < old.NumField(); idx++ {
		payload.Elem().Field(idx + 2).Set(old.Field(idx))
	}

	cmd.Command = payload.Elem().Interface()

	return cmd, nil
}

func (c *Command) MarshalIndentPlist(indent string) ([]byte, error) {
	cmd, err := c.PlistValue()
	if err != nil {
		return nil, err
	}
	return plist.MarshalIndent(cmd, indent)
}

func (c *Command) MarshalPlist(indent string) ([]byte, error) {
	return c.MarshalIndentPlist("")
}

func NewFromType(typ, cmdUUID string) (*Command, error) {
	payload, ok := commands.CommandMap[typ]
	if !ok {
		return nil, ErrUnknownRequestType
	}

	pay := reflect.New(reflect.TypeOf(payload)).Interface().(CommandPayload)

	return &Command{
		CommandUUID: cmdUUID,
		Command:     pay,
	}, nil
}
