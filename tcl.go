package tcl

import (
	"errors"
	"github.com/texas-holdem/tcl/proto"
)

var FullDeck []*texasholdem.Card
func GetShuffledDeck() *texasholdem.Cards {
	cards := make([]*texasholdem.Card, len(FullDeck))
	copy(cards, FullDeck)
	for i := range cards {
		j := Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return &texasholdem.Cards{
		Cards: cards,
	}
}

func computeRoyalFlushScore(hand *texasholdem.Cards) *texasholdem.Score {
	if straightFlushScore := computeStraightFlushScore(hand); nil != straightFlushScore &&
		straightFlushScore.GetKicker1() == texasholdem.Rank_ACE {
		return &texasholdem.Score{
			Category: texasholdem.Score_ROYAL_FLUSH,
		}
	}
	return nil
}

func computeStraightFlushScore(hand *texasholdem.Cards) *texasholdem.Score {
	if nil == computeFlushScore(hand) ||
		nil == computeStraightScore(hand) {
		return nil
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_STRAIGHT_FLUSH,
		Kicker1:  hand.GetCards()[4].GetRank(),
	}
}

func computeFourOfAKindScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	rank := cards[1].GetRank()
	if rank == cards[2].GetRank() &&
		rank == cards[3].GetRank() {
		if rank == cards[0].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FOUR_OF_A_KIND,
				Kicker1:  rank,
				Kicker2:  cards[4].GetRank(),
			}
		} else if rank == cards[4].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FOUR_OF_A_KIND,
				Kicker1:  rank,
				Kicker2:  cards[0].GetRank(),
			}
		}
	}
	return nil
}

func computeFullHouseScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	if cards[0].GetRank() == cards[1].GetRank() &&
		cards[3].GetRank() == cards[4].GetRank() {
		if cards[1].GetRank() == cards[2].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FULL_HOUSE,
				Kicker1:  cards[0].GetRank(),
				Kicker2:  cards[3].GetRank(),
			}
		} else if cards[2].GetRank() == cards[3].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FULL_HOUSE,
				Kicker1:  cards[2].GetRank(),
				Kicker2:  cards[0].GetRank(),
			}
		}
	}
	return nil
}

func computeFlushScore(hand *texasholdem.Cards) *texasholdem.Score {
	suit := hand.GetCards()[0].GetSuit()
	for i := 1; i < 5; i++ {
		if hand.GetCards()[i].GetSuit() != suit {
			return nil
		}
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_FLUSH,
		Kicker1:  hand.GetCards()[4].GetRank(),
	}
}

var StraightSpecialCase = []texasholdem.Rank{texasholdem.Rank_TWO,
	texasholdem.Rank_THREE,
	texasholdem.Rank_FOUR,
	texasholdem.Rank_FIVE,
	texasholdem.Rank_ACE}

func computeStraightScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	isSpecialCase := true
	for i := 0; i < 5; i++ {
		if cards[i].GetRank() != StraightSpecialCase[i] {
			isSpecialCase = false
			break
		}
	}
	if isSpecialCase {
		return &texasholdem.Score{
			Kicker1: texasholdem.Rank_FIVE,
		}
	}
	rank := cards[0].GetRank()
	for i := 1; i < 5; i++ {
		currentRank := cards[i].GetRank()
		if int(currentRank) != int(rank)+i {
			return nil
		}
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_STRAIGHT,
		Kicker1:  cards[4].GetRank(),
	}
}

func computeThreeOfAKindScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	if cards[0].GetRank() == cards[1].GetRank() &&
		cards[1].GetRank() == cards[2].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker1:  cards[0].GetRank(),
			Kicker2:  cards[4].GetRank(),
			Kicker3:  cards[3].GetRank(),
		}
	} else if cards[1].GetRank() == cards[2].GetRank() &&
		cards[2].GetRank() == cards[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker1:  cards[1].GetRank(),
			Kicker2:  cards[4].GetRank(),
			Kicker3:  cards[0].GetRank(),
		}
	} else if cards[2].GetRank() == cards[3].GetRank() &&
		cards[3].GetRank() == cards[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker1:  cards[2].GetRank(),
			Kicker2:  cards[1].GetRank(),
			Kicker3:  cards[0].GetRank(),
		}
	}
	return nil
}

func computeTwoPairsScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	if cards[0].GetRank() == cards[1].GetRank() &&
		cards[2].GetRank() == cards[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker1:  cards[2].GetRank(),
			Kicker2:  cards[0].GetRank(),
			Kicker3:  cards[4].GetRank(),
		}
	} else if cards[0].GetRank() == cards[1].GetRank() &&
		cards[3].GetRank() == cards[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker1:  cards[3].GetRank(),
			Kicker2:  cards[0].GetRank(),
			Kicker3:  cards[2].GetRank(),
		}
	} else if cards[1].GetRank() == cards[2].GetRank() &&
		cards[3].GetRank() == cards[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker1:  cards[3].GetRank(),
			Kicker2:  cards[1].GetRank(),
			Kicker3:  cards[0].GetRank(),
		}
	}
	return nil
}

func computePairScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	if cards[0].GetRank() == cards[1].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker1:  cards[0].GetRank(),
			Kicker2:  cards[4].GetRank(),
			Kicker3:  cards[3].GetRank(),
			Kicker4:  cards[2].GetRank(),
		}
	} else if cards[1].GetRank() == cards[2].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker1:  cards[1].GetRank(),
			Kicker2:  cards[4].GetRank(),
			Kicker3:  cards[3].GetRank(),
			Kicker4:  cards[0].GetRank(),
		}
	} else if cards[2].GetRank() == cards[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker1:  cards[2].GetRank(),
			Kicker2:  cards[4].GetRank(),
			Kicker3:  cards[1].GetRank(),
			Kicker4:  cards[0].GetRank(),
		}
	}
	if cards[3].GetRank() == cards[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker1:  cards[3].GetRank(),
			Kicker2:  cards[2].GetRank(),
			Kicker3:  cards[1].GetRank(),
			Kicker4:  cards[0].GetRank(),
		}
	}
	return nil
}

func computeHighCardScore(hand *texasholdem.Cards) *texasholdem.Score {
	cards := hand.GetCards()
	return &texasholdem.Score{
		Category: texasholdem.Score_HIGH_CARD,
		Kicker1:  cards[4].GetRank(),
		Kicker2:  cards[3].GetRank(),
		Kicker3:  cards[2].GetRank(),
		Kicker4:  cards[1].GetRank(),
		Kicker5:  cards[0].GetRank(),
	}
}

func ComputeScore(hand *texasholdem.Cards) (*texasholdem.Score, error) {
	if nil == hand || 5 != len(hand.GetCards()) {
		return nil, errors.New("invalid hand")
	}
	computers := []func(cards *texasholdem.Cards)*texasholdem.Score{
		computeRoyalFlushScore,
		computeStraightFlushScore,
		computeFourOfAKindScore,
		computeFullHouseScore,
		computeFlushScore,
		computeStraightScore,
		computeThreeOfAKindScore,
		computeTwoPairsScore,
		computePairScore,
		computeHighCardScore,
	}
	hand = SortedCards(hand)
	for _, computer := range computers {
		if score := computer(hand); nil != score {
			return score, nil
		}
	}
	panic("should not reach")
}

func init() {
	FullDeck = make([]*texasholdem.Card, 52)
	idx := 0
	for rank := texasholdem.Rank_TWO; rank <= texasholdem.Rank_ACE; rank++ {
		for suit := texasholdem.Suit_SPADE; suit <= texasholdem.Suit_DIAMOND; suit++ {
			FullDeck[idx] = &texasholdem.Card{
				Rank: rank,
				Suit: suit,
			}
			idx++
		}
	}
}
