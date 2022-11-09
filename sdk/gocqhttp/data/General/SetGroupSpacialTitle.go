package General

// setGroupSpacialTitle Endpoint set_group_special_title
type setGroupSpacialTitle struct {
	GroupId      int    `json:"group_id"`
	UserId       int    `json:"user_id"`
	SpacialTitle string `json:"special_title"`
	Duration     int    `json:"duration"`
}
