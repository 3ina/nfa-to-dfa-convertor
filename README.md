<p align="center">
    <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" align="center" width="30%">
</p>
<p align="center"><h1 align="center">NFA-TO-DFA-CONVERTOR</h1></p>
<p align="center">
	<em><code>❯ REPLACE-ME</code></em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/3ina/nfa-to-dfa-convertor?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/3ina/nfa-to-dfa-convertor?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/3ina/nfa-to-dfa-convertor?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/3ina/nfa-to-dfa-convertor?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

<code>❯ The NFA to DFA Converter is a Go-based implementation that facilitates the conversion of Non-deterministic Finite Automata (NFA) to Deterministic Finite Automata (DFA). This project provides a structured approach to define, manage, and visualize automata, making it a valuable tool for educational purposes, compiler design, and formal language processing.
</code>

---

##  Features

- **NFA Representation**: Defines the structure for Non-deterministic Finite Automata, including states, transitions, and the handling of epsilon transitions.
  
- **DFA Representation**: Implements the structure for Deterministic Finite Automata, providing methods to visualize states, transitions, and final states.

- **Conversion Logic**: Contains algorithms to convert an NFA to a DFA, ensuring that the resulting DFA accurately represents the behavior of the original NFA.

- **State Management**: Efficiently manages sets of states and their transitions, ensuring that all unique states are accounted for during the conversion process.

- **Epsilon Closure Calculation**: Computes the closure of states that can be reached via epsilon transitions, a critical aspect of the conversion process.

- **User-Friendly Output**: Provides methods to print the DFA’s states, transitions, start state, and final states for easy visualization and debugging.

---

##  Project Structure

```sh
└── nfa-to-dfa-convertor/
    ├── automata
    │   ├── dfa.go
    │   ├── helper.go
    │   └── nfa.go
    ├── go.mod
    └── main.go
```


###  Project Index
<details open>
	<summary><b><code>NFA-TO-DFA-CONVERTOR/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/3ina/nfa-to-dfa-convertor/blob/master/main.go'>main.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/3ina/nfa-to-dfa-convertor/blob/master/go.mod'>go.mod</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- automata Submodule -->
		<summary><b>automata</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/3ina/nfa-to-dfa-convertor/blob/master/automata/dfa.go'>dfa.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/3ina/nfa-to-dfa-convertor/blob/master/automata/helper.go'>helper.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/3ina/nfa-to-dfa-convertor/blob/master/automata/nfa.go'>nfa.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with nfa-to-dfa-convertor, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules


###  Installation

Install nfa-to-dfa-convertor using one of the following methods:

**Build from source:**

1. Clone the nfa-to-dfa-convertor repository:
```sh
❯ git clone https://github.com/3ina/nfa-to-dfa-convertor
```

2. Navigate to the project directory:
```sh
❯ cd nfa-to-dfa-convertor
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```




###  Usage
Run nfa-to-dfa-convertor using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```


---
##  Project Roadmap

- [X] **`Task 1`**: <strike>Implement automata package .</strike>
- [ ] **`Task 2`**: tview for UI.
- [ ] **`Task 3`**: Implement feature three.

---

##  Contributing

- **💬 [Join the Discussions](https://github.com/3ina/nfa-to-dfa-convertor/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/3ina/nfa-to-dfa-convertor/issues)**: Submit bugs found or log feature requests for the `nfa-to-dfa-convertor` project.
- **💡 [Submit Pull Requests](https://github.com/3ina/nfa-to-dfa-convertor/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/3ina/nfa-to-dfa-convertor
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/3ina/nfa-to-dfa-convertor/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=3ina/nfa-to-dfa-convertor">
   </a>
</p>
</details>

---

##  License

This project is protected under the MIT License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/mit/#) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

---
