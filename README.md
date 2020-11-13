# Open Source Issue Tracking System - opists
Manage your project issues.

## Build
[see here](apidoc/Build.md)

## Testing API with cucumber godog

```
cd test/api
go test
```
Output:
```
2020/11/13 13:43:29 GET  GetAllProject 0s
...2020/11/13 13:43:29 GET  GetProject 0s
... 6


2 scenarios (2 passed)
6 steps (6 passed)
32.9977ms
testing: warning: no tests to run
PASS
ok  	github.com/DevBoxFanBoy/opists/tests	0.322s [no tests to run]

Process finished with exit code 0
```