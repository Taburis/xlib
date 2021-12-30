# Test Theory

## Power and Errors
---

A **null hypothesis** $H_0$ is a statement exclusive to the **alternative hypotheses** $H_1$. The hyperthesis test is designed to ansewr the question that should the null hyperthesis be rejected? A test is based on the samples $\boldsymbol{X}=(X_1,\dots, X_n)$ from a random variable $X$ defined on a probability space $(\Omega, \mathfrak{R},\mathbb{P})$ and the support of $X$ is $\mathcal{S}\subset\Omega$. The test rule is 
$$
\begin{aligned}
&\text{Reject } H_0 (\text{Accept } H_1),\quad \text{if } \boldsymbol{X}\in C\\
&\text{Retain } H_0 (\text{Reject } H_1),\quad \text{if } \boldsymbol{X}\in \Omega\setminus C.
\end{aligned}
$$
where $C\subset \mathcal{S}$ is called **critical region**. The possible error is classified as
* **Type I** error: $H_0$ is rejected while it is true (false positive). It is quantified as the **significance level** or **szie** $\alpha$:
$$
\alpha=\max_{H_0}\mathbb{P}_{H_0}(\boldsymbol{X}\in C),
$$
The upper bound of the probability that $\boldsymbol{X}\in C$ (rejecting $H_0$) assuming $H_0$ holds.
* **Type II** error: $H_0$ is retained while it is false (false negative). It is quantified as $\beta=1-\max_{H_1}\gamma_C(H_1)$ where $\gamma_C(H_1)$ is the **power function** of the critical region $C$ for $H_1$:
$$
\gamma_C(H_1)=\mathbb{P}_{H_1}(\boldsymbol{X}\in C).
$$

A **best critical region** of size $\alpha$ is a critical region $C$ with size $\alpha$ has the largest power amount all region with size $\alpha$. 

## Parametric Hypothesis Test
---

A typical parametric hypothesis can be formulated as
$$
H_0: \theta\in \omega_0\quad H_1:\theta\in\omega_1,
$$
where $\omega_0$ and $\omega_1$ are subset of $\Omega$ and mutually disjoint. 

For point wise parametric test:
$$
H_0: \theta=\theta_0\quad H_1:\theta=\theta_1,
$$

The random variable $X$ used in the test derives a measure on the sample space $S$ as 
$$
\mu_{X,\theta}(C) = \int_CL(\theta;\boldsymbol{x})d\mathbb{P}(\boldsymbol{x}), \quad C\in S
$$
where $L(\theta;\boldsymbol{x})$ is known as likelihood:
$$
L(\theta; \boldsymbol{x})=\prod_{i=1}^nf(x_i;\theta),\quad \boldsymbol{x}=(x_1,\dots,x_n),
$$
and $f(x;\theta)$ is the PDF of $X$ with the assumed parameter $\theta$. With this notation, the size of a critical region is $\mu_{X,\theta_0}(C)$ and the power is $\mu_{X,\theta_1}(C)$. A test with the critical region $C$ is **unbiased** if $\mu_{X,\theta_1}(C)\ge\mu_{X,\theta_0}(C)$.

:::info Neyman-Pearson Theorem:
Suppose $C$ is a critical region in sample space such that $\exists k> 0$:
1. $\frac{L(\theta_0,\boldsymbol{x})}{L(\theta_1,\boldsymbol{x})}\le k$ for any $\boldsymbol{x}\in C$.
2. $\frac{L(\theta_0,\boldsymbol{x})}{L(\theta_1,\boldsymbol{x})}\ge k$ for any $\boldsymbol{x}\in E\setminus C$.
Then $C$ is a best critical region of size $\alpha=\max_{H_0}\mathbb{P}_{H_0}(\boldsymbol{X}\in C)$. 
:::

**Proof** Notice that
To show the best power of $C$ which is $\mu_{X,\theta_1}(C)\ge\mu_{X,\theta_1}(A)$ for any other region $A$ such that $\mu_{X,\theta_0}(A)=\alpha$, we decompose $C=(C\cap A)\cup(C\cap (E\setminus A))$ and $A=(C\cap A)\cup(A\cap (E\setminus C))$. By the condition 1 and 2, $\mu_{X,\theta_1}(C)\ge \mu_{X,\theta_1}(E\setminus C)$ implies $\mu_{x,\theta_1}[C\cap (E\setminus A)]\ge \mu_{X,\theta_1}[A\cap (E\setminus C)]$. Using the linearity $\mu_{X,\theta_1}(C\cup D)=\mu_{X,\theta_1}(C)+\mu_{X,\theta_1}(D)$ for any $C\cap D=\varnothing$ then the theorem follows.



