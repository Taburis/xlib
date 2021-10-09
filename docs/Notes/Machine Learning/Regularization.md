# Regularization

A regularization for machine learning is the process of adding information in order to prevent overfitting. 

## Regularization Term

A typical loss function $\mathcal{L}$ usually has two terms
$$
\mathcal{L}=\sigma(x, y|\hat \theta) + \lambda R(f),
$$ 
where $x,y$ stands for the inputs and targets, respectively. The first term $\sigma(x,y|\hat\theta)$ usually measure the variance or the cost of prediction from the model $f$, under the assumption $\theta$ and the second regularization term $R(f)$ represents the penalty from the complexity of the model $f$. A parameter $\lambda$ added here to adjust the weight of the regularization. The philosophy here is that adding extra degree of freedom into the model should only for the case that it leads to significant improve the fitting. But the standard of the "significant improvement" is measured by the regularization term.

### Lp Regularizations

The $Lp$ regulariation term is
$$
R(f) = \sum_i|\theta_i|^p,
$$
where $\theta_i$ are the parameters in the model $f(x,\boldsymbol{\theta})$. For a linear model $f(x)=\theta\cdot x$ where $\theta$ could be a matrix, L1 and L2 regularization terms are mostly used:
* **L0 regularization**: $R(f)=\sum_iI(\theta_i)$, where $I(\theta) = 1$ if $\theta\ne 0$ and $I(\theta)=0$ otherwise.
* **L1 regularization**: $R(f)=\sum_i|\theta_i|$.
* **L2 regularization**: $R(f)=\sum_i\theta_i^2$.

All these regularization term will setup a threshold for judging if adding an extra parameter was worthy. The L0 term penalty for adding extra parameter is a constant, independent on the value of the parameters. This property often leads to a sparse solution (many parameters are 0) minimizing the loss function. 

The penalty from L1 and L2 is depends on the value of the parameters added. The penalty introduced by L1 regularization in parameter space is a diamond-like shapes while it is a sphere introduced by L2. It means that the penalty of parameters increase linearly in L1 while the L2 tolerant more than L1 term for those parameters less than 1. This makes the L1 term derive sparse solution more often than the L2 term for the parameters generally with mean less or around 1. Noticing the performance of the regularization term is highly correlated to the form of the loss term $\sigma(x, y|\hat \theta)$ since it will setup the priority of the prediction costs.

The minimum of loss is reached if 
$$
\nabla\sigma = -\lambda \nabla R,
$$
the gradient of the error term is collinear with but opposite direction to the gradient of regularization term. Notice that the gradient of L1 is always a constant, it is a unit vector pointing to origin where it starts from any axis of weights, and is $\sqrt{n}/n(1, \dots, 1)$ elsewhere. 