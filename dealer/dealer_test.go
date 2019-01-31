package dealer

import (
	"github.com/texas-holdem/tcl"
	"testing"
)

func TestDealer_All(t *testing.T) {
	dealer := new(Dealer)
	dealer.CreateTable(10)
	dealer.AddPlayer(1)
	dealer.AddPlayer(3)
	dealer.AddPlayer(5)
	dealer.AddPlayer(7)
	for i := 0; i < 1000; i++ {
		t.Log("")
		t.Logf("round: %d", i)
		dealer.StartNewRound()
		playerCards := dealer.DealCardsToPlayers()
		playerCardsString := "player cards: { "
		for _, cards := range playerCards {
			playerCardsString += tcl.CardsDebugString(cards) + ", "
		}
		playerCardsString += " }"
		t.Log(playerCardsString)
		communityCards := dealer.DealFlopCards()
		communityCards = append(communityCards, dealer.DealTurnCard())
		communityCards = append(communityCards, dealer.DealRiverCard())
		t.Logf("commnity cards: %s", tcl.CardsDebugString(communityCards))
		playerScores := dealer.ComputeScores()
		for i, score := range playerScores {
			if len(playerCards[i]) == 2 {
				t.Logf("player %d score: %v", i, score)
			} else {
				t.Logf("player %d score: nil", i)
			}
		}
	}
}
