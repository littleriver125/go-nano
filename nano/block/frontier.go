package block

import (
	"bytes"

	"github.com/alexbakker/gonano/nano"
	"github.com/alexbakker/gonano/nano/internal/util"
)

const (
	FrontierSize = nano.AddressSize + HashSize
)

type Frontier struct {
	Address nano.Address
	Hash    Hash
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (f *Frontier) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	var err error
	if _, err = buf.Write(f.Address[:]); err != nil {
		return nil, err
	}

	if _, err = buf.Write(f.Hash[:]); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (f *Frontier) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)

	var err error
	if _, err = reader.Read(f.Address[:]); err != nil {
		return err
	}

	if _, err = reader.Read(f.Hash[:]); err != nil {
		return err
	}

	return util.AssertReaderEOF(reader)
}

func (f *Frontier) IsZero() bool {
	return f.Hash.IsZero()
}
