# Special Distributions

## Random Variables and Distributions
---
Given a probability space $(\Omega, \mathfrak{R},\mathbb{P})$ where $\Omega$ is the event pool and $\mathfrak{R}$ is a $\sigma$-algebra defined on $\Omega$ and $\mathbb{P}$ is a probability measure defined on $\mathfrak{R}$ and $\mathbb{P}(\Omega)=1$. A random variable $X:\Omega\to \mathbb{R}$ is a $\mathfrak{R}$-measurable function which means for any Borel set $B\in\mathbb{R}$:
$$
X^{-1}(B)=\lbrace\omega|\omega\in\Omega, X(\omega)\in B\rbrace \in\mathfrak{R}.
$$

A **distriubtion measure** $\mu_X$ of $X$ is a pushforward measure induced by $X$ follows $\mu_X(B)=X_*\mathbb{P}(B)=\mathbb{P}[X^{-1}(B)]$. The Randon-Nikodym's theorem make sure a non-negative $f(x)$ exists such that
$$
\mu_X(B)=\int_Bf(x)dx,
$$
which is called the **probability density function** (PDF) and a **cumulated distribution function** (CDP) $F(x)$ is defined as $F(x)=\mathbb{P}\lbrace X\le x\rbrace$. 

A **characteristic function** $\varphi(t)$ for the random variable is defined as $\varphi(t)=\mathbb{E}(e^{iXt})$. Suppose $X=\sum_ia_iX_i$ where $X_i$ are mutually independent, then 
$$
\varphi_X(t)=\prod_i\varphi_{X_i}(a_it).
$$
Furthermore, a moment generate function is defined as $M(t)=\varphi(-it)$ and 
$$
\mathbb{E}(X^k)=\frac{d^kM(t)}{dt^k}\big|_{t=0}.
$$

## Binomial and Beta Distribution
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

### Multinomial Distributions
The Bernoulli trail can be extended to the case that 3 possible mutually exclusive outcomes from each trails. The random variable represents the outcome of $n$ sequence trials has a **trinominal distribution** $f(x,y)$:
$$
f(x,y)=\frac{n!}{x!y!(n-x-y)!}p_1^xp_2^y(1-p_1-p_2)^{n-x-y},
$$ 
where $x+y\le n$. The moment generate function is
$$
M(t_1,t_2)=(p_1e^{t_1}+p_2e^{t_2}+1-p_1-p_2)^n
$$

To generalize this idea to the case that $k$ outcomes for each trial, and let the random vector $\boldsymbol{X}=(X_1,\dots,X_k)$ be the outcomes where $X_i$ is the counts of $i$-th outcome occurance in $n$ trials. This leads to a **multinomial distribution** with a PDF
$$
f(x_1,\dots,x_k)=\frac{n!}{\prod_{i=1}^kx_i!}\prod_{i=1}^kp_i^{x_i},\quad \sum_{i=1}^np_i=1,
$$
where $p_i$ stands for the probability of the $i$-th outcome occurs in each trial. It is easy to varify that
1. $\mathbb{E}(X_i)=np_i$ and $\text{Var}(X_i)=np_i(1-p_i)$.
2. $\text{Cov}(X_iX_j)=-np_ip_j$. 

### Beta Distribution
A beta distribution $B(\alpha,\beta)$ is defined as the PDF:
$$
f(x;\alpha,\beta)= \frac{1}{B(\alpha,\beta)}x^{\alpha-1}(1-x)^{\beta-1}=\frac{\Gamma(\alpha+\beta)}{\Gamma(\alpha)\Gamma(\beta)}x^{\alpha-1}(1-x)^{\beta-1}.
$$
This is generalized from substituting the factorial in the binomial distribution by the gamma function. The **cumulative function** is given by
$$
F(x;,\alpha,\beta)=\frac{B(x;\alpha,\beta)}{B(\alpha,\beta)},\quad B(x;\alpha,\beta) = \int_0^xt^{\alpha-1}(1-t)^{\beta-1}dt.
$$
The properties of the beta distribution:
* **mean** $\mu=\mathbb{E}(X)=\frac{\alpha}{\alpha+\beta}$.
* **variance** $\text{var}(X)= \frac{\alpha\beta}{(\alpha+\beta)^2(\alpha+\beta+1)}$.

Similar method can be used to extend the multinomial distribution as well. So that we have a multi-variate beta distribution:
$$
f(x_1,\dots, x_n;\alpha_1,\dots, \alpha_n)=\frac{1}{B(\alpha_1,\dots,\alpha_n)}\prod_{i=1}^np_i^{x_i}, \quad B(\alpha_1,\dots,\alpha_n)=\frac{\prod_{i=1}^n\Gamma(\alpha_i)}{\Gamma(\sum_{i=1}^n{\alpha_i})}.
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


## Joint Distributions
---

