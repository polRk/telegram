// Bot API 5.7

package telegram

import "encoding/json"

// APIResponse is a response from the Telegram API with the result
// stored raw.
type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// Error is an error containing extra information returned by the Telegram API.
type Error struct {
	Code    int
	Message string
	*ResponseParameters
}

// Error message string.
func (e Error) Error() string {
	return e.Message
}

type AllowedUpdate string

const (
	AllowedUpdateMessage           AllowedUpdate = "message"
	AllowedUpdateEditedChannelPost AllowedUpdate = "edited_channel_post"
	AllowedUpdateCallbackQuery     AllowedUpdate = "callback_query"
	AllowedUpdateChatMember        AllowedUpdate = "chat_member"
)

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	// URL is a webhook URL, may be empty if webhook is not set up.
	URL string `json:"url"`

	// HasCustomCertificate is True, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// PendingUpdateCount is a number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`

	// IPAddress is a webhook IP address.
	//
	// Optional.
	IPAddress string `json:"ip_address,omitempty"`

	// LastErrorDate is a unix time for the most recent error that happened when trying to deliver an update via webhook.
	//
	// Optional.
	LastErrorDate int `json:"last_error_date,omitempty"`

	// LastErrorMessage is an error message in human-readable format for the most recent error that happened
	// when trying to deliver an update via webhook.
	//
	// Optional.
	LastErrorMessage int `json:"last_error_message,omitempty"`

	// LastSynchronizationErrorDate is unix time of the most recent error that happened
	// when trying to synchronize available updates with Telegram datacenters.
	//
	// Optional.
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`

	// MaxConnections is the maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery.
	//
	// Optional.
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates is an array of the update types you want your bot to receive.
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

// Update is an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// UpdateID is the update's unique identifier.
	// Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you're using Webhooks,
	// since it allows you to ignore repeated updates or to restore the correct update sequence,
	// should they get out of order.
	// If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateID int64 `json:"update_id"`

	// Message is a new incoming message of any kind — text, photo, sticker, etc.
	//
	// Optional.
	Message *Message `json:"message,omitempty"`

	// EditedMessage is a new version of a message that is known to the bot and was edited.
	//
	// Optional.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost is a new incoming channel post of any kind — text, photo, sticker, etc.
	//
	// Optional.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost is a new version of a channel post that is known to the bot and was edited.
	//
	// Optional.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// InlineQuery is a new incoming inline query.
	//
	// Optional.
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// ChosenInlineResult is the result of an inline query that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	//
	// Optional.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// CallbackQuery is a new incoming callback query.
	//
	// Optional.
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// ShippingQuery is a new incoming shipping query.
	// Only for invoices with flexible price.
	//
	// Optional.
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// PreCheckoutQuery is a new incoming pre-checkout query.
	// Contains full information about checkout.
	//
	// Optional.
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Poll is a new poll state.
	// Bots receive only updates about stopped polls and polls, which are sent by the bot.
	//
	// Optional.
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer is a user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	//
	// Optional.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// MyChatMember contains information about the bot's chat member status was updated in a chat.
	// For private chats, this update is received only when the bot is blocked or unblocked by the user.
	//
	// Optional.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// ChatMember contains information about a chat member's status was updated in a chat.
	// The bot must be an administrator in the chat
	// and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	//
	// Optional.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// ChatJoinRequest is a request to join the chat has been sent.
	// The bot must have the can_invite_users administrator right in the chat to receive these updates.
	//
	// Optional.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
}

// User is a Telegram user or bot.
type User struct {
	// ID is a unique identifier for this user or bot.
	ID int64 `json:"id"`

	// IsBot is True, if use is a bot.
	IsBot bool `json:"is_bot"`

	// FirstName is a user's or bot's first name.
	FirstName string `json:"first_name"`

	// LastName is a user's or bot's last name.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// Username is a user's or bot's username.
	//
	// Optional.
	Username string `json:"username,omitempty"`

	// LanguageCode is a IETF language tag of the user's language.
	// https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// Optional.
	LanguageCode string `json:"language_code,omitempty"`

	// CanJoinGroups is True, if the bot can be invited to groups.
	// Returned only in api.GetMe().
	//
	// Optional.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages is True, if privacy mode is disabled for the bot.
	// Returned only in api.GetMe().
	//
	// Optional.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries is True, if the bot supports inline queries.
	// Returned only in api.GetMe().
	//
	// Optional.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

type ChatType string

const (
	ChatTypeSender     ChatType = "sender"
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// Chat represents a chat.
type Chat struct {
	// ID is a unique identifier for this chat.
	ID int64 `json:"id"`

	// Type of chat, can be either ChatTypePrivate, ChatTypeGroup, ChatTypeSuperGroup or ChatTypeChannel.
	Type ChatType `json:"type"`

	// Title for ChatTypeSuperGroup, ChatTypeChannel and ChatTypeGroup chats.
	//
	// Optional.
	Title string `json:"title,omitempty"`

	// Username for ChatTypePrivate, ChatTypeSuperGroup and ChatTypeChannel if available.
	//
	// Optional.
	Username string `json:"username,omitempty"`

	// FirstName of the other party in a ChatTypePrivate.
	//
	// Optional.
	FirstName string `json:"first_name,omitempty"`

	// LastName of the other party in a ChatTypePrivate.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// Photo is a chat photo.
	// Returned only in api.GetChat().
	//
	// Optional.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Bio of the other party in a private chat.
	// Returned only in api.GetChat().
	//
	// Optional.
	Bio string `json:"bio,omitempty"`

	// HasPrivateForwards is True, if privacy settings of the other party in the private chat
	// allows to use tg://user?id=<user_id> links only in chats with the user.
	// Returned only in api.GetChat().
	//
	// Optional.
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`

	// Description for ChatTypeGroup, ChatTypeSuperGroup and ChatTypeChannel chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	Description string `json:"description,omitempty"`

	// InviteLink is a primary invite link, for ChatTypeGroup, ChatTypeSuperGroup and ChatTypeChannel chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage is the most recent pinned message.
	// Returned only in api.GetChat().
	//
	// Optional.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions are default chat member permissions, for ChatTypeGroup and ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay is the minimum allowed delay between consecutive messages sent by each unprivileged user.
	// For ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// MessageAutoDeleteTime is the time in seconds
	// after which all messages sent to the chat will be automatically deleted.
	// Returned only in api.GetChat().
	//
	// Optional.
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`

	// HasProtectedContent is True, if messages from the chat can't be forwarded to other chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	HasProtectedContent int `json:"has_protected_content,omitempty"`

	// StickerSetName ia a name of group sticker set.
	// Returned only in api.GetChat().
	//
	// Optional.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet is True, if the bot can change the group sticker set.
	// Returned only in api.GetChat().
	//
	// Optional.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// LinkedChatID is a unique identifier for the linked chat,
	// i.e. the discussion group identifier for a channel and vice versa.
	// For ChatTypeSuperGroup and ChatTypeChannel.
	// Returned only in api.GetChat().
	//
	// Optional.
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Location is the location to which the ChatTypeSuperGroup is connected.
	// For ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	Location *ChatLocation `json:"location,omitempty"`
}

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

	// VideoChatScheduled is a service message: video chat scheduled.
	//
	// Optional.
	VideoChatScheduled *VoiceChatScheduled `json:"video_chat_scheduled,omitempty"`

	// VideoChatStarted is a service message: video chat started.
	//
	// Optional.
	VideoChatStarted *VoiceChatStarted `json:"video_chat_started,omitempty"`

	// VideoChatStarted is a service message: video chat ended.
	//
	// Optional.
	VideoChatEnded *VoiceChatEnded `json:"video_chat_ended,omitempty"`

	// VideoChatParticipantsInvited is a service message: new participants invited to a video chat.
	//
	// Optional.
	VideoChatParticipantsInvited *VoiceChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`

	// WebAppData is a service message: data sent by a Web App.
	//
	// Optional.
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// ReplyMarkup is an inline keyboard attached to the message.
	// Note: login_url buttons are represented as ordinary url buttons.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageID represents a unique message identifier.
