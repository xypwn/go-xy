# Go XY library

Dependency-free Go utilities.

## /it
It provides some functional iterator-related capabilities (e.g. `Map`, `Filter`, `Uniq`, map `SortedByKey` and many more).

## /profile
Profile provides a super-simple setup for when you *just* want to profile your Go applications. Running a CPU profile is as simple as `defer profile.CPU().Stop()` (note that due to how Go works, `CPU()` [which starts the profile] is not deferred, while `Stop()` is).

See /examples/profileme.

## /ds
DS provides some additional data structures (currently only `Set`).

## /tests
Tests provides a thin wrapper for Go's `testing.T` with simple equality testing.

## /text
Text provides useful text utilities not covered by std strings.

## /digraphs
Digraphs provides utilities for working with directed graphs.