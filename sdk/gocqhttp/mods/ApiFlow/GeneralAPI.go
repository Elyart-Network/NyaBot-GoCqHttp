package ApiFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"github.com/gorilla/websocket"
)

// SendPrivateMsg Endpoint send_private_msg
func SendPrivateMsg(UserId int, GroupId int, Message string, AutoEscape bool) {
	Data := data.SendPrivateMsgData{
		UserId:     UserId,
		GroupId:    GroupId,
		Message:    Message,
		AutoEscape: AutoEscape,
	}
	postRequest("send_private_msg", Data)
}

// SendGroupMsg Endpoint send_group_msg
func SendGroupMsg(GroupId int, Message string, AutoEscape bool) {
	Data := data.SendGroupMsgData{
		GroupId:    GroupId,
		Message:    Message,
		AutoEscape: AutoEscape,
	}
	postRequest("send_group_msg", Data)
}

// SendGroupForwardMsg Endpoint send_group_forward_msg
func SendGroupForwardMsg(GroupId int, Messages interface{}) {
	Data := data.SendGroupForwardMsgData{
		GroupId:  GroupId,
		Messages: Messages,
	}
	postRequest("send_group_forward_msg", Data)
}

// SendMsg Endpoint send_msg
func SendMsg(MessageType string, UserId int, Message string, GroupId int, AutoEscape bool) {
	Data := data.SendMsgData{
		MessageType: MessageType,
		UserId:      UserId,
		GroupId:     GroupId,
		Message:     Message,
		AutoEscape:  AutoEscape,
	}
	postRequest("send_msg", Data)
}

// DeleteMsg Endpoint delete_msg
func DeleteMsg(MessageId int) {
	Data := data.DeleteMsgData{
		MessageId: MessageId,
	}
	postRequest("delete_msg", Data)
}

// GetMsg Endpoint get_msg TODO
func GetMsg(MessageId int) {
	Data := data.GetMsgData{
		MessageId: MessageId,
	}
	postRequest("get_msg", Data)
}

// GetForwardMsg Endpoint get_forward_msg TODO
func GetForwardMsg(MessageId string) {
	Data := data.GetForwardMsgData{
		MessageId: MessageId,
	}
	postRequest("get_forward_msg", Data)
}

// GetImage Endpoint get_image TODO
func GetImage(File string) {
	Data := data.GetImageData{
		File: File,
	}
	postRequest("get_image", Data)
}

// MarkMsgAsRead Endpoint mark_msg_as_read
func MarkMsgAsRead(MessageId int) {
	Data := data.MarkMsgAsReadData{
		MessageId: MessageId,
	}
	postRequest("mark_msg_as_read", Data)
}

// SetGroupKick Endpoint set_group_kick
func SetGroupKick(GroupId int, UserId int, RejectAddRequest bool) {
	Data := data.SetGroupKickData{
		GroupId:          GroupId,
		UserId:           UserId,
		RejectAddRequest: RejectAddRequest,
	}
	postRequest("set_group_kick", Data)
}

// SetGroupBan Endpoint set_group_ban
func SetGroupBan(GroupId int, UserId int, Duration int) {
	Data := data.SetGroupBanData{
		GroupId:  GroupId,
		UserId:   UserId,
		Duration: Duration,
	}
	postRequest("set_group_ban", Data)
}

// SetGroupAnonymousBan Endpoint set_group_anonymous_ban
func SetGroupAnonymousBan(GroupId int, Flag string, Duration int) {
	Data := data.SetGroupAnonymousBanData{
		GroupId:  GroupId,
		Flag:     Flag,
		Duration: Duration,
	}
	postRequest("set_group_anonymous_ban", Data)
}

// SetGroupWholeBan Endpoint set_group_whole_ban
func SetGroupWholeBan(GroupId int, Enable bool) {
	Data := data.SetGroupWholeBanData{
		GroupId: GroupId,
		Enable:  Enable,
	}
	postRequest("set_group_whole_ban", Data)
}

// SetGroupAdmin Endpoint set_group_admin
func SetGroupAdmin(GroupId int, UserId int, Enable bool) {
	Data := data.SetGroupAdminData{
		GroupId: GroupId,
		UserId:  UserId,
		Enable:  Enable,
	}
	postRequest("set_group_admin", Data)
}

