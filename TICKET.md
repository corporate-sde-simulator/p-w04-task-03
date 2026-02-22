# PLATFORM-2895: Investigate leader election failures in distributed cluster

**Status:** In Progress · **Priority:** Critical
**Sprint:** Sprint 26 · **Story Points:** 8
**Reporter:** Vikram Patel (Distributed Systems Lead) · **Assignee:** You (Intern)
**Due:** End of sprint (Friday)
**Labels:** `backend`, `golang`, `distributed`, `investigation`
**Task Type:** Code Debugging

---

## Description

Production is experiencing intermittent leader election failures. When a cluster node goes down, the remaining nodes sometimes fail to elect a new leader, leaving the cluster in a leaderless state for several minutes.

**This is a DEBUGGING task.** The bugs are in the code but there are NO hint comments. You must investigate the symptoms, form hypotheses, and find the root cause.

## Symptoms Reported

- Cluster of 5 nodes: when leader (node-1) is killed, nodes 2-5 fail to elect a new leader ~40% of the time
- Logs show `election timeout` followed by all nodes starting new elections simultaneously
- When it eventually works, the wrong node sometimes wins (highest ID instead of highest log index)
- Node manager shows stale heartbeat data even after node removal

## What You Have

- `src/bullyElection.go` — Bully election algorithm implementation
- `src/nodeManager.go` — Cluster node lifecycle management
- `tests/` — Test files that currently FAIL

## Acceptance Criteria

- [ ] Root cause identified and documented in a code comment
- [ ] All bugs found and fixed
- [ ] All unit tests pass
- [ ] Election succeeds reliably when leader goes down
