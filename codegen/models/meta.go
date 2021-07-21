package models

import (
	"errors"
	"strconv"
)

var (
	ErrNotFound = errors.New("not found")
)

type Meta struct {
	Fields map[string]string
}

func (m *Meta) getMap() map[string]string {
	if m.Fields == nil {
		m.Fields = map[string]string{}
	}
	return m.Fields
}

func (m *Meta) GetInt(name string) (int, error) {
	d := m.getMap()
	id, ok := d[name]
	if !ok {
		return 0, ErrNotFound
	}

	idInt, err := strconv.Atoi(id)
	return idInt, err
}

func (m *Meta) GetString(name string) (string, error) {
	d := m.getMap()
	id, ok := d[name]
	if !ok {
		return "", ErrNotFound
	}
	return id, nil
}
