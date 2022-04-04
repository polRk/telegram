// Bot API 5.7

package api

type ChatMemberStatus string

const (
	ChatMemberStatusOwner         ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusBanned        ChatMemberStatus = "kicked"
)

// ChatMember is an information about one member of a chat.
// Currently, the following 6 types of chat members are supported:
type ChatMember struct {
	// Status is the member's status in the chat.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`

	// IsMember is True, if the user is a member of the chat at the moment of the request.
	//
	// Optional.
	IsMember bool `json:"is_member,omitempty"`

	// CanBeEdited is True, if the bot is allowed to edit administrator privileges of that user.
	//
	// Optional.
	CanBeEdited bool `json:"can_be_edited,omitempty"`

	// IsAnonymous is True, if the user's presence in the chat is hidden.
	//
	// Optional.
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// CanManageChat is True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous administrators in supergroups
	// and ignore slow mode. Implied by any other administrator privilege.
	//
	// Optional.
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// CanDeleteMessages is True, if the administrator can delete messages of other users.
	//
	// Optional.
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// CanManageVoiceChats is True, if the administrator can manage voice chats.
	//
	// Optional.
	CanManageVoiceChats bool `json:"can_manage_voice_chats,omitempty"`

	// CanRestrictMembers is True, if the administrator can restrict, ban or unban chat members.
	//
	// Optional.
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// CanPromoteMembers is True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly
	// (promoted by administrators that were appointed by the user).
	//
	// Optional.
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// CanChangeInfo is True, if the user is allowed to change the chat title, photo and other settings.
	//
	// Optional.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// CanInviteUsers is True, if the user is allowed to invite new users to the chat.
	//
	// Optional.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// CanPostMessages is True, if the administrator can post in the channel.
	// ChatTypeChannel only.
	//
	// Optional.
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages is True, if the administrator can edit messages of other users and can pin messages.
	// ChatTypeChannel only.
	//
	// Optional.
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages is True, if the user is allowed to pin messages; groups and supergroups only.
	//
	// Optional.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// CanSendMessages is True, if the user is allowed to send text messages, contacts, locations and venues.
	//
	// Optional.
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// CanSendMediaMessages is True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes.
	//
	// Optional.
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// CanSendPolls is True, if the user is allowed to send polls.
	//
	// Optional.
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages is True, if the user is allowed to send animations, games, stickers and use inline bots.
	//
	// Optional.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews is True, if the user is allowed to add web page previews to their messages.
	//
	// Optional.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// UntilDate is a date when restrictions will be lifted for this user; unix time.
	// If 0, then the user is restricted/banned forever.
	//
	// Optional.
	UntilDate int `json:"until_date,omitempty"`

	// CustomTitle is a custom title for this user.
	//
	// Optional.
	CustomTitle string `json:"custom_title,omitempty"`
}
