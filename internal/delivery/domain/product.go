package domain

import (
	errs "github.com/rattapon001/inventory-ddd/internal/delivery/errors"
)

type ProductStatus string

const (
	ProductStatusPending ProductStatus = "pending"
	ProductStatusPass    ProductStatus = "pass"
	ProductStatusReject  ProductStatus = "reject"
)

type Product struct {
	Name   string
	Sku    string
	Qty    int16
	Status ProductStatus
}

func NewProduct(name string, sku string, qty int16) *Product {
	return &Product{
		Name:   name,
		Sku:    sku,
		Qty:    qty,
		Status: ProductStatusPending,
	}
}

func (p *Product) Pass() error {
	if p.Status != ProductStatusPending {
		return errs.ErrCannotPassStatusNotPending
	}

	p.Status = ProductStatusPass
	return nil
}

func (p *Product) Reject() error {
	if p.Status != ProductStatusPending {
		return errs.ErrCannotRejectStatusNotPending
	}

	p.Status = ProductStatusReject
	return nil
}
