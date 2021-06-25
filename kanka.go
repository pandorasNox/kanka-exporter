package main

import "time"

type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  string      `json:"next"`
}

type Meta struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	LastPage    int    `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

type MetaInfo interface {
	PaginationLinks() Links
	MetaInfo() Meta
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
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (c Characters) PaginationLinks() Links {
	return c.Links
}

func (c Characters) MetaInfo() Meta {
	return c.Meta
}

type Entities struct {
	Data []struct {
		ID                  int           `json:"id"`
		Name                string        `json:"name"`
		Type                string        `json:"type"`
		ChildID             int           `json:"child_id"`
		Tags                []interface{} `json:"tags"`
		IsPrivate           bool          `json:"is_private"`
		IsTemplate          bool          `json:"is_template"`
		CampaignID          int           `json:"campaign_id"`
		IsAttributesPrivate bool          `json:"is_attributes_private"`
		Tooltip             string        `json:"tooltip"`
		HeaderImage         interface{}   `json:"header_image"`
		ImageUUID           interface{}   `json:"image_uuid"`
		CreatedAt           time.Time     `json:"created_at"`
		CreatedBy           int           `json:"created_by"`
		UpdatedAt           time.Time     `json:"updated_at"`
		UpdatedBy           int           `json:"updated_by"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (e Entities) PaginationLinks() Links {
	return e.Links
}

func (e Entities) MetaInfo() Meta {
	return e.Meta
}

