package General

// sendGroupForwardMsg Endpoint send_group_forward_msg
type sendGroupForwardMsg struct {
	GroupId  int    `json:"group_id"`
	Messages string `json:"messages"`
}
