# PR Review - Leader election using Bully algorithm (by Deepak)

## Reviewer: Vikram Patel
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `bullyElection.go`

> **Bug #1:** Election starts even when higher-ID node is alive but should only start when coordinator is down
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `nodeManager.go`

> **Bug #2:** Victory message is sent before all lower-ID nodes have responded to election
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Deepak**
> Acknowledged. I have documented the issues for whoever picks this up.
