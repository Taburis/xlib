# Concurrency

The purposes to program using concurrency are: dividing the concerns or functionality, and performance (parallelism). The latter is taking the advantages of multi-cores can handle multi-tasks at the same time. The former purpose is separating different functionalities required to run simultaneously into subprocess or thread. 

A process can have many threads which can sharing memories to each other. But communicate between processes is tricky as each process are protect by the system. Files or other ports may needed for the inter-process communication.

The approach to the concurrency can be classified into the following types:
1. Multi-processes: Dividing the functionality or programs into multiple, separate processes. Inter-communicating through the system operation.
2. Multi-threads: Implement the application in one process but using seveeral threads as light weighted processes. The threads communicate through the shared memories. (Supported since C++ 11)

## Threads Management
---
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
---
The **race condition** is a case that multi-threads can access data in a undefined orders. It may reasults undefined behaviors. For instance, there is a data with a structure containing part A and B that correlated with each other. It is possible that one thread modified the part A while another thread modified the part B under a race condition. It may decorrelated the A and B which causes problem. A conception called **invariant** from the race condition is that the structure of the data is unchanged after an operation performed even in a race condition. For instance, a bi-directional list is still a bi-directional list after the thread modify the node under a race condition.

Several ideas can be used to avoid the problem in a race condition:
1. Lock: Using a lock to protect the data accessing by one thread from others.
2. Lock-free programming: Grouping the data together, make it invisible to other threads while one thread is accessing to it.
3. Transaction: Recording all data reading and modification opeartors from different threads in a transaction log. Updateing the data in one step. Restart the transaction if any conflicts happen.

Besides of problematic race condition, a **dead lock** can also happen: Suppose two threads are locked and their next destiny is the data hold by each other. They will have to wait the other to release so to proceed. But this won't happen and caused the dead lock. The following methods can be used to safely share data between threads.

### Lockable Objects
A value `m` is a lockable type if the following functions are defined: `m.lock()`, `m.unlock()`, and `m.try_lock()`. 

Lockable objects provided by C++ STL `<mutex>` header:
* `std::mutex`: A lock object provided `lock()`, `unlock()` functions to hold and release the data accessing by the scoped function.
* `std::shared_mutex`: Similar as the `std::mutex`, but can provide read-only accessibility for many threads with `std::shared_lock<std::shared_mutex>`. It also can provide the exclusive accessibility with `std::lock_guard<std::shared_mutex>` or `std::unique_lock<std::shared_mutex>`.
* `std::recursive_mutex`: Similar as the `std::mutex`, but this one can call `lock()` multiple times. But release the object will need the same times of calling `unlock()`.  
* `std::lock_guard<typename T>`: A guard object can accept `std::mutex` and `std::shared_mutex` to provide a exclusive accessibility. It unlocks the data when it is going to be destructed.
* `std::unique_lock<typename T>`: A object similar to the `std::lock_guard<typename T>` with a unique semantics. It is moveable but not copyable. It can be used to transfer the mutex ownership between scopes.
* `std::shared_lock<typename T>`: A guard provides the read-only multi-accessibility with `std::shared_lock<std::shared_mutex>`.
* `std::lock(...)`: A function can lock multi-mutex at once without deadlock risk.

::: tip
The shared_mutex can be used to lock the data rarely be updated by using `std::shared_lock<std::shared_mutex>`. This provides read-only multi-threads accessibility. It also can provides the exclusive accessibility when the data needs to be updated by using `std::unqiue_lock<std::shared_mutex>` or `std::lock_guard<std::shared_mutex>`. 
:::

