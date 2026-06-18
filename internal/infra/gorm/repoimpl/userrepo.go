package repoimpl

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/entity"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/enum"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/repo"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm/po"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/x"
	"gorm.io/gorm"
)

type UserRepo struct {
	*Base[uint, entity.User, *entity.User, *po.User]
	logger *slog.Logger
}

func (r *UserRepo) Bind(txer repo.Txer) (repo.UserRepo, error) {
	db, ok := txer.Tx().(*gorm.DB)
	if !ok {
		return nil, errors.New("invalid gorm transaction")
	}
	return NewUserRepo(r.logger, db), nil
}

func NewUserRepo(logger *slog.Logger, db *gorm.DB) *UserRepo {
	r := &UserRepo{
		Base:   NewBase[uint, entity.User, *entity.User, *po.User](db),
		logger: x.SLogWithin(logger, &UserRepo{}),
	}
	var _ repo.UserRepo = r
	return r
}

func (r *UserRepo) Query(ctx context.Context, key string) ([]*entity.User, error) {
	userPool := []*po.User{
		{
			Model: gorm.Model{
				ID:        100,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "Tom",
			Gender:   enum.GenderMale,
			Birthday: time.Date(1987, 11, 29, 0, 0, 0, 0, time.Local),
		},
		{
			Model: gorm.Model{
				ID:        101,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "Jerry",
			Gender:   enum.GenderFemale,
			Birthday: time.Date(1991, 1, 17, 0, 0, 0, 0, time.Local),
		},
	}
	var userPOs []*po.User
	for _, u := range userPool {
		if strings.Contains(u.Name, key) {
			userPOs = append(userPOs, u)
		}
	}

	var users []*entity.User
	for _, u := range userPOs {
		users = append(users, u.ToEntity())
	}
	return users, nil
}
