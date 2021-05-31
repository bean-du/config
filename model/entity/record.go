package entity

type Record struct {
	ID        int    `gorm:"primary_key;column:id;" json:"id"`
	Key       string `gorm:"column:key;type:VARCHAR;size:200;" json:"key"`
	Value     string `gorm:"column:value;type:VARCHAR;size:200;" json:"value"`
	Version   string `gorm:"column:version;type:VARCHAR;size:200;" json:"version"`
	Pointer   string `gorm:"column:pointer;type:VARCHAR;size:200;" json:"pointer"`
	CreatedAt string `gorm:"column:created_at;type:TIMESTAMPTZ;" json:"created_at"`
}

func (k *Record) TableName() string {
	return "record"
}
