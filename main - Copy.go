package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	_ "os"
)

/*
Shops have many Products
Shops have many Orders
Products have many Line Items
Orders have many Line Items
*/
/*
type Shop struct {
  // gorm.Model
  ID       int       `json:"id,omitempty"`
  Products []Product `json:"products"`
  Orders   []Order   `json:"orders"`
  Name     string    `json:"name"`
}

type Product struct {
  // gorm.Model
  ID       int `json:"id,omitempty"`
  Shop     `json:"shop"`
  Items    []LineItem `json:"items"`
  Name     string     `json:"name"`
  Price    int        `json:"price"`
  Quantity int        `json:"quantity"`
}

type Order struct {
  // gorm.Model
  ID    int `json:"id,omitempty"`
  Shop  `json:"shop"`
  Items []LineItem `json:"items"`
}

type LineItem struct {
  // gorm.Model
  ID       int `json:"id,omitempty"`
  Product  `json:"product"`
  Order    `json:"order"`
  Quantity int `json:"quantity"`
}
*/
/*
type Shop struct {
	ID       string `json:"id,omitempty"`
	Products string `json:"products"`
	Orders   string `json:"orders"`
	Name     string `json:"name"`
}

type Product struct {
	ID       string `json:"id,omitempty"`
	Shop     string `json:"shop"`
	Items    string `json:"items"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Order struct {
	ID    string `json:"id,omitempty"`
	Shop  string `json:"shop"`
	Items string `json:"items"`
}

type LineItem struct {
	ID       string `json:"id,omitempty"`
	Product  string `json:"product"`
	Order    string `json:"order"`
	Quantity string `json:"quantity"`
}
*/


type Shop struct {
  Id       int       `json:"id,omitempty"`
  Products []*Product `json:"products"`
  Orders   []*Order   `json:"orders"`
  Name     string    `json:"name"`
}

type Product struct {
  Id       int `json:"id,omitempty"`
  *Shop     `json:"shop"`
  Items    []*LineItem `json:"items"`
  Name     string     `json:"name"`
  Price    int        `json:"price"`
  Quantity int        `json:"quantity"`
}

type Order struct {
  Id    int `json:"id,omitempty"`
  *Shop  `json:"shop"`
  Items []*LineItem `json:"items"`
}

type LineItem struct {
  Id       int `json:"id,omitempty"`
  *Product  `json:"product"`
  *Order    `json:"order"`
  Quantity int `json:"quantity"`
}

// var db *gorm.DB

/*
func graphqlHandler(schema graphql.Schema) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	}
}
*/

func main() {
	// router := mux.NewRouter()
	// db, err := gorm.Open(
	// 	"postgres",
	// 	"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
	// 		" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+
	// 		os.Getenv("PASSWORD"))

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// defer db.Close()

	// db.AutoMigrate(&Resource{})

	// router.HandleFunc("/resources", GetResources).Methods("GET")
	// router.HandleFunc("/resources/{id}", GetResource).Methods("GET")
	// router.HandleFunc("/resources", CreateResource).Methods("POST")
	// router.HandleFunc("/resources/{id}", DeleteResource).Methods("DELETE")

	// log.Fatal(http.ListenAndServe(":"+os.Genenv("PORT"), router))

	// schema, err := graphql.NewSchema(graphql.SchemaConfig{})
	// http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	// 	result := graphql.Do(graphql.Params{
	// 		Schema:        schema,
	// 		RequestString: r.URL.Query().Get("query"),
	// 	})
	// 	json.NewEncoder(w).Encode(result)
	// })

/*
	fmt.Println("LEZ BEGIN")
	shopType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"products": &graphql.Field{
				Type: graphql.String,
			},
      "orders": &graphql.Field{
        Type: graphql.String,
      },
      "name": &graphql.Field{
        Type: graphql.String,
      },
		},
	})
	productType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Product",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "shop": &graphql.Field{
        Type: graphql.String,
      },
      "items": &graphql.Field{
        Type: graphql.String,
      },
      "name": &graphql.Field{
        Type: graphql.String,
      },
      "price": &graphql.Field{
        Type: graphql.String,
      },
      "quantity": &graphql.Field{
        Type: graphql.String,
      },
    },
	})
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Order",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "shop": &graphql.Field{
        Type: graphql.String,
      },
      "items": &graphql.Field{
        Type: graphql.String,
      },
	})
	lineItemType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "LineItem",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "product": &graphql.Field{
        Type: graphql.String,
      },
      "order": &graphql.Field{
        Type: graphql.String,
      },
      "quantity": &graphql.Field{
        Type: graphql.String,
      },
	})
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: graphql.Fields{},
	})
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: graphql.Fields{},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
*/

    fmt.Println("LEZ BEGIN")
  shopType := graphql.NewObject(graphql.ObjectConfig{
    Name: "Shop",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "products": &graphql.Field{
        Type: graphql.String,
      },
      "orders": &graphql.Field{
        Type: graphql.String,
      },
      "name": &graphql.Field{
        Type: graphql.String,
      },
    },
  })
  productType := graphql.NewObject(graphql.ObjectConfig{
    Name:   "Product",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "shop": &graphql.Field{
        Type: graphql.String,
      },
      "items": &graphql.Field{
        Type: graphql.String,
      },
      "name": &graphql.Field{
        Type: graphql.String,
      },
      "price": &graphql.Field{
        Type: graphql.String,
      },
      "quantity": &graphql.Field{
        Type: graphql.String,
      },
    },
  })
  orderType := graphql.NewObject(graphql.ObjectConfig{
    Name:   "Order",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "shop": &graphql.Field{
        Type: graphql.String,
      },
      "items": &graphql.Field{
        Type: graphql.String,
      },
  })
  lineItemType := graphql.NewObject(graphql.ObjectConfig{
    Name:   "LineItem",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.String,
      },
      "product": &graphql.Field{
        Type: graphql.String,
      },
      "order": &graphql.Field{
        Type: graphql.String,
      },
      "quantity": &graphql.Field{
        Type: graphql.String,
      },
  })
  rootQuery := graphql.NewObject(graphql.ObjectConfig{
    Name:   "RootQuery",
    Fields: graphql.Fields{},
  })
  rootMutation := graphql.NewObject(graphql.ObjectConfig{
    Name:   "RootMutation",
    Fields: graphql.Fields{},
  })
  schema, _ := graphql.NewSchema(graphql.SchemaConfig{
    Query:    rootQuery,
    Mutation: rootMutation,
  })
	http.HandleFunc("/graphql", graphqlHandler(schema))
	log.Println("Listenin and servin baby")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
