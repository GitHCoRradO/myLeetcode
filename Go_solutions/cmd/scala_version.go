package main
/*
object Solution {
    def copyRandomList(head: node): node = {
    val visited = scala.collection.mutable.HashMap[node, node]()

    def getClonedNode(node: node): node = {
      if (node != null) {
        if (visited.contains(node)) {
          return visited(node)
        } else {
          visited.put(node, new node(node.value))
          return visited(node)
        }
      }
      null
    }

    if (head == null) return null
    var oldNode = head
    var newNode = new node(oldNode.value)
    visited.put(oldNode, newNode)
    while (oldNode != null) {
      newNode.random = getClonedNode(oldNode.random)
      newNode.next = getClonedNode(oldNode.next)
      oldNode = oldNode.next
      newNode = newNode.next
    }
    visited(head)
  }
}

 */
