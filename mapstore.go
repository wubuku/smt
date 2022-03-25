package smt

import (
	"fmt"
)

// MapStore is a key-value store.
type MapStore interface {
	Get(key []byte) ([]byte, error)     // Get gets the value for a key.
	Set(key []byte, value []byte) error // Set updates the value for a key.
	Delete(key []byte) error            // Delete deletes a key.
}

// InvalidKeyError is thrown when a key that does not exist is being accessed.
type InvalidKeyError struct {
	Key []byte
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("invalid key: %x", e.Key)
}

// SimpleMap is a simple in-memory map.
type SimpleMap struct {
	m map[string][]byte
}

// NewSimpleMap creates a new empty SimpleMap.
func NewSimpleMap() *SimpleMap {
	return &SimpleMap{
		m: make(map[string][]byte),
	}
}

// Get gets the value for a key.
func (sm *SimpleMap) Get(key []byte) ([]byte, error) {
	if value, ok := sm.m[string(key)]; ok {
		return value, nil
	}
	return nil, &InvalidKeyError{Key: key}
}

// Set updates the value for a key.
func (sm *SimpleMap) Set(key []byte, value []byte) error {
	sm.m[string(key)] = value
	return nil
}

// Delete deletes a key.
func (sm *SimpleMap) Delete(key []byte) error {
	_, ok := sm.m[string(key)]
	if ok {
		delete(sm.m, string(key))
		return nil
	}
	return &InvalidKeyError{Key: key}
}

type SmtValueStore interface {
	MapStore
	Immutable() bool
	GetForRoot(key []byte, smtRoot []byte) ([]byte, error)     // Get gets the value for a key.
	SetForRoot(key []byte, smtRoot []byte, value []byte) error // Set updates the value for a key.
}

func NewSmtValueStore(mapStore MapStore) SmtValueStore {
	svm, b := mapStore.(SmtValueStore)
	if b {
		return svm
	} else {
		return &MapStoreSmtValueStoreWrapper{
			mapStore: mapStore,
		}
	}
}

type MapStoreSmtValueStoreWrapper struct {
	mapStore MapStore
}

func (sm *MapStoreSmtValueStoreWrapper) Get(key []byte) ([]byte, error) {
	return sm.mapStore.Get(key)
}

// Set updates the value for a key.
func (sm *MapStoreSmtValueStoreWrapper) Set(key []byte, value []byte) error {
	return sm.mapStore.Set(key, value)
}

// Delete deletes a key.
func (sm *MapStoreSmtValueStoreWrapper) Delete(key []byte) error {
	return sm.mapStore.Delete(key)
}

func (sm *MapStoreSmtValueStoreWrapper) Immutable() bool {
	return false
}

func (sm *MapStoreSmtValueStoreWrapper) GetForRoot(key []byte, smtRoot []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented GetForRoot")
}

func (sm *MapStoreSmtValueStoreWrapper) SetForRoot(key []byte, smtRoot []byte, value []byte) error {
	return fmt.Errorf("not implemented SetForRoot")
}

type SimpleSmtValueStore struct {
	SimpleMap
}

func NewSimpleSmtValueMap() *SimpleSmtValueStore {
	svm := SimpleSmtValueStore{}
	svm.m = make(map[string][]byte)
	return &svm
}

func (sm *SimpleSmtValueStore) Immutable() bool {
	return false
}

func (sm *SimpleSmtValueStore) GetForRoot(key []byte, smtRoot []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented GetForRoot")
}

func (sm *SimpleSmtValueStore) SetForRoot(key []byte, smtRoot []byte, value []byte) error {
	return fmt.Errorf("not implemented SetForRoot")
}
