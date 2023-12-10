package day_10

import (
	"aoc-2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Direction string
type TileContent string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

func (d *Direction) opposite() Direction {
	switch *d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}

	return ""
}

const (
	UpDown    TileContent = "|"
	LeftRight TileContent = "-"
	UpRight   TileContent = "L"
	UpLeft    TileContent = "J"
	DownRight TileContent = "F"
	DownLeft  TileContent = "7"
	Ground    TileContent = "."
	Start     TileContent = "S"
)

type Maze struct {
	box   utils.Box
	tiles [][]Tile
}

func (m *Maze) getTile(position Position, direction Direction) (Tile, error) {
	if !position.canMove(direction, m.box) {
		return Tile{}, fmt.Errorf("cannot move %v for maze %v", direction, m.box)
	}

	x := position.point.X
	y := position.point.Y

	switch direction {
	case Up:
		y--
	case Down:
		y++
	case Left:
		x--
	case Right:
		x++
	}

	return m.tiles[y][x], nil
}

type Tile struct {
	content     TileContent
	connections [2]Direction
	position    Position
}

func parseTile(char TileContent, position Position) Tile {
	var direction [2]Direction

	switch char {
	case UpDown:
		direction = [2]Direction{Up, Down}
	case LeftRight:
		direction = [2]Direction{Left, Right}
	case UpRight:
		direction = [2]Direction{Up, Right}
	case UpLeft:
		direction = [2]Direction{Up, Left}
	case DownRight:
		direction = [2]Direction{Down, Right}
	case DownLeft:
		direction = [2]Direction{Down, Left}
	}

	return Tile{char, direction, position}
}

type Position struct {
	point utils.Point
}

func (p *Position) canMove(direction Direction, box utils.Box) bool {
	switch direction {
	case Up:
		return box.IsWithin(utils.Point{X: p.point.X, Y: p.point.Y - 1})
	case Down:
		return box.IsWithin(utils.Point{X: p.point.X, Y: p.point.Y + 1})
	case Left:
		return box.IsWithin(utils.Point{X: p.point.X - 1, Y: p.point.Y})
	case Right:
		return box.IsWithin(utils.Point{X: p.point.X + 1, Y: p.point.Y})
	}

	return false
}

type Cursor struct {
	// pos x, y
	tile Tile
}

func (t *Tile) canMove(direction Direction, maze Maze) (bool, error) {
	sibling, err := maze.getTile(t.position, direction)
	if err != nil {
		return false, err
	}

	if sibling.content == Ground {
		return false, nil
	}

	if sibling.content == Start {
		return true, nil
	}

	for _, connection := range sibling.connections {
		if connection == direction.opposite() {
			return true, nil
		}
	}

	return false, nil
}

type AreaTile string

const (
	AreaTileGround AreaTile = "."
	AreaTileWall   AreaTile = "#"
	AreaTileNest   AreaTile = "N"
)

type Area struct {
	points    [][]utils.Point
	areaTiles [][]AreaTile
}

func (a *Area) sortPoints() {
	for _, row := range a.points {
		for i := 0; i < len(row)-1; i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i].X > row[j].X {
					row[i], row[j] = row[j], row[i]
				}
			}
		}
	}
}

func (a *Area) getPointsByColumns(width int) [][]utils.Point {
	byColumns := make([][]utils.Point, width)
	for i := range byColumns {
		byColumns[i] = []utils.Point{}
	}

	for _, row := range a.points {
		for _, point := range row {
			byColumns[point.X] = append(byColumns[point.X], point)
		}
	}

	return byColumns
}

func (a *Area) buildAreaTile(width int) {
	a.areaTiles = make([][]AreaTile, len(a.points))
	for i := range a.areaTiles {
		a.areaTiles[i] = make([]AreaTile, width)
		for j := range a.areaTiles[i] {
			a.areaTiles[i][j] = AreaTileGround
		}
	}

	for y, row := range a.points {
		for _, point := range row {
			a.areaTiles[y][point.X] = AreaTileWall
		}
	}
}

/* func safePoint(point, max int) int {
	if point < 0 {
		return 0
	}

	if point >= max {
		return max - 1
	}

	return point
} */

