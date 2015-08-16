// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv4

// <no value>

type NotificationSupportedReport struct {
	Properties1 struct {
		NumberOfBitMasks byte

		V1Alarm bool
	}

	BitMask byte
}

func ParseNotificationSupportedReport(payload []byte) NotificationSupportedReport {
	val := NotificationSupportedReport{}

	i := 2

	val.Properties1.NumberOfBitMasks = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.V1Alarm = true
	} else {
		val.Properties1.V1Alarm = false
	}

	i += 1

	val.BitMask = payload[i]
	i++

	return val
}
