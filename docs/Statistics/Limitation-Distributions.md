# Limitation Distributions


## Convergence of Ranodm Variable
---
Given a sequence of random variables $X$ and $X_1,\dots, X_n$, there are several convergence definition:
1. **Convegence in law** or convergence in cumulative distribution (or convergence weakly) denoting as $X_n\xrightarrow{d}X$ if
$$
\lim_{n\to\infty}F_n(x)=F(x),
$$
where $F_n(x)$ and $F(x)$ are cumulative distribution of $X_n$ and $X$, respectively. 
2. **Convergence in probability** denoting as $X_n\xrightarrow{P} X$ if
$$
\lim_{n\to\infty} \mathbb{P}(|X_n-X|>0) = 0,
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

:::info Theorem:
1. **Weak law of Large Numbers** Let $\lbrace X_n\rbrace$ be a size $n$ sequence of iid (identical independent) random variables having mean and variance $\mu$ and $\sigma^2$. Let $\overline{X}_n=\sum_{i=1}^nX_i/n$, then
$$
\overline{X}_n\xrightarrow{P}\mu.
$$
2. If $X_n\xrightarrow{P}X$ and $Y_n\xrightarrow{P}Y$, then $X_n+Y_n\xrightarrow{P}X+Y$ and $X_nY_n\xrightarrow{P}XY$. If $a\in\mathbb{R}$, then $aX_n\xrightarrow{P}aX$.
3. Let $M_{X_n}(t)$ be the moment generate function for $X_n$ and exists $[-h,h]$ in $\mathbb{R}$ such that all $M_{X_n}(t)$ is finite for $|t|\le h$. Then $\lim_{n\to\infty}M_{X_n}(t)=M(t)$, for some subinterval in $[-h,h]$, is equivalent to $X_n\xrightarrow{d}X$.
:::

Proof
1. From the Chebyshev inequality, we have
$$
\mathbb{P}((X-\mu)^2\ge k^2\sigma^2)\le \frac{1}{k^2\sigma^2}\int_A(X-\mu)^2d\mathbb{P}(x)\le \frac{\mathbb{E}[(X-\mu)^2]}{k^2\sigma^2}.
$$
This is equivalent to $\mathbb{E}(|X-\mu|\ge k\sigma)\le 1/k^2$. Since the variance of $\overline{X}_n$ is $sigma/\sqrt{n}$, so we have
$$
\mathbb{P}(|X_n-\mu|\ge \epsilon)=\mathbb{P}[|X_n-\mu|\ge (\epsilon\sqrt{n}/\sigma)(\sigma/\sqrt{n})]\le \frac{\sigma^2}{n\epsilon^2}\to 0.
$$
2. The proof follows from the inequality
$$
|X_n-X|+|Y_n-Y|\ge |(X_n+Y_n)-(X+Y)|.
$$
For the multiplication case, noticing
$$
X_nY_n=X^2_n/2+Y_n^2/2-(X_n-Y_n)^2/2\xrightarrow{P}X^2/2+Y^2/2-(X-Y)^2/2=XY.
3. To prove the necessarity, let $f_n(x)$ be the PDF of $X_n$ and the convergence in distribution impliese that $f_n(x)\xrightarrow{a.s.}f(x)$. Since the $e^{itx}$ is bounded, and $\mathbb{E}(e^{itX})\le \mathbb{E}(X)$, the converge of the characteristic function follows from the Lebesgue's dominated convergence theorem.
The sufficiency follows from the inverse Fourier transform existence. Since the moment generate function $M_X(t)=\varphi(it)$, the proof concluded.

Suppose $X_n\xrightarrow{d} X$, then the distribution (CDF or PDF) of $X$ is the **asymptotic distribution** of the sequence $\lbrace X_n\rbrace$. The central limit theorem shows that the sample mean $\overline{X}_n$ has a normal $N(\mu,\sigma^2)$ as the asymptotic distribution.

## Central Limit Theorem
---
Central Limit Theorem (CLT) has a vary versions. Here is the classical one.
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

