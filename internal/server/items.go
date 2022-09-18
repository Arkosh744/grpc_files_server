package server

import (
	"context"
	items "github.com/Arkosh744/grpc_files_server/gen/item"
)

type ItemsService interface {
	Fetch(ctx context.Context, request *items.FetchRequest) (*items.FetchResponse, error)
	List(ctx context.Context, request *items.ListRequest) (*items.ListResponse, error)
}

type ItemsServiceServer struct {
	service ItemsService
}

func NewItemServer(service ItemsService) *ItemsServiceServer {
	return &ItemsServiceServer{
		service: service,
	}
}

func (h *ItemsServiceServer) Fetch(ctx context.Context, request *items.FetchRequest) (*items.FetchResponse, error) {
	return h.service.Fetch(ctx, request)
}

func (h *ItemsServiceServer) List(ctx context.Context, request *items.ListRequest) (*items.ListResponse, error) {
	return h.service.List(ctx, request)
}
