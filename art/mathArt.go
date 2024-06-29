package art

import (
	"fmt"
	"sort"
)

type Event struct {
  coord       int32
  isStart     bool
  isHorizontal bool
  index       int32
}

type Drawing struct {
  Strokes    int32
  Vectors    []int32
  Directions string
  PlusSigns  int64
  X          int32
  Y          int32
  xKeys      []int32
  yKeys      []int32
  xCheck     map[int32]bool
  yCheck     map[int32]bool
  vertLines  map[int32][2]int32
  horLines   map[int32][2]int32
}

func newDrawing(N int32, L []int32, D string) *Drawing {
  return &Drawing{
    PlusSigns:  0,
    Strokes:    N,
    Vectors:    L,
    Directions: D,
    X:          0,
    Y:          0,
    xCheck:     make(map[int32]bool),
    yCheck:     make(map[int32]bool),
    vertLines:  make(map[int32][2]int32),
    horLines:   make(map[int32][2]int32),
  }
}

func (d *Drawing) checkIntersections() {
  events := make([]Event, 0, 2*len(d.xKeys)+2*len(d.yKeys))

  // Create events for horizontal lines
  for y, line := range d.horLines {
    events = append(events, Event{coord: line[0], isStart: true, isHorizontal: true, index: y})
    events = append(events, Event{coord: line[1], isStart: false, isHorizontal: true, index: y})
  }

  // Create events for vertical lines
  for x, line := range d.vertLines {
    events = append(events, Event{coord: x, isStart: true, isHorizontal: false, index: line[0]})
    events = append(events, Event{coord: x, isStart: false, isHorizontal: false, index: line[1]})
  }

  // Sort events by coordinate, breaking ties by isStart (start before end)
  sort.Slice(events, func(i, j int) bool {
    if events[i].coord == events[j].coord {
      return events[i].isStart && !events[j].isStart
    }
    return events[i].coord < events[j].coord
  })

  activeIntervals := make(map[int32]int32)

  // Process events
  for _, event := range events {
    if event.isHorizontal {
      if event.isStart {
        activeIntervals[event.index]++
      } else {
        activeIntervals[event.index]--
        if activeIntervals[event.index] == 0 {
          delete(activeIntervals, event.index)
        }
      }
    } else {
      if event.isStart {
        for interval := range activeIntervals {
          if interval > event.index && interval < d.vertLines[event.coord][1] {
						fmt.Printf("Adding intersection. interval: %d\n", interval)
            d.PlusSigns++
          }
        }
      }
    }
  }
}

func (d *Drawing) updateMaps() {
  if !d.xCheck[d.X] {
    d.xCheck[d.X] = true
    d.vertLines[d.X] = [2]int32{d.Y, d.Y}
    d.xKeys = append(d.xKeys, d.X)
  } else {
    line := d.vertLines[d.X]
    if d.Y < line[0] {
      line[0] = d.Y
    } else if d.Y > line[1] {
      line[1] = d.Y
    }
    d.vertLines[d.X] = line
  }

  if !d.yCheck[d.Y] {
    d.yCheck[d.Y] = true
    d.horLines[d.Y] = [2]int32{d.X, d.X}
    d.yKeys = append(d.yKeys, d.Y)
  } else {
    line := d.horLines[d.Y]
    if d.X < line[0] {
      line[0] = d.X
    } else if d.X > line[1] {
      line[1] = d.X
    }
    d.horLines[d.Y] = line
  }
}

func (d *Drawing) setNewCoord(index int) {
  dir := string(d.Directions[index])
  vector := d.Vectors[index]
  if dir == "U" {
    d.Y += vector
    d.updateMaps()
  } else if dir == "L" {
    d.X -= vector
    d.updateMaps()
  } else if dir == "R" {
    d.X += vector
    d.updateMaps()
  } else if dir == "D" {
    d.Y -= vector
    d.updateMaps()
  } else {
    fmt.Println("Invalid direction")
  }
}

func getPlusSignCount(N int32, L []int32, D string) int64 {
  d := newDrawing(N, L, D)
  d.yCheck[0] = true
  d.xCheck[0] = true
  d.xKeys = append(d.xKeys, 0)
  d.yKeys = append(d.yKeys, 0)

  for i := range d.Vectors {
    d.setNewCoord(i)
  }

  d.checkIntersections()

  return d.PlusSigns
}

func Solve(N int32, L []int32, D string) int64 {
  return getPlusSignCount(N, L, D)
}