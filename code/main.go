package main

import (
	"github.com/eraykalkan/gosamples/webservice/controllers"
	"net/http"
)

func main() {

	/*u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)*/

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
