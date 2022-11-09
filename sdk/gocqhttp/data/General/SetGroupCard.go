package General

// SetGroupCard Endpoint set_group_card
type SetGroupCard struct {
	GroupId int    `json:"group_id"`
	UserId  int    `json:"user_id"`
	Card    string `json:"card"`
}
