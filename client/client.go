package client

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/amanenk/cq-provider-msgraph/client/services"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	abstractions "github.com/microsoft/kiota/abstractions/go"
	azureAuth "github.com/microsoft/kiota/authentication/go/azure"
	microsoftgraph "github.com/microsoftgraph/msgraph-sdk-go"
)

type Client struct {
	logger   hclog.Logger
	TenantId string
	Adapter  *microsoftgraph.GraphRequestAdapter
	Services *Services
}

type Services struct {
	Groups services.GroupsClient
}

func NewMsgraphClient(log hclog.Logger, tenantId string, services *Services) *Client {
	return &Client{
		logger:   log,
		TenantId: tenantId,
		Services: services,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	settings, err := auth.GetSettingsFromEnvironment()
	if err != nil {
		return nil, err
	}

	c, err := settings.GetClientCredentials()
	if err != nil {
		return nil, err
	}

	cred, err := azidentity.NewClientSecretCredential(
		c.TenantID,
		c.ClientID,
		c.ClientSecret,
		nil,
	)
	if err != nil {
		fmt.Printf("Error creating credentials: %v\n", err)
	}

	auth, err := azureAuth.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		return nil, err
	}

	adapter, err := microsoftgraph.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}

	services := InitServices(adapter)
	client := NewMsgraphClient(logger, c.TenantID, services)
	client.Adapter = adapter

	// Return the initialized client. It will be passed to your resources
	return client, nil
}

func InitServices(requestAdapter abstractions.RequestAdapter) *Services {
	graphClient := microsoftgraph.NewGraphServiceClient(requestAdapter)
	return &Services{
		Groups: graphClient.Groups(),
	}
}
