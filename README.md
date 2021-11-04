# Senmai

A Wrapper Client for Google Spreadsheet API (Sheets API)

## PREPARATION

### Service Account and Key File

1. Create a service account on Google Cloud Platform. Refer: https://cloud.google.com/iam/docs/creating-managing-service-accounts
2. Generate and Download a key file (json) that relates the service account you created. Refer: https://cloud.google.com/iam/docs/creating-managing-service-account-keys

### Sheets API

1. Enable Sheets API on your project. Refer: https://console.developers.google.com/apis/api/sheets.googleapis.com/overview

### Spreadsheet

1. Create a spreadsheet on your google drive.
2. Specify your table schema on the spreadsheet. Row '1' is column name cell. And rows '2' or greater are data cell. Cell '1A' must be 'id'.
3. Invite the service account as a editor, without notification.

## SYNOPSIS

### Init Senmai Client

```
	// KeyFile path of your key file you generated at PREPARATION section 
	const KeyFile = "/path/to/identity.json"

	// SpreadsheetID is identity string of spreadsheet that you want to manipulate. Refer: https://developers.google.com/sheets/api/guides/concepts
	const SpreadsheetID = "1234567890qwertyuiopASDFGHJKL"

	// SheetName is a name of a sheet you want to manipulate 
	SheetName = "book"

	// Init Senmai Client
	sm, err := senmai.New(KeyFile)
	if err != nil {
		log.Fatal(err.Error())
	}
```

### Get Table struct

```
	// Get Table struct that named "book"
	book, err := sm.GetTable(SpreadsheetID, SheetName)
	if err != nil {
		log.Fatal(err.Error())
	}
```

### Fetch a row

```
	// Fetch
	row, err := book.FetchRow("mybook0001")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", row)
```

### Update a row

```
	// Update
	row.Cols["price_yen"] = "4000"
	err = row.Update()
	if err != nil {
		log.Fatal(err.Error())
	}
```

### Create a new row

```
	// Build a new row struct
	b1 := book.NewRow()

	// Fill-in field values
	b1.Cols = map[string]string{
		"id":         "mybook0003",
		"name":       "Slack: Getting Past Burnout, Busywork, and the Myth of Total Efficiency",
		"author":     "Tom DeMarco",
		"price_yen":  "2800",
	}

	// Insert into spreadsheet
	err = b1.Insert()
	if err != nil {
		log.Fatal(err.Error())
	}
```

