/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */
/* Leetcode Evaluation Runtime around 410ms */
import scala.collection.mutable.ArrayBuffer
object Solution {
  def addTwoNumbers(l1: ListNode, l2: ListNode): ListNode = {
    var l1node = l1
    var l2node = l2
    var sl1: String = l1node.x.toString
    var sl2: String = l2node.x.toString
    while (l1node.next != null) {
      sl1 += l1node.next.x.toString
      l1node = l1node.next
    }
    while (l2node.next != null) {
      sl2 += l2node.next.x.toString
      l2node = l2node.next
    }
    sl1 = sl1.reverse
    sl2 = sl2.reverse
    val bi1 = BigInt(sl1)
    val bi2 = BigInt(sl2)
    val bi3 = bi1 + bi2
    val resultStr: String = bi3.toString.reverse
    val arrayNodes = new ArrayBuffer[ListNode]
    for (i <- 0 until resultStr.length){
      val node: ListNode = new ListNode(resultStr(i).toInt - 48)
      arrayNodes += node
    }
    if (arrayNodes.length == 2) {
      arrayNodes(0).next = arrayNodes(1)
    } else if (arrayNodes.length > 2) {
      for (i <- 0 until arrayNodes.length - 2) {
        arrayNodes(i).next = arrayNodes(i + 1)
      }
      arrayNodes(arrayNodes.length - 2).next = arrayNodes(arrayNodes.length - 1)
    }
    arrayNodes(0)
  }
}