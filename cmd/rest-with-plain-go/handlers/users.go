package handlers

import (
	"adzo261/backend-with-go/cmd/rest-with-plain-go/models"
	"adzo261/backend-with-go/cmd/rest-with-plain-go/utils"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Users is an http.Handler with logger Depedency Injected so that it can be changed as needed.
type Users struct {
	logger *log.Logger
}

func NewUsers(logger *log.Logger) *Users {
	return &Users{logger}
}

//Handler should implement ServeHTTP method as it satisfies Handler interface
func (uh *Users) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		uh.getUsers(w, r)
		return
	}
	if r.Method == http.MethodPost {
		uh.addUser(w, r)
		return
	}
	if r.Method == http.MethodPut {
		uh.logger.Println("PUT", r.URL.Path)

		// Catch '/id' from the url
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		uh.logger.Printf("Found Matches: %s\n", g)

		//If found more than one `/id` format, url is invalid
		if len(g) != 1 {
			uh.logger.Println("Invalid URI -  More than one id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		//Found no capture group
		if len(g[0]) != 2 {
			uh.logger.Println("Invalid URI - No capture group found")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			uh.logger.Println("Invalid URI unable to convert to numer", idString)
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		uh.updateUser(id, w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (uh *Users) getUsers(w http.ResponseWriter, r *http.Request) {
	uh.logger.Println("Handle GET Users")

	users := models.GetUsers()
	err := utils.ToJSON(w, users)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusBadRequest)
	}
}

func (uh *Users) addUser(w http.ResponseWriter, r *http.Request) {
	uh.logger.Println("Handle POST User")

	user := &models.User{}
	err := utils.FromJSON(r.Body, user)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	models.AddUser(user)
}

func (uh *Users) updateUser(id int, w http.ResponseWriter, r *http.Request) {
	uh.logger.Println("Handle Put User")

	user := &models.User{}
	err := utils.FromJSON(r.Body, user)
	fmt.Println(user, err)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = models.UpdateUser(id, user)
	if err == models.ErrUserNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
