package main

import "fmt"

func isPalindrome(n int) bool {
    originalN := n
    newN := 0

    for ; n > 0 ; {
        remainder := n % 10
        n /= 10

        newN = (newN * 10) + remainder
    }

    return newN == originalN
}

func LargestPalindromeProduct() {
    /**
     * A palindromic number reads the same both ways.
     * The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
     * Find the largest palindrome made from the product of two 3-digit numbers
     * answer: 906609
     */
     smallestN := 100
     largestN := 999

     result := 0;

     for i := smallestN; i <= largestN; i++ {
         for j := smallestN; j <= largestN; j++ {
             numToCheck := i * j
             if isPalindrome(numToCheck) && numToCheck > result {
                 result = numToCheck
             }
         }
     }


     fmt.Println(result)
}