type MessageID struct {
	// MessageID is a unique message identifier.
	MessageID int64 `json:"message_id"`
}

type MessageEntityType string

const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashTag       MessageEntityType = "hashtag"
	MessageEntityTypeCashTag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeSpoiler       MessageEntityType = "spoiler"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

// MessageEntity is one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type is a type of the entity.
	// Currently, can be:
	// MessageEntityTypeMention (@username),
	// MessageEntityTypeHashTag (#hashtag),
	// MessageEntityTypeCashTag ($USD),
	// MessageEntityTypeBotCommand (/start@jobs_bot),
	// MessageEntityTypeURL (https://telegram.org),
	// MessageEntityTypeEmail (do-not-reply@telegram.org),
	// MessageEntityTypePhoneNumber (+1-212-555-0123),
	// MessageEntityTypeBold (bold text),
	// MessageEntityTypeItalic (italic text),
	// MessageEntityTypeUnderline (underlined text),
	// MessageEntityTypeStrikethrough (strikethrough text),
	// MessageEntityTypeSpoiler (spoiler message),
	// MessageEntityTypeCode (monowidth string),
	// MessageEntityTypePre (monowidth block),
	// MessageEntityTypeTextLink (for clickable text URLs),
	// MessageEntityTypeTextMention (for users without usernames).
	Type MessageEntityType `json:"type"`

	// Offset is an offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`

	// Length is a length of the entity in UTF-16 code units.
	Length int `json:"length"`

	// URL is an url that will be opened after user taps on the text.
	// For MessageEntityTypeTextLink only.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// User is the mentioned user.
	// For MessageEntityTypeTextMention only.
	//
	// Optional.
	User *User `json:"user,omitempty"`

	// Language is the programming language of the entity text.
	// For MessageEntityTypePre only.
	//
	// Optional.
	Language string `json:"language,omitempty"`
}

// PhotoSize is one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a photo width.
	Width int `json:"width"`

	// Height is a photo height.
	Height int `json:"height"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// Animation is an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a video width as defined by sender.
	Width int `json:"width"`

	// Height is a video height as defined by sender.
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb is an animation thumbnail as defined by sender.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is original animation filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// Audio is an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// Performer of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Performer string `json:"performer,omitempty"`

	// Title of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Title string `json:"title,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is MIME type of the file as defined by sender
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`

	// Thumb is a thumbnail of the album cover to which the music file belongs.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Document is a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Thumb is a document thumbnail as defined by sender.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// Video is a video file.
type Video struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a video width as defined by sender.
	Width int `json:"width"`

	// Height is a video height as defined by sender.
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb is a video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is Mime type of file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// VideoNote is a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Length is a video width and height (diameter of the video message) as defined by sender.
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb is a video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// Voice is a voice note.
type Voice struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// MimeType is MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// Contact is a phone contact.
type Contact struct {
	// PhoneNumber is a Contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// FirstName is a Contact's first name.
	FirstName string `json:"first_name"`

	// LastName is a Contact's last name.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// UserID is a Contact's user identifier in Telegram.
	//
	// Optional.
	UserID int64 `json:"user_id,omitempty"`

	// VCard is an additional data about the contact in the form of a vCard.
	//
	// Optional.
	VCard string `json:"v_card,omitempty"`
}

// Dice is an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based.
	Emoji string `json:"emoji"`

	// Value of the dice,
	// 1-6 for “🎲”,
	// “🎯” and “🎳” base emoji,
	// 1-5 for “🏀” and “⚽” base emoji,
	// 1-64 for “🎰” base emoji.
	Value int `json:"value"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Text is an option text, 1-100 characters.
	Text string `json:"text"`

	// VoterCount is a number of users that voted for this option.
	VoterCount int `json:"voter_count"`
}

// PollAnswer is an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID is a unique poll identifier.
	PollID string `json:"poll_id"`

	// User who changed the answer to the poll.
	User *User `json:"user"`

	// OptionIDs are 0-based identifiers of answer options, chosen by the user.
	// May be empty if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}

type PollType string

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

// Poll contains information about a poll.
type Poll struct {
	// ID is a unique poll identifier.
	ID string `json:"id"`

	// Question is a Poll question, 1-300 characters.
	Question string `json:"question"`

	// Options is list of poll options.
	Options []*PollOption `json:"options"`

	// TotalVoterCount is a total number of users that voted in the poll.
	TotalVoterCount int `json:"total_voter_count"`

	// IsClosed is True, if the poll is closed.
	IsClosed bool `json:"is_closed"`

	// IsAnonymous is True, if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`

	// Type is a PollType, currently can be PollTypeRegular or PollTypeQuiz
	Type PollType `json:"type"`

	// AllowsMultipleAnswers is True, if the poll allows multiple answers.
	//
	// Optional.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// CorrectOptionID is a 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode,
	// which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	//
	// Optional.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation is a text that is shown when a user chooses an incorrect answer or taps on the lamp icon
	// in a quiz-style poll, 0-200 characters.
	//
	// Optional.
	Explanation string `json:"explanation,omitempty"`

	// ExplanationEntities is a special entities like usernames, URLs, bot commands, etc. that appear in the explanation.
	//
	// Optional.
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod is an amount of time in seconds the poll will be active after creation.
	//
	// Optional.
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate is a point in time (Unix timestamp) when the poll will be automatically closed.
	//
	// Optional.
	CloseDate int `json:"close_date,omitempty"`
}

