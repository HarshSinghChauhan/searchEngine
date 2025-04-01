package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// Corrected struct with valid Parquet tags
type ParquetRecord struct {
	Message        string `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	MessageRaw     string `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	StructuredData string `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Tag            string `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Sender         string `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Groupings      string `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Event          string `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	EventId        int64  `parquet:"name=EventId, type=INT64"`
	NanoTimeStamp  int64  `parquet:"name=NanoTimeStamp, type=INT64"`
	Namespace      string `parquet:"name=Namespace, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}

var (
	dataLock sync.RWMutex
	data     []ParquetRecord
)

// LoadParquetFiles reads and parses Parquet files into memory
func LoadParquetFiles(filePath string) error {
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		return fmt.Errorf("failed to open Parquet file: %v", err)
	}
	defer fr.Close()

	// Pass a pointer to the ParquetRecord struct (new(ParquetRecord)) instead of a value type
	pr, err := reader.NewParquetReader(fr, new(ParquetRecord), 1) // Pass pointer to struct here
	if err != nil {
		return fmt.Errorf("failed to create Parquet reader: %v", err)
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	log.Printf("Number of rows in Parquet file: %d", num)

	if num == 0 {
		return fmt.Errorf("no records found in Parquet file")
	}

	var temp []ParquetRecord
	for i := 0; i < num; i++ {
		var record ParquetRecord
		if err := pr.Read(&record); err != nil {
			return fmt.Errorf("failed to read Parquet data: %v", err)
		}
		temp = append(temp, record) // Append each record dynamically
	}

	log.Printf("Loaded %d records from %s", len(temp), filePath)
	dataLock.Lock()
	data = append(data, temp...)
	dataLock.Unlock()

	return nil
}
