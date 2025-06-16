package graphql

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewGraphQLService(creds *credentials.Credentials) *GraphQLService {
	return &GraphQLService{
		creds,
	}
}
