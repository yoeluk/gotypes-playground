package main

import (
	"fmt"
)

type GtOrdering[S any] interface {
	Gt(a S, b S) bool
}

// ShowGtOrdering generic show gt ordering function
func ShowGtOrdering[O GtOrdering[S], S fmt.Stringer](o O, a S, b S) {
	fmt.Printf("%s is greater than %s? %t \n", a.String(), b.String(), o.Gt(a, b))
}

type Ordering[S any] interface {
	Equiv(a S, b S) bool
	Lt(a S, b S) bool
	Gt(a S, b S) bool
}

// ShowOrdering generic show ordering function
func ShowOrdering[O Ordering[S], S fmt.Stringer](o O, a S, b S) {
	ShowOrderingCurrying[O, S](o)(a, b)
}

// ShowOrderingCurrying currying generic show ordering function
func ShowOrderingCurrying[O Ordering[S], S fmt.Stringer](o O) func(S, S) {
	return func(a S, b S) {
		fmt.Printf("%s is greater than %s? %t \n", a.String(), b.String(), o.Gt(a, b))
		fmt.Printf("%s is less than %s? %t \n", a.String(), b.String(), o.Lt(a, b))
		fmt.Printf("%s is equivalent to %s? %t \n", a.String(), b.String(), o.Equiv(a, b))
	}
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// StringerSignedInteger is the intersection of signed integers
// with those that have a String() -> string method
type StringerSignedInteger interface {
	SignedInteger
	String() string
}

// ByFactorShowOrdering generic show ordering function
func ByFactorShowOrdering[O Ordering[S], S StringerSignedInteger](o O, a S, b S) {
	ShowOrderingCurrying[O, S](o)(-2*a, -3*b)
}

/*	IntOrdering for int types
 */
type IntOrdering[I ~int] struct{}

func (*IntOrdering[int]) Equiv(a int, b int) bool {
	return a == b
}

func (*IntOrdering[int]) Gt(a int, b int) bool {
	return a > b
}

func (*IntOrdering[int]) Lt(a int, b int) bool {
	return a < b
}

// myInt is a type alias; int can't be used in its place
type myInt int

// String for myInt type
func (b myInt) String() string {
	return fmt.Sprintf("%d", b)
}

/*	Butterfly Ordering
 */
type Butterfly struct {
	WingSpam    int32
	IsNocturnal bool
}

func (*Butterfly) Equiv(a Butterfly, b Butterfly) bool {
	return a.WingSpam == b.WingSpam && a.IsNocturnal == b.IsNocturnal
}

func (*Butterfly) Gt(a Butterfly, b Butterfly) bool {
	return a.IsNocturnal && !b.IsNocturnal
}

func (*Butterfly) Lt(a Butterfly, b Butterfly) bool {
	return b.IsNocturnal && !a.IsNocturnal
}

// String for Butterfly type
func (b Butterfly) String() string {
	return fmt.Sprintf("ButterFly{%d, %t}", b.WingSpam, b.IsNocturnal)
}

func main() {
	io := &IntOrdering[myInt]{}

	// instantiate a generic ShowGtOrdering for IntOrdering
	intGtPrinter := ShowGtOrdering[*IntOrdering[myInt], myInt]

	// use the show function to show gt for 2 myInts
	intGtPrinter(io, myInt(9), myInt(5))

	// type inference here isn't too bad... we don't need the instantiation
	// although it could be helpful for passing around
	// the compiler can infer the types from the arguments at the call-site
	ShowGtOrdering(io, myInt(9), myInt(5))

	// io is full ordering so I can use it as Ordering[myInt] too
	ShowOrdering(io, myInt(9), myInt(5))

	// ButterflyOrdering
	oneButterfly := Butterfly{WingSpam: 2, IsNocturnal: true}
	twoButterfly := Butterfly{WingSpam: 5, IsNocturnal: false}

	bo := &Butterfly{}
	ShowOrdering(bo, oneButterfly, twoButterfly)

	// Currying ordering
	boOrdCurried := ShowOrderingCurrying[*Butterfly, Butterfly](bo)

	// calling curried ordering
	boOrdCurried(oneButterfly, twoButterfly)

	ByFactorShowOrdering(io, myInt(2), myInt(5))

	// ByFactorShowOrdering(bo, oneButterfly, twoButterfly) -- compiler error: Butterfly does not implement StringerSignedInteger
	ints := make([]myInt, 0, 0)
	ints = append(ints, 3, 2, 5, 8, 0, 40, 7, 11, 10)
	sorted := Sorted(io, ints)

	fmt.Printf("\n")
	fmt.Printf("%v", sorted)
}
