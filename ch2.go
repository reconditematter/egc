package egc

import (
	"github.com/reconditematter/cds"
	"sort"
	"sync"
)

// CH2 -- computes the convex hull of a collection of points in the plane.
// The function implements Graham's scan algorithm with Andrew's modification.
// It computes both the lower hull and the upper hull. The hull vertices are listed in
// counter-clockwise order. The function modifies the input `ps` by reordering it.
//
// Reference: R.L. Graham, An efficient algorithm for determining the convex hull of a
// finite planar set, Inform. Process. Lett., 1:132-133 (1972).
//
// DOI: https://dx.doi.org/10.1016/0020-0190(72)90045-2
//
// Reference: A.M. Andrew, Another efficient algorithm for convex hulls in two dimensions,
// Inform. Process. Lett., 9:216-219 (1979).
//
// DOI: https://dx.doi.org/10.1016/0020-0190(79)90072-3
func CH2(ps []Point2) (lower, upper []Point2) {
	//
	// Two special cases: n=0 or n=1.
	//
	n := len(ps)
	if n == 0 {
		lower, upper = []Point2{}, []Point2{}
		return
	}
	if n == 1 {
		lower, upper = []Point2{ps[0]}, []Point2{ps[0]}
		return
	}
	//
	// Sort the input in (x,y)-order.
	//
	sort.Sort(p2slice(ps))
	//
	// nort -- no right turn
	//
	nort := func(list *cds.Seq, p Point2) bool {
		return Orientation2(list.Get(list.Size()-2).(Point2), list.Gethi().(Point2), p) <= 0
	}
	//
	// Build the lower hull.
	//
	lowerhull := cds.NewSeq()
	for i := 0; i < n; i++ {
		pi := ps[i]
		for lowerhull.Size() > 1 && nort(lowerhull, pi) {
			_ = lowerhull.Pophi()
		}
		lowerhull.Addhi(pi)
	}
	//
	// Build the upper hull.
	//
	upperhull := cds.NewSeq()
	for i := n - 1; i >= 0; i-- {
		pi := ps[i]
		for upperhull.Size() > 1 && nort(upperhull, pi) {
			_ = upperhull.Pophi()
		}
		upperhull.Addhi(pi)
	}
	//
	// Special cases.
	//
	if lowerhull.Size() == 2 && lowerhull.Getlo().(Point2).CmpXY(lowerhull.Get(1).(Point2)) == 0 {
		_ = lowerhull.Pophi()
	}
	if upperhull.Size() == 2 && upperhull.Getlo().(Point2).CmpXY(upperhull.Get(1).(Point2)) == 0 {
		_ = upperhull.Pophi()
	}
	//
	//
	//
	lower = make([]Point2, lowerhull.Size())
	for i := range lower {
		lower[i] = lowerhull.Get(i).(Point2)
	}
	upper = make([]Point2, upperhull.Size())
	for i := range upper {
		upper[i] = upperhull.Get(i).(Point2)
	}
	return
}

// ParCH2 -- computes the convex hull of a collection of points in the plane.
// The function implements Graham's scan algorithm with Andrew's modification.
// It computes both the lower hull and the upper hull. The hull vertices are listed in
// counter-clockwise order. The function modifies the input `ps` by reordering it.
// The implementation uses `ncpu` goroutines running concurrently to accelerate
// the computations.
//
// Reference: R.L. Graham, An efficient algorithm for determining the convex hull of a
// finite planar set, Inform. Process. Lett., 1:132-133 (1972).
//
// DOI: https://dx.doi.org/10.1016/0020-0190(72)90045-2
//
// Reference: A.M. Andrew, Another efficient algorithm for convex hulls in two dimensions,
// Inform. Process. Lett., 9:216-219 (1979).
//
// DOI: https://dx.doi.org/10.1016/0020-0190(79)90072-3
func ParCH2(ncpu int, ps []Point2) (lower, upper []Point2) {
	if ncpu < 1 {
		ncpu = 1
	}
	if ncpu > 16 {
		ncpu = 16
	}
	n := len(ps)
	if n < ncpu {
		// Use the sequential algorithm.
		lower, upper = CH2(ps)
		return
	}
	//
	// Parallel loop.
	//
	var mu sync.Mutex
	var wg sync.WaitGroup
	coll := make([]Point2, 0)
	for cpu := 0; cpu < ncpu; cpu++ {
		wg.Add(1)
		first, limit := cpu*n/ncpu, (cpu+1)*n/ncpu
		go func(first, limit int) {
			lowerhull, upperhull := CH2(ps[first:limit])
			mu.Lock()
			coll = append(coll, lowerhull...)
			coll = append(coll, upperhull...)
			mu.Unlock()
			wg.Done()
		}(first, limit)
	}
	wg.Wait()
	//
	//
	//
	lower, upper = CH2(coll)
	return
}

// Sort interface implementation.
type p2slice []Point2

func (a p2slice) Len() int           { return len(a) }
func (a p2slice) Less(i, j int) bool { return a[i].CmpXY(a[j]) < 0 }
func (a p2slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
