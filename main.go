package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.Path("/api").HandlerFunc(handleGrootGraphQL)
	handler := cors.AllowAll().Handler(router)
	logrus.Fatal(http.ListenAndServe(":8000", handler))
}

var rootDevexQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"queryGroot": QueryGroot,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addGroot":    AddGroot,
		"updateGroot": UpdateGroot,
	},
})

var schemaDevex, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootDevexQuery,
	Mutation: rootMutation,
})

func executeQuery(query map[string]interface{}, schema graphql.Schema) *graphql.Result {

	params := graphql.Params{
		Schema:        schema,
		RequestString: query["query"].(string),
	}

	if query["variables"] != nil {
		params.VariableValues = query["variables"].(map[string]interface{})
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		logrus.WithFields(z.Fields(logrus.Fields{"wrong result, unexpected errors:": result.Errors})).Error(ModuleName)
	}

	return result
}

func handleGrootGraphQL(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var g map[string]interface{}
	if r.Method == http.MethodGet {
		g = make(map[string]interface{})
		g["query"] = r.URL.Query().Get("query")
		result := executeQuery(g, schemaDevex)
		json.NewEncoder(w).Encode(result)
	}

	if r.Method == http.MethodPost {
		data, _ := ioutil.ReadAll(r.Body)

		err := json.Unmarshal(data, &g)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}

		result := executeQuery(g, schemaDevex)
		json.NewEncoder(w).Encode(result)
	}
}
