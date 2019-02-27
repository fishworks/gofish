package receipt

import (
	"encoding/json"
	"io"
	"time"
)

const ReceiptFilename = "INSTALL_RECEIPT.json"

type InstallReceipt struct {
	// The name of the food installed
	Name string `json:"name"`
	// Which rig did this food come from?
	Rig string `json:"rig"`
	// Time that this food was last modified (upgraded)
	LastModified time.Time `json:"last-modified"`
	// What version of GoFish was this last modified with?
	GoFishVersion string `json:"gofish-version"`
}

// NewFromReader reads in an install receipt from an io.Reader. Useful when reading from a file stream.
func NewFromReader(r io.Reader) (*InstallReceipt, error) {
	var receipt InstallReceipt
	err := json.NewDecoder(r).Decode(&receipt)
	return &receipt, err
}

// Save writes the install receipt into the given io.Writer.
func (i *InstallReceipt) Save(w io.Writer) error {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	w.Write([]byte("\n"))
	return err
}
