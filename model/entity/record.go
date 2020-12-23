package entity

import "time"

type Record struct {
	ID        int64     `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Key       string    `gorm:"column:key;type:VARCHAR;size:200;" json:"key"`
	Value     string    `gorm:"column:value;type:VARCHAR;size:200;" json:"value"`
	Version   string    `gorm:"column:version;type:VARCHAR;size:200;" json:"version"`
	Pointer   string    `gorm:"column:pointer;type:VARCHAR;size:200;" json:"pointer"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;" json:"created_at"`
}

func (k *Record) TableName() string {
	return "record"
}
