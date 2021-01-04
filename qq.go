// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package egc

import (
	"math/big"
)

// QQ -- represents a rational number of arbitrary precision.
type QQ struct {
	r big.Rat
}

var (
	qq0  = Ratio(0, 1)
	qq1  = Ratio(1, 1)
	qq2  = Ratio(2, 1)
	qq3  = Ratio(3, 1)
	qq4  = Ratio(4, 1)
	qq5  = Ratio(5, 1)
	qq6  = Ratio(6, 1)
	qq7  = Ratio(7, 1)
	qq8  = Ratio(8, 1)
	qq9  = Ratio(9, 1)
	qq10 = Ratio(10, 1)
)

var (
	qq2r  = Ratio(1, 2)
	qq3r  = Ratio(1, 3)
	qq4r  = Ratio(1, 4)
	qq5r  = Ratio(1, 5)
	qq6r  = Ratio(1, 6)
	qq7r  = Ratio(1, 7)
	qq8r  = Ratio(1, 8)
	qq9r  = Ratio(1, 9)
	qq10r = Ratio(1, 10)
)

// Ratio -- returns the rational number p/q.
func Ratio(p, q int64) QQ {
	var c big.Rat
	c.SetFrac64(p, q)
	return QQ{c}
}

// Neg -- returns -a.
func (a QQ) Neg() QQ {
	var c big.Rat
	c.Neg(&a.r)
	return QQ{c}
}

// Inv -- returns 1/a.
func (a QQ) Inv() QQ {
	var c big.Rat
	c.Inv(&a.r)
	return QQ{c}
}

// Abs -- returns |a|.
func (a QQ) Abs() QQ {
	var c big.Rat
	c.Abs(&a.r)
	return QQ{c}
}

// Sq -- returns a².
func (a QQ) Sq() QQ {
	var c big.Rat
	c.Mul(&a.r, &a.r)
	return QQ{c}
}

// Add -- returns a+b.
func (a QQ) Add(b QQ) QQ {
	var c big.Rat
	c.Add(&a.r, &b.r)
	return QQ{c}
}

// Sub -- returns a-b.
func (a QQ) Sub(b QQ) QQ {
	var c big.Rat
	c.Sub(&a.r, &b.r)
	return QQ{c}
}

// Mul -- returns a*b.
func (a QQ) Mul(b QQ) QQ {
	var c big.Rat
	c.Mul(&a.r, &b.r)
	return QQ{c}
}

// Div -- returns a/b.
func (a QQ) Div(b QQ) QQ {
	var c big.Rat
	c.Quo(&a.r, &b.r)
	return QQ{c}
}

// Sgn -- returns -1, 0, or +1 corresponding to a<0, a=0, or a>0.
func (a QQ) Sgn() int {
	return a.r.Sign()
}

// Cmp -- returns -1, 0, or +1 corresponding to a<b, a=b, or a>b.
func (a QQ) Cmp(b QQ) int {
	return a.r.Cmp(&b.r)
}

// Min -- returns min{a,b}.
func (a QQ) Min(b QQ) QQ {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// Max -- returns max{a,b}.
func (a QQ) Max(b QQ) QQ {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// String -- returns a string representation of `a` in the form "p/q".
func (a QQ) String() string {
	return a.r.String()
}
