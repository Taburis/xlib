# Optimization

Given a system formulated by a function $\mathcal{L}(\boldsymbol{x}|\boldsymbol{\theta})$, the optimization is to find a point $\boldsymbol{\hat\theta}$ that maximize or minimize the function $\mathcal{L}(\boldsymbol{x}|\boldsymbol{\theta})$ over all possible $\boldsymbol{x}$. 

## Gradient Searching
One of the most practical solutions is based on the fact that gradient $\nabla\mathcal{L}$ always pointed to the fast direction that increase the $\mathcal{L}$, and the gradient is 0 at the maximum or minimum points of $\mathcal{L}$. The idea of gradient searching can be illustrated by the gradient descent algorithm (searching for minimum point) as an example:
1. Start at a random point in parameter space $\boldsymbol{\theta}$
2. Update the point following the gradent 
$$
\boldsymbol{\theta} \to \boldsymbol{\theta} - r\nabla_\theta\mathcal{L}(\boldsymbol{x}|\boldsymbol{\theta}),
$$
where $r$ is a learning rate.
3. Repeat the steps above until converge.

In real practice, despite the cost of calculating the gradient is expensive for each step, but also several points need to be aware:
* The gradient usually comes from a grid difference. It is critical to have a good estimation on the gradient.
* The gradient searching may leads to a saddle point where $\nabla_\theta \mathcal{L}=0$ but it is neither locally maximum nor minimum.
* Choosing the learning rate is subtle task. A larger learning rate may needs fewer iterations to approach the solution, but it is also easier to past the minimum point so that the searching will move around the solution back and forth.
* Converge standard can be tricky, it depends on the precision of the model.
* Training sample may include noise so that the gradient may not precise. 

Based on the idea of gradient searching, a variety of algorithm developed targeting to solve the optimization problem in a more practical way.

### Batch Gradient Descent

To suppress the error of the gradient raising from the fluctuation from input samples, the gradient is estimated by the mean of the gradients obtained from each of the input samples. The main purpose of the BGD is to reduce the error of the gradients at a cost of calculating the gradients for every data in the training batch.

### Stochastic Gradient Descent
The stochastic gradient descent (SGD) uses a gradient calculated from a random picked sample form training batch as the estimation for ground truth gradient. The advantage of this method is that it will only calculate the gradient once per each batch (independent on the batch size), it significantly simplified the gradient computation at the cost that the error of the SGD will makes the $\mathcal{L}$ fluctuation during the searching iteration.

Another SGD variantion is specific for the case that $\mathcal{L}$ is a summation over all the samples:
$$
\mathcal{L}=\sum_iL(\boldsymbol{x}_i | \boldsymbol{\theta}),
$$
where $\boldsymbol{x}_i$ stands for the inputs from a samples. Since the similarity, the graident of $\nabla_\theta\mathcal{L}$ is approximated by $\nabla_\theta L$ at a randomly picked sample $\boldsymbol{x}_i$. 

### Momentum Estimation
The algorithm of the momentum estimation updating the parameters based on
$$
\begin{aligned}
\boldsymbol{\theta} &= \boldsymbol{\theta} +\alpha \Delta\boldsymbol{\theta},\\
\Delta\boldsymbol{\theta} &= \Delta\boldsymbol{\theta} - r\nabla_\theta \mathcal{L},
\end{aligned}
$$
where $\Delta\boldsymbol{\theta}$ is 0 at the very begining, and $\alpha$ is a decay parameter in $[0,1]$.
The idea comes from imaging the searching as a point moving in parameter space at a velocity $\Delta\boldsymbol{\theta}$ with a friction decay parameter $\alpha$, while $\mathcal{L}$ is a potential function so that $\nabla_\theta \mathcal{L}$ play the role of force to change the velocity of $\Delta\boldsymbol{\theta}$. This method can supressed the disturbs from gradient fluctuation effectively so that $\mathcal{L}$ may decrease smoothly.

### Adaptive Gradient Descent

The adaptive gradient descent (AdaGD) introduced a variating learning step defined by the inverse of the length of the gradient. The algorithm is similar to SGD except moving parameters as
$$
\boldsymbol{\theta}_{t+1}=\boldsymbol{\theta}_t-\frac{r}{\sqrt{S_{t+1}}}\nabla_\theta\mathcal{L}_t,\quad S_{t+1}=S_t+|\nabla_\theta\mathcal{L}|^2.
$$
The gradient will be shrinked by a factor of $1/\sqrt{S_t}$ at $t$-th step. The $S_t$ increase monotonically which leads to the learning step are decreasing as the iteration number increase. 

### Root Mean Square Propagation

The Root Mean Square Propagation (RMSProp) is a similar algorithm like AdaGD by variating the learning step according to the length of the gradient. Unlike the AdaGD which has a monotonic decrasing learning step, RMSProp adjust the learning rate as:

$$
\boldsymbol{\theta}_{t+1}=\boldsymbol{\theta}_t-\frac{r}{\sqrt{v_{t+1}}+\epsilon}\nabla_\theta\mathcal{L}_t,\quad v_{t+1}=\gamma v_t+(1-\gamma)|\nabla_\theta\mathcal{L}|^2,
$$
where $\gamma\in [0, 1]$ is called "forgetting factor" as the history sum of gradient length will decay by a factor of $\gamma$ for each iteraction. The small $\epsilon$ constant is added here to avoid singularity.

### Adaptive Momentum Estimation
[Reference](https://arxiv.org/pdf/1412.6980.pdf)

The momentum estimation reliably reduces the flucuation by introduced a "inertial" and "friction" for searching. The Adaptive Momentum (Adam) estimation extended the momentum estimation algorithm by introduce an adaptive learning step similar to the RMSProp. It updates the parameters based on the formula:
$$
\begin{aligned}
m_{t+1} &= \beta_1 m_t+(1-\beta_1)\nabla_\theta\mathcal{L},\\
v_{t+1} &= \beta_1 v_t+(1-\beta_2)\Vert\nabla_\theta\mathcal{L}\Vert^2,\\
m &= \frac{m_{t+1}}{1-\beta_1^{t+1}},\\
v &= \frac{v_{t+1}}{1-\beta_2^{t+1}},\\
\boldsymbol{\theta}_{t+1} &= \boldsymbol{\theta}_t-r\frac{m}{\sqrt{v}+\epsilon},
\end{aligned}
$$
where $\beta_1,\beta_2$ is the decay parameter for the learning step and the momentum, respectively. The value range is suppose to be $0\le\beta\le 1$ and the default values are $\beta_1=0.9,\beta_2=0.99$. The 3rd and 4th equations are **bias correction** since a moving average $m = \beta m +(1-\beta)G$ can be expressed as
$$
m = 1-\beta\sum_{i=1}^N\beta^{N-i}G
$$ 
and the expectation of this quantity is
$$
\mathbb{E}(m)= \mathbb{E}(G)(1-\beta^N),
$$
indicates that an unbiased estimator for $G$ is $\hat m = m/(1-beta^N)$.