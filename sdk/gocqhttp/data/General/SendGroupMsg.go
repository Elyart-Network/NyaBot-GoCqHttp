package General

// SendGroupMsg Endpoint send_group_msg
type SendGroupMsg struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
