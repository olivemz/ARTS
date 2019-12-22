## Review
### Centrifuge: a reliable system for delivering billions of events per day
https://segment.com/blog/introducing-centrifuge/ \
An introduce to a fullly distributed job scheduler that built by segment. And it analytics the different approaches.
```
The problem with using any sort of queue is that you are fundamentally limited in terms of how you access data. After all, a queue only supports two operations (push and pop).
```
### Queue approach to resolve this issue
1. A single queue
backpressure issue.
2. queues per desitnation
2.1. hard to fairly use the resource, some desitination is more busy than others. 
2.2. Since it share same source, when reach 429 rate limit error from third party, several options are available.
we can try and send a hard cap of 1,000 messages per second, but this delays traffic for B and C by 50 seconds.\
we can try and send more messages to the API for customer A, but we will see 429 (rate limit exceeded) errors. We’ll want to retry those failed messages, possibly causing more slowdowns for B and C.\
we can detect that we are nearing a rate-limit after sending 1,000 messages for customer A in the first second, so we can then copy the next 49,000 messages for customer A into a dead-letter queue, and allow the traffic for B and C to proceed.\
Either of above approaches are not ideal.
3. queue per source and destination
end up with a lot queue to manage

### Getting to virtual queues
1. maintain thousands of little queues. And among these queues, easilily decide to consume messages at different rates from customers A, B, C.
2. Allow to reorder message
3. evenly distribute the workload between different workers.


