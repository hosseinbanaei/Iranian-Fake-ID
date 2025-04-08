# 🎯 Fake National ID Generation with Go Language

This project is a simple educational tool that generates fake national ID numbers using the **Go (Golang)** programming language.
The generated IDs include a valid check digit and can be used for **testing and educational purposes**.

⚠️ This project is strictly intended for **educational and test data generation** purposes. Its use in official or illegal situations is **strongly prohibited**.

---

## 🔎 Structure of Iranian National ID

An Iranian national ID consists of **10 digits**:

- **The first 3 digits:** Issuer's district code (in real IDs).
- **Digits 4 to 9:** Individual identifier (for distinguishing individuals).
- **The 10th digit:** Check digit or validation digit.

📌 In this project, for simplicity, **the first 9 digits are generated randomly**, so **the first 3 digits, which represent the issuance location, are also random**. Therefore, the generated IDs are only algorithmically valid but do not match the real administrative structure.

---

## 🧮 Check Digit Calculation Algorithm

1. Take the first 9 digits.
2. Multiply each digit by a number from 10 to 2.
3. Sum these multiplications.
4. Calculate the remainder of the sum divided by 11.
5. If the remainder is less than 2, the check digit is equal to the remainder.
6. If the remainder is 2 or more, the check digit is equal to **11 minus the remainder**.
7. If the final result is 10 or 11, the check digit will be **zero**.

---

## ⚙️ How to Run the Program

1. Make sure Go is installed on your system. (use `go version` command)
    - You can download and install Go on `https://go.dev/dl`.
2. Save the `main.go` file or clone the project.
3. Run the program with the following command:

```bash
go run main.go
```

4. During execution, the program will ask you how many national IDs you want to generate.

---

## 💻 Sample Output

```
How many fake national IDs should be generated? 3
Generated national IDs:
1. 4278951062
2. 3087154975
3. 5120936241
```

---

## ⚠️ Important Notice

The generated IDs in this project are only for **practice and testing purposes** and are not valid according to the Iranian Civil Registration Organization’s standards.
Any illegal use or misuse of this project **is the responsibility of the user**.

---

## 📝 License

This project is licensed under the **MIT** license.
You are free to use it for personal and educational purposes.

---

## 👤 Author

This project was created by [Hossein Banaei](https://github.com/hosseinbanaei).
If you find it useful, I would appreciate a [⭐](https://github.com/hosseinbanaei/Iranian-Fake-ID/stargazers) on GitHub.