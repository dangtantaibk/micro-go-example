package helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
	"tng/common/utils/cfgutil"
	"tng/common/utils/redisutil"
	"tng/user-profile-service/dtos"
)

const (
	SessionExpireTime = 60 * 60 * 24 * 7 * time.Second
	SessionTTLMinTime = 60 * 60 * 24 * 1 * time.Second
	TokenField        = "token"
	InfoField         = "info"
)

var (
	hashKey = cfgutil.Load("SESSION_KEY")
)

type TokenInfo struct {
	UserId string
	Token  string
}
type helper struct {
	redisCache redisutil.Cache
}

func NewHelper(redisCache redisutil.Cache) Helper {
	return &helper{
		redisCache: redisCache,
	}
}

type Helper interface {
	GenSessionToken() string
	StoreUserSession(tokenInfo *TokenInfo, userSession *dtos.UserSessionInfo) error
	UpdateUserSession(tokenInfo *TokenInfo, userSession *dtos.UserSessionInfo) error
	RemoveUserSession(tokenInfo *TokenInfo) error
	GetUserSession(tokenInfo *TokenInfo) (*dtos.UserSessionInfo, error)
	ValidateTTL(tokenInfo *TokenInfo) error
	GenToken(tokenInfo *TokenInfo) (string, error)
	GetSessionToken(userSession *dtos.UserSessionInfo) (string, error)
	GetTokenInfo(token string) (*TokenInfo, error)
}

func (h *helper) GenSessionToken() string {
	u := uuid.NewV4()
	return u.String()
}

func (h *helper) StoreUserSession(tokenInfo *TokenInfo, userSession *dtos.UserSessionInfo) error {
	key := tokenInfo.UserId
	err := h.redisCache.HSet(key, TokenField, tokenInfo.Token)
	if err != nil {
		return err
	}

	val, _ := json.Marshal(userSession)
	str:=string(val)
	fmt.Println("str: ", str)
	err = h.redisCache.HSet(key, InfoField, string(val))
	if err != nil {
		return err
	}

	h.redisCache.Expire(key, int64(SessionExpireTime))
	return nil
}

func (h *helper) UpdateUserSession(tokenInfo *TokenInfo, userSession *dtos.UserSessionInfo) error {
	key := tokenInfo.UserId

	val, err := json.Marshal(userSession)
	if err != nil {
		return err
	}

	err = h.redisCache.HSet(key, InfoField, val)
	if err != nil {
		return err
	}

	h.redisCache.Expire(key, int64(SessionExpireTime))
	return nil
}

func (h *helper) RemoveUserSession(tokenInfo *TokenInfo) error {
	_, err := h.redisCache.Del(tokenInfo.UserId)
	if err != nil {
		return err
	}

	return err
}

func (h *helper) GetUserSession(tokenInfo *TokenInfo) (*dtos.UserSessionInfo, error) {
	key := tokenInfo.UserId

	val, err := h.redisCache.HGet(key, InfoField)
	if err == redis.Nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	userSession := dtos.UserSessionInfo{}
	jsonInput, err := strconv.Unquote(string(val))
	err = json.Unmarshal([]byte(jsonInput), &userSession)
	if err != nil {
		return nil, err
	}

	return &userSession, nil
}

func (h *helper) ValidateTTL(tokenInfo *TokenInfo) error {
	key := tokenInfo.UserId
	duration, err := h.redisCache.TLL(key)
	if err != nil {
		return err
	}
	if duration > SessionTTLMinTime {
		return nil
	}

	return h.redisCache.ExpireWithDuration(key, SessionExpireTime)
}

func (h *helper) GenToken(tokenInfo *TokenInfo) (string, error) {
	data, err := json.Marshal(tokenInfo)
	if err != nil {
		return "", err
	}

	key := hashKey
	encrypted := h.simpleEncryptDecrypt(string(data), key)
	newToken := base64.StdEncoding.EncodeToString([]byte(encrypted))
	return newToken, nil
}

func (h *helper) GetSessionToken(userSession *dtos.UserSessionInfo) (string, error) {
	key := userSession.ID
	field := TokenField
	token, err := h.redisCache.HGet(string(key), field)
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (h *helper) GetTokenInfo(token string) (*TokenInfo, error) {
	encrypted, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	key := hashKey
	data := h.simpleEncryptDecrypt(string(encrypted), key)

	tokenInfo := &TokenInfo{}
	err = json.Unmarshal([]byte(data), &tokenInfo)
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func (h *helper) simpleEncryptDecrypt(input, key string) (output string) {
	for i := range input {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}
