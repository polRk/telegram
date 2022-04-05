package api

type TelegramTypes interface {
	Update | User | Chat | Message | MessageID | MessageEntity | PhotoSize | Animation | Audio | Document |
		Video | VideoNote | Voice | Contact | Dice | PollOption | PollAnswer | Poll | Location | Venue |
		ProximityAlertTriggered | MessageAutoDeleteTimerChanged | VoiceChatScheduled | VoiceChatStarted |
		VoiceChatEnded | VoiceChatParticipantsInvited | UserProfilePhotos | File | ReplyKeyboardMarkup |
		KeyboardButton | KeyboardButtonPollType | ReplyKeyboardRemove | InlineKeyboardMarkup | InlineKeyboardButton |
		LoginURL | CallbackQuery | ForceReply | ChatPhoto | ChatInviteLink | ChatMember | ChatMemberOwner |
		ChatMemberAdministrator | ChatMemberRestricted | ChatMemberLeft | ChatMemberBanned | ChatMemberUpdated |
		ChatJoinRequest | ChatPermissions | ChatLocation
}
