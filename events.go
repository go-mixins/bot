package bot

// UpdateType modelled after Telegram update types
type UpdateType int

// Possible update types
const (
	UpdCallbackQuery UpdateType = iota
	UpdChannelPost
	UpdChosenInlineResult
	UpdEditedChannelPost
	UpdEditedMessage
	UpdInlineQuery
	UpdShippingQuery
	UpdPreCheckoutQuery
	UpdMessage
)

// MessageType modelled after Telegram message subtypes
type MessageType int

// Message subtypes
const (
	MsgVoice MessageType = iota
	MsgVideoNote
	MsgVideo
	MsgVenue
	MsgText
	MsgSupergroupChatCreated
	MsgSuccessfulPayment
	MsgSticker
	MsgPinnedMessage
	MsgPhoto
	MsgNewChatTitle
	MsgNewChatPhoto
	MsgNewChatMembers
	MsgMigrateToChatID
	MsgMigrateFromChatID
	MsgLocation
	MsgLeftChatMember
	MsgInvoice
	MsgGroupChatCreated
	MsgGame
	MsgDocument
	MsgDeleteChatPhoto
	MsgContact
	MsgChannelChatCreated
	MsgAudio
)
