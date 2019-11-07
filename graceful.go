package main

import (
        "context"
        "fmt"
        "net/http"
        "os"
        "os/signal"
        "time"

        "github.com/go-chi/chi"
        "github.com/go-chi/chi/middleware"
        "github.com/go-chi/valve"
)

func main() {

        // Our graceful valve shut-off package to manage code preemption and
        // shutdown signaling.
        valv := valve.New()
        baseCtx := valv.Context()

        // Example of a long running background worker thing..
        go func(ctx context.Context) {
                var flg int
                for {
                        if flg == 1 {
                                fmt.Println("background goroutine exit")
                                break
                        }
                        //
                        <-time.After(1 * time.Second)

                        func() {

                                valve.Lever(ctx).Open()
                                defer valve.Lever(ctx).Close()

                                // actual code doing stuff..
                                fmt.Println("tick..")
                                time.Sleep(2 * time.Second)
                                // end-logic

                                // signal control..
                                select {
                                case <-valve.Lever(ctx).Stop():
                                        fmt.Println("valve is closed")
                                        flg = 1
                                        return

                                case <-ctx.Done():
                                        fmt.Println("context is cancelled, go home.")
                                        flg = 1
                                        return
                                default:
                                }

                        }()

                }
        }(baseCtx)

        // HTTP service running in this program as well. The valve context is set
        // as a base context on the server listener at the point where we instantiate
        // the server - look lower.
        r := chi.NewRouter()
        r.Use(middleware.RequestID)
        r.Use(middleware.Logger)

        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte("sup"))
        })

        r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {

                valve.Lever(r.Context()).Open()
                defer valve.Lever(r.Context()).Close()

                select {
                case <-valve.Lever(r.Context()).Stop():
                        fmt.Println("valve is closed. finish up..")

                case <-time.After(2 * time.Second):
                        // The above channel simulates some hard work.
                        // We want this handler to complete successfully during a shutdown signal,
                        // so consider the work here as some background routine to fetch a long running
                        // search query to find as many results as possible, but, instead we cut it short
                        // and respond with what we have so far. How a shutdown is handled is entirely
                        // up to the developer, as some code blocks are preemptable, and others are not.
                        for i := 0; i < 34; i++ {
                                fmt.Println("busy", i)
                                time.Sleep(1 * time.Second)
                        }
                }

                w.Write([]byte(fmt.Sprintf("all done.\n")))
        })

        srv := http.Server{Addr: ":3333", Handler: chi.ServerBaseContext(baseCtx, r)}

        go func() {
                srv.ListenAndServe()
                fmt.Println("server exit")
        }()
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)

        <-c
        func() {
                // sig is a ^C, handle it
                fmt.Println("shutting down..")
                fmt.Println("valv blocked")
                fmt.Println()
                t0:=time.Now()
                // first valv
                valv.Shutdown(20 * time.Second) //block
                d0:=time.Since(t0)
                fmt.Println("d0",d0)
                fmt.Println()

                fmt.Println("ctx0")
                // create context with timeout
                ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
                fmt.Println("ctx1")
                defer cancel()
                defer fmt.Println("defer print")
                fmt.Println("ctx2")
                // start http shutdown
                fmt.Println("shutdown blocked 20 second")
           t1:=time.Now()
                err := srv.Shutdown(ctx)
             d1:=time.Since(t1)
                fmt.Println("shutdown err", err,"d1",d1)

                t := time.Now()
                fmt.Println("select blocked")
                // verify, in worst case call cancel via defer
                select {
                case <-time.After(21 * time.Second):
                        fmt.Println("not all connections done")
                case <-ctx.Done():
                        d := time.Since(t)
                        fmt.Println("done", d)
                }

        }()
        fmt.Println("another 10 seconds")

        time.Sleep(10 * time.Second)
}
