package main

import (
	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	"github.com/nazwa/ripple-price-alerts/watcher"
	"time"
)

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}

	Jobs []*watcher.WatcherStruct
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) init() {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		logger.Error(err)
		return
	}

	job, err := watcher.LoadFromFile(dir + "/jobs.json")
	if err != nil {
		logger.Error(err)
		return
	}
	p.Jobs = append(p.Jobs, job)

}

func (p *program) run() error {
	// Dirty hack to force goroutine swap
	time.Sleep(200 * time.Millisecond)

	p.init()

	for {

		// Loop through all jobs
		for _, job := range p.Jobs {
			job.Update()
			// And all pairs within each job
			for _, item := range job.Pairs {
				if item.Error != nil {
					logger.Error(item.Error)
				} else {
					if item.NewAlertDetected() {
						item.MarkReported()
						if err := job.Notify(item); err != nil {
							logger.Error(err)
						}
					}
				}
			}
		}

		time.Sleep(10 * time.Second)
	}

	/*
		for {
			select {
			case <-p.exit:
				ticker.Stop()
				return nil
			default:

			}
		}*/
}

func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}
