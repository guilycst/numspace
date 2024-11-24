# NumSpace

**A journey through applied mathematics with a machine learning bias.**

NumSpace is a **work-in-progress** repository aimed at implementing mathematical concepts — with a strong focus on **linear algebra** — as they are studied and explored. This project is designed to highlight the practical applicability of mathematical principles, especially in areas like **Machine Learning** and **computational applications**.

⚠️ **Note**:  
This repository is in an **experimental stage** and is **not intended for production use**. The primary goal is **learning and experimentation**. Contributions, feedback, and ideas are welcome to help shape its development!

---

## Intended Module Tree

Below is the planned directory and module structure for NumSpace:

```
numspace/
│
├── algebra/                 # Focused on general algebra and linear algebra
│   ├── vectors.go           # Operations on vectors (addition, dot product, norm)
│   ├── matrices.go          # Operations on matrices (multiplication, transpose, inverse)
│   ├── tensors.go           # Operations on tensors
│   └── eigen/               # Subpackage for eigenvalues and eigenvectors
│       ├── decomposition.go # LU decomposition, QR decomposition, etc.
│
├── calculus/                # Focused on numerical calculus
│   ├── differentiation.go   # Numerical differentiation
│   ├── integration.go       # Numerical integration
│
├── geometry/                # Geometric operations
│   ├── transformations.go   # Rotation, scaling, translation
│   ├── projections.go       # Projections onto planes
│
├── stats/                   # Statistics and probability
│   ├── distributions.go     # Probability distributions (normal, binomial, etc.)
│   ├── regression.go        # Linear regression, correlation analysis
│
└── utils/                   # Utility functions
    ├── errors.go            # Error handling
    ├── format.go            # Formatting and debugging tools
    ├── random.go            # Random number generation
```

---

### Goals and Features

- **Modular Design:**  
  The repository is structured into well-defined modules that align with key mathematical fields, ensuring clarity and ease of use.

- **Focus on Linear Algebra:**  
  Core implementations such as matrix operations, eigenvalue decomposition, and tensor manipulations are prioritized.

- **Applicability in Machine Learning:**  
  While not exclusively focused on ML, many implemented concepts directly support common ML applications.

---

### How to Contribute

If you’d like to contribute, here’s how you can help:
1. **Report Issues**: Found a bug or have a suggestion? Open an issue!
2. **Submit Pull Requests**: Contributions to improve existing code or add new features are always welcome.
3. **Share Feedback**: If you have ideas for modules or enhancements, let us know.

