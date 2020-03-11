# Note of Lec One

Topics: implementation, performance, fault tolerance, consistency



**CASE STUDY**: MapReduce

Examples of MapReduce: Word counter, Distributed Grep, Count of URL Access Frequency,
Reverse Web-Link Graph, Term-Vector per Host, Inverted Index, Distributed Sort..

###### Q: Do goroutines run in parallel?
###### Q: How do Go channels work? How does Go make sure they are synchronized between the many possible goroutines?
###### Q: How can I have a goroutine wait for input from any one of a number of different channels? Trying to receive on any one channel blocks if there's nothing to read, preventing the goroutine from checking other channels.