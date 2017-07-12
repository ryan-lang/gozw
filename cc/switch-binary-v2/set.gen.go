// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchbinaryv2

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x25),
		Command:      cc.CommandID(0x01),
		Version:      2,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	TargetValue byte

	Duration byte
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x25
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "SWITCH_BINARY_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.TargetValue) %d<=%d", len(payload), i)
	}

	cmd.TargetValue = payload[i]
	i++

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.Duration) %d<=%d", len(payload), i)
	}

	cmd.Duration = payload[i]
	i++

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.TargetValue)

	payload = append(payload, cmd.Duration)

	return
}