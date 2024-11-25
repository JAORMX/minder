// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: datasources.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const addDataSourceFunction = `-- name: AddDataSourceFunction :one

INSERT INTO data_sources_functions (data_source_id, project_id, name, type, definition)
VALUES ($1, $2, $3, $4, $5) RETURNING id, name, type, data_source_id, definition, created_at, updated_at, project_id
`

type AddDataSourceFunctionParams struct {
	DataSourceID uuid.UUID       `json:"data_source_id"`
	ProjectID    uuid.UUID       `json:"project_id"`
	Name         string          `json:"name"`
	Type         string          `json:"type"`
	Definition   json.RawMessage `json:"definition"`
}

// AddDataSourceFunction adds a function to a datasource.
func (q *Queries) AddDataSourceFunction(ctx context.Context, arg AddDataSourceFunctionParams) (DataSourcesFunction, error) {
	row := q.db.QueryRowContext(ctx, addDataSourceFunction,
		arg.DataSourceID,
		arg.ProjectID,
		arg.Name,
		arg.Type,
		arg.Definition,
	)
	var i DataSourcesFunction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.DataSourceID,
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}

const createDataSource = `-- name: CreateDataSource :one

INSERT INTO data_sources (project_id, name, display_name)
VALUES ($1, $2, $3) RETURNING id, name, display_name, project_id, created_at, updated_at
`

type CreateDataSourceParams struct {
	ProjectID   uuid.UUID `json:"project_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
}

