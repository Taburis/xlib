# Sufficiency and Information

Given a sample $\lbrace X\rbrace_n$ draw from a distribution $f_\theta$ where $\theta\in \Theta$ is some parameters unknown. A statistics $T=T(\boldsymbol{X})$ is the called **sufficient** if the conditional probability $\mathbb{P}(\boldsymbol{X}|T=t)$ is independent on the $\theta$. It means that once the $T$ is given, there's no other statistics, say $H=H(\boldsymbol{X})$, can be used to determine the $\theta$ since $\mathbb{P}(H=h|T=t)=\mathbb{P}(\boldsymbol{X}_h|T=t)$ has no dependence on the $\theta$, where $H(\boldsymbol{X}_h)=h$. Suppose the distribution of $T$ is $g_\theta$ and follows the Bayesian's formula, the definition of $T$ is sufficient statistics if
$$
\frac{\mathbb{P}(\boldsymbol{X}=\boldsymbol{x}, T=t)}{\mathbb{P}(T=t)}\to \frac{f_\theta(\boldsymbol{x})}{g_\theta(t)}=H(x_1,\dots, x_n),\quad t=T(\boldsymbol{x})=T(x_1,\dots, x_n),
$$
where $H(\boldsymbol{x})$ is an arbitrary function independent on the $\theta$. 

:::note The Neyman **factorization theorem**:
The theorem states that if $T$ is a sufficient statistics for $\theta$ if and only if there are two non-negative functions $k_1,k_2$ such that
$$
f_\theta(x_1)\dots f_\theta(x_n)=k_1[T(x_1,\dots, x_n), \theta]k_2(x_1,\dots, x_n).
$$
:::
**Proof**: There are one-to-one mappings that $T_i=T_i(\boldsymbol{X})$ and $T=T_1$ (these could be $T_j=X_j$ for $j>1$), so that $X_i=w_i(T_1,\dots, T_n)$. Then the $f_\theta(x_1,\dots, x_n)=h_\theta(t_1,\dots, t_n)$ so that we can obtain the distribution of $T_1$ as
$$
g_\theta(t)=\int\dots\int h(t_1,\dots, t_n)dt_2\dots dt_n=k_1(T_1,\theta)\int\dots \int |J|k_2(w_1,\dots, w_n)dt_2\dots dt_n,
$$
where $|J|$ is the Jacobian of transformation that $dx_1\dots dx_n=|J|dt_1\dots dt_n$. Clearly the integral has no dependence on the $\theta$ so that the definition of the sufficiency can be reproduced.

## The Role of Sufficient Statistics

Although we know that the sufficient statistics carried all the information about $\theta$, the follow theorems detailed the role of the sufficient statistics in estimator.
:::note Rao-Black Theorem:
Suppose $X_1,\dots,X_n$ denote a sequence of iid random samples from a distribution $f_\theta$, where $\theta$ is unknown, and $T=T(\boldsymbol{X})$ is a sufficient statistics for $\theta$, and $Y=Y(\boldsymbol{X})$ is an unbaised estimator of $\theta$, then the statistics defined following
$$
\varphi(T) = \mathbb{E}(Y|T)
$$
is an unbiased estimator of $\theta$ with the variance no larger than that of $Y$.
:::
**Proof**: This follows from the inequality for two random variable $X,Y$ that 
$$
\mathbb{E}(X)=\mathbb{E}[\mathbb{E}(X|Y)],
$$
and 
$$
\text{Var}(X)\ge \text{Var}[\mathbb{E}(X|Y)].
$$
Let $X=Y$ and $Y=T$ the conclusion follows immediately.

The Rao-Black theorem tells us that if a minimum variance unbiased estimation (MVUE) for $\theta$ exists, then it can be a function of sufficient statistics. Because that if $Y$ is an MVUE of $\theta$, then $\mathbb{E}(Y|T)$ is a unbiased estimator of $\theta$ no worse than the $Y$. Practically, we can restrict the searching of MVUE to the functions of a sufficient statistics.
On the other hand, the sufficient statistics connected with the maximum likelihood estimator (MLE) by the following theorem:

