# Chi-Square Test
 
## Pearson's chi-squared test
---
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

## Fisher's Exact Test
---

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