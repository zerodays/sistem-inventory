package handle

import (
	"encoding/json"
	"net/http"
)

func MilestoneInfoHandle(w http.ResponseWriter, _ *http.Request) {
	res, _ := json.Marshal(map[string]interface{}{
		"clani":         []string{"mm8155", "vd3217"},
		"opis_projekta": "Sistem za vodenje trenutnih projektov, izplačil in inventarja podjetja.",
		"mikrostoritve": []string{"http://20.62.232.75/api/v1/items", "http://20.62.234.124:8080/api/v1/users"},
		"github":        []string{"https://github.com/zerodays/sistem-inventory", "https://github.com/zerodays/sistem-users"},
		"travis":        []string{"https://travis-ci.org/github/zerodays/sistem-inventory", "https://github.com/zerodays/sistem-users"},
		"dockerhub":     []string{"https://hub.docker.com/r/zdays/sistem-inventory", "https://hub.docker.com/r/zdays/sistem-users"},
	})

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