:::note Theorem:
Suppose $T$ is a sufficient statistics for $\theta$ and a maximum likelihood estimator $\hat\theta$ of $\theta$ exists uniquely, then $\hat\theta=\theta(T)$, a function of $T$. 
:::
**Proof**: By the definition of the sufficiency, the likelihood function can be epxressed 
$$
L(\theta, x_1,\dots, x_n)=\prod_if_\theta(x_i)=g_\theta(t)H(x_1,\dots, x_n),
$$
and the MLE $\hat\theta$ maximizing the $L$ will maximize the $g_\theta(t)$ as well since no $\theta$ dependent in $H$. So that it must be a function of $T$ alone.

We know the MLEs are asymptotic unbiased estimators for $\theta$, so we usually obtain a unbiased estimator from sufficient statistics by the procedure:
1. Found a sufficient statistics $T$ for $\theta$.
2. Obtain the MLE $\hat\theta=\hat\theta(T)$ by the equation $\partial\ell(\theta)/\partial\theta=0$.
3. Calculate the $\mathbb{E}(\hat\theta)$ and express the $\theta=\theta(T)$ to obtain the unbiased estimator.

## Completeness

Let $Z$ be a random variable with a PDF or PMF as one member of the family functions $\lbrace h(z,\theta):\theta\in\Theta\rbrace$. It is called **complete family** if $\mathbb{E}(\varphi(z))=0$ implies that $\varphi(z)=0$ except on those points with probability equal 0.

The completeness of the distributions ensure that if there are two unbiased estimator for $\theta$, say $\varphi(T), \psi(T)$, as function of a sufficient statistics $T$, then $\varphi(T)=\psi(T)$ except some points with probability 0. A **complete sufficient statistics** is a sufficient statistics with a distribution belong to a complete family. 

:::note Lehmann and Scheffe:
Let $X_1,\dots, X_n$ be a random sample from a distribution $f_\theta,\theta\in\Theta$, and $T$ is a sufficient statistics with distribution belong to a complete family $g_\theta(T)$ (complete sufficient statistics). If there is a function of $T$, $\varphi(T)$, is an unbiased estimator of $\theta$, then it is unique MVUE of $\theta$. 
:::
**Proof**: The uniqueness follows from the discussion above. To show the minimum variance property, noticed that if there is any other unbiased estimator as functions of $X_i$, there must be an unbiased estimator as function of $T$ with no larger variance based on the Rho-Blackwell's theorem. The uniqueness of this estimator implies that it is MVUE for $\theta$.  

## Exponentail Class of Distributions

A family of the distributions with the following form is called the **exponential class**:
$$
f(x,\theta) = \exp\left\lbrace p(\theta)K(x)+H(x)+q(\theta)\right\rbrace,\quad x\in S
$$
where $S$ is the support of $X$. Furthermore, a exponential class is regular if it satisfies the following conditions:
1. $S$, the support of $X$, doesn't depends on $\theta$. 
2. $p(\theta)$ is a non-trivial continuous function of $\theta\in\Theta$. 
3. This condition is a little difference for discrete and continuous cases:
    * If $X$ is **continuous**, then each of $K(x)$ shouldn't be a constant and $H(x)$ is a conitnuous function of $x\in S$. 
    * If $X$ is **discrete**, then $K(x)$ has to be a non-trivial function of $x\in S$. 


