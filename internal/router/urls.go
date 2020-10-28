package router

import (
	"github.com/zerodays/sistem-inventory/internal/handle"
	"net/http"
)

// apiRoutes returns routes. It is in function instead of a variable to allow
// lazy loading.
func apiRoutes() []Route {
	return []Route{
		{
			Name:              "list_items",
			Path:              "/items",
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.ListItemsHandle),
		},
		{
			Name:              "new_item",
			Path:              "/items",
			CustomContentType: true,
			POST:               http.HandlerFunc(handle.NewItem),
		},
	}
}
