package messages

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type MessageStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *MessageStore {
	return &MessageStore{
		db: db,
	}
}

func (ms *MessageStore) Get(query store.Query) (*model.Message, error) {
	var m model.Message
	q := ms.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ms *MessageStore) GetAll(query store.Query) ([]*model.Message, error) {
	var messages []*model.Message
	q := ms.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.Find(&messages).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return messages, nil
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
	return ms.db.Model(u).Update(u).Error
}

func (ms *MessageStore) Delete(u *model.Message) error {
	return ms.db.Model(u).Delete(u).Error
}
