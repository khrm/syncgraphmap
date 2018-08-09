# syncgraphmap
Benchmark of Different Implementations of Concurrent Graph using Map


## Benchmark

```shell
[khurram]# go test -bench=. -benchtime=10s -timeout 30m && go test -bench=. -benchtime=10s -timeout 30m -cpu 8 && go test -bench=. -benchtime=10s -timeout 30m -cpu 4 && go test -bench=. -benchtime=10s -timeout 30m -cpu 2
goos: linux
goarch: amd64
BenchmarkGraphSingleMutex-16      	  1000000	     23413 ns/op
BenchmarkGraphDoubleMutex-16      	 3000000	      9023 ns/op
BenchmarkGraphSTD-16              	 3000000	      7542 ns/op
BenchmarkGraphHybrid-16           	 2000000	      7105 ns/op
BenchmarkGraphSingleMutex90-16    	 1000000	     33578 ns/op
BenchmarkGraphDoubleMutex90-16    	 2000000	     12960 ns/op
BenchmarkGraphSTD90-16            	 1000000	     14734 ns/op
BenchmarkGraphHybrid90-16         	 1000000	     12253 ns/op
BenchmarkGraphSingleMutex50-16    	 1000000	     26550 ns/op
BenchmarkGraphDoubleMutex50-16    	 2000000	      7660 ns/op
BenchmarkGraphSTD50-16            	 1000000	     10797 ns/op
BenchmarkHybrid50-16              	 1000000	     10870 ns/op
PASS
ok  	github.com/khrm/syncgraphmap	393.166s
goos: linux
goarch: amd64
BenchmarkGraphSingleMutex-8     	 1000000	     19016 ns/op
BenchmarkGraphDoubleMutex-8     	 3000000	      6235 ns/op
BenchmarkGraphSTD-8             	 5000000	     17911 ns/op
BenchmarkGraphHybrid-8          	 1000000	     16694 ns/op
BenchmarkGraphSingleMutex90-8   	 1000000	     40189 ns/op
BenchmarkGraphDoubleMutex90-8   	 2000000	      7267 ns/op
BenchmarkGraphSTD90-8           	 1000000	     10616 ns/op
BenchmarkGraphHybrid90-8        	 2000000	     15396 ns/op
BenchmarkGraphSingleMutex50-8   	 1000000	     24331 ns/op
BenchmarkGraphDoubleMutex50-8   	 3000000	      6318 ns/op
BenchmarkGraphSTD50-8           	 2000000	     10192 ns/op
BenchmarkHybrid50-8             	 2000000	     11435 ns/op
PASS
ok  	github.com/khrm/syncgraphmap	577.223s
goos: linux
goarch: amd64
BenchmarkGraphSingleMutex-4     	 1000000	     18188 ns/op
BenchmarkGraphDoubleMutex-4     	 3000000	      4858 ns/op
BenchmarkGraphSTD-4             	 3000000	      5726 ns/op
BenchmarkGraphHybrid-4          	 2000000	      7470 ns/op
BenchmarkGraphSingleMutex90-4   	 1000000	     37638 ns/op
BenchmarkGraphDoubleMutex90-4   	 2000000	     13232 ns/op
BenchmarkGraphSTD90-4           	 1000000	     10489 ns/op
BenchmarkGraphHybrid90-4        	 1000000	     12797 ns/op
BenchmarkGraphSingleMutex50-4   	 1000000	     19618 ns/op
BenchmarkGraphDoubleMutex50-4   	 3000000	      9370 ns/op
BenchmarkGraphSTD50-4           	 2000000	     11293 ns/op
BenchmarkHybrid50-4             	 2000000	     12088 ns/op
PASS
ok  	github.com/khrm/syncgraphmap	458.935s
goos: linux
goarch: amd64
BenchmarkGraphSingleMutex-2     	 1000000	     19500 ns/op
BenchmarkGraphDoubleMutex-2     	 2000000	      8406 ns/op
BenchmarkGraphSTD-2             	 2000000	     10126 ns/op
BenchmarkGraphHybrid-2          	 2000000	     10476 ns/op
BenchmarkGraphSingleMutex90-2   	 1000000	     52655 ns/op
BenchmarkGraphDoubleMutex90-2   	 1000000	     18469 ns/op
BenchmarkGraphSTD90-2           	 1000000	     22334 ns/op
BenchmarkGraphHybrid90-2        	 2000000	     16947 ns/op
BenchmarkGraphSingleMutex50-2   	 1000000	     14013 ns/op
BenchmarkGraphDoubleMutex50-2   	 2000000	     11412 ns/op
BenchmarkGraphSTD50-2           	 2000000	     13874 ns/op
BenchmarkHybrid50-2             	 2000000	     13719 ns/op
PASS
ok  	github.com/khrm/syncgraphmap	473.606s

```
