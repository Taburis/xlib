# Lebesgue Integral

## Definitions

The logic of the definition of the Lebesgue integral is that the lebesgue integral can be defined clearly on simple function, and extend it to general function by approximating the general function by a sequence of simple functions. 

### Definition on Simple Functions

Given a measure $\mu$ defined on a ring $\mathfrak{S}_{\mu}$ over a set $X$, a simple function is a $\mu$-measurable function $f:X\to \mathbb{R}$ that has only countable distinct values. The Lebesgue integral of a simple function $f$ over a set $A$ is defined as the summation following if it was convergent absolutely:
$$
\int_Af(x)d\mu = \sum_{y_i}y_i\mu(X_{y_i}), \quad X_{y_i}=\lbrace x\in A| f(x) = y_i\rbrace.
$$
A simple function is called $\mu$-integrable if its integral exists.

The Lebesgue integral of simple functions have the following properties:
1. Suppose $A=\cup_i B_i$ where $B_i\cap B_j=\varnothing$ and $f$ has only one value, say $c_k$ over $B_k$, then the sufficient and necessary condition for the integral of $f$ to exist is that the following summation converent absolutely
$$
\int_Af(x)d\mu = \sum_kc_k\mu(B_k).
$$
2. The integral is linear, that is for two simple functions $f,g$ defined on $A$, then $\int_Afd\mu +\int_Agd\mu = \int_A(f+g)d\mu$. And for any number $k$, $k\int_Afd\mu = \int(kf)d\mu$.
3. A simple function $f$ bounded on $A$, $M\ge |f(x)|$ over $A$, is integrable, and $|\int_Af(x)d\mu |\le M\mu(A)$.

**Proof** 
1. The summation of the integral can be expressed as
$$
\sum_iy_i\mu(X_{y_i})=\sum_iy_i\sum_{c_k=y_i}\mu(B_k)=\sum_kc_k\mu(B_k), 
$$
the second step comes from the fact that the measure is non-negative defined.
2. Suppose $A=\cup F_k = \cup G_k$ such that satisfies the condition in 1. The function $f$ has a value $f_i$ over $F_i$ and $g$ has the value of $g_j$ over $G_j$.
Then
$$
\int_A(f+g)d\mu = \sum_{i,j}(f_i+g_j)\mu(F_i\cap G_j).
$$
But $\sum_i\mu(F_i+G_j) = \mu(F_i)$ and $\sum_j\mu(F_i+G_j) = \mu(G_j)$. This conclude the prove. The linearity of the scale multiplication is a straightforward factorization. 
3. Let's $S_y=\sum_{y_i\le y}|y_i|\mu(X_{y_i})$. It is clear that $S_y$ is montonic sequence and it is bounded that $S_y\le M\mu(A)$ which prove that $S_y$ convergent to $\int_A|f(x)|d\mu$ so that the integral of $f$ is convergent absolutely. Furthermore, $|\int_Af(x)d\mu|\le \int_A|f(x)|d\mu\le M\mu(A)$. 

### General Definition

The **Lebesgue integral** for a $\mu$-measurable function $f$ is defined as the following limit if it exists:
$$
\int_Af(x)d\mu = \lim_{n\to\infty}\int_Af_n(x)d\mu,
$$
where $\lbrace f_n\rbrace$ is a sequence of simple $\mu$-measurable functions convergent to $f(x)$ uniformly. 

**Remark**: This definition is valid since the following conditions hold:
1. The limit exists for any $f_n$ uniformly convergent to $f$.
Suppose $f_n,f_m$ are functions uniqformly convergenet to $f$, then
$$
\left |\int_A f_n(x)d\mu -\int_A f_m(x)d\mu\right| \le \mu(A)\text{sup}\lbrace|f_n(x)-f_m(x)|: x\in A\rbrace.
$$
2. The integral is independent on the choice of $f_n$. 
Suppose $f_n$ and $g_m$ are two sequneces of simple functions convergent to $f$ uniformly, then
$$
\left |\int_A f_n(x)d\mu -\int_A g_m(x)d\mu\right| \le \mu(A)\text{sup}\lbrace|f_n(x)-f(x)|: x\in A\rbrace+\mu(A)\text{sup}\lbrace|g_m(x)-f(x)|: x\in A\rbrace.
$$
3. It is obviously that this definition is the same as the Lebesgue integral definition for simple functions.

