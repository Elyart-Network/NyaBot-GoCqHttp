package General

// setGroupLeave Endpoint set_group_leave
type setGroupLeave struct {
	GroupId   int  `json:"group_id"`
	IsDismiss bool `json:"is_dismiss"`
}
