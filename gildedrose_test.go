package gildedrose

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_cover_aged_brie(t *testing.T) {
	var (
		assert = assert.New(t)
	)

	var (
		item Item

		expectedQuality = 2
		expectedSellIn  = -1
	)

	{
		item = Item{
			name: "Aged Brie",
		}
	}

	UpdateQuality([]*Item{&item})

	assert.Equal(expectedQuality, item.quality)
	assert.Equal(expectedSellIn, item.sellIn)
}
