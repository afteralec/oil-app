package rooms

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"petrichormud.com/app/internal/shared"
	"petrichormud.com/app/internal/tests"
)

func TestGraph(t *testing.T) {
	i := shared.SetupInterfaces()
	defer i.Close()

	testRIDs := []int64{}
	for n := 0; n < 15; n++ {
		rid := tests.CreateTestRoom(t, &i, tests.TestRoom)
		testRIDs = append(testRIDs, rid)
		defer tests.DeleteTestRoom(t, &i, rid)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[0],
		To:        testRIDs[1],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}
	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionEast,
		ID:        testRIDs[0],
		To:        testRIDs[4],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorthwest,
		ID:        testRIDs[1],
		To:        testRIDs[2],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}
	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[1],
		To:        testRIDs[3],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionEast,
		ID:        testRIDs[2],
		To:        testRIDs[5],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}
	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionEast,
		ID:        testRIDs[2],
		To:        testRIDs[6],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[3],
		To:        testRIDs[6],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	room, err := i.Queries.GetRoom(context.Background(), testRIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	graph, err := BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	node := &graph
	require.Equal(t, testRIDs[1], node.GetExit(DirectionNorth).ID)
	require.True(t, node.IsExitEmpty(DirectionSoutheast))

	node = node.GetExit(DirectionNorth)
	require.Equal(t, testRIDs[2], node.GetExit(DirectionNorthwest).ID)
	require.Equal(t, testRIDs[3], node.GetExit(DirectionNorth).ID)

	node = node.GetExit(DirectionNorthwest)
	require.Equal(t, int64(0), node.GetExit(DirectionNorth).ID)
	require.True(t, node.GetExit(DirectionNorth).IsExitEmpty(DirectionSouth))
	require.Equal(t, testRIDs[6], node.GetExit(DirectionEast).ID)
}

func TestNodeGetExit(t *testing.T) {
	i := shared.SetupInterfaces()
	defer i.Close()

	rid := tests.CreateTestRoom(t, &i, tests.TestRoom)
	defer tests.DeleteTestRoom(t, &i, rid)
	room, err := i.Queries.GetRoom(context.Background(), rid)
	if err != nil {
		t.Fatal(err)
	}

	node, err := BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range DirectionsList {
		require.Equal(t, int64(0), node.GetExit(dir).ID)
	}

	ridTwo := tests.CreateTestRoom(t, &i, tests.TestRoom)
	defer tests.DeleteTestRoom(t, &i, ridTwo)

	if err := Link(LinkParams{
		Queries:   i.Queries,
		ID:        rid,
		To:        ridTwo,
		Direction: DirectionSouth,
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	room, err = i.Queries.GetRoom(context.Background(), rid)
	if err != nil {
		t.Fatal(err)
	}

	node, err = BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, ridTwo, node.GetExit(DirectionSouth).ID)
}

func TestNodeIsExitEmpty(t *testing.T) {
	i := shared.SetupInterfaces()
	defer i.Close()

	rid := tests.CreateTestRoom(t, &i, tests.TestRoom)
	defer tests.DeleteTestRoom(t, &i, rid)
	room, err := i.Queries.GetRoom(context.Background(), rid)
	if err != nil {
		t.Fatal(err)
	}

	node, err := BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range DirectionsList {
		require.True(t, node.IsExitEmpty(dir))
	}

	ridTwo := tests.CreateTestRoom(t, &i, tests.TestRoom)
	defer tests.DeleteTestRoom(t, &i, ridTwo)

	if err := Link(LinkParams{
		Queries:   i.Queries,
		ID:        rid,
		To:        ridTwo,
		Direction: DirectionSouth,
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	room, err = i.Queries.GetRoom(context.Background(), rid)
	if err != nil {
		t.Fatal(err)
	}

	node, err = BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	require.False(t, node.IsExitEmpty(DirectionSouth))
}

func TestEmptyBindMatrix(t *testing.T) {
	matrix := EmptyBindMatrix()
	require.Equal(t, 5, len(matrix))
	require.Equal(t, 5, len(matrix[0]))
}

func TestIsValidMatrixCoordinate(t *testing.T) {
	matrix := EmptyBindMatrix()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			require.True(t, IsValidMatrixCoordinate(matrix, i, j))
		}
	}

	require.False(t, IsValidMatrixCoordinate(matrix, 5, 4))
	require.False(t, IsValidMatrixCoordinate(matrix, 4, 5))
}

func TestBindMatrix(t *testing.T) {
	i := shared.SetupInterfaces()
	defer i.Close()

	testRIDs := []int64{}
	for n := 0; n < 15; n++ {
		rid := tests.CreateTestRoom(t, &i, tests.TestRoom)
		testRIDs = append(testRIDs, rid)
		defer tests.DeleteTestRoom(t, &i, rid)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[0],
		To:        testRIDs[1],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}
	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionEast,
		ID:        testRIDs[0],
		To:        testRIDs[4],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorthwest,
		ID:        testRIDs[1],
		To:        testRIDs[2],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}
	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[1],
		To:        testRIDs[3],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionEast,
		ID:        testRIDs[2],
		To:        testRIDs[5],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	if err := Link(LinkParams{
		Queries:   i.Queries,
		Direction: DirectionNorth,
		ID:        testRIDs[3],
		To:        testRIDs[6],
		TwoWay:    true,
	}); err != nil {
		t.Fatal(err)
	}

	room, err := i.Queries.GetRoom(context.Background(), testRIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	graph, err := BuildGraph(BuildGraphParams{
		Queries:  i.Queries,
		Room:     &room,
		Depth:    0,
		MaxDepth: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	matrix := graph.BindMatrix(EmptyBindMatrix(), 2, 2)
	for _, row := range matrix {
		rowIDs := []int64{}
		for _, room := range row {
			rowIDs = append(rowIDs, room["ID"].(int64))
		}
		log.Println(rowIDs)
	}
}
