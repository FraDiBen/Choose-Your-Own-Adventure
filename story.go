package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Reader struct {
	filename string
	logger   *log.Logger
}

func (r Reader) Read() Story {
	fp, err := os.Open(r.filename)
	if err != nil {
		r.logger.Println("can't open filename ", r.filename)
		return nil
	}
	stories := Story{}
	err = json.NewDecoder(fp).Decode(&stories)
	if err != nil {
		r.logger.Println("can't decode json file", r.filename)
		return nil
	}
	r.logger.Printf("file %s correctly parsed\n", r.filename)
	return stories
}

type Story map[string]Arc

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string     `json:"title"`
	Story   []string   `json:"story"`
	Options []*option  `json:"options,omitempty"`
}

func (a Arc) StoryParagraph() string {
	return strings.Join(a.Story, "\n\n")
}
