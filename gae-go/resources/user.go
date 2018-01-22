package resources

import (
	"net/http"

	"github.com/google/uuid"
	util "github.com/kazukousen/gae-go/gae-go/http"
	"github.com/kazukousen/gae-go/gae-go/models"
	"google.golang.org/appengine"
)

func init() {
	http.Handle("/users/", util.Chain(util.ResourceHandler(users{})))
}

type users struct {
	util.ResourceBase
}

func (c users) Get(r *http.Request) (util.Status, interface{}) {
	id := r.URL.Path[len("/users/"):]
	if len(id) == 0 {
		id = uuid.New().String()
	}
	user, found := models.FindUser(appengine.NewContext(r), id)
	if !found {
		return util.Fail(http.StatusNotFound, id+" was created."), user
	}
	return util.Success(http.StatusOK), user
}
