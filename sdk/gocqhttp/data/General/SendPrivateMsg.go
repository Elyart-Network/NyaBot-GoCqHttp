package General

// SendPrivateMsg Endpoint send_private_msg
type SendPrivateMsg struct {
	UserId     int    `json:"user_id"`
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
