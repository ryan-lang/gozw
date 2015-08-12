// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package language

import "encoding/binary"

// <no value>

type LanguageSet struct {
	Language uint32

	Country uint16
}

func ParseLanguageSet(payload []byte) LanguageSet {
	val := LanguageSet{}

	i := 2

	val.Language = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	val.Country = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return val
}