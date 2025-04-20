def sliding_window(s: str) -> str:
    # Initialize data structure
    window = {}
    left, right = 0, 0
    max_length = 0
    start = 0

    while right < len(s):
        # Character to add
        c = s[right]
        window[c] = window.get(c, 0) + 1
        right += 1  # Expand window

        # Perform updates within the window
        while len(window) > 0 and window[c] > 1:
            # Character to remove
            d = s[left]
            window[d] -= 1
            if window[d] == 0:
                del window[d]
            left += 1  # Shrink window

        # Update max_length and start position
        if right - left > max_length:
            max_length = right - left
            start = left

    # Return the longest substring
    return s[start:start + max_length]


# Main function to test the sliding window algorithm
if __name__ == "__main__":
    test_string = "abcabcbb"
    result = sliding_window(test_string)
    print("Longest substring without repeating characters:", result)