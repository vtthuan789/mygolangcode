package gohttpmock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mock map[string]*Mock
}

func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body != "" {
		body = strings.ReplaceAll(body, "\t", "")
		body = strings.ReplaceAll(body, "\n", "")
	}

	return body
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mock[key] = &mock
}

func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mock = make(map[string]*Mock)
}
