package pagerduty

import (
	"context"
	"net/http"
	"testing"
)

// Create Slack Connection
func TestSlackConnection_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write([]byte(`{"slack_connection":{"id":"1"}}`))
	})

	client := defaultTestClient(server.URL, "foo")
	input := SlackConnection{}

	res, err := client.CreateSlackConnectionWithContext(context.Background(), "1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SlackConnection{ID: "1"}
	testEqual(t, want, res)
}

// List Slack Connections
func TestSlackConnection_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"slack_connections":[{"id":"1"}]}`))
	})

	client := defaultTestClient(server.URL, "foo")
	o := ListSlackConnectionsOptions{
		Limit:  0,
		Offset: 0,
	}

	res, err := client.ListSlackConnectionsWithContext(context.Background(), "1", o)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListSlackConnectionsResponse{
		APIListObject:    APIListObject{},
		SlackConnections: []SlackConnection{{ID: "1"}},
	}
	testEqual(t, want, res)
}

// Get Slack Connection
func TestSlackConnection_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{
  "slack_connection": {
    "id": "A12BCDE",
    "source_id": "A1234B5",
    "source_name": "test_team",
    "source_type": "team_reference",
    "channel_id": "A123B456C7D",
    "channel_name": "random",
    "notification_type": "responder",
    "config": {
      "events": ["incident.acknowledged"],
      "priorities": ["ABCDEF1"],
      "urgency": "high"
    }
  }
}`))
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetSlackConnectionWithContext(context.Background(), "1", "2")
	if err != nil {
		t.Fatal(err)
	}

	wantConfigUrgency := "high"
	want := &SlackConnection{
		ID:               "A12BCDE",
		SourceID:         "A1234B5",
		SourceName:       "test_team",
		SourceType:       "team_reference",
		ChannelID:        "A123B456C7D",
		ChannelName:      "random",
		NotificationType: "responder",
		Config: &SlackConnectionConfig{
			Events:     []string{"incident.acknowledged"},
			Urgency:    &wantConfigUrgency,
			Priorities: []string{"ABCDEF1"},
		},
	}
	testEqual(t, want, res)
}

// Update Slack Connection
func TestSlackConnection_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, _ = w.Write([]byte(`{"slack_connection":{"id":"1"}}`))
	})

	client := defaultTestClient(server.URL, "foo")
	input := SlackConnection{ID: "1"}

	res, err := client.UpdateSlackConnectionWithContext(context.Background(), "1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SlackConnection{ID: "1"}
	testEqual(t, want, res)
}

// Delete Slack Connection
func TestSlackConnection_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		_, _ = w.Write([]byte(`{"slack_connection":{"id":"1"}}`))
	})

	client := defaultTestClient(server.URL, "foo")

	err := client.DeleteSlackConnectionWithContext(context.Background(), "1", "1")
	if err != nil {
		t.Fatal(err)
	}
}
