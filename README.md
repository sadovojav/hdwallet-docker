# Hierarchical Deterministic Wallet in Docker
Python-based library for the implementation of a hierarchical deterministic wallet generator for more than 140+ multiple cryptocurrencies. It allows the handling of multiple coins, multiple accounts, external and internal chains per account and millions of addresses per chain.

## Usage
```shell script
git clone https://github.com/sadovojav/hdwallet-docker.git
cd hdwallet-docker
docker compose up -d
```

```curl
curl 'http://localhost:8080?symbol=BTC'
```

<details open>
  <summary>Output</summary><br/>

```json5
{
    "cryptocurrency": "Tron",
    "symbol": "TRX",
    "network": "mainnet",
    "strength": 128,
    "entropy": "266680c642e2732b2bdee756998135d6",
    "mnemonic": "chapter cross cover magic cheese night quarter table final slot estate promote",
    "language": "english",
    "passphrase": null,
    "seed": "51c4e98b434a4ac0981d2c40bf04b68e4b86c9c4234f52e760226edc7b51373bbee209bc5684c105df63a54bf24e3fe55bafe6cb52725def634ec74f2e7c274a",
    "root_xprivate_key": "xprv9s21ZrQH143K3Y8L5ip6Z9iAaggF3H8s4BqAppGHHhxhrUiZZxSZCZkZmtepSd7c7YDheZNfVBr8AxG8dNQY3NHiQqJz83tWKWxj1E2LJHd",
    "root_xpublic_key": "xpub661MyMwAqRbcG2CoBkM6vHeu8iWjSjriRQkmdCftr3VgjH3i7VkokN53dCeB2vc8tULG3uMuQ4gRHDSSUc2pytymnQTbjSVxgsT8aibc3vd",
    "xprivate_key": "xprvA3u1UbyRNZRYK2KfEPpVPEbZTL1A3zm9odbrbBZ5PXn2TD8JmMKfpbQ5tpjoL97YgeZBHtpSyRgoLmePPMWF7r71mnoBqMdb6HxPYitntuD",
    "xpublic_key": "xpub6GtMt7WKCvyqXWQ8LRMVkNYJ1MqeTTV1ArXTPZxgwsK1L1TTJtdvNPiZk7eeCJTiCKURbobtXTWrww3hDcPmAxDQYuMxb7eEX7a5mNLr6EG",
    "uncompressed": "1a2ecdc532b30c965791971b1a1af8949c5330fcd1ed7c4200e2cd52a10509341495e9f16c31ab3ec3cfe367a948b48f018e0197ac39fcabe1912a20b48854d7",
    "compressed": "031a2ecdc532b30c965791971b1a1af8949c5330fcd1ed7c4200e2cd52a1050934",
    "chain_code": "4768f69e764673ed8e7549eda1b78d985312f441d740071d5bae3338f69f36cd",
    "private_key": "2283085b88ac8ef9c3d5af76966df10158ca31b117bdf9fc5f55883a6efedcdc",
    "public_key": "031a2ecdc532b30c965791971b1a1af8949c5330fcd1ed7c4200e2cd52a1050934",
    "wif": "KxNoCi6241vuAQbNs46NkSPW8eL1rwQQnkebYzazoNC9F4QxW5Hy",
    "finger_print": "283a2b2f",
    "semantic": "p2pkh",
    "path": "m/44'/195'/0'/0/0",
    "hash": "283a2b2f3c940e2d7e124db3fb44d6ddf557f85a",
    "addresses": {
        "p2pkh": "TQ4P6QLSdoJzVuzDxw1Dx48ReiV3TyBjtS",
        "p2sh": "3Bob6RgmcZzsTjegD2BCiS4wy5gXc3vcxt",
        "p2wpkh": "bc1q9qazkteujs8z6lsjfkelk3xkmh6407z6h3jhrr",
        "p2wpkh_in_p2sh": "37xonYHzobauumo8z8sM6hHE9gKjxReYAV",
        "p2wsh": "bc1qz3t28u2mvdg9ejwd8mg94v7mpnfnsjc32326ve6dnjyq57l0saksnupx0y",
        "p2wsh_in_p2sh": "33GkBu869xCfcYeouBu5qAr7543oT5SSwF"
    }
}
```
</details>
