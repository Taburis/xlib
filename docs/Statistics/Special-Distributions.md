# Special Distributions

## Normal Distributions

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

**Multivariate Normal Distribution**: A random vector $\boldsymbol{X}=(X_1,\dots, X_m)$ forms by $m$ normal distribution $X_i\sim N_m(\mu_i,\sigma_i)$, then the random vector forms $n$-dimensional normal distribution $\boldsymbol{X}\sim N_m(\boldsymbol{\mu}, \Sigma)$ where $\Sigma$ is a covariance matrix which is positive and symmetric. 

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