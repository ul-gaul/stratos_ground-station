package controller

import (
	"github.com/ul-gaul/stratos_ground-station/backend/acquisition"
	"github.com/ul-gaul/stratos_ground-station/backend/acquisition/packet"
	"github.com/ul-gaul/stratos_ground-station/backend/acquisition/persistence/read"
	"github.com/ul-gaul/stratos_ground-station/backend/cmd"
)

/*
// Add frontend functions here.
// Example:

func (c *controller) Hello() string {
    return "World!"
}

// javascript
backend.Hello().then(console.log) // => "World!"
*/

func (c *controller) Connect() ([]packet.TransformedData, error) {
	csv := cmd.GetCSVFile()
	if csv != "" {
		return read.Csv(csv)
	}
	return nil, nil
}

func (c *controller) Data() []packet.TransformedData {
	return acquisition.TestData()
}

func (c *controller) OpenDataFile() ([]packet.TransformedData, error) {
	filename := c.r.Dialog.SelectFile("Select a file", "*.csv")

	if filename != "" {
		return read.Csv(filename)
	}

	return nil, nil
}
