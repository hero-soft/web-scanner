package trunkrecorder

type MessageEnvelope struct {
	MessageType string `json:"type"`
	InstanceID  string `json:"instanceId"`
	InstanceKey string `json:"instanceKey"`
}

type CallsActive struct {
	MessageType string `json:"type"`
	InstanceID  string `json:"instanceId"`
	InstanceKey string `json:"instanceKey"`
	Calls       []Call `json:"calls"`
}

type Call struct {
	ID           string `json:"id"`
	Analog       string `json:"analog"`
	Conventional string `json:"conventional"`
	Elasped      string `json:"elasped"`
	Emergency    string `json:"emergency"`
	Encrypted    string `json:"encrypted"`
	Freq         string `json:"freq"`
	Length       string `json:"length"`
	Phase2       string `json:"phase_2"`
	RecNum       string `json:"rec_num"`
	RecState     string `json:"rec_state"`
	ShortName    string `json:"short_name"`
	SrcId        string `json:"src_id"`
	SrcNum       string `json:"src_num"`
	StartTime    string `json:"start_time"`
	State        string `json:"state"`
	StopTime     string `json:"stop_time"`
	SysNum       string `json:"sys_num"`
	Talkgroup    string `json:"talkgroup"`
	TalkgroupTag string `json:"talkgrouptag"`
}

type AudioUpload struct {
	APIKey     string `json:"api_key,omitempty"`
	CallLength int    `json:"call_length,omitempty"`
	Emergency  bool   `json:"emergency,omitempty"`
}

// api_key:[1234]
// call_length:[19]
// emergency:[0]
// freq:[855162500]
// freq_list:[[{"freq": 855162500, "time": 1667753976, "pos": 0.00, "len": 18.95, "error_count": "0", "spike_count": "0"}]]
// source_list:[[{ "pos": 0.00, "src": 32199 }]]
// start_time:[1667753976]
// stop_time:[1667753995]
// talkgroup_num:[36816]]
// audio/36816-1667753976_855162500-call_16452.m4a
