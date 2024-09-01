package models

import "time"

type Account interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPass() string
	GetCreatedAt() time.Time
	GetUpdatedAt() *time.Time
	Touch()
}

type account struct {
	model
	name      string
	email     string
	pass      string
	updatedAt *time.Time
}

func (m *account) GetID() string {
	return m.id
}

func (u *account) GetName() string {
	return u.name
}

func (u *account) GetEmail() string {
	return u.email
}

func (u *account) GetPass() string {
	return u.pass
}

func (u *account) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *account) GetUpdatedAt() *time.Time {
	return u.updatedAt
}

func (u *account) Touch() {
	now := time.Now()
	u.updatedAt = &now
}

func NewAccount(name, email, pass string) Account {
	return &account{
		model: model{
			id:        generateUUID(),
			createdAt: time.Now(),
		},
		name:      name,
		email:     email,
		pass:      pass,
		updatedAt: nil,
	}
}

func NewExistingAccount(id, name, email, pass string, createdAt, updatedAt time.Time) Account {
	return &account{
		model: model{
			id:        id,
			createdAt: createdAt,
		},
		name:      name,
		email:     email,
		pass:      pass,
		updatedAt: &updatedAt,
	}
}
