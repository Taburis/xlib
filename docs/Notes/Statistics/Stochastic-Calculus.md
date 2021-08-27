# Stochastic Calculus

## Ito Integral
---

Starting from the **Ito Integral definition for simple process**:
Given a probability space $(\Omega,\mathcal{F},\mathbb{P})$ and a Brownian motion $W(t)$ with a filter $\mathcal{F}(t)$, a stochastic process $f_n(t)\in\mathcal{L_2}$, $\mathbb{E}(\int_0^Tf_n^2(t)dt)<\infty$, is called **simple** if the function of $f_n(t)$ adapted to the filter $\mathcal{F}(t)$ are $f_n(t)=f_n(t_i)$ (left endpoint value) when $t\in[t_i,t_{i+1}),i=0,\dots, n$. The Ito integral $I_n(T)$ for a simple stochastic process $f_n(t)\in\mathcal{L}_2$ is defined as:

$$
I_n(T)=\int_0^Tf_n(t)dW(t)=\sum_{i=0}^{n-1}f_n(t_i)[W(t_{i+1})-W(t_i)].
$$
Notice that $I_n(T)$ is a $\mathcal{F}(T)$-measurable function $I(T):\mathcal{\Omega}\to \mathbb{R}$. 

Given a $f(t)\in\mathcal{L}_2$ that is
$$
\mathbb{E}\int_0^Tf^2(t)dt<\infty,
$$
and assuming it is a.e. continuous which means that there is a sequence of simple processes $f_n(t)\xrightarrow{a.e.}f(t)$ such that
$$
\mathbb{E}\int_0^T\left(f(t)-f_n(t)\right)^2dt=0.
$$

The Ito integral $I(T)$ for $f(t)$ is defined as
$$
I(T)=\int_0^Tf(t)dW(t)=\lim_{n\to\infty}\int_0^Tf_n(t)dW(t)=\lim_{n\to\infty}I_n(T).
$$


:::note Theorems of Ito Integral
* Linearity: For $c\in\mathbb{R}$ and another $\mathcal{F}(t)$-measurable simple function $g(t)$, $c\int_0^T f(t)dt = \int_0^T cf(t)dt$ and $\int_0^T [f(t)+g(t)]dt = \int_0^T f(t)dt+\int_0^Tg(t)dt$;
* Adapted: $I(t)$ is continuous and $\mathcal{F}(t)$-measurable;
* Martingale: $\mathbb{E}[I(t)|\tau] = I(\tau)$
* Ito isometry: $\mathbb{E} I^2(t)=\mathbb{E} \int_0^tf^2(u)du$;
* Quadratic variation: Suppose $\lbrace t_0=0,\dots, t_n=T\rbrace$ is a Riemann's partition for the interval $[0,T]$, which is a partition that $\lim_{n\to\infty}\Pi\to 0$ where
$\Vert \Pi\Vert=\max_i|t_{i+1}-t_i|$. The quadratic variation $[I,I] (t)$ is defined as
$$
[I,I](t)=\lim_{\Pi\to 0}\sum_{i}[I(t_{i+1})-I(t_i)]^2.
$$
Suppose $I(t)=\int_0^tf(u)dW(u)$, an Ito integral of process $f(t)$, then $[I,I](t)=\int_0^t f^2(u)du$.
:::

### Ito-Doeblin Formula

Let $f(x,t)$ be a function for which the partial derivatives $\partial_tf,\partial_xf$ and $\partial_x^2f$ are defined and continuous, and let $W(t)$ be a Brownian motion. Then, for every $T\ge0$:

$$
f(T,W(T))=f(0,W(0))+\int_0^T \partial_tf(t,W(t))dt+\int_0^T\partial_xf(t,W(t))dW(t)+\frac{1}{2}\int_0^T\partial_x^2f(t,W(t))dt.
$$

A differetial form of Ito-Doeblin formula is
$$
df(t,W(t))=\partial_tf(t,W(t))dt+\partial_xf(t,W(t))dW(t)+\frac{1}{2}\partial_{x}^2f(t,W(t))dt.
$$


**Proof**: Expand $f$ into Talyor's series:

$$
df=\partial_t f dt+\partial_xfdx+\frac{1}{2}\partial_x^2f(dx)^2+o(dt,dx),
$$

and substitute the $x$ by a function $W(T)$, this variation is valid as well:

$$
df(t,W)=\partial_tf(t,W)dt+\partial_xf(t,W)dW+\frac{1}{2}\partial_x^2f(t,W)(dW)^2+o[dt,(dW)^2],
$$
since $(dW)^2\to dt$ almost surely, it proves the Ito-Doeblin formula hold almost surely.

This formula can also be used to calcualte the Ito integral.

### Ito Process
A Ito process $X(t)$ is a stochastic process with the form
$$
X(t)=X(0)+\int_0^t\Delta(u)dW(u)+\int_0^t\Theta(u)du,
$$
where $X(0)$ is a scale, $W(u)$ is a Brownian process, $\Delta(u)$ and $\Theta(u)$ are adapted stochastic processes.

:::note Theorems
Suppose $X(t)$ is a Ito process, then:
1. The quadratic variation of a Ito process $X(t)$ is
$$
[X,X](t)=\int_0^t\Delta^2(u)du.
$$
2. for a function $f(t,x)$ that $\partial_tf(t,x), \partial_xf(t,x)$ and $\partial^2_xf(t,x)$ are defined and continuous. Then
$$
\begin{aligned}
f(T,X(T))=&f(0,X(0))+\int_0^Tf_t(t,X(t))dt+\int_0^Tf_x(t,X(t))dX(t)\\
&+\frac{1}{2}\int_0^Tf_{xx}(t,X(t))d[X,X](t)\\
=& f(0,X(0))+\int_0^Tf_t(t,X(t))dt+\int_0^Tf_x(t,X(t))\Delta(t)dW(t)\\
&+\int_0^Tf_x(t,X(t))\Theta(t)dt+\frac{1}{2}\int_0^Tf_{xx}(t,X(t))\Delta^2(t)dt.
\end{aligned}
$$
:::

**Proof**: Let's suppose $I(t)=\int_0^t\Delta(u)dW(u)$ and $R(t)=\int_0^t\Theta(u)du$, then
1. $(dX)^2=(dI(t)+dR(t))^2 = dI^2(t)+dR^2(t)+2dI(t)dR(t)$. $dI^2(t)=\Delta^2(t)dt$ from the property of Ito process. The rest term vanishes since they are higher order differentiate form.
2. The previous theorem implies that $d^2X(t)=\Delta(t)dt$, then the Taylor expansion of $f(t,X)$ up to the first order of $dt$ is
$$
f(t,X)=f(t_0,X_0)+f_t(t_0,X_0)dt+f_x(t_0,X_0)dX+\frac{1}{2}f(t_0,X_0)d^2X+\mathcal{O}(dt).
$$
where $dX= \Delta(t)dW+\Theta dt$, and the result follows by putting the differentiate form back to the equation.