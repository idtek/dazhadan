package card

import (
	"testing"
	"github.com/bmizerany/assert"
)

func TestGetCardsType(t *testing.T) {
	card1 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	card2 := &Card{CardType: CardType_Club,CardNo: 1,}	//梅花1
	card3 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	card4 := &Card{CardType: CardType_Club,CardNo: 2,}	//梅花2
	card5 := &Card{CardType: CardType_Heart,CardNo: 2,}	//红桃2
	card6 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2
	card7 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2
	card8 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2
	card9 := &Card{CardType: CardType_Diamond,CardNo: 2,}	//方片2

	card10 := &Card{CardType: CardType_BlackJoker,CardNo: 14,}	//小王
	card11 := &Card{CardType: CardType_BlackJoker,CardNo: 14,}	//小王
	card12 := &Card{CardType: CardType_RedJoker,CardNo: 14,}	//大王
	card13 := &Card{CardType: CardType_RedJoker,CardNo: 14,}	//大王

	card14 := &Card{CardType: CardType_Diamond,CardNo: 3,}	//方片3
	card15 := &Card{CardType: CardType_Diamond,CardNo: 4,}	//方片4
	card16 := &Card{CardType: CardType_Diamond,CardNo: 5,}	//方片5
	card17 := &Card{CardType: CardType_Diamond,CardNo: 6,}	//方片6
	card18 := &Card{CardType: CardType_Diamond,CardNo: 7,}	//方片7
	card19 := &Card{CardType: CardType_Diamond,CardNo: 3,}	//方片3
	card20 := &Card{CardType: CardType_Diamond,CardNo: 4,}	//方片3
	card21 := &Card{CardType: CardType_Diamond,CardNo: 5,}	//方片3

	cards1 := make([]*Card, 0)
	cards1 = append(cards1, card1)
	cards1 = append(cards1, card2)
	cards1 = append(cards1, card3)
	cards1 = append(cards1, card4)
	drop_cards1 := CreateNewCards(cards1)

	cards2 := make([]*Card, 0)
	cards2 = append(cards2, card4)
	cards2 = append(cards2, card5)
	cards2 = append(cards2, card6)
	cards2 = append(cards2, card9)
	drop_cards2 := CreateNewCards(cards2)

	cards3 := make([]*Card, 0)
	cards3 = append(cards3, card10)
	cards3 = append(cards3, card11)
	cards3 = append(cards3, card12)
	cards3 = append(cards3, card13)
	drop_cards3 := CreateNewCards(cards3)

	cards4 := make([]*Card, 0)
	cards4 = append(cards4, card10)
	cards4 = append(cards4, card11)
	cards4 = append(cards4, card12)
	drop_cards4 := CreateNewCards(cards4)

	cards5 := make([]*Card, 0)
	//cards5 = append(cards5, card1)
	//cards5 = append(cards5, card4)
	cards5 = append(cards5, card14)
	cards5 = append(cards5, card15)
	cards5 = append(cards5, card16)
	cards5 = append(cards5, card17)
	cards5 = append(cards5, card18)
	drop_cards5 := CreateNewCards(cards5)

	cards6 := make([]*Card, 0)
	cards6 = append(cards6, card1)
	cards6 = append(cards6, card2)
	cards6 = append(cards6, card5)
	cards6 = append(cards6, card6)
	drop_cards6 := CreateNewCards(cards6)

	cards7 := make([]*Card, 0)
	cards7 = append(cards7, card4)
	cards7 = append(cards7, card5)
	cards7 = append(cards7, card6)
	cards7 = append(cards7, card7)
	cards7 = append(cards7, card8)
	cards7 = append(cards7, card9)
	cards7 = append(cards7, card10)
	cards7 = append(cards7, card11)
	cards7 = append(cards7, card12)
	cards7 = append(cards7, card13)
	drop_cards7 := CreateNewCards(cards7)

	cards8 := make([]*Card, 0)
	cards8 = append(cards8, card1)
	cards8 = append(cards8, card2)
	cards8 = append(cards8, card3)
	cards8 = append(cards8, card4)
	cards8 = append(cards8, card5)
	cards8 = append(cards8, card6)
	cards8 = append(cards8, card14)
	cards8 = append(cards8, card15)
	cards8 = append(cards8, card16)
	cards8 = append(cards8, card17)
	drop_cards8 := CreateNewCards(cards8)

	cards9 := make([]*Card, 0)
	cards9 = append(cards9, card14)
	cards9 = append(cards9, card15)
	cards9 = append(cards9, card16)
	cards9 = append(cards9, card19)
	cards9 = append(cards9, card20)
	cards9 = append(cards9, card21)
	drop_cards9 := CreateNewCards(cards9)

	t.Log(GetCardsType(drop_cards1, true))
	t.Log(GetCardsType(drop_cards2, false))
	t.Log(GetCardsType(drop_cards3, false))
	t.Log(GetCardsType(drop_cards4, false))
	t.Log(GetCardsType(drop_cards5, false))
	t.Log(GetCardsType(drop_cards6, false))
	t.Log(GetCardsType(drop_cards7, false))
	t.Log(GetCardsType(drop_cards7, true))
	t.Log(GetCardsType(drop_cards8, false))
	t.Log(GetCardsType(drop_cards9, false))
}

