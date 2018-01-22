package models

import (
	"context"
	"fmt"
	"sync"

	"github.com/kazukousen/gae-go/gae-go/gcp/datastore"
)

var (
	kind   string
	csonce sync.Once
)

func init() {
	csonce.Do(func() {
		kind = "User"
	})
}

// User is
type User struct {
	Value string
}

// FindUser returns
func FindUser(ctx context.Context, id string) (user *User, found bool) {

	key := datastore.GetNameKey(ctx, kind, id)

	user = &User{}
	if err := datastore.Get(ctx, key, user); err != nil || user == nil {
		if err.Error() == "datastore: no such entity" {
			user := &User{Value: "a"}
			if err = datastore.Put(ctx, key, user); err != nil {
				fmt.Println(err.Error())
			}
			return user, false
		}
		fmt.Println(err.Error())
		return nil, false
	}

	return user, true
}
