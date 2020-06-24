package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(client client.Client) Client {
	return Client{
		client: client,
	}
}

type Client struct {
	client client.Client
}

// Get provides a thin wrapper and client.client to access corev1.Service types
func (c Client) Get(key client.ObjectKey) (corev1.Service, error) {
	svc := corev1.Service{}
	if err := c.client.Get(context.TODO(), key, &svc); err != nil {
		return corev1.Service{}, err
	}
	return svc, nil
}

// Update provides a thin wrapper and client.Client to update corev1.Service types
func (c Client) Update(svc corev1.Service) error {
	if err := c.client.Update(context.TODO(), &svc); err != nil {
		return err
	}
	return nil
}

// Create provides a thin wrapper and client.Client to create corev1.Service types
func (c Client) Create(svc corev1.Service) error {
	if err := c.client.Create(context.TODO(), &svc); err != nil {
		return err
	}
	return nil
}

// Delete provides a thin wrapper and client.Client to delete corev1.Service types
func (c Client) Delete(svc corev1.Service) error {
	if err := c.client.Delete(context.TODO(), &svc); err != nil {
		return err
	}
	return nil
}