// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package egc

// Point2 -- represents a 2-dimensional point with rational Cartesian coordinates (x,y).
type Point2 struct {
	c [2]QQ
}

// NewPoint2 -- returns point (x,y).
func NewPoint2(x, y QQ) Point2 {
	return Point2{[2]QQ{x, y}}
}

// X -- returns p.x.
func (p Point2) X() QQ {
	return p.c[0]
}

// Y -- returns p.y.
func (p Point2) Y() QQ {
	return p.c[1]
}

// C -- C(0) returns p.x; C(1) returns p.y.
func (p Point2) C(i int) QQ {
	return p.c[i]
}

// Plus -- returns p+u (translation by +u).
func (p Point2) Plus(u Vec2) Point2 {
	return Point2{[2]QQ{p.c[0].Add(u.c[0]), p.c[1].Add(u.c[1])}}
}

// Minus -- returns p-u (translation by -u).
func (p Point2) Minus(u Vec2) Point2 {
	return Point2{[2]QQ{p.c[0].Sub(u.c[0]), p.c[1].Sub(u.c[1])}}
}

// CmpXY -- compares the coordinates of `p` and `q` lexicographically in xy-order.
func (p Point2) CmpXY(q Point2) int {
	cx := p.c[0].Cmp(q.c[0])
	if cx != 0 {
		return cx
	}
	cy := p.c[1].Cmp(q.c[1])
	return cy
}

// CmpYX -- compares the coordinates of `p` and `q` lexicographically in yx-order.
func (p Point2) CmpYX(q Point2) int {
	cy := p.c[1].Cmp(q.c[1])
	if cy != 0 {
		return cy
	}
	cx := p.c[0].Cmp(q.c[0])
	return cx
}

// Mid -- returns the midpoint between `p` and `q`.
func (p Point2) Mid(q Point2) Point2 {
	x := p.c[0].Add(q.c[0])
	y := p.c[1].Add(q.c[1])
	return Point2{[2]QQ{x.Mul(qq2r), y.Mul(qq2r)}}
}

// Centroid -- returns the centroid of `p`, `q`, and `r`.
func (p Point2) Centroid(q, r Point2) Point2 {
	x := p.c[0].Add(q.c[0]).Add(r.c[0])
	y := p.c[1].Add(q.c[1]).Add(r.c[1])
	return Point2{[2]QQ{x.Mul(qq3r), y.Mul(qq3r)}}
}

// Cir -- returns the circumcenter of `p`, `q`, and `r`.
// This method causes a runtime panic when the three points are collinear.
func (p Point2) Cir(q, r Point2) Point2 {
	cct := func(dqx, dqy, drx, dry QQ) (QQ, QQ) {
		q2 := dqx.Sq().Add(dqy.Sq())
		r2 := drx.Sq().Add(dry.Sq())
		den := Det2x2([2][2]QQ{{dqx, dqy}, {drx, dry}}).Mul(qq2)
		if den.Sgn() == 0 {
			panic("egc.Point2.Cir: collinear points")
		}
		dcx := Det2x2([2][2]QQ{{dry, dqy}, {r2, q2}}).Div(den)
		dcy := Det2x2([2][2]QQ{{drx, dqx}, {r2, q2}}).Div(den)
		return dcx, dcy.Neg()
	}
	//
	x, y := cct(q.c[0].Sub(p.c[0]), q.c[1].Sub(p.c[1]), r.c[0].Sub(p.c[0]), r.c[1].Sub(p.c[1]))
	return Point2{[2]QQ{x.Add(p.c[0]), y.Add(p.c[1])}}
}

// DistSq -- returns the square of the distance between `p` and `q`.
func (p Point2) DistSq(q Point2) QQ {
	dx := p.c[0].Sub(q.c[0])
	dy := p.c[1].Sub(q.c[1])
	return dx.Sq().Add(dy.Sq())
}

// Orientation2 -- returns the orientation of `p`, `q`, and `r` (in that order).
//
//     -1 = counterclockwise,
//      0 = collinear,
//     +1 = clockwise.
func Orientation2(p, q, r Point2) int {
	D := [3][3]QQ{{qq1, p.c[0], p.c[1]}, {qq1, q.c[0], q.c[1]}, {qq1, r.c[0], r.c[1]}}
	return Det3x3(D).Sgn()
}

// String -- returns a string representation of `p` in the form "(x,y)".
func (p Point2) String() string {
	return "(" + p.c[0].String() + "," + p.c[1].String() + ")"
}