func intersectionCount(point int, wallIndices [][]int) int {
	count := 0
	for _, wallIndex := range wallIndices {
		if point <= wallIndex[0] && point <= wallIndex[1] {
			count++
		}
	}

	return count % len(wallIndices)
}

func (a *Area) getTilesInsideArea(boxWidth int) int {
	a.sortPoints()
	count := 0
	// max := len(a.areaTiles[0])
	byColumns := a.getPointsByColumns(boxWidth)
	fmt.Println(byColumns)

	re := regexp.MustCompile(`(#+)`)
	for y, row := range a.areaTiles {

		stringContent := make([]string, len(row))
		for x, tile := range row {
			stringContent[x] = string(tile)
		}

		wallIndices := re.FindAllStringIndex(strings.Join(stringContent, ""), -1)
		for x, tile := range row {
			isWall := tile == AreaTileWall

			intersectionCount := intersectionCount(x, wallIndices)

			if len(byColumns[x]) == 0 {
				fmt.Printf("no point in column %v\n", x)
				continue
			}

			pastFirstWallVertical := y > byColumns[x][0].Y
			pastLastWallVertical := y > byColumns[x][len(byColumns[x])-1].Y

			if !isWall && intersectionCount == 1 {
				if pastFirstWallVertical && !pastLastWallVertical {
					a.areaTiles[y][x] = AreaTileNest
					count++
					// }
				}

			}
		}
	}

	return count
}

func getPossibleDirections(tile Tile, from Direction, maze Maze) []Direction {
	directions := []Direction{}

	if tile.content == Ground || tile.content == Start {
		for _, direction := range []Direction{Up, Down, Left, Right} {
			canMove, _ := tile.canMove(direction, maze)
			if canMove {
				directions = append(directions, direction)
			}
		}
	}

	for _, connection := range tile.connections {
		if connection != from {
			directions = append(directions, connection)
		}
	}

	fmt.Printf("Tile: %v, from %v, directions: %v\n", tile, from, directions)

	return directions
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	fileContent := [][]Tile{}
	result := 0
	var startPosition Position

	for scanner.Scan() {
		line := scanner.Text()
		lineContent := []Tile{}

		for i, char := range strings.Split(line, "") {
			position := Position{utils.Point{X: i, Y: len(fileContent)}}

			tile := parseTile(TileContent(char), position)
			lineContent = append(lineContent, tile)

			if tile.content == Start {
				startPosition = position
			}
		}

		fileContent = append(fileContent, lineContent)
	}

	box := utils.Box{Length: len(fileContent), Width: len(fileContent[0])}
	maze := Maze{box, fileContent}
	startTile := fileContent[startPosition.point.Y][startPosition.point.X]
	startDirections := getPossibleDirections(startTile, "", maze)

	fmt.Printf("START HERE: %v\n\n\n", startTile)
	yArea := make([][]utils.Point, box.Length)
	area := Area{points: yArea}
	area.points[startPosition.point.Y] = append(area.points[startPosition.point.Y], startPosition.point)

	direction := startDirections[1]
	nextTile, _ := maze.getTile(startPosition, direction)
	cursor := Cursor{nextTile}
	area.points[cursor.tile.position.point.Y] = append(area.points[cursor.tile.position.point.Y], cursor.tile.position.point)

	from := direction.opposite()
	count := 1

	for cursor.tile.content != Start {
		fmt.Printf("Cursor: %v, \n", cursor)
		directions := getPossibleDirections(cursor.tile, from, maze)
		fmt.Printf("Directions: %v\n\n\n", directions)

		if len(directions) == 0 {
			break
		}

		if len(directions) == 1 {
			direction := directions[0]
			from = direction.opposite()

			tile, err := maze.getTile(cursor.tile.position, direction)
			if err != nil {
				panic(err)
			}

			cursor.tile = tile
			area.points[cursor.tile.position.point.Y] = append(area.points[cursor.tile.position.point.Y], cursor.tile.position.point)
		}

		count++
	}

	area.buildAreaTile(box.Width)
	result = area.getTilesInsideArea(box.Width)

	for _, row := range area.areaTiles {
		fmt.Println(row)
	}

	return strconv.Itoa(result)

}
