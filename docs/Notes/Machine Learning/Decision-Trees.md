# Decision Trees

The decision tree algorithms are supervised learning for solving regression and classification problems. The problems can be formulated as follows:
Given datasets containing labels $\lbrace y_n\rbrace$ and features $\lbrace x_n\rbrace$ paired as $(x_i,y_i)$, we need to construct a function or machine $y=f(x)$ to predict the label $y$ based on the input features $x$. The label could be discrete (classification) or continuous (regression) case.

## Decision Tree Algorithm
---

The Classification and Regression Tree (CART) is a term referring general decision tree algorithms. In CART algorithms, the machine $f(x)$ is implemented as a tree. Mostly a binary tree is used for simplicity (ideally, any tree can be split into a binary tree). To build a CART, a metric to measure the performance of the model is needed. This metric is a functional evaluated on the samples using the $i$-th feature. With it, an information gain at a split point on the feature can be defined accordingly. Furthermore, one can search the split point to maximize the information gain, and build a child node if necessary. 

Given $F$ features $\lbrace X_F\rbrace$, the feature boundary or decision boundary $B$ is the direct product of $F$ intervals 
$$
B=[a_1,b_1]\otimes\dots \otimes[a_F,b_F].
$$
Split the $i$-th feature from the boundary $B$ at $p$ will lead to two sub boundaries:
$$
\begin{aligned}
B_L&=[a_1,b_1]\otimes\dots \otimes[a_{i},p]\otimes\cdots\otimes[a_F,b_F]\\
B_R&=[a_1,b_1]\otimes\dots \otimes[p, b_i]\otimes\cdots\otimes[a_F,b_F].
\end{aligned}
$$

The idea of the decision tree learning is to partition the features into small boundaries iteratively until the partition is fine enough for prediction. The iterations built a CART can be illustrated as the following:
```python
# Initiate the feature boundary
boundary = feature_boundary()
root = Node(boundary)
node_list = list()
node_list.append(root)

while len(node_list) >0 :
	node = node_list.pop(0)
	# select the feature candidate for exploring, return the feature index i 
	i = feature_select(node, features, labels)
	# calculate the splitting point leads to a maximum information gain by using the i-th feature
	gain, split_point = split_search(node, features[index], labels)

	# skip if no split point exists: feature exhausted
	if not split_point : 
		#assign label
		node.label = get_label()
	# skip if the gain is too small
	if gain < gain_threshold: continue

	# keep the split point for the current node, split the node into left/right child nodes,
	# for the selection: feature[i] <= split_point -> node.left
	#                    feature[i] >  split_point -> node.right
	node.not_leaf = 1 # 
	node.split_point = split_point
	node.feature_index = i
	# Split the boundary for the i-th feature
	boundary_left = split_left(i, split_point)
	boundary_right = split_right(i, split_point)
	node.left = Node(boundary_left)
	node.right= Node(boundary_right)
	node_list.append(node.left)
	node_list.append(node.right)
``` 

### Classification Trees

Suppose the $F$ features $\lbrace X_F\rbrace$ assigned for each data labeled by $y$ where $C$ possible classes in total, several metrics can be used for classification problem. Given a feature boundary $B$, these metrics can be:

* **Entropy** is a metric used for discrete features. The entropy $H_i(S_{B})$ with respect to the $i$-th feature component, where $S_{B}$ is the samples falling in the feature boundary $B$, is defined as
$$
H(S_{B})=\sum_{k=1}^C\sum_{x\in B_i}-p_k(x)\ln p_k(x),
$$
where $p_k(x)$ is the probability to find a $k$-th class data satisfying $X_i=x$ in the sample $S_{B}$. In general, smaller entropy means less randomness and better information. $H_i(S_{B})=0$ means a perfect classification by using the $i$-th feature. The entropy will calculated for all features and the feature with the **smallest** entropy is picked for splitting. Given a split that $B\to B_L, B_R$, the **information gain** $G(B, B_L, B_R)$ is defined as
$$
G_H(B, B_{L}, B_{R}) = H(B)-\alpha\cdot H(B_L)-(1-\alpha)\cdot H(B_R),
$$
where $\alpha$ is the fraction of the data from $B$ falls into $B_L$. The split point can be determined from searching the maximum information gain from the possible split points along the $X_i$. 
* **Gini Impurity** $I_G(B)$ is defined as
$$
I_G(B) = \sum_{i=1}^C\sum_{x\in B_i}p_k(x)(1-p_k(x))=1-\sum_{k=1}^Cp^2_k(x),
$$ 
where $p_k(x)$ is the probability of finding a $k$-th class data with $X_i=x$ from the sample in $B$. A perfect classification, only one class in the boundary, leads to 0 impurity, so it is smaller the better. Suppose a split that divide the boundary into two subsets $B_L, B_R$, and let $\alpha$ and $1-\alpha$ stands for the fraction of original samples fall into the two subsets. A new impurity for this split is defined as
$$
I_G(B_L,B_R)= \alpha\cdot I_G(B_L)+(1-\alpha)\cdot I_G(B_R).
$$
The **impurity gain** $G(B_L,B_R)$ is defined as the difference:
$$
G_I(B,B_L,B_R)=I_G(B)-I_G(B_L,B_R).
$$
Again, the split point is searched through the $X_i$ to find the point maximum the gain $G_I$. 

:::tip Impurity vs Entropy
Practicing experience shows that entropy method is slightly better performed than the impurity. However, the impurity metric is faster than the entropy calculation due to the relative simple expression. 
:::

:::note Extend to Continuous features
To apply these metrics for continuous features, the feature ranges are needed to be quantized into subintervals, then the summation over the discrete feature values in the metrics above, the features are summed over the subintervals for the continuous features. The intervals are usually divided into two subintervals so that a binary search can be used to find the maximum split point.
:::

### Regression Trees

For regression problem, the input features are supposed to be continuous. For the regression tree, a linear approximation, mostly the constant approximation, is driven by using the data in the feature boundary $B$, so that a more fine partition will leads to a better approximation. Let the $\hat y_j(x)$ be the $j$-th of the output from the regression tree, then $\hat y_j(x)=C_{B_x}$ where $C_{B_x}$ is a constant and $B_x$ is the finest feature boundary contained the $x$. A residual error $R^2$ is often used to measure the performance:
$$
R^2(S_B) = \frac{1}{|S_B|}\sum_{j}\sum_{(x,y)\in S}(y_j-\hat y_j(x))^2,
$$
where $S_B$ is the sample sets contained in the feature boundary $B$ and the element is $S_B$ is the pair $(x_i,y_i)$, and $|S_B|$ is the number of elements in $S_B$. For a split $B\to (B_L,B_R)$, the variance reduction gain $G(S_B, S_{B_L}, S_{B_R})$ is defined as 
$$
G(S_B, S_{B_L}, S_{B_R}) = R^2(S_B)-\alpha\cdot R^2(S_{B_L})-(1-\alpha)\cdot R^2(S_{B_R}),
$$
where $\alpha = |S_{B_L}|/|S_B|$ is the fraction of the sample in $S_{B_L}$ relative to $S_B$.

A splitting threshold $G_0$ is set to prevent over splitting. Only the gain larger than the threshold, $G(S_B, S_{B_L},S_{B_R})\ge G_0$, will leads to a new split. 

## Gradient Boosting
---

