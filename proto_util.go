package tcl

import (
	"fmt"
	"github.com/texas-holdem/tcl/proto"
	"sort"
	"strings"
)

var SuitChar = []rune("?♠♥♣♦")
var RankChar = []rune("?23456789TJQKA")

func CardToString(card *texasholdem.Card) string {
	return fmt.Sprintf("%s%s", string(SuitChar[int(card.GetSuit())]), string(RankChar[int(card.GetRank())]))
}

func CardsDebugString(cards []*texasholdem.Card) string {
	cardStrings := make([]string, len(cards))
	for i, card := range cards {
		cardStrings[i] = CardToString(card)
	}
	return fmt.Sprintf("[%s]", strings.Join(cardStrings, ", "))
}

func SortedCards(hand []*texasholdem.Card) []*texasholdem.Card {
	cards := make([]*texasholdem.Card, len(hand))
	copy(cards, hand)
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].GetRank() < cards[j].GetRank()
	})
	return cards
}

var AbbrToRankMap map[uint8]texasholdem.Rank
var AbbrToSuitMap map[uint8]texasholdem.Suit

func AbbrToCard(abbr string) *texasholdem.Card {
	if len(abbr) != 2 {
		return nil
	}
	rank := AbbrToRankMap[abbr[0]]
	suit := AbbrToSuitMap[abbr[1]]
	return &texasholdem.Card{
		Rank: rank,
		Suit: suit,
	}
}

func AbbrsToCards(abbrs []string) []*texasholdem.Card {
	cards := make([]*texasholdem.Card, len(abbrs))
	for i := 0; i < len(abbrs); i++ {
		cards[i] = AbbrToCard(abbrs[i])
	}
	return cards
}

func compareScore(score1, score2 *texasholdem.Score) int {
	if score1.GetCategory() < score2.GetCategory() {
		return -1
	} else if score1.GetCategory() > score2.GetCategory() {
		return 1
	}
	for i := range score1.GetKicker() {
		if score1.GetKicker()[i] < score2.GetKicker()[i] {
			return -1
		} else if score1.GetKicker()[i] > score2.GetKicker()[i] {
			return 1
		}
	}
	return 0
}

const RawRankChar = " 23456789tjqka"
const RawSuitChar = " shcd"

func init() {
	AbbrToRankMap = make(map[uint8]texasholdem.Rank)
	for i := 0; i < len(RawRankChar); i++ {
		AbbrToRankMap[RawRankChar[i]] = texasholdem.Rank(i)
	}
	AbbrToSuitMap = make(map[uint8]texasholdem.Suit)
	for i := 0; i < len(RawSuitChar); i++ {
		AbbrToSuitMap[RawSuitChar[i]] = texasholdem.Suit(i)
	}
}
