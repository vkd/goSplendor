package splendor

import (
	"math/rand"
	"time"
)

type CardDeck struct {
	first *acard
	last  *acard

	ln  int
	rnd *rand.Rand
}

type acard struct {
	prev *acard
	next *acard

	card interface{}
}

func (c *CardDeck) Len() int {
	return c.ln
}

func (c *CardDeck) Get(i int) interface{} {
	if i < 0 || i >= c.ln {
		return nil
	}
	return c.get(i)
}

func (c *CardDeck) PopRandom() interface{} {
	if c.ln == 0 {
		return nil
	}

	if c.rnd == nil {
		c.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return c.pop(c.rnd.Intn(c.ln))
}

func (c *CardDeck) Pop(i int) interface{} {
	if i < 0 || i >= c.ln {
		return nil
	}
	return c.pop(i)
}

func (c *CardDeck) Put(card interface{}) {
	new_card := &acard{nil, nil, card}
	if c.last == nil {
		c.first = new_card
	} else {
		new_card.prev = c.last
		c.last.next = new_card
	}
	c.last = new_card
	c.ln++
}

func (c *CardDeck) get(i int) interface{} {
	card := c.first
	for j := 0; j < i; j++ {
		if card == nil {
			return nil
		}
		card = card.next
	}
	return card.card
}

func (c *CardDeck) pop(i int) interface{} {
	card := c.first
	for j := 0; j < i; j++ {
		if card == nil {
			return nil
		}
		card = card.next
	}
	c.ln--
	if card.prev != nil {
		card.prev.next = card.next
	}
	if card.next != nil {
		card.next.prev = card.prev
	}
	return card.card
}
