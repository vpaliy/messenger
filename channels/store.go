package Subscription

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type ChannelStore struct {
	db *gorm.DB
}

type SubscriptionStore struct {
	db *gorm.DB
}

func NewChannelStore(db *gorm.DB) *ChannelStore {
	return &ChannelStore{
		db: db,
	}
}

func NewSubscriptionStore(db *gorm.DB) *SubscriptionStore {
	return &SubscriptionStore{
		db: db,
	}
}

func (ms *ChannelStore) Get(query store.Query) (*model.Channel, error) {
	var m model.Channel
	if err := ms.db.Where(query.ToMap()).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ms *ChannelStore) GetAll(query store.Query) ([]*model.Channel, error) {
	return nil, nil
}

func (ms *ChannelStore) Create(u *model.Channel) error {
	return ms.db.Create(u).Error
}

func (ms *ChannelStore) Update(u *model.Channel) error {
	return ms.db.Model(u).Update(u).Error
}

func (ms *ChannelStore) Delete(u *model.Channel) error {
	return ms.db.Model(u).Delete(u).Error
}

// SubscriptionStore methods

func (ms *SubscriptionStore) Get(query store.Query) (*model.Subscription, error) {
	var m model.Subscription
	if err := ms.db.Where(query.ToMap()).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ms *SubscriptionStore) GetAll(query store.Query) ([]*model.Subscription, error) {
	return nil, nil
}

func (ms *SubscriptionStore) Create(c *model.Channel, s *model.Subscription) error {
	return db.Model(&c).Association("Subscriptions").Append(s).Error
}

func (ms *SubscriptionStore) Update(c *model.Channel, s *model.Subscription) error {
	return db.Model(&c).Association("Subscriptions").Update(s).Error
}

func (ms *SubscriptionStore) Delete(c *model.Channel, s *model.Subscription) error {
	return db.Model(&c).Association("Subscriptions").Delete(s).Error
}
