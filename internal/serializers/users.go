package serializers

import (
	"fmt"
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type UserRes struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	Username    string         `json:"username"`
	IsAdmin     bool           `json:"is_admin"`
	UserSetting UserSettingRes `json:"user_settings"`
}

func UserResponse(u *db.User) UserRes {
	return UserRes{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		Username:  u.Username,
		IsAdmin:   u.IsAdmin,
		UserSetting: UserSettingRes{
			ID:            u.UserSetting.ID,
			Country:       u.UserSetting.Country,
			Timezone:      u.UserSetting.Timezone,
			Language:      u.UserSetting.Language,
			Theme:         u.UserSetting.Theme,
			FullName:      u.UserSetting.FullName,
			Email:         u.UserSetting.Email,
			Birthday:      fmt.Sprintf("%s", u.UserSetting.Birthday.Format("2006-01-02")),
			UseMetric:     u.UserSetting.UseMetric,
			Sex:           u.UserSetting.Sex,
			Height:        u.UserSetting.Height,
			BodyFat:       u.UserSetting.BodyFat,
			Avatar:        u.UserSetting.Avatar,
			ActivityLevel: u.UserSetting.ActivityLevel,
		},
	}
}
