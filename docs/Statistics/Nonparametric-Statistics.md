# Nonparametric Statistics

## Empirical Distributions

Given a random variable with CDF $F(x)$, a sample $\lbrace X_n\rbrace$ from the $X$ forms a distribution $\hat F_n(x)$ known as **empricial distribution** (ECDF) defined as
$$
\hat F(x)=\frac{1}{n}\sum_{i=1}^nI(X_i\le x),
$$
where $I(X\le x)$ is an indicator function for the set satisfying the condition $X\le x$. Since the distribution of $X\le x$ is $F(x)$ which is the probability of $I(X\le x)=1$, the distribution of $n\hat F_n(x)$ is a binomial distribution $b(n,F(x))$ for any fixed $x$. Due to the asymptotic behavior of the binomial distribution, we have
$$
\sqrt{n}\left[\hat F_n(x)-F(x)\right]\xrightarrow{d} N\left(0, F(x)(1-F(x))\right).
$$

* **Mean** of $\hat F_n(x)$ is the $\bar x$, sample mean.
* **Variance** of $\hat F_n(x)$ is 
$$
\text{Var}(X)=\frac{1}{n}\sum_{i=1}^n(x_i-\bar x)^2,
$$
which is $\frac{n-1}{n}S^2$, where $S$ is the sample variance.

:::note **Glivenko-Cantelli** Theorem: 
Let $\lbrace X\rbrace_n$ is iid random variables with CDF $F$, where $F(x)$ is continuous in $\mathbb{R}$. The empricial distribution function $\hat F_n(x)$ converge to $F$ uniformly almost surely:
$$
\max_{x}|\hat F_n(x)-F(x)| \xrightarrow{a.s.} 0.
$$
:::

## Location Model

Given a random variable $X$ with PDF $f_X(x)$ and CDF $F_X(x)$, a **location functional** is a functional $T[F_X]$ satisfying the relation:
$$
T[F_{X+a}]=T[F_X]+a,\quad T[F_{aX}] = aT[F_X],\quad a\ne 0,
$$
where $a$ is a constant. For any location functional who is symmetry about $a$, which means that $T[F_{X-a}]=T[F_{a-X}]$, then $T[F_X]=a$, is a constant. 

A scale functional is a function $T[F_X]$ such that
$$
T[F_{aX}]=|a|T[F_X],\quad T[F_{X+a}]=T[F_X], \quad a \ne 0.
$$

For instance, the mean functional $T[F_X]=\mathbb{E}(X)$ and median $T[F_X]=F^{-1}_X(1/2)$ is a location functional while the standard deviation functional $T[F_X]=\sqrt{\text{Var}(X)}$ is a scale functional.

A random variable $Y$ follows a **location Model** $L_X=T[F_X]$ where $T$ is a location functional of $X$, means that
$$
Y=L_X+\epsilon,
$$ 
where $\epsilon$ is a random variable with pdf $f$ and CDF $F$. Given a i.i.d. sample represented by $\lbrace X_n\rbrace$, the distance between the $X_i$ and $X$ is defined by the variable $\epsilon_i=X_i-L_X$. Then $\epsilon_i$ are iid with PDF $f$ and CDF $F$. 