# linkcheck-action

Github action mainly for [Project-based learning](https://github.com/practical-tutorials/project-based-learning) repository. Scans README.md for broken links. You need to trigger this workflow manually after installing on the repository.

## Benchmark
Depends of how fast web servers answering :) Seriously, because of Golang's awesome goroutines, it's really fast.

## TODO
- log with colors - logrus
- write better regex to parse links
- able to scan multiple markdown files, not just one formatted specifically