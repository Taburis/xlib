# Generalized Inverse

## G-Inverse

Suppose $A$ is a $m\times n$ matrix, the generalized inverse of $A$ is a $n\times m$ matrix, denoted as $A^-$, such that
$$
AA^-A=A,
$$
which is usually called G-inverse of $A$. 

* The G-inverse $A^-$ is **always exists** for any $A$.
* **Idenpotent**: $AA^-AA^-=AA^-$, which means that $AA^-$ is a projection into $\mathcal{C}(A)$ (column space of $A$).
* Given linear equations $\boldsymbol{y} = A\boldsymbol{x}$, $A^-\boldsymbol{y} = \boldsymbol{x}$ if $\boldsymbol{y} \in\mathcal{C}(A)$ ($\boldsymbol{y}$ in the column space of $A$). 
* $A^-$ may not be unique, and $A^-=A^{-1}$ if A is a full rank matrix.


## Moore-Penrose Inverse

A G-inverse $A^-$ is a **Moore-Penrose inverse** for $A$ if
1. Hermitian: $(A^+A)^*=A^+A$ and $(AA^+)^*=AA^+$
2. Mutually G-inverse: $A^+AA^+=A$ and $AA^+A=A$.

Here $A^*$ stands for the Hermitian transpose.

It is easy to get the properties that:
* For any matrix $A$, its Moore-Penrose inverse $A^+$ exists and is unique.
* $(A^+)^+=A$,
* $(A^*)^+=(A^+)^*$,
* $(AA^*)^+ = (A^*)^+A^+$.
Because of the idenpotent and symmetry of $A^+$, the operators $AA^+$ and $A^+A$ are orthogonal projection into $\mathcal{C}(A)$.

The Moore-Penrose inverse of $A$ can be expressed as
* $A^+=(A^*A)^{-1}A^*$ and $A^+A=I$ if $A$ has linearly independent columns. 
* $A^+=A^*(AA^*)^{-1}$ and $AA^+=I$ if $A$ has linearly independent rows.
These conclusion comes from the fact that $A^*A$ is full rank matrix if $A$ has linearly independent columns, similar argument holds for the latter case.
