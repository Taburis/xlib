# Asymptotic Distributions


## Convergence of Ranodm Variable
---
Given a sequence of random variables $X$ and $X_1,\dots, X_n$, there are several convergence definition:
1. **Convegence in law**, or convergence in (cumulative) distribution, or convergence weakly denoting as $X_n\xrightarrow{d}X$ if
$$
\lim_{n\to\infty}F_n(x)=F(x),
$$
where $F_n(x)$ and $F(x)$ are cumulative distribution of $X_n$ and $X$, respectively. 
2. **Convergence in probability** denoting as $X_n\xrightarrow{P} X$ if
$$
\lim_{n\to\infty} \mathbb{P}\left\lbrace \omega\in \Omega:|X_n(\omega)-X(\omega)|>0\right \rbrace = 0,
$$ 
3. **Convergence almost surely** denoting as $X_n\xrightarrow{a.s.} X$ if
$$
\mathbb{P}\left\lbrace \omega\in \Omega: \lim_{n\to\infty}X_n(\omega)=X(\omega)\right\rbrace = 1.
$$

:::tip Relations:
* Convergence in CDF is equivalent to the convergence of PDF almost surely. 
* $X_n\xrightarrow{P}X$ $\implies$ $X_n\xrightarrow{d}X$. 
* $X_n\xrightarrow{a.s.}X$ $\implies$ $X_n\xrightarrow{P}X$. 
* $X_n\xrightarrow{a.s.}X$ and $f_n(x)\xrightarrow{a.s.}f(x)$ is different where the latter means $X_n\xrightarrow{d}X$.
:::

:::note Theorems:
1. If $X_n\xrightarrow{P}X$ and $Y_n\xrightarrow{P}Y$, then $X_n+Y_n\xrightarrow{P}X+Y$ and $X_nY_n\xrightarrow{P}XY$. If $a\in\mathbb{R}$, then $aX_n\xrightarrow{P}aX$.
2. Let $M_{X_n}(t)$ be the moment generate function for $X_n$ and exists $[-h,h]$ in $\mathbb{R}$ such that all $M_{X_n}(t)$ is finite for $|t|\le h$. Then $\lim_{n\to\infty}M_{X_n}(t)=M(t)$, for some subinterval in $[-h,h]$, is equivalent to $X_n\xrightarrow{d}X$.
3. If 
$$
\lim_{n\to\infty}\mathbb{E}(X_n-X)^2\to 0,
$$
then $X_n\xrightarrow{P}X$.  
4. If $X_n\xrightarrow{P}c$ and $f$ is a function continuous at $c$, then $f(X_n)\xrightarrow{P}f(c)$.
5. $X_n\xrightarrow{d}X$ if and only if $\mathbb{E}f(X_n)\to \mathbb{E}f(X)$ for all bounded continuous function $f$. As a corollary, if $f$ is a continuous function and $X_n\xrightarrow{d} X$, then $f(X_n)\xrightarrow{d}f(X)$. 
6. Suppose $X_n\xrightarrow{X}$, $A_n\xrightarrow{P}a$, and $B_n\xrightarrow{P}b$, then $A_n+B_nX_n\xrightarrow{d}a+bX$.
:::

**Proof**
1. The proof follows from the inequality
$$
|X_n-X|+|Y_n-Y|\ge |(X_n+Y_n)-(X+Y)|.
$$
For the multiplication case, noticing
$$
X_nY_n=X^2_n/2+Y_n^2/2-(X_n-Y_n)^2/2\xrightarrow{P}X^2/2+Y^2/2-(X-Y)^2/2=XY.
2. To prove the necessarity, let $f_n(x)$ be the PDF of $X_n$ and the convergence in distribution impliese that $f_n(x)\xrightarrow{a.s.}f(x)$. Since the $e^{itx}$ is bounded, and $\mathbb{E}(e^{itX})\le \mathbb{E}(X)$, the converge of the characteristic function follows from the Lebesgue's dominated convergence theorem.
The sufficiency follows from the inverse Fourier transform existence. Since the moment generate function $M_X(t)=\varphi(it)$, the proof concluded.

### Law of Large numbers
:::info Law of Large Numbers:
Let $\lbrace X_n\rbrace$ be a size $n$ sequence of iid (identical independent) random variables having mean $\mu$. Let $\overline{X}_n=\sum_{i=1}^nX_i/n$, then
$$
\text{Weak Law:}\quad\overline{X}_n\xrightarrow{P}\mu
$$
$$
\text{Strong Law:}\quad\overline{X}_n\xrightarrow{a.s.}\mu.
$$
:::

**Proof**:
From the Chebyshev inequality, we have
$$
\mathbb{P}((X-\mu)^2\ge k^2\sigma^2)\le \frac{1}{k^2\sigma^2}\int_A(X-\mu)^2d\mathbb{P}(x)\le \frac{\mathbb{E}[(X-\mu)^2]}{k^2\sigma^2}.
$$
This is equivalent to $\mathbb{E}(|X-\mu|\ge k\sigma)\le 1/k^2$. Since the variance of $\overline{X}_n$ is $\sigma/\sqrt{n}$, so we have
$$
\mathbb{P}(|X_n-\mu|\ge \epsilon)=\mathbb{P}[|X_n-\mu|\ge (\epsilon\sqrt{n}/\sigma)(\sigma/\sqrt{n})]\le \frac{\sigma^2}{n\epsilon^2}\to 0.
$$

