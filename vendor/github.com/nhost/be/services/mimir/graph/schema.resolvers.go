package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	"github.com/nhost/be/services/mimir/graph/generated"
	"github.com/nhost/be/services/mimir/model"
)

// UpdateConfig is the resolver for the setConfig field.
func (r *mutationResolver) UpdateConfig(ctx context.Context, appID string, config model.ConfigConfigUpdateInput) (*model.ConfigConfig, error) {
	return r.updateConfig(ctx, appID, config)
}

// ReplaceConfig is the resolver for the replaceConfig field.
func (r *mutationResolver) ReplaceConfig(ctx context.Context, appID string, config model.ConfigConfigInsertInput) (*model.ConfigConfig, error) {
	return r.replaceConfig(ctx, appID, config)
}

// InsertConfig is the resolver for the insertConfig field.
func (r *mutationResolver) InsertConfig(ctx context.Context, appID string, config model.ConfigConfigInsertInput, systemConfig model.ConfigSystemConfigInsertInput, secrets []*model.ConfigEnvironmentVariableInsertInput) (*model.ConfigInsertConfigResponse, error) {
	return r.insertConfig(ctx, appID, config, systemConfig, secrets)
}

// DeleteConfig is the resolver for the deleteConfig field.
func (r *mutationResolver) DeleteConfig(ctx context.Context, appID string) (*model.ConfigConfig, error) {
	return r.deleteConfig(ctx, appID)
}

// InsertSecret is the resolver for the insertSecret field.
func (r *mutationResolver) InsertSecret(ctx context.Context, appID string, secret model.ConfigEnvironmentVariableInsertInput) (*model.ConfigEnvironmentVariable, error) {
	return r.insertSecret(ctx, appID, secret)
}

// UpdateSecret is the resolver for the updateSecret field.
func (r *mutationResolver) UpdateSecret(ctx context.Context, appID string, secret model.ConfigEnvironmentVariableInsertInput) (*model.ConfigEnvironmentVariable, error) {
	return r.updateSecret(ctx, appID, secret)
}

// DeleteSecret is the resolver for the deleteSecret field.
func (r *mutationResolver) DeleteSecret(ctx context.Context, appID string, key string) (*model.ConfigEnvironmentVariable, error) {
	return r.deleteSecret(ctx, appID, key)
}

// UpdateSystemConfig is the resolver for the updateSystemConfig field.
func (r *mutationResolver) UpdateSystemConfig(ctx context.Context, appID string, systemConfig model.ConfigSystemConfigUpdateInput) (*model.ConfigSystemConfig, error) {
	return r.updateSystemConfig(ctx, appID, systemConfig)
}

// InsertRunServiceConfig is the resolver for the insertRunServiceConfig field.
func (r *mutationResolver) InsertRunServiceConfig(ctx context.Context, appID string, serviceID string, config model.ConfigRunServiceConfigInsertInput) (*model.ConfigRunServiceConfig, error) {
	return r.insertRunServiceConfig(ctx, appID, serviceID, config)
}

// UpdateRunServiceConfig is the resolver for the updateRunServiceConfig field.
func (r *mutationResolver) UpdateRunServiceConfig(ctx context.Context, appID string, serviceID string, config model.ConfigRunServiceConfigUpdateInput) (*model.ConfigRunServiceConfig, error) {
	return r.updateRunServiceConfig(ctx, appID, serviceID, config)
}

// ReplaceRunServiceConfig is the resolver for the replaceRunServiceConfig field.
func (r *mutationResolver) ReplaceRunServiceConfig(ctx context.Context, appID string, serviceID string, config model.ConfigRunServiceConfigInsertInput) (*model.ConfigRunServiceConfig, error) {
	return r.replaceRunServiceConfig(ctx, appID, serviceID, config)
}

// DeleteRunServiceConfig is the resolver for the deleteRunServiceConfig field.
func (r *mutationResolver) DeleteRunServiceConfig(ctx context.Context, appID string, serviceID string) (*model.ConfigRunServiceConfig, error) {
	return r.deleteRunServiceConfig(ctx, appID, serviceID)
}

// ConfigRawJSON is the resolver for the configRawJSON field.
func (r *queryResolver) ConfigRawJSON(ctx context.Context, appID string, resolve bool) (string, error) {
	return r.configRawJSON(ctx, appID, resolve)
}

// Config is the resolver for the config field.
func (r *queryResolver) Config(ctx context.Context, appID string, resolve bool) (*model.ConfigConfig, error) {
	return r.config(ctx, appID, resolve)
}

// Configs is the resolver for the configs field.
func (r *queryResolver) Configs(ctx context.Context, resolve bool, where *model.ConfigConfigComparisonExp) ([]*model.ConfigAppConfig, error) {
	return r.configs(ctx, resolve, where)
}

// AppSecrets is the resolver for the appSecrets field.
func (r *queryResolver) AppSecrets(ctx context.Context, appID string) ([]*model.ConfigEnvironmentVariable, error) {
	return r.appSecrets(ctx, appID)
}

// AppsSecrets is the resolver for the appsSecrets field.
func (r *queryResolver) AppsSecrets(ctx context.Context) ([]*model.ConfigAppSecrets, error) {
	return r.appsSecrets(ctx), nil
}

// SystemConfig is the resolver for the systemConfig field.
func (r *queryResolver) SystemConfig(ctx context.Context, appID string) (*model.ConfigSystemConfig, error) {
	return r.systemConfig(ctx, appID)
}

// SystemConfigs is the resolver for the systemConfigs field.
func (r *queryResolver) SystemConfigs(ctx context.Context, where *model.ConfigSystemConfigComparisonExp) ([]*model.ConfigAppSystemConfig, error) {
	return r.systemConfigs(ctx, where)
}

// RunServiceConfigRawJSON is the resolver for the runServiceConfigRawJSON field.
func (r *queryResolver) RunServiceConfigRawJSON(ctx context.Context, appID string, serviceID string, resolve bool) (string, error) {
	return r.runServiceConfigRawJSON(ctx, appID, serviceID, resolve)
}

// RunServiceConfig is the resolver for the runServiceConfig field.
func (r *queryResolver) RunServiceConfig(ctx context.Context, appID string, serviceID string, resolve bool) (*model.ConfigRunServiceConfig, error) {
	return r.runServiceConfig(ctx, appID, serviceID, resolve)
}

// RunServiceConfigs is the resolver for the runServiceConfigs field.
func (r *queryResolver) RunServiceConfigs(ctx context.Context, appID string, resolve bool) ([]*model.ConfigRunServiceConfigWithID, error) {
	return r.runServiceConfigs(ctx, &appID, resolve, nil)
}

// RunServiceConfigsAll is the resolver for the runServiceConfigsAll field.
func (r *queryResolver) RunServiceConfigsAll(ctx context.Context, resolve bool, where *model.ConfigRunServiceConfigComparisonExp) ([]*model.ConfigRunServiceConfigWithID, error) {
	return r.runServiceConfigs(ctx, nil, resolve, where)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
