package repository

import (
	"context"
	items "github.com/Arkosh744/grpc_files_server/gen/item"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Items struct {
	db *mongo.Database
}

func NewItems(db *mongo.Database) *Items {
	return &Items{
		db: db,
	}
}

func (r *Items) GetByName(ctx context.Context, name string) (items.Item, error) {
	var result items.Item
	err := r.db.Collection("items").FindOne(ctx, bson.M{"name": name}).Decode(&result)
	return result, err
}

func (r *Items) InsertOne(ctx context.Context, item items.Item) error {
	_, err := r.db.Collection("items").InsertOne(ctx, item)
	return err
}

func (r *Items) UpdateOne(ctx context.Context, item items.Item) error {
	_, err := r.db.Collection("items").UpdateOne(ctx, bson.M{"name": item.Name}, bson.M{
		"$set": bson.M{"price": item.Price, "updated_at": item.UpdatedAt},
		"$inc": bson.M{"changes": 1},
	})
	return err
}

func (r *Items) List(ctx context.Context, request *items.ListRequest) (*items.ListResponse, error) {
	findOpt := options.Find()
	if request.SortingAsc == true {
		findOpt.SetSort(bson.D{{request.SortingName, 1}})
	} else {
		findOpt.SetSort(bson.D{{request.SortingName, -1}})
	}
	findOpt.SetLimit(request.Limit)
	findOpt.SetSkip(request.Offset)
	cursor, err := r.db.Collection("items").Find(ctx, bson.D{}, findOpt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var resultItems []*items.ListResponseItems
	if err := cursor.All(ctx, &resultItems); err != nil {
		return nil, err
	}
	resultList := &items.ListResponse{Items: resultItems}
	return resultList, nil
}