// CreateDataSource creates a new datasource in a given project.
func (q *Queries) CreateDataSource(ctx context.Context, arg CreateDataSourceParams) (DataSource, error) {
	row := q.db.QueryRowContext(ctx, createDataSource, arg.ProjectID, arg.Name, arg.DisplayName)
	var i DataSource
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DisplayName,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteDataSource = `-- name: DeleteDataSource :one
DELETE FROM data_sources
WHERE id = $1 AND project_id = $2
RETURNING id, name, display_name, project_id, created_at, updated_at
`

type DeleteDataSourceParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) DeleteDataSource(ctx context.Context, arg DeleteDataSourceParams) (DataSource, error) {
	row := q.db.QueryRowContext(ctx, deleteDataSource, arg.ID, arg.ProjectID)
	var i DataSource
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DisplayName,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteDataSourceFunction = `-- name: DeleteDataSourceFunction :one
DELETE FROM data_sources_functions
WHERE data_source_id = $1 AND name = $2 AND project_id = $3
RETURNING id, name, type, data_source_id, definition, created_at, updated_at, project_id
`

type DeleteDataSourceFunctionParams struct {
	DataSourceID uuid.UUID `json:"data_source_id"`
	Name         string    `json:"name"`
	ProjectID    uuid.UUID `json:"project_id"`
}

func (q *Queries) DeleteDataSourceFunction(ctx context.Context, arg DeleteDataSourceFunctionParams) (DataSourcesFunction, error) {
	row := q.db.QueryRowContext(ctx, deleteDataSourceFunction, arg.DataSourceID, arg.Name, arg.ProjectID)
	var i DataSourcesFunction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.DataSourceID,
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}

const getDataSource = `-- name: GetDataSource :one

SELECT id, name, display_name, project_id, created_at, updated_at FROM data_sources
WHERE id = $1 AND project_id = ANY($2::uuid[])
`

type GetDataSourceParams struct {
	ID       uuid.UUID   `json:"id"`
	Projects []uuid.UUID `json:"projects"`
}

// GetDataSource retrieves a datasource by its id and a project hierarchy.
//
// Note that to get a datasource for a given project, one can simply
// pass one project id in the project_id array.
func (q *Queries) GetDataSource(ctx context.Context, arg GetDataSourceParams) (DataSource, error) {
	row := q.db.QueryRowContext(ctx, getDataSource, arg.ID, pq.Array(arg.Projects))
	var i DataSource
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DisplayName,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDataSourceByName = `-- name: GetDataSourceByName :one

SELECT id, name, display_name, project_id, created_at, updated_at FROM data_sources
WHERE name = $1 AND project_id = ANY($2::uuid[])
`

type GetDataSourceByNameParams struct {
	Name     string      `json:"name"`
	Projects []uuid.UUID `json:"projects"`
}

// GetDataSourceByName retrieves a datasource by its name and
// a project hierarchy.
//
// Note that to get a datasource for a given project, one can simply
// pass one project id in the project_id array.
func (q *Queries) GetDataSourceByName(ctx context.Context, arg GetDataSourceByNameParams) (DataSource, error) {
	row := q.db.QueryRowContext(ctx, getDataSourceByName, arg.Name, pq.Array(arg.Projects))
	var i DataSource
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DisplayName,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDataSourceFunctions = `-- name: ListDataSourceFunctions :many

SELECT id, name, type, data_source_id, definition, created_at, updated_at, project_id FROM data_sources_functions
WHERE data_source_id = $1 AND project_id = $2
`

type ListDataSourceFunctionsParams struct {
	DataSourceID uuid.UUID `json:"data_source_id"`
	ProjectID    uuid.UUID `json:"project_id"`
}

// ListDataSourceFunctions retrieves all functions for a datasource.
func (q *Queries) ListDataSourceFunctions(ctx context.Context, arg ListDataSourceFunctionsParams) ([]DataSourcesFunction, error) {
	rows, err := q.db.QueryContext(ctx, listDataSourceFunctions, arg.DataSourceID, arg.ProjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DataSourcesFunction{}
	for rows.Next() {
		var i DataSourcesFunction
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.DataSourceID,
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDataSources = `-- name: ListDataSources :many

SELECT id, name, display_name, project_id, created_at, updated_at FROM data_sources
WHERE project_id = ANY($1::uuid[])
`

// ListDataSources retrieves all datasources for project hierarchy.
//
// Note that to get a datasource for a given project, one can simply
// pass one project id in the project_id array.
func (q *Queries) ListDataSources(ctx context.Context, projects []uuid.UUID) ([]DataSource, error) {
	rows, err := q.db.QueryContext(ctx, listDataSources, pq.Array(projects))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DataSource{}
	for rows.Next() {
		var i DataSource
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.DisplayName,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDataSource = `-- name: UpdateDataSource :one

UPDATE data_sources
SET display_name = $3
WHERE id = $1 AND project_id = $2
RETURNING id, name, display_name, project_id, created_at, updated_at
`

type UpdateDataSourceParams struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"project_id"`
	DisplayName string    `json:"display_name"`
}

// UpdateDataSource updates a datasource in a given project.
func (q *Queries) UpdateDataSource(ctx context.Context, arg UpdateDataSourceParams) (DataSource, error) {
	row := q.db.QueryRowContext(ctx, updateDataSource, arg.ID, arg.ProjectID, arg.DisplayName)
	var i DataSource
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DisplayName,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateDataSourceFunction = `-- name: UpdateDataSourceFunction :one

UPDATE data_sources_functions
SET type = $3, definition = $4, updated_at = NOW()
WHERE data_source_id = $1 AND project_id = $5 AND name = $2
RETURNING id, name, type, data_source_id, definition, created_at, updated_at, project_id
`

type UpdateDataSourceFunctionParams struct {
	DataSourceID uuid.UUID       `json:"data_source_id"`
	Name         string          `json:"name"`
	Type         string          `json:"type"`
	Definition   json.RawMessage `json:"definition"`
	ProjectID    uuid.UUID       `json:"project_id"`
}

// UpdateDataSourceFunction updates a function in a datasource. We're
// only able to update the type and definition of the function.
func (q *Queries) UpdateDataSourceFunction(ctx context.Context, arg UpdateDataSourceFunctionParams) (DataSourcesFunction, error) {
	row := q.db.QueryRowContext(ctx, updateDataSourceFunction,
		arg.DataSourceID,
		arg.Name,
		arg.Type,
		arg.Definition,
		arg.ProjectID,
	)
	var i DataSourcesFunction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.DataSourceID,
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}
