package datastore

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Get retrives Datastore Entity
func Get(ctx context.Context, key *datastore.Key, i interface{}) error {
	err := datastore.Get(ctx, key, i)
	return err
}

// Put returns Datastore Entity
func Put(ctx context.Context, key *datastore.Key, i interface{}) error {
	_, err := datastore.Put(ctx, key, i)
	return err
}

// GetNameKey returns converted datastore.Key from string
func GetNameKey(ctx context.Context, kind string, name string) *datastore.Key {
	return datastore.NewKey(ctx, kind, name, 0, nil)
}
