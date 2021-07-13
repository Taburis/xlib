# Kernel Method

[Reference](https://arxiv.org/pdf/math/0701907.pdf)

Given a set of samples $\lbrace x_i, y_i\rbrace$, the kernel method is used to analysis the pattern of the given samples by comparing the samples with the inputs. 

## Kernel 

A kernel is defined as a function $K$ that
$$
K(x,x') = \langle \Psi(x), \Psi(x')\rangle,
$$
where $\Psi(x)$ is called a feature map, and $\langle \cdot ,\cdot \rangle$ is a linear inner product whose definition is model-dependent. This inner production measures the similarity between the features and patterns. 



## Support Vector Machine

The Support Vector Machine (SVM) is an algorithm designed for solving the classification problem. Given a set of samples $(x_i,y_i)$ where $i=1,\dots,N$ and $y_i=\pm 1$ labeling the sample into two classes, the SVM tries to find a hyper-surface to separate these classes.

### Linear Support Vector Machine

The linear SVM $f(\boldsymbol{x})$ is a simplification of SVM separating classes by a hyperplane determined by $(\boldsymbol{n}, \boldsymbol{b})$:
$$
\boldsymbol{n}\cdot\boldsymbol{x}-b =0,
$$
where $\boldsymbol{n}$ is the vector normal to the hyperplane, and $\boldsymbol{b}$ is the offset. Noticing that for any $\boldsymbol{x}$, the quantity $d/\Vert \boldsymbol{n}\Vert$ defines the distance of $\boldsymbol{x}$ to the hyperplane, where
$$
d=\boldsymbol{n}\cdot\boldsymbol{x}-b,
$$
and the sign of $d$ indicates which side the point is located. The linear SVM $f(\boldsymbol{x}|\boldsymbol{n},b)$ is defined as
$$
f(\boldsymbol{x}) = \text{sgn}(d),
$$
which is find a hyperplane $(\boldsymbol{n},b)$ such that all the samples labeled by $y=1$ appears on one side of it and the samples with $y=-1$ on the other side. A separable condition is that $yf(\boldsymbol{x})\ge 1$ for all $(\boldsymbol{x},y)$ in samples. One definition of the hyperplane is the parameters $(\boldsymbol{n},b)$ minimizing the loss
$$
\mathcal{L}(\boldsymbol{n},b)=\frac{1}{N}\sum_{i=1}^N\text{max}[0, c-y_i(\boldsymbol{n}\cdot \boldsymbol{x}_i-b)]+\lambda\Vert \boldsymbol{n}\Vert^2,
$$
where $c$ is a margin of the hyperplane and the last term is a L2 regularization. Due to the separable condition, only the mislabeled class may contribute to this loss (if $c=0$). The fit minimizing this loss by gradient descend will ping down the parameters of the hyperplane.