// SetGroupCard Endpoint set_group_card
func SetGroupCard(GroupId int, UserId int, Card string) {
	Data := data.SetGroupCardData{
		GroupId: GroupId,
		UserId:  UserId,
		Card:    Card,
	}
	postRequest("set_group_card", Data)
}

// SetGroupName Endpoint set_group_name
func SetGroupName(GroupId int, GroupName string) {
	Data := data.SetGroupNameData{
		GroupId:   GroupId,
		GroupName: GroupName,
	}
	postRequest("set_group_name", Data)
}

// SetGroupLeave Endpoint set_group_leave
func SetGroupLeave(GroupId int, IsDismiss bool) {
	Data := data.SetGroupLeaveData{
		GroupId:   GroupId,
		IsDismiss: IsDismiss,
	}
	postRequest("set_group_leave", Data)
}

// SetGroupSpecialTitle Endpoint set_group_spacial_title
func SetGroupSpecialTitle(GroupId int, UserId int, SpacialTitle string, Duration int) {
	Data := data.SetGroupSpecialTitleData{
		GroupId:      GroupId,
		UserId:       UserId,
		SpacialTitle: SpacialTitle,
		Duration:     Duration,
	}
	postRequest("set_group_special_title", Data)
}

// SendGroupSign Endpoint set_group_sign
func SendGroupSign(GroupId int) {
	Data := data.SendGroupSignData{GroupId: GroupId}
	postRequest("set_group_sign", Data)
}

// SetFriendAddRequest Endpoint set_friend_add_request
func SetFriendAddRequest(Flag string, Approve bool, Remark string) {
	Data := data.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	postRequest("set_friend_add_request", Data)
}

// SetGroupAddRequest Endpoint set_group_add_request
func SetGroupAddRequest(Flag string, SubType string, Approve bool, Reason string) {
	Data := data.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: SubType,
		Approve: Approve,
		Reason:  Reason,
	}
	postRequest("set_group_add_request", Data)
}

// GetLoginInfo Endpoint get_login_info TODO
func GetLoginInfo(c *websocket.Conn) {
	postRequest("get_login_info", nil)
}

// QidianGetAccountInfo Endpoint qidian_get_account_info TODO
func QidianGetAccountInfo(c *websocket.Conn) {
	postRequest("qidian_get_account_info", nil)
}

// SetQQProfile Endpoint set_qq_profile
func SetQQProfile(Nickname string, Company string, Email string, College string, PersonalNote string) {
	Data := data.SetQQProfileData{
		Nickname:     Nickname,
		Company:      Company,
		Email:        Email,
		College:      College,
		PersonalNote: PersonalNote,
	}
	postRequest("set_qq_profile", Data)
}

// GetStrangerInfo Endpoint get_stranger_info TODO
func GetStrangerInfo(UserId int, NoCache bool) {
	Data := data.GetStrangerInfoData{
		UserId:  UserId,
		NoCache: NoCache,
	}
	postRequest("get_stranger_info", Data)
}

// GetFriendList Endpoint get_friend_list TODO
func GetFriendList(c *websocket.Conn) {
	postRequest("get_friend_list", nil)
}

// GetUnidirectionalFriendList Endpoint get_unidirectional_friend_list TODO
func GetUnidirectionalFriendList(c *websocket.Conn) {
	postRequest("get_unidirectional_friend_list", nil)
}

// DeleteFriend Endpoint delete_friend
func DeleteFriend(FriendId int) {
	Data := data.DeleteFriendData{FriendId: FriendId}
	postRequest("delete_friend", Data)
}

// GetGroupInfo Endpoint get_group_info TODO
func GetGroupInfo(GroupId int, NoCache bool) {
	Data := data.GetGroupInfoData{
		GroupId: GroupId,
		NoCache: NoCache,
	}
	postRequest("get_group_info", Data)
}

// GetGroupList Endpoint get_group_list TODO
func GetGroupList(c *websocket.Conn) {
	postRequest("get_group_list", nil)
}

// GetGroupMemberInfo Endpoint get_group_member_info TODO
func GetGroupMemberInfo(GroupId int, UserId int, NoCache bool) {
	Data := data.GetGroupMemberInfoData{
		GroupId: GroupId,
		UserId:  UserId,
		NoCache: NoCache,
	}
	postRequest("get_group_member_info", Data)
}

// GetGroupMemberList Endpoint get_group_member_list TODO
func GetGroupMemberList(GroupId int, NoCache bool) {
	Data := data.GetGroupMemberListData{
		GroupId: GroupId,
		NoCache: NoCache,
	}
	postRequest("get_group_member_list", Data)
}

