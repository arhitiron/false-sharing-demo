package report

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"

	"github.com/arhitiron/false-sharing-demo/dto"
)

const benchmarkUnit = "ns/op"

// Returns grouped results
func getBenchmarkResults(writerPairs, writeOperations []int) map[int]map[int][]dto.Report {
	reports := map[int]map[int][]dto.Report{}
	for _, wp := range writerPairs {
		wps := map[int][]dto.Report{}
		for _, wo := range writeOperations {
			wps[wo] = getResultsByPairsAndOperations(wp, wo)
		}
		reports[wp] = wps
	}
	return reports
}

// Returns result based on writers amount
func getResultsByPairsAndOperations(writerPairs, writeOperations int) []dto.Report {
	workersNum := len(workers)
	var results []dto.Report
	for i := 0; i < workersNum; i++ {
		for j := i + 1; j < workersNum; j++ {
			w1 := workers[i]
			w2 := workers[j]
			r := testing.Benchmark(func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					runWorkers(w1, w2, writerPairs, writeOperations)
				}
			})
			res := int(r.T) / r.N
			results = append(results, dto.Report{
				Fields:          []int{i, j},
				WriterPairs:     writerPairs,
				WriteOperations: writeOperations,
				Result:          res,
				Unit:            benchmarkUnit,
			})
		}
	}
	return results
}

func runWorkers(w1, w2 fieldWorker, writerPairs, writeOperations int) {
	d := &dto.Dummy{}
	group := errgroup.Group{}
	for i := 0; i < writerPairs; i++ {
		group.Go(w1(d, writeOperations))
		group.Go(w2(d, writeOperations))
	}
	err := group.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
