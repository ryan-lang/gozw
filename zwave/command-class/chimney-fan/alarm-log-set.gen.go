// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

// <no value>

type ChimneyFanAlarmLogSet struct {
	Message byte
}

func ParseChimneyFanAlarmLogSet(payload []byte) ChimneyFanAlarmLogSet {
	val := ChimneyFanAlarmLogSet{}

	i := 2

	val.Message = payload[i]
	i++

	return val
}