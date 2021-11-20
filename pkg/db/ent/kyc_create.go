// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
)

// KycCreate is the builder for creating a Kyc entity.
type KycCreate struct {
	config
	mutation *KycMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (kc *KycCreate) SetUserID(u uuid.UUID) *KycCreate {
	kc.mutation.SetUserID(u)
	return kc
}

// SetAppID sets the "app_id" field.
func (kc *KycCreate) SetAppID(u uuid.UUID) *KycCreate {
	kc.mutation.SetAppID(u)
	return kc
}

// SetFirstName sets the "first_name" field.
func (kc *KycCreate) SetFirstName(s string) *KycCreate {
	kc.mutation.SetFirstName(s)
	return kc
}

// SetLastName sets the "last_name" field.
func (kc *KycCreate) SetLastName(s string) *KycCreate {
	kc.mutation.SetLastName(s)
	return kc
}

// SetRegion sets the "region" field.
func (kc *KycCreate) SetRegion(s string) *KycCreate {
	kc.mutation.SetRegion(s)
	return kc
}

// SetCardType sets the "card_type" field.
func (kc *KycCreate) SetCardType(s string) *KycCreate {
	kc.mutation.SetCardType(s)
	return kc
}

// SetCardID sets the "card_id" field.
func (kc *KycCreate) SetCardID(s string) *KycCreate {
	kc.mutation.SetCardID(s)
	return kc
}

// SetFrontCardImg sets the "front_card_img" field.
func (kc *KycCreate) SetFrontCardImg(s string) *KycCreate {
	kc.mutation.SetFrontCardImg(s)
	return kc
}

// SetBackCardImg sets the "back_card_img" field.
func (kc *KycCreate) SetBackCardImg(s string) *KycCreate {
	kc.mutation.SetBackCardImg(s)
	return kc
}

// SetUserHandlingCardImg sets the "user_handling_card_img" field.
func (kc *KycCreate) SetUserHandlingCardImg(s string) *KycCreate {
	kc.mutation.SetUserHandlingCardImg(s)
	return kc
}

// SetReviewStatus sets the "review_status" field.
func (kc *KycCreate) SetReviewStatus(s string) *KycCreate {
	kc.mutation.SetReviewStatus(s)
	return kc
}

// SetCreateAt sets the "create_at" field.
func (kc *KycCreate) SetCreateAt(u uint32) *KycCreate {
	kc.mutation.SetCreateAt(u)
	return kc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (kc *KycCreate) SetNillableCreateAt(u *uint32) *KycCreate {
	if u != nil {
		kc.SetCreateAt(*u)
	}
	return kc
}

// SetUpdateAt sets the "update_at" field.
func (kc *KycCreate) SetUpdateAt(u uint32) *KycCreate {
	kc.mutation.SetUpdateAt(u)
	return kc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (kc *KycCreate) SetNillableUpdateAt(u *uint32) *KycCreate {
	if u != nil {
		kc.SetUpdateAt(*u)
	}
	return kc
}

// SetID sets the "id" field.
func (kc *KycCreate) SetID(u uuid.UUID) *KycCreate {
	kc.mutation.SetID(u)
	return kc
}

// Mutation returns the KycMutation object of the builder.
func (kc *KycCreate) Mutation() *KycMutation {
	return kc.mutation
}

// Save creates the Kyc in the database.
func (kc *KycCreate) Save(ctx context.Context) (*Kyc, error) {
	var (
		err  error
		node *Kyc
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KycCreate) SaveX(ctx context.Context) *Kyc {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KycCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KycCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KycCreate) defaults() {
	if _, ok := kc.mutation.CreateAt(); !ok {
		v := kyc.DefaultCreateAt()
		kc.mutation.SetCreateAt(v)
	}
	if _, ok := kc.mutation.UpdateAt(); !ok {
		v := kyc.DefaultUpdateAt()
		kc.mutation.SetUpdateAt(v)
	}
	if _, ok := kc.mutation.ID(); !ok {
		v := kyc.DefaultID()
		kc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KycCreate) check() error {
	if _, ok := kc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := kc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := kc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "first_name"`)}
	}
	if _, ok := kc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "last_name"`)}
	}
	if _, ok := kc.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`ent: missing required field "region"`)}
	}
	if _, ok := kc.mutation.CardType(); !ok {
		return &ValidationError{Name: "card_type", err: errors.New(`ent: missing required field "card_type"`)}
	}
	if _, ok := kc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card_id", err: errors.New(`ent: missing required field "card_id"`)}
	}
	if _, ok := kc.mutation.FrontCardImg(); !ok {
		return &ValidationError{Name: "front_card_img", err: errors.New(`ent: missing required field "front_card_img"`)}
	}
	if _, ok := kc.mutation.BackCardImg(); !ok {
		return &ValidationError{Name: "back_card_img", err: errors.New(`ent: missing required field "back_card_img"`)}
	}
	if _, ok := kc.mutation.UserHandlingCardImg(); !ok {
		return &ValidationError{Name: "user_handling_card_img", err: errors.New(`ent: missing required field "user_handling_card_img"`)}
	}
	if _, ok := kc.mutation.ReviewStatus(); !ok {
		return &ValidationError{Name: "review_status", err: errors.New(`ent: missing required field "review_status"`)}
	}
	if _, ok := kc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := kc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	return nil
}

func (kc *KycCreate) sqlSave(ctx context.Context) (*Kyc, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (kc *KycCreate) createSpec() (*Kyc, *sqlgraph.CreateSpec) {
	var (
		_node = &Kyc{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kyc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		}
	)
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := kc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := kc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := kc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := kc.mutation.Region(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldRegion,
		})
		_node.Region = value
	}
	if value, ok := kc.mutation.CardType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardType,
		})
		_node.CardType = value
	}
	if value, ok := kc.mutation.CardID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardID,
		})
		_node.CardID = value
	}
	if value, ok := kc.mutation.FrontCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontCardImg,
		})
		_node.FrontCardImg = value
	}
	if value, ok := kc.mutation.BackCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackCardImg,
		})
		_node.BackCardImg = value
	}
	if value, ok := kc.mutation.UserHandlingCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldUserHandlingCardImg,
		})
		_node.UserHandlingCardImg = value
	}
	if value, ok := kc.mutation.ReviewStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldReviewStatus,
		})
		_node.ReviewStatus = value
	}
	if value, ok := kc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := kc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	return _node, _spec
}

// KycCreateBulk is the builder for creating many Kyc entities in bulk.
type KycCreateBulk struct {
	config
	builders []*KycCreate
}

// Save creates the Kyc entities in the database.
func (kcb *KycCreateBulk) Save(ctx context.Context) ([]*Kyc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kyc, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KycMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KycCreateBulk) SaveX(ctx context.Context) []*Kyc {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KycCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KycCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}