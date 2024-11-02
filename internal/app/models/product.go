package models

import "time"

type Product interface {
	GetID() string
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetPrice() int64
	SetPrice(price int64)
	GetStock() int
	SetStock(qty int)
	GetImageUrl() string
	SetImageUrl(url string)
	GetExternalID() string
	SetExternalID(externalId string)
	GetCreatedAt() time.Time
	GetUpdatedAt() *time.Time
	Touch()
}

type product struct {
	model
	name          string
	description   string
	price         int64
	stockQuantity int
	imageUrl      string
	externalID    string
	updatedAt     *time.Time
}

func (m *product) GetID() string {
	return m.id
}

func (m *product) GetName() string {
	return m.name
}

func (m *product) SetName(name string) {
	m.name = name
}

func (m *product) GetDescription() string {
	return m.description
}

func (m *product) SetDescription(description string) {
	m.description = description
}

func (m *product) GetPrice() int64 {
	return m.price
}

func (m *product) SetPrice(price int64) {
	m.price = price
}

func (m *product) GetStock() int {
	return m.stockQuantity
}

func (m *product) SetStock(qty int) {
	m.stockQuantity = qty
}

func (m *product) GetImageUrl() string {
	return m.imageUrl
}

func (m *product) SetImageUrl(url string) {
	m.imageUrl = url
}

func (m *product) GetExternalID() string {
	return m.externalID
}

func (m *product) SetExternalID(externalId string) {
	m.externalID = externalId
}

func (m *product) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *product) GetUpdatedAt() *time.Time {
	return m.updatedAt
}

func (m *product) Touch() {
	now := time.Now()
	m.updatedAt = &now
}

func NewProduct(name string, description string, price int64, stockQty int, imageUrl string) Product {
	return &product{
		model: model{
			id:        generateUUID(),
			createdAt: time.Now(),
		},
		name:          name,
		description:   description,
		price:         price,
		stockQuantity: stockQty,
		imageUrl:      imageUrl,
		externalID:    "",
		updatedAt:     nil,
	}
}

func NewExistingProduct(id, name string, description string, price int64, stockQty int, imageUrl string, externalID string, createdAt time.Time, updatedAt *time.Time) Product {
	return &product{
		model: model{
			id:        id,
			createdAt: createdAt,
		},
		name:          name,
		description:   description,
		price:         price,
		stockQuantity: stockQty,
		imageUrl:      imageUrl,
		externalID:    externalID,
		updatedAt:     updatedAt,
	}
}
