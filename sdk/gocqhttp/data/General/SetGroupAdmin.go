package General

// SetGroupAdminData Endpoint set_group_admin
type SetGroupAdminData struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	Enable  bool `json:"enable"`
}
