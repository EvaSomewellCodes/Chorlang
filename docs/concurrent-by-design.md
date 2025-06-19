# Concurrent by Design

Concurrency is woven into ChoreLang using goroutines and channels, mirroring
Go's model. Any function can be launched with `start`, and channels coordinate
data between routines. This allows programs to scale without complex threading.
