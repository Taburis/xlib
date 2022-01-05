# Linear Models

The follow theorems are the mathematical foundation for many models rely on normal distributions. By the CLT, these models are practically useful in many cases. 

## Linear Gauss-Markov Model
---
A **Gauss-Markov Assumption** for a $m$-dimensional random variable $\boldsymbol{Y}\sim N(\boldsymbol{\mu}_Y, I_m\sigma^2)$ can be expressed as a location model $\boldsymbol{Y}=\boldsymbol{\mu}_Y+\boldsymbol{\epsilon}$ with assumption that $\boldsymbol{\epsilon}\sim N(0,I_m\sigma^2)$. The linear model further assume that $\boldsymbol{\mu_Y}=X\boldsymbol{\theta}$ so that we have
$$
\boldsymbol{y}=X\boldsymbol{\theta}+\epsilon,\quad \epsilon\sim N(0,I_m\sigma).
$$

The linear model under Gauss-Markov assumption is **estimable** for $\boldsymbol{\theta}$ only if the column space of $X$ samples is full ranked. The estimable of $\boldsymbol{\theta}$ is that if $\theta_1\ne \theta_2$, then $f(\theta_1)\ne f(\theta_2)$. This property of $f(\theta)$ is also known as **identifiable**. 
Given two $\boldsymbol{\theta}_1\ne\boldsymbol{\theta}_2$, non-trivial solution of $X(\boldsymbol{\theta}_1-\boldsymbol{\theta}_2)=0$ implies that column space of $X$ is not full ranked and $\boldsymbol{\theta}_1-\boldsymbol{\theta}_2$ are in the null space of $X$. So linear model is estimable only for those $\boldsymbol{\theta}\notin \text{Null}(X)$.

To quantify the quality of the regression, let $\hat\boldsymbol{y}$ be an estimator for $\boldsymbol{y}$, then we have the following term:
* The Sum of Square of Total (SST): $\Vert\boldsymbol{y}-\bar y\Vert^2 = \sum_i(y_i-\bar y)^2$.
* The Sum of Square of Regression (SSR): $\Vert\boldsymbol{\hat y}-\bar y\Vert^2=\sum_i(\hat y_i-\bar y)^2$.
* The Sum of Square of Error (SSE): $\Vert\boldsymbol{y}-\boldsymbol{\hat y}\Vert^2=\sum_i(y_i-\hat y_i)^2$.

### The Best Linear Unbiased Estimator

A Best Linear Unbiased Estimator (BLUE) of $\boldsymbol{y}$ is a linear unbiased estimator that minimize the SSE. Furthermore, we have:
1. The BLUE of $\boldsymbol{y}$ for the linear model above is $X\boldsymbol{\hat \theta}_{\text{LSE}}$ where $\boldsymbol{\hat \theta}_{\text{LSE}}$ **least square estimator** for $\theta$:
$$
\begin{aligned}
\boldsymbol{\hat \theta}_{\text{LSE}} & = (X^TX)^{-1}X^T\boldsymbol{y},\\
\boldsymbol{\hat y}_{\text{BLUE}}&=X\boldsymbol{\hat \theta}_{\text{LSE}}=P_X\boldsymbol{y},
\end{aligned}
$$
where $P_X=XX^+$ is known as a projection matrix (symmetric and idempotent) and $X^+=(X^TX)^{-1}X^T$ is the Moore-Penrose inverse of $X$.
2. The (SSE) denoted as $R_0^2$ satisfies $R^2_0/\sigma^2\sim\chi^2(m-r)$ where $r$ is the rank of $X$. And $\hat\Sigma_y=R^2_0/(m-r)$ is a unbiased estimator for the $\sigma^2$.
3. Suppose there's a constant term such that $\boldsymbol{y}=X\boldsymbol{\theta}+\beta\cdot \boldsymbol{1}_n$, then the equation can be reduced to $\boldsymbol{y}_0=X_0\boldsymbol{\theta}$ where $\boldsymbol{y}_0=\boldsymbol{y}-\overline{y}$ and $X_0=X-\overline{X}$.
4. $\text{SSR}\perp\text{SSE}$ and $\text{SST}=\text{SSR}+\text{SSE}$.


