package data

// SendPrivateMsgData Endpoint send_private_msg
type SendPrivateMsgData struct {
	UserId     int    `json:"user_id"`
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

// SendGroupMsgData Endpoint send_group_msg
type SendGroupMsgData struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

// SendGroupForwardMsgData Endpoint send_group_forward_msg
type SendGroupForwardMsgData struct {
	GroupId  int    `json:"group_id"`
	Messages string `json:"messages"`
}

// SendMsgData Endpoint send_msg
type SendMsgData struct {
	MessageType string `json:"message_type"`
	UserId      int    `json:"user_id"`
	GroupId     int    `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  bool   `json:"auto_escape"`
}

// DeleteMsgData Endpoint delete_msg
type DeleteMsgData struct {
	MessageId int `json:"message_id"`
}

// GetMsgData Endpoint get_msg
type GetMsgData struct {
	MessageId int `json:"message_id"`
}

// GetForwardMsgData Endpoint get_forward_msg
type GetForwardMsgData struct {
	MessageID string `json:"message_id"`
}

// GetImageData Endpoint get_image
type GetImageData struct {
	File string `json:"file"`
}

// MarkMsgAsReadData Endpoint mark_msg_as_read
type MarkMsgAsReadData struct {
	MessageId int `json:"message_id"`
}

// SetGroupKickData Endpoint set_group_kick
type SetGroupKickData struct {
	GroupId          int  `json:"group_id"`
	UserId           int  `json:"user_id"`
	RejectAddRequest bool `json:"reject_add_request"`
}

// SetGroupBanData Endpoint set_group_ban
type SetGroupBanData struct {
	GroupId  int `json:"group_id"`
	UserId   int `json:"user_id"`
	Duration int `json:"duration"`
}

// SetGroupAnonymousBanData Endpoint set_group_anonymous_ban
type SetGroupAnonymousBanData struct {
	GroupId  int    `json:"group_id"`
	Flag     string `json:"flag"`
	Duration int    `json:"duration"`
}

// SetGroupWholeBanData Endpoint set_group_whole_ban
type SetGroupWholeBanData struct {
	GroupId int  `json:"group_id"`
	Enable  bool `json:"enable"`
}

// SetGroupAdminData Endpoint set_group_admin
type SetGroupAdminData struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	Enable  bool `json:"enable"`
}

// SetGroupCardData Endpoint set_group_card
type SetGroupCardData struct {
	GroupId int    `json:"group_id"`
	UserId  int    `json:"user_id"`
	Card    string `json:"card"`
}

// SetGroupNameData Endpoint set_group_name
type SetGroupNameData struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}

// SetGroupLeaveData Endpoint set_group_leave
type SetGroupLeaveData struct {
	GroupId   int  `json:"group_id"`
	IsDismiss bool `json:"is_dismiss"`
}

// SetGroupSpacialTitleData Endpoint set_group_special_title
type SetGroupSpacialTitleData struct {
	GroupId      int    `json:"group_id"`
	UserId       int    `json:"user_id"`
	SpacialTitle string `json:"special_title"`
	Duration     int    `json:"duration"`
}

// SetGroupSignData Endpoint set_group_sign
type SetGroupSignData struct {
	GroupId int `json:"group_id"`
}
