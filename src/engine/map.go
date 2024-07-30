package engine

import (
	"bufio"
	"log"
	"os"
)

type TileType int

const (
	TileEmpty     TileType = 0
	TileOuterWall TileType = 1
	TileInnerWall TileType = 2
)

type Map [][]TileType

func ReadMap(filename string) Map {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var gameMap Map

	for scanner.Scan() {
		var row []TileType
		line := scanner.Text()
		for _, char := range line {
			switch char {
			case '0':
				row = append(row, TileEmpty)
			case '1':
				row = append(row, TileOuterWall)
			case '2':
				row = append(row, TileInnerWall)
			default:
				log.Fatalf("invalid character in map file: %c", char)
			}
		}
		gameMap = append(gameMap, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	return gameMap
}
