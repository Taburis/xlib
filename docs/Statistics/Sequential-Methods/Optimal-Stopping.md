# Optimal Stopping Problem

## Stopping Times

Suppose the $sigma$-algebra $\mathcal{F}_n, n=1,\dots$ forms a filtration for a sequential observations. A random variable $N$ could take values in $\lbrace 1,2,\dots, \infty\rbrace$ is called a stopping time with respect to the filtration $\mathcal{F}_n$ if 
$$
\lbrace N=n\rbrace \in \mathcal{F}_n,\quad \forall n\ge 0.
$$
Usually, a stopping time indicates a stop condition satisfied after some observations have been taken. Hence it should be a function of other random variables. Moreover, a stopping time can also be represented as disjoint union of the event $A\in \mathcal{F}$ which is
$$
\left(\lbrace N=n\rbrace \cap A\right)\in \mathcal{F}_n,\quad  \forall n\ge 0.
$$
It also implies that $N$ is $\mathcal{F}_N$ measurable but $N+1$ is not.[section](/Statistics/Special-Distributions/Normal-Distributions#Multivariate-Normal-Distribution)

:::note Wald's Fundamental Identity:
Suppose 
:::