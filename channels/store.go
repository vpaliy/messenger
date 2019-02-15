package channels

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
	return &ChannelStore{db}
}

func NewSubscriptionStore(db *gorm.DB) *SubscriptionStore {
	return &SubscriptionStore{db}
}

func (cs *ChannelStore) Get(query store.Query) (*model.Channel, error) {
	var m model.Channel
	q := cs.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.Find(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (cs *ChannelStore) GetAll(query store.Query) ([]*model.Channel, error) {
	var channels []*model.Channel
	q := cs.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.Find(&channels).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return channels, nil
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

func (cs *ChannelStore) Create(c *model.Channel) error {
	tx := cs.db.Begin()
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

func (cs *ChannelStore) Update(c *model.Channel) error {
	tx := cs.db.Begin()
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

func (cs *ChannelStore) Delete(u *model.Channel) error {
	return cs.db.Model(u).Delete(u).Error
}

// SubscriptionStore methods
func (ss *SubscriptionStore) Get(query store.Query) (*model.Subscription, error) {
	var m model.Subscription
	q := ss.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.Find(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ss *SubscriptionStore) GetAll(query store.Query) ([]*model.Subscription, error) {
	var subs []*model.Subscription
	q := ss.db.Where(query.Selection()).Limit(query.Limit())
	if r := query.TimeRange(); r != nil {
		q.Where("created_at BETWEEN ? AND ?", r.From, r.To)
	}
	for _, entity := range query.Preloads() {
		q = q.Preload(entity)
	}
	if err := q.Find(&subs).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return subs, nil
}

func (ss *SubscriptionStore) Create(c *model.Channel, s *model.Subscription) error {
	err := ss.db.Model(c).Association("Members").Append(s).Error
	if err != nil {
		return err
	}
	return ss.db.Model(s).Preload("User").First(s).Error
}

func (ss *SubscriptionStore) Update(c *model.Channel, s *model.Subscription) error {
	return nil
}

func (ss *SubscriptionStore) Delete(c *model.Channel, s *model.Subscription) error {
	return ss.db.Model(&c).Association("Members").Delete(s).Error
}