// GetGroupHonorInfo Endpoint get_group_honor_info TODO
func GetGroupHonorInfo(GroupId int, Type string) {
	Data := data.GetGroupHonorInfoData{
		GroupId: GroupId,
		Type:    Type,
	}
	postRequest("get_group_honor_info", Data)
}

// CanSendImage Endpoint can_send_image TODO
func CanSendImage(c *websocket.Conn) {
	postRequest("can_send_image", nil)
}

// GetVersionInfo Endpoint get_version_info TODO
func GetVersionInfo(c *websocket.Conn) {
	postRequest("get_version_info", nil)
}

// SetRestart Endpoint set_restart
func SetRestart(Delay int) {
	Data := data.SetRestartData{Delay: Delay}
	postRequest("set_restart", Data)
}

// SetGroupPortrait Endpoint set_group_portrait
func SetGroupPortrait(GroupId int, File string, Cache int) {
	Data := data.SetGroupPortraitData{
		GroupId: GroupId,
		File:    File,
		Cache:   Cache,
	}
	postRequest("set_group_portrait", Data)
}

// GetWordSlices Endpoint .get_word_slices TODO
func GetWordSlices(Content string) {
	Data := data.GetWordSlicesData{Content: Content}
	postRequest("get_word_slices", Data)
}

// OCRImage Endpoint ocr_image TODO
func OCRImage(Image string) {
	Data := data.OCRImageData{Image: Image}
	postRequest("ocr_image", Data)
}

// GetGroupSystemMsg Endpoint get_group_system_msg TODO
func GetGroupSystemMsg(c *websocket.Conn) {
	postRequest("get_group_system_msg", nil)
}

// UploadPrivateFile Endpoint upload_private_file
func UploadPrivateFile(UserId int, File string, Name string) {
	Data := data.UploadPrivateFileData{
		UserId: UserId,
		File:   File,
		Name:   Name,
	}
	postRequest("upload_private_file", Data)
}

// UploadGroupFile Endpoint upload_group_file
func UploadGroupFile(GroupId int, File string, Name string, Folder string) {
	Data := data.UploadGroupFileData{
		GroupId: GroupId,
		File:    File,
		Name:    Name,
		Folder:  Folder,
	}
	postRequest("upload_group_file", Data)
}

// GetGroupFileSystemInfo Endpoint get_group_file_system_info TODO
func GetGroupFileSystemInfo(GroupId int) {
	Data := data.GetGroupFileSystemInfoData{GroupId: GroupId}
	postRequest("get_group_file_system_info", Data)
}

// GetGroupRootFiles Endpoint get_group_root_files TODO
func GetGroupRootFiles(GroupId int) {
	Data := data.GetGroupRootFilesData{GroupId: GroupId}
	postRequest("get_group_root_files", Data)
}

// GetGroupFilesByFolder Endpoint get_group_files_by_folder TODO
func GetGroupFilesByFolder(GroupId int, FolderId string) {
	Data := data.GetGroupFilesByFolderData{
		GroupId:  GroupId,
		FolderId: FolderId,
	}
	postRequest("get_group_files_by_folder", Data)
}

// CreateGroupFileFolder Endpoint create_group_file_folder
func CreateGroupFileFolder(GroupId int, Name string, ParentId string) {
	Data := data.CreateGroupFileFolderData{
		GroupId:  GroupId,
		Name:     Name,
		ParentId: ParentId,
	}
	postRequest("create_group_file_folder", Data)
}

// DeleteGroupFolder Endpoint delete_group_folder
func DeleteGroupFolder(GroupId int, FolderId string) {
	Data := data.DeleteGroupFolderData{
		GroupId:  GroupId,
		FolderId: FolderId,
	}
	postRequest("delete_group_folder", Data)
}

// DeleteGroupFile Endpoint delete_group_file
func DeleteGroupFile(GroupId int, FileId string, BusId int) {
	Data := data.DeleteGroupFileData{
		GroupId: GroupId,
		FileId:  FileId,
		BusId:   BusId,
	}
	postRequest("delete_group_file", Data)
}

// GetGroupFileUrl Endpoint get_group_file_url TODO
func GetGroupFileUrl(GroupId int, FileId string, BusId int) {
	Data := data.GetGroupFileUrlData{
		GroupId: GroupId,
		FileId:  FileId,
		BusId:   BusId,
	}
	postRequest("get_group_file_url", Data)
}