The following theorem shows that there always exists a sufficient statistics for the regular exponential class
:::note Theorem:
For a given size of random sample $X_1,\dots, X_n$ with distribution from a regualr exponential class. Then the statistics $T=\sum_iK(x_i)$ is sufficient with the following properties:
1. The PDF or PMF for $T$ is:
$$
f_T(t, \theta)=R(t)\exp\lbrace p(\theta)t+nq(\theta)\rbrace,\quad t\in S_T,
$$
where $S_T$ is the support of $T$. Both $S_T$ and $R(T)$ is independent on $\theta$. 
2. The expectation and variance is:
$$
\begin{aligned}
E(T)&=-nq'(\theta)/p'(\theta),\\
\text{Var}(T)&=\frac{n}{p'(\theta)^3}\left\lbrace p''(\theta) q'(\theta) - q''(\theta)p'(\theta)\right\rbrace,
\end{aligned}
3. If $\theta\in (\alpha,\beta)$, then the $T$ is a complete sufficient statistics for $\theta\in (\alpha,\beta)$
:::
**Proof**: To prove the points one-by-one:
1. Let's construct the following transform that $T_1=\sum_iK(X_i)$ and $T_j=X_j$ for $j>1$. Then the Jacobian is the function of $J$. The sample PDF is 
$$
\begin{aligned}
f(X_1,\dots, X_n,\theta) d^nx&= \exp\left\lbrace p(\theta)\sum_{i=1}^nK(x_i)+\sum_{i=1}^nH(x_i)+nq(\theta)\right\rbrace dx_1\dots dx_n\\
&=\exp\left\lbrace p(\theta)t_1+nq(\theta)\right\rbrace\exp\left\lbrace\sum_{i=2}^nH(t_i)+H'(t_1,\dots, t_n)\right\rbrace |J|dt_1\dots dt_n.
\end{aligned}
$$
The distribution of $T_1$ follows from the intergration over the $t_2,\dots, t_n$. The second exponential alone with the Jacobian $J$ forms the $R(t_1)$ after the integration. Furthermore, the distribution above also shows that $T_1$ is sufficient statistics based on the factorization theorem.

2. To calculate the expectation of $T$, we noticed that $\mathbb{E}(T) = n\mathbb{E}[K(X)]$ and
$$
\frac{d}{d\theta}\int_a^b f_\theta(x) dx=[q'(\theta)+p'(\theta)K(x)]\cdot \int_b^a f_\theta(x)dx = 0
$$
The expectation follows. The variance can be obtained in the similar way (second derivative with respect to $\theta$).

3. To ensure $\mathbb{E}(\varphi(T))=0$, it implies that 
$$
\int_{S_T}\varphi(t)R(t)\exp\lbrace p(\theta)t\rbrace dt=0,
$$
since the $q(\theta)\ne 0$ due to the regulation requirement. The left side of the equation is a Laplace transformation and the only function with 0 as Laplace transform is 0 itself. However, the $p(\theta)$ is non-trivial either by the regulation condition. The only possible way is either $\varphi(t)$ or $R(t)$ vanishes. However, the vanishing of $R(t)$ means that the probability vanishes at that point. It follows that $\varphi(t)$ has to be 0 for the point $t$ with non-vanishing probability. 

The third point of the theorem implies that if we can found a unbiased estimator for $\theta\in (\alpha,\beta)$ as a function of $T=\sum_iK(X_i)$, then it must be a MVUE for $\theta$.

## Multiple Parameter Cases
The exponential class can be extended to the cases when the distribution has more than one parameters, say $f(x,\boldsymbol{\theta})$ where $\boldsymbol{\theta}\in \Theta$ is a vector with $m$ dimensions. The definition of the exponential class becomes:
$$
f(x,\boldsymbol{\theta})=\exp\left\lbrace \sum_{j=1}^mp_j(\boldsymbol{\theta})K_j(x)+H(x)+q(\boldsymbol{\theta})\right\rbrace,\quad x\in S.
$$
A **regular exponential class** requires the following conditions:
1. The support $S$ is independent on $\boldsymbol{\theta}$.
2. The parameter space $\Theta$ contained at least one interior point.
3. The function $p_i(\boldsymbol{\theta})\not\equiv 0, j=1,\dots, m$ are continuous and indepdent with each other. 
4. Depends on the random variable:
    * For the continuous case: the derivatives $K_j'(x), j=1,\dots, m$ are conitnuous for $x\in(a,b)$ and no one is a linear homogeneous function of the others. $H(x)$ is a continuous function in $(a,b)$ as well.
    * For the discrete case: $K_j(x)\not\equiv 0, j=1,\dots,m$ for $x\in S$ and no one is linear homogeneious function of the others.

Note that the linear homogeneious function means that $f(cx)=cf(x)$ for any scalar $c$. 

The probability to have sample $x_1,\dots, x_n$ from this class is
$$
f(x_1,\dots,x_n,\boldsymbol{\theta}) = \prod_if(x_i,\boldsymbol{\theta})=\exp\left\lbrace \sum_{i=1}^mp_j(\boldsymbol{\theta}) \sum_{i=1}^nK_j(x_i)+nq(\boldsymbol{\theta})\right\rbrace\exp\left\lbrace \sum_{i=1}^nH(x_i)\right\rbrace.
$$
:::note Sufficient Statistics for Multi-parameter Exponential Class:
For the regular exponential class above, there are $m$ sufficeint statistics:
$$
T_i= \sum_{j=1}^nK_i(x_j),
$$
with a joint probability distribution:
$$
f_T(\boldsymbol{t},\boldsymbol{\theta})=R(\boldsymbol{t})\exp\left\lbrace \sum_{j=1}^mp_j(\boldsymbol{\theta})t_j+nq(\boldsymbol{\theta})\right \rbrace.
$$
These sufficient statistics are complete if $n>m$.
:::

The exponential class covered a very often used distributions:

### Normal Distributions
Given a normal distribuiton $N(\mu,\sigma)$, it is a typical exponential class with a PDF
$$
f(x,\mu,\sigma) = \exp\left[-\frac{x^2}{2\sigma^2}-\frac{\mu^2}{2\sigma^2}+\frac{\mu x}{\sigma^2}-\ln \sqrt{2\pi\sigma^2}\right],
$$
so that the sufficient statistics are:
$$
T_1 = \sum_jX_j^2, \quad T_2 = \sum_jX_j.
$$
The transform $Z_1 = T_2/n$ and $Z_2=(T_1-T_2^2)/(n-1)$ is also sufficient and $\mathbb{E}(Z_1)=\mu$, $\mathbb{E}(Z_2)=\sigma^2$. It follows that $Z_1, Z_2$ are MVUEs for $\mu$ and $\sigma$, respectively.

For the multivariate normal distribution, it can be expressed as the exponential class form:
$$
f(\boldsymbol{X},\boldsymbol{\mu},\Sigma)=\exp\left\lbrace -\frac{\boldsymbol{X}^T\Sigma^{-1}\boldsymbol{X}}{2}+\boldsymbol{\mu}^T\Sigma^{-1}\boldsymbol{X}-\frac{\boldsymbol{\mu}\Sigma^{-1}\boldsymbol{\mu}}{2}-\frac{1}{2}\ln\det\Sigma-\frac{m}{2}\ln 2\pi\right\rbrace.
$$
Hence given a samples $\boldsymbol{X}_1,\dots,\boldsymbol{X}_n$ the sufficient statistics are:
$$
\boldsymbol{T}_1=\sum_{i=1}^n\boldsymbol{X}_i,\quad \boldsymbol{T}_2=\sum_{i=1}^n\boldsymbol{X}_i\boldsymbol{X}^T_i.
$$
The MVUEs for the mean and variance can be obtained by calculating the expectation sufficient statistics in marginal distributions:
$$
\hat \mu_j=\frac{1}{n}\sum_{i=1}^nX_{ij}=\overline{X}_j,\quad \hat\sigma_j^2=\frac{1}{n-1}\sum_{i=1}^n(X_{ij}-\overline{X}_j)^2,
$$
where $X_{ij}$ is the $i$-th component of $\boldsymbol{X}_j$.

### Multinomial Distributions
The multinomial distribution describes the samples draw from a Bernoulli experiment with $m$ possible outcomes. This Bernoulli experiment has $p_1,\dots, p_m$ parameters stand for the $i$-th outcome for each trial. Let $X_i$ with possible value $0$ or $1$ stands for the Bernoulli outcome. The PMF can be formed into
$$
f(x,p_1,\dots, p_{m-1})=\prod_{i=1}^{m-1}p_i^{x_i}\left(1-\sum_{i=1}^{m-1}p_i\right)^{1-\sum_{i=1}^{m-1}x_i}=\exp\left[\sum_{i=1}^{m-1}x_i\ln \Bigg(\frac{p_i}{1-\sum_{j=1}^{m-1}p_j}\Bigg)+\ln\Bigg(1-\sum_{i=1}^{m-1}p_i\Bigg)\right],
$$
It is a regular exponential class and the sufficient statistics are
$$
T_i = \sum_{j=1}^nX_{ij}, \quad j=1,\dots, m-1,
$$
where $n$ is the sample size and $X_ij$ is the $i$-th outcome in $j$-th trial. Calculating the expectation of $T_i$, we get the MVUE is $\hat p_i=T_i/n$. Furthermore, since $\mathbb{E}(T_iT_j) = (n-1)np_ip_j$. A trick can be used that given $T_i$, the conditional distribution of $\mathbb{E}(T_j|T_i=t)$ is a binomial distribution $b[n-T_i, p_j/(1-p_i)]$. The expectation of $\mathbb{E}(Y_jY_i)=\mathbb{E}[Y_i\mathbb{E}(Y_j|Y_i)]$. Hence the MVUE for $p_ip_j$ is $Y_iY_j/[n(n-1)]$.