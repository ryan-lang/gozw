// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationcommandconfiguration

// <no value>

type CommandConfigurationReport struct {
	GroupingIdentifier byte

	NodeId byte

	Properties1 struct {
		ReportsToFollow byte

		First bool
	}

	CommandLength byte

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte
}

func ParseCommandConfigurationReport(payload []byte) CommandConfigurationReport {
	val := CommandConfigurationReport{}

	i := 2

	val.GroupingIdentifier = payload[i]
	i++

	val.NodeId = payload[i]
	i++

	val.Properties1.ReportsToFollow = (payload[i] & 0x0F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.First = true
	} else {
		val.Properties1.First = false
	}

	i += 1

	val.CommandLength = payload[i]
	i++

	val.CommandClassIdentifier = payload[i]
	i++

	val.CommandIdentifier = payload[i]
	i++

	val.CommandByte = payload[i:]

	return val
}
