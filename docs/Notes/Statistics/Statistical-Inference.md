# Statistical Inference

Given a random variable $X$, the samples draw from $X$ will have the same distribution as $X$, and denote them as random variables $X_1,\dots,X_n$ where $n$ is called **sample size**. A functional $T(X_1,\dots,X_n)$ (since a random variable is a function) on the samples is called **statistics**. 

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

The quartiles are used to illustrate the distribution of samples: $Q_0=\xi_{0}$, $Q_1=\xi_{0.25}$, $Q_2=\xi_{0.5}$, $Q_3=\xi_{0.75}$, and $Q_4=\xi_{1}$. The $Q_0$ and $Q_1$ is the minimum and maximum of the samples. The difference of $I=Q_3-Q_1$ is called the **interquanrtile range**. Based on this, the term is defined the fence as:
* **Inner fence**: $[Q_1-1.5I, Q_3+1.5I]$
* **Outer fence**: $[Q_1-3I, Q_3+3I]$
* **Outlier** is the region outside the inner fence. The region between the inner fence and outter fence is called **mild outlier** and the region outside the outer fence is called **extreme outlier**.

The boxplot is a summary plot for showing the quartiles and each of outliers are usually needs to be plot as dots. The example of boxplot is shown below:
![boxplot](/img/docs/Image_boxplot.png)