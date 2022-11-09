package General

// sendMsg Endpoint send_msg
type sendMsg struct {
	MessageType string `json:"message_type"`
	UserId      int    `json:"user_id"`
	GroupId     int    `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  bool   `json:"auto_escape"`
}
