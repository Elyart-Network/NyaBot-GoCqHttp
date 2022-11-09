package General

// SendGroupMsgData Endpoint send_group_msg
type SendGroupMsgData struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