**Proof**: 
1. It is unbiased since $\mathbb{E}(XX^+\boldsymbol{y})=XX^+X\boldsymbol{\theta}=X\boldsymbol{\theta} = \boldsymbol{y}$ where the last step comes from the property of Moore-Penrose inverse $A=AA^+A$. 
It has the least variance: Suppose $A\boldsymbol{y}$ is another unbiased estimator for $\boldsymbol{y}$, then
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
which leads the **normal equation**: 
$$
X^TX\boldsymbol{\theta} = X^T\boldsymbol{y},
$$
where the following equations are used
$$
\frac{\partial \boldsymbol{y}^T\boldsymbol{x}}{\partial\boldsymbol{x}}=\boldsymbol{y},\quad
\frac{\partial \boldsymbol{y}^T\boldsymbol{x}}{\partial\boldsymbol{x}^T}=\boldsymbol{y}^T,\quad 
\frac{\partial \boldsymbol{x}^T}{\partial\boldsymbol{x}}=I_n,\quad \frac{\partial \boldsymbol{x}^TA\boldsymbol{x}}{\partial\boldsymbol{x}}=(A+A^T)\boldsymbol{x}.
$$
2. Due to $P_X$ is the projection operator, then $I_m-P_X$ is symmetric and idempotent. Furthermore, $I_n=P_X+(I_m-P_X)$ and hence the Cochran's theorem shows that $\Vert P_X\boldsymbol{y}\Vert^2/\sigma^2\sim \chi^2(r)$ while $R_0^2=\Vert (I-P_X)\boldsymbol{y}\Vert^2/\sigma^2\sim \chi^2(m-r)$. 
3. To minimize the error, we have
$$
\frac{\partial}{\partial \boldsymbol{\theta}} (\boldsymbol{y}-X\boldsymbol{\theta}-\beta\cdot\boldsymbol{1}_n)^T(\boldsymbol{y}-X\boldsymbol{\theta}-\beta\cdot\boldsymbol{1}_n)=0,
$$
which leads to
$$
\begin{aligned}
\beta &= \frac{1}{n}\left(\boldsymbol{1}^T_n\boldsymbol{y}-\boldsymbol{1}^T_nX\boldsymbol{\theta}\right)\\
& = \overline{y}-\overline{X}\boldsymbol{\theta}.
\end{aligned}
$$
Put this back to the formular leads to the reduction.
4. By the 3rd point, the model estimator is $\hat\boldsymbol{y}-\bar\boldsymbol{y}=P_X(\boldsymbol{y}-\bar\boldsymbol{y})$ where $X$ is the $X_0$ in the 3rd point. The explicity expand SST shows that
$$
\begin{aligned}
\Vert \boldsymbol{y}-\bar \boldsymbol{y}\Vert^2&= \Vert (\boldsymbol{y}-\bar\boldsymbol{y})-(\hat \boldsymbol{y}-\bar\boldsymbol{y})+\hat\boldsymbol{y}-\bar y\Vert^2\\
&=\Vert \hat \boldsymbol{y}-\bar \boldsymbol{y}\Vert^2+\Vert(\boldsymbol{y}-\bar\boldsymbol{y})-(\hat \boldsymbol{y}-\bar\boldsymbol{y})\Vert^2+2[(\boldsymbol{y}-\bar\boldsymbol{y})-(\hat \boldsymbol{y}-\bar\boldsymbol{y})]^T(\hat \boldsymbol{y}-\bar\boldsymbol{y}),\\
&=\Vert \hat \boldsymbol{y}-\bar \boldsymbol{y}\Vert^2+\Vert\boldsymbol{y}-\hat \boldsymbol{y}\Vert^2+2(\boldsymbol{y}-\bar\boldsymbol{y})^T(I-P_X)P_X(\boldsymbol{y}-\bar\boldsymbol{y}).
\end{aligned}
$$
where the last term vanishes. It means $(\boldsymbol{y}-\bar \boldsymbol{y})\perp (\boldsymbol{y}-\hat\boldsymbol{y})$.

:::info
A explicit solution: $X^+X=X(X^TX)^-X^T$ where $X^-$ is a general inverse of $X$. But in practice, it is usually approximated by a gradient descent search to minimize the residual.
:::