type Locations struct {
	Data []struct {
		ID               int           `json:"id"`
		Name             string        `json:"name"`
		Entry            string        `json:"entry"`
		EntryParsed      string        `json:"entry_parsed"`
		Image            interface{}   `json:"image"`
		ImageFull        string        `json:"image_full"`
		ImageThumb       string        `json:"image_thumb"`
		HasCustomImage   bool          `json:"has_custom_image"`
		IsPrivate        bool          `json:"is_private"`
		IsTemplate       bool          `json:"is_template"`
		EntityID         int           `json:"entity_id"`
		Tags             []interface{} `json:"tags"`
		CreatedAt        time.Time     `json:"created_at"`
		CreatedBy        int           `json:"created_by"`
		UpdatedAt        time.Time     `json:"updated_at"`
		UpdatedBy        int           `json:"updated_by"`
		Type             string        `json:"type"`
		Map              string        `json:"map"`
		IsMapPrivate     int           `json:"is_map_private"`
		ParentLocationID int           `json:"parent_location_id"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (l Locations) PaginationLinks() Links {
	return l.Links
}

func (l Locations) MetaInfo() Meta {
	return l.Meta
}

type Families struct {
	Data []struct {
		ID             int           `json:"id"`
		Name           string        `json:"name"`
		Entry          string        `json:"entry"`
		EntryParsed    string        `json:"entry_parsed"`
		Image          interface{}   `json:"image"`
		ImageFull      string        `json:"image_full"`
		ImageThumb     string        `json:"image_thumb"`
		HasCustomImage bool          `json:"has_custom_image"`
		IsPrivate      bool          `json:"is_private"`
		IsTemplate     bool          `json:"is_template"`
		EntityID       int           `json:"entity_id"`
		Tags           []interface{} `json:"tags"`
		CreatedAt      time.Time     `json:"created_at"`
		CreatedBy      int           `json:"created_by"`
		UpdatedAt      time.Time     `json:"updated_at"`
		UpdatedBy      int           `json:"updated_by"`
		LocationID     int           `json:"location_id"`
		Type           string        `json:"type"`
		FamilyID       interface{}   `json:"family_id"`
		Members        []int         `json:"members"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (f Families) PaginationLinks() Links {
	return f.Links
}

func (f Families) MetaInfo() Meta {
	return f.Meta
}

type Organisations struct {
	Data []struct {
		ID             int           `json:"id"`
		Name           string        `json:"name"`
		Entry          string        `json:"entry"`
		EntryParsed    string        `json:"entry_parsed"`
		Image          interface{}   `json:"image"`
		ImageFull      string        `json:"image_full"`
		ImageThumb     string        `json:"image_thumb"`
		HasCustomImage bool          `json:"has_custom_image"`
		IsPrivate      bool          `json:"is_private"`
		IsTemplate     bool          `json:"is_template"`
		EntityID       int           `json:"entity_id"`
		Tags           []interface{} `json:"tags"`
		CreatedAt      time.Time     `json:"created_at"`
		CreatedBy      int           `json:"created_by"`
		UpdatedAt      time.Time     `json:"updated_at"`
		UpdatedBy      int           `json:"updated_by"`
		LocationID     int           `json:"location_id"`
		Type           string        `json:"type"`
		OrganisationID int           `json:"organisation_id"`
		Members        []struct {
			CharacterID    int         `json:"character_id"`
			CreatedAt      time.Time   `json:"created_at"`
			CreatedBy      interface{} `json:"created_by"`
			ID             int         `json:"id"`
			IsPrivate      bool        `json:"is_private"`
			OrganisationID int         `json:"organisation_id"`
			Role           interface{} `json:"role"`
			UpdatedAt      time.Time   `json:"updated_at"`
			UpdatedBy      interface{} `json:"updated_by"`
		} `json:"members"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (o Organisations) PaginationLinks() Links {
	return o.Links
}

func (o Organisations) MetaInfo() Meta {
	return o.Meta
}

type Items struct {
	Data []struct {
		ID             int           `json:"id"`
		Name           string        `json:"name"`
		Entry          string        `json:"entry"`
		EntryParsed    string        `json:"entry_parsed"`
		Image          interface{}   `json:"image"`
		ImageFull      string        `json:"image_full"`
		ImageThumb     string        `json:"image_thumb"`
		HasCustomImage bool          `json:"has_custom_image"`
		IsPrivate      bool          `json:"is_private"`
		IsTemplate     bool          `json:"is_template"`
		EntityID       int           `json:"entity_id"`
		Tags           []interface{} `json:"tags"`
		CreatedAt      time.Time     `json:"created_at"`
		CreatedBy      int           `json:"created_by"`
		UpdatedAt      time.Time     `json:"updated_at"`
		UpdatedBy      int           `json:"updated_by"`
		LocationID     int           `json:"location_id"`
		CharacterID    interface{}   `json:"character_id"`
		Type           string        `json:"type"`
		Price          interface{}   `json:"price"`
		Size           string        `json:"size"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (i Items) PaginationLinks() Links {
	return i.Links
}

func (i Items) MetaInfo() Meta {
	return i.Meta
}

type Notes struct {
	Data []struct {
		ID             int         `json:"id"`
		Name           string      `json:"name"`
		Entry          string      `json:"entry"`
		EntryParsed    string      `json:"entry_parsed"`
		Image          string      `json:"image"`
		ImageFull      string      `json:"image_full"`
		ImageThumb     string      `json:"image_thumb"`
		HasCustomImage bool        `json:"has_custom_image"`
		IsPrivate      bool        `json:"is_private"`
		IsTemplate     bool        `json:"is_template"`
		EntityID       int         `json:"entity_id"`
		Tags           []int       `json:"tags"`
		CreatedAt      time.Time   `json:"created_at"`
		CreatedBy      int         `json:"created_by"`
		UpdatedAt      time.Time   `json:"updated_at"`
		UpdatedBy      int         `json:"updated_by"`
		Type           string      `json:"type"`
		NoteID         interface{} `json:"note_id"`
		IsPinned       int         `json:"is_pinned"`
	} `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
	Sync  time.Time `json:"sync"`
}

func (n Notes) PaginationLinks() Links {
	return n.Links
}

func (n Notes) MetaInfo() Meta {
	return n.Meta
}
