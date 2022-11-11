package funcs

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"github.com/gorilla/websocket"
)

// SendPrivateMsg Endpoint send_private_msg
func SendPrivateMsg(c *websocket.Conn, UserId int, GroupId int, Message string, AutoEscape bool) {
	Data := data.SendPrivateMsgData{
		UserId:     UserId,
		GroupId:    GroupId,
		Message:    Message,
		AutoEscape: AutoEscape,
	}
	WsSend(c, "send_private_msg", Data)
}

// SendGroupMsg Endpoint send_group_msg
func SendGroupMsg(c *websocket.Conn, GroupId int, Message string, AutoEscape bool) {
	Data := data.SendGroupMsgData{
		GroupId:    GroupId,
		Message:    Message,
		AutoEscape: AutoEscape,
	}
	WsSend(c, "send_group_msg", Data)
}

// SendGroupForwardMsg Endpoint send_group_forward_msg
func SendGroupForwardMsg(c *websocket.Conn, GroupID int, Messages string) {
	Data := data.SendGroupForwardMsgData{
		GroupId:  GroupID,
		Messages: Messages,
	}
	WsSend(c, "send_group_forward_msg", Data)
}

// SendMsg Endpoint send_msg
func SendMsg(c *websocket.Conn, MessageType string, UserID int, Message string, GroupID int, AutoEscape bool) {
	Data := data.SendMsgData{
		MessageType: MessageType,
		UserId:      UserID,
		GroupId:     GroupID,
		Message:     Message,
		AutoEscape:  AutoEscape,
	}
	WsSend(c, "send_msg", Data)
}

// DeleteMsg Endpoint delete_msg
func DeleteMsg(c *websocket.Conn, MessageID int) {
	Data := data.DeleteMsgData{
		MessageId: MessageID,
	}
	WsSend(c, "delete_msg", Data)
}

// GetMsg Endpoint get_msg
func GetMsg(c *websocket.Conn, MessageID int) {
	Data := data.GetMsgData{
		MessageId: MessageID,
	}
	WsSend(c, "get_msg", Data)
}

// GetForwardMsg Endpoint get_forward_msg
func GetForwardMsg(c *websocket.Conn, MessageID string) {
	Data := data.GetForwardMsgData{
		MessageID: MessageID,
	}
	WsSend(c, "get_forward_msg", Data)
}

// GetImage Endpoint get_image
func GetImage(c *websocket.Conn, File string) {
	Data := data.GetImageData{
		File: File,
	}
	WsSend(c, "get_image", Data)
}

// MarkMsgAsRead Endpoint mark_msg_as_read
func MarkMsgAsRead(c *websocket.Conn, MessageID int) {
	Data := data.MarkMsgAsReadData{
		MessageId: MessageID,
	}
	WsSend(c, "mark_msg_as_read", Data)
}

// SetGroupKick Endpoint set_group_kick
func SetGroupKick(c *websocket.Conn, GroupID int, UserID int, RejectAddRequest bool) {
	Data := data.SetGroupKickData{
		GroupId:          GroupID,
		UserId:           UserID,
		RejectAddRequest: RejectAddRequest,
	}
	WsSend(c, "set_group_kick", Data)
}

// SetGroupBan Endpoint set_group_ban
func SetGroupBan(c *websocket.Conn, GroupID int, UserID int, Duration int) {
	Data := data.SetGroupBanData{
		GroupId:  GroupID,
		UserId:   UserID,
		Duration: Duration,
	}
	WsSend(c, "set_group_ban", Data)
}

// SetGroupAnonymousBan Endpoint set_group_anonymous_ban
func SetGroupAnonymousBan(c *websocket.Conn, GroupID int, Flag string, Duration int) {
	Data := data.SetGroupAnonymousBanData{
		GroupId:  GroupID,
		Flag:     Flag,
		Duration: Duration,
	}
	WsSend(c, "set_group_anonymous_ban", Data)
}

// SetGroupWholeBan Endpoint set_group_whole_ban
func SetGroupWholeBan(c *websocket.Conn, GroupID int, Enable bool) {
	Data := data.SetGroupWholeBanData{
		GroupId: GroupID,
		Enable:  Enable,
	}
	WsSend(c, "set_group_whole_ban", Data)
}

// SetGroupAdmin Endpoint set_group_admin
func SetGroupAdmin(c *websocket.Conn, GroupID int, UserID int, Enable bool) {
	Data := data.SetGroupAdminData{
		GroupId: GroupID,
		UserId:  UserID,
		Enable:  Enable,
	}
	WsSend(c, "set_group_admin", Data)
}

// SetGroupCard Endpoint set_group_card
func SetGroupCard(c *websocket.Conn, GroupID int, UserID int, Card string) {
	Data := data.SetGroupCardData{
		GroupId: GroupID,
		UserId:  UserID,
		Card:    Card,
	}
	WsSend(c, "set_group_card", Data)
}

// SetGroupName Endpoint set_group_name
func SetGroupName(c *websocket.Conn, GroupID int, GroupName string) {
	Data := data.SetGroupNameData{
		GroupId:   GroupID,
		GroupName: GroupName,
	}
	WsSend(c, "set_group_name", Data)
}

// SetGroupLeave Endpoint set_group_leave
func SetGroupLeave(c *websocket.Conn, GroupID int, IsDismiss bool) {
	Data := data.SetGroupLeaveData{
		GroupId:   GroupID,
		IsDismiss: IsDismiss,
	}
	WsSend(c, "set_group_leave", Data)
}

// SetGroupSpecialTitle Endpoint set_group_spacial_title
func SetGroupSpecialTitle(c *websocket.Conn, GroupID int, UserID int, SpacialTitle string, Duration int) {
	Data := data.SetGroupSpecialTitleData{
		GroupId:      GroupID,
		UserId:       UserID,
		SpacialTitle: SpacialTitle,
		Duration:     Duration,
	}
	WsSend(c, "set_group_special_title", Data)
}
