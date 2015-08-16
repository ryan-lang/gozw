// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv2

// <no value>

type MultiChannelEndPointReport struct {
	Properties1 struct {
		Identical bool

		Dynamic bool
	}

	Properties2 struct {
		EndPoints byte
	}
}

func ParseMultiChannelEndPointReport(payload []byte) MultiChannelEndPointReport {
	val := MultiChannelEndPointReport{}

	i := 2

	if payload[i]&0x40 == 0x40 {
		val.Properties1.Identical = true
	} else {
		val.Properties1.Identical = false
	}

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Dynamic = true
	} else {
		val.Properties1.Dynamic = false
	}

	i += 1

	val.Properties2.EndPoints = (payload[i] & 0x7F)

	i += 1

	return val
}