// Location is a point on the map.
type Location struct {
	// Longitude
	Longitude float64 `json:"longitude"`

	// Latitude
	Latitude float64 `json:"latitude"`

	// HorizontalAccuracy is the radius of uncertainty for the location, measured in meters.
	//
	// Optional.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod is the time in seconds relative to the message sending date, during which the location can be updated.
	// For active live locations only.
	//
	// Optional.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Heading is the direction in degrees in which user is moving.
	// For active live locations only.
	//
	// Optional.
	Heading int64 `json:"heading,omitempty"`

	// ProximityAlertRadius is a maximum distance in meters for proximity alerts about approaching another chat member.
	//
	// Optional.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}

// Venue is a venue.
type Venue struct {
	// Location is a venue location.
	// Can't be a live location.
	Location *Location `json:"location"`

	// Title is a name of the venue.
	Title string `json:"title"`

	// Address of the venue.
	Address string `json:"address"`

	// Foursquare identifier of the venue.
	//
	// Optional.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType is type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	//
	// Optional.
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID is a Google Places identifier of the venue.
	//
	// Optional.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType is a Google Places type of the venue.
	// See: https://developers.google.com/maps/documentation/places/web-service/supported_types
	// Optional.
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// WebAppData is a data sent from a Web App to the bot.
type WebAppData struct {
	// Data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ButtonText is a text of the web_app keyboard button, from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// ProximityAlertTriggered is the content of a service message,
// sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler is the User that triggered the alert.
	Traveler *User `json:"traveler"`

	// Watcher is the User that set the alert.
	Watcher *User `json:"watcher"`

	// Distance between the users.
	Distance int `json:"distance"`
}

// MessageAutoDeleteTimerChanged is a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// MessageAutoDeleteTime is a new auto-delete time in seconds for messages in the chat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VoiceChatScheduled is a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	// StartDate is a point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator.
	StartDate int `json:"start_date"`
}

// VoiceChatStarted is a service message about a voice chat started in the chat.
// Currently holds no information.
type VoiceChatStarted struct {
}

// VoiceChatEnded is a service message about a voice chat ended in the chat.
type VoiceChatEnded struct {
	// Duration is a voice chat duration in seconds.
	Duration int `json:"duration"`
}

// VoiceChatParticipantsInvited is a service message about new members invited to a voice chat.
type VoiceChatParticipantsInvited struct {
	// Users is an array of new members that were invited to the voice chat.
	//
	// Optional.
	Users []*User `json:"users"`
}

// UserProfilePhotos is a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount is a total number of profile pictures the target user has.
	TotalCount int `json:"total_count"`

	// Photos is an array of requested profile pictures (in up to 4 sizes each)
	Photos []*PhotoSize `json:"photos"`
}

// File is a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize in bytes, if known.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`

	// FilePath, Use https://api.telegram.org/file/bot<token>/<file_path>.
	//
	// Optional.
	FilePath string `json:"file_path,omitempty"`
}

// WebAppInfo is an information about a Web App.
type WebAppInfo struct {
	// URL is an HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps.
	// https://core.telegram.org/bots/webapps#initializing-web-apps
	URL string `json:"url"`
}

// ReplyKeyboardMarkup is a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an Array of KeyboardButton objects.
	Keyboard [][]*KeyboardButton `json:"keyboard"`

	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit.
	// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	//
	// Optional.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat -
	// the user can press a special button in the input field to see the custom keyboard again.
	// Defaults to false.
	//
	// Optional.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active.
	// 1-64 characters.
	//
	// Optional.
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective is showing the keyboard to specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// KeyboardButton is one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields RequestContact, RequestLocation, and RequestPoll are mutually exclusive.
type KeyboardButton struct {
	// Text of the button.
	// If none of the optional fields are used, it will be sent as a message when the button is pressed.
	Text string `json:"text"`

	// RequestContact If True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in ChatTypePrivate.
	//
	// Optional.
	RequestContact bool `json:"request_contact,omitempty"`

	// RequestLocation if True, the user's current location will be sent when the button is pressed.
	// Available in ChatTypePrivate.
	//
	// Optional.
	RequestLocation bool `json:"request_location,omitempty"`

	// RequestPoll if specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in ChatTypePrivate.
	//
	// Optional.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// WebApp described Web App will be launched when the button is pressed.
	// The Web App will be able to send a “web_app_data” service message.
	// Available in ChatTypePrivate.
	//
	// Optional.
	WebApp *WebAppInfo `json:"web_app"`
}

// KeyboardButtonPollType is a type of poll,
// which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type of allowed to create a poll.
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	//
	// Optional.
	Type PollType `json:"type,omitempty"`
}

// ReplyKeyboardRemove remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard requests clients to remove the custom keyboard.
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Selective is removing the keyboard for specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//
	// Optional.
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an Array of InlineKeyboardButton objects.
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text is a label text on the button.
	Text string `json:"text"`
	// URL is HTTP or tg:// url to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username,
	// if this is allowed by their privacy settings.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// LoginURL is An HTTP URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	//
	// Optional.
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// CallbackData is data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	//
	// Optional.
	CallbackData string `json:"callback_data,omitempty"`

	// WebApp is a description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in ChatTypePrivate between a user and the bot.
	//
	// Optional.
	WebApp *WebAppInfo `json:"web_app"`

	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	//
	// Optional.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat if set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	//
	// Optional.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// CallbackGame is a description of the game that will be launched when the user presses the button.
	//
	// Optional.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay is a flag if True, send a Pay button.
	//
	// NOTE: This type of button must always be the first button in the first row
	// and can only be used in invoice messages.
	//
	// Optional.
	Pay bool `json:"pay,omitempty"`
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	// URL is an HTTP URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data,
	// the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	// NOTE: You must always check the hash of the received data
	// to verify the authentication and the integrity of the data as described in Checking authorization.
	URL string `json:"url"`

	// ForwardText is a new text of the button in forwarded messages.
	//
	// Optional.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername is a username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details.
	// If not specified, the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	//
	// Optional.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess pass True to request the permission for your bot to send messages to the user.
	//
	// Optional.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// CallbackQuery is an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	// ID is a unique identifier for this query.
	ID string `json:"id"`

	// From is a Sender.
	From *User `json:"from"`

	// Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old.
	//
	// Optional.
	Message *Message `json:"message,omitempty"`

	// InlineMessageID is an identifier of the message sent via the bot in inline mode, that originated the query.
	//
	// Optional.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ChatInstance a global identifier,
	// uniquely corresponding to the chat to which the message with the callback button was sent.
	// Useful for high scores in games.
	//
	// Optional.
	ChatInstance string `json:"chat_instance"`

	// Data associated with the callback button.
	// Be aware that a bad client can send arbitrary data in this field.
	//
	// Optional.
	Data string `json:"data,omitempty"`

	// GameShortName is a short name of a Game to be returned, serves as the unique identifier for the game.
	//
	// Optional.
	GameShortName string `json:"game_short_name,omitempty"`
}

