// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitorv2

// <no value>

type MeterTblTableIdReport struct {
	Properties1 struct {
		NumberOfCharacters byte
	}

	MeterIdCharacter []byte
}

func ParseMeterTblTableIdReport(payload []byte) MeterTblTableIdReport {
	val := MeterTblTableIdReport{}

	i := 2

	val.Properties1.NumberOfCharacters = (payload[i] & 0x1F)

	i += 1

	val.MeterIdCharacter = payload[i : i+0]
	i += 0

	return val
}
