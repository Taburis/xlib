# Bayesian Statistics

Given a probability space $(\Omega, \mathcal{R},\mathbb{P})$, the Bayes' theorem should the conditional probability follows the equation
$$
\mathbb{P}(A|B)= \frac{\mathbb{P}(A\cup B)}{\mathbb{B}}, 
$$
where $A, B$ are two subsets of $\Omega$. 

Let $X$ be a random variable with PDF $f(x,\theta)$ where $\theta$ is a parameter unknow. The $\theta$ is assumed following a **prior** probability with PDF $h(\theta)$. Given a mutual indepedent samples $\lbrace X_n\rbrace$ drew from $X$, the disbtribution of the $\theta$ with the samples $p(\theta|\boldsymbol{X})$ is called **posterior distribution**. With the Bayes' theorem, the prior and posterior pdf is connected by
$$
p(\theta|\boldsymbol{X})= \frac{L(\boldsymbol{X},\theta)h(\theta)}{L(\boldsymbol{X})},\quad L(\boldsymbol{X},\theta)=\prod_{i=1}^nf(X_i,\theta),\quad L(\boldsymbol{X})=\int_\Theta L(\boldsymbol{X},\theta)d\theta
$$
where $L(\boldsymbol{X},\theta)h(\theta) = \mathbb{P}(\boldsymbol{X}, \theta)$. $L(\boldsymbol{X},\theta)$ is actually the likelihood function of the samples. The truth paramenter of $\theta$ should be the one maximize the probability $p(\theta|\boldsymbol{X})$:
$$
\hat \theta = \argmax_{\theta\in \Theta} p(\theta|\boldsymbol{X}).
$$

To find the $\hat\theta$, the term $L(\boldsymbol{X})$ is not important as it is a constant relative to $\theta$. Then the relation reduced to
$$
p(\theta|\boldsymbol{X})\propto L(\boldsymbol{X},\theta)h(\theta).
$$

Furthermore, if there's a statistics $Y=S(\boldsymbol{X})$ such that $L(\boldsymbol{X},\theta)\propto g(\boldsymbol{X}|\theta)$ where $g(y)$ is the PDF of $Y$. Then we reduced to a general form
$$
p(\theta|\boldsymbol{X})\propto g(\boldsymbol{X}|\theta)h(\theta)
$$