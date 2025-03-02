package global

import (
	"database/sql"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/logger"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)
