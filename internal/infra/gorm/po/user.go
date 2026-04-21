package po

import (
	"time"

	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/entity"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/enum"

	"gorm.io/gorm"
)

//go:generate go tool shoot map -path=../../../domain/entity -file=$GOFILE

type User struct {
	gorm.Model
	Name     string      `gorm:"not null;comment:姓名"`
	Gender   enum.Gender `gorm:"not null;comment:性别"`
	Birthday time.Time   `gorm:"not null;comment:生日"`
}

func (User) TableComment() string {
	return "用户"
}

func (q *User) readEntity(*entity.User) {
	q.CreatedAt = time.Time{}
	q.UpdatedAt = time.Time{}
	q.DeletedAt = gorm.DeletedAt{}
}