// GetStatus Endpoint get_status TODO
func GetStatus(c *websocket.Conn) {
	postRequest("get_status", nil)
}

// GetGroupAtAllRemain Endpoint get_group_at_all_remain TODO
func GetGroupAtAllRemain(GroupId int) {
	Data := data.GetGroupAtAllRemainData{GroupId: GroupId}
	postRequest("get_group_at_all_remain", Data)
}

// HandleQuickOperation Endpoint .handle_quick_operation
func HandleQuickOperation(Context string, Operation string) {
	Data := data.HandleQuickOperationData{
		Context:   Context,
		Operation: Operation,
	}
	postRequest("handle_quick_operation", Data)
}

// SendGroupNotice Endpoint _send_group_notice
func SendGroupNotice(GroupId int, Content string, Image string) {
	Data := data.SendGroupNoticeData{
		GroupId: GroupId,
		Content: Content,
		Image:   Image,
	}
	postRequest("send_group_notice", Data)
}

// GetGroupNotice Endpoint _get_group_notice TODO
func GetGroupNotice(GroupId int) {
	Data := data.GetGroupNoticeData{GroupId: GroupId}
	postRequest("get_group_notice", Data)
}

// ReloadEventFilter Endpoint reload_event_filter
func ReloadEventFilter(file string) {
	Data := data.ReloadEventFilterData{File: file}
	postRequest("reload_event_filter", Data)
}

// DownloadFile Endpoint download_file TODO
func DownloadFile(Url string, ThreadCount int, Headers string) {
	Data := data.DownloadFileData{
		Url:         Url,
		ThreadCount: ThreadCount,
		Headers:     Headers,
	}
	postRequest("download_file", Data)
}

// GetOnlineClients Endpoint get_online_clients TODO
func GetOnlineClients(NoCache bool) {
	Data := data.GetOnlineClientsData{NoCache: NoCache}
	postRequest("get_online_clients", Data)
}

// GetGroupMsgHistory Endpoint get_group_msg_history TODO
func GetGroupMsgHistory(MessageSeq int, GroupId int) {
	Data := data.GetGroupMsgHistoryData{
		MessageSeq: MessageSeq,
		GroupId:    GroupId,
	}
	postRequest("get_group_msg_history", Data)
}

// SetEssenceMsg Endpoint set_essence_msg
func SetEssenceMsg(MessageId int) {
	Data := data.SetEssenceMsgData{MessageId: MessageId}
	postRequest("set_essence_msg", Data)
}

// DeleteEssenceMsg Endpoint delete_essence_msg
func DeleteEssenceMsg(MessageId int) {
	Data := data.DeleteEssenceMsgData{MessageId: MessageId}
	postRequest("delete_essence_msg", Data)
}

// GetEssenceMsgList Endpoint get_essence_msg_list TODO
func GetEssenceMsgList(GroupId int) {
	Data := data.GetEssenceMsgListData{GroupId: GroupId}
	postRequest("get_essence_msg_list", Data)
}

// CheckUrlSafely Endpoint check_url_safely TODO
func CheckUrlSafely(Url string) {
	Data := data.CheckUrlSafelyData{Url: Url}
	postRequest("check_url_safely", Data)
}

// GetModelShow Endpoint _get_module_show TODO
func GetModelShow(Model string) {
	Data := data.GetModelShowData{Model: Model}
	postRequest("_get_module_show", Data)
}

// SetModelShow Endpoint _set_module_show
func SetModelShow(Model string, ModelShow string) {
	Data := data.SetModelShowData{
		Model:     Model,
		ModelShow: ModelShow,
	}
	postRequest("_set_module_show", Data)
}

// DeleteUnidirectionalFriend Endpoint delete_unidirectional_friend
func DeleteUnidirectionalFriend(UserId int) {
	Data := data.DeleteUnidirectionalFriendData{UserId: UserId}
	postRequest("delete_unidirectional_friend", Data)
}

// SendPrivateForwardMsg Endpoint send_private_forward_msg TODO
func SendPrivateForwardMsg(UserId int, Messages interface{}) {
	Data := data.SendPrivateForwardMsgData{
		UserId:   UserId,
		Messages: Messages,
	}
	postRequest("send_private_forward_msg", Data)
}
