// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package prepayment

// <no value>

type PrepaymentBalanceGet struct {
	Properties1 struct {
		BalanceType byte
	}
}

func ParsePrepaymentBalanceGet(payload []byte) PrepaymentBalanceGet {
	val := PrepaymentBalanceGet{}

	i := 2

	val.Properties1.BalanceType = (payload[i] & 0xC0) << 6

	i += 1

	return val
}
