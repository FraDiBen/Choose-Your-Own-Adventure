package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func makeReaderForTesting(fn string) (Reader, *bytes.Buffer) {
	lines := ""
	buffer := bytes.NewBufferString(lines)
	l := log.New(buffer, "", log.Llongfile)
	return Reader{
		filename: fn,
		logger:   l,
	}, buffer
}

func TestReader_Read(t *testing.T) {

	// todo: what about a more meaningful test for the reader?
	t.Run("read -- happy path", func(t *testing.T) {
		r, b := makeReaderForTesting("gopher.json")
		stories := r.Read()
		assert.Len(t, stories, 7)
		assert.Equal(t, stories["sean-kelly"], Arc{
			Title: "Exit Stage Left",
			Story: []string{
				"As you begin walking up to the fox-man you hear him introduce himself as Sean Kelly. While waiting in line you decide to do a little research to see what types of work Sean is into.",
				"A few clicks later and you drop your phone in horror. This guy's online handle is \"StabbyCutyou\". The stories about New York being dangerous were true!",
				"Without a thought you grab your gopher buddy and head for the door. \"I'll explain when we get to the hotel\" you tell him.",
				"After arriving at your hotel you both decide that you have had enough adventure. First thing tomorrow morning you are heading home.",
			},
			Options: []*option{
				{
					Text: "You change your flight to leave early and head to the airport in the morning.",
					Arc:  "home",
				},
			},
		})
		fmt.Println(b.String())
	})

	//todo: add all sad paths
}
