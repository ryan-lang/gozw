// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

// <no value>

type SecurityMessageEncapsulation struct {
	InitializationVectorByte []byte

	Properties1 struct {
		SequenceCounter byte

		Sequenced bool

		SecondFrame bool
	}

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte

	ReceiversNonceIdentifier byte

	MessageAuthenticationCodeByte []byte
}

func ParseSecurityMessageEncapsulation(payload []byte) SecurityMessageEncapsulation {
	val := SecurityMessageEncapsulation{}

	i := 2

	val.InitializationVectorByte = payload[i : i+8]

	i += 8

	val.Properties1.SequenceCounter = (payload[i] & 0x0F)

	if payload[i]&0x10 == 0x10 {
		val.Properties1.Sequenced = true
	} else {
		val.Properties1.Sequenced = false
	}

	if payload[i]&0x20 == 0x20 {
		val.Properties1.SecondFrame = true
	} else {
		val.Properties1.SecondFrame = false
	}

	i += 1

	val.CommandClassIdentifier = payload[i]
	i++

	val.CommandIdentifier = payload[i]
	i++

	val.CommandByte = payload[i:]

	val.ReceiversNonceIdentifier = payload[i]
	i++

	val.MessageAuthenticationCodeByte = payload[i : i+8]

	i += 8

	return val
}
