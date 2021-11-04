package main

import (
	"fmt"

	"github.com/ytnobody/herschel"
)

// Table table struct
type Table struct {
	t             *herschel.Table
	s             *Senmai
	Fields        []string
	SpreadsheetID string
	SheetName     string
}

// FetchRow fetch a row that exact matched to 'A' column
func (t Table) FetchRow(id string) (*Row, error) {
	idx := t.t.IndexOfRowWithPrefix(id)
	if idx < 0 {
		return nil, fmt.Errorf("could not found any row that have an id:%s", id)
	}
	vs := t.getStringValuesAtRow(idx)
	row := t.NewRow()
	row.index = idx
	for i, f := range t.Fields {
		row.Cols[f] = vs[i]
	}
	return row, nil
}

// GetRows return row count of table
func (t Table) GetRows() int {
	return t.t.GetRows()
}

// NewRow build and return a new Row struct that relates this table
func (t Table) NewRow() *Row {
	cols := map[string]string{}
	for _, f := range t.Fields {
		cols[f] = ""
	}
	row := &Row{
		t:    &t,
		s:    t.s,
		Cols: cols,
	}
	return row
}

func (t Table) getStringValuesAtRow(row int) []string {
	vs := t.t.GetValuesAtRow(row)
	values := []string{}
	for _, v := range vs {
		values = append(values, fmt.Sprintf("%s", v))
	}
	return values
}

func (t Table) findVoidRowIndex() int {
	i := t.t.IndexOfRowWithPrefix("")
	if i < 0 {
		i = t.GetRows()
	}
	return i
}

func (t Table) commit() error {
	return t.s.Client.WriteTable(t.SpreadsheetID, t.SheetName, t.t)
}
