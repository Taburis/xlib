# Concurrency

The purposes to program using concurrency are: dividing the concerns or functionality, and performance (parallelism). The latter is taking the advantages of multi-cores can handle multi-tasks at the same time. The former purpose is separating different functionalities required to run simultaneously into subprocess or thread. 

A process can have many threads which can sharing memories to each other. But communicate between processes is tricky as each process are protect by the system. Files or other ports may needed for the inter-process communication.

The approach to the concurrency can be classified into the following types:
1. Multi-processes: Dividing the functionality or programs into multiple, separate processes. Inter-communicating through the system operation.
2. Multi-threads: Implement the application in one process but using seveeral threads as light weighted processes. The threads communicate through the shared memories. (Supported since C++ 11)

## Threads Management

Launching a thread:
```cpp
#include <thread>
class task; // the opeartor () is define for this object
void work_to_do()
void main(){
    task x;
    // passing the object x which x() is defined to a thread
    std::thread t(x);
    t.detach(); //let it run in the background
    // passing the 
    std::thread my_thread(work_to_do)
    my_thread.join(); // make the main wait my_thread to finish before close
}
```
Noticing that if passing the class to thread by using the default constructor is 
```cpp
std::thead my_thread((task()));
// or
std::thead my_thread{(task())};
```
to prevent the confusion from defining a function named `my_thread` with returning a thread and taking a function pointer with shape `f()` as the input.

Using `my_thread.join()` to let the function to wait the finish of `my_thread` before it move on. Each thread can only be joined once. A `thread_guard` can be used to prevent the function closed before the thread finished:
```cpp
class thread_guard
{
    std::thread &t;
    public:
        explicit thread_guard(std::thread &t_): t(t_) {}
        ~thread_guard(){
            if(t.joinable()){ t.join()}
        }
        // disable the default constructor to avoid the automate creation by the compiler
        thread_guard(thread_guard const &) = delete; 
        // disable the assignment or copy to contrain the life scope for thread_guard
        thread_guard & operator = (thread_guard const &) = delete;
};
// usage this guard
void work(); // work needs to be run in parallel
void f(){
    std::thread t(work);
    thread_guard g(t);
    other_work_in_current_thread();
}
``` 
The destructor of `g` will be called when the `f` reach to the end, it will join the thread `t` if it was still runing so that it won't be terminated due to the termination of `f`.

`my_thread.detach()` will let the thread run in the background. It will not joinable once it has been detached. 

Passing arguments to a thread can be done in the following types:
```cpp
void f(std::string const &);
void func(data &);
void oops(){
    char buffer[1024];
    sprintf(buffer, "%i%", some_param);
    // convert the char pointer into std::string explicity,
    // otherwise the buffer will be converted to std::string in the detached thread.
    // It has the risk that the function oops may finish before the buffer has been converted. 
    std::thread t(f, std::string(buffer));
    t.detach();

    data x;
    //use std::ref to force it pass the reference instead of copy the object.
    std::thread t2(f, std::ref(x));
    t2.join();
}
```

The thread can be past as variables by using `std::move`. It can be past into or out of a function. 
```cpp
std::thread t1 (some_work);
// change the owner from t1 to t2
std::thread t2 = std::move(t1);
t1 = std::thread( other_work);
// if the thread hold by t2 was still running, it will be terminated.
t2 = std::move(t1);
```

To prevent the owner changing causes the confusing of the thread lifetime, a conception the scope thread can be implemented as
```cpp
class scope_thread
{
    std::thread t;
    public : 
        explicit scope_thread(std::thread t_): t(std::move(t_)){
            if(!t.joinable()){
                throw std::logic_error("No thread");
            }
        }
        ~scope_thread(){
            t.join();
        }
        //delete the copy and assignment to prevent the scope changing for scope_thread
        scope_thread(scope_thread const &) = delete;
        scope_thread & operator = (scope_thread const &) = delete;
};
// usage
class job;
void f(){
    scope_thread t{std::thread(job())};
    the_rest_work_in_current_thread();
}
```
In this case, the thread is kept by the `scope_thread` object and ensured to be finished before the scope is closed. This is important if the job may changing any variables in current scope, it will cause segmentation problem once the scope is closed but the thread is still trying to access these variables.

Each thread has a unique `std::thread::id` can be obtained by `get_id()`. The IDs can be compared but may not have any meaning as it may system depends. Only the equal of two IDs means that both threads are either the same or not exist. 

## Sharing Data