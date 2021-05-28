package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := &model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) ListTag(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := &model.Tag{
		Name:  name,
		State: state,
	}
	return tag.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) Count(name string, state uint8) (int, error) {
	tag := &model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}
