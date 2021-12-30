# Boostrap Methods

## Boostrap Estimation Procedure
---
Given a $x_1,\dots, x_n$ draw for a random variable $X$ with the CDF $F$, the empirical cumulative distribution function (ECDF) $\hat F_n(x)$ is defined as
$$
\hat F(x)=\frac{1}{n}\sum_{i=1}^nI(X_i\le x),
$$
where $I(X\le x)$ is an indicator function for the set satisfying the condition $X\le x$. By the law of the large numbers, $\hat F_n\xrightarrow{P}F$, the ECDF is a reasonable approximation to the CDF. The new samples $\lbrace X_n^*\rbrace$ draw from the ECDF is called the **boostrap samples** (which equivalent to the resampling with replacement to the original samples). The distribution of the boostrap samples can be think of a conditional probability
$$
\mathbb{P}(X^*<x|\boldsymbol{X}=\boldsymbol{x}) = \hat F_n(X^*<x). 
$$

In boostrap estimation procedure, the estimation $\hat\theta = \hat \theta(\boldsymbol{X})$ from the original sample $\boldsymbol{X}$ is used as the underlying truth and the error is studied with the estimation $\hat\theta^*=\hat\theta(\boldsymbol{X}^*)$ from the boostrap samples $\boldsymbol{X}^*$. 

## Bias Reduction
---
As a sample augmentation method, the boostrap samples can be used to estimate the bias of an estimator $\hat \theta$ for a parameter $\theta$. This estimation bias is measured by
$$
b = \mathbb{E}(\hat \theta-\theta).
$$

This bias can also be measured in the boostrap samples 
$$
\hat b =\mathbb{E}[\hat\theta-\hat\theta^*|\boldsymbol{X}=\boldsymbol{x}],
$$
where $\hat \theta$ stands for the estimator applied to the boostrap samples. Then the new estimator $\hat \theta -\hat b$ should have less bias. Let $\hat\theta^*_i$ detnotes the estimation from the $i$-th boostrap sample $\boldsymbol{x}^*_i$, then the average 
$$
\frac{1}{M}\sum_{i=1}^M(\hat \theta^*_i-\hat\theta)
$$
should well approximated to the $b$ when the $M$ is large according to the large number theorem.

## Percentile Bootstrap Confidence Intervals
---
Let $\lbrace X_n\rbrace$ be a sample with size $n$ and $\hat \theta(\boldsymbol{X})$ is an estimator, the bootstrap confidence interval is obtained from the following procedure:
1. Draw a size $n$ sample $\lbrace X_n^*\rbrace$ from the sample $\lbrace X_n\rbrace$ with replacement.
2. Obtained the estimated value $\hat \theta = \hat \theta(\boldsymbol{X}^*)$.
3. Repeat the step 1 and 2 $N$ times to form the order statistics $\hat\theta^*_1\le \dots\le\hat \theta_N^*$, and $m=\lfloor N\alpha/2\rfloor$, then the interval:
$$
(\hat \theta^*_m, \hat\theta^*_{N+1-m})
$$
is the $\alpha/2\times 100\%$ confidence interval for $\theta$. $N$ should be a large number ($N >3000$ usually).
