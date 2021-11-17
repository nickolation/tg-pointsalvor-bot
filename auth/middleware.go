package auth

import (
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/nickolation/pointsalvor"
)

//custom auth mod --> zap
const (
	authIdentity = "auth"
	authKey      = "auth!!chat_id:%d!!project_id:%d"
	authTable    = "table:[%d]"
)

//custom errors --> zap
var (
	//parse key
	errNilKey     = fmt.Errorf("key is nil: parse isn'd valid operation")
	errInvalidKey = fmt.Errorf("key isn't valid structure: parse isn'd valid operation")
	errForeignKey = fmt.Errorf("key is foreign for auth-data: parse isn'd valid operation")
	errNilId      = fmt.Errorf("id in key-string is nil: parse isn'd valid operation")

	//auth
	errAuth      = fmt.Errorf("internal auth error: auth isn' valid ")
	errNilChatId = fmt.Errorf("chatId is nil: auth locked")

	//registr
	errNilOpt    = fmt.Errorf("keyopt data is nil: registration locked")
	errIdProject = fmt.Errorf("id project is nil: registration locked")
)

//Storage with part of key auth-data: auth!!...
type KeyTokenOpt struct {
	//identitty of user
	ChatId int64

	//used in set-handlers for identity of entities
	ProjId string

	//todoist token required for performing the request to api
	Token string
}

//Status authentification used in validate middleware method
//Services to make bool validation process without lock and returning methods
type StatusAuth struct {
	//status of authorization
	//true is auth - false against
	status bool
}

//makes new pointsalvor-agent with token is need to use sdk-method
func linkAgent(token string) (*sdk.Agent, error) {
	ag, err := sdk.NewAgent(token)
	if err != nil {
		return nil, err
	}

	return ag, nil
}

//make name of table
//used in signUp method for sending to sdk
func callTable(chatId int64) string {
	return fmt.Sprintf(authTable, chatId)
}

//make auth-key value: auth!!chat_id:<>!!projId:<>
func makeKey(chatId, projId int64) (string, error) {
	if projId == 0 {
		return "", errIdProject
	}

	return fmt.Sprintf(authKey, chatId, projId), nil
}

//key allows the struct: auth!chat_id:<...>!!proj_id:<...>
//parse key to custom struct contains need parts
func parseAuthKey(key string) (*KeyTokenOpt, error) {
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

		//		string value now 
		//		later --> int64
		projId = chatPart[strings.Index(projPart, ":")+1:]
	)

	//string chatId --> int64 allows the telegram notation
	c, err := strconv.Atoi(chatPart[strings.Index(chatPart, ":")+1:])
	if err != nil {
		return nil, err
	}
	chatId := int64(c)

	if chatId == 0 || projId == "" {
		return nil, errNilId
	}

	return &KeyTokenOpt{
		ChatId: chatId,
		ProjId: projId,
	}, nil
}

//Validate user for existing the key:auth!! with same chat_id
//Doesn't lock the process but make statusAuth: status - false
func isAuth(key string, chatId int64) StatusAuth {
	parts, err := parseAuthKey(key)
	if err != nil {
		return StatusAuth{
			status: false,
		}
	}

	if parts.ChatId == chatId {
		return StatusAuth{
			status: true,
		}
	}

	return StatusAuth{
		status: false,
	}
}
