// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

// <no value>

type NodeAddStatus struct {
	SeqNo byte

	Status byte

	NewNodeId byte

	NodeInfoLength byte

	Properties1 struct {
		Capability byte

		Listening bool
	}

	Properties2 struct {
		Security byte

		Opt bool
	}

	BasicDeviceClass byte

	GenericDeviceClass byte

	SpecificDeviceClass byte

	CommandClass []byte
}

func ParseNodeAddStatus(payload []byte) NodeAddStatus {
	val := NodeAddStatus{}

	i := 2

	val.SeqNo = payload[i]
	i++

	val.Status = payload[i]
	i++

	val.NewNodeId = payload[i]
	i++

	val.NodeInfoLength = payload[i]
	i++

	val.Properties1.Capability = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Listening = true
	} else {
		val.Properties1.Listening = false
	}

	i += 1

	val.Properties2.Security = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Properties2.Opt = true
	} else {
		val.Properties2.Opt = false
	}

	i += 1

	val.BasicDeviceClass = payload[i]
	i++

	val.GenericDeviceClass = payload[i]
	i++

	val.SpecificDeviceClass = payload[i]
	i++

	val.CommandClass = payload[i:]

	return val
}
