package gcqdata

// GetGuildServiceProfileResp Endpoint get_guild_service_profile
type GetGuildServiceProfileResp struct {
	Nickname  string `json:"nickname"`
	TinyID    string `json:"tiny_id"`
	AvatarURL string `json:"avatar_url"`
}

// GetGuildListResp Endpoint get_guild_list
type GetGuildListResp struct {
	GuildID        string `json:"guild_id"`
	GuildName      string `json:"guild_name"`
	GuildDisplayID int    `json:"guild_display_id"`
}

// GetGuildMetaByGuestData Endpoint get_guild_meta_by_guest
type GetGuildMetaByGuestData struct {
	GuildID string `json:"guild_id"`
}

// GetGuildMetaByGuestResp Endpoint get_guild_meta_by_guest
type GetGuildMetaByGuestResp struct {
	GuildID        string `json:"guild_id"`
	GuildName      string `json:"guild_name"`
	GuildProfile   string `json:"guild_profile"`
	CreateTime     int    `json:"create_time"`
	MaxMemberCount int    `json:"max_member_count"`
	MaxRobotCount  int    `json:"max_robot_count"`
	MaxAdminCount  int    `json:"max_admin_count"`
	MemberCount    int    `json:"member_count"`
	OwnerId        string `json:"owner_id"`
}

// GetGuildChannelListData Endpoint get_guild_channel_list
type GetGuildChannelListData struct {
	GuildID string `json:"guild_id"`
	Nocache bool   `json:"no_cache"`
}

// GetGuildChannelListResp Endpoint get_guild_channel_list
type GetGuildChannelListResp struct {
	OwnerGuildId    string      `json:"owner_guild_id"`
	ChannelId       string      `json:"channel_id"`
	ChannelType     int         `json:"channel_type"`
	ChannelName     string      `json:"channel_name"`
	CreateTime      int         `json:"create_time"`
	CreatorTinyId   string      `json:"creator_tiny_id"`
	TalkPermission  int         `json:"talk_permission"`
	VisibleType     int         `json:"visible_type"`
	CurrentSlowMode int         `json:"current_slow_mode"`
	SlowModes       []slowModes `json:"slow_modes"`
}

type slowModes struct {
	SlowModeKey    int    `json:"slow_mode_key"`
	SlowModeText   string `json:"slow_mode_text"`
	SpeakFrequency int    `json:"speak_frequency"`
	SlowModeCircle int    `json:"slow_mode_circle"`
}

// GetGuildMemberListData Endpoint get_guild_member_list
type GetGuildMemberListData struct {
	GuildID   string `json:"guild_id"`
	NextToken string `json:"next_token"`
}

// GetGuildMemberListResp Endpoint get_guild_member_list
type GetGuildMemberListResp struct {
	Members   []guildMemberInfo `json:"members"`
	Finished  bool              `json:"finished"`
	NextToken string            `json:"next_token"`
}

type guildMemberInfo struct {
	TinyID   string `json:"tiny_id"`
	Title    string `json:"title"`
	Nickname string `json:"nickname"`
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

// GetGuildMemberProfileData Endpoint get_guild_member_profile
type GetGuildMemberProfileData struct {
	GuildID string `json:"guild_id"`
	UserID  string `json:"user_id"`
}

// GetGuildMemberProfileResp Endpoint get_guild_member_profile
type GetGuildMemberProfileResp struct {
	TinyID    string     `json:"tiny_id"`
	Nickname  string     `json:"nickname"`
	AvatarUrl string     `json:"avatar_url"`
	JoinTime  int        `json:"join_time"`
	Roles     []roleInfo `json:"roles"`
}

type roleInfo struct {
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

// SendGuildChannelMessageData Endpoint send_guild_channel_msg
type SendGuildChannelMessageData struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}

// SendGuildChannelMessageResp Endpoint send_guild_channel_msg
type SendGuildChannelMessageResp struct {
	MessageID string `json:"message_id"`
}

// GetTopicChannelFeedsData Endpoint get_topic_channel_feeds
type GetTopicChannelFeedsData struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}

// DeleteGuildRoleData Endpoint delete_guild_role
type DeleteGuildRoleData struct {
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
}

// GetGuildMsgData Endpoint get_guild_msg
type GetGuildMsgData struct {
	MessageID string `json:"message_id"`
	Nocache   bool   `json:"no_cache"`
}

// GetGuildRolesData Endpoint get_guild_roles
type GetGuildRolesData struct {
	GuildID string `json:"guild_id"`
}

// SetGuildMemberRoleData Endpoint set_guild_member_role
type SetGuildMemberRoleData struct {
	GuildID string `json:"guild_id"`
	Set     bool   `json:"set"`
	RoleID  string `json:"role_id"`
	Users   string `json:"users"`
}

// UpdateGuildRoleData Endpoint update_guild_role
type UpdateGuildRoleData struct {
	GuildID     string `json:"guild_id"`
	RoleID      string `json:"role_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Independent bool   `json:"independent"`
}

// CreateGuildRoleData Endpoint create_guild_role
type CreateGuildRoleData struct {
	GuildID      string `json:"guild_id"`
	Color        string `json:"color"`
	Name         string `json:"name"`
	Independent  bool   `json:"independent"`
	InitialUsers string `json:"initial_users"`
}
