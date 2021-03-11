package picqer

import (
	"encoding/json"
	"strconv"
)

type PicklistService interface {
	List(int) ([]Picklist, error)
	Get(int) (*Picklist, error)
	Create(Picklist) (*Picklist, error)
	Update(Picklist) (*Picklist, error)
	Delete(int64) error
}

type PicklistServiceOp struct {
	client *Client
}

// Picklist is a single Picqer picklist respsonse
type Picklist struct {
	IDPicklist       int       `json:"idpicklist,omitempty"`
	PicklistID       string    `json:"picklistid,omitempty"`
	IDCustomer       int       `json:"idcustomer,omitempty"`
	IDOrder          int       `json:"idorder,omitempty"`
	IDWarehouse      int       `json:"idwarehouse,omitempty"`
	DeliveryName     string    `json:"deliveryname,omitempty"`
	DeliveryContact  string    `json:"deliverycontact,omitempty"`
	DeliveryAddress  string    `json:"deliveryaddress,omitempty"`
	DeliveryAddress2 string    `json:"deliveryaddress2,omitempty"`
	DeliveryZipCode  string    `json:"deliveryzipcode,omitempty"`
	DeliveryCity     string    `json:"deliverycity,omitempty"`
	DeliveryRegion   string    `json:"deliveryregion,omitempty"`
	DeliveryCountry  string    `json:"deliverycountry,omitempty"`
	EmailAddress     string    `json:"emailaddress,omitempty"`
	Telephone        string    `json:"telephone,omitempty"`
	Reference        string    `json:"reference,omitempty"`
	AssignedToIDUser int       `json:"assigned_to_iduser,omitempty"`
	Invoiced         bool      `json:"invoiced,omitempty"`
	Urgent           bool      `json:"urgent,omitempty"`
	Status           string    `json:"status,omitempty"`
	TotalProducts    int       `json:"totalproducts,omitempty"`
	TotalPicked      int       `json:"totalpicked,omitempty"`
	SnoozedUntil     string    `json:"snoozed_until,omitempty"`
	Created          string    `json:"created,omitempty"`
	Updated          string    `json:"updated,omitempty"`
	Products         []Product `json:"products,omitempty"`
}

// List all picklists
func (s *PicklistServiceOp) List(offset int) ([]Picklist, error) {
	p := []Picklist{}
	offsetStr := strconv.Itoa(offset)

	res, err := s.client.NewRequest("GET", "backorders?offset="+offsetStr, nil)
	if err != nil {
		return p, err
	}

	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// Create a new picklist
func (s *PicklistServiceOp) Create(p Picklist) (*Picklist, error) {
	return &p, nil
}

// Update an existing picklist
func (s *PicklistServiceOp) Update(p Picklist) (*Picklist, error) {
	return &p, nil
}

// Get a picklist by ID
func (s *PicklistServiceOp) Get(id int) (*Picklist, error) {
	p := Picklist{}
	idStr := strconv.Itoa(id)

	res, err := s.client.NewRequest("GET", "picklists/"+idStr, nil)
	if err != nil {
		return &p, err
	}

	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

// Delete a picklist by ID
func (s *PicklistServiceOp) Delete(id int64) error {
	return nil
}
