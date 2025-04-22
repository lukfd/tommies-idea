package model

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Idea struct {
	ID          int      `json:"id"`
	Timestamp   string   `json:"timestamp"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Writer      string   `json:"writer"`
	Tags        []string `json:"tags"`
}

func (r *RedisClient) GetBytes(idea Idea) ([]byte, error) {
	return json.Marshal(idea)
}

func (r *RedisClient) AddIdea(idea Idea) {
	key := fmt.Sprintf("idea:%d", idea.ID)
	timestamp := time.Now().Unix()
	tags := strings.Join(idea.Tags, ",")

	_, err := r.Client.HSet(r.Ctx, key, "timestamp", timestamp, "title", idea.Title, "description", idea.Description, "writer", idea.Writer, "tags", tags).Result()
	if err != nil {
		log.Printf("Error adding idea to Redis: %v", err)
	} else {
		log.Printf("Successfully added idea with ID %d to Redis", idea.ID)
	}
}

func (r *RedisClient) GetIdea(id int) (Idea, error) {
	key := fmt.Sprintf("idea:%d", id)
	res, err := r.Client.HGetAll(r.Ctx, key).Result()
	if err != nil {
		return Idea{}, fmt.Errorf("error fetching idea from Redis: %v", err)
	}

	if len(res) == 0 {
		return Idea{}, fmt.Errorf("no idea found with ID %d", id)
	}

	tags := strings.Split(res["tags"], ",")
	idea := Idea{
		ID:          id,
		Timestamp:   res["timestamp"],
		Title:       res["title"],
		Description: res["description"],
		Writer:      res["writer"],
		Tags:        tags,
	}

	return idea, nil
}

func (r *RedisClient) GetIdeas() ([]Idea, error) {
	keys, err := r.Client.Keys(r.Ctx, "idea:*").Result()
	if err != nil {
		return nil, fmt.Errorf("error fetching ideas from Redis: %v", err)
	}

	var ideas []Idea
	for _, key := range keys {
		res, err := r.Client.HGetAll(r.Ctx, key).Result()
		if err != nil {
			log.Printf("Error fetching idea from Redis: %v", err)
			continue
		}

		tags := strings.Split(res["tags"], ",")
		idStr := strings.TrimPrefix(key, "idea:")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Error converting id to integer: %v", err)
			continue
		}
		idea := Idea{
			ID:          id,
			Timestamp:   res["timestamp"],
			Title:       res["title"],
			Description: res["description"],
			Writer:      res["writer"],
			Tags:        tags,
		}
		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (r *RedisClient) DeleteIdea(id int) error {
	key := fmt.Sprintf("idea:%d", id)
	_, err := r.Client.Del(r.Ctx, key).Result()
	if err != nil {
		return fmt.Errorf("error deleting idea from Redis: %v", err)
	}

	log.Printf("Successfully deleted idea with ID %d from Redis", id)
	return nil
}
