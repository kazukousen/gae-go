package resources

import (
	"net/http"

	util "github.com/kazukousen/gae-go/gae-go/http"
	"github.com/kazukousen/gae-go/gae-go/models"
	"google.golang.org/appengine"
)

func init() {
	http.Handle("/users/", util.Chain(util.ResourceHandler(user{})))
}

type user struct {
	util.ResourceBase
}

func (u user) Get(r *http.Request) (util.Status, interface{}) {
	if id := r.URL.Path[len("/users/"):]; len(id) != 0 {
		u, found := models.FindUser(appengine.NewContext(r), id)
		if !found {
			return util.Success(http.StatusOK), u
		}
		return util.Success(http.StatusOK), u
	}
	return util.FailSimple(http.StatusBadRequest), models.User{}
}
