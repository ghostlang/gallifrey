package gallifrey

// IntervalTree is a tree of intervals
type IntervalTree interface {
	Insert(...Interval)
	// Intersection(Interval) int64
	Contains(Interval) bool
}

// intervalTree is a Discrete Interval Encoding Tree, allowing insertion of ranges of
// integers and fast intersection and membership calculation.
type intervalTree struct {
	root *node
}

// NewIntervalTree returns a new intervalTree.
func NewIntervalTree() IntervalTree {
	return &intervalTree{}
}

// Insert adds a new range of integers to the tree.
func (d *intervalTree) Insert(intervals ...Interval) {
	for _, i := range intervals {
		d.root = insert(i.Start(), i.End(), d.root)
	}
}

// Balance balances the tree using the DSW algorithm. It is most efficient to
// do this after the tree is complete.
func (d *intervalTree) Balance() {
	if d.root != nil {
		d.root = balance(d.root)
	}
}

// Intersection finds the intersection of the range of integers specified with
// any of the members of the tree. It returns the number of members in common.
func (d *intervalTree) Intersection(i Interval) int64 {
	return intersection(i.Start(), i.End(), d.root)
}

// IntersectionAll finds the number of members in common between two intervalTrees.
func (d *intervalTree) IntersectionAll(other *intervalTree) int64 {
	return intersectionAll(d.root, other)
}

// Total finds the number of integers represented by this tree.
func (d *intervalTree) Total() int64 {
	return total(d.root)
}

// Contains returns whether all of the range specified is contained within this
// diet.
func (d *intervalTree) Contains(i Interval) bool {
	return intersection(i.Start(), i.End(), d.root) == i.End()-i.Start()+1
}

type node struct {
	min   int64
	max   int64
	left  *node
	right *node
}

func splitMax(min, max int64, left, right *node) (int64, int64, *node) {
	if right == nil {
		return min, max, left
	}
	u, v, rprime := splitMax(right.min, right.max, right.left, right.right)
	newd := &node{min, max, left, rprime}
	return u, v, newd
}

func splitMin(min, max int64, left, right *node) (int64, int64, *node) {
	if left == nil {
		return min, max, right
	}
	u, v, lprime := splitMin(left.min, left.max, left.left, left.right)
	newd := &node{min, max, lprime, right}
	return u, v, newd
}

func joinLeft(min, max int64, left, right *node) *node {
	if left != nil {
		xprime, yprime, lprime := splitMax(left.min, left.max, left.left, left.right)
		if yprime+1 == min {
			return &node{xprime, max, lprime, right}
		}
	}
	return &node{min, max, left, right}
}

func joinRight(min, max int64, left, right *node) *node {
	if right != nil {
		xprime, yprime, rprime := splitMin(right.min, right.max, right.left, right.right)
		if max+1 == xprime {
			return &node{min, yprime, left, rprime}
		}
	}
	return &node{min, max, left, right}
}

func insert(x, y int64, d *node) *node {
	if d == nil {
		return &node{x, y, nil, nil}
	}
	switch {
	case x >= d.min && y <= d.max: // Contained within. Do nothing.
		return d

	case y < d.min: // Does not overlap. Is less.
		if y+1 == d.min {
			return joinLeft(x, d.max, d.left, d.right)
		}
		return &node{d.min, d.max, insert(x, y, d.left), d.right}

	case x > d.max: // Does not overlap. Is greater.
		if x == d.max+1 {
			return joinRight(d.min, y, d.left, d.right)
		}
		return &node{d.min, d.max, d.left, insert(x, y, d.right)}

	case x < d.min && y <= d.max: // Overlaps on the left
		return joinLeft(x, d.max, d.left, d.right)

	case x >= d.min && y > d.max: // Overlaps on the right
		return joinRight(d.min, y, d.left, d.right)

	case x < d.min && y > d.max: // Overlaps on left and right
		left := joinLeft(x, d.max, d.left, d.right)
		return joinRight(left.min, y, left.left, left.right)
	}
	return d
}

func intersection(l, r int64, d *node) int64 {
	if d == nil {
		return 0
	}
	if l > d.max {
		if d.right == nil {
			return 0
		}
		return intersection(l, r, d.right)
	}
	if r < d.min {
		if d.left == nil {
			return 0
		}
		return intersection(l, r, d.left)
	}
	if l >= d.min {
		if r <= d.max {
			return r - l + 1
		}
		isection := d.max - l + 1
		if d.right != nil {
			isection += intersection(d.max+1, r, d.right)
		}
		return isection
	}
	if r <= d.max {
		isection := r - d.min + 1
		if d.left != nil {
			isection += intersection(l, d.min-1, d.left)
		}
		return isection
	}
	if l <= d.min && r >= d.max {
		isection := d.max - d.min + 1
		if d.left != nil {
			isection += intersection(l, d.min-1, d.left)
		}
		if d.right != nil {
			isection += intersection(d.max+1, r, d.right)
		}
		return isection
	}
	return 0
}

func compress(root *node, count int) *node {
	var (
		child   *node
		scanner *node
		i       int
	)
	for i = 0; i < count; i++ {
		if scanner == nil {
			child = root
			root = child.right
		} else {
			child = scanner.right
			scanner.right = child.right
		}
		scanner = child.right
		child.right = scanner.left
		scanner.left = child
	}
	return root
}

// nearestPow2 calculates 2^(floor(log2(i)))
func nearestPow2(i int) int {
	r := 1
	for r <= i {
		r <<= 1
	}
	return r >> 1
}

func balance(root *node) *node {
	// Convert to a linked list
	tail := root
	rest := tail.right
	var size int
	for rest != nil {
		if rest.left == nil {
			tail = rest
			rest = rest.right
			size++
		} else {
			temp := rest.left
			rest.left = temp.right
			temp.right = rest
			rest = temp
			tail.right = temp
		}
	}
	// Now execute a series of rotations to balance
	leaves := size + 1 - nearestPow2(size+1)
	root = compress(root, leaves)
	size -= leaves
	for size > 1 {
		root = compress(root, size>>1)
		size >>= 1
	}
	// Return the new root
	return root
}

func intersectionAll(d *node, other *intervalTree) int64 {
	if d == nil {
		return 0
	}
	return other.Intersection(&interval{d.min, d.max}) + intersectionAll(d.left, other) + intersectionAll(d.right, other)
}

func total(d *node) int64 {
	if d == nil {
		return 0
	}
	return d.max - d.min + 1 + total(d.left) + total(d.right)
}
