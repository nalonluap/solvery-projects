package tests

import (
	"simple-cache/pkg/cache"
	"testing"
	"time"
)

// TestNewCache проверяет создание нового кэша.
func TestNewCache(t *testing.T) {
	c := cache.NewCache(5)
	if c == nil {
		t.Error("NewCache() не создал новый экземпляр кэша.")
	}
}

// TestCacheAddGet проверяет добавление и получение элементов из кэша.
func TestCacheAddGet(t *testing.T) {
	c := cache.NewCache(5)
	user := cache.User{ID: "123", Orders: nil}
	c.Add(user)

	gotUser, ok := c.Get("123")
	if !ok || gotUser.ID != user.ID {
		t.Errorf("Get() не вернул ожидаемого пользователя: получено %v, ожидалось %v", gotUser, user)
	}
}

// TestCacheExpiration проверяет, что элементы удаляются из кэша после истечения TTL.
func TestCacheExpiration(t *testing.T) {
	c := cache.NewCache(5)
	user := cache.User{ID: "123", Orders: nil}
	c.Add(user)

	time.Sleep(6 * time.Second) // Ждем, пока элемент истечет

	_, ok := c.Get("123")
	if ok {
		t.Error("Элемент кэша не был удален после истечения TTL.")
	}
}

// Дополнительные тесты могут быть добавлены для проверки других функций и случаев использования.
