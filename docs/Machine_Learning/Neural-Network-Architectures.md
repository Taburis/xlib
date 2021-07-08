#  Neural Network Architectures

## Convolution
Refering to this [introduction](https://cs231n.github.io/convolutional-networks/).

A typical mathematical definition of convolution $A\circ B$ for two matrix $A, B$ is
$$
(A\circ B)_{j,k}= \sum_{a,b}A_{a,b} B_{i-a,j-b}.
$$
This operation equivalent to scan the input tensor with a equal step by a filter, which is exactly the purpose of the convolution layer.

Given an input tensor $x$ with the shape $W_0\times H_0\times D_0$, a 2D convolutional layer is determined by
* Kernal size: $W\times H$ in 2D and $F$ filters will have $W\times H\times D$ kernel $K$ has $(W\times H\times D_0+1)\times D$ independent variables (biases included) based on the sharing parameter strategy. (They are $F$ filters and each of them has $D_0$ layers ($L^j_i$, the i-th layer of the j-th filter) with shape $W\times H$. One bias is assigned to each filter.)
* Then the j-th depth convolution layer output is determined by:
$$
O^j=b_j+\sum_{i=1}^{D_0}L^j_i\circ x_i,
$$
where a convolution operation is used here.
* Suppose the stride length is $S\le \min(W,H)$ and $P$ zero paddings for each side. The output tensor shape $W_1\times H_1\times D_1$ is determined
$$
\begin{aligned}
W_1 &= \lfloor(W_0-W+2P)/S\rfloor+1,\\
H_1 &= \lfloor(H_0-H+2P)/S\rfloor+1,
\end{aligned}
$$
where $\lfloor A\rfloor$ means the largest lower bound of integer approximation. 

## Max pool layer

The function of the max pool layer is boosting the traing speed and preventing the over-fitting by reducing the outputs from selecting the maximum value among the inputs within a loca region as the outputs. For instance, for a size $W\times H$ max pooling layer, it will slice the input into a many small region with shape $W\times H$, (padding or stride may apply). And only the maximum value of each slices will be the output from that region.

## Batch Normalization
The motivation of the batch normalization is based on the imperical conclusion that the NN training can be boosted and stablized if the activations are normalized to have a standard Gaussian distribution over the training samples. The batch normalization (BN) is targetting to practically implement this idea and preserve the representation of each layer. Since we know that normalized activation $X\sim N(0,1)$ which means it will shift the mean value of the original output $x$. This can potentially contradict to what this layer is suppose to represent. So an extra linear transform is applied to $X$. 

Suppose the mini-batch of samples leads to the inputs $x=\lbrace x_1,\dots, x_n\rbrace$ to BN, then the BN algorithm $y = BN(x|\gamma, \beta)$, where $\gamma,\beta$ are trainable parameters, is:
1. Calculate the mean and variance:
$$
\begin{aligned}
\mu &= \frac{1}{n}\sum_i x_i, \\
\sigma^2 &= \frac{1}{n}\sum_i(x_i-\mu)^2,
\end{aligned}
$$
2. Standardized the $x\to X$ by $X=(x-\mu)/(\sigma+\epsilon)$ where $\epsilon>0$ is a small number added for stability purpose. 
3. Apply the linear transformation to get the final ouput $y=\gamma X+\beta$.

Noticed that the transform will leads to original $x$ if $\gamma,\beta$ was trained to be $\gamma=\sigma$ and $\beta=\mu$. 

## Residual Connection (ResNet)

[Reference](https://arxiv.org/pdf/1512.03385.pdf)

Learning the ground truth $T(x)$ for input $x$ is straightfoward. However, the disadvantages come from the fact that the gradient is getting weaker when backward propagatig longer layers. This weakening makes the features extracted from the first few layers are hard to be trained properly which caused the input features are smeared. To solve this problem, instead of fitting the ground truth $T(x)$, we can try to fit the resdiual of the ground truth, $R(x)=T(x)-x$, respected to the input $x$. For doing this, a residual connection, known as shortcut connection, is introduced:

![ResNet Architecture](/img/docs/Image_arch_ResNet.png)

The output from layer $F(x|\theta_i)$ will needs to combined the shortcut to form the input for the next layer 
$$
x_i = F(x_{i-1}|\theta_i)+P\cdot x_{i-1},
$$
where $x_i$ is the output from the $i$-th layers and the $x_0$ stands for the raw input. The linear operator $P$ is used for projection purpose in the case that the channel number is different between the $x_i$ and $x_{i-1}$ ($P=1$ if the channel numbers are the same). For practice purpose, the projection can be done by a $1\times 1$ convolution. Furthermore, the normalization is included in $F(x_{i-1}|\theta_i)$. 

The **Advantage** of this structure is that the input features $x$ are allowed to be propagated even if the residual features $R_i(x)$ hadn't effectively extracted in the first few layers (due to the weakening of the gradients). This can benefits in large model training and reduce the error mostly raised by the non-proper features extracted from inputs.