# Conditional Probability 

## Information
---
Let $(\Omega, \mathcal{F}, \mathbb{P})$ be a probability space where $\mathcal{F}$ is a $\sigma$-algebra over the sample space $\Omega$. The probability space $\Omega$ contained all the possible samples. Samples draw from $\Omega$ have the maximum randomness, which means no information given for the samples. However, if the sample possible candidates are from a subset of $\Omega$, the randomness reduced a little bit. The most certain case is that there is only one possbile candidate for the samples, the lowest randomness. The information we have for the random space can be formulated into the $\sigma$-algebra generated from a subset of the probability space. Suppose $\mathcal{G}$ is a $\sigma$-algebra with $\mathcal{G}\subset \mathcal{F}$, it represents that information we have for $\mathcal{G}$ is no less than what we have for $\mathcal{F}$.

To illustrate the information gain in a sequential experiments, the information at the time $t\in[0,T]$ is represented by a $\sigma$-sub-algebra $\mathcal{F}(t)\subset \mathcal{F}$. The collections $\lbrace \mathcal{F}(t)|t\in[0,T]\rbrace$ is called a **filtration**. It can be used to represent the shrink of the randomness in a time sequence experiments or sampling.  

## Independence
---
Let $\mathcal{H,G}$ be two sub-$\sigma$-algebra standing for two different types of events. The two algebra are independent with each other if
$$
\mathbb{P}(A\cap B)=\mathbb{P}(A)\cdot \mathbb{P}(B),
$$
for any $A\in \mathcal{H}, B\in\mathcal{G}$. The independence between two random variables means $\sigma(X)$ is independent on $\sigma(Y)$, where $\sigma(X)$ is definted as the minimum $\sigma$-algebra for the collection $\lbrace X^{-1}(B)\rbrace$ of all Borel sets $B$. 

:::note Theorem:
Suppose random variables $X,Y$ are mutually independent, $X\perp Y$, and $f,g$ are Borel-measurable functions on $\mathbb{R}$, then:
1. the functions are mutually independent, $f(X)\perp g(Y)$.
2. The joint distribution measure is factorizable:
$$
\mu_{X\times Y}(A\times B)=\mu_X(A)\cdot \mu_Y(B)
$$
for all Borel sets $A\subset \sigma(X)$ and $B\subset\sigma(Y)$.
3. The density distribution is factorizable:
$$
f_{X\times Y}(x,y)=f_X(x)\cdot f_Y(y),\quad x,y\in \mathbb{R}.
$$
4. The correlation between $\text{Cov}(X,Y)$ is 0 if $X\perp Y$.
:::
**Proof**: 
1. For any Borel set $B$, the inverse image $f^{-1}(B)$ is a Borel set for any Borel-measurable function $f$. It means that $\sigma(f(X))\subset \sigma(X)$ and hence the independency follows.
2. This conclusion is a direct conclusion from the definition of the condition.
3. Since $\text{Cov}=\mathbb{E}(XY)-\mathbb{E}X\cdot\mathbb{E}Y$, and the independence means that the right side of the equation is equal. 

However, the inverse of the 3rd point is not true.

## Joint and Marginal Densities
---
Given two random variables $X\in \mathbb{R}^m$ and $Y\in\mathbb{R}^n$ with measure $\mu_X$ and $\mu_Y$, respectively. The direct product random variable $Z=(X,Y)\in\mathbb{R}^{m+n}$ has the measure $\mu_Z=\mu_X\times \mu_Y$ leads to a denstiy $f_Z=d\mu_Z/dz$ called the **joint density** of $X$ and $Y$. Then 

Let the $1_S(x)$ be the support function for the set $S$ that the function to be 1 for $x\in S$ and 0 elsewhere. The probability of $Z\in S$ is defined as
$$
\mathbb{P}(Z\in S) = \int_{R^{m+n}}1_S(x,y)f_Z(x,y)dxdy,
$$
The probability of $X\in A$ is given by
$$
\mathbb{P}(X\in A) = \mathbb{P}(Z\in A\times R^n) = \int_{R^m}\int_{R^n} 1_{A\times R^{n}}(x,y) f_Z(x,y)dxdy=\int_A\left\lbrace\int_{R^n}f_Z(x,y)dy\right\rbrace dx.
$$
Likewise, the density of $X$ is defined as
$$
f_X(x)=\int f_Z(x,y)dy,
$$
is called the **marginal density**. The same definition can be derived for $Y$ as well.




## Conditional Distributions and Expectation
---

Let $(\Omega, \mathcal{F}, \mathbb{P})$ be a probability space. The conditional probability of $Y\in B$ given $X=x$ is defined as
$$
\mathbb{P}(Y\in B|X =x)=\frac{\mathbb{P}(Y\in B, X=x)}{\mathbb{P}(X=x)},
$$
where $B$ is a Borel set. This defined a probability measure for set $B$ with respect to the value $x$, which is known as the **conditional denstiy**. A function  $\mathbb{P}(Y|X)$ is the coniditional density for $Y$ given $X$ if:
1. $\mathbb{P}(B\in\mathcal{F}|X=x)$ is a probability measure for all $x$;
2. $\mathbb{P}(B|X=x)$ is a measurable function for any Borel set $B\in \mathcal{F}$;
3. For any Borel sets $A,B$,
$$
\mathbb{P}(X\in A, Y\in B)=\int_A \mathbb{P}(B|X=x)d\mathbb{P}(x).
$$
The conditional distribution exists and can be obtained explicity for two random vectors $X,Y$ by the formula:
$$
p_{Y|X}(y|x)=\frac{f_{Z}(x,y)}{f_X(x)},
$$
where $Z=(X,Y)$ and $f_Z$ is the joint distribution for $Z$ and $f_X$ is the marginal distribution for $X$.

Furthermore, given a random vector $X$ as a time series where $x_k$ is the $k$-th step in the series. The probability of $X$ can be expanded into the production of the conditional distributions
$$
p_X(x_1,\dots,x_n) = p(x_1)p(x_2|x_1)p(x_3|x_1,x_2)\dots p(x_n|x_1,\dots, x_{n-1}).
$$
This formula can be used to build models easily since the conditional probability is easy to model in some cases.

On the other hand, ff the $B$ is all the probability space, then the kernel $\mathbb{P}(B|X=x)$ above becomes the expectation of $Y$ with a condition that $X=x$. This leads to the definition of the conditional expectation below.

### Conditional Expectation
Let $(\Omega, \mathcal{F}, \mathbb{P})$ be a probability space, $\mathcal{G}\subset \mathcal{F}$ is a sub-$\sigma$-algebra, and $X$ is a integrable random variable. The conditional expectation of $X$ given condition $\mathcal{G}$, denoted as $\mathbb{E}[X|\mathcal{G}]$, is the random variable satisfying
1. $\mathbb{E}[X|\mathcal{G}]$ is $\mathcal{G}$-measurable;
2. The following euqation holds
$$
\int_A\mathbb{E}[X|\mathcal{G}](\omega)d\mathbb{P}(\omega) = \int_A X(\omega)d\mathbb{P}(\omega),\quad \forall A\in \mathcal{G}.
$$

Here are the **properties** about the conditional expectation
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