// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/internal/ent/beacon"
	"github.com/kcarretto/realm/tavern/internal/ent/predicate"
	"github.com/kcarretto/realm/tavern/internal/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetKind sets the "kind" field.
func (tu *TagUpdate) SetKind(t tag.Kind) *TagUpdate {
	tu.mutation.SetKind(t)
	return tu
}

// AddBeaconIDs adds the "beacons" edge to the Beacon entity by IDs.
func (tu *TagUpdate) AddBeaconIDs(ids ...int) *TagUpdate {
	tu.mutation.AddBeaconIDs(ids...)
	return tu
}

// AddBeacons adds the "beacons" edges to the Beacon entity.
func (tu *TagUpdate) AddBeacons(b ...*Beacon) *TagUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.AddBeaconIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearBeacons clears all "beacons" edges to the Beacon entity.
func (tu *TagUpdate) ClearBeacons() *TagUpdate {
	tu.mutation.ClearBeacons()
	return tu
}

// RemoveBeaconIDs removes the "beacons" edge to Beacon entities by IDs.
func (tu *TagUpdate) RemoveBeaconIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveBeaconIDs(ids...)
	return tu
}

// RemoveBeacons removes "beacons" edges to Beacon entities.
func (tu *TagUpdate) RemoveBeacons(b ...*Beacon) *TagUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.RemoveBeaconIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TagMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Kind(); ok {
		if err := tag.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Tag.kind": %w`, err)}
		}
	}
	return nil
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Kind(); ok {
		_spec.SetField(tag.FieldKind, field.TypeEnum, value)
	}
	if tu.mutation.BeaconsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedBeaconsIDs(); len(nodes) > 0 && !tu.mutation.BeaconsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BeaconsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetKind sets the "kind" field.
func (tuo *TagUpdateOne) SetKind(t tag.Kind) *TagUpdateOne {
	tuo.mutation.SetKind(t)
	return tuo
}

// AddBeaconIDs adds the "beacons" edge to the Beacon entity by IDs.
func (tuo *TagUpdateOne) AddBeaconIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddBeaconIDs(ids...)
	return tuo
}

// AddBeacons adds the "beacons" edges to the Beacon entity.
func (tuo *TagUpdateOne) AddBeacons(b ...*Beacon) *TagUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.AddBeaconIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearBeacons clears all "beacons" edges to the Beacon entity.
func (tuo *TagUpdateOne) ClearBeacons() *TagUpdateOne {
	tuo.mutation.ClearBeacons()
	return tuo
}

// RemoveBeaconIDs removes the "beacons" edge to Beacon entities by IDs.
func (tuo *TagUpdateOne) RemoveBeaconIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveBeaconIDs(ids...)
	return tuo
}

// RemoveBeacons removes "beacons" edges to Beacon entities.
func (tuo *TagUpdateOne) RemoveBeacons(b ...*Beacon) *TagUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.RemoveBeaconIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (tuo *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks[*Tag, TagMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Kind(); ok {
		if err := tag.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Tag.kind": %w`, err)}
		}
	}
	return nil
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Kind(); ok {
		_spec.SetField(tag.FieldKind, field.TypeEnum, value)
	}
	if tuo.mutation.BeaconsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedBeaconsIDs(); len(nodes) > 0 && !tuo.mutation.BeaconsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BeaconsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BeaconsTable,
			Columns: tag.BeaconsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: beacon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}