package pagerduty

import (
	"context"

	"github.com/google/go-querystring/query"
)

type SlackConnection struct {
	ID               string                 `json:"id,omitempty"`
	SourceID         string                 `json:"source_id"`
	SourceName       string                 `json:"source_name"`
	SourceType       string                 `json:"source_type"`
	ChannelID        string                 `json:"channel_id"`
	ChannelName      string                 `json:"channel_name"`
	NotificationType string                 `json:"notification_type"`
	Config           *SlackConnectionConfig `json:"config"`
}

type SlackConnectionConfig struct {
	Events     []string `json:"events"`
	Urgency    *string  `json:"urgency"`
	Priorities []string `json:"priorities"`
}

type SlackConnectionResponse struct {
	SlackConnection *SlackConnection `json:"slack_connection"`
}

func (c *Client) CreateSlackConnectionWithContext(ctx context.Context, id string, conn SlackConnection) (*SlackConnection, error) {
	d := map[string]SlackConnection{
		"slack_connection": conn,
	}

	resp, err := c.post(ctx, "/integration-slack/workspaces/"+id+"/connections", d, nil)
	if err != nil {
		return nil, err
	}

	var response SlackConnectionResponse
	if err := c.decodeJSON(resp, &response); err != nil {
		return nil, err
	}

	return response.SlackConnection, nil
}

func (c *Client) GetSlackConnectionWithContext(ctx context.Context, teamID, connID string) (*SlackConnection, error) {
	resp, err := c.get(ctx, "/integration-slack/workspaces/"+teamID+"/connections/"+connID, nil)
	if err != nil {
		return nil, err
	}

	var response SlackConnectionResponse
	if err := c.decodeJSON(resp, &response); err != nil {
		return nil, err
	}

	return response.SlackConnection, nil
}

func (c *Client) UpdateSlackConnectionWithContext(ctx context.Context, teamID string, conn SlackConnection) (*SlackConnection, error) {
	d := map[string]SlackConnection{
		"slack_connection": conn,
	}

	resp, err := c.put(ctx, "/integration-slack/workspaces/"+teamID+"/connections/"+conn.ID, d, nil)
	if err != nil {
		return nil, err
	}

	var response SlackConnectionResponse
	if err := c.decodeJSON(resp, &response); err != nil {
		return nil, err
	}

	return response.SlackConnection, nil
}

func (c *Client) DeleteSlackConnectionWithContext(ctx context.Context, teamID, connID string) error {
	_, err := c.delete(ctx, "/integration-slack/workspaces/"+teamID+"/connections/"+connID)
	return err
}

type ListSlackConnectionsOptions struct {
	Limit  uint `url:"limit,omitempty"`
	Offset uint `url:"offset,omitempty"`
}

type ListSlackConnectionsResponse struct {
	APIListObject
	SlackConnections []SlackConnection `json:"slack_connections"`
}

func (c *Client) ListSlackConnectionsWithContext(ctx context.Context, teamID string, o ListSlackConnectionsOptions) (*ListSlackConnectionsResponse, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	resp, err := c.get(ctx, "/integration-slack/workspaces/"+teamID+"/connections?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var response ListSlackConnectionsResponse
	if err := c.decodeJSON(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
