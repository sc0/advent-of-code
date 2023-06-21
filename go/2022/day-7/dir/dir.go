package dir

type Dir struct {
	Tree      map[string]*Dir
	Parent    *Dir
	FilesSize int
}

func (d Dir) CalculateSize() int {
	result := d.FilesSize

	for _, dir := range d.Tree {
		result += dir.CalculateSize()
	}

	return result
}

func (d Dir) Part1() int {
	thresh := 100000
	result := 0
	size := d.CalculateSize()
	if size < thresh {
		result += size
	}

	for _, dir := range d.Tree {
		result += dir.Part1()
	}

	return result
}

func (d Dir) Part2(freeSpace int) int {
	req := 30_000_000
	taken := d.CalculateSize()
	target := freeSpace + taken

  var candidates []int

	if target >= req {
    candidates = append(candidates, taken)
	}

	for _, dir := range d.Tree {
		size := dir.Part2(freeSpace)
		if target + size >= req {
			candidates = append(candidates, size)
		}
	}

	if len(candidates) == 0 {
		return 99_999_999_999
	}

	min := candidates[0]
	for _, c := range candidates {
		if c < min {
			min = c
		}
	}

	return min
}

func New(parent *Dir) Dir {
	return Dir{make(map[string]*Dir), parent, 0}
}
