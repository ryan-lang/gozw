// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

// <no value>

type MultiChannelCmdEncap struct {
	Properties1 struct {
		SourceEndPoint byte
	}

	Properties2 struct {
		DestinationEndPoint byte

		BitAddress bool
	}

	CommandClass byte

	Command byte

	Parameter []byte
}

func ParseMultiChannelCmdEncap(payload []byte) MultiChannelCmdEncap {
	val := MultiChannelCmdEncap{}

	i := 2

	val.Properties1.SourceEndPoint = (payload[i] & 0x7F)

	i += 1

	val.Properties2.DestinationEndPoint = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Properties2.BitAddress = true
	} else {
		val.Properties2.BitAddress = false
	}

	i += 1

	val.CommandClass = payload[i]
	i++

	val.Command = payload[i]
	i++

	val.Parameter = payload[i:]

	return val
}