// ForceReply showing a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces
// without having to sacrifice privacy mode.
type ForceReply struct {
	// ForceReply shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'.
	ForceReply bool `json:"force_reply"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the reply is active; 1-64 characters.
	//
	// Optional
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective is force reply from specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// ChatPhoto is Telegram chat photo.
type ChatPhoto struct {
	// SmallFileID is a file identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// SmallFileUniqueID is a unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// BigFileID is a file identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// BigFileUniqueID is a unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatInviteLink is an invitation link for a chat.
type ChatInviteLink struct {
	// InviteLink.
	// If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator of the link.
	Creator *User `json:"creator"`

	// CreatesJoinRequest is True, if users joining the chat via the link need to be approved by chat administrators.
	CreatesJoinRequest bool `json:"creates_join_request"`

	// IsPrimary is True, if the link is primary.
	IsPrimary bool `json:"is_primary"`

	// IsRevoked is True, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`

	// Name is an invitation link name.
	//
	// Optional.
	Name string `json:"name,omitempty"`

	// ExpireDate is a point in time (Unix timestamp) when the link will expire or has been expired.
	//
	// Optional.
	ExpireDate int `json:"expire_date,omitempty"`

	// MemberLimit is a maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999.
	//
	// Optional.
	MemberLimit int `json:"member_limit,omitempty"`

	// PendingJoinRequestCount is a number of pending join requests created using this link.
	//
	// Optional.
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

// ChatAdministratorRights is the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// IsAnonymous is True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat is True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members,
	// see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages is True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats is True, if the administrator can manage video chats.
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers is True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers is True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators
	// that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user).
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo is True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers is True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostMessages is True, if the administrator can post in the channel;
	// ChatTypeChannel only.
	//
	// Optional.
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages is True, if the administrator can edit messages of other users and can pin messages;
	// ChatTypeChannel only.
	//
	// Optional.
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages is True, if the user is allowed to pin messages;
	// ChatTypeGroup and ChatTypeSuperGroup only.
	//
	// Optional.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

// ChatMember is an information about one member of a chat.
// Currently, the following 6 types of chat members are supported:
// ChatMemberOwner,
// ChatMemberAdministrator,
// ChatMemberMember,
// ChatMemberRestricted,
// ChatMemberLeft,
// ChatMemberBanned.
type ChatMember struct {
	ChatMemberOwner
	ChatMemberAdministrator
	ChatMemberMember
	ChatMemberRestricted
	ChatMemberLeft
	ChatMemberBanned
	ChatMemberUpdated
}

// ChatMemberOwner is a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	// Status is the member's status in the chat, always ChatMemberStatusCreator.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`

	// IsAnonymous is True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`

	// CustomTitle is a custom title for this user.
	//
	// Optional.
	CustomTitle string `json:"custom_title,omitempty"`
}

// ChatMemberAdministrator is a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	// Status is the member's status in the chat, always ChatMemberStatusAdministrator.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`

	// CanBeEdited is True, if the bot is allowed to edit administrator privileges of that user.
	CanBeEdited bool `json:"can_be_edited"`

	// IsAnonymous is True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat is True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous administrators in supergroups
	// and ignore slow mode. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages is True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats is True, if the administrator can manage video chats.
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers is True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers is True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly
	// (promoted by administrators that were appointed by the user).
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo is True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers is True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostMessages is True, if the administrator can post in the channel.
	// ChatTypeChannel only.
	CanPostMessages bool `json:"can_post_messages"`

	// CanEditMessages is True, if the administrator can edit messages of other users and can pin messages.
	// ChatTypeChannel only.
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages is True, if the user is allowed to pin messages.
	// ChatTypeGroup, ChatTypeSuperGroup only.
	//
	CanPinMessages bool `json:"can_pin_messages"`

	// CustomTitle is a custom title for this user.
	// Optional.
	CustomTitle string `json:"custom_title,omitempty"`
}

// ChatMemberMember is a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	// Status is the member's status in the chat, always ChatMemberStatusMember.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`
}

// ChatMemberRestricted is a chat member that is under certain restrictions in the chat.
// ChatTypeSuperGroup only.
type ChatMemberRestricted struct {
	// Status is the member's status in the chat, always ChatMemberStatusRestricted.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`

	// IsMember is True, if the user is a member of the chat at the moment of the request.
	IsMember bool `json:"is_member"`

	// CanChangeInfo is True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers is True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPinMessages is True, if the user is allowed to pin messages.
	CanPinMessages bool `json:"can_pin_messages"`

	// CanSendMessages is True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages"`

	// CanSendMediaMessages is True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes.
	CanSendMediaMessages bool `json:"can_send_media_messages"`

	// CanSendPolls is True, if the user is allowed to send polls.
	CanSendPolls bool `json:"can_send_polls"`

	// CanSendOtherMessages is True, if the user is allowed to send animations, games, stickers and use inline bots.
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// CanAddWebPagePreviews is True, if the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// UntilDate is a date when restrictions will be lifted for this user; unix time.
	// If 0, then the user is restricted forever.
	UntilDate int `json:"until_date"`
}

// ChatMemberLeft is a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	// Status is the member's status in the chat, always ChatMemberStatusLeft.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`
}

// ChatMemberBanned is a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	// Status is the member's status in the chat, always ChatMemberStatusKicked.
	Status ChatMemberStatus `json:"status"`

	// User is an information about the user.
	User *User `json:"user"`

	// UntilDate is a date when restrictions will be lifted for this user; unix time.
	// If 0, then the user is banned forever.
	UntilDate int `json:"until_date"`
}

// ChatMemberUpdated is a changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to.
	Chat *Chat `json:"chat"`

	// From is a performer of the action, which resulted in the change.
	From *User `json:"from"`

	// Date the change was done in Unix time
	Date int `json:"date"`

	// OldChatMember is previous information about the chat member.
	OldChatMember *ChatMember `json:"old_chat_member"`

	// NewChatMember is new information about the chat member.
	NewChatMember *ChatMember `json:"new_chat_member"`

	// InviteLink is a chat invite link, which was used by the user to join the chat.
	// For joining by invite link events only.
	//
	// Optional.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatJoinRequest is a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent.
	Chat *Chat `json:"chat"`

	// User that sent the join request.
	User *User `json:"user"`

	// Date the request was sent in Unix time.
	Date int `json:"date"`

	// Bio of the user.
	//
	// Optional.
	Bio string `json:"bio,omitempty"`

	// InviteLink is a chat invite link that was used by the user to send the join request.
	//
	// Optional.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages is True, if the user is allowed to send text messages, contacts, locations and venues.
	//
	// Optional.
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// CanSendMediaMessages is True, if the user is allowed to send audios, documents, photos, videos, video notes
	// and voice notes, implies can_send_messages.
	//
	// Optional.
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// CanSendPolls is True, if the user is allowed to send polls, implies can_send_messages.
	//
	// Optional.
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages is True, if the user is allowed to send animations, games, stickers
	// and use inline bots, implies can_send_media_messages.
	//
	// Optional.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews is True, if the user is allowed to add web page previews to their messages,
	// implies can_send_media_messages.
	//
	// Optional.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// CanChangeInfo is True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	//
	// Optional.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// CanInviteUsers is True, if the user is allowed to invite new users to the chat.
	//
	// Optional.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// CanPinMessages is True, if the user is allowed to pin messages.
	// Ignored in public supergroups.
	//
	// Optional.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation is a location to which a chat is connected.
