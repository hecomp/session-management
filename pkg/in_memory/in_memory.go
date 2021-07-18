package in_memory

import (
	"sync"
	"time"
)

type MemStore interface {
	Commit(sessionId string, b []byte, expiration time.Time) error
	Delete(sessionId string) error
	Reset(sessionId string) error
	Find(sessionId string) ([]byte, bool, error)
}

type item struct {
	object     []byte
	expiration int64
}

// InMemStore represents the session in-memory store.
type InMemStore struct {
	items       map[string]item
	mu          sync.RWMutex
	stopCleanup chan bool
}

// NewInMemStore returns a new InMemStore instance, with a background cleanup goroutine that
// runs every minute to remove expired session data.
func NewInMemStore() *InMemStore {
	return NewWithCleanupInterval(time.Minute)
}

// NewWithCleanupInterval returns a new InMemStore instance. The cleanupInterval
// parameter controls how frequently expired session data is removed by the
// background cleanup goroutine. Setting it to 0 prevents the cleanup goroutine
// from running (i.e. expired sessions will not be removed).
func NewWithCleanupInterval(cleanupInterval time.Duration) *InMemStore {
	m := &InMemStore{
		items: make(map[string]item),
	}

	if cleanupInterval > 0 {
		go m.startCleanup(cleanupInterval)
	}

	return m
}

// Find returns the data for a given session sessionId from the InMemStore instance.
// If the session sessionId is not found or is expired, the returned exists flag will
// be set to false.
func (m *InMemStore) Find(sessionId string) ([]byte, bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, found := m.items[sessionId]
	if !found {
		return nil, false, nil
	}

	if time.Now().UnixNano() > item.expiration {
		return nil, false, nil
	}

	return item.object, true, nil
}

// Commit adds a session sessionId and data to the InMemStore instance with the given
// expiration time. If the session sessionId already exists, then the data and expiration
// time are updated.
func (m *InMemStore) Commit(sessionId string, b []byte, expiration time.Time) error {
	m.mu.Lock()
	m.items[sessionId] = item{
		object:     b,
		expiration: expiration.UnixNano(),
	}
	m.mu.Unlock()

	return nil
}

// Delete removes a session sessionId and corresponding data from the InMemStore
// instance.
func (m *InMemStore) Delete(sessionId string) error {
	m.mu.Lock()
	delete(m.items, sessionId)
	m.mu.Unlock()

	return nil
}

// Reset extend a session ttl from the InMemStore instance.
func (m *InMemStore) Reset(sessionId string, expiration time.Time) ([]byte, bool, error) {
	m.mu.RLock()
	item, found := m.items[sessionId]
	if !found {
		return nil, false, nil
	}

	if time.Now().UnixNano() > item.expiration {
		return nil, false, nil
	}
	m.mu.RUnlock()

	m.mu.Lock()
	item.expiration = expiration.UnixNano()
	m.items[sessionId] = item
	m.mu.Unlock()
		
	return item.object, true, nil
	
}

// startCleanup
func (m *InMemStore) startCleanup(interval time.Duration) {
	m.stopCleanup = make(chan bool)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			m.deleteExpired()
		case <-m.stopCleanup:
			ticker.Stop()
			return
		}
	}
}

// StopCleanup terminates the background cleanup goroutine for the InMemStore
// instance. It's rare to terminate this; generally InMemStore instances and
// their cleanup goroutines are intended to be long-lived and run for the lifetime
// of your application.
//
// There may be occasions though when your use of the InMemStore is transient.
// An example is creating a new InMemStore instance in a test function. In this
// scenario, the cleanup goroutine (which will run forever) will prevent the
// InMemStore object from being garbage collected even after the test function
// has finished. You can prevent this by manually calling StopCleanup.
func (m *InMemStore) StopCleanup() {
	if m.stopCleanup != nil {
		m.stopCleanup <- true
	}
}

func (m *InMemStore) deleteExpired() {
	now := time.Now().UnixNano()
	m.mu.Lock()
	for sessionId, item := range m.items {
		if now > item.expiration {
			delete(m.items, sessionId)
		}
	}
	m.mu.Unlock()
}