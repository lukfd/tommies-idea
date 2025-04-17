package routes

import (
	"fmt"
	"net/http"

	"github.com/lukfd/redis-demo/internal"
)

func Idea(w http.ResponseWriter, r *http.Request, redis *internal.RedisClient) {
	id := r.PathValue("id")
	val, err := redis.Client.Get(redis.Ctx, id).Result()
	if err != nil {
		http.Error(w, "Error fetching data from Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the value from Redis
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Value from Redis: %s", val)))
}