Usage of the lockable object `mutex` with `std::lock_guard<std::mutex>` or `std::unqiue_lock<std::mutex>`:
```cpp
std::list<int> data; // the data func will access to
std::mutex locker;
void func(int value){
    // locker guard will call the locker.lock() to preserve the data accessed by the func
    // and will call locker.unlock() at the end of this function.
    std::lock_guard<std::mutex> guard(locker);
    // or  std::unique_lock<std::mutex> lk(locker);
    data.push_back(value); // modify the data
}
```
A `scoped_lock` guard is avaliable in C++17 compiler.
:::danger 
Don't pass pointer or reference to protected data outside the scope of the lock, wether by returning them from a function, storing them in externally visible memory, or passing them as arguments to user-supplied functions. Programs can still access the data by these pointers or references, which makes the lock nonsense. 
:::

The ownership of the `std::unique_lock` can be transferred due to the unique semantics (move semantics but not copyable). The function own the lock can access the data blocked by the lock:
```cpp
std::unique_lock<std::mutex> get_lock(){
    extern std::mutex m; // due to the mutex is neither copyable nor movable.
    std::unique_lock<std::mutex> lk(m);
    prepare_data();
    return lk; 
}
void process_data(){
    std::unique_lock<std::mutex> lk(get_lock());
    do_something_with_data();
}
```
Here the `unqiue_lock` serves like a gateway to access the data.

### Block Multiple Locks

Using `std::lock()` to lock multiple `mutex` can avoid deadlock. The `std::lock()` locks one `mutex` and try to lock another `mutex`. An exception will throw if it failed and the held `mutex` will be released. Here is an example to implement a `swap` function:
```cpp
class X{
    private:
        std::mutex m;
        data_type data;
    public:
        X(data_type &d): data(d){}
        friend void swap(X& lhs, X& rhs){
            if(&lhs==&rhs) return;
            std::lock(lhs.m, rhs.m);
            // using std::adopt_lock flag indicates that the mutex is already locked.
            std::lock_guard<std::mutex> lock_a(lhs.m, std::adopt_lock);
            std::lock_guard<std::mutex> lock_b(rhs.m, std::adopt_lock);
            swap(lhs.data, rhs.data);
        }
        // using the unique_lock to implement the swap is a little bit different
        friend void swap(X& lhs, X& rhs){
            if(&lhs==&rhs) return;
            // std::defer_lock will let the mutex remain unlocked.
            std::unique_lock<std::mutex> lock_a (lhs.m, std::defer_lock);
            std::unique_lock<std::mutex> lock_b (rhs.m, std::defer_lock);
            std::lock(lock_a, lock_b);
            swap(lhs.data, rhs.data);
        }
}
```


### Granularity of Lock

Locking the function can be subtle. The problem can happen if locking the function in a very fine granularity. The race condition can happen in between calling functions and here is few problems can happen in a `stack` object: 
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
On the other side, locking functions in too coarsely will slow down the performance. If copy the data is cheap, unlock the function after copy the data can be a good idea. One example is an object concurrent safe but is expensive to initialization. It is important to secure the initialization to avoid multi-threads creation at the same time:
```cpp
std::shared_ptr<some_resource> resource_ptr;
std::mutex m;
void foo(){
    std::unique_lock<std::mutex> lk(m);
    if(! resource_ptr){
        resource_ptr.reset(new some_resource);
    }
    lk.unlokc();
    resource_ptr->do_something();
}
```
The code above presents the logic but it has potential problematic racing: The race can happen after checked the `if` statement. One double-checked locking is still not perfect solution. A `std::call_once` can be used to solve this problem:
```cpp
class X{
    private:
        data_format data;
        std::once_flag init_flag;
        void init(){ data = create_data();} // func to initialize the data
    public :
        X(){}
        void call(){
            std::call_once(init_flag, &X::init, this);
            do_something(data);
        }
}
```
The `std::call_once(std::once_flag, Callable&& f, Args && ...)` will call the function `f` with the arguments `arg` is the `once_flag` indicate no call previously. `std::once_flag` can't be copy or move.  

### Locking Order and Lock Hierarchy

