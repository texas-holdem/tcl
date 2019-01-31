package dealer

import (
	"errors"
	"github.com/texas-holdem/tcl"
	"github.com/texas-holdem/tcl/proto"
)

type Dealer struct {
	tableSize		int
	seatHasPlayer	[]bool
	deck			[]*texasholdem.Card
	playerCards		[][]*texasholdem.Card
	flopCards		[]*texasholdem.Card
	turnCard		*texasholdem.Card
	riverCard		*texasholdem.Card
	playerScores	[]*texasholdem.Score
}

func (dealer *Dealer) CreateTable(n int) {
	dealer.tableSize = n
	dealer.seatHasPlayer = make([]bool, n)

	tcl.UpdateRandomSeedWithCurrentTimestamp()
}

func (dealer *Dealer) AddPlayer(seatIndex int) error {
	if dealer.seatHasPlayer[seatIndex] {
		return errors.New("seat already has player")
	}
	dealer.seatHasPlayer[seatIndex] = true
	return nil
}

func (dealer *Dealer) RemovePlayer(seatIndex int) error {
	if !dealer.seatHasPlayer[seatIndex] {
		return errors.New("no player at that seat")
	}
	dealer.seatHasPlayer[seatIndex] = false
	return nil
}

func (dealer *Dealer) StartNewRound() {
	dealer.deck = tcl.GetShuffledDeck()
}

func (dealer *Dealer) NotifyPlayerRaise(playerSeat int, amount int) {

}

func (dealer *Dealer) NotifyPlayerFold(playerSeat int) {

}

func (dealer *Dealer) getCardsFromDeck(n int) []*texasholdem.Card {
	cards := dealer.deck[:n]
	dealer.deck = dealer.deck[n:]
	return cards
}

func (dealer *Dealer) DealCardsToPlayers() [][]*texasholdem.Card {
	dealer.playerCards = make([][]*texasholdem.Card, dealer.tableSize)
	for i := 0; i < dealer.tableSize; i++ {
		if dealer.seatHasPlayer[i] {
			dealer.playerCards[i] = dealer.getCardsFromDeck(2)
		}
	}
	return dealer.playerCards
}

func (dealer *Dealer) DealFlopCards() []*texasholdem.Card {
	dealer.getCardsFromDeck(1) // burn 1 card before dealing flop cards
	dealer.flopCards = dealer.getCardsFromDeck(3)
	return dealer.flopCards
}

func (dealer *Dealer) DealTurnCard() *texasholdem.Card {
	dealer.getCardsFromDeck(1) // burn 1 card before dealing turn cards
	dealer.turnCard = dealer.getCardsFromDeck(1)[0]
	return dealer.turnCard
}

func (dealer *Dealer) DealRiverCard() *texasholdem.Card {
	dealer.getCardsFromDeck(1) // burn 1 card before dealing river cards
	dealer.riverCard = dealer.getCardsFromDeck(1)[0]
	return dealer.riverCard
}

func (dealer *Dealer) ComputeScores() []*texasholdem.Score {
	communityCards := append(dealer.flopCards, dealer.turnCard, dealer.riverCard)
	dealer.playerScores = make([]*texasholdem.Score, dealer.tableSize)
	for i := 0; i < dealer.tableSize; i++ {
		if dealer.seatHasPlayer[i] {
			allCards := append(communityCards, dealer.playerCards[i]...)
			dealer.playerScores[i], _ = tcl.ComputeHighestScore(allCards)
		}
	}
	return dealer.playerScores
}

func (dealer *Dealer) NotifyRoundEnd() {

}
