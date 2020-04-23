package main

import "github.com/arhitiron/false-sharing-demo/report"

func main() {
	rw := report.NewReportsWriter()
	if err := rw.Run(); err != nil {
		panic(err)
	}
}
