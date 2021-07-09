# Optimization

Given a system formulated by a function $\mathcal{L}(\boldsymbol{\theta})$, the optimization is to find a point $\boldsymbol{\hat\theta}$ that maximize or minimize the function $\mathcal{L}(\boldsymbol{\theta})$. 

## Gradient Searching
One of the most practical solutions is based on the fact that gradient $\nabla\mathcal{L}$ always pointed to the fast direction that increase the $\mathcal{L}$, and the gradient is 0 at the maximum or minimum points of $\mathcal{L}$. The idea of gradient searching can be illustrated by the gradient descent algorithm (searching for minimum point) as an example:
1. Start at a random point in parameter space $\boldsymbol{\theta}$
2. Update the point following the gradent 
$$
\boldsymbol{\theta} \to \boldsymbol{\theta} - r\nabla\mathcal{L}(\boldsymbol{\theta}),
$$
where $r$ is a learning rate.
3. Repeat the steps above until converge.

In real practice, despite the cost of calculating the gradient is expensive for each step, but also several points need to be aware:
* The gradient usually comes from a grid difference. It is critical to have a good estimation on the gradient.
* The gradient searching may leads to a stable point that $\nabla \mathcal{L}=0$ but it is neither locally maximum nor minimum.
* Choosing the learning rate is subtle task. A larger learning rate may needs fewer iterations to approach the solution, but it is also easier to past the minimum point so that the searching will move around the solution back and forth.
* Converge standard can be tricky, it depends on the precision of the model.
* Training sample may include noise so that the gradient may not precise. 

Based on the idea of gradient searching, a variety of algorithm developed targeting to solve the optimization problem in a more practical way.

### Batch Gradient Descent

To suppress the error of the gradient raising from the fluctuation from input samples, the gradient is estimated by the mean of the gradients obtained from each of the input samples. The main purpose of the BGD is to reduce the error of the gradients at a cost of calculating the gradients for every data in the training batch.

### Stochastic Gradient Descent
The stochastic gradient descent (SGD) uses a gradient calculated from a random picked sample form training batch as the estimation for ground truth gradient. The advantage of this method is that it will only calculate the gradient once per each batch (independent on the batch size), it significantly simplified the gradient computation at the cost that the error of the SGD will makes the $\mathcal{L}$ fluctuation during the searching iteration.

###