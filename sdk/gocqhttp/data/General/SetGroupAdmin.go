package General

// setGroupAdmin Endpoint set_group_admin
type setGroupAdmin struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	Enable  bool `json:"enable"`
}
