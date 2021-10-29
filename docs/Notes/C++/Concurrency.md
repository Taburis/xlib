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

The **race condition** is a case that multi-threads can access data in a undefined orders. It may reasults undefined behaviors. For instance, there is a data with a structure containing part A and B that correlated with each other. It is possible that one thread modified the part A while another thread modified the part B under a race condition. It may decorrelated the A and B which causes problem. A conception called **invariant** from the race condition is that the structure of the data is unchanged after an operation performed even in a race condition. For instance, a bi-directional list is still a bi-directional list after the thread modify the node under a race condition.

Several ideas can be used to avoid the problem in a race condition:
1. Lock-free programming: Grouping the data together, make it invisible to other threads while one thread is accessing to it.
2. Transaction: Recording all data reading and modification opeartors from different threads in a transaction log. Updateing the data in one step. Restart the transaction if any conflicts happen.

Locking the data accessed by a function `func` using `mutex`:
```cpp
std::list<int> data; // the data func will access to
std::mutex locker;
void func(int value){
    // locker guard will call the locker.lock() to preserve the data accessed by the func
    // and will call locker.unlock() at the end of this function.
    std::lock_guard<std::mutex> guard(locker);
    data.push_back(value); // modify the data
}
```
A `scoped_lock` guard is avaliable in C++17 compiler.
:::danger 
Don't pass pointer or reference to protected data outside the scope of the lock, wether by returning them from a function, storing them in externally visible memory, or passing them as arguments to user-supplied functions. Programs can still access the data by these pointers or references, which makes the lock nonsense. 
:::

The race condition may also leads to the following problems:
* Race in status check: Although `mutex` can protect the data during accessing, the race condition can also cause problem. For instance, a `std::stack::empty()` may return `false`, and other thread follows imediately popped one elements that empty the stack. Since the `empty()` query happened and mutex unlocked the stack, so the pop can be performed. The same reason can cause problem for functions like `size()`, `empty()`, et al.
* Race causes the copy unsafe: `pop()` function will copy the object to the target and delete the top element in stack. However, if the copy happens after the deletion, it may cause problem that the element has been removed but the copy failed. For this problem, there is several solutions:
1. Pass in reference: create a variable to contain the value first, then pass the reference to the variable
```cpp
std::vector<int> data;
stack.pop(data);
```
2. Move constructor: If the data popped out has a move constructor or a constructor don't throw exception. Using them during pop(). This is because the exception is the only problem interupt the `pop()` in the middle.
3. A pointer to the item: The pointer can be copied without throwing an exception. To simplify the memory management, one can use the `std::shared_ptr` as the object will be delete automatically if the last pointer to the object is delete. So the basic solution is to **adjust the granlarity of lock on data**. Many of the problems above can be classified a too small lock granularity. But the trade off increasing the granularity can be the concurrency performance degeneration. 

Here is a example of thread safe stack
```cpp
#include <exception>
#include <memory>
#include <mutex>
#include <stack>

// define a exception of the empty stack
struct empty_stack: std::exception{
    const char * what() const throw();
};
template <typename T>
class threadsafe_stack{
    private:
        std::stack<T> data;
        mutable std::mutex m;
    public:
        threadsafe_stack(){}
        threadsafe_stack(const threadsafe_stack & other){
            std::lock_guard<std::mutex> lock(other.m);
            data = other.data;
        }
        //make the stack uncopiable to increase the safety.
        threadsafe_stack & operator = (const threadsafe_stack &) = delete;
        void push(T value){
            std::lock_guard<std::mutex> lock(m);
            data.push(std::move(value));
        }
        std::shared_ptr<T> pop(){
            std::lock_guard<std::mutex> lock(m);
            //checking if the stack is empty to avoid pop on 0 size stack
            if(data.empty()) throw empty_stack();
            // return the pointer to prevent the throw during copy
            std::shared_ptr<T> const res (std::make_shared<T>(data.top()));
            data.pop();
            return res;
        }
        // provide other method for fexibility
        void pop(T& value){
            std::lock_guard<std::mutex> lock(m);
            //checking if the stack is empty to avoid pop on 0 size stack
            if(data.empty()) throw empty_stack();
            value = data.top();
            data.pop();
        }
        // this empty function might be redundant in this threadsafe stack
        bool empty() const {
            std::lock_guard<std::mutex> lock(m);
            return data.empty();
        }
};
```

A **dead lock** can also happen: Suppose two threads are locked and their next destiny is the data hold by each other. They will have to wait the other to release so to proceed. But this won't happen and caused the dead lock.