// go run github.com/99designs/gqlgen generate -v

package graph

import (
	"github.com/userq11/meetmeup/graph/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver type
type Resolver struct {
	Domain *domain.Domain
}
