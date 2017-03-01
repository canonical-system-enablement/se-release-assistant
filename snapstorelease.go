package main

import (
	"fmt"
	"github.com/bergotorino/go-trello"
	"strings"
)

// SE Sprint board ID
const sprintBoardId string = "5722667126680b7e86626557"

const mpChecklistName string = "MPs"
const relatedStoriesChecklistName string = "Stories"

type SnapsToRelease struct {
	board trello.Board
}

func (t *SeClient) SnapsToRelease() (*SnapsToRelease, error) {
	board, err := t.trello.Board(sprintBoardId)
	if err != nil {
		return nil, err
	}
	return &SnapsToRelease{board: *board}, nil
}

func (s *SnapsToRelease) MergeProposals() ([]string, error) {

	var retval []string

	releaseCards, err := s.cardsOfSnapsToRelease()
	if err != nil {
		return nil, err
	}

	// The MPs are tracked on a checklist called "MPs" for each card
	// that is being worked on during the sprint. Additionally the MPs
	// might be tracked on the release card too on a checklist with the
	// "MPs" name.

	// For each card of a snap to release find out related stories and
	// for each of this story find out the MPs. Also try to fetch MPs
	// from the release card directly
	var relatedStories []string
	for _, rc := range releaseCards {

		// Get the checklists
		checklists, err := rc.Checklists()
		if err != nil {
			return nil, err
		}

		// For each checklist on the release card
		for _, checklist := range checklists {
			// "MPs"
			if checklist.Name == mpChecklistName {
				// Each item on this checklist is a MP
				for _, citem := range checklist.CheckItems {
					retval = append(retval, citem.Name)
				}
				continue
			}

			// "Stories"
			if checklist.Name == relatedStoriesChecklistName {

				// Each item on this checklist is a link
				// to the card.
				// Lets save them and process in another step
				for _, citem := range checklist.CheckItems {
					relatedStories = append(relatedStories, citem.Name)
				}
				continue
			}
		}
	}

	// At this point we have MP fetched directly from the release cards
	// as well as a collection of related stories from which we need to
	// obtain the MPs. Sadly we do not have the story trello-like-id but
	// just an url (long or short) therefore we need to look at each card
	// on the board.

	// Get the lists
	lists, err := s.board.Lists()
	if err != nil {
		return nil, err
	}

	for _, l := range lists {
		// Find the correct swimlanes, skip "Snaps to Release" and
		// "Snaps Being Released" lanes.
		if strings.Contains(l.Name, "Snaps") {
			continue
		}

		cards, err := l.Cards()
		if err != nil {
			return nil, err
		}

		for _, c := range cards {
			// Get the checklists
			checklists, err := c.Checklists()
			if err != nil {
				return nil, err
			}

			// For each checklist on the card
			for _, checklist := range checklists {
				// "MPs"
				if checklist.Name == mpChecklistName {
					// Each item on this checklist is a MP
					for _, citem := range checklist.CheckItems {
						retval = append(retval, citem.Name)
					}
					continue
				}
			}
		}
	}

	return unique(retval), nil
}

func (s *SnapsToRelease) cardsOfSnapsToRelease() (cards []trello.Card, err error) {
	// Get the lists
	lists, err := s.board.Lists()
	if err != nil {
		return nil, err
	}

	for _, l := range lists {

		// Find the correct swimlane
		if !strings.Contains(l.Name, "Snaps to Release") {
			continue
		}

		cards, err := l.Cards()
		if err != nil {
			return nil, err
		}
		for i, c := range cards {
			// Remove the README card
			if strings.Contains(c.Name, "How to use this column") {
				cards = cards[:i+copy(cards[i:], cards[i+1:])]
			}
		}
		return cards, nil
	}

	return nil, fmt.Errorf("Snap to Release swimlanie is no more")
}

func unique(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
