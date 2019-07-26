import scala.math.pow

/* Leetcode Evaluation Runtime around 230ms */
object Solution {
  def reverse(x: Int): Int = {
    x >= 0 match {
      case true  => {
        val s: String = x.toString.reverse
        if (BigInt(s) >= (pow(2, 31)).toLong) 0 else BigInt(s).toInt
      }
      case false => {
        val j = 0 - BigInt(x)
        val s: String = j.toString.reverse
        val i = BigInt(s).unary_-
        if (i < (0 - pow(2, 31)).toLong) 0 else i.toInt
      }
    }
  }
}