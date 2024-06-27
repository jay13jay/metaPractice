package spiral

//
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    m := len(matrix) // number of rows
    n := len(matrix[0]) // number of columns
    size := m * n // total number of elements
    ret := make([]int, size) // result slice

    c := newCrawler(m, n) // create a new crawler to traverse the matrix

    // traverse matrix: set current position value, then move, until
    // whole matrix is traversed
    for i := 0; i < size; i++ {
        ret[i] = matrix[c.Pos[0]][c.Pos[1]]
        c.move()
    }

    return ret
}

type Crawler struct {
    Size   int       // Total values in the matrix
    CurDir int       // Current direction
    Pos    [2]int    // Current position
    Dirs   [4][2]int // Array of directions
    lBound [2]int    // Lower bound
    hBound [2]int    // Higher bound
}

func newCrawler(m, n int) *Crawler {
    return &Crawler{
        Size:   m * n,
        CurDir: 0,
        Pos:    [2]int{0, 0},
        Dirs: [4][2]int{
            {0, 1},  // right
            {1, 0},  // down
            {0, -1}, // left
            {-1, 0}, // up
        },
        lBound: [2]int{0, 0},
        hBound: [2]int{m - 1, n - 1},
    }
}

func (c *Crawler) move() {
    nextRow := c.Pos[0] + c.Dirs[c.CurDir][0]
    nextCol := c.Pos[1] + c.Dirs[c.CurDir][1]

    // check if the next move would put the coordinates
    // outside of the matrix' bounds
    if nextRow < c.lBound[0] || nextRow > c.hBound[0] ||
        nextCol < c.lBound[1] || nextCol > c.hBound[1] {
        // adjust the boundaries and change direction
        switch c.CurDir {
        case 0: // right
            c.lBound[0]++
        case 1: // down
            c.hBound[1]--
        case 2: // left
            c.hBound[0]--
        case 3: // up
            c.lBound[1]++
        }
        c.CurDir = (c.CurDir + 1) % 4 // rotate direction
    }

    // move to the next coordinate
    c.Pos[0] += c.Dirs[c.CurDir][0]
    c.Pos[1] += c.Dirs[c.CurDir][1]
}

func Solve(matrix [][]int) []int {
    return spiralOrder(matrix)
}
