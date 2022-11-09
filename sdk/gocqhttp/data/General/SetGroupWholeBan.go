package General

// setGroupWholeBan Endpoint set_group_whole_ban
type setGroupWholeBan struct {
	GroupId int  `json:"group_id"`
	Enable  bool `json:"enable"`
}