One of the most efficient way to prevent dead lock is to fix the order of locking and unlocking. A concept of hierarchy lock is a lock assigned a hierachy so that a lock can only happen if the lock hierachy is lower than the current hierachy held. 
```cpp
hierachical_mutex high_level_m(1000), low_level_m(500), other_level_m(600);
int do_low_level_stuff();
int low_level_func(){
    // the function be called by other functions should have lower hierarchy;
    std::lock_guard<hierarchical_mutex> lk(low_level_m);
    return do_low_level_stuff();
}
int high_level_func(){
    // the high hierarchcial lock used for high level function;
    std::lock_guard<hierarchical_mutex> lk(high_level_m);
    do_high_level_stuff(low_level_func());
}
```
In this way, the order of locks are defined by the function levels. Any case calling a high level locks in lower level function should throw an exception in run time, and hence will not be able to cause a dead lock. The implementation of the hierachical lock is not part of STL but easy to make:
```cpp
class hierarchical_mutex{
    std::mutex internal_mutex;
    // assign the hierachy value for each hierarchical mutex
    unsigned long const hierarchy_value;
    // track the previous thread hierarchy value, it needs to be stored
    // if the current mutex is unlocked.
    unsigned long previous_hierarchy_value;
    //a thread-global variable to keep the hierarchy value of current mutex
    static thread_local unsigned long this_thread_herarchy_value;
    void check_for_herarchy_violation(){
        if(this_thread_hierarch_value <= hierarchy_value)
            throw std::logic_error("mutex hierarchy violated");
    }
    void update_hierachy_value(){
        previous_hierarchy_value = this_thread_hierarchy_value;
        this_thread_hierarchy_value = hierarchy_value;
    }
    public: 
        explicit hierarchical_mutex (unsigend long value):
            hierarchy_value (value),
            previous_hierarchy_value(0){}
        void lock(){
            check_for_hierarchy_violation();
            internal_mutex.lock();
            update_hierarchy_value();
        }
        void unlock(){
            //by the definition, if the current hierarchy value isn't the hierarchy
            //value of this mutex, the current thread is locked by other mutex so 
            //unlock this mutex will mess up the lock hierarchy.
            if(this_hierarchy_value != hierarchy_value)
                throw std::logic_error("mutex hierarchy violated");
            this_thread_hierarchy_value= previous_hierarchy_value;
            internal_mutex.unlock();
        }
        bool try_lock(){
            check_for_hierarchy_violation();
            if(!internal_mutex.try_lock()) return false;
            update_hierarchy_value();
            return true;
        }
};
//init the current thread hierarchy with the maximum value to allow any hierachy lock
// is allowed at the begining of the thread. 
thread_local unsigned long 
    hierarchical_mutex::this_thread_hierarchy_value(ULONG_MAX);
```

## Synchronization
---
Several methods can be used to synchronize the threads:


### Conditional variables

This object is contained in the library`<condition_variable>`. It is primarily used for conditional synchronizing between threads. 
```cpp
std::condition_variable cond;
std::mutex mut;
dataType data;
void data_prepare(){
    {
        std::unique_lock<std::mutex> lk(mut);
        prepare_some_data(data);
    }
    // notify_one will notify one of many threads contained the condidtion variable in wait.
    cond.notify_one();
    // notify_all is an alternative way to notify all the waited condition variables.
    cond.notify_all();
}
void process_thread(){
    std::unique_lock<std::mutex> lk(mut); // using unique_lock here for feasible unlock
    // cond will triggered after been notified, the wait function will lock the mutex and check
    // if the  condition is satified. If so, the thread is activated. If not, then it will unlock 
    // the mutex and put the thread into waiting status.
    cond.wait(
        lk, []{return !data.isready();});
    process(data);
    lk.unlock();
}
```

### Future as Placeholder

This library builds threads running with return objects. The threads accessing to these returned objects are automatically synchronized. This library is `<future>` and the returned object wrapper is `std::future<T>` where `T` stands for the returned value type. 

