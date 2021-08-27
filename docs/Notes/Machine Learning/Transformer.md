# Transformer


## Architecture

The architecture of a transform is illustrated in the figure below:

![Transform Architecture](/img/docs/Image_arch_Transform.png)


The attention layer is designed to select and propagate the input features so that the selected features have longer lifetime which indicates the importance. This feature emphasis mechanism is called "Attention". Mathematically expression for the attention "A = A(x,y)" can be illustrated as

$$
A(x,y) = P(x,y)\cdot y,
$$
where $x$ is the the input control, $y$ is the input features, and the $P(x)$ is a projection operator determined by $x, y$ here to preserve part of the input features $y$ in the output.

### Scaled Dot-product Attention

The inputs of the scaled dot-product attention are queries, keys, and values, denoted as $\boldsymbol{Q}, \boldsymbol{K}, \boldsymbol{V}$ respectively. The dimension of the inputs denoted as $d_q\times d_m, d_k\times d_m$ and $d_k\times d_v$, where $d_m$ is the depth of each key, queries, and $d_v$ is the depth of return values, since the key and values are one-on-one mapping with each other, the number of value candidates is $d_k$ as well. The scaled attention takes the form that

$$
\text{Attention}(\boldsymbol{Q}, \boldsymbol{K}, \boldsymbol{V})=\boldsymbol{V}\cdot\text{Softmax}\left(\frac{\boldsymbol{Q}\cdot \boldsymbol{K}^T}{\sqrt{d_k}}\right),
$$

where the softmax function $\sigma(\boldsymbol{x})$ is a mapping $\sigma:\mathbb{R}^{d}\to[0,1]^d$, $k$ is dimension of the input vector. The explicity definition is

$$
\sigma(\boldsymbol{x})_i = \frac{e^{x_i}}{\sum_{k=1}^d e^{x_k}}.
$$

Since the $d_q$ is the number of queries input and $d_k$ is the number of candidate keys needs to be check. The product $\boldsymbol{Q}\cdot \boldsymbol{K}^T$ will return a tensor with shape $d_q\times d_k$ where the $(i,j)$-th element represents the matching between the $i$-th query with respect to $j$-th key. The last dot-production will project the corresponding values from $\boldsymbol{V}$ as return value for that query so that the output shape should be $d_q\times d_v$.  

### Multi-head Attention

It is generalization of the scaled dot-product attention to a joint form:

$$
\begin{aligned}
\text{MultiHead}(Q, K, V) &= W^OT(Q,K,V),\\
T_i(Q,K,V) &= \text{Attention}[(W^Q\cdot Q)_i,(W^K\cdot K)_i,(W^V\cdot V)_i],
\end{aligned}
$$
where the $W^Q, W^K$ and $W^V$ is a linear transform operator mapping $Q,K,V$ to have the same dimensionality form. $W^O$ is an extra linear transform for the outputs from the joint scaled dot-product attentions.

The schematic view of the multi-head attention:

![Multi-head attentions](/img/docs/Image_multi_head_attention.png)

### Position-wise Feed-foward Networks
Fully connected feed-foward networks $f(x)$ follow the attention layers:

$$
f(x) = \max(0, W\cdot x+b), 
$$

### Encoder

A encoder layer consists by two sublayer: one multi-head attention follows a simple, position-wise fully connected FFN. These two sub-layers are sequenced through a residual connection with normalization. A encoder in the transformer is a series of 6 encoder layers. 

### Decoder

A decoder layer consists by three sub-layers series connected: two multi-head attention layers and a position-wise fully connected FFN. The residual connection is also used to connecting the sub-layers. An other difference between the encoder and decoder is that the the second MHA layer of decoder can take extra input as token for MHA.

## Encoding

The inputs and outputs should be encoded with their position and time information. It is also needed to be embedded into a $d_{\text{model}}$ dimensional vector to consist with the model input shape. The time information is formulated by a look ahead mask. The positional information will be encoded into a positional sequence. These information will be as input for the model.

### Position Encoding

The inputs/outputs will be translated into a vector of tokens with dimensional $d$. 

Since the transformer is going to process the input sequence stride by stride, like a 1D-convolutional layer with kernel size $d$, where $d$ is the input size of the transformer. Although the relative or absolute position information may be crutial, it is lost while the sequence has been sliced to feed to the transformer. We can encode the position information $P(p,i)$, where $p=1,\dots,n$ is the position index and $i=1,\dots,d$ is the channel index, and iput them along with the sliced subsequence. One of the encoding method is:

$$
\begin{aligned}
P(p,2i)& = \sin(p/10000^{2i/d}),\\
P(p,2i+1)& = \cos(p/10000^{2i/d}).
\end{aligned}
$$ 

### Masking

The input tokens are pad into a certain size vector, a mask tensor is needed to indicate which entry is the padding (**Padding mask**). Another mask is needed to hide the information after a certain time when predicting the future output (**Look ahead mask**). These mask will make the input digits at where the masking location to negative infinity (smallest number in practice). 