package models

import "time"

type Category interface {
	GetID() string
	GetName() string
	GetCreatedAt() time.Time
}

type catyegory struct {
	model
	name string
}

func (m *catyegory) GetID() string {
	return m.id
}

func (m *catyegory) GetName() string {
	return m.name
}

func (m *catyegory) GetCreatedAt() time.Time {
	return m.createdAt
}

func NewCategory(name string) Category {
	return &catyegory{
		model: model{
			id:        generateUUID(),
			createdAt: time.Now(),
		},
		name: name,
	}
}

func NewExistingCategory(id, name string, createdAt time.Time) Category {
	return &catyegory{
		model: model{
			id:        id,
			createdAt: createdAt,
		},
		name: name,
	}
}
