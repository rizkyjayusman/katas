package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func TestPlusFiveDexterityVest1(t *testing.T) {
	item := &gildedrose.Item{"+5 Dexterity Vest", -1, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"+5 Dexterity Vest", -2, 18}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = -1
}

func TestPlusFiveDexterityVest2(t *testing.T) {
	item := &gildedrose.Item{"+5 Dexterity Vest", 10, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"+5 Dexterity Vest", 9, 19}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = -1
}

func TestAgedBrie1(t *testing.T) {
	item := &gildedrose.Item{"Aged Brie", 2, 0}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Aged Brie", 1, 1}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestAgedBrie2(t *testing.T) {
	item := &gildedrose.Item{"Aged Brie", -1, 0}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Aged Brie", -2, 2}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestElixirOfTheMongoose1(t *testing.T) {
	item := &gildedrose.Item{"Elixir of the Mongoose", 5, 7}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Elixir of the Mongoose", 4, 6}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestSulfurasHandOfRagnaros1(t *testing.T) {
	item := &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 80}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 80}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = 0
	// sell in = 0
}

func TestSulfurasHandOfRagnaros2(t *testing.T) {
	item := &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 30}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 31}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = 0
	// sell in = 0
}

func TestBackstagePassesToATAFKAL80ETCConcert(t *testing.T) {
	item := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 14, 21}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestBackstagePassesToATAFKAL80ETCConcert2(t *testing.T) {
	item := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 6, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 5, 22}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestBackstagePassesToATAFKAL80ETCConcert3(t *testing.T) {
	item := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 9, 22}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}

func TestBackstagePassesToATAFKAL80ETCConcert4(t *testing.T) {
	item := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", -1, 20}
	gildedrose.UpdateItemQuality(item)

	expected := &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", -2, 0}
	if item.Quality != expected.Quality {
		t.Error("value of quality is not match")
	}

	if item.SellIn != expected.SellIn {
		t.Error("value of sell in is not match")
	}

	// quality = -1
	// sell in = +1
}