type ChatLocation struct {
	// Location is the location to which the supergroup is connected.
	// Can't be a live location.
	Location *Location

	// Address is the location address, defined by the chat owner.
	Address string `json:"address"`
}

// BotCommand is a bot command.
type BotCommand struct {
	// Command is a text of the command; 1-32 characters.
	// Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`

	// Description of the command; 1-256 characters.
	Description string `json:"description"`
}

type BotCommandScopeType string

const (
	BotCommandScopeTypeDefault               BotCommandScopeType = "default"
	BotCommandScopeTypeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	BotCommandScopeTypeAllGroupChats         BotCommandScopeType = "all_group_chats"
	BotCommandScopeTypeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	BotCommandScopeTypeChat                  BotCommandScopeType = "chat"
	BotCommandScopeTypeChatAdministrators    BotCommandScopeType = "chat_administrators"
	BotCommandScopeTypeChatMember            BotCommandScopeType = "chat_member"
)

// BotCommandScope is the scope to which bot commands are applied.
// Currently, the following 7 scopes are supported:
// BotCommandScopeDefault,
// BotCommandScopeAllPrivateChats,
// BotCommandScopeAllGroupChats,
// BotCommandScopeAllChatAdministrators,
// BotCommandScopeChat,
// BotCommandScopeChatAdministrators,
// BotCommandScopeChatMember.
type BotCommandScope struct {
	BotCommandScopeDefault
	BotCommandScopeAllPrivateChats
	BotCommandScopeAllGroupChats
	BotCommandScopeAllChatAdministrators
	BotCommandScopeChat
	BotCommandScopeChatAdministrators
	BotCommandScopeChatMember
}

// BotCommandScopeDefault is the default scope of bot commands.
// Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	// Type is a scope type, must be BotCommandScopeTypeDefault.
	Type BotCommandScopeType `json:"type"`
}

// BotCommandScopeAllPrivateChats is the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Type is a scope type, must be BotCommandScopeTypeAllPrivateChats.
	Type BotCommandScopeType `json:"type"`
}

// BotCommandScopeAllGroupChats is the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Type is a scope type, must be BotCommandScopeTypeAllGroupChats.
	Type BotCommandScopeType `json:"type"`
}

// BotCommandScopeAllChatAdministrators is the scope of bot commands,
// covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Type is a scope type, must be BotCommandScopeTypeAllChatAdministrators.
	Type BotCommandScopeType `json:"type"`
}

// BotCommandScopeChat is the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Type is a scope type, must be BotCommandScopeTypeChat.
	Type BotCommandScopeType `json:"type"`

	// ChatID is a unique identifier for the target chat or username of the target supergroup
	// (in the format @supergroupusername).
	ChatID int64 `json:"chat_id"`
}

// BotCommandScopeChatAdministrators is the scope of bot commands,
// covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Type is a scope type, must be BotCommandScopeTypeChatAdministrators.
	Type BotCommandScopeType `json:"type"`

	// ChatID is a unique identifier for the target chat or username of the target supergroup
	// (in the format @supergroupusername).
	ChatID int64 `json:"chat_id"`
}

// BotCommandScopeChatMember is the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	// Type is a scope type, must be BotCommandScopeTypeChatMember.
	Type BotCommandScopeType `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID int64 `json:"chat_id"`

	// ChatID is a unique identifier of the target user.
	UserID int64 `json:"user_id"`
}

type MenuButtonType string

const (
	MenuButtonTypeCommands MenuButtonType = "commands"
	MenuButtonTypeWebApp   MenuButtonType = "web_app"
	MenuButtonTypeDefault  MenuButtonType = "default"
)

// MenuButton describes the bot's menu button in a private chat.
// It should be one of:
// MenuButtonCommands,
// MenuButtonWebApp,
// MenuButtonDefault.
//
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat.
// Otherwise the default menu button is applied.
// By default, the menu button opens the list of bot commands.
type MenuButton struct {
	// Type of the button.
	Type MenuButtonType `json:"type"`

	// Text on the button.
	Text string `json:"text,omitempty"`

	// WebApp is a description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// MenuButtonCommands is a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	// Type of the button, must be MenuButtonTypeCommands.
	Type MenuButtonType `json:"type"`
}

// MenuButtonWebApp is a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	// Type of the button, must be MenuButtonTypeWebApp.
	Type MenuButtonType `json:"type"`

	// Text on the button.
	Text string `json:"text"`

	// WebApp is a description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	WebApp *WebAppInfo `json:"web_app"`
}

// MenuButtonDefault is that no specific value for the menu button was set.
type MenuButtonDefault struct {
	// Type of the button, must be MenuButtonTypeDefault.
	Type MenuButtonType `json:"type"`
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	//
	// Optional.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// In case of exceeding flood control, the number of seconds left to wait
	// before the request can be repeated.
	//
	// Optional.
	RetryAfter int `json:"retry_after,omitempty"`
}

// Sticker is a sticker.
type Sticker struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a sticker width.
	Width int `json:"width"`

	// Height is a sticker height.
	Height int `json:"height"`

	// IsAnimated is True, if the sticker is animated.
	IsAnimated bool `json:"is_animated"`

	// IsVideo is True, if the sticker is video sticker.
	IsVideo bool `json:"is_video"`

	// Thumb is a sticker thumbnail in the .WEBP or .JPG format.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Emoji associated with the sticker.
	//
	// Optional.
	Emoji string `json:"emoji,omitempty"`

	// SetName is a name of the sticker set to which the sticker belongs.
	//
	// Optional.
	SetName string `json:"set_name,omitempty"`

	// MaskPosition is the position where the mask should be placed.
	// For mask stickers.
	//
	// Optional.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}

// StickerSet is a sticker set.
type StickerSet struct {
	// Name of sticker set.
	Name string `json:"name"`

	// Title of sticker set.
	Title string `json:"title"`

	// IsAnimated is True, if the sticker set contains animated stickers.
	IsAnimated bool `json:"is_animated"`

	// IsVideo is True, if the sticker set contains video stickers.
	IsVideo bool `json:"is_video"`

	// ContainsMasks is True, if the sticker set contains masks.
	ContainsMasks bool `json:"contains_masks"`

	// Stickers is an array of all set stickers.
	Stickers []*Sticker `json:"stickers"`

	// Thumb is a thumbnail of sticker set in the .WEBP, .TGS, or .WEBM format.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

type MaskPositionPoint = string

const (
	MaskPositionPointForeHead MaskPositionPoint = "forehead"
	MaskPositionPointEyes     MaskPositionPoint = "eyes"
	MaskPositionPointMouth    MaskPositionPoint = "mouth"
	MaskPositionPointChin     MaskPositionPoint = "chin"
)

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// Point is the part of the face relative to which the mask should be placed.
	// One of MaskPositionPointForeHead, MaskPositionPointEyes, MaskPositionPointMouth, or MaskPositionPointChin.
	Point MaskPositionPoint `json:"point"`

	// XShift is a shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`

	// YShift is a shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`

	// Scale is a mask scaling coefficient.
	// For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// InlineQuery is an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// ID is a unique identifier for this query.
	ID string `json:"id"`

	// From is a Sender.
	From *User `json:"from"`

	// Query is a text of the query (up to 256 characters).
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot.
	Offset string `json:"offset"`

	// ChatType is a type of the chat, from which the inline query was sent.
	// Can be either ChatTypeSender for a private chat with the inline query sender,
	// ChatTypePrivate, ChatTypeGroup, ChatTypeSuperGroup, or ChatTypeChannel.
	// The chat type should be always known for requests sent from official clients
	// and most third-party clients, unless the request was sent from a secret chat.
	//
	// Optional.
	ChatType ChatType `json:"chat_type,omitempty"`

	// Location is a sender location, only for bots that request user location.
	//
	// Optional.
	Location *Location `json:"location,omitempty"`
}

