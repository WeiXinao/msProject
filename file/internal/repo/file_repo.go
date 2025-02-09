package repo

import (
	"context"

	"github.com/WeiXinao/msProject/file/internal/domain"
	"github.com/WeiXinao/msProject/file/internal/repo/dao"
	"github.com/jinzhu/copier"
)

type FileRepo interface {
	FindByIds(ctx context.Context, ids []int64) (list []*domain.File, err error)
	FindByTaskCode(ctx context.Context, taskCode int64) (list []*domain.SourceLink, err error)
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

// FindByIds implements FileRepo.
func (f *fileRepo) FindByIds(ctx context.Context, ids []int64) (list []*domain.File, err error) {
	files, err := f.dao.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	list = make([]*domain.File, 0)
	err = copier.CopyWithOption(&list, files, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}
	 return
}

// FindByTaskCode implements FileRepo.
func (f *fileRepo) FindByTaskCode(ctx context.Context, taskCode int64) (list []*domain.SourceLink, err error) {
	list = make([]*domain.SourceLink, 0)
	links, err := f.dao.FindByTaskCode(ctx, taskCode)
	if err != nil {
		return nil, err
	}
	err = copier.CopyWithOption(&list, links, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func NewFileRepo(dao dao.FileDao) FileRepo {
	return &fileRepo{
		dao: dao,
	}
}
