// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

// <no value>

type ChimneyFanSetupSet struct {
	Mode byte

	BoostTime byte

	StopTime byte

	MinSpeed byte

	Properties1 struct {
		Size1 byte

		Scale1 byte

		Precision1 byte
	}

	StartTemperature []byte

	Properties2 struct {
		Size2 byte

		Scale2 byte

		Precision2 byte
	}

	StopTemperature []byte

	Properties3 struct {
		Size3 byte

		Scale3 byte

		Precision3 byte
	}

	AlarmTemperatureValue []byte
}

func ParseChimneyFanSetupSet(payload []byte) ChimneyFanSetupSet {
	val := ChimneyFanSetupSet{}

	i := 2

	val.Mode = payload[i]
	i++

	val.BoostTime = payload[i]
	i++

	val.StopTime = payload[i]
	i++

	val.MinSpeed = payload[i]
	i++

	val.Properties1.Size1 = (payload[i] & 0x07)

	val.Properties1.Scale1 = (payload[i] & 0x18) << 3

	val.Properties1.Precision1 = (payload[i] & 0xE0) << 5

	i += 1

	val.StartTemperature = payload[i : i+4]
	i += 4

	val.Properties2.Size2 = (payload[i] & 0x07)

	val.Properties2.Scale2 = (payload[i] & 0x18) << 3

	val.Properties2.Precision2 = (payload[i] & 0xE0) << 5

	i += 1

	val.StopTemperature = payload[i : i+6]
	i += 6

	val.Properties3.Size3 = (payload[i] & 0x07)

	val.Properties3.Scale3 = (payload[i] & 0x18) << 3

	val.Properties3.Precision3 = (payload[i] & 0xE0) << 5

	i += 1

	val.AlarmTemperatureValue = payload[i : i+8]
	i += 8

	return val
}
