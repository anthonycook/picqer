package picqer

import "time"

type PurchaseOrderService interface {
	List() ([]PurchaseOrder, error)
	Get(int64) (*PurchaseOrder, error)
	Create(PurchaseOrder) (*PurchaseOrder, error)
	Update(PurchaseOrder) (*PurchaseOrder, error)
	Delete(int64) error
}

type PurchaseOrderServiceOp struct {
	client *Client
}

type PurchaseOrder struct {
	IDPurchaseOrder int       `json:"idpurchaseorder,omitempty"`
	IDSupplier      int       `json:"idsupplier,omitempty"`
	IDWarehouse     int       `json:"idwarehouse,omitempty"`
	PurchaseOrderID string    `json:"purchaseorderid,omitempty"`
	SupplierName    string    `json:"supplier_name,omitempty"`
	SupplierOrderID int       `json:"supplier_orderid,omitempty"`
	Status          string    `json:"status,omitempty"`
	Remarks         string    `json:"remarks,omitempty"`
	DeliveryDate    time.Time `json:"delivery_date,omitempty"`
	Created         time.Time `json:"created,omitempty"`
	Updated         time.Time `json:"updated,omitempty"`
	Products        []Product `json:"products,omitempty"`
}

func (s *PurchaseOrderServiceOp) List() ([]PurchaseOrder, error) {
	p := []PurchaseOrder{}
	return p, nil
}

func (s *PurchaseOrderServiceOp) Create(p PurchaseOrder) (*PurchaseOrder, error) {
	return &p, nil
}

func (s *PurchaseOrderServiceOp) Update(p PurchaseOrder) (*PurchaseOrder, error) {
	return &p, nil
}

func (s *PurchaseOrderServiceOp) Get(id int64) (*PurchaseOrder, error) {
	p := PurchaseOrder{}
	return &p, nil
}

func (s *PurchaseOrderServiceOp) Delete(id int64) error {
	return nil
}
