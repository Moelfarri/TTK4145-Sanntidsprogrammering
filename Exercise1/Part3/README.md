# Reasons for concurrency and parallelism 

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency is about an application being able to execute a minimum of two tasks or more virtually at the same time.  
 Parallelism on the other hand does not need two or more tasks to individually exist. It physically runs parts of the same task/multiple tasks at the same time using multi-core infrastructure of the CPU, by assigning one core to each sub-task.
 So the difference is that Concurrency is about two tasks being able to run and complete in overlapping time periods, wihle parallelism is tasks running at the same time on a multi-core processor. 
 
 
 ### Why have machines become increasingly multicore in the past decade?
 > There are a lot of challenges in the development of CPUs and one of them is increasing performance without raising the processor clock speed. And as a result Multi-core processors have become more common. 
The advantages of Multi-core CPU:
- Works faster for certain programs
- Computer will not get as hot when it is turned on
- Computer needs less power because it can turn off some section if they are unused
- Parallelism 
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > *Your answer here*
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > *Your answer here*
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > *Your answer here*
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > *Your answer here*
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > *Your answer here*
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > *Your answer here*
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > *Your answer here*
