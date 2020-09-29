package amoeba

import (
	"time"
)

var schemaMap map[string]*Schema
var mountRules func() (map[string]*Schema, error)

func Start(r func() (map[string]*Schema, error)) {
	mountRules = r
	m := &Mount{
		ticker: time.NewTicker(5 * time.Second),
	}
	m.start()
}

func (m *Mount) start() {
	newSchemaMap, err := mountRules()
	if err != nil {
		// TODO ERROR
		return
	}
	schemaMap = newSchemaMap
	go withRecover(m.reload)
}

type Mount struct {
	ticker *time.Ticker
}

func (m *Mount) reload() {
	defer m.ticker.Stop()
	for {
		select {
		case <-m.ticker.C:
			newSchemaMap, err := mountRules()
			if err != nil {
				// TODO ERROR
			}
			schemaMap = newSchemaMap
		}
	}
}
