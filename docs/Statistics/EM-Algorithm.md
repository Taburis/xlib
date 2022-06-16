# Expectation Maximization Algorithm

## Introduction
---
The Expectation-Maximization Algorithm (EM-Algorithm) is often used to estimate the parameters of the models with latent (unobservable) variables by maximizing the associated likelihood function. To state the idea of this algorithm, we define the following elements:
* $X$, a observable random variable with PDF $f_X(x)$ leads to the samples $\lbrace X_n \rbrace$. 
* $Z_i,i=1,\dots,K$, are hidden random variable (latent variable) with PDF $f_Z(z)$ and $X\perp Z$.
* $L(X,\theta)$ is the **observed likelihood function** that $X$ is partially known, and the $L(X,Z,\theta)$ is the **complete likelihood function** that $X,Z$ are known.

To find the maximum of the likelihood function, EM algorithm initializing the parameters by a first guess and then update these parameters iteratively to reach the upper bound. For the details, let the $\hat\theta^{(m)}$ be the parameters at the $m$-th iterations, the algorithm can be summarized as
1. **E-step**: Compute the function
$$
Q(\theta|\hat\theta^{(m)},\boldsymbol{x})=\mathbb{E}_{\hat\theta^{(m)}}\log L(\boldsymbol{x},\boldsymbol{z},\theta),
$$
where the expectation are sum over the hidden variable $Z$. Noticing that the explicity expression of this definition might not be unique. As an illustration, here we expand it by using the Bayes' formula such that
$$
Q(\theta|\hat\theta^{(m)},\boldsymbol{x})=\int\log L(\boldsymbol{x},\boldsymbol{z},\theta)k(\boldsymbol{z}|\hat\theta^{(m)},\boldsymbol{x})d\boldsymbol{z}
$$
where $k(\boldsymbol{z}|\theta,\boldsymbol{x})$ is defined as
$$
k(\boldsymbol{z}|\theta,\boldsymbol{x})=\frac{L(X,Z,\theta)}{L(X,\theta)}.
$$
2. **M-step**: Computing the estimator for $m+1$ step:
$$
\hat\theta^{(m+1)}=\argmax_{\theta\in\Omega}Q(\theta|\theta^{(m)},\boldsymbol{x}).
$$


## Explanation
---
The observed likelihood function connects with the complete likelihood function by the Bayes' theorem which leads to the function $k(\boldsymbol{z}|\theta,\boldsymbol{x})$. By definition, we have
$$
\log L(\boldsymbol{x},\theta)=\int \log L(\boldsymbol{x},\theta)k(\boldsymbol{z}|\theta,\boldsymbol{x})dz = \int \log \left\lbrace L(\boldsymbol{x},\boldsymbol{z},\theta)-\log k(\boldsymbol{z}|\theta,\boldsymbol{x})\right\rbrace k(\boldsymbol{z}|\theta,\boldsymbol{x})dz,
$$
where the first equation hold since the integral over $z$ has to be 1.

With this relationship and suppose the $m$-th step estimation $\hat\theta^{(m)}=\theta_0$, the posterior log-likelihood function based on the estimation $\theta=\theta_0$ and known samples $\boldsymbol{x}$ is given by using the equation above:
$$
\begin{aligned}
\log L(\boldsymbol{x},\theta)&=\int \log L(\boldsymbol{x},\theta)k(\boldsymbol{z}|\theta_0,\boldsymbol{x})dz,\\
&=\int \left\lbrace \log L(\boldsymbol{x},\boldsymbol{z},\theta) - \log k(\boldsymbol{z}|\theta,\boldsymbol{x})\right\rbrace k(\boldsymbol{z}|\theta_0,\boldsymbol{x})d\boldsymbol{z}.
\end{aligned}
$$
Furthermore, we need to show that $L(\hat \theta^{(m+1)}|\boldsymbol{x})\ge L(\hat \theta^{(m)}|\boldsymbol{x})$. This is equivalent to show that the second term above is decreasing with the iterations since the first term is insured to increase by required the maximization. In fact, the second term is a expectation $\mathbb{E}_{\hat\theta^{(m)}}\log k$ under the $k$ measure with assuming $\theta=\hat\theta^{(m)}$. And $\log(x)$ is concave function which have $\mathbb{E}(\log (x))\ge \log(\mathbb{E}(x))$ by Jensen's inequality. On the other hand, we notice that
$$
\begin{aligned}
\int \log\left[\frac{g(x)}{f(x)}\right]f(x)dx &\le \log \int  \frac{g(x)}{f(x)} f(x)dx\\
&=\log\int g(x)dx,
\end{aligned}
$$
This means that $\int \log(g)f(x)dx$ reach the upper bound of the integral if $g(x)=f(x)$ over all functions such that $\int g(x)dx=1$. This proved that the second term is a constant for all $\hat\theta$.
