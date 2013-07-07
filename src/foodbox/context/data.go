package context

import (
	"appengine/datastore"
	"appengine"
)

// ------------------------
// The default query engine
// ------------------------

type QueryEngine interface {
	NewQuery(entityType string) RecordIterator
	NewEntity(entityType string, entity interface{}) error
	SaveEntity(entityType string, entityKey string, entity interface{}) error
	DeleteEntity(key string)
	GetEntity(key string, entity interface{})
}

type DatastoreQueryEngine struct {
	context appengine.Context
}

func NewQueryEngine(context appengine.Context) QueryEngine {
	return &DatastoreQueryEngine {
		context,
	}
}

func (engine *DatastoreQueryEngine) NewQuery(entityType string) RecordIterator {
	q := datastore.NewQuery(entityType)
	return &DatastoreRecordIterator {
		iterator: q.Run(engine.context),
	}
}

func (engine *DatastoreQueryEngine) NewEntity(entityType string, entity interface{}) error {
	_, err := datastore.Put(
		engine.context,
		datastore.NewIncompleteKey(engine.context, entityType, nil),
		entity)
	return err
}

func (engine *DatastoreQueryEngine) SaveEntity(entityType string, entityKey string, entity interface{}) error {
	key, err := datastore.DecodeKey(entityKey)
	if err != nil {
		return err
	}
	
	_, err = datastore.Put(
		engine.context,
		key,
		entity)
	return err
}

func (engine *DatastoreQueryEngine) DeleteEntity(key string) {
	appengineKey, _ := datastore.DecodeKey(key)
	datastore.Delete(engine.context, appengineKey)
}

func (engine *DatastoreQueryEngine) GetEntity(key string, entity interface{}) {
	appengineKey, _ := datastore.DecodeKey(key)
	datastore.Get(engine.context, appengineKey, entity)
}

// ----------------------------
// The default record iterator
// ----------------------------

type RecordIterator interface {
	Next(dst interface{}) (string, bool)
}

type DatastoreRecordIterator struct {
	iterator *datastore.Iterator
}

func (iterator *DatastoreRecordIterator) Next(dst interface{}) (key string, more bool) {
	appengineKey, err := iterator.iterator.Next(dst)
	if err != nil {
		more = false
		key = ""
		return
	}
	
	more = true
	key = appengineKey.Encode()
	return
}