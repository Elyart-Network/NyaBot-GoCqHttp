package General

// setGroupAnonymousBan Endpoint set_group_anonymous_ban
type setGroupAnonymousBan struct {
	GroupId  int    `json:"group_id"`
	Flag     string `json:"flag"`
	Duration int    `json:"duration"`
}
