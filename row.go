package senmai

import "fmt"

// Row row struct
type Row struct {
	t     *Table
	s     *Senmai
	index int
	Cols  map[string]string
}

// Insert insert this row into sheet
func (row Row) Insert() error {
	cnt := row.t.findVoidRowIndex()
	row.index = cnt
	return row.writeValues(false)
}

// Update update this row
func (row Row) Update() error {
	return row.writeValues(true)
}

func (row Row) writeValues(allowDuplicate bool) error {
	values := []interface{}{}
	for _, k := range row.t.Fields {
		v, ok := row.Cols[k]
		if !ok {
			v = ""
		}
		values = append(values, v)
	}
	id := row.Cols["id"]

	if !allowDuplicate {
		dup, _ := row.t.FetchRow(id)
		if dup != nil {
			return fmt.Errorf("duplicate entry: id=%s", id)
		}
	}

	row.t.t.PutValuesAtRow(row.index, values...)
	return row.t.commit()
}
