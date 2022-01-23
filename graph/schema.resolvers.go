package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ControlServer/graph/generated"
	"ControlServer/graph/model"
	"ControlServer/internal/device"
	"context"
)

func (r *mutationResolver) RunCommand(ctx context.Context, input model.Command) (*model.CommandOutput, error) {
	return &model.CommandOutput{Output: "Noting"}, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	return device.GetDevices()
}

func (r *queryResolver) Processors(ctx context.Context, id string) ([]*model.ProcessorInfo, error) {
	return device.GetProcessors(id)
}

func (r *queryResolver) Bios(ctx context.Context, id string) ([]*model.BiosInfo, error) {
	return device.GetBiosInfo(id)
}

func (r *queryResolver) SysInfo(ctx context.Context, id string) ([]*model.SysInfo, error) {
	return device.GetSysInfo(id)
}

func (r *queryResolver) BaseBoards(ctx context.Context, id string) ([]*model.BaseBoardInfo, error) {
	return device.GetBaseBoards(id)
}

func (r *queryResolver) SysEnclosure(ctx context.Context, id string) ([]*model.SysEnclosureInfo, error) {
	return device.GetSysEnclosure(id)
}

func (r *queryResolver) SysSlots(ctx context.Context, id string) ([]*model.SysSlotInfo, error) {
	return device.GetSysSlots(id)
}

func (r *queryResolver) PhysMem(ctx context.Context, id string) ([]*model.PhysMemInfo, error) {
	return device.GetPhysMem(id)
}

func (r *queryResolver) Memory(ctx context.Context, id string) ([]*model.MemoryInfo, error) {
	return device.GetMemory(id)
}

func (r *queryResolver) OemStrings(ctx context.Context, id string) ([]*model.OemStringsInfo, error) {
	return device.GetOemStrings(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
