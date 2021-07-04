# Set


A set $A$ is a collection of distinct elements. 

## Operations

Given two sets $A,B$, the basic operations on sets are
1. Union $A\cup B = \lbrace x\in A \text{ or } x\in B\rbrace$,
2. Intersection $A\cap B = \lbrace x\in A \text{ and } x\in B\rbrace$. 
3. Difference $A\setminus B = \lbrace x\in A\text{ and } x\notin B\rbrace$.
4. Symmetric difference $A\triangle B = \lbrace x \in A\cup B \text{ and } x\notin A\cap B\rbrace $.

Relationship between the operations:
* $(A\cup B)\cap C=(A\cap C)\cup (B\cap C)$,
* $(A\cap B)\cup C=(A\cup C)\cap (B\cup C)$,
* $A\triangle B = (A\setminus B)\cup (B\setminus A)$
* The duality
$$
\begin{aligned}
S\setminus\bigcup_\alpha A_\alpha &= \bigcap_\alpha (S\setminus A_\alpha),\\
S\setminus\bigcap_\alpha A_\alpha &= \bigcup_\alpha (S\setminus A_\alpha).
\end{aligned}
$$

## Mapping of Sets
A mapping of set between sets $M,N$ is a function $f:M\to N$ that for any $x\in M$, there is a unique $f(x)\in N$, that is $f(M)\subset N$.
The inverse of the mapping $f^{-1}(y) = \lbrace x\in M \mid f(x) = y\rbrace$. Notice that $f^{-1}(y)$ may not be unique.


1. $f^{-1}(A\cup B) = f^{-1}(A)\cup f^{-1}(B)$,
2. $f^{-1}(A\cap B) = f^{-1}(A)\cap f^{-1}(B)$,
3. $f^{-1}(A\triangle B) = f^{-1}(A)\triangle f^{-1}(B)$.
4. Suppose $A\subset M$, $f^{-1}(M\setminus A) = f^{-1}(M)\setminus f^{-1}(A)$, 
5. $f(A\cup B) = f(A)\cup f(B)$.
6. $f(A\cap B) \subset f(A)\cap f(B)$. 

Basically the 1 to 4 equations implies that inverse mapping is permutable with all the set operations. But the intersection is not permute with mapping due to the fact that two distinct elements can have the same image.

**Proof**
1. Suppose $x\in f^{-1}(A\cup B)$, then $f(x)\in A\cup B$ which means that $x$ comes from either $f^{-1}(A)$ or $f^{-1}(B)$ which is the right side of the equation.
2. Suppose $f(x)\in A\cap B$, then $x$ needs to be in $f^{-1}(A)$ and $f^{-1}(B)$, but also $x\in f^{-1}(A\cap B)$ by definition. 
3. Noticing $A\triangle B = (A\setminus B)\cup (B\setminus A)$, this conclude the proof by using the previous two equations.
4. Suppose $f(x) \in M\setminus A$, then $x$ should be in $f^{-1}(M)$ but not in $f^{-1}(A)$ which is $f^{-1}(M)\setminus f^{-1}(A)$.  
5. Suppose $x \in A\cup B$,  then $f(x)$ has to be in either in $f(A)$ or $f(B)$ which is $f(A)\cup f(B)$.
6. This relation is obvious, here I show that inverse subset is not hold. In fact, it can be an point $y\in A\setminus B$ and $x\in B\setminus A$ such that $f(x)=f(y)$. In this case, $x\notin A\cap B$ but $f(x) \in f(A)\cap f(B)$. 

## Collection of Sets
Given a set $X$, a set of the subsets of $X$ is called a collection. A collection can also be a ring or semi-ring if certain conditions were satisfied.

### Semi-Ring of Sets

A collection of set $\mathcal{S}$ is called a semi-ring if it satisfies the conditions:
1. $\varnothing \in \mathcal{S}$,
2. If $A,B\in\mathcal{S}$, then $A\cap B\in \mathcal{S}$. 
3. Finite parition: If $A,A_1\in \mathcal{S}$ and $A_1\subset A$, then there are subsets $A_2, \dots, A_n$ in $\mathcal{S}$ such that
$$
A=\bigcup_{k=1}^nA_k,\quad A_i\cap A_j =\varnothing, \forall i\ne j.
$$

