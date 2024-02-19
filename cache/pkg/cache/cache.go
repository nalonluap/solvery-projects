package cache

import (
	"time"
)

type User struct {
	ID     string
	Orders []Order
}

type Order struct {
	ID     string
	Amount int
}

type item struct {
	value      User
	expiration int64
}

type Cache struct {
	items map[string]item
	ttl   int64 // Время жизни элементов в кэше в секундах
}

func NewCache(ttl int64) *Cache {
	return &Cache{
		items: make(map[string]item),
		ttl:   ttl,
	}
}

func (c *Cache) Add(user User) {
	expiration := time.Now().Add(time.Duration(c.ttl) * time.Second).Unix()
	c.items[user.ID] = item{
		value:      user,
		expiration: expiration,
	}
}

func (c *Cache) Get(userID string) (User, bool) {
	item, found := c.items[userID]
	if !found {
		return User{}, false
	}

	if time.Now().Unix() > item.expiration {
		delete(c.items, userID)
		return User{}, false
	}

	return item.value, true
}
