package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/lukfd/redis-demo/model"
)

func GetIdea(w http.ResponseWriter, r *http.Request, redis *model.RedisClient) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	// Convert ID to int
	ideaId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	idea, err := redis.GetIdea(ideaId)

	if err != nil {
		http.Error(w, "Error fetching idea from Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := redis.GetBytes(idea)
	if err != nil {
		http.Error(w, "Error converting idea to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIdeas(w http.ResponseWriter, r *http.Request, redis *model.RedisClient) {
	ideas, err := redis.GetIdeas()
	if err != nil {
		http.Error(w, "Error fetching ideas from Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(ideas)
	if err != nil {
		http.Error(w, "Error converting ideas to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NewIdea(w http.ResponseWriter, r *http.Request, redis *model.RedisClient) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req model.Idea
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	redis.AddIdea(req)
	w.WriteHeader(http.StatusCreated)
}
