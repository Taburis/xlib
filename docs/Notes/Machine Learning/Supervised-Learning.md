# Supervised Learning

Given a probability space $(\Omega,\mathcal{R}, \mathbb{P})$, where $\Omega$ stands for the events of interests. 
* The features are random variables $X:\Omega\to \mathbb{R}^F$ where $F$ is known as the feature size. Features are coordinates for events in feature space, and the $i$-th component is denoted as $X_i$. 
* Labels are random variables $Y:\Omega\to \mathbb{R}^L$ and $L$ is the dimension labels. 
* Samples $S$ is a set of pairs $(X_i,Y_i)$, $i=1,\dots, m$ where $m$ is known as sample size.
* A relationship is a function $f:\mathbb{R}^F\to\mathbb{R}^L$ building connection between the features and labels. The sample $S$ should be, at least, the support of the function $f$.
* Objective function or Loss function $L_S(f)$ is a functional evaluated on the supports which is the sample $S$.

Finding the relationship $f$ is the ultimate goal for statistical learning. For supervised learning, the labels are given, or known in samples. The functional $L_S(f)$ measured the performance of $f$ under the support $S$ so that it is usually defined as
$$
L_S(f)=\int_{S\subset \Omega} \mathcal{L}(Y,f\circ X)(\omega) d\mathbb{P}(\omega),
$$
where $(f\circ X)(\omega) =f[X(\omega)]$, and $\mathcal{L}$ is the loss measure for the difference between the labels and features. This form reduced to summation form for the discrete samples. The minimization of $L_S$ leads to the require that 
$$
\delta L_S(f)= \int_S\frac{\delta\mathcal{L}}{\delta f}\delta (f\circ X) (\omega)d\mathbb{P}(\omega)=0,
$$
which requires the equation
$$
\frac{\delta \mathcal{L}}{\delta f}=0.
$$
The whole machine learning is about to solve this equation for the given loss function $\mathcal{L}$ and the support $S$.

## Loss Functions

Loss functions are supposed to quantify the difference between the predicted labels $\hat Y = f(X)$ and the truth $Y$. On the other side, it could also be used to evaluate the model to prevent the over fitting. 

### Variance

For continuous labels $Y$, the variance $\text{Var}(Y-f(X))$ is used as the error function:
$$
\text{Var}(Y-\hat Y) = \int_S(Y-\hat Y)^2(\omega)d\mathbb{P}(\omega)
$$