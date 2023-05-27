# replacements
Simple replacement strings for Go.  The old school way of doing parameterized string replacements.  The InfluxDB client for Go doesn't support params (at least not for the OSS version of the server).  While I created this for Flux queries, this code is useful for the generic problem of needing to do string replacements based on some settings and/or environment variables.

## Example
This is an example of how to create a reusable query for InfluxDB.  I use it to get the data to be able to calculate the efficiency of my dehumidifier.

```go
package main

import (
	"fmt"
	"time"

	"github.com/neel-m/replacements"
)

func main() {
	var diffQuery = `
data = from(bucket: "##BUCKET##")
	|> range(start: ##START##, stop: ##STOP##)
	|> filter(fn: (r) => r["NodeName"] == "##NODENAME##")
	|> filter(fn: (r) => r["_measurement"] == "##MEASUREMENT##")
	|> filter(fn: (r) => r["_field"] == "##FIELD##")

data
	|> first()
	|> yield(name: "first")

data
	|> last()
	|> yield(name: "last")
	`
	now := time.Now().UTC()
	start := now.Add(-time.Hour * 24)
	params := map[string]string{
		"START":       start.Format(time.RFC3339),
		"STOP":        now.Format(time.RFC3339),
		"BUCKET":      "year",
		"NODENAME":    "dehumid",
		"MEASUREMENT": "sensor",
		"FIELD":       "Total",
	}
	queryString := replacements.ReplacePlaceholders(diffQuery, params)
	fmt.Println(queryString)
}

```
This outputs
```
data = from(bucket: "year")
        |> range(start: 2023-05-26T14:23:08Z, stop: 2023-05-27T14:23:08Z)
        |> filter(fn: (r) => r["NodeName"] == "dehumid")
        |> filter(fn: (r) => r["_measurement"] == "sensor")
        |> filter(fn: (r) => r["_field"] == "Total")

data
        |> first()
        |> yield(name: "first")

data
        |> last()
        |> yield(name: "last")
```
which is a valid Flux query that includes programmatically generated times and other fields.
