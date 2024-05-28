package api

import (
	"context"
	"github.com/zeabur/cli/pkg/model"
)

func (c *Client) CreateProject(ctx context.Context, region string, name *string) (*model.Project, error) {
	var mutation struct {
		CreateProject model.Project `graphql:"createProject(region: $region, name: $name)"`
	}

	err := c.Mutate(ctx, &mutation, map[string]interface{}{
		"region": region,
		"name":   name,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.CreateProject, nil
}

func (c *Client) GetProject(ctx context.Context, projectID string) (*model.Project, error) {
	var query struct {
		Project model.Project `graphql:"project(_id: $id)"`
	}

	err := c.Query(ctx, &query, map[string]interface{}{
		"id": ObjectID(projectID),
	})

	if err != nil {
		return nil, err
	}

	return &query.Project, nil
}

// GetProjectByOwnerAndName returns a project by its owner and name.
func (c *Client) GetProjectByOwnerAndName(ctx context.Context, owner, name string) (*model.Project, error) {
	var query struct {
		Project model.Project `graphql:"project(owner: $owner, name: $name)"`
	}

	err := c.Query(ctx, &query, map[string]interface{}{
		"owner": owner,
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	return &query.Project, nil
}

// DeleteProject deletes a project by its ID.
func (c *Client) DeleteProject(ctx context.Context, id string) error {
	var mutation struct {
		DeleteProject bool `graphql:"deleteProject(_id: $id)"`
	}

	err := c.Mutate(ctx, &mutation, map[string]interface{}{
		"id": ObjectID(id),
	})

	if err != nil {
		return err
	}

	return nil
}
