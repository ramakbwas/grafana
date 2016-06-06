package tsdb

type Executor interface {
	Execute(queries QuerySlice, context *QueryContext) *BatchResult
}

var registry map[string]GetExecutorFn

type GetExecutorFn func(dsInfo *DataSourceInfo) Executor

func init() {
	registry = make(map[string]GetExecutorFn)
}

func getExecutorFor(dsInfo *DataSourceInfo) Executor {
	if fn, exists := registry[dsInfo.Type]; exists {
		return fn(dsInfo)
	}
	return nil
}

func RegisterExecutor(dsType string, fn GetExecutorFn) {
	registry[dsType] = fn
}
