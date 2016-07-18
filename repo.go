package gorepo

import (
	"fmt"
	"sort"
	"sync"
)

type (
	Repository interface {
		GetAll() (interface{}, error)
		Get(int64) (interface{}, error)
		Save(interface{}) error
		Remove(...int64) (interface{}, error)
	}
)

var (
	reposMu sync.RWMutex
	repos   = make(map[string]Repository)
)

func Register(name string, r Repository) {
	reposMu.Lock()
	defer reposMu.Unlock()
	if r == nil {
		panic("repository: Register repository is nil")
	}
	if _, dup := repos[name]; dup {
		panic("repository: Register called twice for repository " + name)
	}
	repos[name] = r
}

func Unregister(name string) {
	reposMu.Lock()
	defer reposMu.Unlock()
	delete(repos, name)
}

func Repositories() []string {
	reposMu.RLock()
	defer reposMu.RUnlock()
	var list []string
	for name := range repos {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func Get(name string) (Repository, error) {
	reposMu.RLock()
	defer reposMu.RUnlock()
	action, ok := repos[name]
	if !ok {
		return nil, fmt.Errorf("repository: unknown repository %q (forgotten import?)", name)
	}
	return action, nil
}

func unregisterAllRepos() {
	reposMu.Lock()
	defer reposMu.Unlock()
	repos = make(map[string]Repository)
}
