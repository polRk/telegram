// Bot API 5.7

package api

// Message is a Telegram message.
type Message struct {
	// MessageID is a unique message identifier inside this chat.
	MessageID int64 `json:"message_id"`

	// From is the sender of the message.
	// Empty for messages sent to channel.
	// For backward compatibility, the field contains a fake sender user in non-channel chats,
	// if the message was sent on behalf of a chat.
	//
	// Optional.
	From *User `json:"from,omitempty"`

	// SenderChat is the sender of the message, sent on behalf of a chat.
	// For example, the channel itself for channel posts, the supergroup itself for messages
	// from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group.
	// For backward compatibility, the field from contains a fake sender user in non-channel chats,
	// if the message was sent on behalf of a chat.
	//
	// Optional.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Date of the message was sent in Unix time.
	Date int `json:"date"`

	// Chat is the conversation the message belongs to.
	Chat *Chat `json:"chat"`

	// ForwardFrom for forwarded messages, sender of the original message.
	//
	// Optional.
	ForwardFrom *User `json:"forward_from,omitempty"`

	// ForwardFromChat for messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat.
	//
	// Optional.
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// ForwardFromMessageID for messages forwarded from channels, identifier of the original message in the channel.
	//
	// Optional.
	ForwardFromMessageID int `json:"forward_from_message_id,omitempty"`

	// ForwardSignature for forwarded messages that were originally sent in channels
	// or by an anonymous chat administrator, signature of the message sender if present.
	//
	// Optional.
	ForwardSignature string `json:"forward_signature,omitempty"`

	// ForwardSenderName is the sender's name for messages forwarded from users
	// who disallow adding a link to their account in forwarded messages.
	//
	// Optional.
	ForwardSenderName string `json:"forward_sender_name,omitempty"`

	// ForwardDate for forwarded messages, date the original message was sent in Unix time.
	//
	// Optional.
	ForwardDate int `json:"forward_date,omitempty"`

	// IsAutomaticForward is True, if the message is a channel post
	// that was automatically forwarded to the connected discussion group.
	//
	// Optional.
	IsAutomaticForward int `json:"is_automatic_forward,omitempty"`

	// ReplyToMessage for replies, the original message.
	// Note that the Message object in this field will not contain further reply_to_message fields
	// even if it itself is a reply.
	//
	// Optional.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// ViaBot is a bot through which the message was sent.
	//
	// Optional.
	ViaBot *User `json:"via_bot,omitempty"`

	// EditDate of the message was last edited in Unix time.
	//
	// Optional.
	EditDate int `json:"edit_date,omitempty"`

	// HasProtectedContent is True, if the message can't be forwarded.
	//
	// Optional.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// MediaGroupID is the unique identifier of a media message group this message belongs to.
	//
	// Optional.
	MediaGroupID string `json:"media_group_id,omitempty"`

	// AuthorSignature is the signature of the post author for messages in channels,
	// or the custom title of an anonymous group administrator.
	//
	// Optional.
	AuthorSignature string `json:"author_signature,omitempty"`

	// Text is for text messages, the actual UTF-8 text of the message, 0-4096 characters.
	//
	// Optional.
	Text string `json:"text,omitempty"`

	// Entities is an array of special entities like usernames, URLs, bot commands, etc. that appear in the text.
	//
	// Optional.
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Animation is an information about the animation. Message is an animation.
	// For backward compatibility, when this field is set, the document field will also be set.
	//
	// Optional.
	Animation *Animation `json:"animation,omitempty"`

	// Audio is an information about the audio file. Message is an audio.
	//
	// Optional.
	Audio *Audio `json:"audio,omitempty"`

	// Document is an information about the general file. Message is a general file.
	//
	// Optional.
	Document *Document `json:"document,omitempty"`

	// Photo is an available sizes of the photo. Message is a photo.
	//
	// Optional.
	Photo []*PhotoSize `json:"photo,omitempty"`

	// Sticker is an information about the sticker. Message is a sticker.
	//
	// Optional.
	Sticker *Sticker `json:"sticker,omitempty"`

	// Video is an information about the video. Message is a video.
	//
	// Optional.
	Video *Video `json:"video,omitempty"`

	// VideoNote is an information about the video message. Message is a video note.
	//
	// Optional.
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Voice is an information about the voice file. Message is a voice message.
	//
	// Optional.
	Voice *Voice `json:"voice,omitempty"`

	// Caption is a caption for the animation, audio, document, photo, video or voice, 0-1024 characters.
	//
	// Optional.
	Caption string `json:"caption,omitempty"`

	// CaptionEntities is an array of special entities like usernames, URLs, bot commands, etc. that appear in the caption.
	//
	// Optional.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Contact is an information about the contact. Message is a shared contact.
	//
	// Optional.
	Contact *Contact `json:"contact,omitempty"`

	// Dice is a dice with random value.
	//
	// Optional.
	Dice *Dice `json:"dice,omitempty"`

	// Game is an information about the game. Message is a game.
	//
	// Optional.
	Game *Game `json:"game,omitempty"`

	// Poll is an information about the poll. Message is a native poll.
	//
	// Optional.
	Poll *Poll `json:"poll,omitempty"`

	// Venue is an information about the venue. Message is a venue.
	// For backward compatibility, when this field is set, the location field will also be set.
	//
	// Optional.
	Venue *Venue `json:"venue,omitempty"`

	// Location is an information about the location. Message is a shared location.
	//
	// Optional.
	Location *Location `json:"location,omitempty"`

	// NewChatMembers is an array of new members that were added to the group or supergroup.
	//
	// Optional.
	NewChatMembers []*User `json:"new_chat_members,omitempty"`

	// LeftChatMember is a member was removed from the group.
	//
	// Optional.
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// NewChatTitle is new chat title.
	//
	// Optional.
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// NewChatPhoto is new chat photo.
	//
	// Optional.
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`

	// DeleteChatPhoto is a service message: the chat photo was deleted.
	//
	// Optional.
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// GroupChatCreated is a service message: the group has been created.
	//
	// Optional.
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// SuperGroupChatCreated is a service message: the supergroup has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	//
	// Optional.
	SuperGroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// ChannelChatCreated is a service message: the channel has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	//
	// Optional.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// MessageAutoDeleteTimerChanged is a service message: auto-delete timer settings changed in the chat.
	//
	// Optional.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	// MigrateToChatID is the specified identifier of the group has been migrated to a supergroup.
	//
	// Optional.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// MigrateFromChatID is the specified identifier of the supergroup has been migrated to a group.
	//
	// Optional.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// PinnedMessage is the specified message was pinned.
	//
	// Optional.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Invoice is an information about the invoice. Message is an invoice for a payment.
	//
	// Optional.
	Invoice *Invoice `json:"invoice,omitempty"`

	// SuccessfulPayment is an information about the payment. Message is a service message about a successful payment.
	//
	// Optional.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// ConnectedWebsite is the domain name of the website on which the user has logged in.
	//
	// Optional.
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// PassportData is the Telegram Passport data.
	//
	// Optional.
	PassportData *PassportData `json:"passport_data,omitempty"`

	// ProximityAlertTriggered is a service message.
	// A user in the chat triggered another user's proximity alert while sharing Live Location.
	//
	// Optional.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// VoiceChatScheduled is a service message: voice chat scheduled.
	//
	// Optional.
	VoiceChatScheduled *VoiceChatScheduled `json:"voice_chat_scheduled,omitempty"`

	// VoiceChatStarted is a service message: voice chat started.
	//
	// Optional.
	VoiceChatStarted *VoiceChatStarted `json:"voice_chat_started,omitempty"`

	// VoiceChatStarted is a service message: voice chat ended.
	//
	// Optional.
	VoiceChatEnded *VoiceChatEnded `json:"voice_chat_ended,omitempty"`

	// VoiceChatParticipantsInvited is a service message: new participants invited to a voice chat.
	//
	// Optional.
	VoiceChatParticipantsInvited *VoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`

	// ReplyMarkup is an inline keyboard attached to the message.
	// Note: login_url buttons are represented as ordinary url buttons.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *Message) Validate() bool {
	if len(m.Caption) > 1024 {
		return false
	}

	return true
}
