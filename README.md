<p align="center">
    <p align="center">Advent of Code</p>
</p>
<p align="center"><h1 align="center">AOC</h1></p>
<p align="center">
	<img src="https://img.shields.io/github/license/zsoltdzsugan/aoc?style=default&logo=opensourceinitiative&logoColor=white&color=00f2ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/zsoltdzsugan/aoc?style=default&logo=git&logoColor=white&color=00f2ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/zsoltdzsugan/aoc?style=default&color=00f2ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/zsoltdzsugan/aoc?style=default&color=00f2ff" alt="repo-language-count">
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
  - [ Example](#-example)

---

##  Overview

<code>❯ This is my Advent of Code solutions repositoty. Where you can find all the challenges with my solutions.</code>

---

##  Features

<code>❯ CLI option implemented (for fun)</code>

---

##  Project Structure

```sh
└── aoc/
    ├── LICENSE
    ├── README.md
    ├── cmd
    │   ├── helper
    │   │   └── getModulePath.go
    │   ├── root.go
    │   ├── show_challenge.go
    │   ├── show_puzzle.go
    │   └── show_solution.go
    ├── go.mod
    ├── main.go
    ├── txt_reader
    │   └── txt_reader.go
    └── year2024
        └── day1
            ├── helper
            │   └── sortData.go
            ├── part1
            │   ├── challenge.md
            │   └── solution.go
            ├── part2
            │   ├── challenge.md
            │   └── solution.go
            └── puzzle.txt
```


###  Project Index
<details open>
	<summary><b><code>AOC/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/main.go'>main.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/go.mod'>go.mod</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- year2024 Submodule -->
		<summary><b>year2024</b></summary>
		<blockquote>
			<details>
				<summary><b>day1</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/year2024/day1/puzzle.txt'>puzzle.txt</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
					<details>
						<summary><b>part2</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/year2024/day1/part2/solution.go'>solution.go</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
					<details>
						<summary><b>helper</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/year2024/day1/helper/sortData.go'>sortData.go</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
					<details>
						<summary><b>part1</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/year2024/day1/part1/solution.go'>solution.go</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- txt_reader Submodule -->
		<summary><b>txt_reader</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/txt_reader/txt_reader.go'>txt_reader.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- cmd Submodule -->
		<summary><b>cmd</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/cmd/root.go'>root.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/cmd/show_puzzle.go'>show_puzzle.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/cmd/show_challenge.go'>show_challenge.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/cmd/show_solution.go'>show_solution.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
			<details>
				<summary><b>helper</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/zsoltdzsugan/aoc/blob/master/cmd/helper/getModulePath.go'>getModulePath.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with aoc, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules


###  Installation

Install aoc using one of the following methods:

**Build from source:**

1. Clone the aoc repository:
```sh
❯ git clone https://github.com/zsoltdzsugan/aoc
```

2. Navigate to the project directory:
```sh
❯ cd aoc
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build -o aoc
```

###  Usage
Run aoc using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ aoc
```
### Example
Run aoc with flags. Show Challenge or Solution requires all 3 flags while puzzle input only needs year and day.
```sh
❯ aoc challenge -y 2024 -d 1 -p 1 # --year 2024 --day 1 --part 1 also works
```
```sh
❯ aoc solution -y 2024 -d 1 -p 1
```
```sh
❯ aoc puzzle -y 2024 -d 1
```
