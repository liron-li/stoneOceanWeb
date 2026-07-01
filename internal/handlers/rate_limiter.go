package handlers

import (
	"sync"
	"time"
)

type ipRateLimiter struct {
	mu          sync.Mutex
	maxFailures int
	window      time.Duration
	entries     map[string]ipRateLimitEntry
}

type ipRateLimitEntry struct {
	firstFailure time.Time
	failures     int
}

func newIPRateLimiter(maxFailures int, window time.Duration) *ipRateLimiter {
	return &ipRateLimiter{
		maxFailures: maxFailures,
		window:      window,
		entries:     make(map[string]ipRateLimitEntry),
	}
}

func (l *ipRateLimiter) Allow(ip string, now time.Time) bool {
	if l == nil {
		return true
	}
	ip = normalizeRateLimitIP(ip)
	if now.IsZero() {
		now = time.Now()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	entry, ok := l.entries[ip]
	if !ok {
		return true
	}
	if !entry.firstFailure.IsZero() && now.Sub(entry.firstFailure) >= l.window {
		delete(l.entries, ip)
		return true
	}
	return entry.failures < l.maxFailures
}

func (l *ipRateLimiter) RecordFailure(ip string, now time.Time) {
	if l == nil {
		return
	}
	ip = normalizeRateLimitIP(ip)
	if now.IsZero() {
		now = time.Now()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	entry, ok := l.entries[ip]
	if !ok || (!entry.firstFailure.IsZero() && now.Sub(entry.firstFailure) >= l.window) {
		l.entries[ip] = ipRateLimitEntry{firstFailure: now, failures: 1}
		return
	}
	entry.failures++
	l.entries[ip] = entry
}

func (l *ipRateLimiter) Reset(ip string) {
	if l == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.entries, normalizeRateLimitIP(ip))
}

func normalizeRateLimitIP(ip string) string {
	if ip == "" {
		return "unknown"
	}
	return ip
}
