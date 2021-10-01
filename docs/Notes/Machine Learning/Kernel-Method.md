# Kernel Method
---
[Reference](https://arxiv.org/pdf/math/0701907.pdf)

Given a set of samples $\lbrace x_i, y_i\rbrace$, the kernel method is used to analysis the pattern of the given samples. 

## Mathematical Foundation 
---

### Reproducing Kernel Hilbert Space

A Reproducing Kernel Hilbert Space (RKHS) $\mathcal{H}(X)$ is a Hilbert space consists by functions defined on $X$ such that $L_x(f)$ is continuous (or, equivalently, bouned) for any $f\in\mathcal{H}$, where $L_x:f\to f(x)$ is a linear operator known as **linear evalutaion functional**. The continuous of $L_x$ means that the small $\Vert f-g\Vert$ implies that $|f(x)-g(x)|$ is small for all $x\in X$. 

:::info Riesz Representation Theorem:
Let $H$ be a Hilbert space whose inner product $\langle x, y\rangle$ is linear for the second argument ($y$) and anti-linear for the first argument ($x$), then for every continuous linear functional $\varphi\in H^*$, there is a unique $f_\varphi\in H$ such that $\varphi(f)=\langle f_\varphi, f\rangle$. 
:::

Based on this thoerem, the linear evalutaion functional can be expressed as
$$
f(x)=L_x(f)=\langle K_x,f\rangle, \quad \forall f\in \mathcal{H}(X).
$$
Moreover, $L_y(K_x)=\langle K_y,K_x \rangle=K(x,y)$ and called $K(x,y)$ as the **reproducing kenerl** and $K:X\times X\to \mathbb{F}$. It is symmetric and positive defined. And the point evaluation of any function $f\in\mathcal{H}(X)$ can be expressed as 
$$
f(y)=\sum_i\alpha_iK(y,x_i),
$$ 
where $\alpha_i$ is called the coordinate of the function $f$ under the kernel and $x_i$ is the basis of the support of all functions in $\mathcal{H}(X)$.

:::info Moore-Aronszajn Theorem:
Suppose $K$ is a postive defined symmetry kernel on $X$, then there is a unique Hilbert space of functions on $X$ such that $K$ is a reproducing kernel.
:::

A **gram matrix** is defined as $K_{ij}=K(x_i,x_j)$. If the kernel is positive defined, then $K_{ij}$ is positive defined. 

### Construct a RKHS

A RKHS kernel can be constructed for a given set $X$ from a map $\mathcal{F}(X,\mathbb{F})$, where $\mathcal{F}(X,\mathbb{F})$ is the Hilbert space of functions mapping $X\to\mathbb{F}$. The set $X$ is support of functions in $\mathcal{F}(X,\mathbb{F})$ and hence any element $f,g\in\mathcal{F}(X,\mathbb{F})$ can be expressed as linear combination by the kernels $\Psi(x_i)=K_{x_i}$:
$$
f=\sum_i\alpha_iK_{x_i},\quad g=\sum_j\beta_jK_{x_j},
$$ 
and the inner product for this RKHS is defined as
$$
\langle f,g\rangle = \sum_{ij}\alpha_i\beta_jK(x_i,x_j)=\sum_{ij}\alpha_i\beta_j\langle \Psi(x_i),\Psi(x_j)\rangle.
$$

### Example of Kernels

#### Feature Map
A feature map is a map $\Psi:X\to H$ where $H$ is a Hilbert space called feature space. The kernel is defined by
$$
K(x,y)=\langle \Psi(x),\Psi(y)\rangle_H,
$$
and the symmetry and positive definiteness folllows from the inner production in space $H$. 

#### Polynomial Kernels
A order of $p$-polynomial $a=\sum_{n=1}^pa_nx^n$ can be expressed as a vector $\boldsymbol{a}=(a_0,\dots,a_p)$, a polynomial kernel $K_p(\boldsymbol{x},\boldsymbol{y})$ is defined as
$$
K_p(\boldsymbol{x},\boldsymbol{y})=(\boldsymbol{x}^T\cdot\boldsymbol{y}+c)^p,\quad c\ge 0.
$$

#### Radial Bias Functions
Some kernels can be expressed as
$$
K(x,y)=h(x-y),
$$
and $h$ is called positive defined if the corresponding Gram matrix is positive defined. The Radial Bias Functions (RBF) are functions $h$ can be written as $h(x)=g(\Vert x\Vert)$ where $g:[0,+\infty]\to \mathbb{R}$. 

The example of the RBF kernel: 
* **Gaussian kernel**:
$$
K(\boldsymbol{x},\boldsymbol{y})= \exp{-\frac{\Vert \boldsymbol{x}-\boldsymbol{y}\Vert^2}{2\sigma^2}},
$$
* **Laplace kernel**:
$$
K(\boldsymbol{x},\boldsymbol{y})=\exp{-\alpha\Vert \boldsymbol{x}-\boldsymbol{y}\Vert},\quad \alpha>0
$$


## Support Vector Machine
---

The Support Vector Machine (SVM) is an algorithm designed for solving the classification problem. Given a set of samples $(x_i,y_i)$ where $i=1,\dots,N$ and $y_i=\pm 1$ labeling the sample into two classes, the SVM tries to find a hyper-surface to separate these classes.

### Linear Support Vector Classification

The linear support vector classifier (SVC) $f(\boldsymbol{x})$ is a simplified SVM separating classes by a hyperplane determined by $(\boldsymbol{n}, b)$, where $\boldsymbol{n}$ is the vector normal to the hyperplane and $b$ is a shift:
$$
d(\boldsymbol{x},\boldsymbol{n},b) =0,
$$
where $\boldsymbol{n}$ is the vector normal to the hyperplane, $\boldsymbol{b}$ is the offset, and 
$$
d(\boldsymbol{x},\boldsymbol{n},b)=\boldsymbol{n}\cdot\boldsymbol{x}-b.
$$
Noticing that for a given sample $\boldsymbol{x}$, the distance from the itself to the hyperplane is given by $d(\boldsymbol{x},\boldsymbol{n},b)/\Vert \boldsymbol{n}\Vert$ and the sign of $d$ indicates which side the point is located, the samples labeled by $y=1$ should be located on the side 
$d(\boldsymbol{n},b) \ge 1$, where 1 (it could be any constant) relates to margin of the hyperplane. Then the samples labeled by $y=-1$ should be on the other side defined as
$d(\boldsymbol{n},b)\le 1$. The margin width is $2/\Vert\boldsymbol{n}\Vert$.

The linear SVM $f(\boldsymbol{x}|\boldsymbol{n},b)$ is defined as
$$
f(\boldsymbol{x}) = \Theta(d(\boldsymbol{x},\boldsymbol{n},b)-1)-\Theta(d(\boldsymbol{x},\boldsymbol{n},b)+1),
$$
where $\Theta(x)$ is a step function vanish for $x<0$ and $\Theta(x)=1$ for $x\ge 0$. 

Given a sample $(\boldsymbol{x},y)$, a correct classification will have $yf(\boldsymbol{x})\ge 1$ while the mis-classification leads to $yf(\boldsymbol{x})\le -1$. The loss function of the linear SVM can be constructed as
$$
\mathcal{L}(\boldsymbol{n},b)=\sum_{i=1}^N-y_if(\boldsymbol{x})+\lambda\Vert \boldsymbol{n}\Vert^2/4,
$$
where $\lambda$ is an arbitrary parameter and the second term is the square of the inverse of the margin width $2/\Vert\boldsymbol{n}\Vert$. The fit minimizing this loss by gradient descend will ping down the parameters of the hyperplane. The loss function $\mathcal{L}$ can be viewed as maximizing the margin width (second term) with a constraint on the hyperplane (first term) by using a Lagrange multiplier method. On the other hand, it could also viewed as a cross-entropy with a L2 regularization.


### Linear Support Vector Regression

Unlike the linear SVC trying to separate the samples outside the hyperplane margin, the linear support vector regression (SVR) are trying to find a hyperplane with a boundary to contain as much samples as possible. In this case, the target value $y$ is represented by the offset of the hyperplane and
$$
f(\boldsymbol{x})=\langle \boldsymbol{n},\boldsymbol{x}\rangle+b,
$$
which is the same form as a linear regression. Furthermore, we are expecting
$$
|f(\boldsymbol{x}_i)-y_i|\le \epsilon+\xi_i,
$$
where $2\epsilon/\Vert\boldsymbol{n}\Vert$ is the width of the boundary and $\xi_i$ is a tolerance for $i$-th samples. The hyperplane is determined by minimizing the loss
$$
\mathcal{L} = \sum_i\varphi(y_i-f(\boldsymbol{x}_i))+\lambda\Vert \boldsymbol{n}\Vert^2,
$$
where the $\varphi$ is a loss function customizable. In the case of $\varphi(x)=x^2$, it leads to a least square estimation with a L2 regularization. It could also be $\varphi(x)=\text{max}(0, |x|-\epsilon)$ that penalty only comes from the points outside the boundary. 

### Nonlinear Extension: Kernel Trick


## Kernel Models

A kernel model is the model $f(x)$ with form
$$
f(x)=\sum_{i=1}^D\alpha_i K(x, x_i),
$$
where $D$ stands for the kernel dimensions and $x_i$ can be the support of the kernel $K(x,y)$. A kernel model may or may not needs to be trained. For some cases, the model is the samples itself, like $k$-mean. Others might need to extract the supports from the samples, like the mixture models.

### Density Estimation

Given the samples $\lbrace X_n\rbrace$ drawn from a random variable $X$ with PDF $f(X)$, a natrual point estimation of $f(x)$ comes from a extension of empirical PDF:
$$
\hat f(x)= \frac{|\lbrace X_i\in B_\delta (x)\rbrace|}{n\cdot V[B_\delta(x)]},
$$
where $B_\delta (x)$ is a ball with radius $\delta$ centered at $x$, $|\lbrace X_i\in B_\delta (x)\rbrace|$ is the number of samples falling in the ball, $V[B_\delta(x)]$ is the volumn of the ball, and $n$ is the sample size. This estimation can be smoothed with a kernel estimation:
$$
\hat f(x)=\frac{1}{n\lambda}\sum_{i=1}^nK_\lambda(x,x_i),
$$
where $K_\lambda$ is a kernel with a parameter $\lambda$ as the size. For instance, if we adpat a normal Gaussian as the kernel, we have
$$
\begin{aligned}
\hat f(x)&=\frac{1}{n}\sum_{i=1}^n\varphi_\lambda(x-x_i),\\
&=(\hat f_n\circ \varphi_\lambda)(x)
\end{aligned}
$$
where $\varphi(x)$ stands for a Gaussian with width $\lambda$, $\hat f_n(x)$ is the empirical distribution with szie $n$, and the $\hat f_n\circ \varphi_\lambda$ stands for the convolution between $\hat f_n$ and $\varphi(x)$. No training is needed for this estimator.

### Gaussian Mixture Model

One of most popular mixture models is the Gaussian mixture model (GMM) $f(x)$ that
$$
f(x)=\sum_{i=1}^D\alpha_i\phi(x, \mu_i,\Sigma_i),\quad \sum_{i=1}^D\alpha_i = 1.
$$
where $\phi(x,\mu_i,\Sigma_i)$ is the a (multivariate) Gaussian with mean and varaince $\mu_i, \Sigma_i$. These variables $\mu_i,\Sigma_i$ not known is called **latent variables**. If the covariance matrix can be expressed as $\Sigma_i=\sigma_i\cdot I$, then the Gaussian Mixture Model reduced to the kernel of Radial Bias Functions. 