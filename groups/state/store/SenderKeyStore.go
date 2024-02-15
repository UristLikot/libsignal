package store

import (
	"github.com/UristLikot/libsignal/groups/state/record"
	"github.com/UristLikot/libsignal/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
