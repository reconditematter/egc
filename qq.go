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
