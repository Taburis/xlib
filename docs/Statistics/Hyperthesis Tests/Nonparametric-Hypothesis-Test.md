# Nonparametric Hypothesis Test

Comparing to the parametirc test, the nonparametric tests are more robust: no presumption on the underlying distributions. This kind of test is called **distribution free** test. But for the case which accesses to both parametric and nonparametric tests, the parametric tests have generally higher testing power with good significance level: better performance. 

## Sign Tests
---
Let $\lbrace X_n\rbrace$ be samples draw for a random varaible $X$. The test statistics for the hypothesis about the median $\xi_{0.5}$ of the $X$:
$$
H_0: \xi_{0.5}=\ell \quad H_1: \xi_{0.5}>\ell
$$
can be the sign statistics $S(\ell)$ defined as
$$
S(\ell)=\sum_{i=1}^n I(X_i>\ell),
$$
where $I(X>\ell)$ is a support function for the set $X>\ell$. The $p$-value of the test is given by
$$
\mathbb{P}_{H_0}[S\ge S(\ell)]=1-F_b[S(\ell)] 
$$
where $S(\ell)$ is the realized sign statistics from the sample and $F_b$ is the cumulative mass function of the binomial $b(n,1/2)$. For large $n$, the critical region can be
$$
z\ge z_\alpha,\quad z=\frac{S(\ell)-(n/2)}{\sqrt{n}/2},
$$
where $z_\alpha$ is the z-score with a $\alpha$-significance level. $z_\alpha=F^{-1}(\alpha)$ where $S(\ell)$ is the realized sign statistics from the sample and $F_b$ is the cumulative mass function of the binomial $b(n,1/2)$. The $p$-value is given by $F(Z\ge z)$. 

:::note MATHEMATICAL FOUNDATION
A median of $X$ is a one type of location functional $L_X$ of $X$. The samples $X_i$ can be reformed into a location model
$$
X_i=L_X+\epsilon_i, 
$$
where $\epsilon$ is i.i.d. to the PDF $f$ and CDF $F$. Under the $H_0$, the $I(X>\ell)$ is a Bernoulli distribution with $p=1/2$. Hence we have $S(\ell)\sim b(n, 1/2)$ (binomial distribution). Furthermore, the binomial distribution can be approximated by a normal distribution according to the central limit theory.
:::


## Wilcoxon Signed-Rank Test
---
Let $\lbrace X_n\rbrace$ be samples draw for a random varaible $X$. Assuming the PDF $f_X(x)$ of $X$ is **symmetric** about $0$ and $X$ satisifies the location model, we want to test the hypothesis:
$$
H_0: L_X= 0 \quad H_1: L_X>0.
$$

The test statistics is the Wilcoxon signed-rank statistics $T$ which is defined as
$$
T=\sum_{j=1}^n j\cdot\text{sign}(X_j),\quad \text{sign}(x) =\left\lbrace 
\begin{array}{cc}
1, & x>0\\
0, & x= 0\\
-1, & x<0
\end{array}\right.
$$ 
where the samples are supposed to be rearranged into the order
$$
|x_1|\le |x_2|\le \dots \le |x_n|.
$$
Under the $H_0$ assumption, $T$ satisifies the following properties:
1. Moment generate function of $T$ is
$$
\mathbb{E}(e^{sT})=\frac{1}{2^n}\prod_{j=1}^n(e^{-sj}+e^{sj}) = \prod_{j=1}^n\cosh(sj).
$$
$$
\mathbb{E}(T)=0 \quad \text{Var}(T)=n(n+1)(2n+1)/6.
$$
2. Asymptotic distribution:
$$
\frac{T}{\sqrt{\text{Var}(T)}}\sim N(0,1).
$$

:::note MATHEMATICAL FOUNDATION
First we show that $|X|\perp \text{sign}(X)$:
$$
\begin{aligned}
\mathbb{P}(X<x, \text{sign}(X)=1) &= \mathbb{P}(0<X<x)\\
& = \mathbb{P}(X<x)-\mathbb{P}(X>0)\\
& = \mathbb{P}(X<x)-\frac{1}{2}\\
& = \mathbb{P}(X>0)\cdot\mathbb{P}(|X|<x).
\end{aligned}
$$
The similar way to show $\mathbb{P}(X<x, X<0)=\mathbb{P}(X<0)\cdot\mathbb{P}(|X|<x)$ which conclude the independence. The MGF comes from 
$$
\begin{aligned}
\mathbb{E}(e^{sT})&=\mathbb{E}\left\lbrace \prod_{j=1}^n e^{j\cdot\text{sign}(X_j)}\right\rbrace\\
&=\prod_{j=1}^n\mathbb{E}\left\lbrace e^{sj\cdot \text{sign}(X_j)}\right\rbrace\\
&=\prod_{j=1}^n\frac{1}{2}\left\lbrace e^{sj}+e^{-sj}\right\rbrace.
\end{aligned}
$$
And the first order derivation of MGF with $s=0$ is 0. By the MGF, the variance is given by the sequence
$$
\text{Var}(T)=\sum_{j=1}^n j^2 = \frac{n(n+1)(2n+1)}{6}.
$$
To get this summation, one just need to notice that $k^2-(k-1)^2 = 2k-1$. 
:::