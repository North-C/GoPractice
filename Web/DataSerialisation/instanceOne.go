package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	mdata, err := asn1.Marshal(13)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to marshal: %s\n", err.Error())
		os.Exit(1)
	}

	var n int
	_, err = asn1.Unmarshal(mdata, &n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to Unmarshal: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("After Unmarshal: %d", n)
}
