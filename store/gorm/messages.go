package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type MessageStore struct {
	db *gorm.DB
}

func NewMessageStore(db *gorm.DB) store.MessageStore {
	return &MessageStore{
		db: db,
	}
}

func (s *MessageStore) Fetch(id store.Arg) (*model.Message, error) {
	return nil, nil
}

func (s *MessageStore) GetForChannel(channel store.Arg, args ...store.Option) ([]*model.Message, error) {
	var ms []*model.Message
	options := store.NewOptions(args...)
	tx := s.db.Where("channel_id = ?", channel).Limit(options.Limit).
		Preload("Attachments").
		Preload("User")
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *MessageStore) Search(channel, query store.Arg, args ...store.Option) ([]*model.Message, error) {
	var ms []*model.Message
	options := store.NewOptions(args...)
	tx := s.db.Where("channel_id = ?", channel).
		Where("text LIKE ?", multipleMatch(query)).
		Limit(options.Limit).
		Preload("Attachments").
		Preload("User")
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *MessageStore) GetForUser(user store.Arg, args ...store.Option) ([]*model.Message, error) {
	var ms []*model.Message
	options := store.NewOptions(args...)
	tx := s.db.Where("user_id = ?", user).Limit(options.Limit).
		Preload("Attachments").
		Preload("User")
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (ms *MessageStore) Create(m *model.Message) error {
	tx := ms.db.Begin()
	if err := tx.Create(m).Error; err != nil {
		return err
	}
	attachments := m.Attachments
	for _, a := range attachments {
		if err := tx.Model(m).Association("Attachments").Append(a).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	err := tx.Where(m.ID).
		Preload("User").
		Preload("Attachments").
		Preload("Channel").
		Find(m).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	m.Attachments = attachments
	return tx.Commit().Error
}

func (ms *MessageStore) Update(u *model.Message) error {
	// TODO: update attachments if needed
	return ms.db.Model(u).Update(u).Error
}

func (ms *MessageStore) Delete(u *model.Message) error {
	// TODO: delete attachments too
	return ms.db.Model(u).Delete(u).Error
}
