package report

import (
	"encoding/json"
	"io/ioutil"
)

const reportsFileName = "result.json"

var (
	writerPairs     = []int{1, 2, 10, 100}
	writeOperations = []int{10, 100, 1000, 10000}
)

type (
	ReportsWriter interface {
		Run() error
	}

	reportsWriter struct {
	}
)

func NewReportsWriter() ReportsWriter {
	return &reportsWriter{}
}

func (r *reportsWriter) Run() error {
	results := getBenchmarkResults(writerPairs, writeOperations)
	file, _ := json.MarshalIndent(results, "", " ")

	return ioutil.WriteFile(reportsFileName, file, 0644)
}
