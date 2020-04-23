package report

import (
	"sync/atomic"

	"github.com/arhitiron/false-sharing-demo/dto"
)

type (
	worker      func() error
	fieldWorker func(*dto.Dummy, int) worker
)

var (
	workers = []fieldWorker{
		field0Worker,
		field1Worker,
		field2Worker,
		field3Worker,
		field4Worker,
		field5Worker,
		field6Worker,
		field7Worker,
		field8Worker,
		field9Worker,
		field10Worker,
		field11Worker,
		field12Worker,
		field13Worker,
		field14Worker,
		field15Worker,
	}
)

func field0Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field0, operations)
}

func field1Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field1, operations)
}

func field2Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field2, operations)
}

func field3Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field3, operations)
}

func field4Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field4, operations)
}

func field5Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field5, operations)
}

func field6Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field6, operations)
}

func field7Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field7, operations)
}

func field8Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field8, operations)
}

func field9Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field9, operations)
}

func field10Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field10, operations)
}

func field11Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field11, operations)
}

func field12Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field12, operations)
}

func field13Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field13, operations)
}

func field14Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field14, operations)
}

func field15Worker(dummy *dto.Dummy, operations int) worker {
	return getWorker(&dummy.Field15, operations)
}

func getWorker(field *int64, operations int) worker {
	return func() error {
		for i := 0; i < operations; i++ {
			atomic.AddInt64(field, 1)
		}
		return nil
	}
}
