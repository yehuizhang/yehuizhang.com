package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"yehuizhang.com/go-webapp-gin/src/database"
	"yehuizhang.com/go-webapp-gin/src/forms"
)

// ----------------------------------------------------------------
const PasswordNotMatch = UserCredentialError("PasswordNotMatch")

type UserCredentialError string

func (e UserCredentialError) Error() string {
	return string("UserCredentialError: " + e)
}

func (UserCredentialError) UserCredentialError() {}

// ----------------------------------------------------------------
type UserCredential struct {
	ID        string `json:"user_id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Active    bool   `json:"active,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type AuthHandler struct {
	database *database.Database
}

func newAuthHandler(database *database.Database) *AuthHandler {
	return &AuthHandler{database: database}
}

func (ah AuthHandler) Signup(c *gin.Context) (*UserInfo, *UserCredential, error) {

	credentialInput, err := readCredentialFromContext(c)

	if err != nil {
		return nil, nil, err
	}

	id := uuid.NewV4()

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(credentialInput.Password), 4)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to encrypt user's password. %s", err)
	}

	userCredential := UserCredential{
		ID:        id.String(),
		Username:  credentialInput.Username,
		Password:  string(encryptedPassword),
		Active:    true,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}

	encoded_userCredential, err := json.Marshal(userCredential)

	if err != nil {
		log.Panic("Unable to convert UserCredential to json")
		return nil, nil, err
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userInfo := UserInfo{
		UpdatedAt: time.Now().UnixNano(),
	}
	var encodedUserInfo []byte
	encodedUserInfo, err = json.Marshal(userInfo)
	if err != nil {
		return nil, nil, err
	}

	_, err = ah.database.Redis.Pipelined(dbCtx, func(p redis.Pipeliner) error {
		err := p.Set(dbCtx, createUserCredentialDbKey(userCredential.Username), encoded_userCredential, 0).Err()

		if err != nil {
			return err
		}
		return p.Set(dbCtx, createUserInfoDbKey(userCredential.ID), encodedUserInfo, 0).Err()
	})

	if err != nil {
		log.Panic(err)
		return nil, nil, err
	}
	return &userInfo, &userCredential, nil
}

func (ah AuthHandler) Signin(c *gin.Context) (*UserCredential, error) {
	credentialInput, err := readCredentialFromContext(c)

	if err != nil {
		return nil, err
	}

	dbResult, err := ah.database.Redis.Get(context.Background(), createUserCredentialDbKey(credentialInput.Username)).Result()
	switch {
	case err == redis.Nil:
		return nil, fmt.Errorf("user %s was not found", credentialInput.Username)
	case err != nil:
		return nil, fmt.Errorf("failed to retrieve credential information. %s", err)
	}

	var storedCredential UserCredential
	err = json.Unmarshal([]byte(dbResult), &storedCredential)
	if err != nil {
		return nil, fmt.Errorf("unable to parse UserCredential data from DB. %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedCredential.Password), []byte(credentialInput.Password))

	if err == nil {
		return &storedCredential, nil
	}
	return nil, PasswordNotMatch
}

func readCredentialFromContext(c *gin.Context) (*forms.UserCredential, error) {
	form := forms.UserCredential{}
	err := c.Bind(&form)

	if err != nil {
		return nil, fmt.Errorf("unable to read UserCredential from body: %s", err)
	}
	return &form, nil
}

func createUserCredentialDbKey(username string) string {
	return "user:credential:" + username
}
