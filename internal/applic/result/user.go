package result

import (
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/entity"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/timex"
)

//go:generate go tool shoot new -json -file=$GOFILE
//go:generate go tool shoot map -path=../../domain/entity -way=<-  -file=$GOFILE -i

type User struct {
	infra.Mapper

	ID       uint
	Name     string
	Birthday timex.DateTime
	Age      int
}

func (q *User) readEntity(e *entity.User) {
	q.Age = e.Age()
}
