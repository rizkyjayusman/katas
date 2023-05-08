package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		UpdateItemQuality(items[i])
	}
}

func UpdateItemQuality(item *Item) {
	if item.Name == "+5 Dexterity Vest" || item.Name == "Elixir of the Mongoose" {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 && item.Quality > 0 {
			item.Quality = item.Quality - 1
		}
	} else if item.Name == "Aged Brie" {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 && item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
	} else if item.Name == "Sulfuras, Hand of Ragnaros" {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
	} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}

		if item.Quality < 50 {
			if item.SellIn < 11 {
				item.Quality = item.Quality + 1
			}

			if item.SellIn < 6 {
				item.Quality = item.Quality + 1
			}
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 {
			item.Quality = item.Quality - item.Quality
		}
	}
}
