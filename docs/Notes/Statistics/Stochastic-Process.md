# Stochastic Process

## Definition

A sequence of $\sigma$-algebra $\mathfrak{F}(t)$, ordered by a parameter $t$ is called a **Filtration** if $\mathfrak{F}(s)\subset \mathfrak{F}(t)$ for all $s<t$. A sequence of random variables $X(t)$ where $t\in[0,T]$ is called an **Adapted Stochastic Process** if $X(t)$ is $\mathfrak{F}(t)$-measurable for $t\in[0,T]$. 

* A **Martingale** process is a stochastic process $X(t)$ satisfying 

$$
\mathbb{E}[X(t)|\mathfrak{F}(s)] = X(s), \quad \forall s<t.
$$

* A **Markov Process** $X(t)$ is a stochastic process that 

$$
\mathbb{E}[f(X(t))|\mathfrak{F}(s)] = g(X(s)), \quad \forall s<t,
$$
where $f,g$ are both Borel-measurable functions. A Markov process is also a Martingale. 

## Hidden Markove Model

Let $X(t), Y(t)$ are two discrete-time stochastic processes and the pair $(X_t, Y_t)$, where $(t=1,\dots, n)$, forms a **Hidden Markov Model** (HMM) if 

1. X(t) is a Markov process;
2. $P[Y_t\in A|X_1=x_1,\dots, X_t=x_t] = P[Y_t\in A|X_t=x_t]$ for any $1\le t\le n$ and an arbitrary measurable set $A$.

The states $X_t$ is called hidden states, $Y_t$ is called observable, and $P(Y_t\in A|X_n=x_n)$ is called emission probability. By the Possion's Law, we have
$P(Y) = \sum_xP(Y|X=x)$.