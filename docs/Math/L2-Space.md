# L2 Space

Given a measurable space $(R, \mathfrak{S},\mu)$, a function $f$ defined on this space is square integrable if 
$$
\int_R f^2(x)d\mu,
$$
exists. The collection of all the square integrable functions is denoted as $L_2$.

Suppose $f,g$ are functions in $L_2$ space, then
1. $f*g$ is integrable
2. $f+g\in L_2$ and $\alpha f\in L_2$ where $\alpha$ is a number.

These two properties follows immdiately from Lebesgue's dominated convergence theorem by noticing 
$$
|f(x)g(x)|\le [f^2(x)+g^2(x)]/2,\quad [f(x)+g(x)]\le f^2(x)+2|f(x)g(x)|+g^2(x).
$$

$L_2$ forms a Euclidean space with the inner production defined as
$$
(f,g)=\int_Rf(x)g(x)d\mu.
$$
And the properties above insure that
1. $(f,g)=(g,f)$,
2. $(\alpha\cdot f, g)=\alpha(f,g)$,
3. $(f+g, h)=(f,h)+(g,h)$,
4. $(f,f)\ge 0$, and will equal zero only if $f=0$ a.e.
5. **Schwarz inequality**: $(f,g)^2\le (f,f)(g,g)$