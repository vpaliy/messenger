package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type SubscriptionStore struct {
	db *gorm.DB
}

func NewSubscriptionStore(db *gorm.DB) *SubscriptionStore {
	return &SubscriptionStore{db}
}

func (s *SubscriptionStore) Fetch(id string) (*model.Subscription, error) {
	var m model.Subscription
	// TODO: add searching by channel name as well
	err := s.db.Where("id = ?", id).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (s *SubscriptionStore) FetchAll(user interface{}, args ...store.Option) ([]*model.Subscription, error) {
	var ms []*model.Subscription
	options := store.NewOptions(args...)
	tx := s.db.Where("user_id = ?", user).Limit(options.Limit)
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (ss *SubscriptionStore) Create(c *model.Channel, s *model.Subscription) error {
	err := ss.db.Model(c).Association("Members").Append(s).Error
	if err != nil {
		return err
	}
	return ss.db.Model(s).Preload("User").First(s).Error
}

func (ss *SubscriptionStore) Update(c *model.Channel, s *model.Subscription) error {
	// TODO: implement this one
	return nil
}

func (ss *SubscriptionStore) Delete(c *model.Channel, s *model.Subscription) error {
	return ss.db.Model(&c).Association("Members").Delete(s).Error
}
