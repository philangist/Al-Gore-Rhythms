class Node(object):
     def __init__(self, value, left=None, right=None):
         self.value = value
         self.left = left
         self.right = right
     def __repr__(self):
         return u'Node(value={})'.format(self.value)


def dfs(node, value):
     if not node:
         return None
     print 'Visiting node %d...' % node.value
     if node.value == value:
         print 'Found value %d!' % value
         return node
     if not (node.left or node.right):
         return None
     left_branch = dfs(node.left, value)
     if left_branch:
         return left_branch
     return  dfs(node.right, value)

def dfs_iterative(node, value):
     stack = [node]
     while True:
         if not stack:
             return None
         current_node = stack.pop()
         print 'Visiting node %d...' % current_node.value
         if current_node.value == value:
             print 'Found value %d!' % value
             return current_node
         if not(current_node.left or current_node.right):
             continue
         if current_node.right:
             stack.append(current_node.right)
         if current_node.left:
             stack.append(current_node.left)

def bfs_iterative(node, value):
     queue = [node]
     while queue:
         current_node = queue[0]
         print 'Visiting node %d...' % current_node.value
         if current_node.value == value:
             print 'Found value %d!' % value
             return current_node
         if not(current_node.left or current_node.right):
             queue = queue[1:]
             continue
         if current_node.left:
             queue.append(current_node.left)
         if current_node.right:
             queue.append(current_node.right)
         queue = queue[1:]
