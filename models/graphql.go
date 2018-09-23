package models

import (
	"../db"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"shop":           GetShopQuery(),
		"createShop":     CreateShopMutation(),
		"updateShop":     UpdateShopMutation(),
		"deleteShop":     DeleteShopMutation(),
		"product":        GetProductQuery(),
		"createProduct":  CreateProductMutation(),
		"updateProduct":  UpdateProductMutation(),
		"deleteProduct":  DeleteProductMutation(),
		"order":          GetOrderQuery(),
		"createOrder":    CreateOrderMutation(),
		"updateOrder":    UpdateOrderMutation(),
		"deleteOrder":    DeleteOrderMutation(),
		"lineItem":       GetLineItemQuery(),
		"createLineItem": CreateLineItemMutation(),
		"updateLineItem": UpdateLineItemMutation(),
		"deleteLineItem": DeleteLineItemMutation(),
	}
}

// Abstracted CRUD layers. Will require additional creative abstractions for a generic Create wrapper.
// Currently not implemented due to having issues getting children (interface to *interace conversion).
func CustomGetQuery(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(objectType),
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return GetDatabaseEntry(params.Args, object)
		},
	}
}
func CustomUpdateMutation(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: objectType,
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return UpdateDatabaseEntry(params.Args, object)
		},
	}
}
func CustomDeleteMutation(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: objectType,
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return DeleteDatabaseEntry(params.Args, object)
		},
	}
}

func GetShopQuery() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id":   &graphql.ArgumentConfig{Type: graphql.Int},
		"name": &graphql.ArgumentConfig{Type: graphql.String},
	}
	return CustomGetQuery(ShopType, args, &[]Shop{})
}
func CreateShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Create new shop",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			shop := Shop{
				Name: params.Args["name"].(string),
			}
			CreateDatabaseEntry(&shop)
			return shop, nil
		},
	}
}
func UpdateShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Update existing shop by ID",
		Args: graphql.FieldConfigArgument{
			"id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"name": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			// Have to instantiate outside DB call to properly return children

			// return UpdateDatabaseEntry(params.Args, &Shop{})

			// var shop Shop
			// return UpdateDatabaseEntry(params.Args, &shop)

			var shop Shop
			DB.First(&shop, params.Args["id"]).Updates(params.Args)
			return shop, nil
		},
	}
	// args := graphql.FieldConfigArgument{
	// 	"id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	// 	"name": &graphql.ArgumentConfig{Type: graphql.String},
	// }
	// return CustomUpdateMutation(ShopType, args, &Shop{})
}
func DeleteShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Delete existing shop by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// return DeleteDatabaseEntry(params.Args, &Shop{})

			var shop Shop
			DB.Where(params.Args).Delete(&shop)
			return shop, nil
		},
	}
	// args := graphql.FieldConfigArgument{
	// 	"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	// }
	// return CustomDeleteMutation(ShopType, args, &Shop{})
}

func GetProductQuery() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id":       &graphql.ArgumentConfig{Type: graphql.Int},
		"name":     &graphql.ArgumentConfig{Type: graphql.String},
		"shop_id":  &graphql.ArgumentConfig{Type: graphql.Int},
		"value":    &graphql.ArgumentConfig{Type: graphql.Int},
		"quantity": &graphql.ArgumentConfig{Type: graphql.Int},
	}
	return CustomGetQuery(ProductType, args, &[]Product{})
}
func CreateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ProductType,
		Description: "Create new product",
		Args: graphql.FieldConfigArgument{
			"shop_id":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"value":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"quantity": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			product := Product{
				ShopID:   params.Args["shop_id"].(int),
				Name:     params.Args["name"].(string),
				Value:    params.Args["value"].(int),
				Quantity: params.Args["quantity"].(int),
			}
			CreateDatabaseEntry(&product)
			return product, nil
		},
	}
}
func UpdateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ProductType,
		Description: "Update existing product by ID",
		Args: graphql.FieldConfigArgument{
			"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"shop_id":  &graphql.ArgumentConfig{Type: graphql.Int},
			"name":     &graphql.ArgumentConfig{Type: graphql.String},
			"value":    &graphql.ArgumentConfig{Type: graphql.Int},
			"quantity": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var product Product
			DB.First(&product, params.Args["id"]).Updates(params.Args)
			return product, nil
		},
	}
}
func DeleteProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ProductType,
		Description: "Delete existing product by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var product Product
			DB.Where(params.Args).Delete(&product)
			return product, nil
		},
	}
}

