package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"yehuizhang.com/go-webapp-gin/src/database"
	"yehuizhang.com/go-webapp-gin/src/forms"
)

type UserInfo struct {
	Name      string `json:"name"`
	Birthday  string `json:"birthday,omitempty"`
	Gender    string `json:"gender,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type UserInfoHandler struct {
	database *database.Database
}

func NewUserInfoHandler(database *database.Database) *UserInfoHandler {
	return &UserInfoHandler{database: database}
}

func (ui UserInfoHandler) AddOrUpdate(uid string, input forms.UserInfo) (*UserInfo, error) {

	userInfo := UserInfo{
		Name:      input.Name,
		Birthday:  input.Birthday,
		Gender:    input.Gender,
		PhotoURL:  input.PhotoURL,
		UpdatedAt: time.Now().UnixNano(),
	}

	encoded_userInfo, err := json.Marshal(userInfo)

	if err != nil {
		return nil, err
	}
	err = ui.database.Redis.Set(context.Background(), createUserInfoDbKey(uid), encoded_userInfo, 0).Err()

	if err != nil {
		return nil, fmt.Errorf("error when try to save data to database %s", err)
	}

	return &userInfo, nil
}

func (ui UserInfoHandler) GetUserInfo(uid string) (*UserInfo, error) {

	data, err := ui.database.Redis.Get(context.Background(), createUserInfoDbKey(uid)).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve UserInfo from DB for user %s. %s", uid, err)
	}

	switch {
	case err == redis.Nil:
		return nil, fmt.Errorf("userInfo was not found or user: %s", uid)
	case err != nil:
		return nil, fmt.Errorf("failed to retrieve UserInfo for user: %s. %s", uid, err)
	}
	var userInfo *UserInfo
	err = json.Unmarshal([]byte(data), &userInfo)
	if err != nil {
		return nil, fmt.Errorf("unable to parse UserInfo data from DB. %s", err)
	}

	return userInfo, nil

}

func createUserInfoDbKey(uid string) string {
	return "user:info:" + uid
}
