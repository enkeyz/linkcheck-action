# linkcheck-action

Github action mainly for [Project-based learning](https://github.com/practical-tutorials/project-based-learning) repository. Scans README.md for broken links. You need to trigger this workflow manually after installing on the repository.

## Benchmark
Depends of how fast web servers answering :) Seriously, because of Golang's awesome goroutines, it's really fast.

## TODO
- allow user to set the max number of concurrent health check through workspace config
    - weighted semaphore
- log with colors - logrus
- check why some websites gives 403, even though they're working normally - prob request header issue
- write better regex to parse links