func GetOrderQuery() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id":      &graphql.ArgumentConfig{Type: graphql.Int},
		"shop_id": &graphql.ArgumentConfig{Type: graphql.Int},
	}
	return CustomGetQuery(OrderType, args, &[]Order{})
}
func CreateOrderMutation() *graphql.Field {
	return &graphql.Field{
		Type:        OrderType,
		Description: "Create new order",
		Args: graphql.FieldConfigArgument{
			"shop_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			order := Order{
				ShopID: params.Args["shop_id"].(int),
			}
			CreateDatabaseEntry(&order)
			return order, nil
		},
	}
}
func UpdateOrderMutation() *graphql.Field {
	return &graphql.Field{
		Type:        OrderType,
		Description: "Update existing order by ID",
		Args: graphql.FieldConfigArgument{
			"id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"shop_id": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var order Order
			DB.First(&order, params.Args["id"]).Updates(params.Args)
			return order, nil
		},
	}
}
func DeleteOrderMutation() *graphql.Field {
	return &graphql.Field{
		Type:        OrderType,
		Description: "Delete existing order by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var order Order
			DB.Where(params.Args).Delete(&order)
			return order, nil
		},
	}
}

func GetLineItemQuery() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.Int},
	}
	return CustomGetQuery(LineItemType, args, &[]LineItem{})
}
func CreateLineItemMutation() *graphql.Field {
	return &graphql.Field{
		Type:        LineItemType,
		Description: "Create new line item",
		Args: graphql.FieldConfigArgument{
			"product_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"order_id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"quantity":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			lineItem := LineItem{
				ProductID: params.Args["product_id"].(int),
				OrderID:   params.Args["order_id"].(int),
				Quantity:  params.Args["quantity"].(int),
			}
			CreateDatabaseEntry(&lineItem)
			return lineItem, nil
		},
	}
}
func UpdateLineItemMutation() *graphql.Field {
	return &graphql.Field{
		Type:        LineItemType,
		Description: "Update existing line item by ID",
		Args: graphql.FieldConfigArgument{
			"id":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"product_id": &graphql.ArgumentConfig{Type: graphql.Int},
			"order_id":   &graphql.ArgumentConfig{Type: graphql.Int},
			"quantity":   &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var lineItem LineItem
			DB.First(&lineItem, params.Args["id"]).Updates(params.Args)
			return lineItem, nil
		},
	}
}
func DeleteLineItemMutation() *graphql.Field {
	return &graphql.Field{
		Type:        LineItemType,
		Description: "Delete existing line item by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var lineItem LineItem
			DB.Where(params.Args).Delete(&lineItem)
			return lineItem, nil
		},
	}
}

// Methods to interface with Database through GORM
func GetDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Find(entry)
	return entry, nil
}
func CreateDatabaseEntry(entry interface{}) {
	DB.Create(entry)
}
func UpdateDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.First(entry, args["id"]).Updates(args)
	return entry, nil
}
func DeleteDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Delete(entry)
	return entry, nil
}
func GetChildrenGeneric(parent, children interface{}) (interface{}, error) {
	DB.Model(parent).Related(children)
	return children, nil
}
func GetTotalLineItemPrice(id int, quantity int) int {
	var product Product
	DB.First(&product, id)
	return product.Value * quantity
}
func GetTotalOrderPrice(order Order) int {
	var lineItems []LineItem
	var sum int
	DB.Model(order).Related(&lineItems)
	for _, item := range lineItems {
		sum += GetTotalLineItemPrice(item.ProductID, item.Quantity)
	}
	return sum
}

// Set the database models to our database struct.
// Set the graphQL schemas, and return the handler for use with our graphQL endpoint.
var GraphqlHandler http.Handler

func init() {
	DB = db.DB
	DB.AutoMigrate(&Shop{}, &Product{}, &Order{}, &LineItem{})

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: GetRootFields(),
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}
	GraphqlHandler = handler.New(&handler.Config{
		Schema: &schema,
	})
}
