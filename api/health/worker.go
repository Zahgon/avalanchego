// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package health

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	allTags = []string{AllTag}

	errRestrictedTag  = errors.New("restricted tag")
	errDuplicateCheck = errors.New("duplicate check")
)

type worker struct {
	log           logging.Logger
	name          string
	failingChecks *prometheus.GaugeVec
	checksLock    sync.RWMutex
	checks        map[string]*taggedChecker

	resultsLock                 sync.RWMutex
	results                     map[string]Result
	numFailingApplicationChecks int
	tags                        map[string]set.Set[string] // tag -> set of check names

	startOnce sync.Once
	closeOnce sync.Once
	wg        sync.WaitGroup
	closer    chan struct{}
}

type taggedChecker struct {
	checker            Checker
	isApplicationCheck bool
	tags               []string
}

func newWorker(
	log logging.Logger,
	name string,
	failingChecks *prometheus.GaugeVec,
) *worker {
	_ = "STUB: not implemented"
	// Initialize the number of failing checks to 0 for all checks
	return nil
}

func (w *worker) RegisterCheck(name string, check Checker, tags ...string) error {
	_ = "STUB: not implemented"
	// We ensure [AllTag] isn't contained in [tags] to prevent metrics from
	// double counting.
	return nil
}

// Add the check to each tag

// Add the special AllTag descriptor

// Whenever a new check is added - it is failing

// If this is a new application-wide check, then all of the registered tags
// now have one additional failing check.
/*=healthy*/ /*=register*/

func (w *worker) RegisterMonotonicCheck(name string, checker Checker, tags ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) Results(tags ...string) (map[string]Result, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// if no tags are specified, return all checks

// we always want to include the application tag

func (w *worker) Start(ctx context.Context, freq time.Duration) { _ = "STUB: not implemented"; return }

func (w *worker) Stop() { _ = "STUB: not implemented"; return }

func (w *worker) runChecks(ctx context.Context) {
	_ = "STUB: not implemented"

	// Copy the [w.checks] map to collect the checks that we will be running
	// during this iteration. If [w.checks] is modified during this iteration of
	// [runChecks], then the added check will not be run until the next
	// iteration.
	return
}

func (w *worker) runCheck(ctx context.Context, wg *sync.WaitGroup, name string, check *taggedChecker) {
	_ = "STUB: not implemented"
	return
}

// To avoid any deadlocks when [RegisterCheck] is called with a lock
// that is grabbed by [check.HealthCheck], we ensure that no locks
// are held when [check.HealthCheck] is called.

/*=healthy*/ /*=register*/

/*=healthy*/ /*=register*/

// updateMetrics updates the metrics for the given check. If [healthy] is true,
// then the check is considered healthy and the metrics are decremented.
// Otherwise, the check is considered unhealthy and the metrics are incremented.
// [register] must be true only if this is the first time the check is being
// registered.
func (w *worker) updateMetrics(tc *taggedChecker, healthy bool, register bool) {
	_ = "STUB: not implemented"
	return
}

// Note: [w.tags] will include AllTag.

// If this is the first time this tag was registered, we also need to
// account for the currently failing application-wide checks.
