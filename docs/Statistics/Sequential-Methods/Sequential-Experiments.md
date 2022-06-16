# Sequential Experiments

## Sequential Estimation
---

Let $X_i, i=1,\dots, N$ be iid randomm variables from a one-parameter exponential family with the PDF
$$
f(x,\theta)= h(x)\exp\left\lbrace \eta(\theta)T(x)+A(\theta)\right\rbrace.
$$
Noticing that $N$ here is not a fixed number and $T_i=T(x_i)$ and $\overline{T}_n=\sum_{i=1}^nX_i/n$. Let $e$ be an estimator for $g(\theta)$, and $e_n$ stands for the estimation by using the first $n$ observed data. Then it has to be a function of the $\overline{T}_n$ since $\overline{T}_n$ is sufficient statistics of a regular exponential family, $\hat \lambda_n=\lambda(\overline{T}_n)$. Furthermore, $\overline{T}_n\xrightarrow{d} N(\mu_T, \sigma_T^2)$, by the delta method, we have
$$
\sqrt{n}(\hat \lambda_n-\lambda)\xrightarrow{d} N\left(0, \nu^2(\theta)\right),\quad \nu(\theta)=\lambda'(\mu_T)\sigma_T,
$$
where $\mu_T$ and $\sigma_T$ are functions of $\theta$. This gives an asymptotic $1-\alpha$ confidence interval for $\hat\lambda$:
$$
\left(\hat\lambda_n\pm \frac{z_{\alpha/2}}{\nu(\hat\theta)\sqrt{n}}\right).
$$

### Fixed Width Confidence Intevals

Although the derivation above is carried with a fixed $n$, however, it tell us that the confidence level can be increased conitnuously with continuing sampling. Suppose we would like to shrink the confidence interval to smaller than the precision $w$ given, the precision can always be archieved with $N\to\infty$. The continous sampling until the given precision rearched is an example of the sequential experiments. Suppose $N_w$ be
$$
N_w=\text{inf}\left\lbrace n:w^2n\ge 4z^2_{\alpha/2}\nu^2(\hat\theta)\right\rbrace,
$$
and $N_w$ is a random variable as the $\hat \theta$ vary randomly. Let $\hat \theta_N$ be the estimation for $\theta$ by using the first $N$ observations. By definition of $N_w$, we have
$$
4z^2_{\alpha/2}\nu^2\left(\hat\theta_{N_w}\right)\le w^2N_w< 4z^2_{\alpha/2}\nu^2\left(\hat\theta_{N_{w-1}}\right).
$$
This means that $w^2N\to 4z^2_{\alpha/2}\nu^2\left(\theta\right)$ as $\hat\theta_N\to \theta$. 

Random variables $W_n,n\ge 1$ are uniformly continuous in probability (UCIP) if for all $\epsilon >0$ there exists $\delta>0$ such that
$$
\mathbb{P}\left\lbrace \max_{0\le k\le n\delta}|W_{n+k}-W_n|\right\rbrace < \epsilon,\quad \forall n\ge 1.
$$

:::note Anscombe
If $N_w, w>0$, are positive integer-valued random variable with $w^2N_w\xrightarrow{P}c\in(0,\infty)$ as $w\to 0$, if $n_w=\lfloor c/w^2\rfloor$, and if $W_n, n\ge 1$ are UCIP, then
$$
W_{N_w}-W_{n_w}\xrightarrow{P}0,\quad \text{as}\quad w\to 0. 
$$
:::
**Proof**: Since $W_n$ is UCIP, given a $\delta\ge 0$, exists a $\epsilon\ge 0$ that $\mathbb{P}(\max_{n}|W_{n_w}-W_{n}|>\epsilon)<\epsilon$ where $n$ in $n_w\pm \delta/w^2$. The possibility to have $|W_{N_w}-W_{n_w}|>\epsilon$ is that $|W_{n}-W_{n_w}|>\epsilon$ for some $n$ within $n_w\pm \delta/w^2$, or $w^2|n_w-N_w|\ge \delta$. 
$$
\mathbb{P}(|W_{N_w}-W_{n_w}|>\epsilon) \le \mathbb{P}(w^2|N_w-n_w|\ge \delta)+\mathbb{P}\left(\max_{w^2|n_w-n|\ge \delta}|W_{n_w}-W_n|\ge \epsilon\right).
$$
However, the first term vanish by the requirement that $N_w\xrightarrow{P} n_w$ and the second term less than $\epsilon$ since $W_{n}$ is UCIP. This proved the theorem.

:::note Lemma
Let $X_i,i=1,\dots, n$ be iid random variables with mean $\mu$ and variance $\sigma^2$. Suppose $S_k=\sum_{i=1}^kX_i$, then for any $c>0$:
$$
\mathbb{P}\left(\max_{1< k\le n}|S_k|>c\right)\le \frac{n\sigma^2}{c^2}.
$$
:::
**Proof**: Suppose $A_k$ be the event that $S_k$ is the first partial sum with magnitude at least $c$. By the definition, $A_i,A_j$ is disjoint events if $i\ne j$. And we have:
$$
\mathbb{P}\left(\max_{1\le k\le n}|S_k|\ge c\right)=\mathbb{P}\left(\sum_{k=1}^nA_k\right)=\sum_{k=1}^n\mathbb{P}(A_k)\le \frac{1}{c^2}\mathbb{E}(S^2_n)=\frac{n\sigma^2}{c^2}.
$$
The last inequality comes from the fact that $\mathbb{E}S^2_nA_k\ge \sum_{i=1}^nc^2\mathbb{P}(A_k)$ for all $k\ge 1$. This proved the lemma.


With the lemma, we show that the normalized partial sum $W_n=\sqrt{n}\overline{X}_n$ are UCIP variables. In fact, it can be expressed as 
$$
W_{n+k}-W_n=\sum_{i=n+1}^{n+k}X_i/\sqrt{n}-\left(\sqrt{\frac{n+k}{n}}-1\right)W_{n+k},
$$
so we have
$$
\mathbb{P}\left(\max_{0\le k\le n\delta}|W_{n+k}-W_n|\ge \epsilon\right)\le \mathbb{P}\left((\sqrt{1+\delta}-1)|W_n+k|\ge \epsilon/2\right)+\mathbb{P}\left(\max_{0\le k \le n\delta}\bigg|\sum_{i=n+1}^{n+k}\frac{X_i}{\sqrt{n}}\bigg|\ge \epsilon/2\right).
$$
By Chebyshev's inequality and the lemma, we have the bound
$$
\frac{4(\sqrt{1+\delta}-1)^2\sigma^2}{n\cdot\epsilon^2}+\frac{4k\sigma^2}{n\cdot \epsilon^2}.
$$
This bound vanishes along $n\to \infty$ and $\delta\to 0$ so that $W_n$ is UCIP. By Anscombe's theorem, we have
$$
\sqrt{N}\left(\overline{T}_N-\mu(\theta)\right)\xrightarrow{d}N(0,\text{Var}_\theta(T)).
$$
By the delta method, we proved the conclusion that
$$
\frac{\sqrt{N_w}(\hat\lambda_N-\lambda)}{\nu(\hat\theta_N)}\xrightarrow{d}N(0,1),\quad \nu(\theta) = \lambda'(\mu_T(\theta))\sigma_T(\theta).
$$