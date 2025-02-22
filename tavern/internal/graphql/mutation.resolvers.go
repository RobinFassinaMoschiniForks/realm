package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"strings"
	"time"

	"realm.pub/tavern/internal/auth"
	"realm.pub/tavern/internal/ent"
	"realm.pub/tavern/internal/ent/file"
	"realm.pub/tavern/internal/graphql/generated"
	"realm.pub/tavern/internal/graphql/models"
)

// DropAllData is the resolver for the dropAllData field.
func (r *mutationResolver) DropAllData(ctx context.Context) (bool, error) {
	// Initialize Transaction
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to initialize transaction: %w", err)
	}
	client := tx.Client()

	// Delete relevant ents
	if _, err := client.Beacon.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete beacons: %w", err))
	}
	if _, err := client.HostFile.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete hostfiles: %w", err))
	}
	if _, err := client.HostProcess.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete hostprocesses: %w", err))
	}
	if _, err := client.Host.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete hosts: %w", err))
	}
	if _, err := client.Quest.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete quests: %w", err))
	}
	if _, err := client.Tag.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete tags: %w", err))
	}
	if _, err := client.Task.Delete().Exec(ctx); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to delete tasks: %w", err))
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return false, rollback(tx, fmt.Errorf("failed to commit transaction: %w", err))
	}

	return true, nil
}

// CreateQuest is the resolver for the createQuest field.
func (r *mutationResolver) CreateQuest(ctx context.Context, beaconIDs []int, input ent.CreateQuestInput) (*ent.Quest, error) {
	// Ensure at least one Beacon ID provided
	if beaconIDs == nil || len(beaconIDs) < 1 {
		return nil, fmt.Errorf("must provide at least one beacon id")
	}

	// 1. Begin Transaction
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize transaction: %w", err)
	}
	client := tx.Client()

	// 2. Rollback transaction if we panic
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// 3. Load Tome
	questTome, err := client.Tome.Get(ctx, input.TomeID)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to load tome: %w", err))
	}

	// 4. Load Tome Files (ordered so that hashing is always the same)
	bundleFiles, err := questTome.QueryFiles().
		Order(ent.Asc(file.FieldID)).
		All(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to load tome files: %w", err))
	}

	// 5. Create bundle (if tome has files)
	var bundleID *int
	if len(bundleFiles) > 0 {
		bundle, err := createBundle(ctx, client, bundleFiles)
		if err != nil || bundle == nil {
			return nil, rollback(tx, fmt.Errorf("failed to create bundle: %w", err))
		}
		bundleID = &bundle.ID
	}

	// 6. Get creator from context (if available)
	var creatorID *int
	if creator := auth.UserFromContext(ctx); creator != nil {
		creatorID = &creator.ID
	}

	// 7. Create Quest
	quest, err := client.Quest.Create().
		SetInput(input).
		SetNillableBundleID(bundleID).
		SetEldritchAtCreation(questTome.Eldritch).
		SetParamDefsAtCreation(questTome.ParamDefs).
		SetTome(questTome).
		SetNillableCreatorID(creatorID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to create quest: %w", err))
	}

	// 8. Create tasks for each beacon
	for _, sid := range beaconIDs {
		_, err := client.Task.Create().
			SetQuest(quest).
			SetBeaconID(sid).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, fmt.Errorf("failed to create task for beacon (%q): %w", sid, err))
		}
	}

	// 9. Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to commit transaction: %w", err))
	}

	// 10. Load the quest with our non transactional client (cannot use transaction after commit)
	quest, err = r.client.Quest.Get(ctx, quest.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load created quest: %w", err)
	}

	return quest, nil
}

// UpdateBeacon is the resolver for the updateBeacon field.
func (r *mutationResolver) UpdateBeacon(ctx context.Context, beaconID int, input ent.UpdateBeaconInput) (*ent.Beacon, error) {
	return r.client.Beacon.UpdateOneID(beaconID).SetInput(input).Save(ctx)
}

// UpdateHost is the resolver for the updateHost field.
func (r *mutationResolver) UpdateHost(ctx context.Context, hostID int, input ent.UpdateHostInput) (*ent.Host, error) {
	return r.client.Host.UpdateOneID(hostID).SetInput(input).Save(ctx)
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error) {
	return r.client.Tag.Create().SetInput(input).Save(ctx)
}

// UpdateTag is the resolver for the updateTag field.
func (r *mutationResolver) UpdateTag(ctx context.Context, tagID int, input ent.UpdateTagInput) (*ent.Tag, error) {
	return r.client.Tag.UpdateOneID(tagID).SetInput(input).Save(ctx)
}

// CreateTome is the resolver for the createTome field.
func (r *mutationResolver) CreateTome(ctx context.Context, input ent.CreateTomeInput) (*ent.Tome, error) {
	var uploaderID *int
	if uploader := auth.UserFromContext(ctx); uploader != nil {
		uploaderID = &uploader.ID
	}

	return r.client.Tome.Create().
		SetNillableUploaderID(uploaderID).
		SetInput(input).
		Save(ctx)
}

// UpdateTome is the resolver for the updateTome field.
func (r *mutationResolver) UpdateTome(ctx context.Context, tomeID int, input ent.UpdateTomeInput) (*ent.Tome, error) {
	return r.client.Tome.UpdateOneID(tomeID).SetInput(input).Save(ctx)
}

// DeleteTome is the resolver for the deleteTome field.
func (r *mutationResolver) DeleteTome(ctx context.Context, tomeID int) (int, error) {
	if err := r.client.Tome.DeleteOneID(tomeID).Exec(ctx); err != nil {
		return 0, err
	}
	return tomeID, nil
}

// CreateRepository is the resolver for the createRepository field.
func (r *mutationResolver) CreateRepository(ctx context.Context, input ent.CreateRepositoryInput) (*ent.Repository, error) {
	var ownerID *int
	if owner := auth.UserFromContext(ctx); owner != nil {
		ownerID = &owner.ID
	}

	return r.client.Repository.Create().
		SetInput(input).
		SetNillableOwnerID(ownerID).
		Save(ctx)
}

// ImportRepository is the resolver for the importRepository field.
func (r *mutationResolver) ImportRepository(ctx context.Context, repoID int, input *models.ImportRepositoryInput) (*ent.Repository, error) {
	// Load Repository
	repo, err := r.client.Repository.Get(ctx, repoID)
	if err != nil {
		return nil, err
	}

	// Configure Filters
	filter := func(string) bool { return true }
	if input != nil && input.IncludeDirs != nil {
		filter = func(path string) bool {
			for _, prefix := range input.IncludeDirs {
				// Ignore Leading /
				path = strings.TrimPrefix(path, "/")
				prefix = strings.TrimPrefix(prefix, "/")

				// Include if matching
				if strings.HasPrefix(path, prefix) {
					return true
				}
			}
			return false
		}
	}

	// Import Tomes
	if err := r.importer.Import(ctx, repo, filter); err != nil {
		return nil, err
	}

	return repo.Update().SetLastImportedAt(time.Now()).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.client.User.UpdateOneID(userID).SetInput(input).Save(ctx)
}

// CreateCredential is the resolver for the createCredential field.
func (r *mutationResolver) CreateCredential(ctx context.Context, input ent.CreateHostCredentialInput) (*ent.HostCredential, error) {
	return r.client.HostCredential.Create().SetInput(input).Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
