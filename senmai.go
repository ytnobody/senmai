package main

import (
	"fmt"

	"github.com/ytnobody/herschel"
	"github.com/ytnobody/herschel/option"
)

// Senmai A wrapper client for Google Sheets API
type Senmai struct {
	Client *herschel.Client
}

// New init a Senmai wrapper client
func New(credentialsFile string) (*Senmai, error) {
	option := option.WithServiceAccountCredentials(credentialsFile)
	client, err := herschel.NewClient(option)
	if err != nil {
		return nil, err
	}
	gst := &Senmai{Client: client}
	return gst, nil
}

// GetTable build and return table struct
func (s Senmai) GetTable(spreadsheetID string, sheetName string) (*Table, error) {
	t, err := s.Client.ReadTable(spreadsheetID, sheetName)
	if err != nil {
		return nil, err
	}
	fs := t.GetValuesAtRow(0)
	fields := []string{}
	for _, f := range fs {
		v := fmt.Sprintf("%s", f)
		fields = append(fields, v)
	}
	table := &Table{
		t:             t,
		s:             &s,
		SpreadsheetID: spreadsheetID,
		SheetName:     sheetName,
		Fields:        fields,
	}
	return table, nil
}
