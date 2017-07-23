package weblru

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rileyr/golru/lru"
	"github.com/rileyr/middleware/wares"
)

type Handler struct {
	cache lru.Cache
}

func New(s int) *Handler {
	return &Handler{cache: lru.New(lru.WithSize(s))}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	key := r.URL.Query().Get("key")
	log.Printf("getting: %s\n", key)
	val, ok := h.cache.Get(key)
	if !ok {
		log.Printf("miss: %s\n", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Printf("found: %s\n", key)

	resp := &Request{Key: key, Value: val}
	bts, err := json.Marshal(resp)
	if err != nil {
		h.handleError(w, r, p, err)
		return
	}

	w.Write(bts)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !req.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("adding: %s\n", req.Key)

	ok := h.cache.Add(req.Key, req.Value)
	if ok {
		log.Printf("evicted for add: %s", req.Key)
	}

	bts, _ := json.Marshal(req)
	w.Write(bts)

}

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, r, p, err)
		return
	}

	log.Printf("removing: %s\n", req.Key)
	h.cache.Remove(req.Key)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleError(w http.ResponseWriter, r *http.Request, p httprouter.Params, e error) {
	rid := r.Context().Value(wares.RequestIDContextKey{}).(string)
	log.Printf("ERROR: %s -- %s -- %s -- %s", rid, r.Method, r.URL.Path, e.Error())
	w.WriteHeader(http.StatusInternalServerError)
	return
}