* Asynchronized task `std::async<T>`: Starting a new thread to process the callable function and return in the future. The starting time is customable by the options `std::launch`:
```cpp
int func(data && x);
data x;
// return from async is contained in a future object
std::future<int> f1 = std::async(std::launch::async, func, std::ref(x)) // run in new thread
auto f2 = std::async(std::launch::deferred, func, std::ref(x))// run in wait() or get()
// Implementation chooses
auto f3 = std::async(
    std::launch::deferred | std::launch::async, 
    func, std::ref(x))
```
For deffered, the thread might not be executed ever if it hasn't been used later.
* Function wrapper `packaged_task<T>`: Associated a function to a future. Since the packaged task is moveable, it wrap a callable function and pass it to other threads so that the function can be called concurrently and the returned value can be used in future:
```cpp
std::deque<std::future<data>> queue;
std::mutex mut;
data generate_data(Para & par);
Para par;
void thread_process(){
    for(auto i =0; i<N; i++){
        std::packaged_task<data> task (generate_data, par);
        {
            std::unique_lock<std::mutex> lk(mut);
            // store the returned value contained in the future queue.
            queue.push_back(task.get_future());
        }
        task();
    }
}
// parallel thread to generate the data
std::thread t1(thread_process);
void process_data(){
    while(!queue.empty()){
        std::unique_lock<std::mutex> lk(mut);
        // call get() to retrive the data
        data x = std::move(queue.get());
        lk.unlock();
        processing_data(x);
    }
}
```
* Promising a value `T` in the future `std::promise<T>`: If the program ensures that a variable with type `T` is calculated somewhere in the future, this data can be contained by the `std::promise<T>` to asynchronize the threads. The value for the promised variable can be set by calling `set_value(x)` and this will mark the `promise` to be ready to use. The associated `future` can be obtained from `get_future()`. It can also handle the exception:
```cpp
extern std::promise<double> p1;
try {
    p1.set_value(calculated_value());
}catch{
    p1.set_exception(std::current_exception());
}
```
`promise` enables a more flexible programming synchronization. But it needs more lines for the task can be handled by `std::async` or `std::packaged_task`. 

:::tips Exception Propagation
If there is an exception occurs in return for `furture`, this exception is stored in the `future` as well. Calling the function `get()` for that `future` will throw the same exception again. Destroye the `promise` or `packaged_task` without calling the `set_value()` or invoking the `packaged_task` will also store a `std::future_errc::broken_promise` exception in the associated state.
:::
* Share the future with multi-thread using `std::shared_future<T>`: a `future` object is moveable but not copyable. After the `get()` is called, no value left in that future so no other thread can access to it. To make sure the result is accessible to multi-thread, a shared future can be used instead:
```cpp
std::promise<int> p;
std::future<int> f (p.get_future());
// implicitly cast the future into a shared future
std::shared_future<int> f1(p.get_future());
// call share() to convert the future into a shared future
auto f2 = f.share();
```

### Chrono Library
The date the time utilities are defined in the library `<chrono>`. A **steady** clock is a clock with uniform rate ticks and the time can not be adjusted.

* Specify the tick by `std::ratio<a, b>`: The tick interval is given as every `a/b` second. `chrono` also provides several predefined durations like:
```cpp
std::chrono::nanoseconds;
std::chrono::milliseconds;
std::chrono::seconds;
std::chrono::minutes;
std::chrono::hours;
```
* Duration `sd::chrono::duration<type, std::ratio<a,b>>(m)`: This specify the time interval as `m` specified ratio. The `type` here is the `short, double,int`, etc. Specify the data type stored the duration.  The `std::chrono::duration_cast<>` can convert one type of duration into another. If the digits are longer than the data type of the accepter, then the digits are truncated.
* Time point `std::chorno::time_point<clock, duration>`: The time point is a moment on the associated `clock`. It usually specified by a duration since the clock's epoch. The **epoch** is a time point used as start time for timing convenience. For a given clock, the epoch might not be clear but you can still call the `time_since_epoch` for a given time point. Or the time point can also be obtained by the addition/subtraction between two time points:
```cpp
auto start = std::chrono::high_resolution_clock::now();
auto stop = std::chrono::high_resolution_clock::now();
auto new_time_point = start+std::chrono::milliseconds(500);
```
* `_util` and `_for` extension: These extensions exist for functions like `wait()` or `sleep()`. `_util` extension can be used with time point and the `_for` is used with the duration. For the case that `wait_util(timepoint)`, a `timeout` status will pass to the function.

