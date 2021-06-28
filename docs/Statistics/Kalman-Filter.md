# Kalman Filter

## Estimator Fusion

Let $x_i$ be a random variable with a PDF distribution $X_i\sim p_i(\mu_i, \sigma_i)$. A sequence of estimators on $\mu$ are random variables $x_i,i=1,\dots,N$ with the same mean values $\mu$. A linear estimator from these $X_i$ is a linear combination $Y=\sum_i\alpha_iX_i=\boldsymbol{\alpha}^T\cdot \boldsymbol{X}$ such that $\mathbb{E}(Y)=\mu$, then:

* $\mathbb{E}(Y)=\mu$, $\text{Var}(Y)=\boldsymbol{\alpha}^T\Sigma\boldsymbol{\alpha}$, where $\Sigma$ is covariance matrix of $\boldsymbol{X}$.
* Since the covariance matrix is diagonalizable, we can assume that $X_i\perp X_j$ for any $i\ne j$, then $Y$ is an optimal linear estimator if
$$
\alpha_i= \frac{1/\sigma_i^2}{\sigma^2_Y},\quad \sigma^2_Y=\sum_j\frac{1}{\sigma_j^2},
$$ 
the variance of $Y$ is $\text{Var}(Y)=\sigma^2_y$.

## Incremental Estimator Fusion

Suppose we are observing a quantity $x$ at different time steps, and the estimator at time $t$ is a random variable $X_i\sim p_i(\mu, \sigma_i)$, and these estimators are mutually independent. Let's $Y_k=\sum_{i=1}^k \alpha_iX_i$ be the optimal linear estimator comes from fusing the first $k$ (include the $k$-th) estimators, where $\alpha_i$ is given above. This procedure can be carried out iteratively:

1. Initial estimation $X_1\sim p_1(\mu, \sigma_1^2)$ and $X_2\sim p_2(\mu, \sigma_2^2)$, and let $Y_1 = X_1$.
2. Get $Y_n$ from fusing $X_n$ and $Y_{n-1}$ together
$$
Y_n=\frac{\sigma_{X_{n}}^2}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2}Y_{n-1}+\frac{\sigma_{Y_{n-1}}^2}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2}X_n = Y_{n-1}+K(X_n-Y_{n-1}),\quad K=\frac{\sigma^2_{Y_{n-1}}}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2},
$$ 
   where the $\sigma_{Y_n}=\sigma_{X_n}^2\cdot \sigma_{Y_{n-1}}^2/(\sigma_{Y_{n-1}}^2+\sigma_{X_n}^2)=K\sigma_{X_n}^2=(1-K)\sigma_{Y_{n-1}}^2$.
3. Repeat this procedure until all the estimations have been merge to $Y_N$. 

## Multivariate Estimator Fusion

Let $\boldsymbol{X}_i\sim p(\boldsymbol{\mu},\Sigma_i)$, a set of random variable vectors with dimensional $N$, where $\Sigma_i$ is the covariance matrix, be unbiased estimators for the same quantity $\boldsymbol{\mu}$. The linear unbiased estimator $\boldsymbol{Y}$ is defined as
$$
\boldsymbol{Y} = \sum_iA_i\boldsymbol{X}_i,\quad \sum_iA_i=I
$$
where $I$ is a $N$-dimensional unit matrix and $A$ is a $N\times N$ matrix. Suppose the estimators are mutually independent, i.e. $\boldsymbol{X}_i\perp\boldsymbol{X}_j$, $i\ne j$, then

$$
\mathbb{E}(\boldsymbol{Y}) =\boldsymbol{\mu}, \quad \text{Var}(Y) = \sum_{j}A_j\Sigma_jA^T_j,
$$
The condition for $\boldsymbol{Y}$ to be the best unbiased estimator (BUE) is that
$$
A_i=\Sigma_i^{-1}\Sigma_Y,\quad \Sigma_Y = \left(\sum_{j}\Sigma_j^{-1}\right)^{-1},
$$

and the covariance $\text{Cov}(Y) = \Sigma_Y$. 

The fusion of estimators can be done iteratively as follow:
1. Let $\boldsymbol{Y}_1 =\boldsymbol{X}_1$,
2. Computing the kernel $K$ and update the $Y_{n}$ by the formula 
$$
\boldsymbol{Y}_n=\boldsymbol{Y_{n-1}}+K(\boldsymbol{X}_n-\boldsymbol{Y}_{n-1}),\quad K=\Sigma_{Y_{n-1}}(\Sigma_{Y_{n-1}}+\Sigma_{X_{n}})^{-1},
$$
and update the covariance $\text{Cov}(\boldsymbol{Y}_n) = (I-K)\Sigma_{Y_{n-1}}$.
3. Repeat the step 2 until all the $M$ observations has been merged. 