**Lemma** of semi-rings:
1. Suppose $A_1,\dots, A_n$ in $\mathcal{S}$ are mutually orthogonal sets, and $A\in \mathcal{S}$ that $A_i\subset A$. Then there are $\lbrace A_{n+1}, \dots, A_m\rbrace$ forms a finite partition for $A$:
$$
A=\bigcup_{k=1}^m A_k.\quad A_i\cap A_j =\varnothing \text{ if } i\ne j.
$$
2. Given $A_1,\dots A_h$ in a semi-ring $\mathcal{S}$, then there are mutually orthogonal $\lbrace B_1,\dots, B_m\rbrace$ ($B_i\cap B_j =\varnothing$) in $\mathcal{S}$ such that
$$
A_k = \bigcup_{i\in s_k} B_i,
$$
where $s_k$ is a collection of index sequence. 


**Proof**
1. To prove this lemma by induction, the case $n=1$ is valid by the definition and we assume it is valid that
$$
A= \bigcup_{i=1}^n A_i\cup \bigcup_{k=1}^m B_k,
$$
and denote $B_{i1}=A_{n+1}\cap B_i$, by definition, we have $B_ik$ so that $B_i=\cup_k^{m_i}B_ik$. Then we have the equation
$$
A= \bigcup_{i=1}^{n+1} A_i\cup \bigcup_{i=1}^n\bigcup_{k=2}^{m_i} B_{ik}.
$$
2. This holds for $h=1$ case by definition. Now assume that it is hold upto $h$ and we have the collection $\lbrace  B_1, \dots,  B_s\rbrace$ satisfies the condition 1. Suppose $B_{i1} = A_{h+1}\cap B_i$, we have a finite parition $B_{i1}\cup\dots\cup B_{in} = B_i$ where $\lbrace B_{ik}\rbrace$ in $\mathcal{S}$ by definition. And the collection $\lbrace B_{ik}\rbrace $ forms the new sequence. They are mutually orthogonal since $\lbrace B_i\rbrace $ is mutually orthogonal and conclude the proof.
2. This lemma hold for $n=1$ by definition. By induction, we assume that it was true for $n$, then we have pairwise disjoint sets $\lbrace B_k\rbrace, k=1,\dots,t$ satisfying the lemma conidtion. Now let $B_{s1}=A_{n+1}\cap B_{s}$, then we have $B_{s}=\cup_{k=1}^{m_s}B_{sk}$, and $A_{n+1}$ can be expressed as
$$
A_{n+1}=\bigcup_{i=1}^nB_{i1}\cup\bigcup_{j=1}^l C_j,
$$  
by the first lemma. Since any $A_i,i\le n$ can be expressed as union of some of $B_j$, so they can be expressed as the union of some of $B_{ik}$ as well. Furthermore, $\lbrace B_{ik}\rbrace$ are pairwise disjoint due to the fact that $\lbrace B_j\rbrace$ are pairwise disjoint. So is $\lbrace C_j\rbrace$. then this lemma is hold for $n+1$ case which leads to the conclusion.

### Ring of Sets

Given a collection of sets $\mathcal{F}$, it is called a **ring** if for any two elements $A,B\in\mathcal{F}$:
1. $A\triangle B\in \mathcal{F}$,
2. $A\cap B \in mathcal{F}$. 

These two conditions implies that $A\cup B\in \mathcal{F}$ and $A\setminus B\in \mathcal{F}$ since $A\cup B = (A\triangle B)\triangle (A\cap B)$ and $A\setminus B=A\triangle(A\cap B)$. Basically, a ring is closed for the basic set operations. And also $\varnothing \in \mathcal{F}$ since $A\triangle A =\varnothing$. A ring of sets can have usual **algebraic sense** if we take $A\triangle B$ as addition and $A\cap B$ as multiplication. In this definition, the empty set $\varnothing$ and the union of all sets $E$ equivalents to 
zero and unit. Furthermore, the elements will satisfy the condition
$$
A+A=0,\quad A^2=A.
$$
A ring satifies the condition above is called a **Boolean ring**. It implies that every Boolean can be realized as a ring of sets with the addition and multiplication above.


