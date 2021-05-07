package picqer

import (
	"encoding/json"
)

/*
StatsService is for interfacing with the
stat endpoints of the Picqer API.
See: https://picqer.com/en/api/stats
*/
type StatsService interface {
	Get(string) (*Stat, error)
}

/*
StatServiceOp handles communication with the
stat related methods of Picqers API
See: https://picqer.com/en/api/stats
*/
type StatsServiceOp struct {
	client *Client
}

/*
Stat represents a single Picqer stat
See: https://picqer.com/en/api/stats
*/
type Stat struct {
	Value int `json:"value"`
}

/*
Get returns a single stat
See: https://picqer.com/en/api/stats
*/
func (s *StatsServiceOp) Get(q string) (*Stat, error) {
	stat := Stat{}

	res, err := s.client.NewRequest("GET", "stats/"+q, nil)
	if err != nil {
		return &stat, err
	}

	err = json.NewDecoder(res.Body).Decode(&stat)
	if err != nil {
		return &stat, err
	}

	return &stat, nil
}
