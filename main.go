package main

import "fmt"

// Returns the index of the target if found, otherwise -1.
func BinarySearch(arr []int, val int) int {
    first, last := 0, len(arr)-1

    // Check for first and last position
    if len(arr) > 0 {
        if arr[first] == val {
            return first
        } else if arr[last] == val {
            return last
        }
    }

    for first <= last {
        mid := first + (last-first)/2

        if arr[mid] == val {
            return mid
        } else if arr[mid] < val {
            first = mid + 1
        } else {
            last = mid - 1
        }
    }

    return -1
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    val := 9

    result := BinarySearch(arr, val)
    if result != -1 {
        fmt.Printf("Found %d at index %d\n", val, result)
    } else {
        fmt.Println("Target not found")
    }
}
