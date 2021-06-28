# Kalman Filter

## Estimator Fusion

Let $x_i$ be a random variable with a PDF distribution $X_i\sim p_i(\mu_i, \sigma_i)$. A sequence of estimators on $\mu$ are random variables $x_i,i=1,\dots,N$ with the same mean values $\mu$. A linear estimator from these $X_i$ is a linear combination $Y=\sum_i\alpha_iX_i=\boldsymbol{\alpha}^T\cdot \boldsymbol{X}$ such that $\mathbb{E}(Y)=\mu$, then:

* $\mathbb{E}(Y)=\mu$, $\text{Var}(Y)=\boldsymbol{\alpha}^T\Sigma\boldsymbol{\alpha}$, where $\Sigma$ is covariance matrix of $\boldsymbol{X}$.
* Since the covariance matrix is diagonalizable, we can assume that $X_i\perp X_j$ for any $i\ne j$, then $Y$ is an optimal linear estimator if
$$
\alpha_i= \frac{1/\sigma_i^2}{\sigma^2_Y},\quad \sigma^2_Y=\sum_j\frac{1}{\sigma_j^2},
$$ 
the variance of $Y$ is $\text{Var}(Y)=\sigma^2_y$.

In a special case that $\sigma_i=\sigma$, all the variance are the same, the merged estimator converges to $Y=\sum_iX_i/N$. 

## Incremental Estimator Fusion

Fusion the estimator can also be carried out by merging the estimator one-by-one, and the final result is the same as we fusion them altogether like we did above. This makes a more practical way to fusion the estimators, which is known as **Kalman filtering** procedure.

Suppose we are observing a quantity $x$ at different time steps, and the estimator at time $t$ is a random variable $X_i\sim p_i(\mu, \sigma_i)$, and these estimators are mutually independent. Let's $Y_k=\sum_{i=1}^k \alpha_iX_i$ be the optimal linear estimator comes from fusing the first $k$ (include the $k$-th) estimators, where $\alpha_i$ is given above. This procedure can be carried out iteratively:

1. Initial estimation $X_1\sim p_1(\mu, \sigma_1^2)$ and $X_2\sim p_2(\mu, \sigma_2^2)$, and let $Y_1 = X_1$.
2. Get $Y_n$ from fusing $X_n$ and $Y_{n-1}$ together
$$
Y_n=\frac{\sigma_{X_{n}}^2}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2}Y_{n-1}+\frac{\sigma_{Y_{n-1}}^2}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2}X_n = Y_{n-1}+K(X_n-Y_{n-1}),\quad K=\frac{\sigma^2_{Y_{n-1}}}{\sigma_{Y_{n-1}}^2+\sigma_{X_{n}}^2},
$$ 
   where the $\sigma_{Y_n}=\sigma_{X_n}^2\cdot \sigma_{Y_{n-1}}^2/(\sigma_{Y_{n-1}}^2+\sigma_{X_n}^2)=K\sigma_{X_n}^2=(1-K)\sigma_{Y_{n-1}}^2$.
3. Repeat this procedure until all the estimations have been merge to $Y_N$. 

## Multivariate Estimator Fusion

Now we extend the fusion process into the multivariate cases.

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

## Kalman Filtering on Linear Systems

Let a $m$-dimensional vector $\boldsymbol{X}_t$ stands for the state of a linear system at the time step $t$, the transition of the states obay a linear equation
$$
\boldsymbol{X}_{t+1}=T_t\boldsymbol{X}_{t}+\boldsymbol{W}_t,
$$
where $T_t$ is a transistion matrix and $\boldsymbol{W}_t$ is a vector of random variables that $\boldsymbol{W}_t\sim N_m(0,\Sigma_{W_t})$, Gaussian distribution with covariance matrix $\Sigma_t$. The observation $\boldsymbol{Y}_t$ on the $\boldsymbol{X}_t$ is modeled by a linear equation as follow:
$$
\boldsymbol{Y}_t = H_t\boldsymbol{X}_t+\boldsymbol{V}_t,
$$
where $\boldsymbol{V}_k\sim N_m(0,\Sigma_{V_k})$. 

Denoting $\boldsymbol{\hat X}_{t|t-1}$ as the **priori estimator** of $\boldsymbol{X}_t$ using upto and include the information at the time $t-1$.
The **posteriori estimator** $\boldsymbol{X}_{t|t}$ is the estimator for $\boldsymbol{X}_t$ by using upto and include the information at the time $t$. Iterative fusion procedure build up the relation between the priori and posteriori estimator as:

$$
\begin{aligned}
H \boldsymbol{\hat X}_t &= H\boldsymbol{\hat Y}_{t|t}+H_tK_X(\boldsymbol{y}_n-H\boldsymbol{\hat X}_{t|t-1}),\\
K_X &= \Sigma_{t|t-1}H_t^T(H_t\Sigma_{t|t-1}H^T_t+\Sigma_{Y_t})^{-1},
\end{aligned}
$$
where $\boldsymbol{y}_t$ is the observation at $t$ and $\Sigma_{t|t-1}$ is the covariance matrix of $\boldsymbol{\hat X}_{t|t-1}$. The difference $\boldsymbol{y}_n-H\boldsymbol{\hat X}_{t|t-1}$ is also known as **innovation** or **mesurement residual**. Since we assume that $H^{-1}$ exists, it leads to the Kalman filtering procedure:

1. Setting $\boldsymbol{\hat X}_{1|1} = H_1^{-1}\boldsymbol{y}_1$
2. Update the priori estimator as
$$
\boldsymbol{\hat X}_{t|t-1} = T_t\boldsymbol{X_{t-1|t-1}},
$$
with the covariance matrix
$$
\Sigma_{t|t-1} = T_t\Sigma{t-1|t-1}T_t^T+\Sigma_{W_t}.
$$
3. Update the posterior estimator as
$$
\boldsymbol{\hat X}_{t|t} = \boldsymbol{\hat X}_{t|t-1} + K_{X_t} R_t, \quad K_{X_t} = \Sigma_{t|t-1}H_t^T(H_t\Sigma_{t|t-1}H^T_t+\Sigma_{Y_t})^{-1},
$$
with covariance
$$
\Sigma_{t|t} = (I-K_{X_t}H_t)\Sigma_{t|t-1}.
$$
4. Repeat the procedure until to the end of the time series.