// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

// <no value>

type MultiChannelEndPointFind struct {
	GenericDeviceClass byte

	SpecificDeviceClass byte
}

func ParseMultiChannelEndPointFind(payload []byte) MultiChannelEndPointFind {
	val := MultiChannelEndPointFind{}

	i := 2

	val.GenericDeviceClass = payload[i]
	i++

	val.SpecificDeviceClass = payload[i]
	i++

	return val
}