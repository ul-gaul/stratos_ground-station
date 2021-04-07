package read

import (
    "encoding/csv"
    "github.com/jszwec/csvutil"
    "io"
    "os"
    
    "github.com/ul-gaul/stratos_ground-station/acquisition/packet"
)

type CsvReader struct {
    *csvutil.Decoder
}

func NewCsvReader(reader io.Reader) (CsvReader, error) {
    decoder, err := csvutil.NewDecoder(csv.NewReader(reader))
    return CsvReader{decoder}, err
}

// Read reads the next packet.
func (cr CsvReader) Read() (packet.TransformedData, error) {
    var pkt packet.TransformedData
    err := cr.Decode(&pkt)
    return pkt, err
}

// ReadAll reads the next packet.
func (cr CsvReader) ReadAll() ([]packet.TransformedData, error) {
    var pkts []packet.TransformedData
    err := cr.Decode(&pkts)
    return pkts, err
}

func Csv(path string) ([]packet.TransformedData, error) {
    f, err := os.Open(path)
    if err != nil { return nil, err }
    
    reader, err := NewCsvReader(f)
    if err != nil { return nil, err }
    
    return reader.ReadAll()
}

