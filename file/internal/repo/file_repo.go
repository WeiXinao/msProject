package repo

import (
	"context"

	"github.com/WeiXinao/msProject/file/internal/domain"
	"github.com/WeiXinao/msProject/file/internal/repo/dao"
	"github.com/jinzhu/copier"
)

type FileRepo interface {
	SaveFileAndSourceLink(ctx context.Context, file domain.File, sourceLink domain.SourceLink) error
}

func (f *fileRepo) SaveFileAndSourceLink(ctx context.Context, file domain.File,
	sourceLink domain.SourceLink) error {
	fileEty := dao.File{}
	err := copier.Copy(&fileEty, file)
	if err != nil {
		return err
	}
	sourceLinkEty := dao.SourceLink{}
	err = copier.Copy(&sourceLinkEty, sourceLink)
	if err != nil {
		return err
	}
	return f.dao.SaveFileAndSourceLink(ctx, fileEty, sourceLinkEty)
}

type fileRepo struct {
	dao dao.FileDao
}

func NewFileRepo(dao dao.FileDao) FileRepo {
	return &fileRepo{
		dao: dao,
	}
}