For two random variables $X,Y$, the direct product forms a map $X\otimes Y:\Omega\to \mathbb{R}\otimes\mathbb{R}$, and the corresponding measure is defined as
$$
\mu_{X\otimes Y}(B_1\otimes B_2)=\mathbb{P}[X^{-1}(B_1)\cap Y^{-1}(B_2)],
$$
where $B_1,B_2$ are arbitrary Borel sets in $\mathbb{R}$. The PDF or CDF incuded from this measure is called **joint probability density function** or **joint cumulative distribution function**. Usually, it is easier to construct the joint CDF, denoted as $F_{X,Y}(x,y)$ and the joint PDF denoting as $f_{X,Y}(x,y)$ is obtained by the partial differential. Furthermore, if $X\perp Y$ (independent), then
$$
\mathbb{P}\left\lbrace X^{-1}\big([-\infty,x]\big)\cap Y^{-1}\big([-\infty,y]\big)\right\rbrace =\mathbb{P}\left\lbrace X^{-1}\big([-\infty,x]\big)\right\rbrace \mathbb{P}\left\lbrace Y^{-1}\big([-\infty,y]\big)\right\rbrace,
$$ 
so that $F_{X,Y}(x,y)=F_X(x)F_Y(y)$ and so is the joint PDF. 


### t-Distributions

It will show later that the sample mean and sample variance is independent with each other. Provided the distribution of sample mean and variance is known by CLT, the distribution of normalized sample mean satisifes a $t$**-distribution**. The PDF of the $t$-distribution with degree of freedom $r$, denoted as $t(r)$, is 
$$
f(t)= \frac{\Gamma(\frac{r+1}{2})}{\Gamma(\frac{r}{2})}\frac{1}{\sqrt{\pi r}}(1+t^2/r)^{-(r+1)/2}, \quad t \in \mathbb{R}.
$$

This distribution obtained from the joint PDF of $X\sim N(0,1)$ and $Y\sim \chi^2(r)$ and $X\perp Y$:
$$
f_{X,Y}(x,y)=\frac{e^{-x^2/2}}{\sqrt{2\pi}}\frac{y^{r/2}e^{-y/2}}{\Gamma(r/2)2^{r/2}},\quad x\in\mathbb{R}, y\in[0,+\infty].
$$

The $t$-distribution follows from a substitution that $t=x/\sqrt{y/r}$ and integral over the $y$.

:::info Student's Theorem:
Let $X_1,\dots, X_n$ be i.i.d. random variables each having $X_i\sim N(\mu,\sigma^2)$. The new random variables $\overline{X}$ and $S^2$ defined are defined as
$$
\overline{X}=\sum_{i=1}^nX_i/n,\quad S^2=\sum_{i=1}^n(X_i-\overline{X})^2/(n-1)
$$
satisfies:
1. $\overline{X}\sim N(\mu, \sigma^2/n)$;
2. Independent: $\overline{X}\perp S^2$;
3. $(n-1)S^2/\sigma^2\sim \chi^2(n-1)$;
4. Let $Z=(\overline{X}-\mu)/\sqrt{\frac{\sigma^2}{n}}$, then $T=Z/\sqrt{\frac{S^2}{n-1}}=(\overline{X}-\mu)/(S/\sqrt{n})$ satisfies the $t$-distribution $T\sim t(n-1)$.
:::

**Proof**: 
1. The proof of the 1st refers to the [section](#normal-distributions). 
2. Since all $X_i$ and $\overline{X}$ are normal distributions, then $X_i-\overline{X}\perp\overline{X}$ if and only if $\text{Cov}(X_i-\overline{X},\overline{X})=0$. Straight calcuation shows
$$
\begin{aligned}
\text{Cov}(X_i-\overline{X},\overline{X})&=\mathbb{E}(X_i\overline{X}-X_i\mu-\overline{X}^2+\mu\overline{X})\\
&=\mathbb{E}(X_i\overline{X}-\overline{X}^2)\\
&=\frac{\sigma^2+\mu^2}{n}+\frac{n-1}{n}\mu^2-\frac{\sigma^2}{n}-\mu^2\\
&=0.
\end{aligned}
$$
This proved the 2nd point.
3. $X_i-\mu\sim N(0,\sigma^2)$ and
$$
\begin{aligned}
V&=\sum_i\left[(X_i-\overline{X})+(\overline{X}-\mu)\right]^2/\sigma^2\\
&=\sum_i(X_i-\overline{X})^2/\sigma^2+\left(\frac{\overline{X}-\mu}{\sigma\sqrt{n}}\right)^2,
\end{aligned}
$$ 
where the first term is $(n-1)S^2/\sigma^2$ and independent with the second term. It is known that $V\sim\chi^2(n)$ and the last term is a $\chi^2(1)$ distribution. It implies that the rest is $\chi^2(n-1)$ which conclude this proof.
4. It follows by the $t$-distributions definition along with the 2nd point that $\overline{X}\perp S^2$.

The CLT along with the student's theorem play a important role in sampling inference.

### F-Distributions

Assuming $X\sim\chi^2(r_1)$, $Y\sim\chi^2(r_2)$, $X\perp Y$ and $F=\frac{X/r_1}{Y/r_2}$, using the similar idea we derive the $t$-distribution, the $F$ follows a $F$-distributions with PDF:
$$
f(x)=\frac{\Gamma(\frac{r_1+r_2}{2})}{\Gamma(\frac{r_1}{2})\Gamma(\frac{r_2}{2})}\frac{(r_1/r_2)^{\frac{r_1}{2}}x^{\frac{r_1}{2}-1}}{(1+\frac{r_1}{r_2}x)^{(r_1+r_2)/2}}, \quad x\in [0,+\infty).
$$