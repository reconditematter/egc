// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package egc

// Vec2 -- represents a 2-dimensional vector with rational Cartesian coordinates (x,y).
type Vec2 struct {
	c [2]QQ
}

// NewVec2 -- returns vector (x,y).
func NewVec2(x, y QQ) Vec2 {
	return Vec2{[2]QQ{x, y}}
}

var (
	vecz = Vec2{[2]QQ{qq0, qq0}}
	veci = Vec2{[2]QQ{qq1, qq0}}
	vecj = Vec2{[2]QQ{qq0, qq1}}
)

// ZVec2 -- returns the zero vector (0,0).
func ZVec2() Vec2 {
	return vecz
}

// IVec2 -- returns the unit vector (1,0).
func IVec2() Vec2 {
	return veci
}

// JVec2 -- returns the unit vector (0,1).
func JVec2() Vec2 {
	return vecj
}

// X -- returns u.x.
func (u Vec2) X() QQ {
	return u.c[0]
}

// Y -- returns u.y.
func (u Vec2) Y() QQ {
	return u.c[1]
}

// C -- C(0) returns u.x; C(1) returns u.y.
func (u Vec2) C(i int) QQ {
	return u.c[i]
}

// SumAbs -- returns the L1-norm of `u`.
func (u Vec2) SumAbs() QQ {
	return u.c[0].Abs().Add(u.c[1].Abs())
}

// MaxAbs -- returns the L∞-norm of `u`.
func (u Vec2) MaxAbs() QQ {
	return u.c[0].Abs().Max(u.c[1].Abs())
}

// AbsSq -- returns |u|².
func (u Vec2) AbsSq() QQ {
	return u.c[0].Sq().Add(u.c[1].Sq())
}

// Neg -- returns -u.
func (u Vec2) Neg() Vec2 {
	return Vec2{[2]QQ{u.c[0].Neg(), u.c[1].Neg()}}
}

// Add -- returns u+v.
func (u Vec2) Add(v Vec2) Vec2 {
	return Vec2{[2]QQ{u.c[0].Add(v.c[0]), u.c[1].Add(v.c[1])}}
}

// Sub -- returns u-v.
func (u Vec2) Sub(v Vec2) Vec2 {
	return Vec2{[2]QQ{u.c[0].Sub(v.c[0]), u.c[1].Sub(v.c[1])}}
}

// Mul -- returns s*u.
func (u Vec2) Mul(s QQ) Vec2 {
	return Vec2{[2]QQ{s.Mul(u.c[0]), s.Mul(u.c[1])}}
}

// Div -- returns (1/s)*u.
func (u Vec2) Div(s QQ) Vec2 {
	return Vec2{[2]QQ{u.c[0].Div(s), u.c[1].Div(s)}}
}

// Dot -- returns the scalar (dot) product of `u` and `v`.
func (u Vec2) Dot(v Vec2) QQ {
	return u.c[0].Mul(v.c[0]).Add(u.c[1].Mul(v.c[1]))
}

// String -- returns a string representation of `u` in the form "(x,y)".
func (u Vec2) String() string {
	return "(" + u.c[0].String() + "," + u.c[1].String() + ")"
}
