# t- and Z-Tests


## t-Tests
---
The $t$-Test is a general term referring to all the hypothesis tests using a statistics following the $t$-distribution.


### Independent two-sample $t$-test

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

## Z-Tests
---

Given a sample $\lbrace X_n\rbrace$ drawn from a distribution with mean $\mu$ and variance $\sigma$. When the $\sigma$ is known, we define the statistics $Z=(\overline{X}-\mu_0)/\sqrt{\sigma/n}$, which has $Z\sim N(0,1)$ assuming $\mu=\mu_0$. And it can be used to test the hypothesis below (with the $p$-value followed):  
1. $H_0: \mu\ge \mu_0$, $p$-value: $\mathbb{P}(Z<z)=\Phi(z)$ (left side tail)
2. $H_0: \mu\le \mu_0$, $p$-value: $\mathbb{P}(Z>z)=1-\Phi(z)=P(-z)$ (right side tail)
3. $H_0: \mu=\mu_0$, $p$-value: $\mathbb{P}(|Z|<z)=2\Phi(-|z|)$ (two side tails)

where $z$ is the value of $Z$ obtained from the sample, and $\Phi(z)$ is the cumulative function of standard normal distribution. The $p$-value stands for the extreme case of $Z$ assuming the null hypothesis is true. If $\mu\ge \mu_0$, then extreme case that having a large $Z$ should not be small. It means tha a small $p$-value means that null hypothesis is unlikely to be true. The $p$-values picked for other two hypothesises based on the similar logic.

:::tip T-test vs Z-test
Usually the variance is unknown and needs to be estimated by sample variance $S^2/(n-1)$. In this case, Z-test is an approximation to the t-test. Furthermore, the t-test is favored for the small sample size (30 or less) case. 
:::
