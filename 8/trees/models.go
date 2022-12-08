package trees

// Map
type Map struct {
	Width  int
	Height int
	Trees  [][]*Tree
}

type Tree struct {
	Height  uint
	Top     bool
	Bottom  bool
	Left    bool
	Right   bool
	VTop    int
	VBottom int
	VLeft   int
	VRight  int
}

func (t *Tree) IsVisible() bool {
	return t.Top || t.Bottom || t.Left || t.Right
}
func (t *Tree) CalcScore() int {
	return t.VTop * t.VBottom * t.VLeft * t.VRight
}

// NewMap creates a new map
func NewMap() *Map {
	return &Map{}
}

func (m *Map) Check(i, j int) bool {
	t := m.Trees[i][j]
	if t.IsVisible() {
		return true
	}

	// Check top
	if m.CheckTop(i, j, t.Height) {
		t.Top = true
		return true
	}
	// Check left
	if m.CheckLeft(i, j, t.Height) {
		t.Left = true
		return true
	}
	// Check bottom
	if m.CheckBottom(i, j, t.Height) {
		t.Bottom = true
		return true
	}
	// Check right
	if m.CheckRight(i, j, t.Height) {
		t.Right = true
		return true
	}

	return false
}

func (m *Map) CheckTop(i, j int, h uint) bool {
	for k := i - 1; k >= 0; k-- {
		t := m.Trees[k][j]
		if t.Height >= h {
			return false
		}
		if t.Top {
			return true
		}
	}
	return true
}

func (m *Map) CheckLeft(i, j int, h uint) bool {
	for k := j - 1; k >= 0; k-- {
		t := m.Trees[i][k]
		if t.Height >= h {
			return false
		}
		if t.Left {
			return true
		}
	}
	return true
}

func (m *Map) CheckBottom(i, j int, h uint) bool {
	for k := i + 1; k < m.Height; k++ {
		t := m.Trees[k][j]
		if t.Height >= h {
			return false
		}
		if t.Bottom {
			return true
		}
	}
	return true
}

func (m *Map) CheckRight(i, j int, h uint) bool {
	for k := j + 1; k < m.Width; k++ {
		t := m.Trees[i][k]
		if t.Height >= h {
			return false
		}
		if t.Right {
			return true
		}
	}
	return true
}

func (m *Map) CalcVisibilityScore(i, j int) int {
	t := m.Trees[i][j]

	// Calc v top
	t.VTop = m.CalcVisibilityScoreTop(i, j, t.Height)
	// Calc v left
	t.VLeft = m.CalcVisibilityScoreLeft(i, j, t.Height)
	// Calc v bottom
	t.VBottom = m.CalcVisibilityScoreBottom(i, j, t.Height)
	// Calc v right
	t.VRight = m.CalcVisibilityScoreRight(i, j, t.Height)
	//fmt.Println(t.VTop, t.VLeft, t.VBottom, t.VRight)
	return t.CalcScore()
}

func (m *Map) CalcVisibilityScoreTop(i, j int, h uint) int {
	v := 0
	for k := i - 1; k >= 0; k-- {
		t := m.Trees[k][j]
		v += 1
		if t.Height >= h {
			return v
		}
	}
	return v
}

func (m *Map) CalcVisibilityScoreLeft(i, j int, h uint) int {
	v := 0
	for k := j - 1; k >= 0; k-- {
		t := m.Trees[i][k]
		v += 1
		if t.Height >= h {
			return v
		}
	}
	return v
}

func (m *Map) CalcVisibilityScoreBottom(i, j int, h uint) int {
	v := 0
	for k := i + 1; k < m.Height; k++ {
		t := m.Trees[k][j]
		if t.Height >= h {
			return v
		}
		v += 1
	}
	return v
}

func (m *Map) CalcVisibilityScoreRight(i, j int, h uint) int {
	v := 0
	for k := j + 1; k < m.Width; k++ {
		t := m.Trees[i][k]
		v += 1
		if t.Height >= h {
			return v
		}
	}
	return v
}
