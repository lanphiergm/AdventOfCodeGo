package utils

// Whether or not the map contains the specified coordinate
func ContainsCoord(coords map[Coord]struct{}, test Coord) bool {
	_, ok := coords[test]
	return ok
}
