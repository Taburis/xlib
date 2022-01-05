# Mixture Model

## Definition
---
A finite mixture model assumes that the observation represented by the random variables $X$ is draw from a distribution that composed by $K$ sub-distributions (component). To formulate the mixture model, define the following variables:
1. $Z_j,j=1,\dots, K$ is the latent variables can taken values $z_j=0,1$ represents the data comes from the $j$-th component.
2. $X_i,i=1,\dots, N$ stands for the observations from the model where $N$ is the sample size.
3. $\theta_j,j=1,\dots, K$ is the parameters for the $j$-th component.

By the Bayes' law, the probability of $X$ can be expressed as
$$
\mathbb{P}(x)=\sum_{j=1}^K\mathbb{P}(x|Z_j=1)\mathbb{P}(Z_j=1),
$$
where $\mathbb{P}(x|Z_j=1)$ is the probability of the $X$ when it comes from the $j$-th component. If we assume that all the components have a distribution $p(x,\theta)$, this leads to the probability distribution for mixture model:
$$
p(x)= \sum_{j=1}^K\pi_jp(x,\theta_j),
$$
where $\pi_j=\mathbb{P}(Z_j=1)$, $\sum_{j=1}^K\pi_j=1$, and $\theta_j$. The multivariated case can be defined similarly.

## Expectation and Maximization Algorithm
---
The expectation and maximization (EM) algorithm can be used to learn the parameters of the mixture model. It works well especially for normal distributions.
Suppose $\theta_j$ stands for the parameters of the $j$-th components, then the parameters of mixture models are 
$$
\lbrace \pi_j,\theta_j\rbrace\quad j = 1, \dots, K.
$$
Given the samples $X_j,j=1,\dots,N$, the latent random vector associated to the $i$-th samples is denoted as $\boldsymbol{Z}_i$ where $i=1,\dots, N$. And let $z_{ij}$ be the $j$-th component of $\boldsymbol{z}_i$ and there is only one component of $\boldsymbol{z}_i$ is 1 and rest must be vanish by the definition. Then we have
$$
p(z_{ij}=1|\boldsymbol{x},\boldsymbol{\pi},\boldsymbol{\theta})=\frac{L(\boldsymbol{x},z_{ij}=1,\boldsymbol{\pi},\boldsymbol{\theta})}{L(\boldsymbol{x},\boldsymbol{\pi},\boldsymbol{\theta})}=\frac{\pi_jp(x_i,\theta_j)}{\sum_{j=1}^K\pi_jp(x_i,\theta_j)}.
$$

The complete likelihood function is
$$
L(\boldsymbol{x},Z,\boldsymbol{\pi},\boldsymbol{\theta})=\prod_{i=1}^N\prod_{j=1}^{K}\left[\pi_jp(x_i,\theta_j)\right]^{z_{ij}}
$$
The log-likelihood function is then
$$
\ell(\boldsymbol{x},Z,\boldsymbol{\pi}, \boldsymbol{\theta})=\sum_{i=1}^N\sum_{j=1}^{K}z_{ij}\left[\ln\pi_j+\ln p(x_i,\theta_j)\right]
$$
So that we can define the expectation of the log-likelihood function as
$$
\begin{aligned}
Q(\boldsymbol{\pi},\boldsymbol{\theta}|\boldsymbol{\pi}^*,\boldsymbol{\theta}^*,\boldsymbol{x})=\mathbb{E}_{\pi^*,\theta^*}\ell(\boldsymbol{x},Z,\boldsymbol{\pi}, \boldsymbol{\theta}) = \sum_{i=1}^N\sum_{j=1}^K\mathbb{E}_{\pi^*,\theta^*}(z_{ij})\left[\ln\pi_j+\ln p(x_i,\theta_j)\right],
\end{aligned}
$$
where $\mathbb{E}_{\pi^*,\theta^*}(z_{ij}) = p(z_{ij}=1|\boldsymbol{x},\boldsymbol{\pi}^*,\boldsymbol{\theta}^*)$. Until this step, we have all the formula we need for EM algorithm, the optimization is:
1. Initiate with $\boldsymbol{\theta}^{(0)}$ and $\boldsymbol{\pi}^{(0)}$;
2. Calculate the expectation 
$$
Q(\boldsymbol{\pi},\boldsymbol{\theta}|\boldsymbol{\pi}^{(m)},\boldsymbol{\theta}^{(m)},\boldsymbol{x})
$$
using the $m$-th step iteration of parameters $\boldsymbol{\pi}^{(m)}$ and $\boldsymbol{\theta}^{(m)}$;
3. Update the parameters by
$$
\left\lbrace \boldsymbol{\pi}^{(m+1)}, \boldsymbol{\theta}^{(m+1)}\right\rbrace=\argmax_{\boldsymbol{\pi},\boldsymbol{\theta}} Q(\boldsymbol{\pi},\boldsymbol{\theta}|\boldsymbol{\pi}^{(m)},\boldsymbol{\theta}^{(m)},\boldsymbol{x}).
$$
Repeat the 2 and 3 steps until the converge condition reached.


## Gaussian Mixture
---
One widely used distribution for the mixture model is a normal distribution $p(\boldsymbol{x},\boldsymbol{\theta})=N(\boldsymbol{x},\boldsymbol{\mu},\Sigma)$. In this case, the EM algorithm can be further simplified due the maximum estimation of the mean and the variance of a normal distribution is well defined. 

Let's define 
$$
p_{ij}^{(m)} = p\left(z_{ij}=1|x_i, \pi^{(m)}_j,\boldsymbol{\mu}_i^{(m)},\Sigma_i^{(m)}\right)=\frac{\pi_j^{(m)}N\left(\boldsymbol{x}_i,\boldsymbol{\mu}_j^{(m)},\Sigma_j^{(m)}\right)}{\sum_{j=1}^K\pi_j^{(m)}N\left(\boldsymbol{x}_i,\boldsymbol{\mu}_j^{(m)},\Sigma_j^{(m)}\right)}
$$

With this function, the estimations at the $m+1$ step based on the $m$-th iterations can be expressed as
$$
\begin{aligned}
\pi^{(m+1)}_j&= \frac{1}{N}\sum_{i=1}^Np_{ij}^{(m)},\\
\boldsymbol{\mu}^{(m+1)}_j&=\frac{\sum_{i=1}^Np^{(m)}_{ij}\boldsymbol{x}_i}{\sum_{i=1}^Np^{(m)}_{ij}},\\
\Sigma^{(m+1)}_j&=\frac{1}{\sum_{i=1}^Np_{ij}^{(m)}}\sum_{i=1}^N\left\lbrace p_{ij}^{(m)}\left[\boldsymbol{x}_i-\boldsymbol{\mu}_i^{(m+1)}\right]\cdot \left[\boldsymbol{x}_i-\boldsymbol{\mu}_i^{(m+1)}\right]^T\right\rbrace.
\end{aligned}
$$