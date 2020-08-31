// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/peanut-cc/ent_orm_notes/user_index2/ent/city"
	"github.com/peanut-cc/ent_orm_notes/user_index2/ent/street"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCity   = "City"
	TypeStreet = "Street"
)

// CityMutation represents an operation that mutate the Cities
// nodes in the graph.
type CityMutation struct {
	config
	op             Op
	typ            string
	id             *int
	name           *string
	clearedFields  map[string]struct{}
	streets        map[int]struct{}
	removedstreets map[int]struct{}
	done           bool
	oldValue       func(context.Context) (*City, error)
}

var _ ent.Mutation = (*CityMutation)(nil)

// cityOption allows to manage the mutation configuration using functional options.
type cityOption func(*CityMutation)

// newCityMutation creates new mutation for $n.Name.
func newCityMutation(c config, op Op, opts ...cityOption) *CityMutation {
	m := &CityMutation{
		config:        c,
		op:            op,
		typ:           TypeCity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCityID sets the id field of the mutation.
func withCityID(id int) cityOption {
	return func(m *CityMutation) {
		var (
			err   error
			once  sync.Once
			value *City
		)
		m.oldValue = func(ctx context.Context) (*City, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().City.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCity sets the old City of the mutation.
func withCity(node *City) cityOption {
	return func(m *CityMutation) {
		m.oldValue = func(context.Context) (*City, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CityMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *CityMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *CityMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the City.
// If the City object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *CityMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *CityMutation) ResetName() {
	m.name = nil
}

// AddStreetIDs adds the streets edge to Street by ids.
func (m *CityMutation) AddStreetIDs(ids ...int) {
	if m.streets == nil {
		m.streets = make(map[int]struct{})
	}
	for i := range ids {
		m.streets[ids[i]] = struct{}{}
	}
}

// RemoveStreetIDs removes the streets edge to Street by ids.
func (m *CityMutation) RemoveStreetIDs(ids ...int) {
	if m.removedstreets == nil {
		m.removedstreets = make(map[int]struct{})
	}
	for i := range ids {
		m.removedstreets[ids[i]] = struct{}{}
	}
}

// RemovedStreets returns the removed ids of streets.
func (m *CityMutation) RemovedStreetsIDs() (ids []int) {
	for id := range m.removedstreets {
		ids = append(ids, id)
	}
	return
}

// StreetsIDs returns the streets ids in the mutation.
func (m *CityMutation) StreetsIDs() (ids []int) {
	for id := range m.streets {
		ids = append(ids, id)
	}
	return
}

// ResetStreets reset all changes of the "streets" edge.
func (m *CityMutation) ResetStreets() {
	m.streets = nil
	m.removedstreets = nil
}

// Op returns the operation name.
func (m *CityMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (City).
func (m *CityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CityMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, city.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CityMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case city.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *CityMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case city.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown City field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CityMutation) SetField(name string, value ent.Value) error {
	switch name {
	case city.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CityMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CityMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CityMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown City numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CityMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CityMutation) ClearField(name string) error {
	return fmt.Errorf("unknown City nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CityMutation) ResetField(name string) error {
	switch name {
	case city.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CityMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.streets != nil {
		edges = append(edges, city.EdgeStreets)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CityMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case city.EdgeStreets:
		ids := make([]ent.Value, 0, len(m.streets))
		for id := range m.streets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedstreets != nil {
		edges = append(edges, city.EdgeStreets)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CityMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case city.EdgeStreets:
		ids := make([]ent.Value, 0, len(m.removedstreets))
		for id := range m.removedstreets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CityMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CityMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown City unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CityMutation) ResetEdge(name string) error {
	switch name {
	case city.EdgeStreets:
		m.ResetStreets()
		return nil
	}
	return fmt.Errorf("unknown City edge %s", name)
}

// StreetMutation represents an operation that mutate the Streets
// nodes in the graph.
type StreetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	city          *int
	clearedcity   bool
	done          bool
	oldValue      func(context.Context) (*Street, error)
}

var _ ent.Mutation = (*StreetMutation)(nil)

// streetOption allows to manage the mutation configuration using functional options.
type streetOption func(*StreetMutation)

// newStreetMutation creates new mutation for $n.Name.
func newStreetMutation(c config, op Op, opts ...streetOption) *StreetMutation {
	m := &StreetMutation{
		config:        c,
		op:            op,
		typ:           TypeStreet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStreetID sets the id field of the mutation.
func withStreetID(id int) streetOption {
	return func(m *StreetMutation) {
		var (
			err   error
			once  sync.Once
			value *Street
		)
		m.oldValue = func(ctx context.Context) (*Street, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Street.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStreet sets the old Street of the mutation.
func withStreet(node *Street) streetOption {
	return func(m *StreetMutation) {
		m.oldValue = func(context.Context) (*Street, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StreetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StreetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *StreetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *StreetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *StreetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Street.
// If the Street object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *StreetMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *StreetMutation) ResetName() {
	m.name = nil
}

// SetCityID sets the city edge to City by id.
func (m *StreetMutation) SetCityID(id int) {
	m.city = &id
}

// ClearCity clears the city edge to City.
func (m *StreetMutation) ClearCity() {
	m.clearedcity = true
}

// CityCleared returns if the edge city was cleared.
func (m *StreetMutation) CityCleared() bool {
	return m.clearedcity
}

// CityID returns the city id in the mutation.
func (m *StreetMutation) CityID() (id int, exists bool) {
	if m.city != nil {
		return *m.city, true
	}
	return
}

// CityIDs returns the city ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// CityID instead. It exists only for internal usage by the builders.
func (m *StreetMutation) CityIDs() (ids []int) {
	if id := m.city; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCity reset all changes of the "city" edge.
func (m *StreetMutation) ResetCity() {
	m.city = nil
	m.clearedcity = false
}

// Op returns the operation name.
func (m *StreetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Street).
func (m *StreetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *StreetMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, street.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *StreetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case street.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *StreetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case street.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Street field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *StreetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case street.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Street field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *StreetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *StreetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *StreetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Street numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *StreetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *StreetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *StreetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Street nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *StreetMutation) ResetField(name string) error {
	switch name {
	case street.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Street field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *StreetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.city != nil {
		edges = append(edges, street.EdgeCity)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *StreetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case street.EdgeCity:
		if id := m.city; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *StreetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *StreetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *StreetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcity {
		edges = append(edges, street.EdgeCity)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *StreetMutation) EdgeCleared(name string) bool {
	switch name {
	case street.EdgeCity:
		return m.clearedcity
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *StreetMutation) ClearEdge(name string) error {
	switch name {
	case street.EdgeCity:
		m.ClearCity()
		return nil
	}
	return fmt.Errorf("unknown Street unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *StreetMutation) ResetEdge(name string) error {
	switch name {
	case street.EdgeCity:
		m.ResetCity()
		return nil
	}
	return fmt.Errorf("unknown Street edge %s", name)
}