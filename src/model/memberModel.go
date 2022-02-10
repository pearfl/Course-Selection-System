package model

import (
	"course_select/src/database"
	types "course_select/src/global"
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Member struct {
	UserID    string         `json:"user_id" form:"user_id" gorm:"primary_key"`
	Nickname  string         `json:"nickname" form:"nickname"`
	Username  string         `json:"username" form:"username"`
	Password  string         `json:"password" form:"password"`
	UserType  types.UserType `json:"user_type" form:"user_type"`
	IsDeleted bool           `json:"is_deleted" form:"is_deleted"`
}

func (Member) TableName() string {
	return "member"
}

func (member *Member) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("user_id", uuid)
}

// func md5V(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func (model *Member) CreateMember(newMember Member) (string, error) {
	return "不知道", nil
}

/*
@title	GetMember
@description	基于user_id查询一个成员信息
@auth	马信宏	时间(2022/2/10   14:42)
*/

func GetMember(user_id string) (Member, error) {
	var result Member
	err := database.MySqlDb.First(&Member{}, "user_id = ?", user_id).Scan(&result).Error
	return result, err
}

// GetAllMembers 返回所有成员
func (member *Member) GetAllMembers(offset, limit int) ([]Member, error) {
	var ans []Member
	err := database.MySqlDb.Limit(limit).Offset(offset).Find(&ans).Error
	if err != nil {
		return ans, err
	}
	return ans, nil
}

func UpdateMember(user_id string, nickname string) error {
	result := db.Model(&Member{}).Where("user_id = ?", user_id).Update("nickname", nickname)
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteMember(user_id string) error {
	result := db.Model(&Member{}).Where("user_id = ?", user_id).Update("is_deleted", true)
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}
