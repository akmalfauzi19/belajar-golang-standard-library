package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"hafizh", "akmal", "fauzi"})
	_ = writer.Write([]string{"muhammad", "fadli"})
	_ = writer.Write([]string{"john", "doe"})

	writer.Flush()
}
