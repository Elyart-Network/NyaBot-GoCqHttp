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
	GroupId  int         `json:"group_id"`
	Messages interface{} `json:"messages"`
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
	MessageId string `json:"message_id"`
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

// SetGroupSpecialTitleData Endpoint set_group_special_title
type SetGroupSpecialTitleData struct {
	GroupId      int    `json:"group_id"`
	UserId       int    `json:"user_id"`
	SpacialTitle string `json:"special_title"`
	Duration     int    `json:"duration"`
}

// SendGroupSignData Endpoint set_group_sign
type SendGroupSignData struct {
	GroupId int `json:"group_id"`
}

// SetFriendAddRequestData Endpoint set_friend_add_request
type SetFriendAddRequestData struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve"`
	Remark  string `json:"remark"`
}

// SetGroupAddRequestData Endpoint set_group_add_request
type SetGroupAddRequestData struct {
	Flag    string `json:"flag"`
	SubType string `json:"sub_type"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}

// GetLoginInfoData Endpoint get_login_info
type GetLoginInfoData struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
}

// QidianGetAccountInfoData Endpoint qidian_get_account_info
type QidianGetAccountInfoData struct {
	MasterId   int    `json:"master_id"`
	ExtName    string `json:"ext_name"`
	CreateTime int    `json:"create_time"`
}

// SetQQProfileData Endpoint set_qq_profile
type SetQQProfileData struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	College      string `json:"college"`
	PersonalNote string `json:"personal_note"`
}

// GetStrangerInfoData Endpoint get_stranger_info
type GetStrangerInfoData struct {
	UserId  int  `json:"user_id"`
	NoCache bool `json:"no_cache"`
}

// GetFriendListData Endpoint get_friend_list
type GetFriendListData struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

// GetUnidirectionalFriendListData Endpoint get_unidirectional_friend_list
type GetUnidirectionalFriendListData struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Source   string `json:"source"`
}

// DeleteFriendData Endpoint delete_friend
type DeleteFriendData struct {
	FriendId int `json:"friend_id"`
}

// GetGroupInfoData Endpoint get_group_info
type GetGroupInfoData struct {
	GroupId int  `json:"group_id"`
	NoCache bool `json:"no_cache"`
}

// GetGroupMemberInfoData Endpoint get_group_member_info
type GetGroupMemberInfoData struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	NoCache bool `json:"no_cache"`
}

// GetGroupMemberListData Endpoint get_group_member_list
type GetGroupMemberListData struct {
	GroupId int  `json:"group_id"`
	NoCache bool `json:"no_cache"`
}

// GetGroupHonorInfoData Endpoint get_group_honor_info
type GetGroupHonorInfoData struct {
	GroupId int    `json:"group_id"`
	Type    string `json:"type"`
}

// SetRestartData Endpoint set_restart
type SetRestartData struct {
	Delay int `json:"delay"`
}

// SetGroupPortraitData Endpoint set_group_portrait
type SetGroupPortraitData struct {
	GroupId int    `json:"group_id"`
	File    string `json:"file"`
	Cache   int    `json:"cache"`
}

// GetWordSlicesData Endpoint .get_word_slices
type GetWordSlicesData struct {
	Content string `json:"content"`
}

// OCRImageData Endpoint ocr_image
type OCRImageData struct {
	Image string `json:"image"`
}

// UploadPrivateFileData Endpoint upload_private_file
type UploadPrivateFileData struct {
	UserId int    `json:"user_id"`
	File   string `json:"file"`
	Name   string `json:"name"`
}

// UploadGroupFileData Endpoint upload_group_file
type UploadGroupFileData struct {
	GroupId int    `json:"group_id"`
	File    string `json:"file"`
	Name    string `json:"name"`
	Folder  string `json:"folder"`
}

// GetGroupFileSystemInfoData Endpoint get_group_file_system_info
type GetGroupFileSystemInfoData struct {
	GroupId int `json:"group_id"`
}

// GetGroupRootFilesData Endpoint get_group_root_files
type GetGroupRootFilesData struct {
	GroupId int `json:"group_id"`
}

// GetGroupFilesByFolderData Endpoint get_group_files_by_folder
type GetGroupFilesByFolderData struct {
	GroupId  int    `json:"group_id"`
	FolderId string `json:"folder_id"`
}

// CreateGroupFileFolderData Endpoint create_group_file_folder
type CreateGroupFileFolderData struct {
	GroupId  int    `json:"group_id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

// DeleteGroupFolderData Endpoint delete_group_folder
type DeleteGroupFolderData struct {
	GroupId  int    `json:"group_id"`
	FolderId string `json:"folder_id"`
}

// DeleteGroupFileData Endpoint delete_group_file
type DeleteGroupFileData struct {
	GroupId int    `json:"group_id"`
	FileId  string `json:"file_id"`
	BusId   int    `json:"busid"`
}

// GetGroupFileUrlData Endpoint get_group_file_url
type GetGroupFileUrlData struct {
	GroupId int    `json:"group_id"`
	FileId  string `json:"file_id"`
	BusId   int    `json:"busid"`
}

// GetGroupAtAllRemainData Endpoint get_group_at_all_remain
type GetGroupAtAllRemainData struct {
	GroupId int `json:"group_id"`
}

// HandleQuickOperationData Endpoint .handle_quick_operation
type HandleQuickOperationData struct {
	Context   string `json:"context"`
	Operation string `json:"operation"`
}

// SendGroupNoticeData Endpoint _send_group_notice
type SendGroupNoticeData struct {
	GroupId int    `json:"group_id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

// GetGroupNoticeData Endpoint _get_group_notice
type GetGroupNoticeData struct {
	GroupId int `json:"group_id"`
}

// ReloadEventFilterData Endpoint reload_event_filter
type ReloadEventFilterData struct {
	File string `json:"file"`
}

// DownloadFileData Endpoint download_file
type DownloadFileData struct {
	Url         string `json:"url"`
	ThreadCount int    `json:"thread_count"`
	Headers     string `json:"headers"`
}

// GetOnlineClientsData Endpoint get_online_clients
type GetOnlineClientsData struct {
	NoCache bool `json:"no_cache"`
}

// GetGroupMsgHistoryData Endpoint get_group_msg_history
type GetGroupMsgHistoryData struct {
	MessageSeq int `json:"message_seq"`
	GroupId    int `json:"group_id"`
}

// SetEssenceMsgData Endpoint set_essence_msg
type SetEssenceMsgData struct {
	MessageId int `json:"message_id"`
}

// DeleteEssenceMsgData Endpoint delete_essence_msg
type DeleteEssenceMsgData struct {
	MessageId int `json:"message_id"`
}

// GetEssenceMsgListData Endpoint get_essence_msg_list
type GetEssenceMsgListData struct {
	GroupId int `json:"group_id"`
}

// CheckUrlSafelyData Endpoint check_url_safely
type CheckUrlSafelyData struct {
	Url string `json:"url"`
}

// GetModelShowData Endpoint _get_model_show
type GetModelShowData struct {
	Model string `json:"model"`
}

// SetModelShowData Endpoint _set_model_show
type SetModelShowData struct {
	Model     string `json:"model"`
	ModelShow string `json:"model_show"`
}

// DeleteUnidirectionalFriendData Endpoint delete_unidirectional_friend
type DeleteUnidirectionalFriendData struct {
	UserId int `json:"user_id"`
}

// SendPrivateForwardMsgData Endpoint send_private_forward_msg
type SendPrivateForwardMsgData struct {
	UserId   int         `json:"user_id"`
	Messages interface{} `json:"messages"`
}