It because the $(I_m-P_X)$ projects vector to the space perpendicular to the column space of $X$, and it is symmetry and idempotent $(I_m-P_X)^2=(I_m-P_X)$, it derive the first expression immediately. To get the expectation, we have
$$
\begin{aligned}
\mathbb{E}(R^2_0)&=\mathbb{E}[Y^T(I_m-](P_x)Y],\\
&=\mathbb{E}[(\boldsymbol{y}-X\boldsymbol{\theta})^T(I_m-P_x)(\boldsymbol{y}-X\boldsymbol{\theta})],\\
&=\text{tr}(A\Sigma),
\end{aligned}
$$
where we used the properties that $\boldsymbol{y}^TA\boldsymbol{y} = \text{tr}(\boldsymbol{y}^TA\boldsymbol{y})$, $(I_m-P_X)X\theta=0$ and the relation:
$$
\mathbb{E}(\boldsymbol{x}^TA\boldsymbol{x}) = \text{tr}{A\Sigma_x}+\boldsymbol{\mu}^T_xA\boldsymbol{\mu}_x,
$$
and the trace of $P_X$ is the rank of $X$. 

### Linear Estimation and Test

Given a linear model with weight $\boldsymbol{\theta}$, an arbitrary linear estimator of $H\boldsymbol{\theta}$ is estimable if
1. $\mathcal{C}(H^T)\subset \mathcal{C}(X^T)$, where $\mathcal{C}(X^T)$ stands for the column space of $X^T$.
2. The rank of the column space of $H$ is the same of the rank of $H$. 

The reason for the linear estimable conditions above is that the estimator $\boldsymbol{\hat\theta}=A\boldsymbol{y}$, a linear estimation from $\boldsymbol{y}$. Then the estimable means that $\forall \boldsymbol{\theta},\exists B: H\boldsymbol{\theta}=B\boldsymbol{y}$ which means that $\mathcal{C}(H^T)\subset \mathcal{C}(X^T)$. The second point is a minimum requirement for $H$ since if $H$ is a $k\times m$ matrix while $k\le \text{rank}{H}$, then $H$ can be reduced to a smaller matrix which is full rank of column.

For a estimable linear combination $H\boldsymbol{\theta}$, $H\in \mathbb{R}^{k\times h}$ where $h$ is the dimension of the $\boldsymbol{\theta}$, one can form a hypothesis test:
$$
H_0: H\boldsymbol{\theta} = \boldsymbol{a} \quad H_1: H\boldsymbol{\theta}\ne \boldsymbol{a}.
$$
We have the following properties:
1. $R_0^2\perp R^2_1-R^2_0$ where 
$$
R_1^2=\min_{\lbrace\boldsymbol{\theta}: H\boldsymbol{\theta}=\boldsymbol{a}\rbrace} \Vert \boldsymbol{y}-X\boldsymbol{\theta}\Vert ^2
$$
2. For different hypothesis assumption, we have:
$$
H_0: \frac{R_1^2-R_0^2}{\sigma^2}\sim \chi^2(k),\quad H_1:  \frac{R_1^2-R_0^2}{\sigma^2}\sim \chi^2_\delta(k),
$$
where $\chi_\delta^2(k)$ is a non-central chi square distribution with 
$$
\delta = \frac{1}{\sigma^2}(H\boldsymbol{\theta}-\boldsymbol{a})^TV^{-1}(H\boldsymbol{\theta}-\boldsymbol{a}),\quad V=H(X^TX)^-H^T.
$$
3. $F$-statistics for hypothesis
$$
F = \frac{U/k}{R^2_0/(m-r)},\quad U=(H\boldsymbol{\hat\theta}-\boldsymbol{a})^TV^{-1}(H\boldsymbol{\hat\theta}-\boldsymbol{a}),\quad U/\sigma^2\sim \chi^2(k)
$$
which satisfies the distribution:
$$
H_0: F\sim F(k,m-r),\quad H_1:F\sim F_\delta(k,m-r),
$$
where $k=\text{rank}{H}$.

With this hypothesis test method, one can use this to test the importance of the $i$-th component in $\boldsymbol{\theta}$ by setting $H\boldsymbol{\theta}=\theta_i$ with the associated $H_0:H\boldsymbol{\theta}=0$.

