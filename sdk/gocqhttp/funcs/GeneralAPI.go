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
func SendGroupForwardMsg(c *websocket.Conn, GroupId int, Messages interface{}) {
	Data := data.SendGroupForwardMsgData{
		GroupId:  GroupId,
		Messages: Messages,
	}
	WsSend(c, "send_group_forward_msg", Data)
}

// SendMsg Endpoint send_msg
func SendMsg(c *websocket.Conn, MessageType string, UserId int, Message string, GroupId int, AutoEscape bool) {
	Data := data.SendMsgData{
		MessageType: MessageType,
		UserId:      UserId,
		GroupId:     GroupId,
		Message:     Message,
		AutoEscape:  AutoEscape,
	}
	WsSend(c, "send_msg", Data)
}

// DeleteMsg Endpoint delete_msg
func DeleteMsg(c *websocket.Conn, MessageId int) {
	Data := data.DeleteMsgData{
		MessageId: MessageId,
	}
	WsSend(c, "delete_msg", Data)
}

// GetMsg Endpoint get_msg TODO
func GetMsg(c *websocket.Conn, MessageId int) {
	Data := data.GetMsgData{
		MessageId: MessageId,
	}
	WsSend(c, "get_msg", Data)
}

// GetForwardMsg Endpoint get_forward_msg TODO
func GetForwardMsg(c *websocket.Conn, MessageId string) {
	Data := data.GetForwardMsgData{
		MessageId: MessageId,
	}
	WsSend(c, "get_forward_msg", Data)
}

// GetImage Endpoint get_image TODO
func GetImage(c *websocket.Conn, File string) {
	Data := data.GetImageData{
		File: File,
	}
	WsSend(c, "get_image", Data)
}

// MarkMsgAsRead Endpoint mark_msg_as_read
func MarkMsgAsRead(c *websocket.Conn, MessageId int) {
	Data := data.MarkMsgAsReadData{
		MessageId: MessageId,
	}
	WsSend(c, "mark_msg_as_read", Data)
}

// SetGroupKick Endpoint set_group_kick
func SetGroupKick(c *websocket.Conn, GroupId int, UserId int, RejectAddRequest bool) {
	Data := data.SetGroupKickData{
		GroupId:          GroupId,
		UserId:           UserId,
		RejectAddRequest: RejectAddRequest,
	}
	WsSend(c, "set_group_kick", Data)
}

// SetGroupBan Endpoint set_group_ban
func SetGroupBan(c *websocket.Conn, GroupId int, UserId int, Duration int) {
	Data := data.SetGroupBanData{
		GroupId:  GroupId,
		UserId:   UserId,
		Duration: Duration,
	}
	WsSend(c, "set_group_ban", Data)
}

// SetGroupAnonymousBan Endpoint set_group_anonymous_ban
func SetGroupAnonymousBan(c *websocket.Conn, GroupId int, Flag string, Duration int) {
	Data := data.SetGroupAnonymousBanData{
		GroupId:  GroupId,
		Flag:     Flag,
		Duration: Duration,
	}
	WsSend(c, "set_group_anonymous_ban", Data)
}

// SetGroupWholeBan Endpoint set_group_whole_ban
func SetGroupWholeBan(c *websocket.Conn, GroupId int, Enable bool) {
	Data := data.SetGroupWholeBanData{
		GroupId: GroupId,
		Enable:  Enable,
	}
	WsSend(c, "set_group_whole_ban", Data)
}

// SetGroupAdmin Endpoint set_group_admin
func SetGroupAdmin(c *websocket.Conn, GroupId int, UserId int, Enable bool) {
	Data := data.SetGroupAdminData{
		GroupId: GroupId,
		UserId:  UserId,
		Enable:  Enable,
	}
	WsSend(c, "set_group_admin", Data)
}

// SetGroupCard Endpoint set_group_card
func SetGroupCard(c *websocket.Conn, GroupId int, UserId int, Card string) {
	Data := data.SetGroupCardData{
		GroupId: GroupId,
		UserId:  UserId,
		Card:    Card,
	}
	WsSend(c, "set_group_card", Data)
}

// SetGroupName Endpoint set_group_name
func SetGroupName(c *websocket.Conn, GroupId int, GroupName string) {
	Data := data.SetGroupNameData{
		GroupId:   GroupId,
		GroupName: GroupName,
	}
	WsSend(c, "set_group_name", Data)
}

// SetGroupLeave Endpoint set_group_leave
func SetGroupLeave(c *websocket.Conn, GroupId int, IsDismiss bool) {
	Data := data.SetGroupLeaveData{
		GroupId:   GroupId,
		IsDismiss: IsDismiss,
	}
	WsSend(c, "set_group_leave", Data)
}

