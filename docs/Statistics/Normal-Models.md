# Normal Models

The follow theorems are the mathematical foundation for many models rely on normal distributions. By the CLT, these models are practically useful in many cases. 

## Mathematical Foundation
---
:::note Theorems:
Given a random vector $\boldsymbol{X}=(X_1,\dots, X_n)$, then
1. Let $\boldsymbol{X}\sim N_n(0,I_n\sigma^2)$, and $Q=\boldsymbol{X}^TA\boldsymbol{X}/\sigma^2$ where $A$ is a symmetric real matrix. Then the moment generate function
$$
M(t)=\prod_{i=1}^r(1-2t\lambda_i)^{-\frac{1}{2}}=\det(1-2tA)^{-\frac{1}{2}},\quad r=\text{rank}(A).
$$
where $\lambda_i$ is the $i$-th eigenvalue of $A$. Furthermore, $Q\sim\chi^2(r)$ if and only if $A^2=A$ (idempotent).
2. Suppose $\boldsymbol{X}\sim N_n(\boldsymbol{\mu},\Sigma)$, and $Q=(\boldsymbol{X}-\boldsymbol{\mu})^T\Sigma^{-1}(\boldsymbol{X}-\boldsymbol{\mu})$, then $Q\sim \chi^2(n)$.
3. Suppose $\boldsymbol{X}\sim N_n(\boldsymbol{\mu},\Sigma)$ and $Q=(\boldsymbol{X}-\boldsymbol{\mu})^TA(\boldsymbol{X}-\boldsymbol{\mu})$ where $A$ is symmetry. Then $Q\sim\chi^2(r)$ where $r=\text{rank}(A)$ if and only if $A\Sigma A=A$. 
4. Given $\boldsymbol{X}\sim N_n(0,I_n\sigma^2)$, and define the quadratic form $Q_A=\sigma^{-2}\boldsymbol{X}^TA\boldsymbol{X}$, then $Q_A\perp Q_B$ if and only if $AB=0$.
5. **Hogg and Craig's theorem**: Given $Q=\sum_{i=1}^kQ_i$ where $Q,Q_1,\dots,Q_k$ are quadratic form of observations obey $N_n(0,\sigma^2)$, and $Q/\sigma^2\sim\chi^2(n)$ and $Q_i/\sigma^2\sim\chi^2(r_i)$ for $i\le k-1$ and $Q_k$ is non-negative. Then $Q_k/\sigma^2\sim\chi^2(r_k)$ where $n=\sum_{i=1}^kr_k$.  
6. **Cochran's theorem**: Let $\boldsymbol{X}\sim N(0,I_n\sigma^2)$ and 
$$
\sum_{i=1}^nX_i^2=\sum_{j=1}^kQ_j,
$$
where $Q_j=\boldsymbol{X}^TA_j\boldsymbol{X}$ for symmetric $A_j$ with $\text{rank}(A_j)=r_j$. Then $Q_j$ are mutually independent and $Q_j/\sigma^2\sim\chi^2(r_j)$ if and only if $\sum_jr_j=n$.
:::

