// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

// <no value>

type ProtectionReport struct {
	Level struct {
		LocalProtectionState byte
	}

	Level2 struct {
		RfProtectionState byte
	}
}

func ParseProtectionReport(payload []byte) ProtectionReport {
	val := ProtectionReport{}

	i := 2

	val.Level.LocalProtectionState = (payload[i] & 0x0F)

	i += 1

	val.Level2.RfProtectionState = (payload[i] & 0x0F)

	i += 1

	return val
}
