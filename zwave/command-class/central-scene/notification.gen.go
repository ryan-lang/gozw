// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package centralscene

// <no value>

type CentralSceneNotification struct {
	SequenceNumber byte

	KeyAttributes byte

	SceneNumber byte
}

func ParseCentralSceneNotification(payload []byte) CentralSceneNotification {
	val := CentralSceneNotification{}

	i := 2

	val.SequenceNumber = payload[i]
	i++

	val.KeyAttributes = (payload[i] & 0x07)

	i += 1

	val.SceneNumber = payload[i]
	i++

	return val
}