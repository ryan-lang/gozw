// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementproxy

import "errors"

// <no value>

type NodeListGet struct {
	SeqNo byte
}

func (cmd *NodeListGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	return nil
}

func (cmd *NodeListGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	return
}