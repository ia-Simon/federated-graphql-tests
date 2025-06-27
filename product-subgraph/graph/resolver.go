// Package graph holds all the generated code + implementations for
// the functioning of the product-subgraph Graphql API.
package graph

import "database/sql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Datastore *sql.DB
}
