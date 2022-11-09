package General

// SendGroupForwardMsg Endpoint send_group_forward_msg
type SendGroupForwardMsg struct {
	GroupId  int    `json:"group_id"`
	Messages string `json:"messages"`
}
