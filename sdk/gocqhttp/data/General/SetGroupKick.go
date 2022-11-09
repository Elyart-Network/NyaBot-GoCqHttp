package General

// SetGroupKick Endpoint set_group_kick
type SetGroupKick struct {
	GroupId          int  `json:"group_id"`
	UserId           int  `json:"user_id"`
	RejectAddRequest bool `json:"reject_add_request"`
}
