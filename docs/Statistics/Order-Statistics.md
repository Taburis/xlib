# Order Statistics

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
&=\int_a^b \frac{n!}{(k-1)!(n-k)!}F^k(y)[1-F(y)]^{n-k}f(y)dy\\
&=\int_0^1 \frac{n!}{(k-1)!(n-k)!} z^k(1-z)^{n-k}dz\\
&=\frac{k}{n+1},
\end{aligned}
$$
where the last step follows from the iterative partial integrals. It is exactly the definition of $\xi_p$.
3. Suppose $\xi_p$ is the $p$-th quantile, then $F(\xi_p)=p=\mathbb{P}(X<\xi_p)$. For $Y_i<\xi_p<Y_j$ in order statistics, it implies that there are $i$ elements needs to be less then $\xi_p$ while $n-j-1$ elements should be larger than $\xi_p$. The samples of the order statistics obeys a binomial distribution with probability $p$.  This leads to the expression above. 

## For Uniform Distribution
Suppose $U\sim\text{Uniform}(0,1)$ is a uniform distribution with value ranging $[0,1]$. Let the $k$-th order of $n$ iid. $U_i$, $i=1,\dots, n$ is denoted as $U_{k:n}$. Then the expectation is given by
$$
\begin{aligned}
&\mathbb{E}(U^m_{k:n})=\int_0^1\frac{n!}{(k-1)!(n-k)!}x^{k+m-1}(1-x)^{n-k}dx=\frac{\text{Beta}(k+m,n-k+1)\cdot n!}{(k-1)!(n-k)!},\\
&\text{Beta}(x,y)=\frac{\Gamma(x)\Gamma(y)}{\Gamma(x+y)}.
\end{aligned}
$$
Using the formula that $\Gamma(n+1)=n!$ for integer $n>0$, we have
$$
\mathbb{E}(U_{k:n})=\frac{k}{n+1},\quad \text{Var}(U_{k:n})=\frac{k(n+1-k)}{(n+1)^2(n+2)}.
$$

The uniform order statistics build a connection to the order statistics of other random variable $X$ with a CDF $F(x)$ through the formula:
$$
\mathbb{E}[X_{k:n}]=\mathbb{E}[F^{-1}(U_{k:n})],
$$
which can be validated by explicitly computing the expressions on both sides of the equation.

## For Normal Distribution

Given a normal random variable $X\sim N(0,1)$, and $X_{(k)}$ stands for the $k$-th ordered random variable for a size $n$ sample from $X$, then the expectation of $X_{(k)}$ is given by
$$
\mathbb{E}(X_{(k)})=\frac{n!}{(k-1)!(n-k)!}\int_{-\infty}^{+\infty}[\Phi(x)]^{k-1}[1-\Phi(x)]^{n-k}\phi(x)xdx
$$