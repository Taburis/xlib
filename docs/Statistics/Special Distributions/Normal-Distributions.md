# Normal Distributions


## Single Variable Normal Distributions
---
A normal distribution play a important role in statistics due to central limit theorem. A normal distribution $N(\mu,\sigma)$ is determined by mean $\mu$ and variance and $\sigma$. Suppose a random variable obeys the normal distribution, the PDF is
$$
f(x)=\frac{1}{\sqrt{2\pi\sigma^2}}\exp\left(-\frac{(x-\mu)^2}{2\sigma^2}\right).
$$ 
A normal distribution $N(0,1)$ is called a **standard normal distribution**.

:::info Properties: (suppose $X\sim N(\mu_1,\sigma^2)$)
1. Given $Y\sim N(\mu_2,\sigma_2)$ and $X\perp Y$, then $X+Y\sim N(\mu_1+\mu_2,\sigma_1^2+\sigma^2)$
2. Suppose $c\in\mathbb{R}$, then $cX\sim N(c\mu,c^2\sigma^2)$. 
3. If $Z=(X-\mu)/\sigma$, then $Z\sim N(0,1)$ and $Z^2\sim\chi^2(1)$, where $\chi(1)$ is a chi square distribution with degree of freedom 1.
4. The moment generate function of $X$ is $M_X(t)=\exp(\mu t+\sigma^2 t^2/2)$. 
:::

## Multivariate Normal Distribution
---
A random vector $\boldsymbol{X}=(X_1,\dots, X_m)$ forms by $m$ normal distribution $X_i\sim N_m(\mu_i,\sigma_i)$, then the random vector forms $n$-dimensional normal distribution $\boldsymbol{X}\sim N_m(\boldsymbol{\mu}, \Sigma)$ where $\Sigma$ is a covariance matrix which is positive and symmetric. 

:::info Properties:
* The linear transform $\Gamma$ on $\boldsymbol{X}$ will leads to $Y=\Gamma\boldsymbol{X} = N_m(\Gamma\boldsymbol{\mu},\Gamma^T\Sigma\Gamma)$. 
* The moment generate function $M_{\boldsymbol{X}}(\boldsymbol{t})=\exp\lbrace \boldsymbol{t}^T\boldsymbol{\mu}+\boldsymbol{t}^T\Sigma\boldsymbol{t}/2\rbrace$. 
* Suppose 
$$
\boldsymbol{X}=
\begin{pmatrix}
\boldsymbol{X_1}\\
\boldsymbol{X_2}
\end{pmatrix}
,\quad \Sigma= 
\begin{pmatrix}
\Sigma_{11} & \Sigma_{12}\\
\Sigma_{12}^T & \Sigma_{22}
\end{pmatrix},
$$
then $\boldsymbol{X}_1\perp\boldsymbol{X}_2$ if and only if $\Sigma_{12}=0$. And the conditional variable 
$$
\boldsymbol{X}_1|\boldsymbol{X}_2\sim N_m(\boldsymbol{\mu}_1+\Sigma_{12}\Sigma_{22}^{-1}(\boldsymbol{X}_2-\boldsymbol{\mu}_2), \Sigma_{11}-\Sigma_{12}\Sigma_{22}\Sigma_{12}^T)
$$
* Let $W=(\boldsymbol{X}-\boldsymbol{\mu})^T\Sigma^{-1}(\boldsymbol{X}-\boldsymbol{\mu})$, then $W\sim\chi^2(m)$.
:::

Suppose orthogonal transform $\Gamma$ diagonalize $\Sigma$, and $\Gamma\Sigma\Gamma^T=\Lambda$ where the diagonal elements are $\lambda_i$. Suppose $\lambda_1\ge\lambda_2\ge\dots$. The variable $\boldsymbol{Y}=\Gamma(\boldsymbol{X}-\boldsymbol{\mu})$ is called **principal components** and $\text{tr}(\Sigma)$ is called **total variation**. The first principal component $Y_1$ has the largest variation over all linear combination $\boldsymbol{a}^T(\boldsymbol{X}-\boldsymbol{\mu})$, where $\boldsymbol{a}^T\cdot\boldsymbol{a} = 1$. Because the variation is a linear combination of all principal components with the constraints that $\sum_ia_i^2=1$ and $\lambda_1$ is the largest component.


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