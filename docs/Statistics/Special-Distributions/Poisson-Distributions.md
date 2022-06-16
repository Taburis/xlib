# Poisson Distributions

Considering a process generating a non-negative integer number $x$ associated to a interval in $\mathbb{R}$ with length $w$, and let $f(x,w)$ denotes the corresponding probability. It is a **Poisson Process** if it satisfies the following assumption:
1. $f(1,w)=\lambda w+o(w)$ where $o(w)$ is a higher order small quantity comparing to $w$, i.e. $\lim_{w\to0}o(w)/w=0$. 
2. $\sum_{x=2}^\infty f(x,w)=o(w)$. 
3. The probability to get $x$ for non-overlapped intervals are independent.

:::info Properties of a Poisson process:
1. The Poisson distribution is denoted as $p(m)$ where $m=\lambda w$ and
$$
f(x,w)=\frac{m^x}{x!}e^{-m}, \quad x\in \mathbb{N}^+\cup\lbrace 0\rbrace.
$$ 
2. The moment generate function 
$$
M(t)=\exp\lbrace m e^t-1\rbrace.
$$
and hence $\mathbb{E}(X)= m$ and $\text{Var}(X)=m$. 
3. Suppose $X_i\sim p(m_i)$, and $Y=\sum_iX_i$, then $Y\sim p(\sum_im_i)$.
:::

**Proof**
1. Starting from $f(0,w)=1-\lambda w-o(w)$ and $f(0,h+w)=f(0,h)f(0,w)$, the differential of $f(0,w)$ is
$$
\frac{d f(0,h)}{dh}=-\lambda,
$$
and hence $f(0,h)=Ce^{-\lambda h}$. For $f(x,w+dw)=\sum_{k=1}^\infty f(x-k,w)f(k,dw)$, however, $o(w)o(dw)=\sum_{k=2}^\infty f(x-k,w)f(k,dw)$ so that
$$
f(x,w+dw)=f(x,w)[1-\lambda dw-o(dw)]+f(x-1,w)[\lambda dw+o(dw)]+o(w)o(dw),
$$
so we have
$$
\frac{df(x,w)}{dw}= -\lambda+\lambda f(x-1,w).
$$
By induction, we got the Poisson distribution.
2. $M(t)=\sum_xe^{tx}\frac{m^x}{x!}e^{-m} = e^{m(e^t-1)}$. $M'(0)=\mu=m$ and $\text{Var}(X)=M''(0)-\mu^2 = m$.
3. $M_Y(t)=\prod_{i=1}^ne^{m_i(e^t-1)}=\exp\lbrace \sum_{i=1}^nm_i(e^t-1)\rbrace$.
