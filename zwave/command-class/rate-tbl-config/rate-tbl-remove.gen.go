// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblconfig

// <no value>

type RateTblRemove struct {
	Properties1 struct {
		RateParameterSetIds byte
	}

	RateParameterSetId []byte
}

func ParseRateTblRemove(payload []byte) RateTblRemove {
	val := RateTblRemove{}

	i := 2

	val.Properties1.RateParameterSetIds = (payload[i] & 0x3F)

	i += 1

	val.RateParameterSetId = payload[i : i+0]
	i += 0

	return val
}
