# t- and F-Distributions

## t-Distributions
---
It will show later that the sample mean and sample variance is independent with each other. Provided the distribution of sample mean and variance is known by CLT, the distribution of normalized sample mean satisifes a $t$**-distribution**. The PDF of the $t$-distribution with degree of freedom $r$, denoted as $t(r)$, is 
$$
f(t)= \frac{\Gamma(\frac{r+1}{2})}{\Gamma(\frac{r}{2})}\frac{1}{\sqrt{\pi r}}(1+t^2/r)^{-(r+1)/2}, \quad t \in \mathbb{R}.
$$

This distribution obtained from the joint PDF of $X\sim N(0,1)$ and $Y\sim \chi^2(r)$ and $X\perp Y$:
$$
f_{X,Y}(x,y)=\frac{e^{-x^2/2}}{\sqrt{2\pi}}\frac{y^{r/2}e^{-y/2}}{\Gamma(r/2)2^{r/2}},\quad x\in\mathbb{R}, y\in[0,+\infty].
$$

The $t$-distribution follows from a substitution that $t=x/\sqrt{y/r}$ and integral over the $y$.

:::info Student's Theorem:
Let $X_1,\dots, X_n$ be i.i.d. random variables each having $X_i\sim N(\mu,\sigma^2)$. The new random variables $\overline{X}$ and $S^2$ defined are defined as
$$
\overline{X}=\sum_{i=1}^nX_i/n,\quad S^2=\sum_{i=1}^n(X_i-\overline{X})^2/(n-1)
$$
satisfies:
1. $\overline{X}\sim N(\mu, \sigma^2/n)$;
2. Independent: $\overline{X}\perp S^2$;
3. $(n-1)S^2/\sigma^2\sim \chi^2(n-1)$;
4. Let $Z=(\overline{X}-\mu)/\sqrt{\frac{\sigma^2}{n}}$, then $T=Z/\sqrt{\frac{S^2}{n-1}}=(\overline{X}-\mu)/(S/\sqrt{n})$ satisfies the $t$-distribution $T\sim t(n-1)$.
:::

**Proof**: 
1. The proof of the 1st refers to the [section](#normal-distributions). 
2. Since all $X_i$ and $\overline{X}$ are normal distributions, then $X_i-\overline{X}\perp\overline{X}$ if and only if $\text{Cov}(X_i-\overline{X},\overline{X})=0$. Straight calcuation shows
$$
\begin{aligned}
\text{Cov}(X_i-\overline{X},\overline{X})&=\mathbb{E}(X_i\overline{X}-X_i\mu-\overline{X}^2+\mu\overline{X})\\
&=\mathbb{E}(X_i\overline{X}-\overline{X}^2)\\
&=\frac{\sigma^2+\mu^2}{n}+\frac{n-1}{n}\mu^2-\frac{\sigma^2}{n}-\mu^2\\
&=0.
\end{aligned}
$$
This proved the 2nd point.
3. $X_i-\mu\sim N(0,\sigma^2)$ and
$$
\begin{aligned}
V&=\sum_i\left[(X_i-\overline{X})+(\overline{X}-\mu)\right]^2/\sigma^2\\
&=\sum_i(X_i-\overline{X})^2/\sigma^2+\left(\frac{\overline{X}-\mu}{\sigma\sqrt{n}}\right)^2,
\end{aligned}
$$ 
where the first term is $(n-1)S^2/\sigma^2$ and independent with the second term. It is known that $V\sim\chi^2(n)$ and the last term is a $\chi^2(1)$ distribution. It implies that the rest is $\chi^2(n-1)$ which conclude this proof.
4. It follows by the $t$-distributions definition along with the 2nd point that $\overline{X}\perp S^2$.

The CLT along with the student's theorem play a important role in sampling inference.

## F-Distributions
---
Assuming $X\sim\chi^2(r_1)$, $Y\sim\chi^2(r_2)$, $X\perp Y$ and $F=\frac{X/r_1}{Y/r_2}$, using the similar idea we derive the $t$-distribution, the $F$ follows a $F$-distributions with PDF:
$$
f(x)=\frac{\Gamma(\frac{r_1+r_2}{2})}{\Gamma(\frac{r_1}{2})\Gamma(\frac{r_2}{2})}\frac{(r_1/r_2)^{\frac{r_1}{2}}x^{\frac{r_1}{2}-1}}{(1+\frac{r_1}{r_2}x)^{(r_1+r_2)/2}}, \quad x\in [0,+\infty).
$$