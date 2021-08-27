# Tree

## Binary Tree

Binary tree is a tree where have at most two branches (left or right) at each node. The node of a binary tree is usually denoted as $(S, L, R)$ where $S$ is the node itself and $L,R$ is the left or right side descendants. The level (or depth) $n$ of each node is $n$-th descendant of the root node. 

### Traversal
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