type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGif      InlineQueryResultType = "gif"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

// InlineQueryResult is the result of an inline query.
// Telegram clients currently support results of the following 20 types:
// InlineQueryResultCachedAudio,
// InlineQueryResultCachedDocument,
// InlineQueryResultCachedGif,
// InlineQueryResultCachedMpeg4Gif,
// InlineQueryResultCachedPhoto,
// InlineQueryResultCachedSticker,
// InlineQueryResultCachedVideo,
// InlineQueryResultCachedVoice,
// InlineQueryResultArticle,
// InlineQueryResultAudio,
// InlineQueryResultContact,
// InlineQueryResultGame,
// InlineQueryResultDocument,
// InlineQueryResultGif,
// InlineQueryResultLocation,
// InlineQueryResultMpeg4Gif,
// InlineQueryResultPhoto,
// InlineQueryResultVenue,
// InlineQueryResultVideo,
// InlineQueryResultVoice.
type InlineQueryResult struct {
	InlineQueryResultArticle
	InlineQueryResultPhoto
	InlineQueryResultGif
	InlineQueryResultMpeg4Gif
	InlineQueryResultVideo
	InlineQueryResultAudio
	InlineQueryResultVoice
	InlineQueryResultDocument
	InlineQueryResultLocation
	InlineQueryResultVenue
	InlineQueryResultContact
	InlineQueryResultGame
	InlineQueryResultCachedPhoto
	InlineQueryResultCachedGif
	InlineQueryResultCachedMpeg4Gif
	InlineQueryResultCachedSticker
	InlineQueryResultCachedDocument
	InlineQueryResultCachedVideo
	InlineQueryResultCachedAudio
	InlineQueryResultCachedVoice
}

