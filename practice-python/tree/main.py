class Node:
    def __init__(self, p:Node, left:Node,right:Node,data:int):
        self.p = p
        self.left = left
        self.right = right
        self.data = data

def PostorderTreeWalk(root:Node):
    stack = []
    result = []
    last_visit = None
    while root or stack:
        while root:
            stack.append(root)
            root = root.left
        node = stack[-1]
        if node.right is None or node.right == last_visit:
            stack.pop()
            result.append(node.value)
        else:
            node = node.right
    return result
