package General

// SetGroupAnonymousBan Endpoint set_group_anonymous_ban
type SetGroupAnonymousBan struct {
	GroupId  int    `json:"group_id"`
	Flag     string `json:"flag"`
	Duration int    `json:"duration"`
}
