package handle

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/sistem-inventory/internal/logger/item"
	"net/http"
)

func ListItemsHandle(w http.ResponseWriter, r *http.Request) {
	items, err := item.ListItems()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []byte

	// TODO to se sigurno da naredit boljse
	if items != nil {
		res, err = json.Marshal(items)
	} else {
		res =  []byte("[]")
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func NewItem(w http.ResponseWriter, r *http.Request) {
	// TODO authentication
	var item item.Item

	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(item.Name)
	fmt.Println(item.DatePurchased)
	fmt.Println(item.Description)
	fmt.Println(item.Price)

	err = item.Insert()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}