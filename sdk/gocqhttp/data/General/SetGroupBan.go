package General

// SetGroupBan Endpoint set_group_ban
type SetGroupBan struct {
	GroupId  int `json:"group_id"`
	UserId   int `json:"user_id"`
	Duration int `json:"duration"`
}
