# Combinatorics

Many basic combinatorics problems can be classified, or equivalent to several basic combinatorial models.

## Counting Problem

:::note A **pigeon-hole** principle:
Let $n,m$ and $k$ be the positive integer so that $n>km$. Let us distribute $n$ balls into $m$ boxes. There will be at least one box which contained $k+1$ balls.  
:::

### Permutations

Given a tuple $a$ contained $n$ distinct elements so that it can be represented as $a=(1,2,\dots, n)$, a **permutation** $P$ is an one-on-one bijection function mapping $a$ to a new order of elements, for instance, $P(a)=(2,3,\dots, n, 1)$. A permutation can be expressed as a two-line notation:
$$
P=\begin{pmatrix}
1 & 2 & \dots & n\\
P(1) & P(2) & \dots & P(n)
\end{pmatrix}.
$$

* For an $n$-tuple with all elements are distinct, the number of distinct permutations are $n!$. 
* For a $n$-tuple has only $m$ distinct values, and $n_i$ ($i=1,\dots, m$) elements corresponds to the value $i$ satisfying the condition $\sum_in_i=n$. Then the distinct permutation is given by 
$$
\binom{n}{n1,\dots, n_m}=\frac{n!}{\prod_{j=1}^m n_j!},
$$
and this is also known as the multinomial coefficient.

:::tip Stirling's formula
The Stirling's formula gives an approxmation to the factorial $n!$
$$
n!\sim z(n)=\sqrt{2\pi n}\left(\frac{n}{e}\right)^n,
$$
which means that $\lim_{n\to\infty} \frac{n!}{z(n)} =1$. 
:::

### Multisets

A multiset is a modified conception of set, allows for multiple instances for each of its elements. The number of instance given for each element is called the **multiplicity** of that element in the multiset. For instance,  $M=\lbrace a,a,b\rbrace$ is a multiset from a set $S=\lbrace a, b\rbrace$, the multiplicity of $a$ is 2, and $b$ has multiplicity 1. The **size** $n$ of the multiset is the number of instance in it, so that $\sum_{x\in S}m_i=n$.

Given a multiset $A$ with the elements selected from the set $S$ with size $n$, labeled as $i=1,\dots, n$, the multiplicity of $i$ is denoted as $m_i$. The number of possible distinct multiset, $N$, is determined by different condition for the multisets $A$: 
* $\binom{n}{k}$ for $A$ has size $k$ and multiplicity $m_i=1$, $i=1,\dots, n$, where
$$
\binom{n}{k}=\frac{n!}{k!(n-k)!}.
$$
In this case, $A$ can be subset of $S$ with size $k$.
* If $A$ has the size $k$, the number is
$$
\binom{n+k-1}{k}.
$$

**Proof**: 
* It is straightforward, for the first element, there are $n$ candidates, but only $n-1$ candidates left for choice of the second element. So it is $n!/(n-k)!$ possible candidates. Finally, the order of the selected $k$ elements are not important so that the total number should be $n!/(k!(n-k)!)$.
* For the previous selected $k$ elements can be ordered as
$$
1\le a_1<a_2<\dots< a_k\le n,
$$ 
where $a_i$ is the $i$-th element of the multiset $A$. We knew the number of possible distinct multisets with this order is $\binom{n}{k}$. For this case, the elements of $A$ satisfies the relation:
$$
1\le a_1\le a_2\le \dots\le a_k\le n.
$$
We can make trick that changing $a'_i=a_i+i-1$ to avoid them to be equal, then we have 
$$
1\le a'_1< a'_2< \dots< a'_k\le n+k-1,
$$ 
and then the number of different possible array of $a'_i$ is can be obtained from the previous theorem, which is $\binom{n+k-1}{k}$.


### Multinomial Coefficients
The properties of binomial and Multinomial coefficients are:
1. For non-negative integer $k\le n$: 
$$
\binom{n}{k}=\binom{n}{n-k}.
$$
2. For $n\ge 0$:
$$
(x+y)^n=\sum_{k=0}^n\binom{n}{k}x^ky^{n-k},\quad n(1+x)^{n-1}=\sum_{k=1}^nk\binom{n}{k}x^{k-1}
$$
3. For $n\ge 0$:
$$
\sum_{k=1}^n(-1)^k\binom{n}{k}=0.
$$
4. For $n,k\ge 0$:
$$
\binom{n}{k}+\binom{n}{k+1}=\binom{n+1}{k+1},\quad \sum_{j=0}^{n-k}\binom{k+j}{k}=\binom{n+1}{k+1}
$$


### Partitions

A **weak composition** of $n$ is a $r$-tuple $(a_1,\dots, a_r)$ such that 
$$
\sum_{i=1}^ra_i=n.
$$
A weak composition is a **composition** is $a_i\ge 1$ for all $i=1,\dots r$.

For $n,k > 0$, we have:
1. The number weak composition of $n$ into $k$ parts is
$$
\binom{n+k-1}{n}
$$ 
2. The number of composition of $n$ into $k$ parts is
$$
\binom{n-1}{k-1}
$$
3. The number of all compositions of $n$ is $2^{n-1}$.

**Proof**:
1. This problem equivalent to the problem that distributing $n$ identical objects into $k$ labeled groups. To do so, it is equivalent to label each object by a number $i\le k$, assigned to the $i$-th group. The number of possible assignment identical to the number of multisets with size $n$ draw from a set $[k]$ (the sets consists with number from $1$ to $k$). This number is $\binom{n+k-1}{n}$ refer to the [section](#Multisets).
2. This problem can be rephrased into the number of weak compositions of $n-k$ into $k$ parts, by add 1 for each of $k$ elements in the weak composition. Noticing that $\binom{n-1}{n-k}=\binom{n-1}{k-1}$.
3. Let $k$ be the parts of compositions of $n$, then $k\le n$:
$$
\sum_{k=1}^{n}\binom{n-1}{k-1}=2^{n-1}.
$$