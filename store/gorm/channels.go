package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
)

type ChannelStore struct {
	db *gorm.DB
}

func NewChannelStore(db *gorm.DB) store.ChannelStore {
	return &ChannelStore{db}
}

func filterTags(tags []model.Tag) []model.Tag {
	seen := map[string]bool{}
	result := []model.Tag{}
	for _, t := range tags {
		if !seen[t.Tag] {
			result = append(result, t)
		}
		seen[t.Tag] = true
	}
	return result
}

func (s *ChannelStore) Fetch(id string) (*model.Channel, error) {
	var m model.Channel
	err := s.db.Where("id = ?", id).Or("name = ?", id).
		Preload("Tags").
		Preload("Creator").
		Preload("Members").
		Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// fetch channels based on their
func (s *ChannelStore) Search(query string, args ...store.Option) ([]*model.Channel, error) {
	var ms []*model.Channel
	options := store.NewOptions(args...)
	tx := s.db.Where("id = ?", query).Or("name = ?", query).Limit(options.Limit).
		Preload("Tags").
		Preload("Creator").
		Preload("Members")
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *ChannelStore) GetCreatedBy(user string, args ...store.Option) ([]*model.Channel, error) {
	var ms []*model.Channel
	options := store.NewOptions(args...)
	tx := s.db.Where("creator_id = ?", user).Limit(options.Limit).
		Preload("Tags").
		Preload("Creator").
		Preload("Members")
	if tr := options.TimeRange(); tr != nil {
		tx.Where("created_at BETWEEN ? AND ?", tr.From, tr.To)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *ChannelStore) Create(c *model.Channel) error {
	tx := s.db.Begin()
	tags := filterTags(c.Tags)
	if err := tx.Create(c).Error; err != nil {
		return err
	}
	for _, tag := range tags {
		err := tx.Where(&model.Tag{Tag: tag.Tag}).First(&tag).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(c).Association("Tags").Append(tag).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Where(c.ID).Preload("Tags").Preload("Creator").Find(c).Error; err != nil {
		tx.Rollback()
		return err
	}
	c.Tags = tags
	return tx.Commit().Error
}

func (s *ChannelStore) Update(c *model.Channel) error {
	tx := s.db.Begin()
	if err := tx.Model(c).Update(c).Error; err != nil {
		return err
	}
	tags := filterTags(c.Tags)
	for _, tag := range tags {
		err := tx.Where(&model.Tag{Tag: tag.Tag}).First(&tag).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Model(c).Association("Tags").Replace(tags).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Where(c.ID).
		Preload("Tags").
		Preload("Members").
		Preload("Creator").
		Find(c).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *ChannelStore) GetForMember(id string, args ...store.Option) ([]*model.Channel, error) {
	// TODO:
	return nil, nil
}

func (s *ChannelStore) Delete(u *model.Channel) error {
	// TODO: delete all subscriptions as well
	return s.db.Model(u).Delete(u).Error
}
