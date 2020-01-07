# Reasons for concurrency and parallelism 

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency is about an application being able to execute a minimum of two tasks or more virtually at the same time.  
 Parallelism on the other hand does not need two or more tasks to individually exist. It physically runs parts of the same task/multiple tasks at the same time using multi-core infrastructure of the CPU, by assigning one core to each sub-task.
 So the difference is that Concurrency is about two tasks being able to run and complete in overlapping time periods, wihle parallelism is tasks running at the same time on a multi-core processor. 
 
 
 ### Why have machines become increasingly multicore in the past decade?
 > There are a lot of challenges in the development of CPUs and one of them is increasing performance without raising the processor clock speed. And as a result Multi-core processors have become more common. 
The advantages of Multi-core CPU:
>- Works faster for certain programs
>- Computer will not get as hot when it is turned on
>- Computer needs less power because it can turn off some section if they are unused
>- Parallelism 
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > All programs that are designed for multi-core CPUs should neccesarily use concurrency as otherwise only one core will be used and the rest of the CPU will be sitting there and doing nothing. Motivation: Any desktop app, Mobile apps, Video games for the playstation or any other console gaming platform, even cloud computing to able to manipulate multiple servers if the platform you own ends up growing. 
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > It makes the programmer's life harder because as the system scales the more complex it becomes. Also because of the unpredictible nature of tracking down bugs and reproducing them as well as unpredictable scheduling
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > Processes - OS-managed truly conccurent at least in the presence of suitable hardware support. Exists within  their own address space.
 
 > Threads - OS-managed, within the same address space as the parent and all its other threads. Possibly truly concurrent, and multi-tasking is pre-emptive.
 
 > Green threads - These are user-space projections of the same concept as threads, but are not OS-managed. Not truly concurrent, except in the sense that there may be multiple worker threads or processes giving them CPU time concurrently, so this can be considered interleaved or multiplexed.
 
 > Coroutines - This is a general control structure whereby flow control is cooperatively passed between two different routines without returning. Example: During coroutine A's execution, it passes control to coroutine B. Then after some time, the coroutine B passes control back to coroutine A. Since there is dependency between coroutines, and they must run in tandem, they can therefore not be concurrent.
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
> pthread_create() - this creates a new thread.
> threading.Thread() - Creates an instance of the threading.Thread class
> go - Creates a goroutine in golang runtime. This is a way for concurrency
>They all of sort of create? (i am not sure i fully understood this question)
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > *Your answer here*
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > *Your answer here*
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > *Your answer here*
