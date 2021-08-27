# Statistical Inference

Given a random variable $X$, the samples draw from $X$ will have the same distribution as $X$, and denote them as random variables $X_1,\dots,X_n$ where $n$ is called **sample size**. A functional $T(X_1,\dots,X_n)$ (since a random variable is a function) on the samples is called **statistics**. Define the **Sample mean** $\overline{X}=\sum_{i=1}^nX_i/n$ and the **Sample variance** $S^2=\sum_{i=1}^n(X_i-\overline{X})^2/(n-1)$. 

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