// SetGroupSpecialTitle Endpoint set_group_spacial_title
func SetGroupSpecialTitle(c *websocket.Conn, GroupId int, UserId int, SpacialTitle string, Duration int) {
	Data := data.SetGroupSpecialTitleData{
		GroupId:      GroupId,
		UserId:       UserId,
		SpacialTitle: SpacialTitle,
		Duration:     Duration,
	}
	WsSend(c, "set_group_special_title", Data)
}

// SendGroupSign Endpoint set_group_sign
func SendGroupSign(c *websocket.Conn, GroupId int) {
	Data := data.SendGroupSignData{GroupId: GroupId}
	WsSend(c, "set_group_sign", Data)
}

// SetFriendAddRequest Endpoint set_friend_add_request
func SetFriendAddRequest(c *websocket.Conn, Flag string, Approve bool, Remark string) {
	Data := data.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	WsSend(c, "set_friend_add_request", Data)
}

// SetGroupAddRequest Endpoint set_group_add_request
func SetGroupAddRequest(c *websocket.Conn, Flag string, SubType string, Approve bool, Reason string) {
	Data := data.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: SubType,
		Approve: Approve,
		Reason:  Reason,
	}
	WsSend(c, "set_group_add_request", Data)
}

// GetLoginInfo Endpoint get_login_info TODO
func GetLoginInfo(c *websocket.Conn) {
	WsSend(c, "get_login_info", nil)
}

// QidianGetAccountInfo Endpoint qidian_get_account_info TODO
func QidianGetAccountInfo(c *websocket.Conn) {
	WsSend(c, "qidian_get_account_info", nil)
}

// SetQQProfile Endpoint set_qq_profile
func SetQQProfile(c *websocket.Conn, Nickname string, Company string, Email string, College string, PersonalNote string) {
	Data := data.SetQQProfileData{
		Nickname:     Nickname,
		Company:      Company,
		Email:        Email,
		College:      College,
		PersonalNote: PersonalNote,
	}
	WsSend(c, "set_qq_profile", Data)
}

// GetStrangerInfo Endpoint get_stranger_info TODO
func GetStrangerInfo(c *websocket.Conn, UserId int, NoCache bool) {
	Data := data.GetStrangerInfoData{
		UserId:  UserId,
		NoCache: NoCache,
	}
	WsSend(c, "get_stranger_info", Data)
}

// GetFriendList Endpoint get_friend_list TODO
func GetFriendList(c *websocket.Conn) {
	WsSend(c, "get_friend_list", nil)
}

// GetUnidirectionalFriendList Endpoint get_unidirectional_friend_list TODO
func GetUnidirectionalFriendList(c *websocket.Conn) {
	WsSend(c, "get_unidirectional_friend_list", nil)
}

// DeleteFriend Endpoint delete_friend
func DeleteFriend(c *websocket.Conn, FriendId int) {
	Data := data.DeleteFriendData{FriendId: FriendId}
	WsSend(c, "delete_friend", Data)
}

// GetGroupInfo Endpoint get_group_info TODO
func GetGroupInfo(c *websocket.Conn, GroupId int, NoCache bool) {
	Data := data.GetGroupInfoData{
		GroupId: GroupId,
		NoCache: NoCache,
	}
	WsSend(c, "get_group_info", Data)
}

// GetGroupList Endpoint get_group_list TODO
func GetGroupList(c *websocket.Conn) {
	WsSend(c, "get_group_list", nil)
}

// GetGroupMemberInfo Endpoint get_group_member_info TODO
func GetGroupMemberInfo(c *websocket.Conn, GroupId int, UserId int, NoCache bool) {
	Data := data.GetGroupMemberInfoData{
		GroupId: GroupId,
		UserId:  UserId,
		NoCache: NoCache,
	}
	WsSend(c, "get_group_member_info", Data)
}

// GetGroupMemberList Endpoint get_group_member_list TODO
func GetGroupMemberList(c *websocket.Conn, GroupId int, NoCache bool) {
	Data := data.GetGroupMemberListData{
		GroupId: GroupId,
		NoCache: NoCache,
	}
	WsSend(c, "get_group_member_list", Data)
}

// GetGroupHonorInfo Endpoint get_group_honor_info TODO
func GetGroupHonorInfo(c *websocket.Conn, GroupId int, Type string) {
	Data := data.GetGroupHonorInfoData{
		GroupId: GroupId,
		Type:    Type,
	}
	WsSend(c, "get_group_honor_info", Data)
}

