// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev2

// <no value>

type ThermostatModeSet struct {
	Level struct {
		Mode byte
	}
}

func ParseThermostatModeSet(payload []byte) ThermostatModeSet {
	val := ThermostatModeSet{}

	i := 2

	val.Level.Mode = (payload[i] & 0x1F)

	i += 1

	return val
}
