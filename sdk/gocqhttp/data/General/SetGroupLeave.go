package General

// SetGroupLeave Endpoint set_group_leave
type SetGroupLeave struct {
	GroupId   int  `json:"group_id"`
	IsDismiss bool `json:"is_dismiss"`
}
