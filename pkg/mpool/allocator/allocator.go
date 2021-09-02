package allocator

type allocator interface {
	alloc(sz int64)
	free(sz *interface{})
}
