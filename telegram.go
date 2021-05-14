package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

type SetWebhook struct {
	URL string `json:"url"`

	IP                 string   `json:"ip_address,omitempty"`
	MaxConnections     *int     `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates *bool    `json:"drop_pending_updates,omitempty"`
}

func tgSetWebhook(URL string, AllowedUpdates []string) error {
	resp, err := RequestJSONPost(telegramAPIServerPrefix+"/setWebhook", SetWebhook{
		URL:            URL,
		AllowedUpdates: AllowedUpdates,
	})
	log.Println(string(resp.Body()))
	fasthttp.ReleaseResponse(resp)
	return err
}

type SendMessage struct {
	ChatID                   interface{} `json:"chat_id"` //Int or String
	Text                     string      `json:"text"`
	DisableWebPagePreview    bool        `json:"disable_web_page_preview"`
	DisableNotification      bool        `json:"disable_notification"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`

	ParseMode        string          `json:"parse_mode,omitempty"`
	Entities         []MessageEntity `json:"entities,omitempty"`
	ReplyToMessageID *int            `json:"reply_to_message_id,omitempty"`

	Method string `json:"method,omitempty"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`

	URL      string `json:"url,omitempty"`
	User     *User  `json:"user,omitempty"`
	Language string `json:"language,omitempty"`
}

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           *bool  `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool  `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool  `json:"supports_inline_queries,omitempty"`
}

type Update struct {
	UpdateID int `json:"update_id"`

	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
}

type Message struct {
	MessageID int  `json:"message_id"`
	Date      int  `json:"date"`
	Chat      Chat `json:"chat"`

	From                 *User           `json:"from,omitempty"`
	SenderChat           *Chat           `json:"sender_chat,omitempty"`
	ForwardFrom          *User           `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat           `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID *int            `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string          `json:"forward_signature,omitempty"`
	ForwardSenderName    string          `json:"forward_sender_name,omitempty"`
	ForwardDate          *int            `json:"forward_date,omitempty"`
	ReplyToMessage       *Message        `json:"reply_to_message,omitempty"`
	ViaBot               *User           `json:"via_bot,omitempty"`
	EditDate             *int            `json:"edit_date,omitempty"`
	MediaGroupID         string          `json:"media_group_id,omitempty"`
	AuthorSignature      string          `json:"author_signature,omitempty"`
	Text                 string          `json:"text,omitempty"`
	Entities             []MessageEntity `json:"entities,omitempty"`
}

type Chat struct {
	ID   int    `json:"id"`
	Type string `json:"type"`

	Title                 string   `json:"title,omitempty"`
	Username              string   `json:"username,omitempty"`
	FirstName             string   `json:"first_name,omitempty"`
	LastName              string   `json:"last_name,omitempty"`
	Bio                   string   `json:"bio,omitempty"`
	Description           string   `json:"description,omitempty"`
	InviteLink            string   `json:"invite_link,omitempty"`
	PinnedMessage         *Message `json:"pinned_message,omitempty"`
	SlowModeDelay         *int     `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime *int     `json:"message_auto_delete_time,omitempty"`
	LinkedChatID          *int     `json:"linked_chat_id,omitempty"`
}
