package tcl

import (
	"errors"
	"github.com/texas-holdem/tcl/proto"
)

var FullDeck []*texasholdem.Card
func GetShuffledDeck() []*texasholdem.Card {
	cards := make([]*texasholdem.Card, len(FullDeck))
	copy(cards, FullDeck)
	for i := range cards {
		j := Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func computeRoyalFlushScore(hand []*texasholdem.Card) *texasholdem.Score {
	if straightFlushScore := computeStraightFlushScore(hand); nil != straightFlushScore &&
		straightFlushScore.GetKicker()[0] == texasholdem.Rank_ACE {
		return &texasholdem.Score{
			Category: texasholdem.Score_ROYAL_FLUSH,
		}
	}
	return nil
}

func computeStraightFlushScore(hand []*texasholdem.Card) *texasholdem.Score {
	if nil == computeFlushScore(hand) ||
		nil == computeStraightScore(hand) {
		return nil
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_STRAIGHT_FLUSH,
		Kicker:  []texasholdem.Rank{
			hand[4].GetRank(),
		},
	}
}

func computeFourOfAKindScore(hand []*texasholdem.Card) *texasholdem.Score {
	rank := hand[1].GetRank()
	if rank == hand[2].GetRank() &&
		rank == hand[3].GetRank() {
		if rank == hand[0].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FOUR_OF_A_KIND,
				Kicker: []texasholdem.Rank{
					rank,
					hand[4].GetRank(),
				},
			}
		} else if rank == hand[4].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FOUR_OF_A_KIND,
				Kicker: []texasholdem.Rank{
					rank,
					hand[0].GetRank(),
				},
			}
		}
	}
	return nil
}

func computeFullHouseScore(hand []*texasholdem.Card) *texasholdem.Score {
	if hand[0].GetRank() == hand[1].GetRank() &&
		hand[3].GetRank() == hand[4].GetRank() {
		if hand[1].GetRank() == hand[2].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FULL_HOUSE,
				Kicker: []texasholdem.Rank{
					hand[0].GetRank(),
					hand[3].GetRank(),
				},
			}
		} else if hand[2].GetRank() == hand[3].GetRank() {
			return &texasholdem.Score{
				Category: texasholdem.Score_FULL_HOUSE,
				Kicker: []texasholdem.Rank{
					hand[2].GetRank(),
					hand[0].GetRank(),
				},
			}
		}
	}
	return nil
}

func computeFlushScore(hand []*texasholdem.Card) *texasholdem.Score {
	suit := hand[0].GetSuit()
	for i := 1; i < 5; i++ {
		if hand[i].GetSuit() != suit {
			return nil
		}
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_FLUSH,
		Kicker: []texasholdem.Rank{
			hand[4].GetRank(),
		},
	}
}

var StraightSpecialCase = []texasholdem.Rank{texasholdem.Rank_TWO,
	texasholdem.Rank_THREE,
	texasholdem.Rank_FOUR,
	texasholdem.Rank_FIVE,
	texasholdem.Rank_ACE}

func computeStraightScore(hand []*texasholdem.Card) *texasholdem.Score {
	isSpecialCase := true
	for i := 0; i < 5; i++ {
		if hand[i].GetRank() != StraightSpecialCase[i] {
			isSpecialCase = false
			break
		}
	}
	if isSpecialCase {
		return &texasholdem.Score{
			Category: texasholdem.Score_STRAIGHT,
			Kicker: []texasholdem.Rank{
				texasholdem.Rank_FIVE,
			},
		}
	}
	rank := hand[0].GetRank()
	for i := 1; i < 5; i++ {
		currentRank := hand[i].GetRank()
		if int(currentRank) != int(rank)+i {
			return nil
		}
	}
	return &texasholdem.Score{
		Category: texasholdem.Score_STRAIGHT,
		Kicker: []texasholdem.Rank{
			hand[4].GetRank(),
		},
	}
}

