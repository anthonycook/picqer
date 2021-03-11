package picqer

import (
	"encoding/json"
	"strconv"
)

// OrderService is an interface for interfacing with the
// order endpoints of the Picqer API.
// See: https://picqer.com/en/api/orders
type OrderService interface {
	List(int) ([]Order, error)
	Get(int) (*Order, error)
	Create(Order) (*Order, error)
	Update(Order) (*Order, error)
	Delete(int) error
}

// OrderServiceOp handles communication with the order
// related methods of Picqers API
type OrderServiceOp struct {
	client *Client
}

// Order represents a Picqer backorder
type Order struct {
	IDOrder                   int            `json:"idorder,omitempty"`
	IDCustomer                int            `json:"idcustomer,omitempty"`
	IDTemplate                int            `json:"idtemplate,omitempty"`
	IDShippingProviderProfile int            `json:"idshippingprovider_profile,omitempty"`
	OrderID                   string         `json:"orderid,omitempty"`
	DeliveryName              string         `json:"deliveryname,omitempty"`
	DeliveryContactName       string         `json:"deliverycontactname,omitempty"`
	DeliveryAddress           string         `json:"deliveryaddress,omitempty"`
	DeliveryAddress2          string         `json:"deliveryaddress2,omitempty"`
	DeliveryZipCode           string         `json:"deliveryzipcode,omitempty"`
	DeliveryCity              string         `json:"deliverycity,omitempty"`
	DeliveryRegion            string         `json:"deliveryregion,omitempty"`
	DeliveryCountry           string         `json:"deliverycountry,omitempty"`
	FullDeliveryAddress       string         `json:"full_delivery_address,omitempty"`
	InvoiceName               string         `json:"invoicename,omitempty"`
	InvoiceContactName        string         `json:"invoicecontactname,omitempty"`
	InvoiceAddress            string         `json:"invoiceaddress,omitempty"`
	InvoiceAddress2           string         `json:"invoiceaddress2,omitempty"`
	InvoiceZipCode            string         `json:"invoicezipcode,omitempty"`
	InvoiceCity               string         `json:"invoicecity,omitempty"`
	InvoiceRegion             string         `json:"invoiceregion,omitempty"`
	InvoiceCountry            string         `json:"invoicecountry,omitempty"`
	FullInvoiceAddress        string         `json:"full_invoice_address,omitempty"`
	Telephone                 string         `json:"telephone,omitempty"`
	EmailAddress              string         `json:"emailaddress,omitempty"`
	Reference                 string         `json:"reference,omitempty"`
	CustomerRemarks           string         `json:"customer_remarks,omitempty"`
	PickupPointData           string         `json:"pickup_point_data,omitempty"`
	PartialDelivery           bool           `json:"partialdelivery,omitempty"`
	AutoSplit                 bool           `json:"auto_split,omitempty"`
	Invoiced                  bool           `json:"invoiced,omitempty"`
	PreferredDeliveryDate     string         `json:"preferred_delivery_date,omitempty"`
	Discount                  int            `json:"discount,omitempty"`
	CalculateVat              bool           `json:"calculatevat,omitempty"`
	Status                    string         `json:"status,omitempty"`
	PublicStatusPage          string         `json:"public_status_page,omitempty"`
	Created                   string         `json:"created,omitempty"`
	Updated                   string         `json:"updated,omitempty"`
	Warehouses                []int          `json:"warehouses,omitempty"`
	Products                  []Product      `json:"products,omitempty"`
	Tags                      map[string]Tag `json:"tags,omitempty,omitempty"`
	Orderfields               []Orderfields  `json:"orderfields,omitempty"`
}

// Orderfields represents the orderfields of a Picqer order
type Orderfields struct {
	IDorderField int    `json:"idorderfield,omitempty"`
	Title        string `json:"title,omitempty"`
	Value        string `json:"value,omitempty"`
}

// List orders
func (s *OrderServiceOp) List(offset int) ([]Order, error) {
	o := []Order{}
	offsetStr := strconv.Itoa(offset)

	res, err := s.client.NewRequest("GET", "orders?offset="+offsetStr, nil)
	if err != nil {
		return o, err
	}

	err = json.NewDecoder(res.Body).Decode(&o)
	if err != nil {
		return o, err
	}

	return o, nil
}

// Create order
func (s *OrderServiceOp) Create(o Order) (*Order, error) {
	return &o, nil
}

// Update order
func (s *OrderServiceOp) Update(o Order) (*Order, error) {
	return &o, nil
}

// Get individual order
func (s *OrderServiceOp) Get(id int) (*Order, error) {
	o := Order{}
	idStr := strconv.Itoa(id)

	res, err := s.client.NewRequest("GET", "orders/"+idStr, nil)
	if err != nil {
		return &o, err
	}

	err = json.NewDecoder(res.Body).Decode(&o)
	if err != nil {
		return &o, err
	}

	return &o, nil
}

// Delete an order
func (s *OrderServiceOp) Delete(id int) error {
	idStr := strconv.Itoa(id)

	_, err := s.client.NewRequest("DELETE", "orders/"+idStr, nil)
	if err != nil {
		return err
	}

	return nil
}
