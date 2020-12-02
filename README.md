Advent of Code - using Go
===

My attempt at solving Advent of Code Challenges using Go.

Run
---
**Pre Requisite:** If you want to be able to fetch your input on demand, you
need to have your session cookie from the AoC website in a file:
`$REPO/cookie`. It is, however, possible to just put your input in:
`$REPO/$year/data/$day` and the script will not bother downloading new input.

The easiest way to run a solution is by using: `go run $year/$day/main.go`, but
it is also possible to use: `go build $year/$day/main.go && ./main` to run the
binary.

### Tests
If you want to run the tests, you use: `go test $year/$day/*.go`.

2019
---
I used some of the 2019 challenges to get familiar with the how Go wants you to
structure your projects.
Thereafter I wrote the `./util/` package which I thought could come in handy
for the 2020 challenges.

2020
---
