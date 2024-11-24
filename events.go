package revoltgo

import (
	"bytes"
	"fmt"
)

func aeuConstructor() any {
	return new(AbstractEventUpdate)
}

const jsonSkipAheadKeyType = len(`{"type":"`)

// eventTypeFromJSON extracts the event type from the JSON data.
func eventTypeFromJSON(data []byte) (string, error) {
	closingTagIndex := bytes.IndexByte(data[jsonSkipAheadKeyType:], '"')
	if closingTagIndex < 0 {
		return "", fmt.Errorf("closing quote of type field not found")
	}

	result := data[jsonSkipAheadKeyType : jsonSkipAheadKeyType+closingTagIndex]
	return string(result), nil
}

var eventToStruct = map[string]func() any{
	"Error": func() any { return new(EventError) },
	"Bulk":  func() any { return new(EventBulk) },

	"Authenticated": func() any { return new(EventAuthenticated) },
	"Ready":         func() any { return new(EventReady) },
	"Pong":          func() any { return new(EventPong) },
	"Auth":          func() any { return new(EventAuth) },

	/* All update events are abstracted away. */
	"ServerUpdate":       aeuConstructor,
	"ChannelUpdate":      aeuConstructor,
	"ServerRoleUpdate":   aeuConstructor,
	"WebhookUpdate":      aeuConstructor,
	"UserUpdate":         aeuConstructor,
	"ServerMemberUpdate": aeuConstructor,

	"Message":               func() any { return new(EventMessage) },
	"MessageAppend":         func() any { return new(EventMessageAppend) },
	"MessageDelete":         func() any { return new(EventMessageDelete) },
	"BulkMessageDelete":     func() any { return new(EventBulkMessageDelete) },
	"MessageReact":          func() any { return new(EventMessageReact) },
	"MessageUnreact":        func() any { return new(EventMessageUnreact) },
	"MessageUpdate":         func() any { return new(EventMessageUpdate) },
	"MessageRemoveReaction": func() any { return new(EventMessageRemoveReaction) },

	"ChannelCreate":      func() any { return new(EventChannelCreate) },
	"ChannelDelete":      func() any { return new(EventChannelDelete) },
	"ChannelAck":         func() any { return new(EventChannelAck) },
	"ChannelStartTyping": func() any { return new(EventChannelStartTyping) },
	"ChannelStopTyping":  func() any { return new(EventChannelStopTyping) },

	"ChannelGroupJoin":  func() any { return new(EventChannelGroupJoin) },
	"ChannelGroupLeave": func() any { return new(EventChannelGroupLeave) },

	"ServerCreate":      func() any { return new(EventServerCreate) },
	"ServerDelete":      func() any { return new(EventServerDelete) },
	"ServerRoleDelete":  func() any { return new(EventServerRoleDelete) },
	"ServerMemberJoin":  func() any { return new(EventServerMemberJoin) },
	"ServerMemberLeave": func() any { return new(EventServerMemberLeave) },

	"EmojiCreate": func() any { return new(EventEmojiCreate) },
	"EmojiDelete": func() any { return new(EventEmojiDelete) },

	"UserSettingsUpdate": func() any { return new(EventUserSettingsUpdate) },
	"UserRelationship":   func() any { return new(EventUserRelationship) },
	"UserPlatformWipe":   func() any { return new(EventUserPlatformWipe) },

	"WebhookCreate": func() any { return new(EventWebhookCreate) },
	"WebhookDelete": func() any { return new(EventWebhookDelete) },
}
