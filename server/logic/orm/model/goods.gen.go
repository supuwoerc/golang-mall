// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const TableNameGood = "goods"

// Good mapped from table <goods>
type Good struct {
	ID            string          `gorm:"column:id;primaryKey" json:"id"`                       // ID
	Name          string          `gorm:"column:name;not null" json:"name"`                     // 商品名称
	Img           string          `gorm:"column:img;not null" json:"img"`                       // 商品图片
	Price         decimal.Decimal `gorm:"column:price;not null" json:"price"`                   // 现价
	OriginalPrice decimal.Decimal `gorm:"column:original_price;not null" json:"original_price"` // 原价
	Desc          string          `gorm:"column:desc;not null" json:"desc"`                     // 描述
	CreatedAt     time.Time       `gorm:"column:created_at;not null" json:"created_at"`         // 创建时间
	UpdatedAt     time.Time       `gorm:"column:updated_at;not null" json:"updated_at"`         // 更新时间
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`                  // 删除时间
}

// TableName Good's table name
func (*Good) TableName() string {
	return TableNameGood
}
