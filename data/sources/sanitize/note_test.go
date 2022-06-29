package sanitize_test

import (
	"github.com/analogj/justvanish/data/sources/sanitize"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestAppendNewNote(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		currentNotes []string
		newNote      string
		resultNotes  []string
	}{
		{currentNotes: []string{}, newNote: "", resultNotes: []string{}},               //empty string adds nothing
		{currentNotes: []string{"hello"}, newNote: "", resultNotes: []string{"hello"}}, //empty string adds nothing
		{currentNotes: []string{}, newNote: "hello", resultNotes: []string{"hello"}},
		{currentNotes: []string{}, newNote: "hello    ", resultNotes: []string{"hello"}},
		{currentNotes: []string{"hello"}, newNote: "hello", resultNotes: []string{"hello"}},                                           //dont add if already exists
		{currentNotes: []string{"hello with spaces"}, newNote: "hello with spaces      ", resultNotes: []string{"hello with spaces"}}, //dont add if already exists
	}

	//test
	for _, r := range table {
		t.Run(r.newNote, func(t *testing.T) {
			require.Equal(t, r.resultNotes, sanitize.AppendNewNote(r.currentNotes, r.newNote))
		})
	}
}

func TestAppendNewNote_WithMultiple(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		currentNotes []string
		newNote1     string
		newNote2     string
		resultNotes  []string
	}{
		{currentNotes: []string{}, newNote1: "", newNote2: "", resultNotes: []string{}},               //empty string adds nothing
		{currentNotes: []string{}, newNote1: "a", newNote2: "b", resultNotes: []string{"a", "b"}},     //empty string adds nothing
		{currentNotes: []string{"a"}, newNote1: "a   ", newNote2: "   a", resultNotes: []string{"a"}}, //trim
		{currentNotes: []string{"a"}, newNote1: "", newNote2: "   thisisnew", resultNotes: []string{"a", "thisisnew"}},
	}

	//test
	for _, r := range table {
		t.Run(r.newNote1+r.newNote2, func(t *testing.T) {
			require.Equal(t, r.resultNotes, sanitize.AppendNewNote(r.currentNotes, r.newNote1, r.newNote2))
		})
	}
}

func TestAppendNewNote_AfterMultipleCalls(t *testing.T) {
	t.Parallel()
	//setup
	table := []struct {
		currentNotes []string
		newNotes1    []string
		newNotes2    []string
		resultNotes  []string
	}{
		{currentNotes: []string{}, newNotes1: []string{"", ""}, newNotes2: []string{"", ""}, resultNotes: []string{}},
		{currentNotes: []string{"a"}, newNotes1: []string{"d", "c"}, newNotes2: []string{"d", "e"}, resultNotes: []string{"a", "d", "c", "e"}},
		{currentNotes: []string{"kwjeriwea"}, newNotes1: []string{"d", "c"}, newNotes2: []string{"d", "e"}, resultNotes: []string{"kwjeriwea", "d", "c"}},
	}

	//test
	for _, r := range table {
		t.Run(strings.Join(append(r.newNotes1, r.newNotes2...), ","), func(t *testing.T) {
			notes := sanitize.AppendNewNote(r.currentNotes, r.newNotes1...)
			notes = sanitize.AppendNewNote(notes, r.newNotes2...)
			require.Equal(t, r.resultNotes, notes)
		})
	}
}
