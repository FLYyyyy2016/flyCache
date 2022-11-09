# flyCache

灵感来源自：
https://github.com/coocood/freecache
https://github.com/allegro/bigcache
https://github.com/patrickmn/go-cache

### 测试结果
```
goos: linux
goarch: amd64
pkg: github.com/FLYyyyy2016/flyCache/bench
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkCaches
BenchmarkCaches/BigCacheZipfRead
BenchmarkCaches/BigCacheZipfRead-12         	75957837	        14.92 ns/op	      11 B/op	       2 allocs/op
BenchmarkCaches/FastCacheZipfRead
BenchmarkCaches/FastCacheZipfRead-12        	84189523	        13.34 ns/op	       8 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheZipfRead
BenchmarkCaches/FreeCacheZipfRead-12        	32567366	        31.25 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheZipfRead
BenchmarkCaches/GroupCacheZipfRead-12       	55118144	        22.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkCaches/RistrettoZipfRead
BenchmarkCaches/RistrettoZipfRead-12        	68341687	        19.64 ns/op	      25 B/op	       1 allocs/op
BenchmarkCaches/SyncMapZipfRead
BenchmarkCaches/SyncMapZipfRead-12          	196505307	         5.957 ns/op	       0 B/op	       0 allocs/op
BenchmarkCaches/FlyMapZipfRead
BenchmarkCaches/FlyMapZipfRead-12           	120846295	         9.909 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/BigCacheOneKeyRead
BenchmarkCaches/BigCacheOneKeyRead-12       	31816006	        32.14 ns/op	      12 B/op	       3 allocs/op
BenchmarkCaches/FastCacheOneKeyRead
BenchmarkCaches/FastCacheOneKeyRead-12      	27399451	        43.91 ns/op	       8 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheOneKeyRead
BenchmarkCaches/FreeCacheOneKeyRead-12      	 6988585	       169.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheOneKeyRead
BenchmarkCaches/GroupCacheOneKeyRead-12     	10372804	       118.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCaches/RistrettoOneKeyRead
BenchmarkCaches/RistrettoOneKeyRead-12      	25370245	        48.77 ns/op	      24 B/op	       1 allocs/op
BenchmarkCaches/SyncMapOneKeyRead
BenchmarkCaches/SyncMapOneKeyRead-12        	274208064	         4.357 ns/op	       0 B/op	       0 allocs/op
BenchmarkCaches/FlyMapOneKeyRead
BenchmarkCaches/FlyMapOneKeyRead-12         	24544654	        48.23 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/BigCacheOneKeyWrite
BenchmarkCaches/BigCacheOneKeyWrite-12      	 4371194	       314.5 ns/op	      75 B/op	       2 allocs/op
BenchmarkCaches/FastCacheOneKeyWrite
BenchmarkCaches/FastCacheOneKeyWrite-12     	 8419677	       144.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheOneKeyWrite
BenchmarkCaches/FreeCacheOneKeyWrite-12     	 5946896	       188.6 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheOneKeyWrite
BenchmarkCaches/GroupCacheOneKeyWrite-12    	 5794826	       201.9 ns/op	      48 B/op	       4 allocs/op
BenchmarkCaches/RistrettoOneKeyWrite
BenchmarkCaches/RistrettoOneKeyWrite-12     	 4211682	       281.2 ns/op	     100 B/op	       4 allocs/op
BenchmarkCaches/SyncMapOneKeyWrite
BenchmarkCaches/SyncMapOneKeyWrite-12       	 4693840	       253.9 ns/op	      64 B/op	       5 allocs/op
BenchmarkCaches/FlyMapOneKeyWrite
BenchmarkCaches/FlyMapOneKeyWrite-12        	 6405898	       185.9 ns/op	      48 B/op	       4 allocs/op
BenchmarkCaches/BigCacheZipfWrite
BenchmarkCaches/BigCacheZipfWrite-12        	24688425	        57.64 ns/op	      59 B/op	       1 allocs/op
BenchmarkCaches/FastCacheZipfWrite
BenchmarkCaches/FastCacheZipfWrite-12       	46932324	        24.50 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheZipfWrite
BenchmarkCaches/FreeCacheZipfWrite-12       	38385217	        31.55 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheZipfWrite
BenchmarkCaches/GroupCacheZipfWrite-12      	27955227	        44.49 ns/op	      47 B/op	       3 allocs/op
BenchmarkCaches/RistrettoZipfWrite
BenchmarkCaches/RistrettoZipfWrite-12       	17457244	        75.66 ns/op	     100 B/op	       4 allocs/op
BenchmarkCaches/SyncMapZipfWrite
BenchmarkCaches/SyncMapZipfWrite-12         	 4720960	       255.8 ns/op	      63 B/op	       4 allocs/op
BenchmarkCaches/FlyMapZipfWrite
BenchmarkCaches/FlyMapZipfWrite-12          	28813032	        41.09 ns/op	      48 B/op	       3 allocs/op
BenchmarkCaches/BigCacheOneKeyWrite#01
BenchmarkCaches/BigCacheOneKeyWrite#01-12   	 4241176	       273.3 ns/op	      67 B/op	       2 allocs/op
BenchmarkCaches/FastCacheOneKeyWrite#01
BenchmarkCaches/FastCacheOneKeyWrite#01-12  	 8162754	       141.8 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheOneKeyWrite#01
BenchmarkCaches/FreeCacheOneKeyWrite#01-12  	 6663332	       184.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheOneKeyWrite#01
BenchmarkCaches/GroupCacheOneKeyWrite#01-12 	 5574392	       210.5 ns/op	      48 B/op	       4 allocs/op
BenchmarkCaches/RistrettoOneKeyWrite#01
BenchmarkCaches/RistrettoOneKeyWrite#01-12  	 4457104	       273.3 ns/op	     100 B/op	       4 allocs/op
BenchmarkCaches/SyncMapOneKeyWrite#01
BenchmarkCaches/SyncMapOneKeyWrite#01-12    	 4765220	       248.4 ns/op	      64 B/op	       5 allocs/op
BenchmarkCaches/FlyMapOneKeyWrite#01
BenchmarkCaches/FlyMapOneKeyWrite#01-12     	30469768	        41.24 ns/op	      48 B/op	       3 allocs/op
BenchmarkCaches/BigCacheZipfMixed
BenchmarkCaches/BigCacheZipfMixed-12        	16646809	        92.37 ns/op	      14 B/op	       2 allocs/op
BenchmarkCaches/FastCacheZipfMixed
BenchmarkCaches/FastCacheZipfMixed-12       	15871736	        70.47 ns/op	       7 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheZipfMixed
BenchmarkCaches/FreeCacheZipfMixed-12       	40789767	        32.56 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheZipfMixed
BenchmarkCaches/GroupCacheZipfMixed-12      	50358004	        25.18 ns/op	       7 B/op	       0 allocs/op
BenchmarkCaches/RistrettoZipfMixed
BenchmarkCaches/RistrettoZipfMixed-12       	16866546	        72.04 ns/op	      34 B/op	       1 allocs/op
BenchmarkCaches/SyncMapZipfMixed
BenchmarkCaches/SyncMapZipfMixed-12         	144106375	         8.032 ns/op	       5 B/op	       0 allocs/op
BenchmarkCaches/FlyMapZipfMixed
BenchmarkCaches/FlyMapZipfMixed-12          	20244728	        85.67 ns/op	       7 B/op	       1 allocs/op
BenchmarkCaches/BigCacheOneKeyMixed
BenchmarkCaches/BigCacheOneKeyMixed-12      	 5074706	       234.7 ns/op	      13 B/op	       2 allocs/op
BenchmarkCaches/FastCacheOneKeyMixed
BenchmarkCaches/FastCacheOneKeyMixed-12     	 6097554	       194.0 ns/op	       7 B/op	       1 allocs/op
BenchmarkCaches/FreeCacheOneKeyMixed
BenchmarkCaches/FreeCacheOneKeyMixed-12     	 7528140	       161.7 ns/op	       4 B/op	       1 allocs/op
BenchmarkCaches/GroupCacheOneKeyMixed
BenchmarkCaches/GroupCacheOneKeyMixed-12    	 8487646	       145.6 ns/op	      17 B/op	       1 allocs/op
BenchmarkCaches/RistrettoOneKeyMixed
BenchmarkCaches/RistrettoOneKeyMixed-12     	 4797948	       241.9 ns/op	      35 B/op	       1 allocs/op
BenchmarkCaches/SyncMapOneKeyMixed
BenchmarkCaches/SyncMapOneKeyMixed-12       	100000000	        10.76 ns/op	       6 B/op	       0 allocs/op
BenchmarkCaches/FlyMapOneKeyMixed
BenchmarkCaches/FlyMapOneKeyMixed-12        	20071470	        67.51 ns/op	       7 B/op	       1 allocs/op
```
