# Bayesian Statistics

Given a probability space $(\Omega, \mathcal{R},\mathbb{P})$, the Bayes' theorem tells that the conditional probability follows the equation:
$$
\mathbb{P}(A|B)= \frac{\mathbb{P}(A\cup B)}{\mathbb{B}}, 
$$
where $A, B$ are two subsets of $\Omega$. To state the Bayesian statistics, we define the following terms: 
* $X\sim f(x|\theta)$: the random variable following a PDF with a parameter $\theta$ undetermined. 
* Piror distribution $\theta\sim h(\theta)$: an assumed distribution of the $\theta$ before observations.
* Posterior distribution $p(\theta|\boldsymbol{X})$: the realized distribution after samples $\boldsymbol{X}$ were drawn. 

These conceptions are connected by the Bayes' theorem as follows:
$$
p(\theta|\boldsymbol{X})= \frac{L(\boldsymbol{X},\theta)h(\theta)}{P(\boldsymbol{X})},\quad L(\boldsymbol{X},\theta)=\prod_{i=1}^nf(X_i,\theta),\quad P(\boldsymbol{X})=\int_\Theta L(\boldsymbol{X},\theta)d\theta
$$
where $L(\boldsymbol{X},\theta)h(\theta) = \mathbb{P}(\boldsymbol{X}, \theta)$. The likelihood function $L(\boldsymbol{X},\theta)$ represents the probability to obtain the sample $\boldsymbol{X}$ for a fixed parameter $\theta$. The $P(\boldsymbol{X})$ is the marginal probability.

## Baysian Estimation
---
 Based on the maximum likelihood assumption, the truth paramenter of $\theta$ should be the one maximize the probability $p(\theta|\boldsymbol{X})$:
$$
\hat \theta = \argmax_{\theta\in \Theta} p(\theta|\boldsymbol{X}) = \argmax_{\theta\in \Theta} L(\boldsymbol{X},\theta)h(\theta),
$$
the last equation follows from that the denominator $L(\boldsymbol{X})$ has no $\theta$-dependency. This estimation pick the peak of the $p(\theta|\boldsymbol{X})$ as the prediction of $\theta$. However, it is not the unique way to estimate parameters in Bayesian statistics. The idea of the Bayesian point estimation is to estimate the parameter $\hat\theta$ is "close" to the truth $\theta$. And a dicision function known as the **loss function** $\mathcal{L}(\theta,\hat \theta)$ measures difference between the estimation and the truth.  

An estimator for the $\theta$ can be formulised as a function $\hat\theta = \hat\theta(\boldsymbol{X})$. The estimation error is defined by $\mathcal{L}[\theta, \hat \theta]$:
$$
\hat\theta = \argmin \int_{\Theta}\mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]d\mathbb{P}(\theta),
$$
where $\mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]d\mathbb{P}(\theta) = \mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]p(\theta|\boldsymbol{X})d\theta$. Different forms of the loss function can lead to different estimation value for $\theta$. For instance, if $\mathcal{L}=(\theta-\hat\theta)^2$, it would lead to a mean value estimation $\hat\theta = \mathbb{E}(\theta)$. And $\mathcal{L}=|\theta-\hat\theta|$ would lead to the median of the $p(\theta|\boldsymbol{X})$. 

The estimation function can be interpreted as $\mathbb{E}(\mathcal{L}|\boldsymbol{X})$. By definition, it can be expressed as
$$
\int_X\mathbb{E}(\mathcal{L}|\boldsymbol{x})d\mathbb{P}(\boldsymbol{x}) = \int_X\int_\Theta\mathcal{L}(\theta,\hat\theta)p(\theta|\boldsymbol{x}) P(\boldsymbol{x})d\boldsymbol{x}d\theta=\int_X\int_\Theta\mathcal{L}(\theta,\hat\theta)L(\boldsymbol{x}|\theta)h(\theta)d\boldsymbol{x}d\theta.
$$
By defining the **risk function** $R(\theta)$ as
$$
R(\theta) = \int_X\mathcal{L}(\theta,\hat\theta) L(\boldsymbol{x}|\theta)d\boldsymbol{x},
$$
the point estimator can be rewritten as
$$
\hat\theta = \argmin_{\theta}\mathbb{E}[R(\theta)].
$$

A **interval estimation** for parameter $\theta$ is two functions $\beta_0(\boldsymbol{x})$ and $\beta_1(\boldsymbol{x})$ and
$$
\mathbb{P}[\beta_0<\theta <\beta_1|\boldsymbol{x}] = \int_\alpha^\beta p(\theta|\boldsymbol{x})d\theta = \alpha.
$$ 
Then the interval $[\beta_0,\beta_1]$ is called the $\alpha$-credible intervals (distinguish from the confidence interval).

## Hypotheses Test
---
The hypotheses about the parameter $\theta$ can be stated as:
$$
H_0: \theta\in \Theta_0,\quad H_1: \theta\in \Theta_1,
$$
where $\Theta_1\cap\Theta_0=\varnothing$. The conclusion for the test can be made by
$$
\text{Accept }H_0: \mathbb{P}(\theta\in \Theta_0|\boldsymbol{x}) \ge \mathbb{P}(\theta\in\Theta_1|\boldsymbol{x})
$$

## Gibbs Sampling
---

Sampling theorem: 