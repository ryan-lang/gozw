// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchmultilevelv4

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandStopLevelChange cc.CommandID = 0x05

func init() {
	gob.Register(StopLevelChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x26),
		Command:      cc.CommandID(0x05),
		Version:      4,
	}, NewStopLevelChange)
}

func NewStopLevelChange() cc.Command {
	return &StopLevelChange{}
}

// <no value>
type StopLevelChange struct {
}

func (cmd StopLevelChange) CommandClassID() cc.CommandClassID {
	return 0x26
}

func (cmd StopLevelChange) CommandID() cc.CommandID {
	return CommandStopLevelChange
}

func (cmd StopLevelChange) CommandIDString() string {
	return "SWITCH_MULTILEVEL_STOP_LEVEL_CHANGE"
}

func (cmd *StopLevelChange) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *StopLevelChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}