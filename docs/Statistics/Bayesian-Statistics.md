# Bayesian Inference
## Baysian Estimation
---
Suppose a random vector $X$ stands for a sample drawn from a model $\mathbb{P}_\theta$ with a parameter $\theta\in\Theta$ unknown, and $\hat e=e(X)$ is a statistics used to estimate the value $g(\theta)$. Then a **loss function** $\mathcal{L}(\theta,X)$ represents the loss of the value estimated from $\hat e$ comparing with the truth $g(\theta)$. It is natural to assume that $\mathcal{L}\ge 0$ and equal to 0 only if the estimation is the same as the truth. Furthermore, the **risk function** $R(\theta,e)$ defined as
$$
R(\theta)=\mathbb{E}\mathcal{L}(\theta,X),
$$
which is the estimation of the loss over all the samples. 

However, in the Baysian statistics, the possible value of the parameter $\theta$ is assumed to follow a **prior distribution**, say $\theta\sim h$. Then the model $\mathbb{P}_\theta$ is consided as a conditional probability with the parameter is given, $\mathbb{P}(X|\theta)$ and the estimator is the one minimizing the following expectation:
$$
\int_\Theta R(\theta)h(\theta)d\theta.
$$
This type of estimator is called **Bayes**.

More formally, let $(\Omega, \mathcal{R},\mathbb{P})$ be a probability space and define the following distributions: 
* $X\sim f(x|\theta)$: the random variable following a PDF with a parameter $\theta$ undetermined. 
* Prior distribution $\theta\sim h(\theta)$: an assumed distribution of the $\theta$ before observations. 
* Posterior distribution $p(\theta|\boldsymbol{X})$: the realized distribution after samples $\boldsymbol{X}$ were drawn. 

The Bayes' theorem shows
$$
\mathbb{P}(X\in\Omega|\theta=\theta_0)= \frac{\mathbb{P}(X\in\Omega, \theta=\theta_0)}{\mathbb{P}(\theta=\theta_0)}.
$$


These conceptions are connected by the Bayes' theorem as follows:
$$
p(\theta|\boldsymbol{X})= \frac{L(\boldsymbol{X},\theta)h(\theta)}{P(\boldsymbol{X})},\quad L(\boldsymbol{X},\theta)=\prod_{i=1}^nf(X_i,\theta),\quad P(\boldsymbol{X})=\int_\Theta L(\boldsymbol{X},\theta)d\theta
$$
where $L(\boldsymbol{X},\theta)h(\theta) = \mathbb{P}(\boldsymbol{X}, \theta)$. The likelihood function $L(\boldsymbol{X},\theta)$ represents the probability to obtain the sample $\boldsymbol{X}$ for a fixed parameter $\theta$. The $P(\boldsymbol{X})$ is the marginal probability.


:::tip Remark:
Commonly, the parameter with prior distribution should not be viewed as the random in the frequentist sense. It reflects the information about the parameter before observations. With the observation made, the information is improved and reshape the prior distribution to the posterior distribtion. Similarly, the prior and posterior risk can be defined, respectively.
:::


Based on the maximum likelihood assumption, the truth paramenter of $\theta$ should be the one maximize the probability $p(\theta|\boldsymbol{X})$:
$$
\hat \theta = \argmax_{\theta\in \Theta} p(\theta|\boldsymbol{X}) = \argmax_{\theta\in \Theta} L(\boldsymbol{X},\theta)h(\theta),
$$
the last equation follows from that the denominator $L(\boldsymbol{X})$ has no $\theta$-dependency. This estimation pick the peak of the $p(\theta|\boldsymbol{X})$ as the prediction of $\theta$. However, it is not the unique way to estimate parameters in Bayesian statistics. The idea of the Bayesian point estimation is to estimate the parameter $\hat\theta$ is "close" to the truth $\theta$ measured by the loss function. This idea is illustrated by the following theorem:
:::note Theorem:
Suppose $e(X)$ is a statistics used for estimating the $g(\theta)$ and
1. $\mathbb{E}\mathcal{L}(\theta,X)<\infty$ for some values of $\boldsymbol{X}$;
2. The valude $\hat e$ defined as
$$
\hat e = \argmin_{\theta\in\Theta}\int_{\Theta}\mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]d\mathbb{P}(\theta) = \argmin_{\theta\in\Theta}\mathbb{E}[\mathcal{L}(\theta,\boldsymbol{X}=\boldsymbol{x})],
$$
exists almost everywhere for $\boldsymbol{x}$.

Then $\hat e$ is a Bayes estimator for $g(\theta)$. 
:::

As an example, let's say the $g(theta)=\theta$ and the corresponding estimator is $\hat\theta$, then $\mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]d\mathbb{P}(\theta) = \mathcal{L}[\theta, \hat \theta(\boldsymbol{X})]p(\theta|\boldsymbol{X})d\theta$. Different forms of the loss function can lead to different estimation value for $\theta$. For instance, if $\mathcal{L}=(\theta-\hat\theta)^2$, it would lead to a mean value estimation $\hat\theta = \mathbb{E}(\theta)$. And $\mathcal{L}=|\theta-\hat\theta|$ would lead to the median of the $p(\theta|\boldsymbol{X})$. 

The estimation function can be interpreted as $\mathbb{E}(\mathcal{L}|\boldsymbol{X})$. By definition, it can be expressed as
$$
\int_X\mathbb{E}(\mathcal{L}|\boldsymbol{x})d\mathbb{P}(\boldsymbol{x}) = \int_X\int_\Theta\mathcal{L}(\theta,\hat\theta)p(\theta|\boldsymbol{x}) P(\boldsymbol{x})d\boldsymbol{x}d\theta=\int_X\int_\Theta\mathcal{L}(\theta,\hat\theta)L(\boldsymbol{x}|\theta)h(\theta)d\boldsymbol{x}d\theta.
$$
Then the risk function $R(\theta)$ becomes
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