# NumSpace
A journey through applied mathematics with a machine learning bias.

NumSpace is a work-in-progress repository where mathematical concepts — especially those related to linear algebra — are gradually implemented as they are studied and learned. The focus is on practical exploration and the applicability of these concepts in areas like Machine Learning and other computational applications.

⚠️ Note: This project is in an experimental stage and is not intended for production use. It is designed for learning and experimentation. Contributions, feedback, and ideas are welcome!

Let me know if you’d like further refinements!


# Intended module tree

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