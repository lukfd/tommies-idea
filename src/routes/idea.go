package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lukfd/redis-demo/internal"
)

func Idea(w http.ResponseWriter, r *http.Request, redis *internal.RedisClient) {
	id := r.PathValue("id")
	key := fmt.Sprintf("idea:%s", id)
	log.Println("Key: " + key)
	res, err := redis.Client.HGetAll(redis.Ctx, key).Result()
	if err != nil {
		http.Error(w, "Error fetching data from Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the map to a string
	var builder strings.Builder
	for k, v := range res {
		builder.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}

	// Respond with the formatted string
	log.Println("Result: " + builder.String())
	log.Println("Original: " + strconv.Itoa(len(res)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(builder.String()))
}
