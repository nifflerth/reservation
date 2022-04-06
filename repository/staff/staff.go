package staff

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/mongodb"
)

type Repository struct {
	*mongodb.Repository
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Repository, err error) {
	mongoDB, err := mongodb.New(ctx, uri, dbName, collName)
	if err != nil {
		return nil, err
	}
	return &Repository{mongoDB}, nil
}
