# New Features in C++ Library

## Rvalue 

The lvalue and rvalue stands for the value can appear on the left and right side of assigment, respectively. The reference to lvalue and rvalue is
```cpp
int & lref;  // lvalue reference
int && rref; // rvalue reference
int const & ref; // const will convert the ref to the rvalue reference
```
`std::move` function is equivalent to the `static_cast<T&&>` that will cast the object to a rvalue. The purpose for creating the rvalue conception is to solve at least the two problems:
1. Implementing the move semantics.
2. Perfect forwarding.

### Move Semantics
A move semantics means that an object can be moved from one lvalue variable `A` to another say `B`. It actually includes two steps, copy the object in `A` to `B`, and delete the object in `A`. But move semantics stands for a more efficient implementation that swaping the object held by `A` and `B`, and then discard the object `A`.  To implement the move semantics, a rvalue reference `obj&&` is needed here to distinguish the lvalue assignment:
```cpp
obj& obj::operator = (obj && rhs){
    //swap the values between this and rhs.
}
```
In this case, a lvalue input should trigger a normal copy assignment while the rvalue should trigger the move semantics. To ensure a move semantics for an object, the rvalue copy constructor and rvalue assignemnt operator should be defined:
```cpp
//classical copy and assignment constructor
obj::obj(obj & rhs); 
obj & obj::operator = ( obj const &rhs); 
// this func can be call by both lvalue and rvalue, which is ambiguous
//rvalue copy and assignment for move semantics
obj::obj(obj && rhs);
// This extra definition is needed to element the ambiguous for rvalue
obj & obj::operator = ( obj && rhs); 
```
Then calling `obj x = std::move(y)` will call the assignemnt with `std::move(y)` as input first, and follow the copy constructor `obj(ojb &&)`.
:::tip
It should be careful when we implement the rvalue assignemnt. Sometimes we don't delete the object in `lhs` and it will just hold somewhere in the scope and be deleted after the scope closed. But the side effect brings from their destructor can be problematic. So the side effects should be cleared before we lost the access to it.
:::

A rvalue reference is a lvalue if it is held by a variable, otherwise, it is a rvalue:
```cpp
obj && c = obj();
// c will be treated as a lvalue but obj() is a rvalue. 
obj x(c);
// this will call the constructor obj(obj&) instead of obj(obj&&) for rvalue.
```

### Perfect Forwarding
For a template function:
```cpp
template <typename T>
void func(T&& x);
```
The actual type of reference is detemined by the following reference collapsing rules:
* `T& &`  --> `T&`
* `T&& &` --> `T&`
* `T& &&` --> `T&`
* `T&& &&`--> `T&&`

For instance, `func(obj&)` means that `T=obj&` so that if `x` is a lvalue reference, then it will be deducted into `obj &` in `func(x)`.  This feature can be used in a perfect forwarding function `std::forward` defined as
```cpp
template<class S>
S&& forward(typename remove_reference<S>::type& a) noexcept
{
  return static_cast<S&&>(a);
} 
```
will `noexcept` tells the compiler that this function won't throw any exception. This forward function preserves the original value type for `arg` when calling `std::forward<Arg>(arg)` no matter if `arg` is a lvalue or rvalue referred to the object `Arg`.

### Delete and Defaulted Functions

For some cases, a perticular function is not allowed for an object. It can be specified by the key word `detele` to prevent the compiler to pick it up. 
Here are two major usage:
* Prevent the implicit constructor from compiler:
```cpp
class obj{
  public:
    obj(){};
    // removed the copy function to make the object non-copyable
    obj(obj const & x) = delete;
    obj& operator = (obj & x) = delete;
};
```
* To avoid implicit cast to restrict the input types:
```cpp
// prevent the implicit convertion for int as inputs
void func(short);
void func(int) = delete;
```