// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatfanmodev4

// <no value>

type ThermostatFanModeReport struct {
	Properties1 struct {
		Off bool

		FanMode byte
	}
}

func ParseThermostatFanModeReport(payload []byte) ThermostatFanModeReport {
	val := ThermostatFanModeReport{}

	i := 2

	val.Properties1.FanMode = (payload[i] & 0x0F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Off = true
	} else {
		val.Properties1.Off = false
	}

	i += 1

	return val
}
