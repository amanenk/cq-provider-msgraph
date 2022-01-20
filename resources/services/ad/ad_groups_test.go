package ad

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/amanenk/cq-provider-msgraph/client"
	"github.com/amanenk/cq-provider-msgraph/client/services/mocks"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	azureAuth "github.com/microsoft/kiota/authentication/go/azure"
	microsoftgraph "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createADGroupsTestServer(t *testing.T) (*client.Services, error) {
	mux := httprouter.New()
	mux.GET("/v1.0/groups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		groups := []msgraph.Group{
			*fakeGroup(t),
		}

		value, err := json.Marshal(groups)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		resp := msgraph.Paging{
			NextLink: "",
			Value:    value,
		}

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	cred, err := azidentity.NewClientSecretCredential(
		"test",
		"test",
		"test",
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

	ts := httptest.NewTLSServer(mux)
	adapter.SetBaseUrl(ts.URL)
	services := client.InitServices(adapter)
	return services, nil
}

func fakeGroup(t *testing.T) *msgraph.Group {
	var group msgraph.Group
	if err := faker.FakeDataSkipFields(&group, []string{
		"Conversations",
		"Threads",
		"CalendarView",
		"Events",
		"Calendar",
		"Drive",
		"Drives",
		"Sites",
		"Onenote",
	}); err != nil {
		t.Fatal(err)
	}

	group.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	group.Conversations = []msgraph.Conversation{fakeConversation(t)}
	group.Calendar = fakeCalendar(t)
	group.CalendarView = []msgraph.Event{fakeEvent(t)}
	group.Events = []msgraph.Event{fakeEvent(t)}
	group.Drive = fakeDrive(t)
	group.Drives = []msgraph.Drive{*fakeDrive(t)}
	group.Sites = []msgraph.Site{*fakeSite(t)}
	group.Onenote = fakeOnenote(t)
	return &group
}

func fakeOnenote(t *testing.T) *msgraph.Onenote {
	e := msgraph.Onenote{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Notebooks",
		"Sections",
		"SectionGroups",
		"Pages",
	}); err != nil {
		t.Fatal(err)
	}

	return &e
}
func fakeSite(t *testing.T) *msgraph.Site {
	e := msgraph.Site{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"BaseItem",
		"Drive",
		"Drives",
		"Items",
		"Item",
		"Lists",
		"Sites",
		"Onenote",
		"Analytics",
	}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}
func fakeDrive(t *testing.T) *msgraph.Drive {
	e := msgraph.Drive{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"BaseItem",
		"Special",
		"Items",
		"List",
		"Root",
	}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}

func fakeEvent(t *testing.T) msgraph.Event {
	e := msgraph.Event{}
	if err := faker.FakeData(&e.OutlookItem); err != nil {
		t.Fatal(err)
	}
	e.Calendar = fakeCalendar(t)
	return e
}

func fakeCalendar(t *testing.T) *msgraph.Calendar {
	e := msgraph.Calendar{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Events",
		"CalendarView",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}

func fakeConversationThread(t *testing.T) msgraph.ConversationThread {
	e := msgraph.ConversationThread{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Posts",
	}); err != nil {
		t.Fatal(err)
	}
	return e
}

func fakeConversation(t *testing.T) msgraph.Conversation {
	e := msgraph.Conversation{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Threads",
	}); err != nil {
		t.Fatal(err)
	}
	e.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	return e
}

func buildAdGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGroupsClient(ctrl)

	services := client.Services{
		Groups: m,
	}
	groups := groups.GroupsResponse{}

	if err := faker.FakeData(&groups); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().Get(gomock.Any()).Return(
		&groups,
		nil,
	)

	return services
}

func TestADGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	services := buildAdGroupsMock(t, ctrl)
	t.Fatal(services)
	//client.MsgraphMockTestHelper(t, AdGroups(), buildAdGroupsMock, client.TestOptions{})
}

//
//func TestADGroups(t *testing.T) {
//	ADGroupsHelper(t, AdGroups(), createADGroupsTestServer, client.TestOptions{})
//}

func ADGroupsHelper(t *testing.T, table *schema.Table, builder func(*testing.T) (*client.Services, error), options client.TestOptions) { //todo add example config
	cfg := ""
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "msgraph_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				services, err := builder(t)
				if err != nil {
					t.Fatal(err)
				}
				c := client.NewMsgraphClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), "testTenant", services)

				return c, nil
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
			Config: func() provider.Config {
				return &client.Config{}
			},
		},
		Table:          table,
		Config:         cfg,
		SkipEmptyJsonB: options.SkipEmptyJsonB,
	})

}
