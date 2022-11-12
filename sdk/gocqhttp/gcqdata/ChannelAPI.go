package gcqdata

// GetGuildServiceProfileData Endpoint get_guild_service_profile
type GetGuildServiceProfileData struct {
	Nickname  string `json:"nickname"`
	TinyID    string `json:"tiny_id"`
	AvatarURL string `json:"avatar_url"`
}

// GetGuildListData Endpoint get_guild_list
type GetGuildListData struct {
	GuildID        string `json:"guild_id"`
	GuildName      string `json:"guild_name"`
	GuildDisplayID int    `json:"guild_display_id"`
}

// GetGuildMetaByGuestData Endpoint get_guild_meta_by_guest
type GetGuildMetaByGuestData struct {
	GuildID string `json:"guild_id"`
}

// GetGuildChannelListData Endpoint get_guild_channel_list
type GetGuildChannelListData struct {
	GuildID string `json:"guild_id"`
	Nocache bool   `json:"no_cache"`
}

// GetGuildMemberListData Endpoint get_guild_member_list
type GetGuildMemberListData struct {
	GuildID   string `json:"guild_id"`
	NextToken string `json:"next_token"`
}

// GetGuildMemberProfileData Endpoint get_guild_member_profile
type GetGuildMemberProfileData struct {
	GuildID string `json:"guild_id"`
	UserID  string `json:"user_id"`
}

// SendGuildChannelMessageData Endpoint send_guild_channel_msg
type SendGuildChannelMessageData struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
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
