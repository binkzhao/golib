package bitarr

import (
	"testing"
	"fmt"
)

func TestBitArr(t *testing.T)  {
	x, y := New(), New()
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x:", x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println("y:", y.String()) // "{9 42}"
	x.UnionWith(y)
	fmt.Println("x unionWith y:", x.String())         // "{1 9 42 144}"
	fmt.Println("x has 9,123:", x.Has(9), x.Has(123)) // "true false"
	fmt.Println("x len:", x.Len())                    //4
	fmt.Println("y len:", y.Len())                    //2
	x.Remove(42)
	fmt.Println("x after Remove 42:", x.String()) //{1 9 144}
	z := x.Copy()
	fmt.Println("z copy from x:", z.String()) //{1 9 144}
	x.Clear()
	fmt.Println("clear x:", x.String()) //{}
	x.AddAll(1, 2, 9)
	fmt.Println("x addAll 1,2,9:", x.String()) //{1 2 9}
	x.IntersectWith(y)
	fmt.Println("x intersectWith y:", x.String()) //{9}
	x.AddAll(1, 2)
	fmt.Println("x addAll 1,2:", x.String()) //{1 2 9}
	x.DiffWith(y)
	fmt.Println("x differenceWith y:", x.String()) //{1 2}
	x.AddAll(9, 144)
	fmt.Println("x addAll 9,144:", x.String()) //{1 2 9 144}
	x.SymmetricDiff(y)
	fmt.Println("x symmetricDifference y:", x.String()) //{1 2 42 144}
	for _, value := range x.Elems() {
		fmt.Print(value, " ") //1 2 42 144
	}
}