**Proof**: 
1. Since $A$ is symmetry, there's a spectra decomposition of $A=\Gamma \Lambda \Gamma^T$ where $\Lambda$ is diagonal matrix with $\Lambda_{ii}=\lambda_i$ and $\Gamma\Gamma^T=1$. Then the diagonalization $\boldsymbol{Z}=\Gamma\boldsymbol{X}$ leads to $Z\sim N_n(0,I_n\sigma^2)$ and $Q=\sum_i\lambda_iZ_i^2$. The moment generate function is
$$
\begin{aligned}
M_X(t) = M_Z(t)&=\mathbb{E}\left\lbrace e^{\sum_{i=1}^r\lambda_i Z_i}\right\rbrace\\
&=\prod_{i=1}^r\mathbb{E}(e^{t\lambda_iZ_i})\\
&=\prod_{i=1}^r(1-2t\lambda_i)^{-\frac{1}{2}}\\
&=\det(I_n-2tA).
\end{aligned}
$$
Furthermore, if $A^2=A$, then the eigenvalues of $A$ is $\lambda_i=1$ for $i=1,\dots, r$ where $r=\text{rank}(A)$. Put this but to the moment generate function above leads to the exact formula of moment generate function of $\chi^2(r)$ distribution. 
2. Let $\boldsymbol{Y}=\boldsymbol{X}-\boldsymbol{\mu}$, and a spectra decompose exists that $\Sigma=\Gamma^T\Lambda\Gamma$ so that $\boldsymbol{Z}=\Gamma\boldsymbol{Y}\Sigma^{-1/2}$ satisfies $\boldsymbol{Z}\sim N_n(0,I_n)$ since $\Sigma^{-1}=(\Gamma^T\Lambda\Gamma)^{-1}=\Gamma^T\Lambda^{-1}\Gamma$ and $\Gamma$ is a orthogonal matrix. On the other hand, $Q=\boldsymbol{Y}^T\Sigma^{-1}\boldsymbol{Y}=\boldsymbol{Z}^T\boldsymbol{Z}$ so that $Q\sim\chi^2(n)$.
3. Define $Z=\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})$ then $Z\sim N_n(0,I_n)$. On the other hand, $Q=[\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})]^T\Sigma^{1/2}A\Sigma^{1/2}[\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})]$. Based on the 1st theorem, $Q\sim\chi^2(r)$ if and only if $\Sigma^{1/2}A\Sigma^{1/2}$ is idempotent which require $A\Sigma A=A$.
4. To show the sufficiency, we assume $AB=0$ and let $\Gamma_A$ be a $r\times n$, where $r=\text{rank}(A)$, formed by non-trivial eigenvectors of $A$ normalized such that $\Gamma_A\Gamma_A^T=I_r$. So is $\Gamma_B$ of $B$. Then $A=\Gamma_A^T\Lambda_A\Gamma_A$ where $\Lambda_A$ is diagnol matrix with $r$ non-vanished elements. Let's do the transform $\boldsymbol{W}=\Gamma\boldsymbol{X}$ such that $W\sim N_n(0, \Sigma)$ where
$$
\Gamma=\begin{pmatrix}
\Gamma_A\\
\Gamma_B
\end{pmatrix},\quad \Sigma = \begin{pmatrix}
I_r & \Gamma_A^T\Gamma_B\\
\Gamma_B^T\Gamma_A & I_s
\end{pmatrix}.
$$
Furthermore, $AB=(\Gamma_A^T\Lambda_A)\Gamma_A\Gamma_B^T(\Lambda_B\Gamma_B) = 0$ and $\Lambda_A\Gamma_AAB\Gamma_B\Lambda_B=0 =\Lambda_A^2\Gamma_A\Gamma_B^T\Lambda_B^2$ while the $\Lambda_A^2$ and $\Lambda_B^2$ are full rank which implies that $\Gamma_A\Gamma_B^T=0$ which implies that $Q_A\perp Q_B$.
To prove the necessarity, assuming the $Q_A\perp Q_B$ then
$$
\begin{aligned}
\mathbb{E}\left(e^{Q_At_1+Q_Bt_2}\right)&=\mathbb{E}\left(e^{Q_At_1}\right)\mathbb{E}\left(e^{Q_Bt_2}\right)\\
&=\det(1-2t_1A)^{-\frac{1}{2}}\det(1-2t_2B)^{-\frac{1}{2}}\\
&=\det(1-2t_1A-2t_2B)^{-\frac{1}{2}}
\end{aligned}
$$
and adapting a transform diagonalize $A=\Gamma\Lambda\Gamma$, we have the equation
$$
\left\lbrace \prod_{i=1}^r (1-2t_1\lambda_i)\right\rbrace \det(1-2t_2B)=\det(1-2t_1\Lambda-2t_2D),
$$
and $r=\text{rank}(A)$, $D=\Gamma^TB\Gamma$. This equation implies that $\Lambda D=0=AB$ because $\det(AB)=\det(A)\det(B)$. 
5. Suppose $\boldsymbol{X}\sim N_n(0,I_n\sigma^2)$ and $Q=\boldsymbol{X}^TA\boldsymbol{X}, Q_i=\boldsymbol{X}^TA_i\boldsymbol{X}$, where $A$ and $A_i$ are idempotent for $i\ne k$. Then $A^2=A$ and $A^2_1=A_1$ by the 2nd theorem. Since $A$ and $A_1$ are symmetric and positive semi-defined, there's a transform $\Gamma$ such that diagonalized $A,A_1$ simutaneosly. And since $A_2$ is positive semi-defined as well, it has to be
$$
A_1=\begin{pmatrix}
\alpha & \\
 & 0
\end{pmatrix},\quad 
A_2=\begin{pmatrix}\beta & \\
 & 0
\end{pmatrix},\quad A=\begin{pmatrix}I_r & \\
 & 0
\end{pmatrix}.
$$
Now $A_1A=A_1$ based on the form above and this implies that $A_1A_2=0$ which means that $Q_1\perp Q_2$ which means that $Q_2\sim\chi^2(n-r_1)$ and $A_2^2=A_2$. 
For general case, $A=A_1+B$ where $B=\sum_{i=2}^kA_i$ and we have $A_1B=0$ and $B^2=B$ and $Q_B\sim\chi^2(n-r_1)$ from the $k=2$ case we discussed above. This procedure can be repeated until $k=2$. 
6. The necessarity is obviously that $Q_i/\sigma^2\sim\chi^2(r_i)$ and $Q_i\perp Q_j$ implies that $\sum_iQ_i/\sigma^2\sim \chi^2(\sum_ir_i)$ which is $n=\sum_ir_i$ since $\sum_iX^2_i/\sigma^2\sim\chi^2(n)$. 
To show the sufficiency, we first show that $A_j$ is idempotent if $I_n=\sum_jA_j$ and $\sum_jr_j=n$. In fact, for $k=2$, $I=A+B$, by spectra decomposition
$$
I=\Gamma^T(\Lambda_A+D)\Gamma,
$$
where $D=\Gamma B\Gamma^T$ and $\Lambda_A$ can exactly $r_A$ non-vanishing diagonal elements. It shows that $D$ has to be non-vanished in the complement of the column space of $A$ and hence $A$ and $B$ are idempotent. For general case, it can be reduced to $I=A+B$ and with $B=C+D$. But since $B$ is idmpotent, the similar procedure carried in $B$'s invariant subspace will leads to the same result that $C,D$ are idempotent and so on so forth.
This shows that $A_iA_j=0$ if $i\ne j$ and hence $Q_i\perp Q_j$ and hence $Q_j/\sigma^2\sim \chi^2(r_j)$.

