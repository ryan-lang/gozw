// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatfanmodev2

// <no value>

type ThermostatFanModeReport struct {
	Level struct {
		FanMode byte
	}
}

func ParseThermostatFanModeReport(payload []byte) ThermostatFanModeReport {
	val := ThermostatFanModeReport{}

	i := 2

	val.Level.FanMode = (payload[i] & 0x0F)

	i += 1

	return val
}