// CanSendImage Endpoint can_send_image TODO
func CanSendImage(c *websocket.Conn) {
	WsSend(c, "can_send_image", nil)
}

// GetVersionInfo Endpoint get_version_info TODO
func GetVersionInfo(c *websocket.Conn) {
	WsSend(c, "get_version_info", nil)
}

// SetRestart Endpoint set_restart
func SetRestart(c *websocket.Conn, Delay int) {
	Data := data.SetRestartData{Delay: Delay}
	WsSend(c, "set_restart", Data)
}

// SetGroupPortrait Endpoint set_group_portrait
func SetGroupPortrait(c *websocket.Conn, GroupId int, File string, Cache int) {
	Data := data.SetGroupPortraitData{
		GroupId: GroupId,
		File:    File,
		Cache:   Cache,
	}
	WsSend(c, "set_group_portrait", Data)
}

// GetWordSlices Endpoint .get_word_slices TODO
func GetWordSlices(c *websocket.Conn, Content string) {
	Data := data.GetWordSlicesData{Content: Content}
	WsSend(c, "get_word_slices", Data)
}

// OCRImage Endpoint ocr_image TODO
func OCRImage(c *websocket.Conn, Image string) {
	Data := data.OCRImageData{Image: Image}
	WsSend(c, "ocr_image", Data)
}

// GetGroupSystemMsg Endpoint get_group_system_msg TODO
func GetGroupSystemMsg(c *websocket.Conn) {
	WsSend(c, "get_group_system_msg", nil)
}

// UploadPrivateFile Endpoint upload_private_file
func UploadPrivateFile(c *websocket.Conn, UserId int, File string, Name string) {
	Data := data.UploadPrivateFileData{
		UserId: UserId,
		File:   File,
		Name:   Name,
	}
	WsSend(c, "upload_private_file", Data)
}

// UploadGroupFile Endpoint upload_group_file
func UploadGroupFile(c *websocket.Conn, GroupId int, File string, Name string, Folder string) {
	Data := data.UploadGroupFileData{
		GroupId: GroupId,
		File:    File,
		Name:    Name,
		Folder:  Folder,
	}
	WsSend(c, "upload_group_file", Data)
}

// GetGroupFileSystemInfo Endpoint get_group_file_system_info TODO
func GetGroupFileSystemInfo(c *websocket.Conn, GroupId int) {
	Data := data.GetGroupFileSystemInfoData{GroupId: GroupId}
	WsSend(c, "get_group_file_system_info", Data)
}

// GetGroupRootFiles Endpoint get_group_root_files TODO
func GetGroupRootFiles(c *websocket.Conn, GroupId int) {
	Data := data.GetGroupRootFilesData{GroupId: GroupId}
	WsSend(c, "get_group_root_files", Data)
}

// GetGroupFilesByFolder Endpoint get_group_files_by_folder TODO
func GetGroupFilesByFolder(c *websocket.Conn, GroupId int, FolderId string) {
	Data := data.GetGroupFilesByFolderData{
		GroupId:  GroupId,
		FolderId: FolderId,
	}
	WsSend(c, "get_group_files_by_folder", Data)
}

// CreateGroupFileFolder Endpoint create_group_file_folder
func CreateGroupFileFolder(c *websocket.Conn, GroupId int, Name string, ParentId string) {
	Data := data.CreateGroupFileFolderData{
		GroupId:  GroupId,
		Name:     Name,
		ParentId: ParentId,
	}
	WsSend(c, "create_group_file_folder", Data)
}

// DeleteGroupFolder Endpoint delete_group_folder
func DeleteGroupFolder(c *websocket.Conn, GroupId int, FolderId string) {
	Data := data.DeleteGroupFolderData{
		GroupId:  GroupId,
		FolderId: FolderId,
	}
	WsSend(c, "delete_group_folder", Data)
}

// DeleteGroupFile Endpoint delete_group_file
func DeleteGroupFile(c *websocket.Conn, GroupId int, FileId string, BusId int) {
	Data := data.DeleteGroupFileData{
		GroupId: GroupId,
		FileId:  FileId,
		BusId:   BusId,
	}
	WsSend(c, "delete_group_file", Data)
}

// GetGroupFileUrl Endpoint get_group_file_url TODO
func GetGroupFileUrl(c *websocket.Conn, GroupId int, FileId string, BusId int) {
	Data := data.GetGroupFileUrlData{
		GroupId: GroupId,
		FileId:  FileId,
		BusId:   BusId,
	}
	WsSend(c, "get_group_file_url", Data)
}

// GetStatus Endpoint get_status TODO
func GetStatus(c *websocket.Conn) {
	WsSend(c, "get_status", nil)
}

