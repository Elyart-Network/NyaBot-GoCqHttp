package General

// SetGroupWholeBan Endpoint set_group_whole_ban
type SetGroupWholeBan struct {
	GroupId int  `json:"group_id"`
	Enable  bool `json:"enable"`
}
