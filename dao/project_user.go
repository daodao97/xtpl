package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/daodao97/xgo/xdb"
)

var ProjectUser xdb.Model

type ProjectUserRecord struct {
	Id         int       `json:"id"`
	AppID      string    `json:"appid"`
	RefUID     int       `json:"ref_uid"`
	InviteCode string    `json:"invite_code"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	UserName   string    `json:"user_name"`
	AvatarURL  string    `json:"avatar_url"`
	Channel    string    `json:"channel"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *ProjectUserRecord) Record() xdb.Record {
	return xdb.Record{
		"appid":       strings.TrimSpace(p.AppID),
		"ref_uid":     p.RefUID,
		"invite_code": strings.TrimSpace(p.InviteCode),
		"email":       normalizeEmail(p.Email),
		"password":    p.Password,
		"user_name":   strings.TrimSpace(p.UserName),
		"avatar_url":  strings.TrimSpace(p.AvatarURL),
		"channel":     strings.TrimSpace(p.Channel),
	}
}

func (p *ProjectUserRecord) FromRecord(record xdb.Record) {
	p.Id = record.GetInt("id")
	p.AppID = record.GetString("appid")
	p.RefUID = record.GetInt("ref_uid")
	p.InviteCode = record.GetString("invite_code")
	p.Email = record.GetString("email")
	p.Password = record.GetString("password")
	p.UserName = record.GetString("user_name")
	p.AvatarURL = record.GetString("avatar_url")
	p.Channel = record.GetString("channel")
	if created := record.GetTime("created_at"); created != nil {
		p.CreatedAt = *created
	}
	if updated := record.GetTime("updated_at"); updated != nil {
		p.UpdatedAt = *updated
	}
}

func CreateProjectUser(user *ProjectUserRecord) (int64, error) {
	if user == nil {
		return 0, fmt.Errorf("user is nil")
	}
	user.Email = normalizeEmail(user.Email)
	if ProjectUser == nil {
		return 0, fmt.Errorf("project user model not initialized")
	}
	lastId, err := ProjectUser.Insert(user.Record())
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func GetProjectUserByID(id int) (*ProjectUserRecord, error) {
	if ProjectUser == nil {
		return nil, fmt.Errorf("project user model not initialized")
	}
	record, err := ProjectUser.First(xdb.WhereEq("id", id))
	if err != nil {
		return nil, err
	}
	result := &ProjectUserRecord{}
	result.FromRecord(record)
	return result, nil
}

func GetProjectUserByAppAndEmail(appID string, email string) (*ProjectUserRecord, error) {
	if ProjectUser == nil {
		return nil, fmt.Errorf("project user model not initialized")
	}
	record, err := ProjectUser.First(
		xdb.WhereEq("appid", strings.TrimSpace(appID)),
		xdb.WhereEq("email", normalizeEmail(email)),
	)
	if err != nil {
		return nil, err
	}
	result := &ProjectUserRecord{}
	result.FromRecord(record)
	return result, nil
}

func GetProjectUserByInviteCode(appID string, inviteCode string) (*ProjectUserRecord, error) {
	if ProjectUser == nil {
		return nil, fmt.Errorf("project user model not initialized")
	}
	inviteCode = strings.TrimSpace(inviteCode)
	if inviteCode == "" {
		return nil, xdb.ErrNotFound
	}
	record, err := ProjectUser.First(
		xdb.WhereEq("appid", strings.TrimSpace(appID)),
		xdb.WhereEq("invite_code", inviteCode),
	)
	if err != nil {
		return nil, err
	}
	result := &ProjectUserRecord{}
	result.FromRecord(record)
	return result, nil
}

func ListProjectUsers(appID string, page int, size int, opt ...xdb.Option) (int64, []*ProjectUserRecord, error) {
	if ProjectUser == nil {
		return 0, nil, fmt.Errorf("project user model not initialized")
	}
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	baseOpt := []xdb.Option{
		xdb.OrderByDesc("id"),
	}
	if appID != "" {
		baseOpt = append(baseOpt, xdb.WhereEq("appid", strings.TrimSpace(appID)))
	}
	baseOpt = append(baseOpt, opt...)

	total, records, err := ProjectUser.Page(page, size, baseOpt...)
	if err != nil {
		return 0, nil, err
	}
	users := make([]*ProjectUserRecord, 0, len(records))
	for _, record := range records {
		item := &ProjectUserRecord{}
		item.FromRecord(record)
		users = append(users, item)
	}
	return total, users, nil
}

func UpdateProjectUserByID(id int, updates xdb.Record) error {
	if len(updates) == 0 {
		return fmt.Errorf("updates is empty")
	}
	if ProjectUser == nil {
		return fmt.Errorf("project user model not initialized")
	}
	_, err := ProjectUser.Update(updates, xdb.WhereEq("id", id))
	return err
}

func DeleteProjectUserByID(id int) error {
	if ProjectUser == nil {
		return fmt.Errorf("project user model not initialized")
	}
	_, err := ProjectUser.Delete(xdb.WhereEq("id", id))
	return err
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
