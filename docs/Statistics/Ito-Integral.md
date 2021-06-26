# Ito Integral

## Definition

Starting from the **Ito Integral definition for simple functions**:
Given a probability Brownian motion $W(t)$ with a filter $\mathcal{F}(t)$, and $f(t)$ which is a $\mathcal{F}(t)$-measuable simple
 function. The Ito integral $I(T)$ for the simple function $f(t)$, the values are discrete in the interval $(t_0=0,\dots,t_n=T)$, is defined as

$$
I(T)=\int_0^Tf(t)dW(t)=\sum_{i=0}^{m-1}f(t_i)[W(t_{i+1})-W(t_i)].
$$

**Properties** of Ito integral:

* Linearity: For $c\in\mathbb{R}$ and another $\mathcal{F}(t)$-measurable simple function $g(t)$, $c\int_0^T f(t)dt = \int_0^T cf(t)dt$ and $\int_0^T [f(t)+g(t)]dt = \int_0^T f(t)dt+\int_0^Tg(t)dt$;
* Adapted: $I(t)$ is continuous and $\mathcal{F}(t)$-measurable;
* Martingale: $\mathbb{E}[I(t)|\tau] = I(\tau)$
* Ito isometry: $\mathbb{E} I^2(t)=\mathbb{E} \int_0^tf^2(u)du$;
* Quadratic variation: $[I,I] (t)=\int_0^t f^2(u)du$.

## Ito-Doeblin Formula

Let $f(x,t)$ be a function for which the partial derivatives $f_t,f_x$ and $f_{xx}$ are defined and continuous, and let $W(t)$ be a Brownian motion. Then, for every $T\ge0$:

$$
f(T,W(T))=f(0,W(0))+\int_0^T f_t(t,W(t))dt+\int_0^Tf_x(t,W(t))dW(t)+\frac{1}{2}\int_0^Tf_{xx}(t,W(t))dt.
$$

**Proof**: Expand $f$ into Talyor's series:

$$
df=\partial_t f dt+\partial_xfdx+\frac{1}{2}\partial_{xx}f(dx)^2+o(dt,dx),
$$

and substitute the $x$ by a function $W(T)$, this variation is valid as well:

$$
df(t,W)=f_t(t,W)dt+f_x(t,W)dW+\frac{1}{2}f_{xx}(t,W)(dW)^2+o[dt,(dW)^2],
$$
since $(dW)^2\to dt$ almost surely, it proves the Ito-Doeblin formula hold almost surely.

This formula can also be used to calcualte the Ito integral.