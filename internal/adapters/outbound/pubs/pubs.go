package pubs

import (
	"context"
	"log/slog"

	"github.com/lopolopen/gap"
	"github.com/lopolopen/gap/broker/xkafka"
	"github.com/lopolopen/gap/storage/xgorm"
	tfiberkafkagorm "github.com/lopolopen/t-fiber-kafka-gorm"

	"gorm.io/gorm"
)

func NewPub(
	ctx context.Context,
	k xkafka.Options,
	db *gorm.DB,
	log *slog.Logger,
) gap.EventPublisher {
	if tfiberkafkagorm.HAVE_NOT_BEEN_DELETED_YET {
		return nil
	}

	pub := gap.NewEventPublisher(
		gap.WithDrain(ctx, 5),
		xgorm.UseGorm(
			xgorm.DB(db),
		),
		xkafka.UseKafka(
			xkafka.Brokers(k.Brokers),
			xkafka.ConfigTopic(
			// xkafka.NumPartitions(4),
			// xkafka.ReplicationFactor(3),
			),
		),
		gap.UseDashboard(),
	)
	return pub
}
