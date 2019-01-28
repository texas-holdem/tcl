package tcl

import (
	"github.com/golang/protobuf/proto"
	"github.com/texas-holdem/tcl/proto"
	"testing"
)

func TestAbbrToCard(t *testing.T) {
	testCases := []string{"2h", "3s", "tc", "qd", "ac"}
	expected := []*texasholdem.Card{
		{
			Rank: texasholdem.Rank_TWO,
			Suit: texasholdem.Suit_HEART,
		},
		{
			Rank: texasholdem.Rank_THREE,
			Suit: texasholdem.Suit_SPADE,
		},
		{
			Rank: texasholdem.Rank_TEN,
			Suit: texasholdem.Suit_CLUB,
		},
		{
			Rank: texasholdem.Rank_QUEEN,
			Suit: texasholdem.Suit_DIAMOND,
		},
		{
			Rank: texasholdem.Rank_ACE,
			Suit: texasholdem.Suit_CLUB,
		},
	}
	for i := 0; i < len(testCases); i++ {
		actual := AbbrToCard(testCases[i])
		if !proto.Equal(actual, expected[i]) {
			t.Errorf("expected: %s actual: %s", proto.MarshalTextString(expected[i]), proto.MarshalTextString(actual))
		}
	}
}