**Properties**
1. For any constant $k$ and function $f$, then $\int_Akf(x)d\mu = k\int_Af(x)d\mu$.
2. Suppose that $f,g$ are integrable functions, then
$$
\int_Af(x)d\mu+\int_Ag(x)d\mu = \int_A\lbrace f(x)+g(x)\rbrace d\mu.
$$ 
3. Suppose $|f(x)|\le M$ over $A$, then 
$$
\left|\int_Af(x)d\mu\right| \le M\mu(A),
$$
which means that $f(x)$ is integrable.
4. If $f$ is integrable and is non-negative over $A$, then $\int_Af(x)d\mu \ge 0$. If $f\ge g$ and $f,g$ are integrable, then $\int_Af(x)d\mu\ge\int_Ag(x)d\mu$. Furthermore, if $m\le f(x)\le M$, then $m\le \int_Af(x)d\mu/\mu(A)\le M$. 
5. Suppose $A$ have a countable partition $A=\cup_iA_i$, where $A_i$ are pairwise disjoint and $f$ is a integral function, then
$$
\int_Af(x)d\mu = \sum_i\int_{A_i}f(x)d\mu.
$$
It also means that if $f$ is integrable on $A$, then it is integrable on any subsets of $A$.
6. Suppose $\varphi(x)$ is an integrable function on $A$, and $|f(x)|\le \varphi(x)$ over $A$, then $f$ is also integrable on $A$. 
7. Given a function $f$ defined on $A$, then the integral $\int_Af(x)d\mu$ exists if and only if $\int_A|f(x)|d\mu$. 
8. The **Chebyshev inequality**: If $\varphi(x)\ge 0$ on $A$, then
$$
\mu\lbrace x\in A: \varphi(x)\ge 0 \rbrace \le \frac{1}{c}	 \int_A\varphi(x)d\mu.
$$

**Proof**
1. This comes from the factor that scaling a seuqnce by a constant won't change the convergence.
2. Notice that this conclusion holds for simple functions, this is a straightforward conclusion from passing the limit to the approximation sequence.
3. If $f$ is bounded and a sequence of simple functions $f_n$ convergent to $f$ uniformly, then $f_n$ is bounded and integrable. The uniformly convergence of $f_n\to f$ implies that $\int_Af_n(x)d\mu\to \int_Af(x)d\mu$ which conclude the proof.
4. It is obviously if $f$ is a simple function. And this theorem follows $f$ can be approximated by the $f_n$ which approximates $f$ from the upper limit so that $\int_Af_n(x)d\mu\ge 0$ which proved this theorem. The corollary follows from the condition that $f(x)-g(x)\ge 0$, $M-f(x)\ge 0$ and $f(x)-m\ge 0$. 
5. This conclusion holds for any integrable simple function since the integrable convergent absolutely that re-order any number of items won't change the value it convergent to.On the other hand, suppose $f_n\to f$ uniformly, where $f_n$ are simple functions, then we have a $\epsilon > 0$ such that
$$
\sum_i\left | \int_{A_i} f(x)-f_n(x)d\mu \right |\le \sum_i \epsilon\mu(A_i)\le \epsilon\mu(A),
$$
while 
$$
\left |\int_Af(x) f(x) - \int_A f_n(x) d\mu\right|\le \epsilon\mu(A).
$$
It implies that $\sum_i\int_{A_i}f(x)d\mu$ convergent and $\sum_i\int_{A_i}[f(x)-f_n(x)]d\mu\le 2\epsilon \mu(A)$, which proved the theorem.
6. Suppose $f_n(x), \varphi_n(x)$ are simple functions that $|f_n|\le \varphi_n$, then the integrable of $\varphi_n$ implies that the integral of $f_n$ convergent absolutely. For the general case, we can partition the set $A$ into subsets that $|f_n(x)|\le \varphi_n(x)$ on each of the subsets. Since all of the $f_n(x)$ is integrable, the existence of $\int_Af(x)d\mu$ follows from the uniformly convergence $f_n(x)\to f(x)$. 
7. The necessity comes from the theorem 6 above since $|f(x)|\ge f(x)$. For a simple function $f$, the sufficiency is preserved by definition of integral. For a general function $f$ that $|f_n(x)-f(x)|\le \epsilon$, the approximation convergent since
$$
\int_A\Big||f_n(x)|-|f(x)|\Big|d\mu \le \int_A |f(x)-f_n(x)|d\mu\le \epsilon \mu(A).
$$
8. Noticing $A=A\setminus B \cup B$ where $B = \lbrace x\in A: \varphi(x)\ge 0 \rbrace$, then
$$
\int_A \varphi(x)d\mu = \left (\int_{A\setminus B} +\int_B\right)\varphi(x)d\mu\ge \int_B\varphi(x)d\mu\ge c\mu(B).
$$


