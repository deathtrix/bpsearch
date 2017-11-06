package interfaces

// CrawlerInterface interface
type CrawlerInterface interface {
	Start(string, chan<- string)
}

// StoreInterface interface
type StoreInterface interface {
	Put(interface{}, interface{})
	Get(interface{}) (interface{}, bool)
	Remove(interface{})
}

// IndexerInterface interface
type IndexerInterface interface {
	Start(string, StoreInterface)
}
