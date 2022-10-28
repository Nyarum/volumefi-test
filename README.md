# volumefi-test

To run tests:
- go tests .

To run benchmarks:
- go test -bench=.

To build it use:
- go run .

To use it follow the format:
- POST http://localhost:8080/calculate with body:

[["SFO", "EWR"]]

or

[["ATL", "EWR"], ["SFO", "ATL"]]

or 

[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
