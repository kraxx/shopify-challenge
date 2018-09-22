package queries

import (
	"../types"
	"github.com/graphql-go/graphql"
)

func GetShopQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.ShopType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var shops []types.Shop

			// How to obtain data herrrrr
			return shops, nil
		},
	}
}
