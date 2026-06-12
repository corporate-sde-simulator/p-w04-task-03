# Beginner Explanatory Guide: PLATFORM-2895: Investigate leader election failures in distributed cluster

> **Task Type**: Product Task  
> **Domain/Focus**: Distributed Systems, Golang

---

## 1. The Goal (In-Depth Beginner Explanation)

### The Core Problem
In a distributed system, multiple nodes (or computers) work together to perform tasks. One of these nodes is designated as the "leader," which coordinates actions and ensures that all nodes are synchronized. The task at hand addresses a critical issue where, when the leader node fails (for example, if it crashes or is shut down), the remaining nodes sometimes fail to elect a new leader. This failure can leave the cluster in a state where no node is in charge, causing delays and potential data inconsistencies.

The symptoms reported indicate that when the current leader (node-1) is killed, the other nodes (nodes 2-5) attempt to elect a new leader but fail approximately 40% of the time. This is problematic because it can lead to a situation where the cluster is unable to process requests, resulting in downtime for users. Additionally, when a new leader is eventually elected, it may not be the most suitable candidate, as the election process is flawed. Fixing this issue is crucial for maintaining the reliability and efficiency of the distributed system, ensuring that it can recover quickly from node failures.

### Jargon Buster (Key Terms Explained)
* **Leader Election**: This is a process used in distributed systems to designate one node as the leader among a group of nodes. The leader is responsible for coordinating tasks and making decisions. For example, in a cluster of five nodes, if node-1 is the leader and it fails, the other nodes must quickly elect a new leader to maintain operations.

* **Election Timeout**: This term refers to a specific duration that a node waits before it considers the current leader to be unresponsive or failed. If the timeout expires without receiving a heartbeat signal from the leader, the node will initiate a new election. For instance, if node-2 does not hear from node-1 within the election timeout period, it will start the process to elect a new leader.

* **Stale Heartbeat**: In distributed systems, nodes send periodic signals (heartbeats) to indicate they are alive and functioning. A stale heartbeat occurs when a node continues to report that it is active even after it has been removed from the cluster. This can lead to confusion and incorrect assumptions about the state of the cluster.

* **Bully Election Algorithm**: This is a specific algorithm used for leader election in distributed systems. In this algorithm, when a node detects that the leader has failed, it will attempt to become the new leader by communicating with other nodes. If it has a higher ID than the current leader, it will take over as the leader. For example, if node-3 has a higher ID than node-1, it will become the new leader if node-1 fails.

### Expected Outcome
After implementing the solution, the distributed system should reliably elect a new leader whenever the current leader fails. 

**Before**: When the leader node fails, the remaining nodes often fail to elect a new leader, leading to a leaderless state approximately 40% of the time. This results in downtime and potential data inconsistencies.

**After**: The system should successfully elect a new leader every time the current leader fails, ensuring continuous operation and coordination among nodes. All nodes should also have accurate heartbeat data reflecting their current state.

---

## 2. Related Coding Concepts & Syntax (50% Theory, 50% Practice)

### Concept 1: Goroutines and Concurrency in Golang
#### 📘 Theoretical Overview (50%)
* **Why it exists**: Goroutines are a fundamental feature of Golang that allow functions to run concurrently, meaning they can execute simultaneously without blocking each other. This is crucial in distributed systems where multiple nodes need to communicate and perform tasks at the same time. Without concurrency, the system would be slow and inefficient, as it would have to wait for one task to finish before starting another.

* **Key Mechanisms**: Goroutines are lightweight threads managed by the Go runtime. When you start a goroutine, it runs independently of the main program flow. The Go scheduler handles the execution of these goroutines, allowing them to yield control when waiting for I/O operations or other blocking tasks. This mechanism enables efficient use of system resources and improves performance.

#### 💻 Syntax & Practical Examples (50%)
* **Language Syntax**:
  ```go
  go functionName(parameters) // Starts a new goroutine
  ```
  In this syntax, `go` is a keyword that tells the Go runtime to execute `functionName` concurrently with the rest of the program. 

* **Real-World Application**:
  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func printMessage(message string) {
      for i := 0; i < 5; i++ {
          fmt.Println(message)
          time.Sleep(1 * time.Second) // Simulates a delay
      }
  }

  func main() {
      go printMessage("Hello from Goroutine!") // Starts a new goroutine
      printMessage("Hello from Main!") // Runs in the main thread
  }
  ```
  In this example, the `printMessage` function is called as a goroutine, allowing it to run concurrently with the main function. This demonstrates how multiple tasks can be performed simultaneously.

---

## 3. Step-by-Step Logic & Walkthrough

1. **Step 1: Locate and Analyze the Target File**
   * Navigate to the `src` folder in your project directory. The relevant files for this task are `bullyElection.go` and `nodeManager.go`.
   * Open `bullyElection.go` and inspect the sections where the leader election logic is implemented. Look for functions that handle the election process and any conditions that check for the leader's status.

2. **Step 2: Input Verification & Validation**
   * Check for edge cases in the election process. For example, ensure that the system can handle scenarios where multiple nodes attempt to start an election simultaneously. Look for conditions that might lead to race conditions or deadlocks.

3. **Step 3: Core Implementation / Modification**
   * Modify the election logic to ensure that nodes do not start elections at the same time. Implement a mechanism to manage election timeouts more effectively, ensuring that nodes wait for a clear signal before initiating a new election. This may involve adjusting the timeout values or adding additional checks to prevent simultaneous elections.

4. **Step 4: Output Verification & Testing**
   * After making changes, run the existing tests in the `tests/` directory to verify that the modifications work as intended. Use the command line to execute the tests and check for any failures. Ensure that all tests pass and that the system behaves correctly when the leader node is removed.

---

## 4. Detailed Walkthrough of Test Cases

### Test Case 1: Standard / Success Case
* **Description**: This test checks if the system can successfully elect a new leader when the current leader fails.
* **Inputs**:
  ```json
  {
      "nodes": ["node-1", "node-2", "node-3", "node-4", "node-5"],
      "leader": "node-1",
      "action": "kill",
      "target": "node-1"
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The system receives the input indicating that node-1 (the leader) is to be killed.
  2. Node-1 fails, and the remaining nodes detect the failure after the election timeout.
  3. Nodes 2-5 initiate the election process, but due to the implemented changes, they do not start elections simultaneously.
  4. The node with the highest log index is elected as the new leader.
* **Expected Output**: The new leader is successfully elected (e.g., "node-3"), and the system continues to operate without downtime.

### Test Case 2: Edge Case / Validation Fail
* **Description**: This test checks how the system handles a scenario where all nodes attempt to start an election simultaneously.
* **Inputs**:
  ```json
  {
      "nodes": ["node-1", "node-2", "node-3", "node-4", "node-5"],
      "leader": "node-1",
      "action": "kill",
      "target": "node-1"
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The system receives the input to kill node-1.
  2. All remaining nodes detect the failure and attempt to start an election at the same time.
  3. The election logic detects the simultaneous attempts and prevents multiple elections from occurring.
  4. The node with the highest ID or log index is elected as the new leader without conflicts.
* **Expected Output**: The system successfully elects a new leader without any errors or conflicts, demonstrating that the election process is robust against simultaneous attempts.