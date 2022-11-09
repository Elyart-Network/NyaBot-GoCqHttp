package General

// SendGroupForwardMsgData Endpoint send_group_forward_msg
type SendGroupForwardMsgData struct {
	GroupId  int    `json:"group_id"`
	Messages string `json:"messages"`
}
