package picqer

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/*
ProductService is for interfacing with the
product endpoints of the Picqer API.
See: https://picqer.com/en/api/products
*/
type ProductService interface {
	List() (*[]Product, error)
	Get(int) (*Product, error)
	Create(Product) (*Product, error)
	Update(Product) (*Product, error)
	Delete(int) error
}

/*
ProductServiceOp handles communication with the
product related methods of Picqers API
See: https://picqer.com/en/api/products
*/
type ProductServiceOp struct {
	client *Client
}

/*
Product represents a single Picqer product
See: https://picqer.com/en/api/products
*/
type Product struct {
	IDProduct                  int             `json:"idproduct,omitempty"`
	IDVatGroup                 int             `json:"idvatgroup,omitempty"`
	IDSupplier                 int             `json:"idsupplier,omitempty"`
	IDPicklistProduct          int             `json:"idpicklist_product,omitempty"`
	IDOrderProduct             int             `json:"idorder_product,omitempty"`
	IDReturnProductReplacement int             `json:"idreturn_product_replacement,omitempty"`
	Name                       string          `json:"name,omitempty"`
	Price                      float64         `json:"price,omitempty"`
	FixedStockPrice            float64         `json:"fixedstockprice,omitempty"`
	ProductCode                string          `json:"productcode,omitempty"`
	ProductCodeSupplier        string          `json:"productcode_supplier,omitempty"`
	DeliveryTime               string          `json:"deliverytime,omitempty"`
	Description                string          `json:"description,omitempty"`
	Barcode                    string          `json:"barcode,omitempty"`
	UnlimitedStock             bool            `json:"unlimitedstock,omitempty"`
	StockLocation              string          `json:"stocklocation,omitempty"`
	PartOfIDPicklistProduct    int             `json:"partof_idpicklist_product,omitempty"`
	PartofIDorderProduct       int             `json:"partof_idorder_product,omitempty"`
	HasParts                   bool            `json:"has_parts,omitempty"`
	Weight                     int             `json:"weight,omitempty"`
	Length                     int             `json:"length,omitempty"`
	Width                      int             `json:"width,omitempty"`
	Height                     int             `json:"height,omitempty"`
	Active                     bool            `json:"active,omitempty"`
	ProductFields              []Productfields `json:"productfields,omitempty"`
	Images                     []string        `json:"images,omitempty"`
	Stock                      []Stock         `json:"stock,omitempty"`
	Tags                       map[string]Tag  `json:"tags,omitempty"`
	Amount                     int             `json:"amount,omitempty"`
	AmountReceived             int             `json:"amountreceived,omitempty"`
	AmountPicked               int             `json:"amount_picked,omitempty"`
	AmountCancelled            int             `json:"amount_cancelled,omitempty"`
	Remarks                    string          `json:"remarks,omitempty"`
}

type Productfields struct {
	IDProductField int    `json:"idproductfield,omitempty"`
	Title          string `json:"title,omitempty"`
	Value          string `json:"value,omitempty"`
}

type Stock struct {
	IDWarehouse         int `json:"idwarehouse,omitempty"`
	Stock               int `json:"stock,omitempty"`
	Reserved            int `json:"reserved,omitempty"`
	ReservedBackorders  int `json:"reservedbackorders,omitempty"`
	ReservedPicklists   int `json:"reservedpicklists,omitempty"`
	ReservedAllocations int `json:"reservedallocations,omitempty"`
	Freestock           int `json:"freestock,omitempty"`
}

type Tag struct {
	IDTag     int    `json:"idtag,omitempty"`
	Title     string `json:"title,omitempty"`
	Color     string `json:"color,omitempty"`
	Inherit   bool   `json:"inherit,omitempty"`
	TextColor string `json:"textColor,omitempty"`
}

/*
List returns a list of all products and is limited up
to 100 results at a time, it accepts an offest param for pagination.
See: https://picqer.com/en/api/products
*/
func (s *ProductServiceOp) List() (*[]Product, error) {
	p := []Product{}

	res, err := s.client.NewRequest("GET", "products", nil)
	if err != nil {
		return &p, err
	}

	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

/*
Create and save a new product, the response from Picqer is returned
See: https://picqer.com/en/api/products
*/
func (s *ProductServiceOp) Create(p Product) (*Product, error) {
	fmt.Println(s.client.BaseURL)
	return &p, nil
}

// Update an existing product
func (s *ProductServiceOp) Update(p Product) (*Product, error) {
	return &p, nil
}

/*
Get returns a single product matching the prdoivded product ID
See: https://picqer.com/en/api/products
*/
func (s *ProductServiceOp) Get(id int) (*Product, error) {
	p := Product{}
	idString := strconv.Itoa(id)

	res, err := s.client.NewRequest("GET", "products/"+idString, nil)
	if err != nil {
		return &p, err
	}

	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

/*
Delete will delete a product using a provided product ID
See: https://picqer.com/en/api/products
*/
func (s *ProductServiceOp) Delete(id int) error {
	idString := strconv.Itoa(id)

	_, err := s.client.NewRequest("DELETE", "products/"+idString, nil)
	if err != nil {
		return err
	}

	return nil
}
