package model

import "gorm.io/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	model := db.Model(&t)
	if t.Name != "" {
		model = model.Where("name = ?", t.Name)
	}
	model.Where("state = ?", t.State)

	if err := model.Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	model := db.Model(&t)

	model = model.Offset(pageOffset).Limit(pageSize)

	//if pageOffset > 0{
	//}
	if t.Name != "" {
		model = model.Where("name = ?", t.Name)
	}
	model.Where("state = ?", t.State)

	if err = model.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Model(&t).Create(t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&t).Where("id", t.ID).Updates(t).Error
}
