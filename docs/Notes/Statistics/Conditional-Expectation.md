# Conditional Expectation

## Definition

Let $(\Omega, \mathcal{F}, \mathbb{P})$ be a probability space, $\mathcal{G}\subset \mathcal{F}$ is a sub-$\sigma$-algebra, and $X$ is a integrable random variable. The conditional expectation of $X$ given condition $\mathcal{G}$, denoted as $\mathbb{E}[X|\mathcal{G}]$, is the random variable satisfying
1. $\mathbb{E}[X|\mathcal{G}]$ is $\mathcal{G}$-measurable;
2. The following euqation holds
$$
\int_A\mathbb{E}[X|\mathcal{G}](\omega)d\mathbb{P}(\omega) = \int_A X(\omega)d\mathbb{P}(\omega),\quad \forall A\in \mathcal{G}.
$$

## Properties

Here are the properties about the conditional expectation
1. A conditional expection is always exists and is uniquely defined almost surely (The measurement of the set of different points is 0).
2. Given two $\mathcal{G}$-measurable random variables $X,Y$ and $a,b\in\mathbb{R}$, then
$$
\mathbb{E}[aX+bY|\mathcal{G}]=a\mathbb{E}[X|\mathcal{G}]+b\mathbb{E}[Y|\mathcal{G}].
$$
3. $X,Y$ are integrable and $X$ is $\mathcal{G}$-measurable, then
$$
\mathbb{E}[XY|\mathcal{G}] = X\mathbb{E}[Y|\mathcal{G}].
$$
4. Suppose $\mathcal{H}$ is a $\sigma$-algebra and $\mathcal{H}\subset \mathcal{G}$ and $X$ is integrable, then
$$
\mathbb{E}\big[\mathbb{E}[X|\mathcal{G}]\big | \mathcal{H}\big]=\mathbb{E}[X|\mathcal{H}].
$$
5. If $X$ is a integrable random variable that independent to $\mathcal{G}$, then
$$
\mathbb{E}[X|\mathcal{G}]=\mathbb{E}X.
$$
6. Let $\varphi(x)$ be a convex function, $\varphi(tx+(1-t)y)\le t\varphi(x)+(1-t)\varphi(y)$, then 
$$
\mathbb{E}[\varphi(X)|\mathcal{G}]\ge \varphi(\mathbb{E}[X|\mathcal{G}]).
$$

**proof**
1. The conditional expectation is actually the Radon-Nikodym derivative, therefore, the existence is guaranteed by the Radon-Nikodym's theorem. For the almost surely uniqueness, lets assume $Y,Z$ are two random variables could be $\mathbb{E}[X|\mathcal{G}]$ and $A=\lbrace p\in \Omega| Y(p)>Z(p)\rbrace$. By the definition:
$$
\int_A(Y-Z)(\omega)d\mathbb{P}(\omega) = 0,
$$
which can only hold if $\mathbb{P}(A)=0$.
2. This is a straightforward derivation from the linearity of the Lebesgue integral.
3. Noticing that $X$ can be approximated by a series of simple functions consistent by indicator $\mathcal{I}_U$ that is $1$ if $x\in U$ and 0 otherwise. Then it is enough to notice that
$$
\int_A\mathcal{I}_B\mathbb{E}[Y|\mathcal{G}](\omega )d\mathbb{P}(\omega)=\int_{A\cap B}\mathbb{E}[Y|\mathcal{G}](\omega )\mathbb{P}(\omega)=\int_AX(\omega)Y(\omega)d\mathbb{P}(\omega),
$$ 
where the last step comes from the fact that $\mathbb{P}(A\cap B) =\mathbb{P}(A)\cdot\mathbb{P}(B)$ for independence.
4. Since $\mathcal{H}\subset\mathcal{G}\subset\mathcal{F}$, then for any $A\in\mathcal{H}$, it is also in $\mathcal{G},\mathcal{H}$, which means:
$$
\int_A \mathbb{E}\big[\mathbb{E}[X|\mathcal{G}]\big | \mathcal{H}\big](\omega)d\mathbb{P}(\omega ) =\int_A\mathbb{E}[X|\mathcal{G}](\omega)d\mathbb{P}(\omega ) = \int_AX(\omega)d\mathbb{P}(\omega),
$$
other the other side of the equation, we have
$$
\int_A \mathbb{E}\big[\mathbb{E}[X|\mathcal{H}]\big | \mathcal{H}\big](\omega)d\mathbb{P}(\omega ) = \int_AX(\omega)d\mathbb{P}(\omega)
$$
which conclude this theorem.
5. Assuming $X$ is an indicator function with support $B$, by the definition we have
$$
\int_A\mathbb{E}[X|\mathcal{G}](\omega)d\mathbb{P}(\omega) = \int_A X(\omega)d\mathbb{P}(\omega)=\mathbb{P}(A\cap B)=\mathbb{P}(A)\mathbb{E}X = \int_A\mathbb{E}X d\mathbb{P}(\omega).
$$
6. Without lose generality, here we assume the value of $X$ is in $[0,1]$. Based on the definition of a convex function, it can be extended into the following equation hold for any convex function $\varphi$:
$$
\varphi\left(\sum_{i=1}^Nt_ix_i\right) \le \sum_{i=1}^Nt_i\varphi(x_i), \quad \sum_{i=1}^Nt_i = 1,
$$
On the other hand, a random variable can be approximated by a simple function $X_n=\sum_{i=1}^n\mathbb{P}(A_{i,n})i/n$, where $A_{i,n} = X^{-1}[i/n,(i+1)/n]$. Then inequality holds by inductions. 