# Hypothesis Tests

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

### Chi-Square Test
 
#### Pearson's chi-squared test
Given random variables of a multinomial distribution $\lbrace X_k\rbrace$, which $X_i$ stands for the occurrence of the $i$-th outcome and $k$ distinct outcomes in total, and assuming the test hypothesis is that
$$
H_0: p_i=\mu_i,\: \forall i=1,\dots,k.\quad H_1: \text{otherwise}.
$$
The statistics for the hypothesis test is defined as 
$$
Q=\sum_{i=1}^k\frac{(X_i-n\mu_i)^2}{n\mu_i},\quad\sum_i\mu_i=1,\quad Q\sim\chi^2(k-1),
$$
and the critical region of size $\alpha$ is determined according to the $\chi^2(k-1)$ distribution.

:::note Mathematical Foundation
Each multinomial random variable $X_i$ is actually the sum of Bernoulli's random variables $B^i_j$: $X_i=\sum_{j=1}^nB^i_j$. According to the Central Limit Theorem, $\overline{B_i}=\sum_jB^i_j/n=X_i/n$ is approximately a normal distribution with mean $\mathbb{E}(B^i_j)=p_i$ and variance $p_i(1-p_i)$ when $n$, the number of the total trials, is large. Explicitly: 
$$
\sqrt{n}\frac{\overline{B_i}-p_i}{\sqrt{p_i(1-p_i)}}=\frac{X_i-np_i}{\sqrt{np_i(1-p_i)}}\xrightarrow{d} N(0,1),
$$
Define $Y_i=(X_i-np_i)/\sqrt{np_i}$, then $Y_i\sim N(0,1-p_i)$. Furthermore, for the covariance we have
$$
\text{Cov}(Y_i,Y_j)=\text{Cov}(X_i/n,X_j/n)/(n\sqrt{p_ip_j})=-\sqrt{p_ip_j}
$$
and hence $\boldsymbol{Y}\sim N_k(0,\Sigma)$ where $\Sigma_{ii}=1-p_i$ and $\Sigma_{ij}=-\sqrt{p_ip_j}$. Noticing that if we introduce a vector $\boldsymbol{v}=(\sqrt{p_1},\dots,\sqrt{p_k})^T$, then $\Sigma=I_k-\boldsymbol{v}\boldsymbol{v}^T$. Further more, it is easy to verify that $P=\boldsymbol{v}\boldsymbol{v}^T$ is symmetric and idempotent since $P^2=P(\sum_ip_i)=P$ and $\text{rank}(P)=1$. It means that $\text{rank}(\Sigma)=k-1$ and hence there's a transform $\Gamma$ diagonalizing the $\Sigma$ so that $k-1$ diagonal elements are 1 and the rest is 0. Based on the [Cochran's theorem](Normal-Models.md#mathematical-foundation):
$$
\boldsymbol{Y}^T\boldsymbol{Y}=\sum_{i=1}^{k-1}Z_i^2\sim\chi^2(k-1),
$$ 
which proved the conclusion. 
:::

### Fisher's Exact Test

The fact that multinomial distribution approximates to normal distribution in large sample size makes the Chi-square test is not applicable to the case with limited sample size. The Fisher's exact test is introduced to small sample hypothesis test.

Given a $n$-size sample satisfying a multinomial distribution $X_i$ according to features $\lbrace A_k\rbrace$ ($X_i$ is the number of samples labeled as $A_i$). Furthermore, these samples can also be labeled according a other feature set $\lbrace B_h\rbrace$ which satisfies a multinomial distribution denoted as $Y_j$. The hypothesis is
$$
H_0: X_i\perp Y_j \quad H_1:\text{otherwise},
$$
and the random variable $P_{i,j}$ stands for the counts for samples labeled as $(i,j)$.

For the case that $X$ and $Y$ are binomial, the distributions of $P_{i,j}$ has only 4 possible combination and can be represented by a $2\times 2$ table

|   | A1| A2 | Margin B |
|---|----|-----|------------|
|**B1**| a |b | a+b|
|**B2**| c | d| c+d|
|**Margin A**|a+c| b+d| n=a+b+c+d|

Fisher shows that $P_{i,j}$ satisfies the hypergeometric distribution under the null hypothesis. The probability to have the table above is: 
$$
p=\frac{\binom{a+b}{a}\binom{c+d}{c}}{\binom{n}{a+c}}=\frac{\binom{a+b}{b}\binom{c+d}{d}}{\binom{n}{b+d}}=\frac{(a+b)!(c+d)!(a+c)!(b+d)!}{a!b!c!d!n!}.
$$

:::danger
The probability $p$ alone can't tell anything. To understand how rare of this sample, it is crucial to understand where this result located overall the hypergeometric distribution. The idea is that we need to calculate the probability to have a more extreme (imbalanced) result than the current one, which is sum over the probability of all extreme cases. It could be sum over the possible result imbalanced toward one-side or two-sides, depends on the problem we are facing. The cumulative probability of the extreme cases is the power or confidential to the null hypothesis.
:::

### t-Test

The $t$-Test is a general term referring to all the hypothesis tests using a statistics following the $t$-distribution.


#### Independent two-sample $t$-test

Given two samples $\lbrace X_m\rbrace$ and $\lbrace Y_n\rbrace$ with $X_i\perp Y_j$, assuming the mean $\mu_X,\mu_Y$ and variance $\sigma_X,\sigma_Y$ exists for both $X_i,Y_j$, the test hypothesis is
$$
H_0: \mu_X = \mu_Y\quad H_1: \mu_X\ne \mu_Y.
$$
* With a presumption that $\sigma_X=\sigma_Y$, the statistics $T$ for testing is defined as 
$$
T\sim t(m+n-2),\quad T=\frac{\overline{X}-\overline{Y}}{S_p\sqrt{\frac{1}{m}+\frac{1}{n}}}, \quad S_p^2=\frac{(m-1)S^2_X+(n-1)S^2_Y}{m+n-2}.
$$
* Suppose $\sigma_X\ne\sigma_Y$, the statistics $T$ defined for testing is defined as
$$
 T\sim t(\nu),\quad T=\frac{\overline{X}-\overline{Y}}{\sqrt{s^2_{\overline{X}}+s^2_{\overline{Y}}}}, \quad s_{\overline{X}} = \frac{S_X}{\sqrt{n}}
$$
where $s_{\overline{X}}$ is the standard error of the mean $\overline{X}$, $S_X^2$ is the sample variance of $X$, and the degree of freedom $\nu$ is defined as
$$
\nu = \frac{\left(\frac{S^2_X}{m}+\frac{S^2_Y}{n}\right)^2}{\frac{S_X^4}{m^2(m-1)}+\frac{S^4_Y}{n^2(n-1)}}
$$
This test is also known as **Welch's $t$-test**.

:::note Mathematical Foundation 
From the [student's theorem](Special-Distributions.md#t-distributions) we know that following facts:
1. $\overline{X}\sim N(\mu,\sigma^2_x/m)$ and $\overline{X}\sim N(\mu_y,\sigma^2_y/n)$.
2. $(m-1)S_X^2/\sigma^2_X\sim\chi^2(m-1)$ and $(n-1)S_Y^2/\sigma^2_Y\sim\chi^2(n-1)$.
3. $\overline{X}\perp S^2_X$ and $\overline{Y}\perp S^2_Y$.

The null hypothesis assumed that $\mu_X=\mu_Y$ and hence $Z\sim N(0, S_{X+Y}^2)$ where $Z=\overline{X}-\overline{Y}$ and $S_{X+Y}^2=\frac{\sigma^2_X}{m}+\frac{\sigma^2_Y}{n}$. It is easy to prove that $(X_i\overline{X})\perp Z$ and hence $Z\perp S_X^2$, so is $Z\perp S_Y^2$. Now we can define a variance 
$$
S^2=(m-1)S^2_X/\sigma^2_X+(n-1)S_Y^2/\sigma^2_Y,\quad S^2\sim\chi^2(n+m-2)
$$
and $S\perp Z$. This means that if we define 
$$
T=\frac{Z/\sqrt{S^2_{X+Y}}}{\sqrt{S^2/(m+n-2)}},
$$
then $T\sim t(m+n-2)$. This expression can be further simplified if we assuming $\sigma_X=\sigma_Y$, which leads to the $t$-test:
$$
T=\frac{\overline{X}-\overline{Y}}{S_p\sqrt{\frac{1}{m}+\frac{1}{n}}},\quad S_p^2=\frac{(m-1)S^2_X+(n-1)S^2_Y}{m+n-2},
$$
where $S_p^2$ is also known as **pooled standard deviation**.
:::

### Z-Test

Given a sample $\lbrace X_n\rbrace$ drawn from a distribution with mean $\mu$ and variance $\sigma$. When the $\sigma$ is known, we define the statistics $Z=(\overline{X}-\mu_0)/\sqrt{\sigma/n}$, which has $Z\sim N(0,1)$ assuming $\mu=\mu_0$. And it can be used to test the hypothesis below (with the $p$-value followed):  
1. $H_0: \mu\ge \mu_0$, $p$-value: $\mathbb{P}(Z<z)=\Phi(z)$ (left side tail)
2. $H_0: \mu\le \mu_0$, $p$-value: $\mathbb{P}(Z>z)=1-\Phi(z)=P(-z)$ (right side tail)
3. $H_0: \mu=\mu_0$, $p$-value: $\mathbb{P}(|Z|<z)=2\Phi(-|z|)$ (two side tails)

where $z$ is the value of $Z$ obtained from the sample, and $\Phi(z)$ is the cumulative function of standard normal distribution. The $p$-value stands for the extreme case of $Z$ assuming the null hypothesis is true. If $\mu\ge \mu_0$, then extreme case that having a large $Z$ should not be small. It means tha a small $p$-value means that null hypothesis is unlikely to be true. The $p$-values picked for other two hypothesises based on the similar logic.

:::tip T-test vs Z-test
Usually the variance is unknown and needs to be estimated by sample variance $S^2/(n-1)$. In this case, Z-test is an approximation to the t-test. Furthermore, the t-test is favored for the small sample size (30 or less) case. 
:::


## Nonparametric Hypothesis Test

Comparing to the parametirc test, the nonparametric tests are more robust: no presumption on the underlying distributions. This kind of test is called **distribution free** test. But for the case which accesses to both parametric and nonparametric tests, the parametric tests have generally higher testing power with good significance level: better performance. 

### Sign Tests

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


### Wilcoxon Signed-Rank Test

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