package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, cancelCtx := errgroup.WithContext(context.Background())
	
	g.Go(func() error {
		s := http.Server{
		    Addr:    ":8011",
	    }
		go func(ctx context.Context) {
			<-ctx.Done()
			fmt.Println("http routione：other work done")
			s.Shutdown(context.Background())
		}(cancelCtx)

		fmt.Println("http routione 8011：START!")
		return s.ListenAndServe()
	})

	g.Go(func() error {
		s := http.Server{
		    Addr:    ":8012",
	    }
		go func(ctx context.Context) {
			<-ctx.Done()
			fmt.Println("http routione：other work done")
			s.Shutdown(context.Background())
		}(cancelCtx)

		fmt.Println("http routione 8012：START!")
		return s.ListenAndServe()
	})

	g.Go(func() error {
		c := make(chan os.Signal)
	    signal.Notify(c)
	    fmt.Println("signal routine：START!")
	    for {
			select {
			case s := <-c:
				return fmt.Errorf("hookSignal get %v signal", s)
			case <-cancelCtx.Done():
				return fmt.Errorf("signal routine：other work done")
			}
	    }
	})

	if err := g.Wait(); err != nil {
		fmt.Println("error : ", err.Error())
	}

	fmt.Println("DONE!")
}




