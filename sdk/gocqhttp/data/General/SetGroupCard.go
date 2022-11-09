package General

// SetGroupCardData Endpoint set_group_card
type SetGroupCardData struct {
	GroupId int    `json:"group_id"`
	UserId  int    `json:"user_id"`
	Card    string `json:"card"`
}
