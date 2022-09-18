package service

import (
	"context"
	"encoding/csv"
	"errors"
	items "github.com/Arkosh744/grpc_files_server/gen/item"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (items.Item, error)
	InsertOne(ctx context.Context, item items.Item) error
	UpdateOne(ctx context.Context, item items.Item) error
	List(ctx context.Context, request *items.ListRequest) (*items.ListResponse, error)
}

type Items struct {
	repo Repository
}

func NewItems(repo Repository) *Items {
	return &Items{
		repo: repo,
	}
}

func (s *Items) Fetch(ctx context.Context, request *items.FetchRequest) (*items.FetchResponse, error) {
	data, err := getCVSdata(request.Url)
	if err != nil {
		return &items.FetchResponse{Code: 0, Text: "Doesn't get CVS file"}, err
	}
	log.Println(1, data, err)
	for _, v := range data {
		price, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			log.Println(err)
		}
		existingItem, err := s.repo.GetByName(ctx, v[0])
		if err == nil && price != existingItem.Price {
			existingItem.Price = price
			existingItem.UpdatedAt = time.Now()
			err = s.repo.UpdateOne(ctx, existingItem)
			log.Println("updated item: ", existingItem.Name)
			if err != nil {
				log.Println(err)
			}
		} else if err == mongo.ErrNoDocuments {
			log.Println("Inserting new item: ", v[0])
			err := s.repo.InsertOne(ctx, items.Item{
				Name:      v[0],
				Price:     price,
				UpdatedAt: time.Now(),
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
	return &items.FetchResponse{Code: 200, Text: "OK"}, nil
}

func (s *Items) List(ctx context.Context, request *items.ListRequest) (*items.ListResponse, error) {
	log.Println("List request", request)
	return s.repo.List(ctx, request)
}

func getCVSdata(url string) ([][]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, errors.New("Doesn't get CVS file. Status code now: " + res.Status)
	}
	defer res.Body.Close()

	reader := csv.NewReader(res.Body)
	reader.Comma = ';'
	result, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