:::note Mathematical Foundation
With these requirement above, the estimable linear combination of BLUE $\boldsymbol{\hat \theta}$ has the form $H\boldsymbol{\hat\theta}=TX\boldsymbol{\hat\theta}$ where $T\in\mathbb{R}^{k\times m}$ hence
$$
H\boldsymbol{\hat\theta}\sim N_k(H\boldsymbol{\theta}, HH^T\sigma^2),\quad V^{-1/2}H(\boldsymbol{\hat\theta}-\boldsymbol{\theta})/\sigma\sim N_k(0, I_k),
$$ 
where $V=H(X^TX)^-H^T$ because $H\boldsymbol{\hat \theta}=H(X^TX)^-X^T\boldsymbol{y}$. It means that
$$
U\sim \chi^2(k), \quad U = \frac{1}{\sigma^2}(H\boldsymbol{\hat\theta}-H\boldsymbol{\theta})^TV^{-1}(H\boldsymbol{\hat\theta}-H\boldsymbol{\theta}).
$$
Notice that $U\perp R^2_0$ since 

To show the $R^2_0\perp R^2_1-R^2_0$, notice that the $R^2_0=\Vert \boldsymbol{y}-X\boldsymbol{\hat\theta}\Vert^2=\Vert (I-P_X)\boldsymbol{y}\Vert^2$, then
$$
\begin{aligned}
R^2_1&=\min_{\lbrace\boldsymbol{\theta}: H\boldsymbol{\theta}=\boldsymbol{a}\rbrace}\Vert \boldsymbol{y}-X\boldsymbol{\hat \theta}+X\boldsymbol{\hat \theta}-X\boldsymbol{\theta}\Vert^2\\
&=\Vert \boldsymbol{y}-X\boldsymbol{\hat \theta}\Vert^2+\min_{\lbrace\boldsymbol{\theta}: H\boldsymbol{\theta}=\boldsymbol{a}\rbrace}\Vert X\boldsymbol{\hat \theta}-X\boldsymbol{\theta}\Vert^2
\end{aligned}
$$
The last step comes from that $(\boldsymbol{y}-X\boldsymbol{\hat \theta})^TX(\boldsymbol{\hat\theta}-\boldsymbol{\theta})=0$ which means that $R^2_0\perp R^2_1-R^2_0$. Further more, we have $P_X\boldsymbol{y}-X\boldsymbol{\theta} = P_X(\boldsymbol{y}-X\boldsymbol{\theta})$, then $R^2_1-R^2_0=\Vert P_X(\boldsymbol{y}-X\boldsymbol{\theta})\Vert^2$. By the 3rd theorem in the [section](/Statistics/Special-Distributions/Normal-Distributions#multivariate-normal-distribution), $R^2_1-R^2_0\sim \chi^2(k)$ under the $H_0$. Otherwise, it is a skewed chi2 distribution. This proved the 2nd point. Actually, we can prove that this $R^2_1-R^2_0\perp R^2_0$ for arbitrary $\theta$. In fact, $H\boldsymbol{\theta}=TX\boldsymbol{\theta}=TP_XX\boldsymbol{\theta}$, then we have
$$
\begin{aligned}
H\boldsymbol{\hat \theta}-H\boldsymbol{\theta}&=H(X^TX)^-X^T(\boldsymbol{y}-X\boldsymbol{\theta})\\
&=TP_X(\boldsymbol{y}-X\boldsymbol{\theta}).
\end{aligned}
$$
Based on this, we have
$$
U=(\boldsymbol{y}-X\boldsymbol{\theta})^TA(\boldsymbol{y}-X\boldsymbol{\theta}),\quad A=P_XT^T(TP_XT^T)^{-1}TP_X,
$$
where we used the equation $H(X^TX)^-H^T=TP_XT^T$. Furthermore, $A(I-P_X)=0$ which means that $U\perp R^2_0$. Then $U/\sigma^2\sim \chi^2(k)$ and $R^2_0/\sigma^2\sim \chi^2(m-r)$ follows from the Cochram's theorem. The 3rd point follows from the definition of the $F$-distribution.
:::

### One-Way ANOVA

Analysis of Variance (ANOVA) is a method to analyze the mean between the data groups. Given a sequence of i.i.d. random variables $X_j,j=1,\dots,b$ $X_i\sim N(\mu_i,\sigma^2)$ and the samples draw for $X_j$ are denoted as $X_{ij}$ where $i=1,\dots, a$. The model for these samples are
$$
X_{ij}=\mu_j+e_{ij},\quad e\sim N(0,\sigma^2).
$$ 
This model isolated the single factor $\mu$ from variance $e_{ij}$ for data, and is called the $b$-level (the sample size for each group) ANOVA model. Here list several usual inference about this model:
1. The BLUE $\hat\mu_j=\bar y_{.j}$ where $\bar y_{.j}=\sum_{i=1}^by_{ij}/a$.
2. Let $W_j = \bar X_{.j}-\mu_j$, then $W_j\sim N(0,\sigma^2/a)$ and
$$
t_j \sim t(ab-b),\quad t_j=\frac{\sqrt{a}W_j}{\sqrt{R^2_0/(ab-b)}},
$$
where $R^2_0 /\sigma^2= \sum_{ij}(X_{ij}-\bar X_{.j})^2/\sigma^2\sim \chi^2(ab-b)$.
3. Mean test hypothesis:
$$
H_0: \mu_1=\dots = \mu_b = \mu, \quad H_1: \mu_i\ne \mu
$$ 
the $F$-statistics can be used to test the hypothesis:
$$
\frac{(R^2_1-R^2_0)/(b-1)}{R^2_0/(ab-b)}\sim F(b-1, ab-b),
$$
where $R^2_1-R^2_0=a\sum_{j}(\bar X_{.j}-\bar X_{..})^2$ where $\bar X_{..}=\sum_{ij}X_{ij}/ab$.


:::note Mathematical Foundation
1. The ANOVA model is a special case of linear regression by flatten the $y_{ij}$ into $y_i$ where $i=1,\dots, bm$. So that we have
$$
X=\begin{pmatrix}
1_{a\times 1} &&\\
&\ddots & \\
&& 1_{a \times 1}
\end{pmatrix},\quad \boldsymbol{\theta}=\begin{pmatrix}\mu_1 \\ \vdots \\ \mu_b \end{pmatrix}
$$
where $1_{a \times 1}$ is a $a\times 1$ matrix ($a$ dimensional column vector) with all elements are 1, the subscripts here refers the row$\times$\column number, and the elements absence in the matrix above are 0. Notice that $1_{1\times a}\cdot 1_{a\times 1} = a$. The BLUE comes from the formula
$$
\begin{aligned}
\hat \boldsymbol{\theta}&= (X^TX)^{-1}X^T\boldsymbol{y}\\
&=\frac{1}{b}
\begin{pmatrix}
1_{1\times a} &&\\
&\ddots & \\
&& 1_{1 \times a}
\end{pmatrix}\cdot
\begin{pmatrix}
y_{1}\\ \vdots \\ y_{ab} 
\end{pmatrix}
\\
&=\begin{pmatrix}
\bar y_{.1} & \dots \bar y_{.b}
\end{pmatrix}^T,
\end{aligned}
$$
here it is easy to verify $(X^TX)^{-1}=I_{b\times b}/b$. 
2. $R^2_0\sim \chi(ab-b)$ since $\text{rank}(X)=b$. The rest follows from student's theorem.
3. The $H_0$ hypothesis can be rephrased into the test that $\mu_1-\mu_j=0$ for all $j=2,\dots,b$. It hence can be expressed as $H\boldsymbol{\mu}=0$ where 
$$
H=\begin{pmatrix} 
1 & -1 & & &\\
1 &  &-1 & &\\
\vdots & & & \ddots &\\
1 &  & & &-1
\end{pmatrix}_{(b-1)\times b},\quad \text{rank}(H)=b-1.
$$
The empty place is 0 in the matrix above. Assuming the $H_0$ is true, the best estimation of mean is $\mu=\bar X_{..}=\sum_{ij}X_{ij}/ab$. So that $U=R^2_1-R^2_0$ and $U/\sigma^2\sim\chi^2(b-1)$. Since $R^2_0\perp (R^2_1-R^2_0)$, it leads to the $F$-distributions of the ratio.
:::