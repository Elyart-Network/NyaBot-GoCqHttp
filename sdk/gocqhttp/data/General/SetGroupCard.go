package General

// setGroupCard Endpoint set_group_card
type setGroupCard struct {
	GroupId int    `json:"group_id"`
	UserId  int    `json:"user_id"`
	Card    string `json:"card"`
}
