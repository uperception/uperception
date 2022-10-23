package collectors

type Collector interface {
	Collect(url string) error
}
