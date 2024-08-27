// Created by Anh Cao on 27.08.2024.

package cache

import "github.com/AnhCaooo/go-web-server-template-generator/template/internal/models"

var Cache *models.Cache

// initialize a cache instance
func NewCache() {
	Cache = &models.Cache{
		Data: make(map[string]models.CacheValue),
	}
}
