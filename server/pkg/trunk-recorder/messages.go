package trunkrecorder

type MessageEnvelope struct {
	MessageType string `json:"type"`
	InstanceID  string `json:"instanceId"`
	InstanceKey string `json:"instanceKey"`
}
