package controller

import (
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/packet"
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


func (c *controller) Connect() {
    acquisition.SendTestData()
}

func (c *controller) Data() []packet.TransformedData {
    return acquisition.TestData()
}