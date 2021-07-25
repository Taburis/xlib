# Linear Model

## Gauss-Markov Model

A **Gauss-Markov Assumption** for a $m$-dimensional random variable $\boldsymbol{y}$ is that $\boldsymbol{y}\sim N(X\boldsymbol{\theta},I_m\sigma^2)$: assuming the mean of the $y$ can be expressed as a linear combination of a $n$-dimension vector $\boldsymbol{\theta}$ with a $m\times n$ matrix $X$ as the function
$$
\boldsymbol{y}=X\boldsymbol{\theta}+\epsilon,\quad \epsilon\sim N(0,I_m\sigma).
$$

### Identifiable
The linear model under Gauss-Markov assumption is estimable for $\boldsymbol{\theta}$ only if the column space of $X$ samples is full ranked. The estimable of $\boldsymbol{\theta}$ is that if $\theta_1\ne \theta_2$, then $f(\theta_1)\ne f(\theta_2)$. This property of $f(\theta)$ is also known as identifiable. 

Given two $\boldsymbol{\theta}_1\ne\boldsymbol{\theta}_2$, non-trivial soluiton of $X(\boldsymbol{\theta}_1-\boldsymbol{\theta}_2)=0$ implise that column splace of $X$ is not full ranked and $\boldsymbol{\theta}_1-\boldsymbol{\theta}_2$ are in the null space of $X$. 

### The Best Linear Unbiased Estimator
The Best Linear Unbiased Estimator (BLUE) of $\boldsymbol{y}$ is the linear combination of $X\boldsymbol{\hat \theta}_{\text{LSE}}$ where $$ *\boldsymbol{\hat \theta}_{\text{LSE}}$*least square estimator** $\boldsymbol{\hat \theta}_{\text{LSE}}$.
$$
\boldsymbol{\hat y}_{\text{BLUE}}=X\boldsymbol{\hat \theta}_{\text{LSE}}=P_X\boldsymbol{y}=XX^+\boldsymbol{y},
$$
where $X^+$ is the Moore-Penrose inverse of $X$ and $P_X=XX^+$ is known as a projection matrix.

It is unbiased since $\mathbb{E}(XX^+\boldsymbol{y})=XX^+X\boldsymbol{\theta}=X\boldsymbol{\theta} = \boldsymbol{y}$ where the property of Moore-Penrose inverse $A=AA^+A$ used. 

And it has the least variance. Suppose $A\boldsymbol{y}$ is another unbiased estimator for $\boldsymbol{y}$, then
$$
\begin{aligned}
\text{Var}(A\boldsymbol{y}) &=\text{Var}(A\boldsymbol{y}+P_X\boldsymbol{y}-P_X\boldsymbol{y}),\\
& = \text{Var}[(A-P_X)\boldsymbol{y}]+\text{Var}(P_X\boldsymbol{y})+\text{Cov}[(A-P_X)\boldsymbol{y},P_X\boldsymbol{y}],\\
& \ge \text{Var}(P_X\boldsymbol{y}),
\end{aligned}
$$
since the last covariance matrix is vanish due to
$$
\text{Cov}[(B-P_A)\boldsymbol{y},P_A\boldsymbol{y}]=(B-P_A)B\text{Var}(y),
$$
and $(A-P_X)\oplus P_X=A$.

The $\boldsymbol{\hat \theta}_{\text{LSE}}$ is defined by the least square method:
$$
\frac{\partial}{\partial \boldsymbol{\theta}} (\boldsymbol{y}-X\boldsymbol{\theta})^T(\boldsymbol{y}-X\boldsymbol{\theta})=0,
$$
which implise the normal equation 
$$
X^TX\boldsymbol{\theta} = X^T\boldsymbol{y},
$$
where the following equations are used
$$
\frac{\partial \boldsymbol{y}^T\boldsymbol{x}}{\partial\boldsymbol{x}}=\boldsymbol{y}^T,\quad \frac{\partial \boldsymbol{x}^T}{\partial\boldsymbol{x}}=I_n,\quad \frac{\partial \boldsymbol{x}^TA\boldsymbol{x}}{\partial\boldsymbol{x}}=(A+A^T)\boldsymbol{x}.
$$

:::info
A explicity solution $X^+X=X(X^TX)^-X^T$ where $X^-$ is a general inverse of $X$. 
:::

### Residual and Uncertainties
* Residual of the BLUE is defined as $R_0^2 = \Vert \boldsymbol{y}-\boldsymbol{\hat y}_{\text{BLUE}}\Vert= \Vert (I_m-P_X)\boldsymbol{y}\Vert$ and $\mathbb{E}(R_0^2)=\sigma^2/(m-r)$ where $r$ is the rank of $X$. 
* The Unbiased estimator of $\hat\Sigma_y=R^2_0/(m-r)$.

It because the $(I_m-P_X)$ projects vector to the space perpendicular to the column space of $X$, and it is symmetry and idenpotent $(I_m-P_X)^2=(I_m-P_X)$, it derive the first expression immediately. To get the expectation, we have
$$
\begin{aligned}
\mathbb{E}(R^2_0)&=\mathbb{E}[Y^T(I_m-P_x)Y],\\
&=\mathbb{E}[(\boldsymbol{y}-X\boldsymbol{\theta})^T(I_m-P_x)(\boldsymbol{y}-X\boldsymbol{\theta})],\\
&=\text{tr}(A\Sigma),
\end{aligned}
$$
where we used the properties that $\boldsymbol{y}^TA\boldsymbol{y} = \text{tr}(\boldsymbol{y}^TA\boldsymbol{y})$, $(I_m-P_X)X\theta=0$ and the relation:
$$
\mathbb{E}(\boldsymbol{x}^TA\boldsymbol{x}) = \text{tr}{A\Sigma_x}+\boldsymbol{\mu}^T_xA\boldsymbol{\mu}_x,
$$
and the trace of $P_X$ is the rank of $X$. 





