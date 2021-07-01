# Set


A set $A$ is a collection of distinct elements. 

## Operations

Given two sets $A,B$, the basic operations on sets are
1. Union $A\cup B = \lbrace x\in A \text{or} x\in B\rbrace$,
2. Intersection $A\cap B = \lbrace x\in A \text{and} x\in B\rbrace$. 
3. Difference $A\\B = \lbrace x\in \text{and} x\notin B\rbrace$.
4. Symmetric difference $A\triangle B = \lbrace x \in A\cup B \text{and} x\notin A\cap B\rbrace $.

Relationship between the operations:
* $(A\cup B)\cap C=(A\cap C)\cup (B\cap C)$,
* $(A\cap B)\cup C=(A\cup C)\cap (B\cup C)$,
* $A\triangle B = (A\\B)\cup (B\\A)$
* The duality
$$
\begin{aligned}
S\\\bigcup_\alpha A_\alpha &= \bigcap_\alpha (S\\ A_\alpha),\\
S\\\bigcap_\alpha A_\alpha &= \bigcup_\alpha (S\\ A_\alpha).
\end{aligned}
$$

## Mapping of Sets
A mapping of set between sets $M,N$ is a function $f:M\to N$ that for any $x\in M$, there is a unique $f(x)\in N$, that is $f(M)\subset N$.
The inverse of the mapping $f^{-1}(y) = \lbrace x\in M | f(x) = y\rbrace $. Notice that $f^{-1}(y)$ may not be unique.


1. $f^{-1}(A\cup B) = f^{-1}(A)\cup f^{-1}(B)$,
2. $f^{-1}(A\cap B) = f^{-1}(A)\cap f^{-1}(B)$,
3. $f(A\cup B) = f(A)\cup f(B)$,
4. $f^{-1}(N\\ A) = f^{-1}(N)\\ f^{-1}(A)$,
5. $f^{-1}(A\triangle B) = f^{-1}(A)\triangle f^{-1}(B)$.