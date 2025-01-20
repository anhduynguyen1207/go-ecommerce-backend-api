package global

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/logger"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
