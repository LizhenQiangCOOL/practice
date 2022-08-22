from collections import OrderedDict

# 利用orderedDict
class LRU(OrderedDict):
    """limit size, evicting the least recently looked-ip key when full
    python orderedDict 设计实际是 字典 + 双向链表
    """

    def __init__(self, maxsize=128, /, *args, **kwargs):
        self.maxsize = maxsize
        super().__init__(*args, **kwargs)

    def __getitem__(self, key):
        value = super().__getitem__(key)
        self.move_to_end(key)
        return value

    def __setitem__(self, key, value):
        if key in self:
            self.move_to_end(key)
        super().__setitem__(key, value)
        if len(self) > self.maxsize:
            oldest = next(iter(self))
            del self[oldest]


# python 完全手动实现: 字典+双向链表
class Node():

    def __init__(self, key: int, value: int):
        self.key = key
        self.value = value
        self.pre: Node = None
        self.next: Node = None


class DoublyLinkedList():
    """双向链表"""
    def __init__(self):
        # 使用空哨兵
        self.head: Node = Node(0,0)
        self.tail: Node = Node(0,0)

    def append_to_front(self, node: Node):
        node.next = self.head.next
        if node.next:
            node.next.pre = node
        else:
            # 链表插入前为空
            self.tail.next = node
        self.head.next = node
        node.pre = self.head


    def move_to_front(self, node: Node):
        # 如果node为最后一个节点
        if node.next is None:
            #只剩下一个元素
            if self.head.next == self.tail.next:
                return
            else:
                # 需要处理尾指针
                self.tail.next = node.pre
        else:
            # 处理后一个节点的前驱
            node.next.pre = node.pre
        # 处理前一个节点的后继
        node.pre.next = node.next
        # 移动到前面, 也要处理后继和前驱
        node.next = self.head.next
        self.head.next.pre = node

        self.head.next = node
        node.pre = self.head


    def remove_from_tail(self):
        if self.head.next == self.tail.next:
            self.head.next = None
            self.tail.next = None
        else:
            last_node = self.tail.next
            last_node.pre.next = None
            self.tail.next = self.tail.next.pre
            last_node.pre = None

class LRUCache():

    def __init__(self, max_size:int):
        self.max_size = max_size
        self.size = 0
        self.lookup = {}  #key:query, value:node
        self.linked_list = DoublyLinkedList()

    def get(self, key:int)->int:
        """从LRF缓存中查询, 会更新缓存顺序
        """
        node = self.lookup.get(key, None)
        if node is None:
            return -1
        self.linked_list.move_to_front(node)
        return node.value

    def set(self, key:int, value:int) -> None:
        """设置缓存，如果已存在则排到最前，
                    如果不存在则删除超过长度的node再插入新数据到最前
        """
        node = self.lookup.get(key, None)
        if node is not None:
            # key 存在更新值
            node.value = value
            self.linked_list.move_to_front(node)
        else:
            # key不存在
            if self.size == self.max_size:
                last_node = self.linked_list.tail.next
                self.lookup.pop(last_node.key)
                self.linked_list.remove_from_tail()
            else:
                self.size += 1
            new_node = Node(key, value)
            self.linked_list.append_to_front(new_node)
            self.lookup[key] = new_node


if __name__ == "__main__":
    l = LRUCache(2)
    l.set(1,1)
    l.set(2,2)
    l.linked_list.myprint()
    l.get(1)
    l.set(3,3)
    l.linked_list.myprint()
    l.get(2)
    l.set(4,4)
    l.linked_list.myprint()
    l.get(1)
    l.get(3)
    l.linked_list.myprint()
    l.get(4)
    l.linked_list.myprint()
    pass