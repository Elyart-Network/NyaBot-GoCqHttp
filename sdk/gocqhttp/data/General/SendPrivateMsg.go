package General

// sendPrivateMsg Endpoint send_private_msg
type sendPrivateMsg struct {
	UserId     int    `json:"user_id"`
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