type InlineQueryResultArticle[MC InputMessageContent] struct {
	// Type of the result, must be InlineQueryResultTypeArticle.
	Type InlineQueryResultType `json:"type"`

	// ID is a unique identifier for this result, 1-64 Bytes.
	ID string `json:"id"`

	// Title of the result.
	Title string `json:"title"`

	// InputMessageContent is a content of the message to be sent.
	InputMessageContent *MC `json:"input_message_content"`

	// ReplyMarkup is an InlineKeyboardMarkup attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// URL of the result.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// HideURL is used to hide the URL to be shown in the message.
	//
	// Optional.
	HideURL bool `json:"hide_url,omitempty"`

	//  Description is a short description of the result.
	//
	// Optional.
	Description string `json:"description,omitempty"`

	// ThumbURL of the thumbnail for the result
	//
	// Optional.
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth is a width of the Thumbnail.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight is a height of the Thumbnail.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto is a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto[MC InputMessageContent] struct {
	// Type of the result, must be InlineQueryResultTypePhoto.
	Type InlineQueryResultType `json:"type"`

	// ID is a unique identifier for this result, 1-64 Bytes.
	ID string `json:"id"`

	// PhotoURL is a valid URL of the photo.
	// Photo must be in JPEG format.
	// Photo size must not exceed 5MB.
	PhotoURL string `json:"photo_url"`

	// ThumbURL is a URL of the thumbnail for the photo.
	ThumbURL string `json:"thumb_url"`

	// PhotoWidth is a width of the photo.
	//
	// Optional.
	PhotoWidth int `json:"photo_width,omitempty"`

	// PhotoHeight is a height of the photo.
	//
	// Optional.
	PhotoHeight int `json:"photo_height,omitempty"`

	// Title of the result.
	//
	// Optional.
	Title string `json:"title,omitempty"`

	//  Description is a short description of the result.
	//
	// Optional.
	Description string `json:"description,omitempty"`

	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption,omitempty"`

	// ParseMode is a mode for parsing entities in the photo caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	//
	// Optional.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// CaptionEntities is an array of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup is an InlineKeyboardMarkup attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent is a content of the message to be sent instead of the photo.
	//
	// Optional.
	InputMessageContent *MC `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif
// TODO: https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	// Type of the result, must be InlineQueryResultTypeGif.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultMpeg4Gif
// TODO: https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be InlineQueryResultTypeMpeg4Gif.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultVideo
// TODO: https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	// Type of the result, must be InlineQueryResultTypeVideo.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultAudio
// TODO: https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	// Type of the result, must be InlineQueryResultTypeAudio.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultVoice
// TODO: https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	// Type of the result, must be InlineQueryResultTypeVoice.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultDocument
// TODO: https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	// Type of the result, must be InlineQueryResultTypeDocument.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultLocation
// TODO: https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	// Type of the result, must be InlineQueryResultTypeLocation.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultVenue
// TODO: https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	// Type of the result, must be InlineQueryResultTypeVenue.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultContact
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	// Type of the result, must be InlineQueryResultTypeContact.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultGame
// TODO: https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	// Type of the result, must be InlineQueryResultTypeGame.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedPhoto
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be InlineQueryResultTypePhoto.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedGif
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	// Type of the result, must be InlineQueryResultTypeGif.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedMpeg4Gif
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be InlineQueryResultTypeMpeg4Gif.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedSticker
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be InlineQueryResultTypeSticker.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedDocument
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	// Type of the result, must be InlineQueryResultTypeDocument.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedVideo
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	// Type of the result, must be InlineQueryResultTypeVideo.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedAudio
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	// Type of the result, must be InlineQueryResultTypeAudio.
	Type InlineQueryResultType `json:"type"`
}

// InlineQueryResultCachedVoice
// TODO: https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	// Type of the result, must be InlineQueryResultTypeVoice.
	Type InlineQueryResultType `json:"type"`
}

// InputMessageContent is the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 5 types:
// InputTextMessageContent,
// InputLocationMessageContent,
// InputVenueMessageContent,
// InputContactMessageContent,
// InputInvoiceMessageContent.
type InputMessageContent interface {
	InputTextMessageContent |
		InputLocationMessageContent |
		InputVenueMessageContent |
		InputContactMessageContent |
		InputInvoiceMessageContent
}

// InputTextMessageContent
// TODO: https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct{}

// InputLocationMessageContent
// TODO: https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct{}

// InputVenueMessageContent
// TODO: https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct{}

// InputContactMessageContent
// TODO: https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct{}

// InputInvoiceMessageContent
// TODO: https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputInvoiceMessageContent struct{}

// ChosenInlineResult is a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ResultID is the unique identifier for the result that was chosen.
	ResultID string `json:"result_id"`

	// From is the user that chose the result.
	From *User `json:"from"`

	// Location is a sender location, only for bots that require user location.
	//
	// Optional.
	Location *Location `json:"location,omitempty"`

	// InlineMessageID is an identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	//
	// Optional.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Query is the query that was used to obtain the result.
	Query string `json:"query"`
}

// SentWebAppMessage is an information about an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	// InlineMessageID is an identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	//
	// Optional.
	InlineMessageID string `json:"inline_message_id"`
}

// LabeledPrice is a portion of the price for goods or services.
type LabeledPrice struct {
	// Label of the portion.
	Label string `json:"label"`

	// Amount is a price of the product in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	// Title is a product name.
	Title string `json:"title"`

	// Description is a product description.
	Description string `json:"description"`

	// StartParameter is a unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter"`

	// Currency is a three-letter ISO 4217 currency code.
	// See: https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`

	// TotalAmount is a total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64 `json:"total_amount"`
}

// ShippingAddress is a shipping address.
type ShippingAddress struct {
	// CountryCode is a	ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State, if applicable.
	State string `json:"state"`

	// City.
	City string `json:"city"`

	// StreetLine1 is a first line for the address.
	StreetLine1 string `json:"street_line1"`

	// StreetLine2 is a second line for the address.
	StreetLine2 string `json:"street_line2"`

	// PostCode is an address post code.
	PostCode string `json:"post_code"`
}

// OrderInfo is an information about an order.
type OrderInfo struct {
	// Name is a user's name.
	//
	// Optional.
	Name string `json:"name,omitempty"`

	// PhoneNumber is a user's phone number.
	//
	// Optional.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is a user's email.
	//
	// Optional.
	Email string `json:"email,omitempty"`

	// ShippingAddress is a user's shipping address.
	//
	// Optional.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption is one shipping option.
type ShippingOption struct {
	// ID is an option identifier.
	ID string `json:"id"`

	// Title is an option title.
	Title string `json:"title"`

	// Prices is an array of price portions.
	Prices []*LabeledPrice
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Currency is a three-letter ISO 4217 currency code.
	// See: https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`

	// TotalAmount is a total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload is a bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID is an identifier of the shipping option chosen by the user.
	//
	// Optional.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo is an order info provided by the user.
	//
	// Optional.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// TelegramPaymentChargeID is a Telegram payment identifier.
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID is a Provider payment identifier.
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// ShippingQuery is an information about an incoming shipping query.
type ShippingQuery struct {
	// ID is a unique query identifier.
	ID string `json:"id"`

	// From is a User who sent the query.
	From *User `json:"from"`

	// InvoicePayload is a Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingAddress is the User specified shipping address.
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// ID is a unique query identifier.
	ID string `json:"id"`

	// From is a User who sent the query.
	From *User `json:"from"`

	// Currency is a three-letter ISO 4217 currency code.
	// See: https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`

	// TotalAmount is a total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload is a Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID is an identifier of the shipping option chosen by the user.
	//
	// Optional.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo is an order info provided by the user.
	//
	// Optional.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// PassportData contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data is an array with information about documents and other Telegram Passport elements
	// that was shared with the bot.
	Data []*EncryptedPassportElement `json:"data"`

	// Credentials is an encrypted credentials required to decrypt the data.
	Credentials *EncryptedCredentials `json:"credentials"`
}

// PassportFile is a file uploaded to Telegram Passport.
// Currently, all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize in bytes.
	FileSize int `json:"file_size"`

	// FileDate is a unix time when the file was uploaded.
	FileDate int `json:"file_date"`
}

type EncryptedPassportElementType = string

const (
	EncryptedPassportElementTypePersonalDetails EncryptedPassportElementType = "personal_details"
	EncryptedPassportElementTypePassport        EncryptedPassportElementType = "passport"
	EncryptedPassportElementDriverLicense       EncryptedPassportElementType = "driver_license"
	EncryptedPassportElementIdentityCard        EncryptedPassportElementType = "identity_card"
	EncryptedPassportElementInternalPassport    EncryptedPassportElementType = "internal_passport"
	EncryptedPassportElementAddress             EncryptedPassportElementType = "address"
	EncryptedPassportUtilityBill                EncryptedPassportElementType = "utility_bill"
	EncryptedPassportBankStatement              EncryptedPassportElementType = "bank_statement"
	EncryptedPassportRentalAgreement            EncryptedPassportElementType = "rental_agreement"
	EncryptedPassportPassportRegistration       EncryptedPassportElementType = "passport_registration"
	EncryptedPassportTemporaryRegistration      EncryptedPassportElementType = "temporary_registration"
	EncryptedPassportPhoneNumber                EncryptedPassportElementType = "phone_number"
	EncryptedPassportEmail                      EncryptedPassportElementType = "email"
)

// EncryptedPassportElement contains information about documents or other Telegram Passport elements
// shared with the bot by the user.
type EncryptedPassportElement struct {
	// Type is an EncryptedPassportElement type.
	// One of EncryptedPassportElementTypePersonalDetails,
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportElementAddress,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration,
	// EncryptedPassportPhoneNumber,
	// EncryptedPassportEmail.
	Type EncryptedPassportElementType `json:"type"`

	// Data is a Base64-encoded encrypted Telegram Passport element data provided by the user.
	// Available for EncryptedPassportElementTypePersonalDetails,
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport and EncryptedPassportElementAddress types.
	// Can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Data string `json:"data,omitempty"`

	// PhoneNumber is a user's verified phone number, available only for EncryptedPassportPhoneNumber type.
	//
	// Optional.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is a user's verified email address, available only for EncryptedPassportEmail type.
	Email string `json:"email,omitempty"`

	// Files is an array of encrypted files with documents provided by the user.
	// Available for EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration and EncryptedPassportTemporaryRegistration types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Files []*PassportFile `json:"files,omitempty"`

	// FrontSide is an encrypted file with the front side of the document, provided by the user.
	// Available for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard and EncryptedPassportElementInternalPassport.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide is an encrypted file with the reverse side of the document, provided by the user.
	// Available for EncryptedPassportElementDriverLicense and EncryptedPassportElementIdentityCard.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie is an encrypted file with the selfie of the user holding a document, provided by the user.
	// Available for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard and EncryptedPassportElementInternalPassport.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation is an array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration and EncryptedPassportTemporaryRegistration types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Translation []*PassportFile `json:"translation,omitempty"`

	// Hash is a Base64-encoded element hash for using in PassportElementErrorUnspecified.
	Hash string `json:"hash"`
}

// EncryptedCredentials contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Data is a Base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption and authentication.
	Data string `json:"data"`

	// Hash is a Base64-encoded data hash for data authentication.
	Hash string `json:"hash"`

	// Secret is a Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption.
	Secret string `json:"secret"`
}

type PassportElementErrorSource string

const (
	PassportElementErrorSourceData             = "data"
	PassportElementErrorSourceFrontSide        = "front_side"
	PassportElementErrorSourceReverseSide      = "reverse_side"
	PassportElementErrorSourceSelfie           = "selfie"
	PassportElementErrorSourceFile             = "file"
	PassportElementErrorSourceFiles            = "files"
	PassportElementErrorSourceTranslationFile  = "translation_file"
	PassportElementErrorSourceTranslationFiles = "translation_files"
	PassportElementErrorSourceUnspecified      = "unspecified"
)

// PassportElementError is an error in the Telegram Passport element which was submitted that should be resolved by the user.
// It should be one of:
// PassportElementErrorDataField,
// PassportElementErrorFrontSide,
// PassportElementErrorReverseSide,
// PassportElementErrorSelfie,
// PassportElementErrorFile,
// PassportElementErrorFiles,
// PassportElementErrorTranslationFile,
// PassportElementErrorTranslationFiles,
// PassportElementErrorUnspecified.
type PassportElementError struct {
	PassportElementErrorDataField
	PassportElementErrorFrontSide
	PassportElementErrorReverseSide
	PassportElementErrorSelfie
	PassportElementErrorFile
	PassportElementErrorFiles
	PassportElementErrorTranslationFile
	PassportElementErrorTranslationFiles
	PassportElementErrorUnspecified
}

// PassportElementErrorDataField is an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Source is an error source, must be PassportElementErrorSourceData.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementTypePersonalDetails,
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportElementAddress.
	Type EncryptedPassportElementType `json:"type"`

	// FieldName is a name of the data field which has the error.
	FieldName string `json:"field_name"`

	// DataHash is a Base64-encoded data hash.
	DataHash string `json:"data_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorFrontSide is an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Source is an error source, must be PassportElementErrorSourceFrontSide.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport.
	Type EncryptedPassportElementType `json:"type"`

	// FileHash is a Base64-encoded hash of the file with the front side of the document.
	FileHash string `json:"file_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorReverseSide is an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Source is an error source, must be PassportElementErrorSourceReverseSide.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	Type EncryptedPassportElementType `json:"type"`

	// FileHash is a Base64-encoded hash of the file with the reverse side of the document.
	FileHash string `json:"file_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorSelfie is an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Source is an error source, must be PassportElementErrorSourceSelfie.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport.
	Type EncryptedPassportElementType `json:"type"`

	// FileHash is a Base64-encoded hash of the file with the selfie.
	FileHash string `json:"file_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorFile is an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	// Source is an error source, must be PassportElementErrorSourceFile.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration.
	Type EncryptedPassportElementType `json:"type"`

	// FileHash is a Base64-encoded file hash.
	FileHash string `json:"file_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorFiles is an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Source is an error source, must be PassportElementErrorSourceFiles.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration.
	Type EncryptedPassportElementType `json:"type"`

	// FileHashes is an array of base64-encoded file hashes.
	FileHashes []string `json:"file_hashes"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorTranslationFile is an issue with one of the files that constitute the translation of a document.
// The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Source is an error source, must be PassportElementErrorSourceTranslationFile.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration.
	Type EncryptedPassportElementType `json:"type"`

	// FileHash is a Base64-encoded file hash.
	FileHash string `json:"file_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorTranslationFiles is an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Source is an error source, must be PassportElementErrorSourceTranslationFiles.
	Source PassportElementErrorSource `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of:
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration.
	Type EncryptedPassportElementType `json:"type"`

	// FileHashes is an array of base64-encoded file hashes.
	FileHashes []string `json:"file_hashes"`

	// Message of an error.
	Message string `json:"message"`
}

