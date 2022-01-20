package client

import (
	abstractions "github.com/microsoft/kiota/abstractions/go"
	"github.com/microsoft/kiota/abstractions/go/serialization"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

//
//func MsgraphMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, options TestOptions) {
//	t.Helper()
//	ctrl := gomock.NewController(t)
//
//	//todo add example config
//	cfg := ""
//
//	providertest.TestResource(t, providertest.ResourceTestCase{
//		Provider: &provider.Provider{
//			Name:    "msgraph_mock_test_provider",
//			Version: "development",
//			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
//				c := NewMsgraphClient(logging.New(&hclog.LoggerOptions{
//					Level: hclog.Warn,
//				}), "testTenant")
//				c.Adapter.SetBaseUrl("http://localhost:3000")
//				c.Services = builder(t, ctrl)
//				return c, nil
//			},
//			ResourceMap: map[string]*schema.Table{
//				"test_resource": table,
//			},
//			Config: func() provider.Config {
//				return &Config{}
//			},
//		},
//		Table:          table,
//		Config:         cfg,
//		SkipEmptyJsonB: options.SkipEmptyJsonB,
//	})
//}

type TestRequestAdapter struct {
	baseUrl string
	data    map[string]serialization.Parsable
}

func NewTestRequestAdapter(url string, data map[string]serialization.Parsable) *TestRequestAdapter {
	return &TestRequestAdapter{
		baseUrl: url,
		data:    data,
	}
}

func (r *TestRequestAdapter) SendAsync(requestInfo abstractions.RequestInformation, constructor func() serialization.Parsable, responseHandler abstractions.ResponseHandler) (serialization.Parsable, error) {
	//test, err := requestInfo.GetUri()
	//if err != nil {
	//	return nil, err
	//}

	panic("not implemented")
}

// SendCollectionAsync executes the HTTP request specified by the given RequestInformation and returns the deserialized response model collection.
func (r *TestRequestAdapter) SendCollectionAsync(requestInfo abstractions.RequestInformation, constructor func() serialization.Parsable, responseHandler abstractions.ResponseHandler) ([]serialization.Parsable, error) {
	panic("not implemented")
}

// SendPrimitiveAsync executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model.
func (r *TestRequestAdapter) SendPrimitiveAsync(requestInfo abstractions.RequestInformation, typeName string, responseHandler abstractions.ResponseHandler) (interface{}, error) {
	panic("not implemented")
}

// SendPrimitiveCollectionAsync executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model collection.
func (r *TestRequestAdapter) SendPrimitiveCollectionAsync(requestInfo abstractions.RequestInformation, typeName string, responseHandler abstractions.ResponseHandler) ([]interface{}, error) {
	panic("not implemented")
}

// SendNoContentAsync executes the HTTP request specified by the given RequestInformation with no return content.
func (r *TestRequestAdapter) SendNoContentAsync(requestInfo abstractions.RequestInformation, responseHandler abstractions.ResponseHandler) error {
	panic("not implemented")
}

// GetSerializationWriterFactory returns the serialization writer factory currently in use for the request adapter service.
func (r *TestRequestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	panic("not implemented")
}

// EnableBackingStore enables the backing store proxies for the SerializationWriters and ParseNodes in use.
func (r *TestRequestAdapter) EnableBackingStore() {}

// SetBaseUrl sets the base url for every request.
func (r *TestRequestAdapter) SetBaseUrl(_ string) {}

// GetBaseUrl gets the base url for every request.
func (r *TestRequestAdapter) GetBaseUrl() string {
	return r.baseUrl
}