func computeThreeOfAKindScore(hand []*texasholdem.Card) *texasholdem.Score {
	if hand[0].GetRank() == hand[1].GetRank() &&
		hand[1].GetRank() == hand[2].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker: []texasholdem.Rank{
				hand[0].GetRank(),
				hand[4].GetRank(),
				hand[3].GetRank(),
			},
		}
	} else if hand[1].GetRank() == hand[2].GetRank() &&
		hand[2].GetRank() == hand[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker: []texasholdem.Rank{
				hand[1].GetRank(),
				hand[4].GetRank(),
				hand[0].GetRank(),
			},
		}
	} else if hand[2].GetRank() == hand[3].GetRank() &&
		hand[3].GetRank() == hand[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker: []texasholdem.Rank{
				hand[2].GetRank(),
				hand[1].GetRank(),
				hand[0].GetRank(),
			},
		}
	}
	return nil
}

func computeTwoPairsScore(hand []*texasholdem.Card) *texasholdem.Score {
	if hand[0].GetRank() == hand[1].GetRank() &&
		hand[2].GetRank() == hand[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker: []texasholdem.Rank{
				hand[2].GetRank(),
				hand[0].GetRank(),
				hand[4].GetRank(),
			},
		}
	} else if hand[0].GetRank() == hand[1].GetRank() &&
		hand[3].GetRank() == hand[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker: []texasholdem.Rank{
				hand[3].GetRank(),
				hand[0].GetRank(),
				hand[2].GetRank(),
			},
		}
	} else if hand[1].GetRank() == hand[2].GetRank() &&
		hand[3].GetRank() == hand[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker: []texasholdem.Rank{
				hand[3].GetRank(),
				hand[1].GetRank(),
				hand[0].GetRank(),
			},
		}
	}
	return nil
}

func computePairScore(hand []*texasholdem.Card) *texasholdem.Score {
	if hand[0].GetRank() == hand[1].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker: []texasholdem.Rank{
				hand[0].GetRank(),
				hand[4].GetRank(),
				hand[3].GetRank(),
				hand[2].GetRank(),
			},
		}
	} else if hand[1].GetRank() == hand[2].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker: []texasholdem.Rank{
				hand[1].GetRank(),
				hand[4].GetRank(),
				hand[3].GetRank(),
				hand[0].GetRank(),
			},
		}
	} else if hand[2].GetRank() == hand[3].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker: []texasholdem.Rank{
				hand[2].GetRank(),
				hand[4].GetRank(),
				hand[1].GetRank(),
				hand[0].GetRank(),
			},
		}
	}
	if hand[3].GetRank() == hand[4].GetRank() {
		return &texasholdem.Score{
			Category: texasholdem.Score_PAIR,
			Kicker: []texasholdem.Rank{
				hand[3].GetRank(),
				hand[2].GetRank(),
				hand[1].GetRank(),
				hand[0].GetRank(),
			},
		}
	}
	return nil
}

func computeHighCardScore(hand []*texasholdem.Card) *texasholdem.Score {
	return &texasholdem.Score{
		Category: texasholdem.Score_HIGH_CARD,
		Kicker: []texasholdem.Rank{
			hand[4].GetRank(),
			hand[3].GetRank(),
			hand[2].GetRank(),
			hand[1].GetRank(),
			hand[0].GetRank(),
		},
	}
}

func ComputeScore(hand []*texasholdem.Card) (*texasholdem.Score, error) {
	if nil == hand || 5 != len(hand) {
		return nil, errors.New("invalid hand")
	}
	computers := []func(cards []*texasholdem.Card)*texasholdem.Score{
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

func ComputeHighestScore(cards []*texasholdem.Card) (*texasholdem.Score, error) {
	n := len(cards)
	if n < 5 {
		return nil, errors.New("can not compute score with less than 5 cards")
	}
	var maxScore *texasholdem.Score
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for p := k + 1; p < n; p++ {
					for q := p + 1; q < n; q++ {
						hand := []*texasholdem.Card{
							cards[i],
							cards[j],
							cards[k],
							cards[p],
							cards[q],
						}
						currentScore, _ := ComputeScore(hand)
						if compareScore(currentScore, maxScore) > 0 {
							maxScore = currentScore
						}
					}
				}
			}
		}
	}
	return maxScore, nil
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
