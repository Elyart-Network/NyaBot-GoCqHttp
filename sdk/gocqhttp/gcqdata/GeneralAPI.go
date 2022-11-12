package gcqdata

// SendPrivateMsgData Endpoint send_private_msg
type SendPrivateMsgData struct {
	UserId     int    `json:"user_id"`
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

// SendPrivateMsgResp Endpoint send_private_msg
type SendPrivateMsgResp struct {
	MessageId int `json:"message_id"`
}

// SendGroupMsgData Endpoint send_group_msg
type SendGroupMsgData struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

// SendGroupMsgResp Endpoint send_group_msg
type SendGroupMsgResp struct {
	MessageId int `json:"message_id"`
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

// SendMsgResp Endpoint send_msg
type SendMsgResp struct {
	MessageId int `json:"message_id"`
}

// DeleteMsgData Endpoint delete_msg
type DeleteMsgData struct {
	MessageId int `json:"message_id"`
}

// GetMsgData Endpoint get_msg
type GetMsgData struct {
	MessageId int `json:"message_id"`
}

// GetMsgResp Endpoint get_msg
type GetMsgResp struct {
	MessageId  int         `json:"message_id"`
	RealId     int         `json:"real_id"`
	Sender     interface{} `json:"sender"`
	Time       int         `json:"time"`
	Message    string      `json:"message"`
	RawMessage string      `json:"raw_message"`
}

// GetForwardMsgData Endpoint get_forward_msg
type GetForwardMsgData struct {
	MessageId string `json:"message_id"`
}

// GetForwardMsgResp Endpoint get_forward_msg
type GetForwardMsgResp struct {
	Messages interface{} `json:"messages"`
}

// GetImageData Endpoint get_image
type GetImageData struct {
	File string `json:"file"`
}

// GetImageResp Endpoint get_image
type GetImageResp struct {
	Size     int    `json:"size"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
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

// QidianGetAccountInfoResp Endpoint qidian_get_account_info
type QidianGetAccountInfoResp struct {
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

// GetStrangerInfoResp Endpoint get_stranger_info
type GetStrangerInfoResp struct {
	UserId    int    `json:"user_id"`
	Nickname  string `json:"nickname"`
	Sex       string `json:"sex"`
	Age       int    `json:"age"`
	Qid       string `json:"qid"`
	Level     int    `json:"level"`
	LoginDays int    `json:"login_days"`
}

// GetFriendListResp Endpoint get_friend_list
type GetFriendListResp struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

// GetUnidirectionalFriendListResp Endpoint get_unidirectional_friend_list
type GetUnidirectionalFriendListResp struct {
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

// GetGroupInfoResp Endpoint get_group_info
type GetGroupInfoResp struct {
	GroupId         int    `json:"group_id"`
	GroupName       string `json:"group_name"`
	GroupMemo       string `json:"group_memo"`
	GroupCreateTime int    `json:"group_create_time"`
	GroupLevel      int    `json:"group_level"`
	MemberCount     int    `json:"member_count"`
	MaxMemberCount  int    `json:"max_member_count"`
}

// GetGroupListResp Endpoint get_group_list
type GetGroupListResp struct {
	GroupId         int    `json:"group_id"`
	GroupName       string `json:"group_name"`
	GroupMemo       string `json:"group_memo"`
	GroupCreateTime int    `json:"group_create_time"`
	GroupLevel      int    `json:"group_level"`
	MemberCount     int    `json:"member_count"`
	MaxMemberCount  int    `json:"max_member_count"`
}

// GetGroupMemberInfoData Endpoint get_group_member_info
type GetGroupMemberInfoData struct {
	GroupId int  `json:"group_id"`
	UserId  int  `json:"user_id"`
	NoCache bool `json:"no_cache"`
}

// GetGroupMemberInfoResp Endpoint get_group_member_info
type GetGroupMemberInfoResp struct {
	GroupId         int    `json:"group_id"`
	UserId          int    `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int    `json:"age"`
	Area            string `json:"area"`
	JoinTime        int    `json:"join_time"`
	LastSentTime    int    `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int    `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
	ShutUpTimestamp int    `json:"shut_up_timestamp"`
}

// GetGroupMemberListData Endpoint get_group_member_list
type GetGroupMemberListData struct {
	GroupId int  `json:"group_id"`
	NoCache bool `json:"no_cache"`
}

// GetGroupMemberListResp Endpoint get_group_member_list
type GetGroupMemberListResp struct {
	GroupId         int    `json:"group_id"`
	UserId          int    `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int    `json:"age"`
	Area            string `json:"area"`
	JoinTime        int    `json:"join_time"`
	LastSentTime    int    `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int    `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
	ShutUpTimestamp int    `json:"shut_up_timestamp"`
}

// GetGroupHonorInfoData Endpoint get_group_honor_info
type GetGroupHonorInfoData struct {
	GroupId int    `json:"group_id"`
	Type    string `json:"type"`
}

// GetGroupHonorInfoResp Endpoint get_group_honor_info
type GetGroupHonorInfoResp struct {
	GroupId          int              `json:"group_id"`
	CurrentTalkative currentTalkative `json:"current_talkative"`
	TalkativeList    otherTalkative   `json:"talkative_list"`
	PerformerList    otherTalkative   `json:"performer_list"`
	LegendList       otherTalkative   `json:"legend_list"`
	StrongNewbieList otherTalkative   `json:"strong_newbie_list"`
	EmotionList      otherTalkative   `json:"emotion_list"`
}

type currentTalkative struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	DayCount int    `json:"day_count"`
}

type otherTalkative struct {
	UserId      int    `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

// CanSendImageResp Endpoint can_send_image
type CanSendImageResp struct {
	Yes bool `json:"yes"`
}

// CanSendRecordResp Endpoint can_send_record
type CanSendRecordResp struct {
	Yes bool `json:"yes"`
}

// GetVersionInfoResp Endpoint get_version_info
type GetVersionInfoResp struct {
	AppName                  string `json:"app_name"`
	AppVersion               string `json:"app_version"`
	AppFullName              string `json:"app_full_name"`
	ProtocolVersion          int    `json:"protocol_version"`
	CoolqEdition             string `json:"coolq_edition"`
	CoolqDirectory           string `json:"coolq_directory"`
	GoCqHttp                 string `json:"go-cqhttp"`
	PluginVersion            string `json:"plugin_version"`
	PluginBuildNumber        string `json:"plugin_build_number"`
	PluginBuildConfiguration string `json:"plugin_build_configuration"`
	RuntimeVersion           string `json:"runtime_version"`
	RuntimeOs                string `json:"runtime_os"`
	Version                  string `json:"version"`
	Protocol                 string `json:"protocol"`
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

// GetWordSlicesResp Endpoint .get_word_slices
type GetWordSlicesResp struct {
	Slices []string `json:"slices"`
}

// OCRImageData Endpoint ocr_image
type OCRImageData struct {
	Image string `json:"image"`
}

// OCRImageResp Endpoint ocr_image
type OCRImageResp struct {
	Texts    string        `json:"texts"`
	Language textDetection `json:"language"`
}

type textDetection struct {
	Text        string `json:"text"`
	Confidence  int    `json:"confidence"`
	Coordinates int    `json:"coordinates"`
}

// GetGroupSystemMsgResp Endpoint get_group_system_msg
type GetGroupSystemMsgResp struct {
	InvitedRequests []invitedRequests `json:"invited_requests"`
	JoinRequests    []joinRequests    `json:"join_requests"`
}

type invitedRequests struct {
	RequestId   int    `json:"request_id"`
	InvitorUin  int    `json:"invitor_uin"`
	InvitorNick string `json:"invitor_nick"`
	GroupId     int    `json:"group_id"`
	GroupName   string `json:"group_name"`
	Checked     bool   `json:"checked"`
	Actor       int    `json:"actor"`
}

type joinRequests struct {
	RequestId     int    `json:"request_id"`
	RequesterUin  int    `json:"requester_uin"`
	RequesterNick string `json:"requester_nick"`
	Message       string `json:"message"`
	GroupId       int    `json:"group_id"`
	GroupName     string `json:"group_name"`
	Checked       bool   `json:"checked"`
	Actor         int    `json:"actor"`
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

// GetGroupFileSystemInfoResp Endpoint get_group_file_system_info
type GetGroupFileSystemInfoResp struct {
	FileCount  int `json:"file_count"`
	LimitCount int `json:"limit_count"`
	UsedSpace  int `json:"used_space"`
	TotalSpace int `json:"total_space"`
}

// GetGroupRootFilesData Endpoint get_group_root_files
type GetGroupRootFilesData struct {
	GroupId int `json:"group_id"`
}

// GetGroupRootFilesResp Endpoint get_group_root_files
type GetGroupRootFilesResp struct {
	Files   []file   `json:"files"`
	Folders []folder `json:"folders"`
}

// GetGroupFilesByFolderData Endpoint get_group_files_by_folder
type GetGroupFilesByFolderData struct {
	GroupId  int    `json:"group_id"`
	FolderId string `json:"folder_id"`
}

// GetGroupFilesByFolderResp Endpoint get_group_files_by_folder
type GetGroupFilesByFolderResp struct {
	Files   []file   `json:"files"`
	Folders []folder `json:"folders"`
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

// GetGroupFileUrlResp Endpoint get_group_file_url
type GetGroupFileUrlResp struct {
	Url string `json:"url"`
}

type file struct {
	GroupId       int    `json:"group_id"`
	FileId        string `json:"file_id"`
	FileName      string `json:"file_name"`
	Busid         int    `json:"busid"`
	FileSize      int    `json:"file_size"`
	UploadTime    int    `json:"upload_time"`
	DeadTime      int    `json:"dead_time"`
	ModifyTime    int    `json:"modify_time"`
	DownloadTimes int    `json:"download_times"`
	Uploader      int    `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
}

type folder struct {
	GroupId        int    `json:"group_id"`
	FolderId       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int    `json:"create_time"`
	Creator        int    `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int    `json:"total_file_count"`
}

// GetStatusResp Endpoint get_status
type GetStatusResp struct {
	AppInitialized bool       `json:"app_initialized"`
	AppEnabled     bool       `json:"app_enabled"`
	PluginsGood    bool       `json:"plugins_good"`
	AppGood        bool       `json:"app_good"`
	Online         bool       `json:"online"`
	Good           bool       `json:"good"`
	Stat           statistics `json:"stat"`
}

type statistics struct {
	PacketReceived  int `json:"packet_received"`
	PacketSent      int `json:"packet_sent"`
	PacketLost      int `json:"packet_lost"`
	MessageReceived int `json:"message_received"`
	MessageSent     int `json:"message_sent"`
	DisconnectTimes int `json:"disconnect_times"`
	LostTimes       int `json:"lost_times"`
	LastMessageTime int `json:"last_message_time"`
}

// GetGroupAtAllRemainData Endpoint get_group_at_all_remain
type GetGroupAtAllRemainData struct {
	GroupId int `json:"group_id"`
}

// GetGroupAtAllRemainResp Endpoint get_group_at_all_remain
type GetGroupAtAllRemainResp struct {
	CanAtAll                 bool `json:"can_at_all"`
	RemainAtAllCountForGroup int  `json:"remain_at_all_count_for_group"`
	RemainAtAllCountForUin   int  `json:"remain_at_all_count_for_uin"`
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

// GetGroupNoticeResp Endpoint _get_group_notice
type GetGroupNoticeResp struct {
	SenderId    int           `json:"sender_id"`
	PublishTime int           `json:"publish_time"`
	Message     noticeMessage `json:"message"`
}

type noticeMessage struct {
	Text  string         `json:"text"`
	Image []noticeImages `json:"image"`
}

type noticeImages struct {
	Height string `json:"height"`
	Width  string `json:"width"`
	Id     string `json:"id"`
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

// DownloadFileResp Endpoint download_file
type DownloadFileResp struct {
	File string `json:"file"`
}

// GetOnlineClientsData Endpoint get_online_clients
type GetOnlineClientsData struct {
	NoCache bool `json:"no_cache"`
}

// GetOnlineClientsResp Endpoint get_online_clients
type GetOnlineClientsResp struct {
	Counts []device `json:"counts"`
}

type device struct {
	AppId      int    `json:"app_id"`
	DeviceName string `json:"device_name"`
	DeviceKind string `json:"device_kind"`
}

// GetGroupMsgHistoryData Endpoint get_group_msg_history
type GetGroupMsgHistoryData struct {
	MessageSeq int `json:"message_seq"`
	GroupId    int `json:"group_id"`
}

// GetGroupMsgHistoryResp Endpoint get_group_msg_history
type GetGroupMsgHistoryResp struct {
	Messages []string `json:"messages"`
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

// GetEssenceMsgListResp Endpoint get_essence_msg_list
type GetEssenceMsgListResp struct {
	SenderId     int    `json:"sender_id"`
	SenderNick   string `json:"sender_nick"`
	SenderTime   int    `json:"sender_time"`
	OperatorId   int    `json:"operator_id"`
	OperatorNick string `json:"operator_nick"`
	OperatorTime int    `json:"operator_time"`
	MessageId    int    `json:"message_id"`
}

// CheckUrlSafelyData Endpoint check_url_safely
type CheckUrlSafelyData struct {
	Url string `json:"url"`
}

// CheckUrlSafelyResp Endpoint check_url_safely
type CheckUrlSafelyResp struct {
	Level int `json:"level"`
}

// GetModelShowData Endpoint _get_model_show
type GetModelShowData struct {
	Model string `json:"model"`
}

// GetModelShowResp Endpoint _get_model_show
type GetModelShowResp struct {
	Variants []variants `json:"variants"`
}

type variants struct {
	ModelShow string `json:"model_show"`
	NeedPay   bool   `json:"need_pay"`
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

// SendPrivateForwardMsgResp Endpoint send_private_forward_msg
type SendPrivateForwardMsgResp struct {
	MessageId int `json:"message_id"`
}
