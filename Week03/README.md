学习笔记

# to learn

MESI协议

内存屏障

CAS 指令

所有的包都要去看官方文档


先不要关注runtime是如何实现的

# Goroutine


管好生命周期，一定在掌握之中。

重点在怎么把goroutine用好。

## Processes and Threads

进程：
所有资源而运行的容器。这些资源包括内存地址空间、文件句柄、设备和线程。

线程
是操作系统调度的一种执行路径，用于在处理器执行我们在函数中编写的代码。
一个进程从一个线程开始，即主线程，当该线程终止时，进程终止。这是因为主线程是应用程序的原点。
然后，主线程可以依次启动更多的线程，而这些线程可以启动更多的线程。

操作系调度：
无论线程属于哪个进程，操作系统都会安排线程在可用处理器上运行。

## Goroutine有什么区别？

go 可以启动轻量级的线程，其实就是创建了一个goroutine。
goroutines 在绑定到单个操作系统线程的逻辑处理器中运行(P)。
即使使用这个单一的逻辑处理器和操作系统线程，也可以调度数十万 goroutine 以惊人的效率和性能并发运行。

数十万 goroutine 无非是在内存里创建了用户态伪线程。

## 并发和并行

并发不是并行。

并行是指不同执行单元。
并发是指同时执行。

## Keep yourself busy or do the work yourself

例子：
main函数如果退出了。。。。搞个select {} 则使得main函数不退出。--不鼓励这种做法。
log.Fatal/os.exit会使得defer无法正常执行。只在Main和init函数调用。

总之，go func出去完全不知道什么时候结束。

## 启动goroutine 一定是调用者，而不是函数内部。

## 每次都要问2个问题

Any time you start a Goroutine you must ask yourself:

When will it terminate?
What could prevent it from terminating?


# Memory model

本章最重要的。要彻底搞清楚。

什么叫原子？什么叫可见性？happens before-- 要搞清楚谁先执行，谁后执行。

涉及到底层CPU指令，一个多核CPU怎么保证原子递增？
MESI协议。


# Package sync

CAS指令

# chan

go多线程有更优雅的方式解决锁的问题。

# context

context引入之后可以真正非常方便的解决级联的取消。而且是显式的传递