### Functional Programming with Futures.

The functional programming (FP) makes each program like a function which return the output depends only on the input. No external data modification so that no data race and time race problem. This advantage makes it a good way for concurrent programming. Using a quicksort algorithm as an example:
```cpp
template<typename T>
std::list<T> parallel_quick_sort (std::list<T> input){
    if(input.empty()) return input;
    std::list<T> result;
    result.splice(result.begin(), input, input.begin());
    T const & pivot = *result.begin();
    auto divide_point = std::partition(input.begin(), input,end(),
        [&] (T const &t) {return t<pivot;});
    std::list<T> lower_part;
    lower_part.splice(lower_part.end(), input.input.begin(), divide_point);
    std::future<std::list<T>> new_lower(
        // the sorting on the lower half can be done by a new thread.
        // this will create a new thread for each iteration. But if the new threads
        // are too many to handle, the compiler may generate a task instead (deferred mode)
        std::async( &parallel_quick_sort<T>, std::move(lower_part)));
    auto new_higher(
        parallel_quick_sort(std::move(input)));
    result.splice(result.end(), new_higher);
    result.splice(result.begin(), new_lower.get());
    return result;
}
```

You can also use a customized wrapper to integrate the `packaged_task` and `thread`:
```cpp
template<typename F, typename A>
std::future<std::result_of<F(A&&)>::type> 
    spawn_task(F&& f, A&& a){
        typedef std::result_of<F (A&&)>::type result_type;
        std::packaged_task<result_type(A&&)> task(std::move(f));
        std::future<result_type> res(task.get_future());
        std::thread t(std::move(task), std::move(a));
        t.detach();
        return res;
    }
```

**Finite State Machine**  is an object with finite states. It can receive inputs, update its states, and generates the outputs. This model is perticularly well fit for synchronizing operations without sharing data. 

### Experimental Library

The namespace `std::experimental` provided a continuations for `std::experimental::future`: `then(Callable &&f)`. This function will take the future as inputs for the function `f`:
```cpp
int process_result(std::experimental::future<int> input);
std::experimental::future<int> result;
auto fut2 = result.then(process_result);
// after then is called, the result is no longer valid. If you want to acess the future by multiple times,
// a shared_future is needed; The callable function needs to take the shared_future as input.
auto fut = result.share();
int process_shared_result(std::experimental::shared_future<int> input);
auto fut3 = fut.then(process_shared_result);
auto fut4 = fut.then(so_other_stuff);
``` 

Waiting for multiple futures is also implemented in the `experimental` library:
* `when_all`: It takes iterator begin and the end as input to determine if all futures between them are ready. If so it return a future with a type of vector containing all the ready futures:
```cpp
std::experimental::when_all(results.begin(), results.end()).then(
    [](std::future<std::vector<std::experimetnal::future<int>>> ready_results){
        std::vector<std::experimental::future<int>> futures = ready_results.get();
        do_somethin(futures);
    }
)
```
* `when_any`: It triggers the waiting function when any one of the futures is ready:
```cpp
bool DoCheck(
    std::experimental::future<std::experimental::when_any_result<
    std::vector<std::experimental::future<MyData *>>>> input);
std::experimental::when_any(results.begin(), results.end())
    .then(DoCheck);
```