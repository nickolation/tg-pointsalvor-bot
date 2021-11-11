package auth

import (
	"fmt"
	"strings"
)

//custom auth mod --> zap
const (
	authIdentity = "auth"
)

//custom errors --> zap
var (
	errNilKey     = fmt.Errorf("key is nil: parse isn'd valid operation")
	errInvalidKey = fmt.Errorf("key isn't valid structure: parse isn'd valid operation")
	errForeignKey = fmt.Errorf("key is foreign for auth-data: parse isn'd valid operation")
	errNilId      = fmt.Errorf("id in key-string is nil: parse isn'd valid operation")

	errAuth      = fmt.Errorf("internal auth error: auth isn' valid ")
	errNilChatId = fmt.Errorf("chatId is nil: auth locked!")
)

//Storage with part of key auth-data: auth!!...
type KeypartStorage struct {
	//identitty of user
	chatId string

	//used in set-handlers for identity of entities
	projId string
}

//Status authentification used in validate middleware method
//Services to make bool validation process without lock and returning methods
type StatusAuth struct {
	//status of authorization
	//true is auth - false against
	Status bool
}

//key allows the struct: auth!chat_id:<...>!!proj_id:<...>
//parse key to custom struct contains need parts
func ParseAuthKey(key string) (*KeypartStorage, error) {
	if key == "" {
		return nil, errNilKey
	}

	partStorage := strings.Split(key, "!!")

	if len(partStorage) != 3 {
		return nil, errInvalidKey
	}

	if partStorage[0] != authIdentity {
		return nil, errForeignKey
	}

	var (
		chatPart = partStorage[1]
		projPart = partStorage[2]

		chatId = chatPart[strings.Index(chatPart, ":")+1:]
		projId = chatPart[strings.Index(projPart, ":")+1:]
	)

	if chatId == "" || projId == "" {
		return nil, errNilId
	}

	return &KeypartStorage{
		chatId: chatId,
		projId: projId,
	}, nil
}

//Validate user for existing the key:auth!! with same chat_id
//Doesn't lock the process but make statusAuth: status - false
func IsAuth(key string, chatId string) StatusAuth {
	parts, err := ParseAuthKey(key)
	if err != nil {
		return StatusAuth{
			Status: false,
		}
	}

	if parts.chatId == chatId {
		return StatusAuth{
			Status: true,
		}
	}

	return StatusAuth{
		Status: false,
	}
}