Theorems about rings:
1. The intersection of arbitrary number of rings is a ring.
2. Suppose $\mathcal{S}$ is a collection of subsets of $X$, then there is a unique ring $\mathcal{F}(\mathcal{S})$ that containing all the elemnets of $\mathcal{S}$ and is contained in every ring that containing $\mathcal{S}$. Furthermore, suppose $\mathcal{B}$ is a collection of sets that admiting a finite partion for any sets $A\in\mathcal{B}$:
$$
A = \bigcup_{k=1}^nA_k,\quad A_k\in \mathcal{S}.
$$ 
Then it is the ring $\mathcal{F}(\mathcal{S})$.

**Poorf**
1. This is a direct conclusion from the definition of the ring.
2. Based on the first lemma, the intersections of the elements from all the rings containing $\mathcal{S}$ forms the ring $\mathcal{F}(\mathcal{S})$. Now we approve the collection $\mathcal{B}$ is a ring. Suppose $A, B\in\mathcal{B}$, by the assumption we have
$$
A=\bigcup_{k=1}^mA_i,\quad B=\bigcup_{k=1}^nB_k.
$$
With the notation $C_{ij}=A_i\cap B_j$, it is obviously that $C_{ij}\in\mathcal{S}$ and by the lemma 1. of the semi-ring, we have
$$
A_i=\cup_{j=1}^mC_{ij}\bigcup \cup_{k=1}^{h_i}D_{ik},\quad B_j=\cup_{i=1}^mC_{ij}\bigcup \cup_{k=1}^{l_j}E_{kj},
$$
where $D_{ik},E_{kj}\in\mathcal{S}$. This decomposition implies
$$
A\cap B = \bigcup_{i,j} C_{ij},\quad A\triangle B = \cup_{i,k}D_{ik}\bigcup\cup_{j,k}E_{jk}.
$$
Therefore $A\cap B$ and $A\triangle B$ are in the collection $\mathcal{B}$ which means it is a ring. And it is obviously the minimal ring contained the $\mathcal{S}$ so that $\mathcal{B}=\mathcal{F}(\mathcal{S})$. 

Usually we call the $\mathcal{F}(\mathcal{S})$ as **the ring generated by** $\mathcal{S}$. 

### Sigma-Ring and Borel Algebra

* A ring $\mathfrak{R}$ of sets is called $\sigma$-ring if
$$
A=\bigcup_i A_i,\quad A_i\in \mathfrak{R},
$$
implies $A\in\mathcal{R}$ for countable $A_i$.
* A ring $\mathfrak{R}$ of sets is called $\delta$-ring if
$$
A=\bigcap_i A_i,\quad A_i\in \mathfrak{R},
$$
implies $A\in\mathcal{R}$ for countable $A_i$.

If a $\sigma$-ring contained a unit, then it is called a $\sigma$-algebra. So is the definition of $\delta$-algebra. A $\sigma$-algebra is also a $\delta$-algebra, and vice versa since 
$$
\bigcup_iA_i=E\setminus\bigcap_i(E\setminus A_i),\quad \bigcap_iA_i=E\setminus \bigcup_i(E\setminus A_i).
$$
So they are both called as a **Borel algebra (or B-algebra)**. If $\mathfrak{S}$ is a nonempty collection of sets, there is a Borel algebral $\mathfrak{B}(\mathfrak{S})$ containing $\mathfrak{S}$ and contained in every B-algebra containing $\mathfrak{S}$. This $\mathfrak{B}(\mathfrak{S})$ is called the **minimal B-algebra over** $\mathfrak{S}$ or the **Borel closure** of $\mathfrak{S}$. 

Mapping on the Borel algebra $f(\mathfrak{B})$ denotes the image of all the sets $A\in\mathfrak{B}$. The mapping leads to the following properties:
1. If $\mathfrak{R}$ is a ring, then $f^{-1}(\mathfrak{R})$ is a ring.
2. If $\mathfrak{R}$ is an algebra, $f^{-1}(\mathfrak{R})$ is an algebra.
3. If $\mathfrak{R}$ is a B-algebra, $f^{-1}(\mathfrak{R})$ is a B-algebra.
4. $\mathfrak{R}(f^{-1}(\mathfrak{M}) = f^{-1}(\mathfrak{R}(\mathfrak{M}))$ ( inverse mapping is permutable with the ring generation operation).
5. $\mathfrak{B}(f^{-1}(\mathfrak{M}) = f^{-1}(\mathfrak{B}(\mathfrak{M}))$ ( inverse mapping is permutable with the B-algebra generation operation).
