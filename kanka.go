package main

import "time"

type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  string      `json:"next"`
}

type PaginationLinks interface {
	PaginationLinks() Links
}

type Characters struct {
	Data []struct {
		ID                   int           `json:"id"`
		Name                 string        `json:"name"`
		Entry                string        `json:"entry"`
		EntryParsed          string        `json:"entry_parsed"`
		Image                string        `json:"image"`
		ImageFull            string        `json:"image_full"`
		ImageThumb           string        `json:"image_thumb"`
		HasCustomImage       bool          `json:"has_custom_image"`
		IsPrivate            bool          `json:"is_private"`
		IsTemplate           bool          `json:"is_template"`
		EntityID             int           `json:"entity_id"`
		Tags                 []interface{} `json:"tags"`
		CreatedAt            time.Time     `json:"created_at"`
		CreatedBy            int           `json:"created_by"`
		UpdatedAt            time.Time     `json:"updated_at"`
		UpdatedBy            int           `json:"updated_by"`
		LocationID           int           `json:"location_id"`
		Title                string        `json:"title"`
		Age                  string        `json:"age"`
		Sex                  string        `json:"sex"`
		Pronouns             interface{}   `json:"pronouns"`
		RaceID               int           `json:"race_id"`
		Type                 string        `json:"type"`
		FamilyID             int           `json:"family_id"`
		IsDead               bool          `json:"is_dead"`
		Traits               []interface{} `json:"traits"`
		IsPersonalityVisible bool          `json:"is_personality_visible"`
	} `json:"data"`
	Links `json:"links"`
	Meta  struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
	Sync time.Time `json:"sync"`
}

func (c Characters) PaginationLinks() Links {
	return c.Links
}