// GetGroupAtAllRemain Endpoint get_group_at_all_remain TODO
func GetGroupAtAllRemain(c *websocket.Conn, GroupId int) {
	Data := data.GetGroupAtAllRemainData{GroupId: GroupId}
	WsSend(c, "get_group_at_all_remain", Data)
}

// HandleQuickOperation Endpoint .handle_quick_operation
func HandleQuickOperation(c *websocket.Conn, Context string, Operation string) {
	Data := data.HandleQuickOperationData{
		Context:   Context,
		Operation: Operation,
	}
	WsSend(c, "handle_quick_operation", Data)
}

// SendGroupNotice Endpoint _send_group_notice
func SendGroupNotice(c *websocket.Conn, GroupId int, Content string, Image string) {
	Data := data.SendGroupNoticeData{
		GroupId: GroupId,
		Content: Content,
		Image:   Image,
	}
	WsSend(c, "send_group_notice", Data)
}

// GetGroupNotice Endpoint _get_group_notice TODO
func GetGroupNotice(c *websocket.Conn, GroupId int) {
	Data := data.GetGroupNoticeData{GroupId: GroupId}
	WsSend(c, "get_group_notice", Data)
}

// ReloadEventFilter Endpoint reload_event_filter
func ReloadEventFilter(c *websocket.Conn, file string) {
	Data := data.ReloadEventFilterData{File: file}
	WsSend(c, "reload_event_filter", Data)
}

// DownloadFile Endpoint download_file TODO
func DownloadFile(c *websocket.Conn, Url string, ThreadCount int, Headers string) {
	Data := data.DownloadFileData{
		Url:         Url,
		ThreadCount: ThreadCount,
		Headers:     Headers,
	}
	WsSend(c, "download_file", Data)
}

// GetOnlineClients Endpoint get_online_clients TODO
func GetOnlineClients(c *websocket.Conn, NoCache bool) {
	Data := data.GetOnlineClientsData{NoCache: NoCache}
	WsSend(c, "get_online_clients", Data)
}

// GetGroupMsgHistory Endpoint get_group_msg_history TODO
func GetGroupMsgHistory(c *websocket.Conn, MessageSeq int, GroupId int) {
	Data := data.GetGroupMsgHistoryData{
		MessageSeq: MessageSeq,
		GroupId:    GroupId,
	}
	WsSend(c, "get_group_msg_history", Data)
}

// SetEssenceMsg Endpoint set_essence_msg
func SetEssenceMsg(c *websocket.Conn, MessageId int) {
	Data := data.SetEssenceMsgData{MessageId: MessageId}
	WsSend(c, "set_essence_msg", Data)
}

// DeleteEssenceMsg Endpoint delete_essence_msg
func DeleteEssenceMsg(c *websocket.Conn, MessageId int) {
	Data := data.DeleteEssenceMsgData{MessageId: MessageId}
	WsSend(c, "delete_essence_msg", Data)
}

// GetEssenceMsgList Endpoint get_essence_msg_list TODO
func GetEssenceMsgList(c *websocket.Conn, GroupId int) {
	Data := data.GetEssenceMsgListData{GroupId: GroupId}
	WsSend(c, "get_essence_msg_list", Data)
}

// CheckUrlSafely Endpoint check_url_safely TODO
func CheckUrlSafely(c *websocket.Conn, Url string) {
	Data := data.CheckUrlSafelyData{Url: Url}
	WsSend(c, "check_url_safely", Data)
}

// GetModelShow Endpoint _get_module_show TODO
func GetModelShow(c *websocket.Conn, Model string) {
	Data := data.GetModelShowData{Model: Model}
	WsSend(c, "_get_module_show", Data)
}

// SetModelShow Endpoint _set_module_show
func SetModelShow(c *websocket.Conn, Model string, ModelShow string) {
	Data := data.SetModelShowData{
		Model:     Model,
		ModelShow: ModelShow,
	}
	WsSend(c, "_set_module_show", Data)
}

// DeleteUnidirectionalFriend Endpoint delete_unidirectional_friend
func DeleteUnidirectionalFriend(c *websocket.Conn, UserId int) {
	Data := data.DeleteUnidirectionalFriendData{UserId: UserId}
	WsSend(c, "delete_unidirectional_friend", Data)
}

// SendPrivateForwardMsg Endpoint send_private_forward_msg TODO
func SendPrivateForwardMsg(c *websocket.Conn, UserId int, Messages interface{}) {
	Data := data.SendPrivateForwardMsgData{
		UserId:   UserId,
		Messages: Messages,
	}
	WsSend(c, "send_private_forward_msg", Data)
}
