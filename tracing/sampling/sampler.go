package tracing

import (
	math/rand
)

// Sampler determines if a trace should be sampled
type Sampler interface {
	ShouldSample(traceID string) bool
}

// ProbabilisticSampler samples based on probability
type ProbabilisticSampler struct {
	Rate float64
}

func (p *ProbabilisticSampler) ShouldSample(traceID string) bool {
	return rand.Float64() < p.Rate
}

// ErrorSampler always samples errors
type ErrorSampler struct{}

func (e *ErrorSampler) ShouldSample(traceID string) bool {
	return true
}

// CompositeSampler combines multiple samplers
type CompositeSampler struct {
	samplers []Sampler
}

func (c *CompositeSampler) ShouldSample(traceID string) bool {
	for _, s := range c.samplers {
		if s.ShouldSample(traceID) {
			return true
		}
	}
	return false
}
