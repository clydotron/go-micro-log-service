package dao

import (
	"github.com/clydotron/go-micro-log-service/models"

	dao_mongo "github.com/clydotron/go-micro-log-service/dao/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

// #TODO move interfaces into own package?
type LogRepo interface {
	Insert(entry models.LogEntry) error
	All() ([]*models.LogEntry, error)
}

type DataStore struct {
	LogRepo LogRepo
}

func NewDataStore(client *mongo.Client) DataStore {
	return DataStore{
		LogRepo: dao_mongo.NewLogRepo(client),
	}
}
