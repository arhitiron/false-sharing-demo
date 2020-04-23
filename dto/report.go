package dto

type Report struct {
	Fields          []int  `json:"fields"`
	WriterPairs     int    `json:"writerPairs"`
	WriteOperations int    `json:"writeOperations"`
	Result          int    `json:"result"`
	Unit            string `json:"unit"`
}
