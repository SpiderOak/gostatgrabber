/*
go language package for a StatGrabber client

If you use Ganglia or other visual systems monitoring tools, you already know
the benefit of seeing real-time graphs of your CPU use, memory, disk activity,
network packets, and so on.

StatGrabber gives you the ability to visualize the internal operation of your
own software along side of all of your existing system stats. See real-time
graphs of connected users, application transactions, revenue received, or any
other quantitative values useful to you.

It's often very helpful to see the transactions or events your own software is
handling in real time relationship to system CPU, memory, etc.

The client modules simply emit non-blocking UDP packets and get on with their
business, avoiding slowing down their response time.

You can graph 4 types of stats:
 * counters (ex: transactions, revenue)
 * averages (ex: size of transactions)
 * accumulators (ex: bandwidth used) and elapsed
 * time (ex: time per transaction)
*/
package statgrabber
