package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type CreateTagReq struct {
	Name     string `json:"name"`
	CreateBy string `json:"create_by"`
	State    uint8  `json:"state"`
}

type ListTagReq struct {
	Name  string `form:"name"`
	State uint8  `form:"state,default=1"`
}

type CountTagReq struct {
	Name  string `form:"name"`
	State uint8  `form:"state,default=1"`
}

func (svr *Service) CreateTag(param *CreateTagReq) {
	_ = svr.dao.CreateTag(param.Name, param.State, param.CreateBy)
}

func (svr *Service) CountTag(param *CountTagReq) (int, error) {
	return svr.dao.Count(param.Name, param.State)
}

func (svr *Service) ListTag(param *ListTagReq, pager app.Pager) ([]*model.Tag, error) {
	return svr.dao.ListTag(param.Name, param.State, pager.Page, pager.PageSize)
}
