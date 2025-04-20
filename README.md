# Algorithms Learning Notes

## Sliding Window Algorithm

### Overview
The sliding window algorithm is a powerful technique used to solve problems involving arrays or strings. It is particularly useful for problems that require finding subarrays or substrings that satisfy certain conditions.

### Key Concepts
1. **Window Representation**:
   - Use a data structure (e.g., `map`, `int`) to represent the current state of the window.
   - Adjust the data structure based on the problem requirements (e.g., count of elements, sum of elements).

2. **Two Pointers**:
   - Use two pointers (`left` and `right`) to represent the boundaries of the window.
   - Move the `right` pointer to expand the window.
   - Move the `left` pointer to shrink the window when necessary.

3. **Window Updates**:
   - Update the window's data structure as elements are added or removed.
   - Perform any necessary calculations or checks within the window.

4. **Debugging**:
   - Use debug statements (e.g., `fmt.Println`) to print the state of the window during development.
   - Remove debug statements in the final implementation to avoid performance issues.

### Pseudocode Framework
```go
func slidingWindow(s string) {
    var window = ... // Initialize data structure
    left, right := 0, 0

    for right < len(s) {
        c := rune(s[right]) // Character to add
        window[c]++
        right++ // Expand window

        // Perform updates within the window
        ...

        // Shrink window if necessary
        for left < right && window needs shrink {
            d := rune(s[left]) // Character to remove
            window[d]--
            left++ // Shrink window
            // Perform updates after shrinking
            ...
        }
    }
}