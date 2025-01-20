package initialize

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger()
}
