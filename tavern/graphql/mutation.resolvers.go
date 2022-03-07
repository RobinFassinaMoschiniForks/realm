package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kcarretto/realm/tavern/auth"
	"github.com/kcarretto/realm/tavern/ent"
	"github.com/kcarretto/realm/tavern/ent/implant"
	"github.com/kcarretto/realm/tavern/ent/implantconfig"
)

func (r *mutationResolver) Callback(ctx context.Context, info CallbackInput) (*CallbackResponse, error) {
	// Get the viewer
	vc, err := auth.ImplantFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Get the implant config based on the viewer
	configID, err := r.client.ImplantConfig.Query().
		Where(
			implantconfig.AuthToken(vc.AuthToken),
		).
		OnlyID(ctx)
	if err != nil {
		return nil, err
	}

	// Load the target
	target, err := r.client.Target.Get(ctx, info.TargetID)
	if err != nil {
		return nil, err
	}

	// Upsert the implant
	impQuery := r.client.Implant.Query().
		Where(implant.SessionID(info.SessionID))

	var imp *ent.Implant
	if exists := impQuery.Clone().ExistX(ctx); exists {
		impID := impQuery.OnlyIDX(ctx)
		imp = r.client.Implant.UpdateOneID(impID).
			SetProcessName(info.ProcessName).
			SaveX(ctx)
	} else {
		imp = r.client.Implant.Create().
			SetSessionID(info.SessionID).
			SetProcessName(info.ProcessName).
			SetConfigID(configID).
			SetTarget(target).
			SaveX(ctx)
	}

	// Format the response
	resp := CallbackResponse{
		Implant: imp,
	}

	return &resp, nil
}

func (r *mutationResolver) CreateImplantCallbackConfig(ctx context.Context, config CreateImplantCallbackConfigInput) (*ent.ImplantCallbackConfig, error) {
	return r.client.ImplantCallbackConfig.Create().
		SetURI(config.URI).
		SetNillableProxyURI(config.ProxyURI).
		SetNillablePriority(config.Priority).
		SetNillableTimeout(config.Timeout).
		SetNillableInterval(config.Interval).
		SetNillableJitter(config.Jitter).
		Save(ctx)
}

func (r *mutationResolver) UpdateImplantCallbackConfig(ctx context.Context, config UpdateImplantCallbackConfigInput) (*ent.ImplantCallbackConfig, error) {
	cfg, err := r.client.ImplantCallbackConfig.Get(ctx, config.ID)
	if err != nil {
		return nil, err
	}

	mutation := cfg.Update().
		SetNillableProxyURI(config.ProxyURI).
		SetNillablePriority(config.Priority).
		SetNillableTimeout(config.Timeout).
		SetNillableInterval(config.Interval).
		SetNillableJitter(config.Jitter)
	if config.URI != nil {
		mutation = mutation.SetURI(*config.URI)
	}
	return mutation.Save(ctx)
}

func (r *mutationResolver) DeleteImplantCallbackConfig(ctx context.Context, id int) (int, error) {
	return id, r.client.ImplantCallbackConfig.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) CreateImplantServiceConfig(ctx context.Context, config CreateImplantServiceConfigInput) (*ent.ImplantServiceConfig, error) {
	return r.client.ImplantServiceConfig.Create().
		SetName(config.Name).
		SetNillableDescription(config.Description).
		SetExecutablePath(config.ExecutablePath).
		Save(ctx)
}

func (r *mutationResolver) UpdateImplantServiceConfig(ctx context.Context, config UpdateImplantServiceConfigInput) (*ent.ImplantServiceConfig, error) {
	cfg, err := r.client.ImplantServiceConfig.Get(ctx, config.ID)
	if err != nil {
		return nil, err
	}

	mutation := cfg.Update().
		SetNillableDescription(config.Description)
	if config.Name != nil {
		mutation = mutation.SetName(*config.Name)
	}
	if config.ExecutablePath != nil {
		mutation = mutation.SetExecutablePath(*config.ExecutablePath)
	}
	return mutation.Save(ctx)
}

