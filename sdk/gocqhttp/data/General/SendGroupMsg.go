package General

// sendGroupMsg Endpoint send_group_msg
type sendGroupMsg struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
