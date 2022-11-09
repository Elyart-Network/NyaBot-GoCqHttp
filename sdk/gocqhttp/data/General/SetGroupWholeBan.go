package General

// SetGroupWholeBanData Endpoint set_group_whole_ban
type SetGroupWholeBanData struct {
	GroupId int  `json:"group_id"`
	Enable  bool `json:"enable"`
}
