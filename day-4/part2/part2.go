package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid struct {
    Cells [][]rune
    Rows int
    Cols int
}

// function that counts adjacent @ values 
func (g *Grid) CountNeighbors(r, c int) (bool, int)  {
    directions := []struct{dr, dc int}{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},          {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }

    count := 0
    for _, d := range directions {
        nr, nc := r+d.dr, c+d.dc
        if nr >= 0 && nr < g.Rows && nc >= 0 && nc < g.Cols {
            if g.Cells[nr][nc] == '@' {
                count++
            }
        }
    }

    if count <= 3 {
        return true, count
    }
    return false, count
}

// set '@' at given indices to '.' based on indices struct slice
func (g *Grid) RemoveAtIndices(indices []struct{row, col int}) {
    for _, idx := range indices {
        g.Cells[idx.row][idx.col] = '.'
    }
    fmt.Println("Grid after removal:")
    for r := 0; r < g.Rows; r++ {
        for c := 0; c < g.Cols; c++ {
            fmt.Printf("%c", g.Cells[r][c])
        }
        fmt.Println()
    }
    fmt.Println()
}

// Performs a single pass over the grid, counting @ with 3 or fewer neighbors.
// Returns the count and the list of their indices.
func singleGridPass(grid Grid) (int, []struct{row, col int}) {
    sumOk := 0
    indicesToRemove := []struct{row, col int}{}

    fmt.Println("Initial grid:")
    for r := 0; r < grid.Rows; r++ {
        for c := 0; c < grid.Cols; c++ {
            fmt.Printf("%c", grid.Cells[r][c])
            if grid.Cells[r][c] == '@' {
                ok, _ := grid.CountNeighbors(r, c)
                if ok {
                    sumOk++
                    indicesToRemove = append(indicesToRemove, struct{row, col int}{r, c})
                }
            }
        }
        fmt.Println()
    }
    fmt.Println()

    return sumOk, indicesToRemove
}

func readGridFromFile(path string) (Grid, error) {
    file, err := os.Open(path)
    if err != nil {
        return Grid{}, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var cells [][]rune

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text()) // string with the line content
        //fmt.Printf("Read line: %+v\n", line)
        if len(line) == 0 {
            continue
        }
        row := []rune(line) // Convert string to slice of runes
        cells = append(cells, row)
    }

    if err := scanner.Err(); err != nil {
        return Grid{}, err
    }

    if len(cells) == 0 {
        return Grid{}, nil
    }

    grid := Grid{
        Cells: cells,
        Rows:  len(cells),  
        Cols:  len(cells[0]),
    }

    return grid, nil
}

func Part2() {
    grid, err := readGridFromFile("input.txt")
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }
//    fmt.Printf("Grid loaded: %d rows, %d cols\n", grid.Rows, grid.Cols)
//    fmt.Printf("Grid content: %+v\n", grid.Cells)
    totalOk := 0
    for {
        sumOk, indicesToRemove := singleGridPass(grid)
        if sumOk == 0 {
            break
        }
        totalOk += sumOk
        grid.RemoveAtIndices(indicesToRemove)
    }

    fmt.Println()
    fmt.Printf("Total @ with 3 or fewer neighbors: %d\n", totalOk)

}

