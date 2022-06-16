# Binomial and Beta Distribution

## Bernoulli Experiment
---
A **Bernoulli experiment** is a class of statistical experiment that only 2 possible mutually exclusive outcome at a fixed probability for each trail. The random variable $X$ for this experiment is formulated as a Bernoulli distribution $p(x)$ that
$$
p(x)=p^x(1-p)^{1-p},\quad x=0,1.
$$
The $p$ is the probability to get the outcome represented by $1$. For $X$, it is easy to verify that $\mathbb{E}(X)=p$, $\text{Var}(X)=p(1-p)$. 

A random variable represents the outcome of a sequence of $n$ Bernoulli trials forms the **Binomial distribution** $b(n,p)$ where $p$ is the probability to get 1 for each trail. The PMF $f(x)$ is
$$
f(x)=\begin{pmatrix}n \\ x\end{pmatrix}p^{x}(1-p)^{n-x}, \quad \begin{pmatrix}n,x\end{pmatrix}=\frac{n!}{x!(n-x)!}.
$$
where $x=0,1,\dots, n$. 
:::info Properties:
* The moment generate function for binomial distribution is $M(t)=[(1-p)+pe^t]^n$.
* $\mathbb{E}(X)=np$ and $\text{Var}(X)=np(1-p)$. 
* Suppose $X_i\sim b(n_i,p)$ and $Y=\sum_iX_i$, then $Y\sim b(\sum_in_i,p)$. 
:::

## Multinomial Distributions
---
The Bernoulli trail can be extended to the case that 3 possible mutually exclusive outcomes from each trails. The random variable represents the outcome of $n$ sequence trials has a **trinominal distribution** $f(x,y)$:
$$
f(x,y)=\frac{n!}{x!y!(n-x-y)!}p_1^xp_2^y(1-p_1-p_2)^{n-x-y},
$$ 
where $x+y\le n$. The moment generate function is
$$
M(t_1,t_2)=(p_1e^{t_1}+p_2e^{t_2}+1-p_1-p_2)^n
$$

To generalize this idea to the case that $k$ outcomes for each trial, and let the random vector $\boldsymbol{X}=(X_1,\dots,X_k)$ be the outcomes where $X_i$ is the counts of $i$-th outcome occurance in $n$ trials. This leads to a **multinomial distribution** with a PDF
$$
f(x_1,\dots,x_k)=\frac{n!}{\prod_{i=1}^kx_i!}\prod_{i=1}^kp_i^{x_i},\quad \sum_{i=1}^np_i=1,
$$
where $p_i$ stands for the probability of the $i$-th outcome occurs in each trial. It is easy to varify that
1. $\mathbb{E}(X_i)=np_i$ and $\text{Var}(X_i)=np_i(1-p_i)$.
2. $\text{Cov}(X_iX_j)=-np_ip_j$. 

## Beta Distribution
---
A beta distribution $B(\alpha,\beta)$ is defined as the PDF:
$$
f(x;\alpha,\beta)= \frac{1}{B(\alpha,\beta)}x^{\alpha-1}(1-x)^{\beta-1}=\frac{\Gamma(\alpha+\beta)}{\Gamma(\alpha)\Gamma(\beta)}x^{\alpha-1}(1-x)^{\beta-1}.
$$
This is generalized from substituting the factorial in the binomial distribution by the gamma function. The **cumulative function** is given by
$$
F(x;,\alpha,\beta)=\frac{B(x;\alpha,\beta)}{B(\alpha,\beta)},\quad B(x;\alpha,\beta) = \int_0^xt^{\alpha-1}(1-t)^{\beta-1}dt.
$$
The properties of the beta distribution:
* **mean** $\mu=\mathbb{E}(X)=\frac{\alpha}{\alpha+\beta}$.
* **variance** $\text{var}(X)= \frac{\alpha\beta}{(\alpha+\beta)^2(\alpha+\beta+1)}$.

Similar method can be used to extend the multinomial distribution as well. So that we have a multi-variate beta distribution:
$$
f(x_1,\dots, x_n;\alpha_1,\dots, \alpha_n)=\frac{1}{B(\alpha_1,\dots,\alpha_n)}\prod_{i=1}^np_i^{x_i}, \quad B(\alpha_1,\dots,\alpha_n)=\frac{\prod_{i=1}^n\Gamma(\alpha_i)}{\Gamma(\sum_{i=1}^n{\alpha_i})}.
$$