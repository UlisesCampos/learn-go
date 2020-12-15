## Diference between concurrency and parallelism
Parallelism is the process by which a CPU executes multiple tasks simultaneosly.
Concurrency is the CPU's ability to switch between multiple tasks, which start, run, and complete
while overlapping one another.

A channel in Go is the conduit through which concurrency flows. Channels can be undirectional
with data either sent to or received by them - or bidirectional, which can do both.