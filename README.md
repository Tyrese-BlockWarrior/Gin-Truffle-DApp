# Gin-Truffle-DApp

Ethereum Certificate DApp built using Gin, Truffle & Solidity.

## 🛠 Built With

<div align="left">
<a href="https://go.dev/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/go.svg" width="36" height="36" alt="Go" /></a>
<a href="https://gin-gonic.com/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/gin.svg" width="36" height="36" alt="Gin" /></a>
<a href="https://nodejs.org/en/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/nodejs.svg" width="36" height="36" alt="NodeJS" /></a>
<a href="https://trufflesuite.com" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/truffle.svg" width="36" height="36" alt="Truffle" /></a>
<a href="https://soliditylang.org/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/solidity.svg" width="36" height="36" alt="Solidity" /></a>
</div>

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/DEMYSTIF/gin-truffle-dapp.git
cd gin-truffle-dapp
```

Install truffle

```bash
npm i -g truffle
```

Compile contract

```bash
truffle compile
```

Run a blockchain (ganache/geth/foundry) on port 8545

Deploy contract

```bash
truffle migrate
```

Install abigen

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for contract

```bash
abigen --abi Cert.abi --pkg main --type Cert --out Cert.go
```

Start the application

```bash
go run .
```

## 📜 License

Distributed under the MIT License.

## 🎗️ Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/<feature_name>`)
3. Commit your Changes (`git commit -m '<feature_name>_added'`)
4. Push to the Branch (`git push origin feature/<feature_name>`)
5. Open a Pull Request
