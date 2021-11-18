# Nonparametric Statistics

## Empirical Distributions

Given a random variable with CDF $F(x)$, a sample $\lbrace X_n\rbrace$ from the $X$ forms a distribution $\hat F_n(x)$ known as **empricial distribution** (ECDF) defined as
$$
\hat F(x)=\frac{1}{n}\sum_{i=1}^nI(X_i\le x),
$$
where $I(X\le x)$ is an indicator function for the set satisfying the condition $X\le x$. Since the distribution of $X\le x$ is $F(x)$ which is the probability of $I(X\le x)=1$, the distribution of $n\hat F_n(x)$ is a binomial distribution $b(n,F(x))$ for any fixed $x$. By the central limit theory, we have
$$
\sqrt{n}\left[\hat F_n(x)-F(x)\right]\xrightarrow{d} N\left(0, F(x)(1-F(x))\right).
$$

* **Mean** of $\hat F_n(x)$ is the $\bar x$, sample mean.
* **Variance** of $\hat F_n(x)$ is 
$$
\text{Var}(X)=\frac{1}{n}\sum_{i=1}^n(x_i-\bar x)^2,
$$
which is $\frac{n-1}{n}S^2$, where $S$ is the sample variance.

## Order Statistics

Given a sample $\lbrace X_n\rbrace$ draw from a PDF $f(x)$ for random variable $X$, let $\lbrace Y_n\rbrace$ represents the statistics of $\lbrace X_n\rbrace$ with the order that $Y_1<Y_2<\dots<Y_n$. 
1. The probability to have $\lbrace y_n\rbrace$ is denoted as $g(y_1,\dots, y_n)$, then
$$
g(y_1,\dots, y_n)=n!\prod_{i=1}^nf(y_i),\quad y_1<y_2<\dots<y_n.
$$
The PDF of $Y_k$ is given by the marginal PDF
$$
\begin{aligned}
g_k(y_k)&=\int_a^{y_2}\dots\int_a^{y_{k}}\int_{y_k}^b\dots\int_{y_k}^b n!f(y_1)\dots f(y_{k-1})f(y_{k+1})\dots f(y_n)dy_1\dots dy_{k-1}dy_{k+1}\dots y_n\\
&=\frac{n!}{(k-1)!(n-k)!}F^{k-1}(y_k)\left[1-F(y_k)\right]^{n-k}f(y_k),\quad a\le y_k\le b,
\end{aligned}
$$
where $a,b$ is the lower/upper bound of the range of $X$.
2. A $p$-th **Quantile** $\xi_p$ is defined as $\xi_p=F^{-1}(p)$ where the $F(x)$ is the CDF of $X$. For instance, a $0.5$ quantile is the median of $X$. An approximated estimator of the $\xi_p$ is $\hat\xi_p=Y_k$ where $k=\lfloor p(n+1) \rfloor$ and $n$ is the size of the samples. 
3. The probability of $\xi_p\in (Y_i,Y_j)$ for the order statistics $\lbrace Y_n\rbrace$ is given by
$$
\mathbb{P}(Y_i < \xi_p < Y_j) = \sum_{k=i}^{j-1}\binom{n}{k}p^k(1-p)^k,
$$
Let $\gamma= \mathbb{P}(Y_i < \xi_p < Y_j)$, then $(Y_i,Y_j)$ formed the $100\cdot\gamma\%$-**confidential interval**.

**Proof**
1. There are $n!$ possible ways to get the order $y_1<\dots<y_n$ from the samples $\lbrace x_n\rbrace$, which are drawn at the probability $\prod_if(x_i)$.
2. To show if $Y_k$ is an unbiased estimator for $\xi_p$, we show $\mathbb{E}[F(Y_k)]=p$ based on the definition of $\xi_p$. In fact, the expected CDF is given by
$$
\begin{aligned}
\mathbb{E}[F(Y_k)]&=\int_a^bg_k(y)F(y)dy\\
&=\int_a^b \frac{n!}{(k-1)!}{(n-k)!}F^k(y)[1-F(y)]^{n-k}dy\\
&=\int_0^1 \frac{n!}{(k-1)!}{(n-k)!} z^k(1-z)^{n-k}dz\\
&=\frac{k}{n+1},
\end{aligned}
$$
where the last step follows from the iterative partial integrals. It is exactly the definition of $\xi_p$.
3. Suppose $\xi_p$ is the $p$-th quantile, then $F(\xi_p)=p=\mathbb{P}(X<\xi_p)$. For $Y_i<\xi_p<Y_j$ in order statistics, it implies that there are $i$ elements needs to be less then $\xi_p$ while $n-j-1$ elements should be larger than $\xi_p$. The samples of the order statistics obeys a binomial distribution with probability $p$.  This leads to the expression above.

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