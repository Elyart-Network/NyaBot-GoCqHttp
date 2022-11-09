package General

// SetGroupSpacialTitleData Endpoint set_group_special_title
type SetGroupSpacialTitleData struct {
	GroupId      int    `json:"group_id"`
	UserId       int    `json:"user_id"`
	SpacialTitle string `json:"special_title"`
	Duration     int    `json:"duration"`
}
