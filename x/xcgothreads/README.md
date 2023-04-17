# Go goroutines and C threads

* start N goroutines, each printing out their thread id

Strange, expected thread count to be 8 (NumCPU), but it hovers between 8 and 11.

> the thread will be terminated -- "https://pkg.go.dev/runtime#LockOSThread"

So there is some place in the runtime, where a thread is terminated and a new thread is spawned.

Notes:

* we should also see the effects of `GOMAXPROCS`

Ha!

```
$ GOMAXPROCS=1 ./main -n 10000  | jq -r .tid | sort | uniq -c
   3285 1069782
   6716 1069787
```
