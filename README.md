# Algorithms Project Documentation

## Index
- [Overview](#overview)
- [Sliding Window Algorithm](#sliding-window-algorithm)
- [Other Algorithms](#other-algorithms)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Python Cheat Sheet](#python-cheat-sheet)

## Overview
This project contains implementations of various algorithms, focusing on the sliding window technique and other algorithmic strategies. The goal is to provide clear and efficient solutions to common problems encountered in programming and computer science.

## Sliding Window Algorithm
The sliding window algorithm is a powerful technique used to solve problems involving arrays or strings. It is particularly useful for problems that require finding subarrays or substrings that satisfy certain conditions. 

For detailed implementation, refer to the [sliding_window.go](src/sliding_window.py) file.

## Other Algorithms
This section includes various algorithmic techniques and their respective functions. For more information, check the [other_algorithms.go](src/other_algorithms.go) file.

## Usage
To use the algorithms in this project, clone the repository and navigate to the `src` directory. You can run the Go files directly or import them into your own Go projects.

## Contributing
Contributions are not welcome! If you have suggestions for improvements or additional algorithms, please feel free to submit a pull request, but I won't approve it!

## License
This project is licensed under the MIT License.

### Index of Python Cheat Sheet
- [List Operations in Python](#list-operations-in-python)
    - [`pop()`](#pop)
    - [`remove()`](#remove)
    - [`del`](#del)
Here are some essential Python syntax and functions to remember:

### List Operations in Python

#### `pop()`
The `pop()` function removes the last element or the element at the specified index from a list.

- **Requires parameter?** Optional (index).
- **Returns value?** Yes, the value of the deleted element.
- **Error handling:** Raises `IndexError` if the specified index is not found.
- **Time complexity:** O(1) for the last element, O(n) for a specified index.

**Syntax:**
```python
list_name.pop()
list_name.pop(index)
```

#### `remove()`
The `remove()` function removes the first occurrence of the specified element from a list.

- **Requires parameter?** Yes, the value of the element.
- **Returns value?** No.
- **Error handling:** Raises `ValueError` if the specified value is not found.
- **Time complexity:** O(n).

**Syntax:**
```python
list_name.remove(value)
```

#### `del`
The `del` keyword removes an element at a specified index, slices a list, or deletes the entire list.

- **Requires parameter?** Yes, the index or slice.
- **Returns value?** No.
- **Error handling:** Raises `IndexError` if the specified index is not found.
- **Time complexity:** O(n).

**Syntax:**
```python
del list_name
del list_name[index]
del list_name[start:end]
```