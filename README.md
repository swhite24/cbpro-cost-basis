# cbpro-cost-basis

`cbpro-cost-basis` provides a utility for quickly determing cost basis on purchases made from [Coinbase Pro](https://pro.coinbase.com).

## Prerequisites

- Verified Coinbase Pro account with instant deposits
- Obtain an API Key  
  https://docs.pro.coinbase.com/#authentication

## Installation

Check the [releases](https://github.com/swhite24/cbpro-cost-basis/releases) page and grab the binary appropriate for your system.

## Features

- Determine cost basis for purchases of a specific product over a specific period of time

## Usage

```sh
cbpro-cost-basis helps calculate cost basis on purchases from Coinbase Pro over a specified period of time

Usage:
  cbpro-cost-basis [flags]

Flags:
      --end string          End date of order fills to calculate cost basis. (2021-01-01)
  -h, --help                help for cbpro-cost-basis
      --key string          Coinbase Pro API key
      --passphrase string   Coinbase Pro API key passphrase
      --product string      Product to use when calculating (default "BTC-USD")
      --secret string       Coinbase Pro API key secret
      --start string        Start date of order fills to calculate cost basis. (2021-01-01)
```

## Example

```sh
$ cbpro-cost-basis --start 2021-01-01

Product: BTC-USD
Start Date: 2019-01-01
End Date: 2020-01-01
Total Amount Purchased: x.xxxxxxxx
Total Cost: xxx.xx
Cost Basis: xxx.xx

```

## License

See [LICENSE.txt](LICENSE.txt)
