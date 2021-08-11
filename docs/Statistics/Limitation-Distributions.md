# Limitation Distributions

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
:::

:::info Theorem:
1. **Weak law of Large Numbers** Let $\lbrace X_n\rbrace$ be a size $n$ sequence of iid (identical independent) random variables having mean and variance $\mu$ and $\sigma^2$. Let $\overline{X}_n=\sum_{i=1}^nX_i/n$, then
$$
\overline{X}_n\xrightarrow{P}\mu.
$$
2. If $X_n\xrightarrow{P}X$ and $Y_n\xrightarrow{P}Y$, then $X_n+Y_n\xrightarrow{P}X+Y$ and $X_nY_n\xrightarrow{P}XY$. If $a\in\mathbb{R}$, then $aX_n\xrightarrow{P}aX$.
:::

:::note Proof
1. From the Chebyshev inequality, we have
$$
\mathbb{P}((X-\mu)^2\ge k^2\sigma^2)\le \frac{1}{k^2\sigma^2}\int_A(X-\mu)^2d\mathbb{P}(x)\le \frac{\mathbb{E}[(X-\mu)^2]}{k^2\sigma^2}.
$$
This is equivalent to $\mathbb{E}(|X-\mu|\ge k\sigma)\le 1/k^2$. Since the variance of $\overline{X}_n$ is $sigma/\sqrt{n}$, so we have
$$
\mathbb{P}(|X_n-\mu|\ge \epsilon)=\mathbb{P}[|X_n-\mu|\ge (\epsilon\sqrt{n}/\sigma)(\sigma/\sqrt{n})]\le \frac{\sigma^2}{n\epsilon^2}\to 0.
$$
2. 
:::