## Gauss-Markov Model
---
A **Gauss-Markov Assumption** for a $m$-dimensional random variable $\boldsymbol{y}$ is that $\boldsymbol{y}\sim N(X\boldsymbol{\theta},I_m\sigma^2)$: assuming the mean of the $y$ can be expressed as a linear combination of a $n$-dimension vector $\boldsymbol{\theta}$ with a $m\times n$ matrix $X$ as the function
$$
\boldsymbol{y}=X\boldsymbol{\theta}+\epsilon,\quad \epsilon\sim N(0,I_m\sigma).
$$

The linear model under Gauss-Markov assumption is **estimable** for $\boldsymbol{\theta}$ only if the column space of $X$ samples is full ranked. The estimable of $\boldsymbol{\theta}$ is that if $\theta_1\ne \theta_2$, then $f(\theta_1)\ne f(\theta_2)$. This property of $f(\theta)$ is also known as **identifiable**. 
Given two $\boldsymbol{\theta}_1\ne\boldsymbol{\theta}_2$, non-trivial solution of $X(\boldsymbol{\theta}_1-\boldsymbol{\theta}_2)=0$ implies that column space of $X$ is not full ranked and $\boldsymbol{\theta}_1-\boldsymbol{\theta}_2$ are in the null space of $X$. So linear model is estimable only for those $\boldsymbol{\theta}\notin \text{Null}(X)$.

### The Best Linear Unbiased Estimator
* The Best Linear Unbiased Estimator (BLUE) of $\boldsymbol{y}$ for the linear model above is $X\boldsymbol{\hat \theta}_{\text{LSE}}$ where $\boldsymbol{\hat \theta}_{\text{LSE}}$ **least square estimator** for $\theta$:
$$
\boldsymbol{\hat y}_{\text{BLUE}}=X\boldsymbol{\hat \theta}_{\text{LSE}}=P_X\boldsymbol{y}=XX^+\boldsymbol{y}=X(X^TX)^{-1}\boldsymbol{y},
$$
where $X^+=(X^TX)^{-1}X^T$ is the Moore-Penrose inverse of $X$ and $P_X=XX^+$ is known as a projection matrix (symmetric and idempotent). 
* Residual of the BLUE is defined as $R_0^2 = \Vert \boldsymbol{y}-\boldsymbol{\hat y}_{\text{BLUE}}\Vert= \Vert (I_m-P_X)\boldsymbol{y}\Vert$ then $R^2_0/\sigma^2\sim\chi^2(m-r)$ where $r$ is the rank of $X$. So that $\hat\Sigma_y=R^2_0/(m-r)$ is a unbiased estimator for the $\sigma^2$.


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
\frac{\partial \boldsymbol{y}^T\boldsymbol{x}}{\partial\boldsymbol{x}}=\boldsymbol{y}^T,\quad \frac{\partial \boldsymbol{x}^T}{\partial\boldsymbol{x}}=I_n,\quad \frac{\partial \boldsymbol{x}^TA\boldsymbol{x}}{\partial\boldsymbol{x}}=(A+A^T)\boldsymbol{x}.
$$
2. Due to $P_X$ is the projection operator, then $I_m-P_X$ is symmetric and idempotent. Furthermore, $I_n=P_X+(I_m-P_X)$ and hence the Cochran's theorem shows that $\Vert \boldsymbol{y}\Vert/\sigma^2\sim \chi^2(r)$ while $R_0^2=\Vert I-P_X\Vert/\sigma^2\sim \chi^2(m-r)$. 

:::info
A explicit solution: $X^+X=X(X^TX)^-X^T$ where $X^-$ is a general inverse of $X$. But in practice, it is usually approximated by a gradient descent search to minimize the residual.
:::



It because the $(I_m-P_X)$ projects vector to the space perpendicular to the column space of $X$, and it is symmetry and idempotent $(I_m-P_X)^2=(I_m-P_X)$, it derive the first expression immediately. To get the expectation, we have
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

### Linear Estimation

Given a random variable $H\boldsymbol{\theta}$, it is estimable if
1. $\mathcal{C}(H^T)\subset \mathcal{C}(X^T)$, where $\mathcal{C}(X^T)$ stands for the column space of $X$.
2. The rank of the column space of $H$ is the same of the rank of $H$. 


## One-Way ANOVA
---
