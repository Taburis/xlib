# Gamma Distributions


A gamma function $\Gamma(\alpha)$ defined as
$$
\Gamma(\alpha)=\int_0^\infty x^{\alpha-1}e^{-x}dx
$$
exists for $\alpha>0$. The Gamma distribution $\Gamma(\alpha,\beta)$ is defined with the PDF
$$
f(x)= \frac{1}{\Gamma(\alpha)\beta^\alpha}x^{\alpha-1}e^{-x/\beta},\quad 0\le x<\infty,
$$
and $f(x)=0$ for elsewhere. The moment generate function is
$$
M(t)=\frac{1}{(1-\beta t)^\alpha},\quad t< 1/\beta,
$$
and $\mathbb{E}(X)=\alpha\beta$ and $\text{Var}(X)=\alpha\beta^2$. 

## $\chi^2$-Distributions
A **chi square distribution** $\chi^2(r)$ is defined as $\chi^2(r)=\Gamma(r/2,2)$ where $r$ is called the **degree of freedom**. 

:::info Properties:
1. Suppose $X_i\sim \Gamma(\alpha_i,\beta)$ are independent variables and $Y=\sum_iX_i$, then $Y\sim\Gamma(\sum_i\alpha_i,\beta)$. 
2. Suppose $X\sim\chi^2(r)$, then $\mathbb{E}(X)= r$, $\text{Var}(X)=2r$, and $\mathbb{E}(X^k)=2^k\Gamma(r/2+k)/\Gamma(r/2)$ if $k>-r/2$..
3. Let $X_i\sim N(0,1)$ are mutually independent variables, then $X^2\sim \chi^2(1)$. Let $Y=\sum_{i=1}^rX_i^2$ then $Y\sim \chi^2(r)$.
:::

**Proof**
1. Since $X_i$ are mutually independent, then $M_Y(t)=\prod_iM_{X_i}(t)$ which leads to the theorem immdiately. It means that $\sum_i\chi^2(r_i)=\chi^2(\sum_ir_i)$.
2. It follows straight from the calculation of $\mathbb{E}(X^k)$.
3. Let $X\sim N(0,1)$, consider the PDF 
$$
\begin{aligned}
f(x)&=\frac{1}{\sqrt{2\pi}}\int_{-\infty}^\infty x^2e^{-\frac{x^2}{2}}dx\\
&=\frac{1}{\sqrt{2\pi}}\int_0^\infty\sqrt{y}e^{-y/2}dy,
\end{aligned}
$$
which leads to a PDF for $Y=X^2$ as $e^{-y/2}y^{-1/2}/\sqrt{2\pi}$ which is a $\Gamma(1/2,2)=\chi^2(1)$. Furthermore, $Y=\sum_{i=1}^rX^2_i$ is a $\chi^2(r)$ follows from the 1st theorem. 

## Exponential Distributions
---
Considering a Poisson process that $x$ be the numbers generated for a interval with length $w$. We know the $X$ satisfies a Poisson distribution $P(\lambda w)$ for a fixed $w$. Now suppose the random variable $W$ stands for the length of the intevral associated with a certain non-negative interger $x$, then the distribution of $W$ is known as **exponential distribution** with the PDF defined as
$$
f(w)=\frac{\lambda}{\Gamma(x)}(\lambda\omega)^{x-1}e^{-\lambda \omega},
$$
and $f(w)=0$ elsewhere.

**Proof** 
Without loose any generality, we consider the Poisson process generating a non-negative integer $X$ for a fixed interval with length $w$. The CDF to have $X=k$ for $W\le w$ is
$$
\mathbb{P}(X=k, W\le w)=1-\mathbb{P}(X=k,W>w), \quad \mathbb{P}(X=k,W>w)=\sum_{x=0}^{k-1}\mathbb{P}(X=x,W=w),
$$
where the probability of $\mathbb{P}(X=j,W=w)$ satisfies the Poisson distribution, then
$$
\sum_{x=0}^{k-1}\frac{(\lambda w)^{x-1}e^{-\lambda w}}{x!}=\int_{\lambda w}^\infty\frac{z^{k-1}e^{-z}}{(k-1)!}dz,
$$
obtained by integral the right side by part. This leads to
$$
\mathbb{P}(X=k, W\le w) = \int_0^{\lambda w}\frac{z^{k-1}e^{-z}}{(k-1)!}
$$
so that the PDF is $f(w)=d\mathbb{P}(X=k, W\le w)/dw$ which is teh exponential distributions for $w>0$.
