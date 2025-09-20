# Contribution Guidelines

Thank you for considering contributing to the GoStaticServe project! This document provides guidelines and best practices for participating in this project, helping you get started quickly and effectively contribute to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Reporting Issues](#reporting-issues)
  - [Submitting Code](#submitting-code)
  - [Writing Documentation](#writing-documentation)
- [Development Environment Setup](#development-environment-setup)
- [Code Standards](#code-standards)
- [PR Submission Process](#pr-submission-process)
- [Testing](#testing)
- [Feedback](#feedback)

## Code of Conduct

We expect all contributors to follow basic code of conduct when participating in the project:

- Respect other contributors
- Accept constructive criticism
- Focus on the best interests of the project
- Maintain a friendly and inclusive attitude

## How to Contribute

### Reporting Issues

If you find a bug or have a suggestion for a new feature, please submit an issue on GitHub. When submitting an issue, please follow this format:

**Bug Report:**
- Title: Concisely describe the issue
- Environment: Operating system, Go version, browser (if applicable)
- Steps to reproduce: Detailed steps explaining how to reproduce the issue
- Expected behavior: Describe what you expected to happen
- Actual behavior: Describe what actually happened
- Additional information: Such as error logs, screenshots, etc.

**Feature Request:**
- Title: Concisely describe the feature
- Detailed description: Explain the purpose and implementation ideas of the feature
- Priority: Indicate the importance of the feature

### Submitting Code

We welcome you to submit code to fix bugs or add new features. Before submitting code, please ensure:

1. Your code complies with the project's [code standards](#code-standards)
2. Your code has passed all tests
3. Your commit messages are clear and concise

### Writing Documentation

Good documentation is crucial for the project. You can:

- Improve existing documentation
- Add new use cases
- Write API documentation
- Create tutorials or examples

## Development Environment Setup

Please refer to the [Development Guide](doc/development.md) to learn how to set up your local development environment.

## Code Standards

To maintain code consistency and readability, we require all code to follow these standards:

1. **Go Code Style**
   - Follow Go language's standard code style
   - Use `gofmt` to format code
   - Use `golint` to check code quality

2. **Naming Conventions**
   - Package names use lowercase letters, no underscores
   - Function names use camelCase, with uppercase first letter for public functions
   - Variable names use camelCase, with lowercase first letter for private variables

3. **Comment Standards**
   - Add comments for public functions, types, and variables
   - Comments should clearly describe the function and purpose of the code
   - Use GoDoc format for comments

## PR Submission Process

1. **Fork the Repository**: Fork the GoStaticServe repository on GitHub to your own account

2. **Clone the Repository**: Clone the forked repository to your local machine

   ```bash
   git clone https://github.com/YOUR_USERNAME/GoStaticServe.git
   cd GoStaticServe
   ```

3. **Create a Branch**: Create a new branch for your changes

   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bugfix-name
   ```

4. **Make Changes**: Make code modifications on your branch

5. **Commit Changes**: Commit your changes with a clear commit message

   ```bash
   git add .
   git commit -m "Concise commit message"
   ```

6. **Push Changes**: Push your branch to GitHub

   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create Pull Request**: Create a Pull Request on GitHub, describing your changes and the problem you're solving

## Testing

Before submitting code, please ensure your code passes all tests. You can run tests using the following command:

```bash
go test ./...
```

If you add new features, please also add corresponding test cases.

## Feedback

If you encounter any issues during the contribution process, you can get help through:

- Submitting an issue on GitHub
- Contacting project maintainers

Thank you for your support and contributions to the GoStaticServe project!