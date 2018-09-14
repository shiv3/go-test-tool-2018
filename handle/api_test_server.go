package handle

import (
	"encoding/json"
	"net/http"
)

func handlerRuok(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type","text/plain")
		w.Write([]byte("imok"))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handlerJson(w http.ResponseWriter,r *http.Request){
	j,_ := json.Marshal(map[string]interface{}{
		"test":1,
		"test2":"2",
		"test3":[]int{1,2,3},
	})

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type","application/json")
		w.Write(j)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Handler() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/ruok",handlerRuok)
	mux.HandleFunc("/json",handlerJson)
	return mux
}