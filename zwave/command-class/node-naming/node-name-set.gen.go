// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package nodenaming

// <no value>

type NodeNamingNodeNameSet struct {
	Level struct {
		CharPresentation byte
	}

	NodeNameChar string
}

func ParseNodeNamingNodeNameSet(payload []byte) NodeNamingNodeNameSet {
	val := NodeNamingNodeNameSet{}

	i := 2

	val.Level.CharPresentation = (payload[i] & 0x07)

	i += 1

	val.NodeNameChar = string(payload[i : i+16])

	i += 16

	return val
}
