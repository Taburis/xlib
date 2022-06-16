# Supervised Learning

Suppose $(\Omega,\mathcal{R}, \mathbb{P})$ is a probability space and $\Omega$ stands for the events of interests. 
* The features are random variables $X:\Omega\to \mathbb{R}^F$ mapping the events to the feature space $\mathbb{R}^F$, where $F$ is called the feature size. The $i$-th component is denoted as $X_i$. 
* Labels are random variables $Y:\Omega\to \mathbb{R}^L$ and $L$ is the dimension of labels. 
* A sample $S$ is a set of pairs $(X_i,Y_i)$, $i=1,\dots, m$ where $m$ is known as sample size.
* A model is a function $f:\mathbb{R}^F\to\mathbb{R}^L$ building connection between the features and labels. The sample $S$ should be, at least, the support of the function $f$.
* Objective function or Loss function $L_S(f)$ is a functional evaluated on the supports $S$.

Building a model in statistical learning is to estimate a function $f$ such that have a best performance (minimizing the loss function). For supervised learning, the labels are given in samples. The performance of the model $f$ with respect to $L_S(f)$ under the support $S$ is defined as
$$
L_S(f)=\int_{S\subset \Omega} \mathcal{L}(Y,f\circ X)(\omega) d\mathbb{P}(\omega),
$$
where $\mathcal{L}$ is the loss measure for the difference between the labels and features. And $(f\circ X)(\omega) =f[X(\omega)]$. This expression is reduced the integration to summation if the sample is discrete. The minimization of $L_S$ leads to the requirement: 
$$
\delta L_S(f)= \int_S\frac{\delta\mathcal{L}}{\delta f}\delta (f\circ X) (\omega)d\mathbb{P}(\omega)=0,
$$
so that the following equation must be satifised
$$
\frac{\delta \mathcal{L}}{\delta f}=0.
$$
The machine learning is all about how to solve this equation for the given loss function $\mathcal{L}$ and the support $S$.

## Variation of Functions

One practical way to implement the variation of functions $\delta f$ is the functional parameterization. The idea is that the function $f(x,\boldsymbol{\theta})$ contained undetermined enough parameters $\boldsymbol{\theta} =(\theta_1,\dots, \theta_m)$ so that the function shape can be manipulated by assigning different parameters to the function. In this method, the structure of $f(x,\boldsymbol{\theta})$ itself set up a limitation for probing. Like a linear function $f(x,\boldsymbol{\theta}) = \sum_ix_i\\theta_i$ can't capture the non-linear behavior no matter how you choose the parameters. The more flexible function you input, the better learning performance you may reach. But the trade off is that flexibility usually means complexity which may cause difficulty on optimizing. For instance, a neural network can be very flexible with a high enough layers and cells, but it leads to great challenge to optimize (more local minimum, long training iterations, and over-fit handling). The prototype function $f(x,\boldsymbol{\theta})$ also know as the "model". So choose your model wisely.

With the parameterization method, the supervised problem reduced to a optimization problem:
$$
\frac{\delta\mathcal{L}}{\delta f}=\frac{\partial\mathcal{L}(f)}{\partial f}\frac{\partial f}{\partial \theta}=0.
$$ 

## Loss Functions

Loss functions are supposed to quantify the difference between the predicted labels $\hat Y = f(X)$ and the truth $Y$. On the other side, it could also be used to evaluate the model to prevent the over fitting. 

### Variance

For continuous labels $Y$, the variance $\text{Var}(Y-f(X))$ is used as the error function:
$$
\text{Var}(Y-\hat Y) = \int_S(Y-\hat Y)^2(\omega)d\mathbb{P}(\omega)
$$