# Special Distributions

Given a probability space $(\Omega, \mathfrak{R},\mathbb{P})$ where $\Omega$ is the event pool and $\mathfrak{R}$ is a $\sigma$-algebra defined on $\Omega$ and $\mathbb{P}$ is a probability measure defined on $\mathfrak{R}$ and $\mathbb{P}(\Omega)=1$. A random variable $X:\Omega\to \mathbb{R}$ is a $\mathfrak{R}$-measurable function which means for any Borel set $B\in\mathbb{R}$:
$$
X^{-1}(B)=\lbrace\omega|\omega\in\Omega, X(\omega)\in B\rbrace \in\mathfrak{R}.
$$

A **distriubtion measure** $\mu_X$ of $X$ is a pushforward measure induced by $X$ follows $\mu_X(B)=X_*\mathbb{P}(B)=\mathbb{P}[X^{-1}(B)]$. The Randon-Nikodym's theorem make sure a non-negative $f(x)$ exists such that
$$
\mu_X(B)=\int_Bf(x)dx,
$$
which is called the **probability density function** (PDF) and a **cumulated distribution function** (CDP) $F(x)$ is defined as $F(x)=\mathbb{P}\lbrace X\le x\rbrace$.  

## Binomial Distribution
---
### Bernoulli Experiment
A **Bernoulli experiment** is a class of statistical experiment that only 2 possible mutually exclusive outcome at a fixed probability for each trail. The random variable $X$ for this experiment is formulated as a Bernoulli distribution $p(x)$ that
$$
p(x)=p^x(1-p)^{1-p},\quad x=0,1.
$$
The $p$ is the probability to get the outcome represented by $1$. For $X$, it is easy to verify that $\mathbb{E}(X)=p$, $\text{Var}(X)=p(1-p)$. 

A random variable represents the outcome of a sequence of $n$ Bernoulli trials forms the **Binomial distribution** $b(n,p)$ where $p$ is the probability to get 1 for each trail. The PMF $f(x)$ is
$$
f(x)=\begin{pmatrix}n \\ x\end{pmatrix}p^{x}(1-p)^{n-x}, \quad \begin{pmatrix}n,x\end{pmatrix}=\frac{n!}{x!(n-x)!}.
$$
where $x=0,1,\dots, n$. 
:::info Properties:
* The moment generate function for binomial distribution is $M(t)=[(1-p)+pe^t]^n$.
* $\mathbb{E}(X)=np$ and $\text{Var}(X)=np(1-p)$. 
* Suppose $X_i\sim b(n_i,p)$ and $Y=\sum_iX_i$, then $Y\sim b(\sum_in_i,p)$. 
:::

### Other Related Distributions
The Bernoulli trail can be extended to the case that 3 possible mutually exclusive outcomes from each trails. The random variable represents the outcome of $n$ sequence trials has a **trinominal distribution** $f(x,y)$:
$$
f(x,y)=\frac{n!}{x!y!(n-x-y)!}p_1^xp_2^y(1-p_1-p_2)^{n-x-y},
$$ 
where $x+y\le n$. The moment generate function is
$$
M(t_1,t_2)=(p_1e^{t_1}+p_2e^{t_2}+1-p_1-p_2)^n
$$

## Poisson Distributions
---
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


## Normal Distributions
---
A normal distribution play a important role in statistics due to central limit theorem. A normal distribution $N(\mu,\sigma)$ is determined by mean $\mu$ and variance and $\sigma$. Suppose a random variable obeys the normal distribution, the PDF is
$$
f(x)=\frac{1}{\sqrt{2\pi\sigma^2}}\exp\left(-\frac{(x-\mu)^2}{2\sigma^2}\right).
$$ 
A normal distribution $N(0,1)$ is called a **standard normal distribution**.

:::info Properties: (suppose $X\sim N(\mu_1,\sigma^2)$)
1. Given $Y\sim N(\mu_2,\sigma_2)$ and $X\perp Y$, then $X+Y\sim N(\mu_1+\mu_2,\sigma_1^2+\sigma^2)$
2. Suppose $c\in\mathbb{R}$, then $cX\sim N(c\mu,c^2\sigma^2)$. 
3. If $Z=(X-\mu)/\sigma$, then $Z\sim N(0,1)$ and $Z^2\sim\chi^2(1)$, where $\chi(1)$ is a chi square distribution with degree of freedom 1.
4. The moment generate function of $X$ is $M_X(t)=\exp(\mu t+\sigma^2 t^2/2)$. 
:::

