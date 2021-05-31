package repostory

import (
	"coco-tool/config/conf"
	"coco-tool/config/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	RecordRepo = &recordRepo{}
	db         *gorm.DB
)

func Init() {
	db = conf.DB
}

type recordRepo struct{}

func (r *recordRepo) Details(key string) ([]*entity.Record, error) {
	res := make([]*entity.Record, 0, 0)
	if err := db.Model(&entity.Record{}).Where("key = ?", key).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *recordRepo) Keys() ([]string, error) {
	res := make([]*entity.Record, 0, 0)
	if err := db.Select("key").Find(&res).Error; err != nil {
		return nil, err
	}
	keys := make([]string, 0, 0)
	for _, r := range res {
		keys = append(keys, r.Key)
	}
	return keys, nil
}

func (r *recordRepo) Del(key, version string) error {
	if version == "" {
		err := db.Where("key = ?", key).Delete(&entity.Record{}).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	err := db.Where("key = ? and version = ?", key, version).Delete(&entity.Record{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (r *recordRepo) Set(record *entity.Record) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Record{}).Where("key = ?", record.Key).Update("pointer", "no").Error; err != nil {
			return err
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "key"}, {Name: "value"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"pointer":      "yes",
			}),
		}).Create(record).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *recordRepo) Apply(key, version string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Record{}).Where("key = ?", key).Update("pointer", "no").Error; err != nil {
			return err
		}
		if err := tx.Model(&entity.Record{}).Where("key = ? and version = ?", key, version).Update("pointer", "yes").Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *recordRepo) Get(key, version string) (*entity.Record, error) {
	res := new(entity.Record)
	if err := db.Model(&entity.Record{}).Where("key = ? and version = ?", key, version).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
