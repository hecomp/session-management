package in_memory

import (
	"sync"
	"time"

	"github.com/go-kit/kit/log"

	. "github.com/hecomp/session-management/internal/models"
)

// SessionInterval parameter controls how frequently expired session data is removed by the
// background cleanup goroutine
const SessionInterval = 2 * time.Minute

// MemStore
type MemStore interface {
	Commit(sessionId string, b []byte, expiration time.Time) error
	Delete(sessionId string) error
	Reset(sessionId string) error
	Find(sessionId string) ([]byte, bool, error)
	List() map[string]Item
}

// InMemStore represents the session in-memory store.
type InMemStore struct {
	logger      log.Logger
	items       map[string]Item
	mu          sync.RWMutex
	stopCleanup chan bool
}

// NewInMemStore returns a new InMemStore instance, with a background session cleanup goroutine that
// runs every minute to remove expired session data.
func NewInMemStore(logger log.Logger) *InMemStore {
	m := &InMemStore{
		items: make(map[string]Item),
		logger: logger,
	}

	go m.startSessionCleanup(SessionInterval)

	return m
}

// NewWithCleanupInterval returns a new InMemStore instance. The cleanupInterval
// parameter controls how frequently expired session data is removed by the
// background cleanup goroutine. Setting it to 0 prevents the cleanup goroutine
// from running (i.e. expired sessions will not be removed).
//func NewWithCleanupInterval(cleanupInterval time.Duration) *InMemStore {
//	m := &InMemStore{
//		items: make(map[string]Item),
//	}
//
//	if cleanupInterval > 0 {
//		go m.startSessionCleanup(cleanupInterval)
//	}
//
//	return m
//}

// Find returns the sessionId for a given session from the InMemStore instance.
// If the session id is not found or is expired, the returned exists flag will
// be set to false.
func (m *InMemStore) Find(sessionId string) ([]byte, bool, error) {
	m.logger.Log("method", "find", "sessionId", sessionId)
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, found := m.items[sessionId]
	if !found {
		return nil, false, nil
	}

	if time.Now().UnixNano() > item.Expiration {
		m.logger.Log("action", "expired", "sessionId", sessionId)
		return nil, false, nil
	}

	return item.Oject, true, nil
}

// Commit adds a session sessionId and data to the InMemStore instance with the given
// expiration time. If the session sessionId already exists, then the data and expiration
// time are updated.
func (m *InMemStore) Commit(sessionId string, b []byte, expiration time.Time) error {
	m.logger.Log("method", "commit", "sessionId", sessionId)
	m.mu.Lock()
	m.items[sessionId] = Item{
		Oject:     b,
		Expiration: expiration.UnixNano(),
	}
	m.mu.Unlock()

	return nil
}

// Delete removes a session sessionId and corresponding data from the InMemStore
// instance.
func (m *InMemStore) Delete(sessionId string) error {
	m.logger.Log("method", "delete", "sessionId", sessionId)
	m.mu.Lock()
	delete(m.items, sessionId)
	m.mu.Unlock()

	return nil
}

// Reset extend a session ttl from the InMemStore instance.
func (m *InMemStore) Reset(sessionId string, expiration time.Time) ([]byte, bool, error) {
	m.logger.Log("method", "reset")
	m.mu.RLock()
	item, found := m.items[sessionId]
	if !found {
		return nil, false, nil
	}

	if time.Now().UnixNano() > item.Expiration {
		m.logger.Log("action", "expired", "sessionId", sessionId)
		return nil, false, nil
	}
	m.mu.RUnlock()

	m.mu.Lock()
	item.Expiration = expiration.UnixNano()
	m.items[sessionId] = item
	m.mu.Unlock()
		
	return item.Oject, true, nil
	
}

// List
func (m *InMemStore) List() (map[string]Item, error) {
	m.logger.Log("method", "list")
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.items, nil
}

// startSessionCleanup
func (m *InMemStore) startSessionCleanup(interval time.Duration) {
	m.logger.Log("method", "startSessionCleanup")
	m.stopCleanup = make(chan bool)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			m.deleteSessionExpired()
		case <-m.stopCleanup:
			ticker.Stop()
			return
		}
	}
}

// StopSessionCleanup terminates the background cleanup goroutine for the InMemStore
// instance. It's rare to terminate this; generally InMemStore instances and
// their cleanup goroutines are intended to be long-lived and run for the lifetime
// of your application.
//
// There may be occasions though when your use of the InMemStore is transient.
// An example is creating a new InMemStore instance in a test function. In this
// scenario, the cleanup goroutine (which will run forever) will prevent the
// InMemStore object from being garbage collected even after the test function
// has finished. You can prevent this by manually calling StopSessionCleanup.
func (m *InMemStore) StopSessionCleanup() {
	m.logger.Log("stopCleanup")
	if m.stopCleanup != nil {
		m.stopCleanup <- true
	}
}

// deleteSessionExpired
func (m *InMemStore) deleteSessionExpired() {
	m.logger.Log("deleteSessionExpired")
	now := time.Now().UnixNano()
	m.mu.Lock()
	for sessionId, item := range m.items {
		if now > item.Expiration {
			m.logger.Log("action", "session-expired", "sessionId", sessionId)
			delete(m.items, sessionId)
		}
	}
	m.mu.Unlock()
}