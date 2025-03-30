package prioritymq

// PubOption option for producer
type PubOption func(p *PublishOptions)

// PublishOptions publish options
type PublishOptions struct {
	Key string
}

// WithPubName set publish name
func WithPubName(key string) PubOption {
	return func(p *PublishOptions) {
		p.Key = key
	}
}
