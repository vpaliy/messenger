package messages

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type MessagesStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *MessagesStore {
	return &MessagesStore{
		db: db,
	}
}

func (ms *MessagesStore) Get(query store.Query) (*model.Message, error) {
	var m model.Message
	if err := ms.db.Where(query.ToMap()).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ms *MessagesStore) GetAll(query store.Query) ([]*model.Message, error) {
	return nil, nil
}

func (ms *MessagesStore) Create(u *model.Message) error {
	return ms.db.Create(u).Error
}

func (ms *MessagesStore) Update(u *model.Message) error {
	return ms.db.Model(u).Update(u).Error
}

func (ms *MessagesStore) Delete(u *model.Message) error {
	return ms.db.Model(u).Delete(u).Error
}
