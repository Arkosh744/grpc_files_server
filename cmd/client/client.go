package main

import (
	"fmt"
	items "github.com/Arkosh744/grpc_files_server/gen/item"
	"github.com/Arkosh744/grpc_files_server/internal/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	url      string = "http://164.92.251.245:8080/api/v1/products/"
	limit    int64  = 5
	offset   int64  = 5
	sortName string = "name"
	sortAsc  bool   = true
)

func main() {
	var conn *grpc.ClientConn
	cfg, err := config.New("internal/config")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := items.NewItemsServiceClient(conn)

	// Проверяем работу метода Fetch
	responseFetch, err := client.Fetch(ctx, &items.FetchRequest{Url: url})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(responseFetch)

	// Проверяем работу метода List
	responseList, err := client.List(ctx, &items.ListRequest{Offset: offset, Limit: limit, SortingName: sortName, SortingAsc: sortAsc})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(responseList)
}
