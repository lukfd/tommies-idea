package routes

import (
	"fmt"
	"net/http"

	"github.com/lukfd/redis-demo/model"
)

func Keyval(w http.ResponseWriter, r *http.Request, redis *model.RedisClient) {
	key := r.PathValue("key")
	val, err := redis.Client.Get(redis.Ctx, key).Result()
	if err != nil {
		http.Error(w, "Error fetching data from Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the value from Redis
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Value from Redis: %s", val)))
}
