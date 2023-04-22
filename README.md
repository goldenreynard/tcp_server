# TCP Server made using Go

- TCP is the most reliable way for two machines to talk to each other over the network.
- TCP servers are designed to handle and serve multiple TCP connections.
- A TCP server is a simple process that runs on a machine and listens to a port that understands TCP.

## First Part: Accepting a single request 

- First block the port: 1729 for the connection
- Made the connection open
- Make the do function which 
  - Accepts the request
  - Read the connection string
  - Process the Request
  - Write back to the server
  - Close the request

---

## Accepting Multiple Request - Single threaded

- Built on the top of the first request
- We Made the connection to run infinitely using the for loop
- Can accept the multiple request
- First Process the first incoming request then when finished move to the next one
- Single threaded i.e One Request at one time

---

## Multithreading - Parallel Processing

- Accepting multiple request at the same time
- When accepting the request the function which is responsible for processing the request will go to the other thread
- Main thread then can process the next incoming request
- Server can process the multiple request at the same time

> if Large numbers of the request comes then it will be thread overloading which can crash the server limits 
---

## Improvements to the Multithreading

- Limiting the number of threads
- Connection pooling: fixed time for a client to send the request
- Making thread Pool i.e Fixed number of the thread available
- TCP backlog queue configuration i.e how many connection you want in the backlog

> **SOURCE**: [Link](https://www.youtube.com/watch?v=f9gUFy-9uCM)