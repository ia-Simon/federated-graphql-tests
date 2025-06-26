// Package graph implemements resolvers and auto-generated
// code for the inner workings of the GraphQL API.
package graph

import "database/sql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Datastore *sql.DB
}
