# Tree Traversal.

## Binary Tree Traversal

Binary tree is a tree where have at most two branches (left or right) at each node. The node of a binary tree is usually denoted as $(S, L, R)$ where $S$ is the node itself and $L,R$ is the left or right side descendants. The level (or depth) $n$ of each node is $n$-th descendant of the root node. 

There are generally 3 types of order of traversal through a binary tree:

1. **Pre-order** visit the nodes in a order that `node-> left-> right`
2. **In-order**: `left->node->right`
3. **Post-order**: `left->right->node`

These three types of order of traversal can be implemented iterative in a very similar way. Suppose `action` is the function we'd like to impose on each node, then the traversal pesudo-code is:
```py
def iter(node, action):
	if not node: return
	action(node) # pre-order 
	iter(node.left)
	action(node) # in-order 
	iter(node.right)
	action(node) # post-order 
```

The pre-order is a special case of DFS. A BFS is a level-order traversal since it will visit all nodes in the same level before it move to the next level.

## Breadth-first Search (BFS)

This algorithm is for traversing or searching tree or graph structure, denoted as $G=(V,E)$ where $V,E$ are the vertices and edges of the graph $G$. The following pesudo-code illustrated a non-iterative BFS implementation for finding the `target` in a graph or tree starting at node `root`:

```py
def BSF(root, target):
	Q = queue()
	Q.enqueue(root)
	root.discovered = True
	while not Q.empty():
		node = Q.dequeue()
		if node == v: 
			return node
		# label the node is discovered
		node.discovered = True
		for nd in node.neighbour():
			if nd.discovered : continue
			Q.enqueue(nd)
```

This BSF can also be modified to be a non-iterative depth-first search (DFS) implementation by:

1. Using a stack instead of a queue
2. Check  the discovering while getting the element from the stack.

The properties of the BFS is that

* The time complexity is $\mathcal{O}(|V|+|E|)$ where $|V|, |E|$ is the number of vertices and edges, respectively.
* The memory of the BFS can be huge depending on the size of the graph. Typically it is $\mathcal{O}(b^{|V|})$ where $b, |V|$ is the number of the averaging branching for each vertex and the $|V|$ is the average number of vertices in the graph.
* The path it reaching to the target nodes is the one with smallest number of nodes. 


## Depth-first Search (DFS)

A similar algorithm like BFS, with the searching priority focused on the depth of the tree/graph $G=(V,E)$. A iterative pesudo-code of DFS is:

```py
def DFS(root, target):
	if root == target: return root
	root.discovered = True
	for node in root.neighbours:
		if node.discovered: continue
		DFS(node, target)
```

Also a non-iterative implementation is similar to the BFS code:
```py
def DSF(root,target):
	Q = stack()
	Q.push(root)
	root.discovered = True
	while not Q.empty():
		node = Q.pop()
		if node.discovered : return
		if node == target: 
			return node
		# label the node is discovered
		for nd in node.neighbour():
			nd.discovered = True
			Q.enqueue(nd)
	return None
```

* The **time complexity** of the DFS $\mathcal{O}(|E|+|V|)$ in the worest case, the same as the BFS.
* **memory space** of DFS is $\mathcal{O}(b)$, where $b$ is the average branching of the graph and $b=|E|$ in the worest case.