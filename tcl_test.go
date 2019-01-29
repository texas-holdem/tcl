package tcl

import (
	"github.com/golang/protobuf/proto"
	"github.com/texas-holdem/tcl/proto"
	"math"
	"testing"
)

func TestComputeTypicalScores(t *testing.T) {
	testCases := []*texasholdem.Cards{
		AbbrsToCards([]string{"th", "jh", "qh", "kh", "ah"}), // royal flush
		AbbrsToCards([]string{"th", "jh", "qh", "kh", "9h"}), // straight flush
		AbbrsToCards([]string{"ts", "th", "tc", "td", "ah"}), // four of a kind
		AbbrsToCards([]string{"th", "ts", "td", "kh", "ks"}), // full house
		AbbrsToCards([]string{"th", "2h", "4h", "6h", "8h"}), // flush
		AbbrsToCards([]string{"th", "9s", "8d", "6h", "7h"}), // straight
		AbbrsToCards([]string{"th", "2h", "4h", "2d", "2s"}), // three of a kind
		AbbrsToCards([]string{"th", "2h", "ts", "2d", "8h"}), // two pairs
		AbbrsToCards([]string{"th", "2h", "4h", "6h", "2s"}), // pair
		AbbrsToCards([]string{"th", "2d", "4s", "6h", "8h"}), // high card
	}
	expected := []*texasholdem.Score{
		{
			Category: texasholdem.Score_ROYAL_FLUSH,
		},
		{
			Category: texasholdem.Score_STRAIGHT_FLUSH,
			Kicker1:  texasholdem.Rank_KING,
		},
		{
			Category: texasholdem.Score_FOUR_OF_A_KIND,
			Kicker1:  texasholdem.Rank_TEN,
			Kicker2:  texasholdem.Rank_ACE,
		},
		{
			Category: texasholdem.Score_FULL_HOUSE,
			Kicker1:  texasholdem.Rank_TEN,
			Kicker2:  texasholdem.Rank_KING,
		},
		{
			Category: texasholdem.Score_FLUSH,
			Kicker1:  texasholdem.Rank_TEN,
		},
		{
			Category: texasholdem.Score_STRAIGHT,
			Kicker1:  texasholdem.Rank_TEN,
		},
		{
			Category: texasholdem.Score_THREE_OF_A_KIND,
			Kicker1:  texasholdem.Rank_TWO,
			Kicker2:  texasholdem.Rank_TEN,
			Kicker3:  texasholdem.Rank_FOUR,
		},
		{
			Category: texasholdem.Score_TWO_PAIRS,
			Kicker1:  texasholdem.Rank_TEN,
			Kicker2:  texasholdem.Rank_TWO,
			Kicker3:  texasholdem.Rank_EIGHT,
		},
		{
			Category: texasholdem.Score_PAIR,
			Kicker1:  texasholdem.Rank_TWO,
			Kicker2:  texasholdem.Rank_TEN,
			Kicker3:  texasholdem.Rank_SIX,
			Kicker4:  texasholdem.Rank_FOUR,
		},
		{
			Category: texasholdem.Score_HIGH_CARD,
			Kicker1:  texasholdem.Rank_TEN,
			Kicker2:  texasholdem.Rank_EIGHT,
			Kicker3:  texasholdem.Rank_SIX,
			Kicker4:  texasholdem.Rank_FOUR,
			Kicker5:  texasholdem.Rank_TWO,
		},
	}
	for i := 0; i < len(testCases); i++ {
		actual, err := ComputeScore(testCases[i])
		if nil != err {
			t.Error(err)
		}
		if !proto.Equal(actual, expected[i]) {
			t.Errorf("expected: %s actual: %s", proto.MarshalTextString(expected[i]), proto.MarshalTextString(actual))
			t.Errorf("cards: %s", CardsDebugString(testCases[i]))
		}
	}
}

func TestGetShuffledDeck(t *testing.T) {
	n := 10000
	rankSums := make([]int, 52)
	suitSums := make([]int, 52)
	for i := 0; i < n; i++ {
		cards := GetShuffledDeck()
		for j, card := range cards.GetCards() {
			rankSums[j] += int(card.GetRank())
			suitSums[j] += int(card.GetSuit())
		}
	}
	expectedRankSum := int(texasholdem.Rank_TWO + texasholdem.Rank_ACE) * n / 2
	allowedDelta := int(texasholdem.Rank_TWO + texasholdem.Rank_ACE) * int(math.Floor(math.Sqrt(float64(n))))
	rankSumMin := expectedRankSum - allowedDelta
	rankSumMax := expectedRankSum + allowedDelta
	for i := 0; i < 52; i++ {
		if rankSums[i] < rankSumMin || rankSums[i] > rankSumMax {
			t.Errorf("not random enough? rankSums[%d] = %d", i, rankSums[i])
		}
	}
	expectedSuitSum := int(texasholdem.Suit_SPADE + texasholdem.Suit_DIAMOND) * n / 2
	allowedDelta = int(texasholdem.Suit_SPADE + texasholdem.Suit_DIAMOND) * int(math.Floor(math.Sqrt(float64(n))))
	suitSumMin := expectedSuitSum - allowedDelta
	suitSumMax := expectedSuitSum + allowedDelta
	for i := 0; i < 52; i++ {
		if suitSums[i] < suitSumMin || suitSums[i] > suitSumMax {
			t.Errorf("not random enough? suitSums[%d] = %d", i, suitSums[i])
		}
	}
}
