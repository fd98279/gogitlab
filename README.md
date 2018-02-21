# gogitlab

To build:
************  Building Go project: gogitlab  ************
  with GOPATH: C:\workspace\Test\gogitlab
>> Running: C:\Go\bin\go.exe install -v -gcflags "-N -l" ./...
   ^^^ Terminated, exit code: 0 ^^^
************  Build terminated.  ************


To run tests:
************  Building Go project: gogitlab  ************
  with GOPATH: C:\workspace\Test\gogitlab
>> Running: C:\Go\bin\go.exe test -v -gcflags "-N -l" ./...
=== RUN   TestAverage
--- PASS: TestAverage (0.00s)
=== RUN   TestCallURL
Elapsed Time:  0.09123132181818183  seconds
--- PASS: TestCallURL (5.02s)
PASS
ok  	lib/util	5.126s
?   	main	[no test files]
   ^^^ Terminated, exit code: 0 ^^^
************  Build terminated.  ************

To run the code:
C:\workspace\Test\gogitlab\bin>main
Starting http get...
Elapsed Time:  0.09470011886792454  seconds
Finished http get...
