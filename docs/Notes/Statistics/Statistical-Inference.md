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

## Sampling
---
Sampling plays a central role during inference. The population is a pool where samples drawn from. The **sampling error** of a statistics $X$ is the error from the finite sample observations. A randomness is required when generating the samples to ensure the samples are effective. Suppose the population size is $N$, and the size of samples is $n$. We want to know a range $2\sigma_X$ so that $(1-\alpha)\times 100\%$ of values are expected fall within it. and this variance $\sigma_X$ is the $\alpha$-confidence **margin of error** (MOE).

To find the MOE, we first consider that $N\gg n$, which the CLT is applicable. It means that the mean $\overline{X}_n$ over the samples follows a normal distribution $\overline{X}_n\sim N(\overline{X}, \sigma^2/n)$ where $\overline{X}$ is the mean of $X$ over the population. The $\alpha$-confidence margin of error can be estimated from the CTL so that $Z=\sqrt{n}(\overline{X}_n-\overline{X})/\sigma$ and $Z\sim N(0,1)$:
$$
\text{MOE}_\alpha=z_\alpha\times \sqrt{\frac{\sigma^2}{n}},
$$
where $z_\alpha$ is the point that $F(Z=z)=1-\alpha$ where $F(z)$ is the CDF of $N(0,1)$.

For the case the $N$ is not large comparing to $n$, and the outcomes are two mutual exclusive results, say 0 or 1, we can estimate the MOE in more detials. Suppose sampling is performed with replacement (boostrap), the distribution of $X$ follows the binomial distribution. Then we have $\sigma^2=p(1-p)$ is the variance of the binomial distribution, and the binomial distribution can be approximated by normal distribution so that we have
$$
\text{MOE}_\alpha=z_\alpha\times \sqrt{\frac{\sigma^2}{n}}\approx z_\alpha\times \sqrt{\frac{p(1-p)}{n}},
$$
where $p$ is the average of the $X$.
For the non-replacement sampling, it follows the hypergeometric distribution with a variance $np(1-p)(N-n)/(N-1)$ so that it can be approximated as
$$
\text{MOE}_\alpha\approx z_\alpha\times \sqrt{\frac{p(1-p)}{n}}\times \sqrt{1-\frac{n}{N}},
$$
where the $\sqrt{1-n/N}$ is a correction factor for non-replacement sampling.

### Bootstrap Resampling

A bootstrap resampling is an statistics augmentation method. It treat a given sample as population, repeatly draw samples with replacement from that population to generate new bootstrap samples. 

### Percentile Bootstrap Confidence Intervals

Let $\lbrace X_n\rbrace$ be a sample with size $n$ and $\hat \theta(\boldsymbol{X})$ is an estimator, the bootstrap bootstrap confidence interval is obtained from the following procedure:
1. Draw a size $n$ sample $\lbrace X_n^*\rbrace$ from the sample $\lbrace X_n\rbrace$ with replacement.
2. Obtained the estimated value $\hat \theta = \hat \theta(\boldsymbol{X}^*)$.
3. Repeat the step 1 and 2 $N$ times to form the order statistics $\hat\theta^*_1\le \dots\le\hat \theta_N^*$, and $m=\lfloor N\alpha/2\rfloor$, then the interval:
$$
(\hat \theta^*_m, \hat\theta^*_{N+1-m})
$$
is the $\alpha/2\times 100\%$ confidence interval for $\theta$. $N$ should be a large number ($N >3000$ usually).

:::note Explanation
The bootstrap samples $\lbrace X^*_n\rbrace$ follows the empirical distribution $\hat F_n(x)$ from the sample $\lbrace X_n\rbrace$. 
:::