package main

import (
	"encoding/csv"
	"os"
)

func main() {
	write := csv.NewWriter(os.Stdout)
	_ = write.Write([]string{"Thoby, Biba, Junior"})
	_ = write.Write([]string{"Thoby, Biba, Junior"})
	_ = write.Write([]string{"Thoby, Biba, Junior"})
	_ = write.Write([]string{"Thoby, Biba, Junior"})
	write.Flush()
}

// Pages 74
