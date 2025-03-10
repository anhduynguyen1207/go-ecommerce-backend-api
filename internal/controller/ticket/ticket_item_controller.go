package ticket

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/service"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

var TicketItem = new(cTicketItem)

type cTicketItem struct{}

func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, 1)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
