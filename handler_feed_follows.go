package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/imim77/rssag/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	parm := params{}
	err := decoder.Decode(&parm)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    parm.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
	}
	respondWithJSON(w, 200, databaseFeedsFollowsToFeedsFollows(feedfollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "feedFollowID")
	feedfollowid := uuid.MustParse(feedFollowIDstr)
	err := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting the feed follow: %v", err))
	}
	respondWithJSON(w, 200, struct{}{})
}
