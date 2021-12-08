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

The Rao-Black theorem tells us that the minimum variance unbiased estimation (MVUE) must be a function of sufficient statistics. 
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