package main

import (
	"fmt"
	"sort"
)

type MatchInterface interface {
	goalsSort(players []Player) []Player
	ratingSort(players []Player) []Player
	gmSort(players []Player) []Player
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		playerI := players[i]
		playerJ := players[j]

		goalI := playerI.Goals
		goalJ := playerJ.Goals

		if goalI != goalJ {
			return goalI > goalJ
		}

		raitI := playerI.Rating
		raitJ := playerJ.Rating

		if raitI != raitJ {
			return raitI > raitJ
		}

		return playerI.Name < playerJ.Name
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		playerI := players[i]
		playerJ := players[j]

		raitI := playerI.Rating
		raitJ := playerJ.Rating

		if raitI != raitJ {
			return raitI > raitJ
		}
		return playerI.Name < playerJ.Name
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		playerI := players[i]
		playerJ := players[j]

		goalI := playerI.Goals
		goalJ := playerJ.Goals

		missI := playerI.Misses
		missJ := playerJ.Misses
		var gm1, gm2 float64
		if missI != 0 {
			gm1 = float64(goalI) / float64(missI)
		} else {
			gm1 = float64(goalI) / 2
		}

		if missJ != 0 {
			gm2 = float64(goalJ) / float64(missJ)
		} else {
			gm2 = float64(goalJ) / 2
		}
		if gm1 != gm2 {
			return gm1 > gm2
		}

		raitI := playerI.Rating
		raitJ := playerJ.Rating

		if raitI != raitJ {
			return raitI > raitJ
		}
		return playerI.Name < playerJ.Name
	})
	return players
}

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	pl := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	pl.calculateRating()
	return pl
}

func (s *Player) calculateRating() {
	numerator := float64(s.Goals) + float64(s.Assists)/2

	if s.Misses != 0 {
		s.Rating = numerator / float64(s.Misses)
	} else {
		s.Rating = numerator
	}
}

func main() {
	fmt.Println("dkgndgdgnkdg")
}
