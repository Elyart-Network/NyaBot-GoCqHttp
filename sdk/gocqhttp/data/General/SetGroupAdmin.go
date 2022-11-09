package General

// SetGroupAdmin Endpoint set_group_admin
type SetGroupAdmin struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	Enable  bool `json:"enable"`
}
