// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package egc

// Det2x2 -- computes the determinant of a 2x2 matrix.
func Det2x2(a [2][2]QQ) QQ {
	m01 := a[0][0].Mul(a[1][1]).Sub(a[1][0].Mul(a[0][1]))
	return m01
}

// Det3x3 -- computes the determinant of a 3x3 matrix.
func Det3x3(a [3][3]QQ) QQ {
	m01 := a[0][0].Mul(a[1][1]).Sub(a[1][0].Mul(a[0][1]))
	m02 := a[0][0].Mul(a[2][1]).Sub(a[2][0].Mul(a[0][1]))
	m12 := a[1][0].Mul(a[2][1]).Sub(a[2][0].Mul(a[1][1]))
	m012 := m01.Mul(a[2][2]).Sub(m02.Mul(a[1][2])).Add(m12.Mul(a[0][2]))
	return m012
}

// Det4x4 -- computes the determinant of a 4x4 matrix.
func Det4x4(a [4][4]QQ) QQ {
	m01 := a[1][0].Mul(a[0][1]).Sub(a[0][0].Mul(a[1][1]))
	m02 := a[2][0].Mul(a[0][1]).Sub(a[0][0].Mul(a[2][1]))
	m03 := a[3][0].Mul(a[0][1]).Sub(a[0][0].Mul(a[3][1]))
	m12 := a[2][0].Mul(a[1][1]).Sub(a[1][0].Mul(a[2][1]))
	m13 := a[3][0].Mul(a[1][1]).Sub(a[1][0].Mul(a[3][1]))
	m23 := a[3][0].Mul(a[2][1]).Sub(a[2][0].Mul(a[3][1]))
	m012 := m12.Mul(a[0][2]).Sub(m02.Mul(a[1][2])).Add(m01.Mul(a[2][2]))
	m013 := m13.Mul(a[0][2]).Sub(m03.Mul(a[1][2])).Add(m01.Mul(a[3][2]))
	m023 := m23.Mul(a[0][2]).Sub(m03.Mul(a[2][2])).Add(m02.Mul(a[3][2]))
	m123 := m23.Mul(a[1][2]).Sub(m13.Mul(a[2][2])).Add(m12.Mul(a[3][2]))
	m0123 := m123.Mul(a[0][3]).Sub(m023.Mul(a[1][3])).Add(m013.Mul(a[2][3])).Sub(m012.Mul(a[3][3]))
	return m0123
}
