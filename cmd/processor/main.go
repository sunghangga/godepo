package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"godepo/internal/services/collector"
	"godepo/internal/services/detector"
	"godepo/internal/services/flagger"
)

var (
	brokers      = []string{"localhost:9092"}
	runCollector = flag.Bool("collector", false, "run collector processor")
	runFlagger   = flag.Bool("flagger", false, "run flagger processor")
	runDetector  = flag.Bool("detector", false, "run detector processor")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	if *runCollector {
		log.Println("starting collector")
		grp.Go(collector.Run(ctx, brokers))
	}
	if *runFlagger {
		log.Println("starting flagger")
		grp.Go(flagger.Run(ctx, brokers))
	}
	if *runDetector {
		log.Println("starting detector")
		grp.Go(detector.Run(ctx, brokers))
	}

	// Wait for SIGINT/SIGTERM
	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}
	cancel()
	if err := grp.Wait(); err != nil {
		log.Println(err)
	}
	log.Println("done")
}
