// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package prepayment

// <no value>

type PrepaymentSupportedReport struct {
	Properties1 struct {
		TypesSupported byte
	}
}

func ParsePrepaymentSupportedReport(payload []byte) PrepaymentSupportedReport {
	val := PrepaymentSupportedReport{}

	i := 2

	val.Properties1.TypesSupported = (payload[i] & 0x0F)

	i += 1

	return val
}
