package items

import (
	"time"
)

type Item struct {
	Name      string    `bson:"name"`
	Price     float64   `bson:"price"`
	Changes   int64     `bson:"changes"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type SortParams struct {
	Limit       int64
	Offset      int64
	SortingAsc  bool
	SortingName string
}
