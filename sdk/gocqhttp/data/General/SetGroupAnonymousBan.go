package General

// SetGroupAnonymousBanData Endpoint set_group_anonymous_ban
type SetGroupAnonymousBanData struct {
	GroupId  int    `json:"group_id"`
	Flag     string `json:"flag"`
	Duration int    `json:"duration"`
}
