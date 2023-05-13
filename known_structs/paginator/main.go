package main

import (
	"fmt"
	"math"
)

func main() {
	testMe(generates())
}
func testMe(pgs []*Paginator) {
	for _, pg := range pgs {
		fmt.Println("**************************************************")
		fmt.Printf("Items: %d ItemsPerPage: %d\n", pg.Items, pg.ItemsPerPage)

		fmt.Printf("PageCount: %d\n", pg.PageCount())
		fmt.Printf("ItemsCount: %d\n", pg.ItemsCount())
		fmt.Printf("first page %d\n", pg.PageLength(1))
		fmt.Printf("third page %d\n", pg.PageLength(3))
		fmt.Printf("last page %d\n", pg.PageLength(pg.PageCount()-1))
	}

}
func generates() []*Paginator {
	ret := make([]*Paginator, 6)
	ret[0] = New([]int{1, 2}, 2)
	ret[1] = New([]int{1}, 1)
	ret[2] = New([]int{1, 2, 3, 4, 5}, 2)
	ret[3] = New([]int{1, 2, 3, 4, 5, 6, 4}, 3)
	ret[4] = New([]int{1, 2, 3, 22, 33, 3, 4, 5, 4}, 6)
	ret[5] = New([]int{1, 2, 3, 4, 5, 6}, 2)

	return ret
}

// itemscount --> number of items
// pagelength --> number of items per page
// pageCount
// pageindex
// 2 params (items, page divider)
type Paginator struct {
	Items        []int
	ItemsPerPage int
}

func (pg *Paginator) ItemsCount() int {
	return len(pg.Items)
}
func (pg *Paginator) PageCount() int {
	numPages := math.Ceil(float64(pg.ItemsCount()) / float64(pg.ItemsPerPage))
	return int(numPages)
}
func (pg *Paginator) PageLength(pageNum int) int {
	remainingItems := pg.ItemsCount() % pg.ItemsPerPage
	if pg.PageCount() == 1 {
		return pg.ItemsCount()
	}

	if pageNum == pg.PageCount() {
		return remainingItems
	} else if pageNum > pg.PageCount() || pageNum <= 0 {
		return 0
	} else {
		return pg.ItemsPerPage
	}

}
func New(items []int, itemsPerPage int) *Paginator {
	return &Paginator{Items: items, ItemsPerPage: itemsPerPage}
}