// PassportElementErrorUnspecified is an issue in an unspecified place.
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	// Source is an error source, must be PassportElementErrorSourceUnspecified.
	Source PassportElementErrorSource `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue.
	Type EncryptedPassportElementType `json:"type"`

	// ElementHash us a Base64-encoded element hash.
	ElementHash string `json:"element_hash"`

	// Message of an error.
	Message string `json:"message"`
}

// Game is a Telegram game.
type Game struct {
	// Title of the game.
	Title string `json:"title"`

	// Description of the game.
	Description string `json:"description"`

	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`

	// Text is a brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game
	// when the bot calls setGameScore, or manually edited using editMessageText.
	// 0-4096 characters.
	//
	// Optional.
	Text string `json:"text,omitempty"`

	// TextEntities is a special entities that appear in text, such as usernames, URLs, bot commands, etc.
	//
	// Optional.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`

	// Animation that will be displayed in the game message in chats.
	// Upload via BotFather.
	//
	// Optional.
	Animation *Animation `json:"animation,omitempty"`
}

// CallbackGame is a placeholder, currently holds no information.
// Use BotFather to set up your game.
type CallbackGame struct {
}

// GameHighScore is one row of the high scores table for a game.
type GameHighScore struct {
	// Position in high score table for the game.
	Position int `json:"position"`

	// User.
	User *User `json:"user"`

	// Score.
	Score int `json:"score"`
}

type ParseMode string

const (
	ParseModeHTML       = "HTML"
	ParseModeMarkdown   = "Markdown"
	ParseModeMarkdownV2 = "MarkdownV2"
)
