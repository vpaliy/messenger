package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) store.UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Fetch(user string) (*model.User, error) {
	var m model.User
	err := s.db.Where("id = ?", user).Or("username = ?", user).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (s *UserStore) Search(query string, args ...store.Option) ([]*model.User, error) {
	var ms []*model.User
	options := store.NewOptions(args...)
	// find users whose name starts with the given parameter
	tx := s.db.Where("username LIKE ?", query+"%").Limit(options.Limit)
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *UserStore) Create(u *model.User) (err error) {
	return s.db.Create(u).Error
}

func (s *UserStore) Update(u *model.User) error {
	return s.db.Model(u).Update(u).Error
}

func (s *UserStore) Delete(u *model.User) error {
	return s.db.Model(u).Delete(u).Error
}
