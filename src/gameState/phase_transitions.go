package gameState

import (
	"fmt"
	"math/rand"
	"sort"

	"boardInfo"
)

func (g *GlobalState) timeString() string {
	result := fmt.Sprintf("%02d-%02d-%02d", g.Round, g.Phase, g.Turn)
	return result
}

func (g *GlobalState) marketPhase() bool {
	return g.Phase == 0
}
func (g *GlobalState) businessPhase() bool {
	return g.Phase == 1 || g.Phase == 2
}

func (g *GlobalState) currentTurn() string {
	return g.TurnOrder[g.Turn%len(g.TurnOrder)]
}

func (g *Game) beginMarketPhase() {
	g.Round += 1
	g.Phase = 0
	g.Turn = 0
	g.Passes = 0

	g.TurnOrder = make([]string, 0, len(g.Players))
	for name, _ := range g.Players {
		g.TurnOrder = append(g.TurnOrder, name)
	}
	sort.Sort(playerSorter{list: g.TurnOrder, info: g.Players})
}

func (g *Game) beginBusinessPhase() {
	g.Phase += 1
	g.Turn = 0

	g.TurnOrder = make([]string, 0, len(g.Companies))
	for name, company := range g.Companies {
		if company.President != "" {
			g.TurnOrder = append(g.TurnOrder, name)
		}
	}
	sort.Sort(companySorter{list: g.TurnOrder, info: g.Companies})
}

func (g *Game) endMarketTurn(pass bool) {
	g.Turn += 1

	// If every player has passed (in a row) then we are finished with the market phase and need
	// to begin the next business phase.
	if !pass {
		g.Passes = 0
	} else if g.Passes += 1; g.Passes == len(g.TurnOrder) {
		// Any companies that have orphaned stock by the end of the market phase have their stock
		// prices reduced.
		for name, _ := range g.OrphanStocks {
			company := g.Companies[name]
			company.StockPrice = boardInfo.PrevStockPrice(company.StockPrice)
		}
		g.beginBusinessPhase()
	}
}

func (g *Game) endBusinessTurn() {
	if g.Turn += 1; g.Turn == len(g.TurnOrder) {
		if g.Phase == 1 {
			g.beginBusinessPhase()
		} else {
			g.beginMarketPhase()
		}
	}
}

// The sorters implement sort.Interface and allow us to sort lists of the player and company names
// to determine turn order for the next phase.
type (
	playerSorter struct {
		list []string
		info map[string]*Player
	}
	companySorter struct {
		list []string
		info map[string]*Company
	}
)

func (s playerSorter) Len() int {
	return len(s.list)
}
func (s playerSorter) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}
func (s playerSorter) Less(i, j int) bool {
	item1, item2 := s.info[s.list[i]], s.info[s.list[j]]
	if item1.Cash != item2.Cash {
		return item1.Cash < item2.Cash
	} else if item1.NetWorth != item2.NetWorth {
		return item1.NetWorth < item2.NetWorth
	}
	return rand.Float32() < 0.5
}

func (s companySorter) Len() int {
	return len(s.list)
}
func (s companySorter) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}
func (s companySorter) Less(i, j int) bool {
	item1, item2 := s.info[s.list[i]], s.info[s.list[j]]
	if item1.StockPrice != item2.StockPrice {
		return item1.StockPrice > item2.StockPrice
	}
	return item1.PriceChange < item2.PriceChange
}
