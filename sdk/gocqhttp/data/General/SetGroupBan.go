package General

// SetGroupBanData Endpoint set_group_ban
type SetGroupBanData struct {
	GroupId  int `json:"group_id"`
	UserId   int `json:"user_id"`
	Duration int `json:"duration"`
}
