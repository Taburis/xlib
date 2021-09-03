# Statistical Inference

Given a random variable $X$, the samples draw from $X$ will have the same distribution as $X$, and denote them as random variables $X_1,\dots,X_n$ where $n$ is called **sample size**. A functional $T(X_1,\dots,X_n)$ (since a random variable is a function) on the samples is called **statistics**. 

## Order Statistics

Given a sample $\lbrace X_n\rbrace$ draw from a PDF $f(x)$ for random variable $X$, let $\lbrace Y_n\rbrace$ represents the statistics of $\lbrace X_n\rbrace$ with the order that $Y_1<Y_2<\dots<Y_n$. 
1. The probability to have $\lbrace y_n\rbrace$ is denoted as $g(y1,\dots, y_n)$, then
$$
g(y_1,\dots, y_n)=n!\prod_{i=1}^nf(y_i),\quad y_1<y_2<\dots<y_n.
$$
The PDF of $Y_k$ is given by the marginal PDF
$$
\begin{aligned}
g_k(y_k)&=\int_a^{y_2}\dots\int_a^{y_{k}}\int_{y_k}^b\dots\int_{y_k}^b n!f(y_1)\dots f(y_{k-1})f(y_{k+1})\dots f(y_n)dy_1\dots dy_{k-1}dy_{k+1}\dots y_n\\
&=\frac{n!}{(k-1)!(n-k)!}F^{k-1}(y_k)\left[1-F(y_k)\right]^{n-k}f(y_k),\quad a\le y_k\le b,
\end{aligned}
$$
where $a,b$ is the lower/upper bound of the range of $X$.
2. A $p$-th **Quantile** $\xi_p$ is defined as $\xi_p=F^{-1}(p)$ where the $F(x)$ is the CDF of $X$. For instance, a $0.5$ quantile is the median of $X$. An approximated estimator of the $\xi_p$ is $\hat\xi_p=Y_k$ where $k=\lfloor p(n+1) \rfloor$ and $n$ is the size of the samples. 
3. The probability of $\xi_p\in (Y_i,Y_j)$ for the order statistics $\lbrace Y_n\rbrace$ is given by
$$
\mathbb{P}(Y_i < \xi_p < Y_j) = \sum_{k=i}^{j-1}\binom{n}{k}p^k(1-p)^k,
$$
Let $\gamma= \mathbb{P}(Y_i < \xi_p < Y_j)$, then $(Y_i,Y_j)$ formed the $100\cdot\gamma\%$-**confidential interval**.

**Proof**
1. There are $n!$ possible ways to get the order $y_1<\dots<y_n$ from the samples $\lbrace x_n\rbrace$, which are drawn at the probability $\prod_if(x_i)$.
2. To show if $Y_k$ is an unbiased estimator for $\xi_p$, we need to show that $\mathbb{E}(Y_k)=\xi_p$. In fact, the expected CDF is given by
$$
\begin{aligned}
\mathbb{E}[F(Y_k)]&=\int_a^bg_k(y)F(y)dy\\
&=\int_a^b \frac{n!}{(k-1)!}{(n-k)!}F^k(y)[1-F(y)]^{n-k}dy\\
&=\int_0^1 \frac{n!}{(k-1)!}{(n-k)!} z^k(1-z)^{n-k}dz\\
&=\frac{k}{n+1},
\end{aligned}
$$
where the last step follows from the iterative partial integrals. It is exactly the definition of $\xi_p$.
3. Suppose $\xi_p$ is the $p$-th quantile, then $F(\xi_p)=p=\mathbb{P}(X<\xi_p)$. For $Y_i<\xi_p<Y_j$ in order statistics, it implies that there are $i$ elements needs to be less then $\xi_p$ while $n-j-1$ elements should be larger than $\xi_p$. The samples of the order statistics obeys a binomial distribution with probability $p$.  This leads to the expression above.

## Descriptive Statistics

Descriptive statistics provide simple summaries about the samples that have been made.

Given a random variable $X$, the PDF $f(x)$, the CDF $F(x)$, and a sample $\lbrace X_n\rbrace$ drawn for $X$, the very basic statistics for samples are the **Sample mean** $\overline{X}=\sum_{i=1}^nX_i/n$ and the **Sample variance** $S^2=\sum_{i=1}^n(X_i-\overline{X})^2/(n-1)$. 

:::info Theorem:
* Given a statistics $T(\boldsymbol{X})=\sum_ia_iX_i$ and $W=\sum_jb_jY_j$, then
$$
\mathbb{E}(T)=\sum_ia_i\mathbb{E}(X_i),
$$
provided the expectation exists for every $X_i$. And 
$$
\text{Cov}(T,W)=\sum_{ij}a_ib_j\text{Cov}(X_i,Y_j),
$$
provided the variance exists for $X_i,Y_j$. Furthermore,
$$
\text{Var}(T) = \sum_ia^2_i\text{Var}(X_i)+2\sum_{i<j}a_ia_j\text{Cov}(X_i,X_j).
$$
* For sample mean $\overline{X}$ and sample variance $S^2$, such
$$
\mathbb{E}(\overline{X})=\mu,\quad \text{Var}(\overline{X})=\sigma^2/n,
$$
where $\mu$ and $\sigma$ is the mean and variance of the random variable $X$ the samples draw from.
and
$$
\mathbb{E}(S^2)=\sigma^2.
:::

Furthermore, other useful descriptive statistics are defined follows:
1. Mode $x_{\text{max}}$ : the sample that occurs most often in samples, $x_{\text{max}}=\text{argmax}_xf(x)$.
2. Median $m$ : The middle of the samples when list along a monotonic increasing order. Ideally $F(m)=1/2$. It is $0.5$ quantile $m=\xi_{0.5}$ and can be estimated by $y_{(n+1)/2}$ or $(y_{n/2}+y_{n/2+1})/2$ in case the $n$ is even and $y_k$ is the $k$-th largest number of the samples.
3. Standardized Moment: $\tilde{\mu}_k$ for $k$-th moment is defined as
$$
\tilde{\mu}_k=\frac{1}{n}\sum_{i=1}^nz_i=\frac{\mu_k}{\sigma^k}, \quad z_i=\frac{X_i-\mu}{\sigma}, \mu_k=\mathbb{E}\left[(X-\mu)^k\right],\quad \sigma^2=\mathbb{E}\left[(X-\mu)^2\right].
$$
$\tilde{\mu}_1=0$ and $\tilde{\mu}_2=1$ by definition for all random variables. The standardized moment is dimensionless. Starting from the third standardized moment, it has different descriptive meaning:
* **Skewness**: a measure of the asymmetry of the probability distribution of a real-valued random variable. Positive skewness implies the left tail is longer and the negative sign means the right tail is longer.  For a unimodal distribution, the relation between the mean, median and mode are shown for different skewness signs are shown below:
![skewness relationship with mean, median and mode](/img/docs/Image_skewness_mean_median_mode_relation.png)
* **Kurtosis**: The kurtosis describes the spread of the distribution. A higher kurtosis implies the PDF is wider spread. Smaller kurtosis means more concentrated distribution around the mean.

### Boxplot
