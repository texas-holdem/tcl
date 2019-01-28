package tcl

import (
	"github.com/golang/protobuf/proto"
	"github.com/texas-holdem/tcl/proto"
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