## Exchanging the Limits

1. If a sequence $f_n(x)$ converge to $f(x)$ on $A$ and $|f_n(x)|\le \varphi(x)$ for all $n$, where $\varphi(x)$ is a integrable function on $A$, then the limit function $f(x)$ is integrable over $A$ and
$$
\lim_{n\to\infty}\int_Af_n(x)d\mu \to \int_Af(x)d\mu.
$$
2. If $f_n(x)\to f(x)$ and $|f_n(x)|\le M$, then 
$$
\lim_{n\to\infty}\int_Af_n(x)d\mu \to \int_Af(x)d\mu.
$$
3. Given a sequence monotonic increasing function $f_n(x)$, $f_n(x)\le f_m(x)$ for all $n\le m$. Suppose $f_n(x)$ is integrable and the integral is bounded
$$
\int_Af_n(x)d\mu\le K.
$$
Then
$$
\lim_{n\to\infty}f_n(x)\to f(x)
$$
almost everywhere on $A$ and $f(x)$ is integrabel
$$
\lim_{n\to\infty}\int_Af_n(x)d\mu \to \int_Af(x)d\mu.
$$
This conclusion also holds for the monotonic decreasing sequence with a lower bound on their integrals.
4. If $\varphi_n(x)\ge 0$ and 
$$
\sum_{n}\int_A\varphi_n(x)d\mu\le \infty,
$$ 
then $\sum_{n}\varphi_n(x)$ converges almost everywhere on $A$ and
$$
\sum_{n}\int_A\varphi_n(x)d\mu = \int_A\Big[\sum_n\varphi_n(x)\Big]d\mu.
$$

**Proof**
1. Since $|f_n(x)|\le \varphi(x)$, then $f_n(x)$ is integrable. The integrable of $\varphi(x)$ implies that exists a $A_m=\lbrace x: \varphi(x) > m\rbrace$ such that
$$
\int_{A_m} \varphi(x)d\mu\le \epsilon.
$$
Based on Egorov's theorem, for a number $\epsilon/m>0$, there is a decomposition $A\A_m=D\cup E$ such that $D\subset A$ such that $\mu(E)\le \epsilon$ and $f_n(x)$ convergent to $f(x)$ uniformly. Now choosing a $N$ such that for any $n\ge N$:
$$
|f_n(x)-f(x)|\le\epsilon/\mu(D),
$$
then the integral
$$
\int_A[f_n(x)-f(x)]d\mu = \left(\int_{A_m}+\int_{E}+\int_C\right)[f_n(x) -f(x)] d\mu\le 5\epsilon.
$$
2. This is a collorary from the 1st theorem with $\varphi(x)=M$. 
3. 