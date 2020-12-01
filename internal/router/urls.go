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
			AuthorizedOnly:    true,
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.ListItemsHandle),
			POST:              http.HandlerFunc(handle.NewItemHandle),
		},
		{
			Name:              "new_item",
			Path:              "/items/{id}",
			AuthorizedOnly:    true,
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.GetItemHandle),
			PUT:               http.HandlerFunc(handle.EditItemHandle),
			DELETE:            http.HandlerFunc(handle.DeleteItemHandle),
		},
		{
			Name:              "milestone_info",
			Path:              "/milestone_info",
			AuthorizedOnly:    false,
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.MilestoneInfoHandle),
		},
	}
}
