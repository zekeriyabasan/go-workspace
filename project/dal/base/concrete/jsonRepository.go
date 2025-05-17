package concrete

import (
	"bytes"
	"encoding/json"
	"everymore-go/project/dal/base/contracts"
	"fmt"
	"io"
	"net/http"
)

type JsonRepository[T contracts.Entity] struct {
	resource string
	baseURL  string
	client   *http.Client
}

func NewJsonRepository[T contracts.Entity](baseURL string, resource string) *JsonRepository[T] {
	return &JsonRepository[T]{
		resource: resource,
		baseURL:  baseURL,
		client:   &http.Client{},
	}
}

func (r *JsonRepository[T]) Add(entity T) error {
	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s", r.baseURL, r.resource)
	resp, err := r.client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (r *JsonRepository[T]) Delete(id int) error {
	url := fmt.Sprintf("%s/%s/%d", r.baseURL, r.resource, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (r *JsonRepository[T]) GetAll() ([]T, error) {
	url := fmt.Sprintf("%s/%s", r.baseURL, r.resource)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []T
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
