# Manifold 

## Topological Manifold

A $m$-dimensional topological manifold is a Hausdorff space $M$ that $\forall x\in M: \exists U_x$ which is a neighbourhood of $x$ homeomorphic to an open set in $\mathbb{R}^m$. Let $\varphi_U$ be the homeomorphism $\varphi_U:U\to \mathbb{R}^m$, the sets along with the homeomorphism, $(U,\varphi_U)$, is called the **coordinate chart**. The local coordinate of any point $x\in U$ is determined by te mapping $\varphi_U(x) = (u_1,\dots, u_m)$. Two coordinate charts $(U,\varphi_U)$ and $(V,\varphi_V)$ are $C^r$-compatible if the coordinate $x=\varphi_U(p)$ of any $p\in U\cap V$ satisfies $g(f(x))=x$, where $f,g$ are $C^r$ differentable and 
$$
f=\varphi_U\circ\varphi_V^{-1}, \quad g=\varphi_V\circ\varphi_U^{-1}.
$$

## Differentiable Structure

A $C^r$-differentiable structure on a manifold $M$ is a set of coordinate charts 
$$
\mathcal{A}=\lbrace (U,\varphi_U), (V,\varphi_V),\dots\rbrace,
$$
satisfying the following conditions
1. $\lbrace U,V,\dots, \rbrace$ forms an open covering of $M$.
2. Any two coordinate charts are $C^r$-compatible to each other.
3. $\mathcal{A}$ is the **maximal**, that is, if a $(\alpha, \varphi_\alpha)$ is a coordinate chart that $C^r$-compatible with all the charts in $\mathcal{A}$, this chart is also included in $\mathcal{A}$. 


## Differentiable Manifold

A manifold $M$ is called a $C^r$-**differentiable manifold** if there's a $C^r$-differentiable structure defined on $M$, denoted as $(\varphi, M)$ (where $\varphi(p)=\varphi_U(p)$ for $x \in U$). A **smooth manifold** is a $C^\infty$-differentiable manifold. 

* A differentiable manifold is **locally compact** that there is an compact set $V$ for any point $x$ with an open neighbour $U_x$ such that $x\in U_x\subset V$.
* For any $V\ne \varnothing$ and $(U,\varphi_U)$ such that $\overline{V}\subset U$, there is a smooth function $\mathcal{I}_{V\subset U}:M\to\mathbb{R}$ known as **indicator** such that $0\le \mathcal{I}_{V\subset U}\le 1$ and 
$$
\mathcal{I}_{V\subset U}(p) = \left\lbrace
\begin{aligned}
1 & p\in V,\\
0 & p\notin V.
\end{aligned}\right.
$$