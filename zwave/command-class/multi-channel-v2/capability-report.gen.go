// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv2

// <no value>

type MultiChannelCapabilityReport struct {
	Properties1 struct {
		EndPoint byte

		Dynamic bool
	}

	GenericDeviceClass byte

	SpecificDeviceClass byte

	CommandClass []byte
}

func ParseMultiChannelCapabilityReport(payload []byte) MultiChannelCapabilityReport {
	val := MultiChannelCapabilityReport{}

	i := 2

	val.Properties1.EndPoint = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Dynamic = true
	} else {
		val.Properties1.Dynamic = false
	}

	i += 1

	val.GenericDeviceClass = payload[i]
	i++

	val.SpecificDeviceClass = payload[i]
	i++

	val.CommandClass = payload[i:]

	return val
}
