package main

import "context"

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) OrderService {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}
