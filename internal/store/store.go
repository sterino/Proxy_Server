// store/store.go
package store

import (
	models "app/internal/domain/proxy"
	"net/http"
	"sync"
)

type Store struct {
	sync.Map
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Set(id string, req models.RequestProxy, statusCode int, head http.Header) models.ResponseProxy {

	header := make(map[string]string)
	for k, v := range head {
		header[k] = v[0]
	}

	resp := models.ResponseProxy{
		ID:      id,
		Status:  statusCode,
		Headers: header,
		Length:  len(head),
	}
	s.Store(id+"_resp", resp)
	s.Store(id+"_req", req)

	return resp
}

func (s *Store) SetError(id string, req interface{}, statusCode int, err string) models.Error {
	errMap := make(map[string]string)
	errMap["error"] = err

	resp := models.Error{
		ID:      id,
		Status:  statusCode,
		Headers: errMap,
		Message: err,
	}
	s.Store(id+"_resp", resp)
	s.Store(id+"_req", req)

	return resp
}

func (s *Store) Get(id string) (interface{}, interface{}, bool) {
	res, found := s.Load(id + "_resp")
	req, _ := s.Load(id + "_req")

	return res, req, found
}

func (s *Store) GetAll() (map[string]interface{}, bool) {
	resp := make(map[string]interface{})
	s.Range(func(k, v interface{}) bool {
		resp[k.(string)] = v
		return true
	})

	return resp, len(resp) > 0
}
