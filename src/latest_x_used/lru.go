package latest_x_used

type LRU struct {
}

type LRUInterface interface {
	Get()
	Put()
	Purge()
}
