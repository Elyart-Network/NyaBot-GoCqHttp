package General

// SendPrivateMsg Endpoint send_private_msg
type SendPrivateMsg struct {
	UserId     int64  `json:"user_id"`
	GroupId    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
