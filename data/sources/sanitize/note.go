package sanitize

import (
	"strings"
)

func AppendNewNote(currentNotes []string, newNotes ...string) []string {
	for _, newNote := range newNotes {
		newNote = strings.TrimSpace(newNote)
		if len(newNote) == 0 {
			continue
		}
		alreadyExists := false
		for _, currentNote := range currentNotes {
			if strings.Contains(currentNote, newNote) {
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			currentNotes = append(currentNotes, newNote)
		}
	}

	return currentNotes
}
