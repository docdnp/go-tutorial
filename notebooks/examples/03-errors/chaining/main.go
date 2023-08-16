package main

import (
	"database/sql"
	"errors"
	"fmt"

	chain "github.com/g8rswimmer/error-chain"
)

type Entity struct{}

func scanRow(row *sql.Row) (*Entity, error) {
	return nil, sql.ErrNoRows
}

func query(id string) *sql.Row {
	return &sql.Row{}
}

func FetchByID(id string) (*Entity, error) {
	row := query(id)
	e, err := scanRow(row)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		ec := chain.New()
		ec.Add(fmt.Errorf("entity doesn't exist"))
		ec.Add(fmt.Errorf("entity fetch error: %s", id))
		ec.Add(err)
		return nil, ec
	case err != nil:
		return nil, fmt.Errorf("entity fetch err: %s %w", id, err)
	default:
		return e, nil
	}
}

func main() {
	_, err := FetchByID("DummyID")
	fmt.Println(err)
}
