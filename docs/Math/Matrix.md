# Matrix

## Linear Transform

A production of a vector from the right side of a $m\times n$ (row-column number) matrix $A$ defined on a field $F$ is a linear transformation
$$
A:\mathbb{F}^n\to \mathcal{C}(A)\subset\mathbb{F}^m,
$$
where $\mathcal{C}(A)$ is the column space of $A$. Since the matrix can be expressed as $A=(\boldsymbol{a}_1,\dots,\boldsymbol{a}_n)$, and the transform $A\boldsymbol{x}$ can be epxressed as
$$
A\boldsymbol{x}=\sum_{i=1}^nx_i\boldsymbol{a}_i,\quad x=(x1,\dots, x_n).
$$


* The **null space** of $A$, denoted as $\text{Null}(A)$ is the space that $A\boldsymbol{x}=0$, $\forall \boldsymbol{x}\in\text{Null}(A)$. It is a complementary space of the column space of $A$
$$
\mathbb{F}^n=\mathcal{C}(A)+\text{Null}(A^*).
$$


:::tip Isomorphism
The linear transformation is also known as homomorphism. And it is isomorphism if the dimension of $\text{Null}(A)$ is 0.
:::

### Production of Matrices

The production of two matrix $A_{m\times n}$ and $B_{n\times h}$ is defined as
$$
C_{ij}=\sum_{k=1}^{n}A_{ik}B_{kj},
$$
so that $C$ is a $m\times h$ matrix. It is obviously that $\mathcal{C}(C)\subset \mathcal{C}(A)$. On the other, if we represent
$$
A=(\boldsymbol{a}_1,\dots,\boldsymbol{a}_k), \quad C=(\boldsymbol{c}_1,\dots,\boldsymbol{c}_k),
$$
then $\mathcal{C}(C)\subset \mathcal{C}(A)$ means that $\boldsymbol{c}_i=\sum_{k}b_{ik}\boldsymbol{a}_k$ so that $b_{ik}$ forms a matrix that $AB=C$. It leads to the theorem that for a two matrix $A$ and $C$, $\mathcal{C}(C)\subset \mathcal{C}(A)$ if and only if there exists a matrix $B$ such that $AB=C$. 

### Invariant Space and Eigenvectors
A **invariant vector space** $V$ for linear transform $A_{m\times m}$ is a space that $A\boldsymbol{x}\in V$ for all $\boldsymbol{x}\in V$. In the case that the invariant space dimension is 1, it means that any vectors $\boldsymbol{v}\in V$ can only be scaled by a constant factor after the transformation: $A\boldsymbol{v}=\lambda \boldsymbol{v}$. Those vectors are called **eigenvectors**  with a **eigenvalue** $\lambda$. The eigenvalues of a matrix $A$ is determined by the equation
$$
\det(A-\lambda I_m)=0.
$$
Eigenvectors with distinct eigenvalues are linear independent.

The set of all eigenvalues of $A$ is also known as the **spectra** of $A$. 

:::info Theorems:
1. The existence of the solution of the eigen-equation is determined by the properties of the field $\mathbb{F}$ where $A$ is defined. For a algebra complete field such as $\mathbb{C}$, the eigenvalues are always exists for any square matrix $A$. This can be rephrasing into the form that any linear operators defined on $\mathcal{C}$ (or $\mathcal{R}$) will have 1-dimensional (2-dimensional) invariant subspace.
2. Eigenvectors exists for any symmetric matrix $A$ defined on $\mathbb{R}$.
:::

The 2nd theorem is a strength version of the 1st theorem for real symmetric matrices. Based on the 1st theorem, it is guaranteed that 2-dimensional invariant spaces $L$ exist for $A$ and it could be expressed as
$$
A_L=\left(
\begin{array}{cc}
a & b\\
b & c
\end{array}
\right).
$$
For the eigenvalue equation $(a-\lambda)(c-\lambda)-b^2=0$, we have $(a-c)^2+4b^2\ge 0$ which means that real solutions exists for the $\det(A_L)$ so that this 2-dimensional invariant space contained two 1-dimensional invariant spaces.



:::tip Geometric interpretation
A transform defined by a $m\times m$ matrix is combination of projections, rotations, coordinate scaling and reflection. The eigenvector defined the direction of the rotation axis and the eigenvalues defined the coordinate scaling factor along this direction.
:::

## Decomposition

### Eigen-decomposition

Based on the theorem that any the symmetric matrix has eigenvectors, the representation of any quadratic form can be diagonalized under the basis consists by the eigenvectors. Given a symmetric matrix $A_{m\times m}$, the basis transform leads to 
$$
A=U\Lambda U^*,\quad \Lambda = \left(
\begin{array}{ccc}
\lambda_1 && \\
& \ddots &\\
&& \lambda_n\end{array}\right)
$$
where $U=(\boldsymbol{u}_1,\dots, \boldsymbol{u}_m)$ and $\boldsymbol{u_i}$ is a eigenvector of $A$ with a eigenvalue $\lambda_i$. In case $m>n=\text{rank}(A)$	, the matrix $U=(\boldsymbol{u}_1,\dots, \boldsymbol{u}_n, \boldsymbol{v}_{1}, \boldsymbol{v}_{m-n})$ where $\boldsymbol{v_i}$ span the compliment of the column space of $A$ and
$$
\Lambda = \left(
\begin{array}{cccc}
\lambda_1 &&& \\
& \ddots &&\\
&& \lambda_n&\\
&&& 0_{(m-n)\times (m-n)}\end{array}\right)
$$
where $0_{(m-n)\times (m-n)}$ is a $(m-n)\times (m-n)$ 0 matrix.

### Singularity Value Decomposition

Given a matrix $A_{m\times n}$ defined over a field $\mathbb{F}$, a singular value $\sigma\ge 0$ corresponds to a pair of unit vector $\boldsymbol{v}\in\mathbb{F}^n$ and $\boldsymbol{u}\in\mathbb{F}^m$ such that
$$
A\boldsymbol{v}= \sigma \boldsymbol{u},\quad A^*\boldsymbol{u}=\sigma\boldsymbol{v},
$$
where $\boldsymbol{v}$ and $\boldsymbol{u}$ is known as **right-singular** and **left-singular** vector for $A$. This decomposition is known as **singular value decomposition** (SVD).


Given a $m\times n$ matrix defined on a field $\mathcal{F}$, it can be decomposed into the form
$$
A=V\Sigma W^*,
$$
where the $V$ and $W$ are $m\times m$ and $n\times n$ unitary matrix consists by left and right singular vectors, respectively. The $\Sigma$ is a $m\times n$ matrix where the diagonal elements are singular values and the rests are 0.

The connection between the SVD and ED. Say the SVD of $A$ is $A=V\Sigma W^*$, then we have $AA^*=V\Sigma W^*W\Sigma V^*=V\Sigma\Sigma^*V^*$. On the other hand, Since $AA^*$ or $A^*A$ is symmetric, we have the ED $AA^*=V\Lambda_1 V^*$ which shows that $\Sigma\Sigma^*=\Lambda_1$. The same for $A^*A=W^*\Sigma\Sigma^* W$ such that $\Sigma^*\Sigma=\Lambda_2$. Without lose generality, we assume that $m\ge n$ then we have
$$
\Sigma=\left(
\begin{array}{c}
\sqrt{\Lambda_n} \\
0_{(m-n)\times n}
\end{array}
\right),
$$ 
where $\sqrt{\Lambda}_n$ is the $n\times n$ diagonal matrix with square root of the eigenvalues as diagonal elements.