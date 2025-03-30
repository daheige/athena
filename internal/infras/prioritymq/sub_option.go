package prioritymq

// SubOption subscribe option
type SubOption func(s *SubscribeOptions)

// SubscribeOptions subscribe message option
type SubscribeOptions struct {
	// specifies the consumer name
	// for kafka name eq group_id
	Name string

	// priority consumer buffer size
	BufferSize int

	// Commit the offset to the backend for kafka
	// Note: calling Commit performs a blocking synchronous operation.
	CommitOffsetBlock bool
}

// WithSubName set sub name
func WithSubName(name string) SubOption {
	return func(s *SubscribeOptions) {
		s.Name = name
	}
}

// WithBufferSize set priority mq buffer size
func WithBufferSize(size int) SubOption {
	return func(s *SubscribeOptions) {
		s.BufferSize = size
	}
}

// WithCommitOffsetBlock set consumerGroup commitOffsetBlock
func WithCommitOffsetBlock(commitOffsetBlock bool) SubOption {
	return func(s *SubscribeOptions) {
		s.CommitOffsetBlock = commitOffsetBlock
	}
}
