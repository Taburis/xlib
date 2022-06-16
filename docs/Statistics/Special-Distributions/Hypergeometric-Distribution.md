# Hypergeometric Distribution
Consider an experiment that draw $n$ samples (without replacement) from $N=A+B$ candidates, each candidate has two possible exclusive outcome, say positive and negative, where $A$ and $B$ stands for the total number of candidates with outcome positvie and negative, respectively. Like draw $n$ balls from a box contained $N$ balls with color blue and red. Let $M$ be a random variable stands for the total number of outcome positive, then $M$ follows a hypergeometric distribution with a probability mass function $p(m)$:
$$
p(m)= \mathbb{P}(M=m)=\frac{\binom{A}{m}\binom{N-A}{n-m}}{\binom{N}{m}}.
$$

The properties of the hypergeometric distribution are:
* **Mean** $\mathbb{E}(M)= nA/N$ and **varaince** $\text{var}(M)=n\frac{A}{N}\frac{B}{N}\frac{N-n}{N-1}$

## Asymptotic Behavior
---
Suppose the draw is replacement, the probability of outcome with a $m$ positive is also $p=A/N$. Then the distribution of $m$ follows the binomial distribution. In the limit of $N\to\infty$, the hypergeometric PMF can be written as
$$
p(m)=\frac{\binom{pN}{m}\binom{N(1-p)}{n-m}}{\binom{N}{m}}\xrightarrow{N\to\infty} \binom{n}{m}p^m(1-p)^{(n-m)}.
$$

**Proof** In fact, the leading order of binomial coefficient under the limit is
$$
\binom{N}{m}=\frac{N^m\cdot 1\cdot (1-1/N)\cdots(1-\frac{m+1}{N})}{m!}\xrightarrow{N\to\infty} \frac{N^m}{m!}.
$$
Substitute this approximation to the PMF, it follows that the binomial PMF is the limitation. 
