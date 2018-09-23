package handle

import (
	"encoding/json"
	"net/http"
	"time"
)

func handlerRuok(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("imok"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetOnlyJSON(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	}
}

func handlerJSON(w http.ResponseWriter, r *http.Request) {
	k, ok := r.URL.Query()["v"]
	if !ok || len(k) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	j, _ := json.Marshal(map[string]interface{}{
		"test":  1,
		"test2": k[0],
		"test3": []int{1, 2, 3},
	})
	w.Write(j)
}

func handlerSlow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("imok"))
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ruok", handlerRuok)
	mux.HandleFunc("/json", GetOnlyJSON(handlerJSON))
	return mux
}
