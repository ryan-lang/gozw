// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblconfig

import "encoding/binary"

// <no value>

type RateTblSet struct {
	RateParameterSetId byte

	Properties1 struct {
		NumberOfRateChar byte

		RateType byte
	}

	RateCharacter []byte

	StartHourLocalTime byte

	StartMinuteLocalTime byte

	DurationMinute uint16

	Properties2 struct {
		ConsumptionScale byte

		ConsumptionPrecision byte
	}

	MinConsumptionValue uint32

	MaxConsumptionValue uint32

	Properties3 struct {
		MaxDemandScale byte

		MaxDemandPrecision byte
	}

	MaxDemandValue uint32

	DcpRateId byte
}

func ParseRateTblSet(payload []byte) RateTblSet {
	val := RateTblSet{}

	i := 2

	val.RateParameterSetId = payload[i]
	i++

	val.Properties1.NumberOfRateChar = (payload[i] & 0x1F)

	val.Properties1.RateType = (payload[i] & 0x60) << 5

	i += 1

	val.RateCharacter = payload[i : i+1]
	i += 1

	val.StartHourLocalTime = payload[i]
	i++

	val.StartMinuteLocalTime = payload[i]
	i++

	val.DurationMinute = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.Properties2.ConsumptionScale = (payload[i] & 0x1F)

	val.Properties2.ConsumptionPrecision = (payload[i] & 0xE0) << 5

	i += 1

	val.MinConsumptionValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	val.MaxConsumptionValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	val.Properties3.MaxDemandScale = (payload[i] & 0x1F)

	val.Properties3.MaxDemandPrecision = (payload[i] & 0xE0) << 5

	i += 1

	val.MaxDemandValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	val.DcpRateId = payload[i]
	i++

	return val
}
