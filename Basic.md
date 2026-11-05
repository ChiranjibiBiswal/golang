# Why Choose Go (Golang)

Go is a statically typed, compiled language built for simple, reliable and
efficient software — especially networked services and cloud tooling.

## 1. Fast build times
- A small language and a fast compiler mean builds finish in seconds, even for
  large codebases.
- That keeps Docker image builds and CI pipelines quick, shortening the
  developer feedback loop.

## 2. Fast startup
- A Go program compiles to a single self-contained native binary — no VM or
  interpreter to boot — so processes start in milliseconds.
- Fast startup makes containers cheap to scale up and down, which suits
  Kubernetes, autoscaling and serverless workloads.

## 3. Performance and efficiency
- Code is compiled ahead of time directly to machine code.
- A compact runtime, an efficient garbage collector and value-oriented data
  structures keep CPU and memory usage low.

## 4. Concurrency model
- Goroutines are lightweight user-space threads (a few KB of stack, grown on
  demand), so hundreds of thousands can run at once.
- The Go scheduler multiplexes goroutines onto a small pool of OS threads
  (`GOMAXPROCS`, defaulting to the number of CPU cores), so every core is used
  without manual thread management.
- Channels and `select` make communication explicit: "Do not communicate by
  sharing memory; instead, share memory by communicating."

## 5. Static typing and compilation
- Types are checked at compile time, so type mismatches, typos and unused
  variables/imports are caught before the code reaches production.
- The output is one binary, trivial to deploy — nothing to install on the target.

## Trade-offs worth knowing
- Verbose error handling (`if err != nil`) and a deliberately small feature set.
- Generics only arrived in Go 1.18; the type system is still minimal.
- Weaker GUI/desktop and heavy ML/numerics ecosystem than some alternatives.