**Multivariate Normal Distribution**: A random vector $\boldsymbol{X}=(X_1,\dots, X_m)$ forms by $m$ normal distribution $X_i\sim N_m(\mu_i,\sigma_i)$, then the random vector forms $n$-dimensional normal distribution $\boldsymbol{X}\sim N_m(\boldsymbol{\mu}, \Sigma)$ where $\Sigma$ is a covariance matrix which is positive and symmetric. 

:::info Properties:
* The linear transform $\Gamma$ on $\boldsymbol{X}$ will leads to $Y=\Gamma\boldsymbol{X} = N_m(\Gamma\boldsymbol{\mu},\Gamma^T\Sigma\Gamma)$. 
* The moment generate function $M_{\boldsymbol{X}}(\boldsymbol{t})=\exp\lbrace \boldsymbol{t}^T\boldsymbol{\mu}+\boldsymbol{t}^T\Sigma\boldsymbol{t}/2\rbrace$. 
* Suppose 
$$
\boldsymbol{X}=
\begin{pmatrix}
\boldsymbol{X_1}\\
\boldsymbol{X_2}
\end{pmatrix}
,\quad \Sigma= 
\begin{pmatrix}
\Sigma_{11} & \Sigma_{12}\\
\Sigma_{12}^T & \Sigma_{22}
\end{pmatrix},
$$
then $\boldsymbol{X}_1\perp\boldsymbol{X}_2$ if and only if $\Sigma_{12}=0$. And the conditional variable 
$$
\boldsymbol{X}_1|\boldsymbol{X}_2\sim N_m(\boldsymbol{\mu}_1+\Sigma_{12}\Sigma_{22}^{-1}(\boldsymbol{X}_2-\boldsymbol{\mu}_2), \Sigma_{11}-\Sigma_{12}\Sigma_{22}\Sigma_{12}^T)
$$
* Let $W=(\boldsymbol{X}-\boldsymbol{\mu})^T\Sigma^{-1}(\boldsymbol{X}-\boldsymbol{\mu})$, then $W\sim\chi^2(m)$.
:::

Suppose orthogonal transform $\Gamma$ diagonalize $\Sigma$, and $\Gamma\Sigma\Gamma^T=\Lambda$ where the diagonal elements are $\lambda_i$. Suppose $\lambda_1\ge\lambda_2\ge\dots$. The variable $\boldsymbol{Y}=\Gamma(\boldsymbol{X}-\boldsymbol{\mu})$ is called **principal components** and $\text{tr}(\Sigma)$ is called **total variation**. The first principal component $Y_1$ has the largest variation over all linear combination $\boldsymbol{a}^T(\boldsymbol{X}-\boldsymbol{\mu})$, where $\boldsymbol{a}^T\cdot\boldsymbol{a} = 1$. Because the variation is a linear combination of all principal components with the constraints that $\sum_ia_i^2=1$ and $\lambda_1$ is the largest component.

## $\Gamma$, $\beta$ and $\chi^2$ Distributions
---

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

### $\chi^2$-Distributions
A **$chi^2$ distribution**$ $\chi^2(r)$ is defined as $\chi^2(r)=\Gamma(r/2,2)$ where $r$ is called the **degree of freedom**.

:::info Properties:
1. Suppose $X_i\sim \Gamma(\alpha_i,\beta)$ are independent variables and $Y=\sum_iX_i$, then $Y\sim\Gamma(\sum_i\alpha_i,\beta)$. 
2. Suppose $X\sim\chi^2(r)$, then $\mathbb{E}(X^k)=2^k\Gamma(r/2+k)/\Gamma(r/2)$ if $k>-r/2$.
3. Let $X_i\sim N(0,1)$ are mutually independent variables, then $X^2\sim \chi^2(1)$. Let $Y=\sum_{i=1}^rX_i^2$ then $Y\sim \chi^2(r)$.
:::

**Proof**
1. Since $X_i$ are mutually independent, then $M_Y(t)=\prod_iM_{X_i}(t)$ which leads to the theorem immdiately.
2. It follows straight from the calculation of $\mathbb{E}(X^k)$.
3. Let $X\sim N(0,1)$, consider the PDF 
$$
\begin{aligned}
f(x)&=\frac{1}{\sqrt{2\pi}}\int_{-\infty}^\infty x^2e^{-\frac{x^2}{2}}dx\\
&=\frac{1}{\sqrt{2\pi}}\int_0^\infty\sqrt{y}e^{-y/2}dy,
\end{aligned}
$$
which leads to a PDF for $Y=X^2$ as $e^{-y/2}y^{-1/2}/\sqrt{2\pi}$ which is a $\Gamma(1/2,2)=\chi^2(1)$. Furthermore, $Y=\sum_{i=1}^rX^2_i$ is a $\chi^2(r)$ follows from the 1st theorem. 

### Exponential Distributions

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
