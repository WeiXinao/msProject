package dao

import (
	"context"

	"github.com/WeiXinao/msProject/pkg/ormx"
	"github.com/WeiXinao/xkit/slice"
	"xorm.io/xorm"
)

type xormFileDao struct {
	db *xorm.Engine
}

// FindByIds implements FileDao.
func (x *xormFileDao) FindByIds(ctx context.Context, ids []int64) (list []*File, err error) {
	list = make([]*File, 0)	
	err = x.db.Context(ctx).Table(&File{}).
	In("id", slice.Map(ids, func(idx int, src int64) any {
		return src
	})).Find(list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FindByTaskCode implements FileDao.
func (x *xormFileDao) FindByTaskCode(ctx context.Context, taskCode int64) (list []*SourceLink, err error) {
	list = make([]*SourceLink, 0)
	err = x.db.Context(ctx).Where("link_code = ?", taskCode).Find(list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// SaveFile implements FileDao.
func (x *xormFileDao) SaveFile(ctx context.Context, file *File) error {
	_, err := x.db.Context(ctx).InsertOne(file)
	return err
}

// SaveFileAndSourceLink implements FileDao.
func (x *xormFileDao) SaveFileAndSourceLink(ctx context.Context, file File, sourceLink SourceLink) error {
	oldDB := x.db
	defer func()  {
		x.db = oldDB	
	}()
	return ormx.NewTxSession(x.db.Context(ctx)).EngineTx(func(engine *xorm.Engine) error {
		err := x.SaveFile(ctx, &file)
		if err != nil {
			return err
		}
		sourceLink.SourceCode = int64(file.Id)
		return x.SaveSourceLink(ctx, &sourceLink)
	})
}

// SaveSourceLink implements FileDao.
func (x *xormFileDao) SaveSourceLink(ctx context.Context, link *SourceLink) error {
	_, err := x.db.Context(ctx).InsertOne(link)
	return err
}

func NewXormFileDao(engine *xorm.Engine) FileDao {
	engine.Sync(
		new(File),
		new(SourceLink),
	)
	return &xormFileDao{
		db: engine,
	}
}
