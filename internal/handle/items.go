package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerodays/sistem-auth/middleware"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-inventory/internal/handle/errors"
	"github.com/zerodays/sistem-inventory/internal/logger/item"
	"net/http"
	"strconv"
)

func ListItemsHandle(w http.ResponseWriter, r *http.Request) {
	u := middleware.UserFromRequest(r)

	// user needs to have read permission in order to list inventory items
	if !u.HasPermissions(permission.InventoryRead) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	items, err := item.ListItems()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []byte

	// TODO to se sigurno da naredit boljse, hocem nekak dobit takle json: "[]" in ne "null" kadar ni userjev
	if items != nil {
		res, err = json.Marshal(items)
	} else {
		res = []byte("[]")
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(res)
}

func GetItemHandle(w http.ResponseWriter, r *http.Request) {
	u := middleware.UserFromRequest(r)

	// check permission
	if !u.HasPermissions(permission.InventoryRead) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// get item id
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errors.Response(w, errors.InvalidParams)
		return
	}

	// read item from database
	itm, err := item.ForID(id)

	if err != nil {
		errors.Response(w, errors.DatabaseError)
		return
	}

	itemResponse(w, *itm)
}

func NewItemHandle(w http.ResponseWriter, r *http.Request) {
	u := middleware.UserFromRequest(r)

	if !u.HasPermissions(permission.InventoryWrite) {
		// in order to create items, you need to have write permission
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	var itm item.Item
	err := json.NewDecoder(r.Body).Decode(&itm)

	if err != nil {
		errors.Response(w, errors.InvalidJSON)
		return
	}

	err = itm.Insert()

	if err != nil {
		errors.Response(w, errors.DatabaseError)
		return
	}

	itemResponse(w, itm)
}

func itemResponse(w http.ResponseWriter, itm item.Item) {
	res, err := json.Marshal(itm)

	if err != nil {
		errors.Response(w, errors.UnknownError)
		return
	}

	_, _ = w.Write(res)
}
