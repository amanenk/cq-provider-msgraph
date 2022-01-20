package services

import (
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
)

//go:generate mockgen -package=mocks -destination=./mocks/groups.go . GroupsClient
type GroupsClient interface {
	Get(options *groups.GroupsRequestBuilderGetOptions) (*groups.GroupsResponse, error)
}
