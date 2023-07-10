package mapstore

import (
	"day2/domain"
	"errors"
	"fmt"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (m MapStore) Create(c domain.Customer) error {
	if _, ok := m.store[c.ID]; ok {
		fmt.Println("\nKey already exists, can't add")
		return errors.New("\nKey already exists, can't add")
	}
	fmt.Println("\nKey doesn't exists, adding it")
	m.store[c.ID] = c
	return nil
}

func (m MapStore) Delete(s string) error {
	if _, ok := m.store[s]; ok {
		fmt.Println("\nKey exists.. deleting it")
		delete(m.store, s)
		return nil
	}
	fmt.Println("\nKey doesn't exists, adding it")
	return errors.New("Key doesn't exists, cannot delete it")
}

func (m MapStore) Update(s string, c domain.Customer) error {
	if _, ok := m.store[s]; ok {
		fmt.Println("\nKey exists, updating it ")
		m.store[s] = c
		return nil
	}
	fmt.Println("\nKey doesn't exists, cannpt update it")
	return errors.New("\nKey doesn't exists, cannpt update it")
}

func (m MapStore) GetAll() ([]domain.Customer, error) {
	if len(m.store) == 0 {
		fmt.Println("\nmap is empty")
		return nil, errors.New("\nno customers in the map")
	}
	values := make([]domain.Customer, 0, len(m.store))
	for _, v := range m.store {
		values = append(values, v)
	}
	return values, nil
}

func (m MapStore) GetByID(s string) (domain.Customer, error) {
	if _, ok := m.store[s]; ok {
		fmt.Println("\nKey exists, getting value of it ")
		return m.store[s], nil
	}
	fmt.Println("\nKey doesn't exists, cannot find a customer")
	return m.store[s], errors.New("\nKey doesn't exists, cannot find a customer")
}