func TestGetSameCardsNum(t *testing.T) {
	card1 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	card2 := &Card{CardType: CardType_Club,CardNo: 1,}	//梅花1
	card3 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	card4 := &Card{CardType: CardType_Club,CardNo: 2,}	//梅花2
	card5 := &Card{CardType: CardType_Heart,CardNo: 2,}	//红桃2
	card6 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2
	card7 := &Card{CardType: CardType_Spade,CardNo: 5,}	//黑桃5
	card8 := &Card{CardType: CardType_Spade,CardNo: 8,}	//黑桃8
	//card9 := &Card{CardType: CardType_Diamond,CardNo: 2,}	//方片2

	cards1 := make([]*Card, 0)
	cards1 = append(cards1, card1)
	cards1 = append(cards1, card2)
	cards1 = append(cards1, card3)
	cards1 = append(cards1, card4)

	cards2 := make([]*Card, 0)
	cards2 = append(cards2, card1)
	cards2 = append(cards2, card2)
	cards2 = append(cards2, card3)
	cards2 = append(cards2, card4)
	cards2 = append(cards2, card5)
	cards2 = append(cards2, card6)
	cards2 = append(cards2, card7)
	cards2 = append(cards2, card8)
	//cards2 = append(cards2, card9)

	t.Log(GetSameCardsNum(cards1))
	t.Log(GetSameCardsNum(cards2))
}

func TestIs510K(t *testing.T) {
	card1 := &Card{CardType: CardType_Diamond,CardNo: 5,}	//方片5
	card2 := &Card{CardType: CardType_Club,CardNo: 10,}	//梅花10
	card3 := &Card{CardType: CardType_Diamond,CardNo: 13,}	//方片K
	card4 := &Card{CardType: CardType_Diamond,CardNo: 10,}	//方片10
	card5 := &Card{CardType: CardType_Heart,CardNo: 2,}	//红桃2
	card9 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2

	cards1 := make([]*Card, 0)
	cards1 = append(cards1, card1)
	cards1 = append(cards1, card2)
	cards1 = append(cards1, card3)
	t.Log(Is510K(cards1))

	cards2 := make([]*Card, 0)
	cards2 = append(cards2, card1)
	cards2 = append(cards2, card4)
	cards2 = append(cards2, card3)
	cards2 = append(cards2, card9)
	t.Log(Is510K(cards2))

	cards3 := make([]*Card, 0)
	cards3 = append(cards3, card1)
	cards3 = append(cards3, card4)
	cards3 = append(cards3, card3)
	t.Log(Is510K(cards3))

	cards4 := make([]*Card, 0)
	cards4 = append(cards4, card1)
	cards4 = append(cards4, card5)
	cards4 = append(cards4, card3)
	t.Log(Is510K(cards4))
}

func TestIsSameCardType(t *testing.T) {
	card1 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	card2 := &Card{CardType: CardType_Club,CardNo: 1,}	//梅花1
	card3 := &Card{CardType: CardType_Diamond,CardNo: 1,}	//方片1
	//card4 := &Card{CardType: CardType_Club,CardNo: 2,}	//梅花2
	card5 := &Card{CardType: CardType_Heart,CardNo: 2,}	//红桃2
	card6 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2
	card7 := &Card{CardType: CardType_Spade,CardNo: 5,}	//黑桃5
	card8 := &Card{CardType: CardType_Spade,CardNo: 8,}	//黑桃8
	card9 := &Card{CardType: CardType_Spade,CardNo: 2,}	//黑桃2

	cards1 := make([]*Card, 0)
	cards1 = append(cards1, card1)
	cards1 = append(cards1, card2)
	cards1 = append(cards1, card3)
	assert.Equal(t, IsSameCardType(cards1), false)

	cards2 := make([]*Card, 0)
	cards2 = append(cards2, card6)
	cards2 = append(cards2, card7)
	cards2 = append(cards2, card8)
	cards2 = append(cards2, card9)
	assert.Equal(t, IsSameCardType(cards2), true)

	cards3 := make([]*Card, 0)
	cards3 = append(cards3, card5)
	assert.Equal(t, IsSameCardType(cards3), true)

	cards4 := make([]*Card, 0)
	assert.Equal(t, IsSameCardType(cards4), false)
}

func TestIsStraight(t *testing.T) {
	nums1 := []int{2,3,4}
	assert.Equal(t, IsStraight(nums1), false)

	nums2 := []int{3,4}
	assert.Equal(t, IsStraight(nums2), true)

	nums3 := []int{1,2,13,12}
	assert.Equal(t, IsStraight(nums3), true)
}

func TestGetStraightWeight(t *testing.T) {
	nums1 := []int{2,3,4}
	t.Log(GetStraightWeight(nums1))

	nums2 := []int{3,4}
	t.Log(GetStraightWeight(nums2))

	nums3 := []int{1,2,13,12}
	t.Log(GetStraightWeight(nums3))
}