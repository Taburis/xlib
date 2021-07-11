# Quantization

The idea of quantization for a given model $f(x|\boldsymbol{W})$ with weights $\boldsymbol{W}$ is to approximate the original model by weights with lower precision. Many digits of the precision is redundant in terms of the marginal gain from increasing the precision of weights. Such it is possible to reduce the precision significantly with a minor tradeoff on the prediction accuracy. The quantization sliming the weight precision is expected to optimize a model in the following aspects:
* **Smaller storage size**: Quantizing a `float32` model into a `float16` precision will reduce the size by up to 50%.
* Faster performance and less memory usage: Less memory usage comes from the smaller size of weights needs to be loaded. Lower precision results a faster calculation process.

## Affine Quantization

[Reference](https://arxiv.org/pdf/2004.09602.pdf)

A quantization is a mapping $q:[a,b]\to I_q$ where $q$ stands for the target precision represented by a $q$-bit variable and $I_q=\lbrace 0, 1,\dots, 2^{q-1}\rbrace$. A straightforward mapping comes from slicing $[a,b]$ into $2^q$ equal size subintervals $r_i=[is,(i+1)s]$ where $s=(b-a)/2^q$ and map any $x\in a+r_i$ to $i$. So the affine quantization $f_q(x)$ can be represented as
$$
f_q(x) = \text{clip}_{[a,b]}\left\lbrace\bigg\lfloor \frac{x-a}{s}\bigg\rfloor \right\rbrace, \quad s = \frac{b-a}{2^q}.
$$
where $x\in [x_{\text{min}},x_{\text{max}}]$. The clip function here is to strip the under/over-flow values outside $[a,b]$ to constraint the boundary of $x$ into a finite region ($\text{clip}_{[a,b]}(x)=x$ for $x\in [a,b]$ and vanishes outside the interval) 

A de-quantization is a mapping from $I_q$ back to the original range $[a,b]$. It is used to restore the weights quantized before to preserve the layer representation:
$$
d_q(x) = s\cdot x_q+x_0,
$$
where $x_q=f_q(x)$ stands for the quantized value of $x$. 

Noticed that if the targeting precision is a signed $q$-bit variable, then 1 bit will be the position of sign and will only have $2^q-1$ avaliable distinct values.

A **Scale Quantization** is a version of simplified affine quantiztion for the special case that $x$ was clipped into a symmetric interval $[-a,a]$ and the target precision is a signed $q$-bit variable. Then there's no need to align the intervals after scaled. The quantization is then simplified to
$$
\begin{aligned}
f_q(x) &= \text{clip}_{[-a,a]}\left\lbrace\bigg\lfloor \frac{x}{s}\bigg\rfloor \right\rbrace, \quad s = \frac{2a}{2^q-1},\\
d_q(x) &= s\cdot x_q,
\end{aligned}
$$
where only $2^q-1$ of distinct values avaliable for a signed $q$-bits. 

## Quantization Network

[Reference](https://arxiv.org/pdf/1911.09464.pdf)