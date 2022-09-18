package server

import (
	"fmt"
	"net"

	items "github.com/Arkosh744/grpc_files_server/gen/item"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv     *grpc.Server
	itemsServer items.ItemsServiceServer
}

func New(itemsServer items.ItemsServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		itemsServer: itemsServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	items.RegisterItemsServiceServer(s.grpcSrv, s.itemsServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
