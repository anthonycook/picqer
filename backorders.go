package picqer

import (
	"encoding/json"
	"strconv"
)

/*
BackorderService is for interfacing with the
backorder endpoints of the Picqer API.
See: https://picqer.com/en/api/backorders
*/
type BackorderService interface {
	List(int) ([]Backorder, error)
	Get(int) (*Backorder, error)
	GetForOrder(int, int) ([]Backorder, error)
	Delete(int) error
	Process() error
}

/*
BackorderServiceOp handles communication with the
backorder related methods of Picqers API
See: https://picqer.com/en/api/backorders
*/
type BackorderServiceOp struct {
	client *Client
}

/*
Backorder represents a single Picqer backorder
See: https://picqer.com/en/api/backorders
*/
type Backorder struct {
	Idbackorder       int    `json:"idbackorder"`
	Idorder           int    `json:"idorder"`
	Idproduct         int    `json:"idproduct"`
	Idcustomer        int    `json:"idcustomer"`
	Idwarehouse       int    `json:"idwarehouse"`
	Amount            int    `json:"amount"`
	AmountAvailable   int    `json:"amount_available"`
	Priority          int    `json:"priority"`
	PartOfIdbackorder int    `json:"part_of_idbackorder"`
	PartsPerParent    int    `json:"parts_per_parent"`
	CreatedAt         string `json:"created_at"`
	DateAvailable     string `json:"date_available"`
}

/*
List returns a list of all backorders and is limited up
to 100 results at a time, it accepts an offest param for pagination.
See: https://picqer.com/en/api/backorders
*/
func (s *BackorderServiceOp) List(offset int) ([]Backorder, error) {
	b := []Backorder{}
	offsetStr := strconv.Itoa(offset)

	res, err := s.client.NewRequest("GET", "backorders?offset="+offsetStr, nil)
	if err != nil {
		return b, err
	}

	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

/*
Get returns a single backorder matching the prdoivded backorder ID
See: https://picqer.com/en/api/backorders
*/
func (s *BackorderServiceOp) Get(id int) (*Backorder, error) {
	b := Backorder{}
	idStr := strconv.Itoa(id)

	res, err := s.client.NewRequest("GET", "backorders/"+idStr, nil)
	if err != nil {
		return &b, err
	}

	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		return &b, err
	}

	return &b, nil
}

/*
GetForOrder gets all backorders for a given order using a provided order ID
and is limited up to 100 results at a time, it accepts an offest param for pagination.
See: https://picqer.com/en/api/backorders
*/
func (s *BackorderServiceOp) GetForOrder(id, offset int) ([]Backorder, error) {
	b := []Backorder{}
	idStr := strconv.Itoa(id)
	offsetStr := strconv.Itoa(offset)

	res, err := s.client.NewRequest("GET", "orders/"+idStr+"/backorders?offset="+offsetStr, nil)
	if err != nil {
		return b, err
	}

	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

/*
Delete will delete a backorder using a provided backorder ID
See: https://picqer.com/en/api/backorders
*/
func (s *BackorderServiceOp) Delete(id int) error {
	idStr := strconv.Itoa(id)

	_, err := s.client.NewRequest("DELETE", "backorders/"+idStr, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
Process backorders will process all current backorders
See: https://picqer.com/en/api/backorders
*/
func (s *BackorderServiceOp) Process() error {
	_, err := s.client.NewRequest("POST", "backorders/process", nil)
	if err != nil {
		return err
	}

	return nil
}
