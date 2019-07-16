/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */
import scala.collection.mutable
import scala.util.control.Breaks._
/* Leetcode Evaluation Runtime around 670ms */
object Solution {
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    val numzip = nums.zipWithIndex
    val c = for {
      (a, a_index) <- numzip
      b_index <- (a_index + 1) to (nums.size - 1)
      if ((a + nums(b_index)) == target)
    } yield (a_index, b_index)
    Array(c(0)._1, c(0)._2)
  }
}

/* Leetcode Evaluation Runtime around 450ms */
object Solution1 {
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    val numzip = nums.zipWithIndex
    var result = mutable.ArrayBuffer[Int]()
    breakable {
      for ( (a, a_index) <- numzip) {
        for (b_index <- (a_index + 1) to (nums.size - 1) ) {
          if ( a + nums(b_index) == target)
          {   result += a_index
            result += b_index
            break   }
        }
      }
    }
    result.toArray
  }
}