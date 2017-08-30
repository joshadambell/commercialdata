package handlers

import (
	"fmt"
	"net/http"

	"github.com/joshadambell/commercialdata/internal/mongo"
	"goji.io/pat"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	collection := mongo.Session().DB("test").C("people")
	name := pat.Param(r, "name")

	if r.Method == http.MethodGet {
		var result Person
		err := collection.Find(bson.M{"name": name}).One(&result)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, "Hello, %s!", name)
		return
	}

	err := collection.Insert(&Person{Name: name})
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "%s saved!", name)

}