func (r *mutationResolver) DeleteImplantServiceConfig(ctx context.Context, id int) (int, error) {
	return id, r.client.ImplantServiceConfig.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) CreateImplantConfig(ctx context.Context, config CreateImplantConfigInput) (*ent.ImplantConfig, error) {
	return r.client.ImplantConfig.Create().
		SetName(config.Name).
		AddCallbackConfigIDs(config.CallbackConfigIDs...).
		AddServiceConfigIDs(config.ServiceConfigIDs...).
		Save(ctx)
}

func (r *mutationResolver) UpdateImplantConfig(ctx context.Context, config UpdateImplantConfigInput) (*ent.ImplantConfig, error) {
	mutation := r.client.ImplantConfig.UpdateOneID(config.ID).
		AddCallbackConfigIDs(config.AddCallbackConfigIDs...).
		RemoveCallbackConfigIDs(config.RemoveCallbackConfigIDs...).
		AddServiceConfigIDs(config.AddServiceConfigIDs...).
		RemoveServiceConfigIDs(config.RemoveServiceConfigIDs...)
	if config.Name != nil {
		mutation.SetName(*config.Name)
	}
	return mutation.Save(ctx)
}

func (r *mutationResolver) DeleteImplantConfig(ctx context.Context, id int) (int, error) {
	return id, r.client.ImplantConfig.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) CreateDeploymentConfig(ctx context.Context, config CreateDeploymentConfigInput) (*ent.DeploymentConfig, error) {
	return r.client.DeploymentConfig.Create().
		SetName(config.Name).
		SetNillableCmd(config.Cmd).
		SetNillableStartCmd(config.StartCmd).
		SetNillableFileDst(config.FileDst).
		SetNillableImplantConfigID(config.ImplantConfigID).
		SetNillableFileID(config.FileID).
		Save(ctx)
}

func (r *mutationResolver) UpdateDeploymentConfig(ctx context.Context, config UpdateDeploymentConfigInput) (*ent.DeploymentConfig, error) {
	mutation := r.client.DeploymentConfig.UpdateOneID(config.ID).
		SetNillableCmd(config.Cmd).
		SetNillableStartCmd(config.StartCmd).
		SetNillableFileDst(config.FileDst).
		SetNillableImplantConfigID(config.ImplantConfigID).
		SetNillableFileID(config.FileID)
	if config.Name != nil {
		mutation = mutation.SetName(*config.Name)
	}
	return mutation.Save(ctx)
}

func (r *mutationResolver) DeleteDeploymentConfig(ctx context.Context, id int) (int, error) {
	return id, r.client.DeploymentConfig.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) CreateTag(ctx context.Context, tag CreateTagInput) (*ent.Tag, error) {
	return r.client.Tag.Create().
		SetName(tag.Name).
		AddTargetIDs(tag.TargetIDs...).
		Save(ctx)
}

func (r *mutationResolver) UpdateTag(ctx context.Context, tag UpdateTagInput) (*ent.Tag, error) {
	mutation := r.client.Tag.UpdateOneID(tag.ID).
		AddTargetIDs(tag.AddTargetIDs...).
		RemoveTargetIDs(tag.RemoveTargetIDs...)
	if tag.Name != nil {
		mutation = mutation.SetName(*tag.Name)
	}
	return mutation.Save(ctx)
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id int) (int, error) {
	return id, r.client.Tag.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) CreateTarget(ctx context.Context, target CreateTargetInput) (*ent.Target, error) {
	return r.client.Target.Create().
		SetName(target.Name).
		SetForwardConnectIP(target.ForwardConnectIP).
		Save(ctx)
}

func (r *mutationResolver) CreateCredential(ctx context.Context, credential CreateCredentialInput) (*ent.Credential, error) {
	return r.client.Credential.Create().
		SetTargetID(credential.TargetID).
		SetPrincipal(credential.Principal).
		SetSecret(credential.Secret).
		SetKind(credential.Kind).
		Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
