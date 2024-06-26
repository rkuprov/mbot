package store_test

//
// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"testing"
// 	"time"
//
// 	"github.com/stretchr/testify/assert"
// 	"go.etcd.io/bbolt"
//
// 	"github.com/rkuprov/mbot/pkg/store"
// )
//
// func NewTestClient(ctx context.Context) (*store.Store, func()) {
// 	var funcCollect = []func() error{}
// 	db, err := bbolt.Open("test.db", 0600, &bbolt.Options{Timeout: 1 * time.Second})
// 	if err != nil {
// 		return nil, nil
// 	}
// 	funcCollect = append(funcCollect, db.Close)
// 	funcCollect = append(funcCollect, func() error {
// 		fmt.Println("removing test.db")
// 		return os.Remove("test.db")
// 	})
// 	out, err := store.NewWithClient(ctx, db)
// 	if err != nil {
// 		return nil, func() {
// 			for _, f := range funcCollect {
// 				f()
// 			}
// 		}
//
// 	}
//
// 	return out, func() {
// 		for _, f := range funcCollect {
// 			f()
// 		}
// 	}
// }
//
// func TestAddCustomer(t *testing.T) {
// 	ctx := context.Background()
// 	c, cleanup := NewTestClient(ctx)
// 	defer cleanup()
//
// 	id, err := c.CreateCustomer(ctx, store.CustomerCreate{
// 		Name:    "Roman",
// 		Email:   "roman@gmail.com",
// 		Contact: "call me maybe",
// 	})
//
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, id)
//
// 	customer, err := c.GetCustomer(ctx, id)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Roman", customer.Name)
// 	assert.Equal(t, "roman@gmail.com", customer.Email)
// 	assert.Equal(t, "call me maybe", customer.Contact)
// }
//
// func TestGetCustomersAll(t *testing.T) {
// 	ctx := context.Background()
// 	c, cleanup := NewTestClient(ctx)
// 	defer cleanup()
//
// 	id1, err := c.CreateCustomer(ctx, store.CustomerCreate{
// 		Name:    "Roman",
// 		Email:   "roman@gmail.com",
// 		Contact: "call me maybe",
// 	})
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, id1)
//
// 	id2, err := c.CreateCustomer(ctx, store.CustomerCreate{
// 		Name:    "John",
// 		Email:   "john@doe.com",
// 		Contact: "1234567890",
// 	})
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, id2)
//
// 	customers, err := c.GetCustomersAll(ctx)
// 	assert.NoError(t, err)
// 	assert.Len(t, customers, 2)
// 	assert.Equal(t, customers[0].Name, "Roman")
// 	assert.Equal(t, customers[1].Name, "John")
// 	assert.Equal(t, customers[0].Email, "roman@gmail.com")
// 	assert.Equal(t, customers[1].Email, "john@doe.com")
// }
//
// func TestUpdateCustomer(t *testing.T) {
// 	ctx := context.Background()
// 	c, cleanup := NewTestClient(ctx)
// 	defer cleanup()
//
// 	id, err := c.CreateCustomer(ctx, store.CustomerCreate{
// 		Name:    "Roman",
// 		Email:   "roman@gmail.com",
// 		Contact: "call me maybe",
// 	})
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, id)
//
// 	err = c.UpdateCustomer(ctx, id, store.CustomerUpdate{
// 		Name:    "Roman's cat",
// 		Email:   "me@meow.com",
// 		Contact: "meow meow",
// 	})
//
// 	out, err := c.GetCustomer(ctx, id)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Roman's cat", out.Name)
// 	assert.Equal(t, "me@meow.com", out.Email)
// 	assert.Equal(t, "meow meow", out.Contact)
// }
