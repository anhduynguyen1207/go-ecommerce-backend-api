package impl

import (
	"context"
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/database"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/model"
)

type sTicketItem struct {
	//implemeantation interface here
	r *database.Queries
}

func NewTicketItemImpl(r *database.Queries) *sTicketItem {
	return &sTicketItem{
		r: r,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	fmt.Println("CAL service GetTicketItemById...")
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}
	// mapper

	return &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}, nil
}
