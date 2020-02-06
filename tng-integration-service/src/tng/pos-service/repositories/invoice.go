package repositories

import (
	"tng/common/models/pos"
	"tng/common/utils/db"
)

type InvoiceRepository interface {
	Create(*db.DB, *pos.Invoice) error
}

type invoiceRepository struct{}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}

func (r *invoiceRepository) Create(ormer *db.DB, invoice *pos.Invoice) error {
	_, err := ormer.InsertOrUpdate(invoice)
	return err
}
