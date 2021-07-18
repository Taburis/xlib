# Kernel Method

[Reference](https://arxiv.org/pdf/math/0701907.pdf)

Given a set of samples $\lbrace x_i, y_i\rbrace$, the kernel method is used to analysis the pattern of the given samples. 

## Mathematical Foundation 

Let $X$ be a set and $\mathcal{H}$ is Hilbert space of real valued function defined on $X$. It is a **reproducing kernel Hilbert space** if $L_x$ is continuous on any $f\in\mathcal{H}$ where
$$
L_x:f\to f(x),\quad \forall f\in\mathcal{H},
$$
which means for any $f,g\in\mathcal{H}$, the small of $\Vert f-g\Vert$ implies that $|f(x)-g(x)|$ is small for all $x\in X$. A reproducing kernel Hilbert space could be span by a set of kernels, where the kernel is defined as a bilinear function $K$
$$
K(x,x') = \langle \psi(x), \psi(x')\rangle,
$$
where $\psi:X\to\mathcal{H}$ forms a span as $\lbrace \psi(x_1),\dots, \psi(x_n)\rbrace$. So that any function $f\in\mathcal{H}$ can be expressed as a linear combination of kernels such that
$$
f(x)=\sum_{i}\alpha_iK(x,x_i).
$$
A **gram matrix** is defined as $K_{ij}=K(x_i,x_j)$.


## Support Vector Machine

The Support Vector Machine (SVM) is an algorithm designed for solving the classification problem. Given a set of samples $(x_i,y_i)$ where $i=1,\dots,N$ and $y_i=\pm 1$ labeling the sample into two classes, the SVM tries to find a hyper-surface to separate these classes.

### Linear Support Vector Classification

The linear support vector classifier (SVC) $f(\boldsymbol{x})$ is a simplified SVM separating classes by a hyperplane determined by $(\boldsymbol{n}, \boldsymbol{b})$:
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
