package dao

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	dao  *Dao
)

type Dao struct {
	engine *gorm.DB
}

func NewDao(engine *gorm.DB) *Dao {
	once.Do(func() {
		dao = &Dao{engine: engine}
	})
	return dao
}
