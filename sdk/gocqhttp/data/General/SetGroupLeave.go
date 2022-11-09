package General

// SetGroupLeaveData Endpoint set_group_leave
type SetGroupLeaveData struct {
	GroupId   int  `json:"group_id"`
	IsDismiss bool `json:"is_dismiss"`
}
