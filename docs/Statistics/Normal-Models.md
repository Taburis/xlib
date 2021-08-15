# Normal Models

:::note Theorems:
Given a random vector $\boldsymbol{X}=(X_1,\dots, X_n)$, then
1. Let $\boldsymbol{X}\sim N_n(0,I_n\sigma^2)$, and $Q=\boldsymbol{X}^TA\boldsymbol{X}/\sigma^2$ where $A$ is a symmetric real matrix. Then the moment generate function
$$
M(t)=\prod_{i=1}^r(1-2t\lambda_i)^{-\frac{1}{2}}=\det(1-2tA)^{-\frac{1}{2}},\quad r=\text{rank}(A).
$$
where $\lambda_i$ is the $i$-th eignvalue of $A$. Furthermore, $Q\sim\chi^2(r)$ if and only if $A^2=A$ (idempotent).
2. Suppose $\boldsymbol{X}\sim N_n(\boldsymbol{\mu},\Sigma)$, and $Q=(\boldsymbol{X}-\boldsymbol{\mu})^T\Sigma^{-1}(\boldsymbol{X}-\boldsymbol{\mu})$, thne $Q\sim \chi^2(n)$.
3. Suppose $\boldsymbol{X}\sim N_n(\boldsymbol{\mu},\Sigma)$ and $Q=(\boldsymbol{X}-\boldsymbol{\mu})^TA(\boldsymbol{X}-\boldsymbol{\mu})$ where $A$ is symmetry. Then $Q\sim\chi^2(r)$ where $r=\text{rank}(A)$ if and only if $A\Sigma A=A$. 
:::

**Proof**: 
1. Since $A$ is symmetry, there's a spectra decomposition of $A=\Gamma \Lambda \Gamma^T$ where $\Lambda$ is diagnoal matrix with $\Lambda_{ii}=\lambda_i$ and $\Gamma\Gamma^T=1$. Then the diagnolization $\boldsymbol{Z}=\Gamma\boldsymbol{X}$ leads to $Z\sim N_n(0,I_n\sigma^2)$ and $Q=\sum_i\lambda_iZ_i^2$. The moment generate function is
$$
\begin{aligned}
M_X(t) = M_Z(t)&=\mathbb{E}\left\lbrace e^{\sum_{i=1}^r\lambda_i Z_i}\right\rbrace\\
&=\prod_{i=1}^r\mathbb{E}(e^{t\lambda_iZ_i})\\
&=\prod_{i=1}^r(1-2t\lambda_i)^{-\frac{1}{2}}\\
&=\det(I_n-2tA).
\end{aligned}
$$
Furthermore, if $A^2=A$, then the eigenvalues of $A$ is $\lambda_i=1$ for $i=1,\dots, r$ where $r=\text{rank}(A)$. Put this but to the moment generate function above leads to the exact formula of moment generate function of $\chi^2(r)$ distriubtion. 
2. Let $\boldsymbol{Y}=\boldsymbol{X}-\boldsymbol{\mu}$, and a spectra decompose exists that $\Sigma=\Gamma^T\Lambda\Gamma$ so that $\boldsymbol{Z}=\Gamma\boldsymbol{Y}\Sigma^{-1/2}$ satifies $\boldsymbol{Z}\sim N_n(0,I_n)$ since $\Sigma^{-1}=(\Gamma^T\Lambda\Gamma)^{-1}=\Gamma^T\Lambda^{-1}\Gamma$ and $\Gamma$ is a orthogonal matrix. On the other hand, $Q=\boldsymbol{Y}^T\Sigma^{-1}\boldsymbol{Y}=\boldsymbol{Z}^T\boldsymbol{Z}$ so that $Q\sim\chi^2(n)$.
3. Define $Z=\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})$ then $Z\sim N_n(0,I_n)$. On the other hand, $Q=[\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})]^T\Sigma^{1/2}A\Sigma^{1/2}[\Sigma^{-1/2}(\boldsymbol{X}-\boldsymbol{\mu})]$. Based on the 1st theorem, $Q\sim\chi^2(r)$ if and only if $\Sigma^{1/2}A\Sigma^{1/2}$ is idempotent which require $A\Sigma A=A$.