:::note Remark
Although the proof used a finite variance $\sigma<\infty$, but this assumption is **not necessary**. This just makes the proof shorter and easy to understand. The large or infinite variance will leads to a slower convergence. 
:::

### Central Limit Theorem
Suppose $X_n\xrightarrow{d} X$, then the distribution (CDF or PDF) of $X$ is the **asymptotic distribution** of the sequence $\lbrace X_n\rbrace$. If we strengthen the law of large numbers by requiring the variance is finite $\sigma^2<\infty$, we get the central limit theorem which shows that the sample mean $\overline{X}_n$ has a normal $N(\mu,\sigma^2)$ as its asymptotic distribution. Central Limit Theorem (CLT) has a vary versions. Here is the classical one.
:::info Central Limit Theorem 
Let $\lbrace X_n\rbrace$ be a sequence of $n$ i.i.d. random variables having the same mean $\mu$ and variance $\sigma<\infty$. Let $\overline{X}_n=\sum_{i=1}^nX_i/n$, then
$$
\sqrt{n}(\overline{X}_n-\mu)\xrightarrow{d} N(0,\sigma^2).
$$
:::

**Poof**: Let $X=\sqrt{n}(\overline{X}_n-\mu)=[\sum_i(X_i-\mu)]/\sqrt{n}=\sum_iY_i/\sqrt{n}$ where $Y_i=X_i-\mu$, and the characteristic function is defined as 
$$
\varphi_X(t)=\prod_i\varphi_{Y_i}\left(\frac{t}{\sqrt{n}}\right).
$$
As $t$ is finite such that $t/\sqrt{n}\to 0$ as $n\to \infty$, the characteristic function can be approximated by Taylor's formula:
$$
\begin{aligned}
\varphi_{Y_i}(t)&=\mathbb{E}\left[1-\frac{itY_i}{\sqrt{n}}-\frac{t^2Y^2_i}{2n}+o(Y^2_it^2/n)\right]\\
&=1-\frac{(t\sigma)^2}{2n}+o\left(\frac{t^2}{n}\right),
\end{aligned}
$$
where the last step used the $\mathbb{E}(Y_i)=0$ and $\mathbb{E}(Y_i^2)=\sigma^2$ for all $i$.
All the $\varphi_{Y_i}(t)$ has the same form for any $Y_i$ when $t$ is small and this is the key of the proof: The asymptotic behavior of a random variable is dominated only by the mean and variance. The final characteristic function $\varphi_X(t)$ has the form
$$
\varphi_X(t)=\left(1-\frac{(t\sigma)^2}{2n}\right)^{n}\to e^{-t^2\sigma^2/2}|_{n\to\infty}
$$
where the last step comes from $\lim_{n\to\infty}(1+x/n)^{n}=e^x$. This is exactly the characteristic function of normal distribution $N(0,\sigma^2)$.


:::info Delta Method:
With the assumption of the central limit theorem, if $f$ is differentiable at $\mu$, then
$$
\sqrt{n}[f(\overline{X}_n)-f(\mu)]\xrightarrow{d} N(0,[f'(\mu)]^2\sigma^2).
$$
:::


The asymptotic distributiuon of the percentiles is an example of the application about the delta method:
:::note Theorem:
Let $X_1,\dots,X_n$ be iid random variables with CDF $F$. Suppose $\xi_\gamma$ be the $\gamma$-percentile where $\gamma \in [0,]$, and $\hat \xi_\gamma$ is the $\lfloor\gamma n\rfloor$-th order statistics for $X_1,\dots, X_n$. (This means that $F(\xi_\gamma)=\gamma$ and $\hat \xi_\gamma$ is the unbiased estimator for $\xi_\gamma$.) If $F'(\xi_\gamma)$ exists and is finite and positive, then 
$$
\sqrt{n}(\hat \xi_\gamma-\xi_\gamma)\xrightarrow{d} N\left(0, \frac{\gamma(\gamma-1)}{[F'(\xi_\gamma)]^2}\right).
$$
:::



## Uniform Convergence of General Functions
Given a set of iid random variables $\lbrace X\rbrace_n$ with probability $\mathbb{P}$ and a class $\mathcal{F}$ of integrable real-valued functions $f$ defined over the domain of the probability space of $X$. Then we define a random variable as
$$
\Vert \mathbb{P}_n-\mathbb{P}\Vert_{\mathcal{F}}=\sup_{f\in\mathcal{F}}\left|\frac{1}{n}\sum_{i=1}^nf(X_i)-\mathbb{E}[f(X)]\right|,
$$
which measure the upper bound of the deviation between the sample average from the expectation for the class $\mathcal{F}$. And the class $\mathcal{F}$ is called the **Glivenko-Cantelli** class of $\mathbb{P}$ if
$$
\Vert \mathbb{P}_n-\mathbb{P}\Vert_{\mathcal{F}}\xrightarrow{d} 0,
$$
as the $n\to\infty$. 