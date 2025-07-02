# Numscript WASM CLI

This project is a WebAssembly (WASM) port of the [formancehq/numscript](https://github.com/formancehq/numscript) CLI, enabling you to run Numscript execution in any WASM-compatible environment. Originally forked from the official repository, it now provides a standalone WASM binary that maintains full compatibility with the original CLI functionality.

## What is Numscript?

Numscript is a Domain-Specific Language (DSL) designed to help you model complex financial transactions, replacing complex and error-prone custom code with easy-to-read, declarative scripts. Originally developed by Formance, Numscript is used to express financial transactions within the Formance ledger system.

The language provides:
- **Declarative syntax** for modeling financial operations
- **Built-in validation** for transaction integrity  
- **Account management** with metadata support
- **Variable substitution** for dynamic scripts
- **Feature flags** for experimental functionality

You can try Numscript online at [playground.numscript.org](https://playground.numscript.org/).

## Installation

### Prerequisites

You'll need a WebAssembly (WASM) runtime to execute the numscript-wasm binary. This guide uses Wasmtime, but you can also use other WASM runtimes like Wasmer.

**Install Wasmtime:**
```bash
curl https://wasmtime.dev/install.sh -sSf | bash
```

### Download the Binary

Download the WASM binary from the [latest release](https://github.com/PagoPlus/numscript-wasm/releases/latest) for your platform.

## Usage

The CLI provides two commands: `version` and `run`.

### Version Command

Display the current version of the CLI:

```bash
wasmtime numscript.wasm version
```

**Output:**
```bash
# If you are using the released binary this will print the real app version:
v0.1.0
# Otherwise will print dev:
dev
```

### Run Command

Execute a Numscript with the provided input data. The command reads JSON input from stdin containing the script, variables, account metadata, balances, and feature flags.

#### Input Format

The input JSON must follow this structure:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `script` | string | ✅ | The Numscript code to execute |
| `variables` | object | ❌ | [Variables](https://docs.formance.com/modules/numscript/reference/variables) |
| `metadata` | object | ❌ | [Account metadata](https://docs.formance.com/modules/numscript/reference/metadata) |
| `balances` | object | ✅ | Current account balances. [Example below](#balances). |
| `featureFlags` | object | ❌ | Experimental features to enable (empty objects as values). [Example below](#featureflags) |

##### balances
Is an object containing the account name and their respective assets and balances. Ex:
```json
  "balances": {
    "foo": {
      "USD/2": 200,
      "EUR/2": 100
    },
    "bar": {
      "BRL/2": 200,
      "JPY/2": 100
    }
  },
```
###### Asset Format
Assets use the format `CURRENCY/PRECISION`:
- `USD/2` = US Dollars with 2 decimal places
- `EUR/2` = Euros with 2 decimal places  
- `BTC/8` = Bitcoin with 8 decimal places

##### featureFlags
Is an object containing the name of the feature flag you want to use with an empty object as their value. Ex:
```json
  "featureFlags": {
    "experimental-oneof": {},
    "experimental-get-amount-function": {},
    "experimental-mid-script-function-call": {}
  }
```

Numscript does not have an official list of feature flags, but you can see the [source code](https://github.com/formancehq/numscript/blob/v0.0.17/internal/flags/flags.go) and the documentation with the [experimental features usage](https://github.com/formancehq/numscript/blob/main/differences-with-machine.md#new-functionalities-feature-flags)

#### Examples

##### Basic Transfer
Simple account-to-account transfer:

1. **Create an input file** (`example.json`):

```json
{
  "script": "send [USD/2 100] (\n  source = @foo\n  destination = @bar\n)",
  "balances": {
    "foo": {
      "USD/2": 200,
      "EUR/2": 100
    }
  },
  "variables": {},
  "metadata": {},
  "featureFlags": {}
}
```

2. **Execute the script**:

```bash
wasmtime numscript.wasm run < example.json
```

3. **Expected output** (JSON format):

```json
{
  "postings": [
    {
      "source": "foo",
      "destination": "bar", 
      "asset": "USD/2",
      "amount": 100
    }
  ],
  "txMeta": {},
  "accountsMeta": {}
}
```

#### Using variables

1. **Input File:** (`example.json`):
```json
{
  "script": "vars { account $user monetary $fee portion $tax } send $fee (source = $user destination = { $tax to @platform:tax remaining to @platform:revenue })",
  "balances": {
    "users:1234": {
      "USD/2": 10000
    }
  },
  "variables": {
    "user": "users:1234",
    "fee": "USD/2 100",
    "tax": "20%"
  },
  "metadata": {},
  "featureFlags": {
    "experimental-oneof": {},
    "experimental-get-amount-function": {},
    "experimental-mid-script-function-call": {}
  }
}
```

2. **Execute the script**:
```bash
wasmtime numscript.wasm run < example.json
```

3. **Expected output** (JSON format):
```json
{
  "postings": [
    {
      "source": "users:1234",
      "destination": "platform:tax",
      "amount": 20,
      "asset": "USD/2"
    },
    {
      "source": "users:1234",
      "destination": "platform:revenue",
      "amount": 80,
      "asset": "USD/2"
    }
  ],
  "txMeta": {},
  "accountsMeta": {}
}
```


#### Using metadata

1. **Input File:** (`example.json`):
```json
{
  "script": "vars { account $order account $merchant = meta($order, \"merchant\") portion $commission = meta($merchant, \"commission\") } send [USD/2 *] ( source = $order destination = { $commission to @platform:fees remaining to $merchant })",
  "balances": {
    "orders:2345": {
      "USD/2": 1000
    }
  },
  "variables": {
    "order": "orders:2345"
  },
  "metadata": {
    "orders:2345": {
      "merchant": "merchants:1234"
    },
    "merchants:1234": {
      "commission": "15%"
    }
  },
  "featureFlags": {
    "experimental-oneof": {},
    "experimental-get-amount-function": {},
    "experimental-mid-script-function-call": {}
  }
}
```

2. **Execute the script**:
```bash
wasmtime numscript.wasm run < example.json
```

3. **Expected output** (JSON format):
```json
{
  "postings": [
    {
      "source": "orders:2345",
      "destination": "platform:fees",
      "amount": 150,
      "asset": "USD/2"
    },
    {
      "source": "orders:2345",
      "destination": "merchants:1234",
      "amount": 850,
      "asset": "USD/2"
    }
  ],
  "txMeta": {},
  "accountsMeta": {}
}
```

#### Allowing experimental features

If you want to use experimental features, you need to enable them by adding their respective feature flag:

1. **Input File:** (`example.json`):
```json
{
  "script": "vars { monetary $mon = [USD/2 100] number $n = get_amount($mon) } send [USD/2 $n] (source = oneof { @foo @bar } destination = @baz)",
  "balances": {
    "foo": {
      "USD/2": 99
    },
    "bar": {
      "USD/2": 100
    }
  },
  "variables": {},
  "metadata": {},
  "featureFlags": {
    "experimental-oneof": {},
    "experimental-get-amount-function": {},
    "experimental-mid-script-function-call": {}
  }
}
```

2. **Execute the script**:
```bash
wasmtime numscript.wasm run < example.json
```

3. **Expected output** (JSON format):
```json
{
  "postings": [
    {
      "source": "bar",
      "destination": "baz",
      "amount": 100,
      "asset": "USD/2"
    }
  ],
  "txMeta": {},
  "accountsMeta": {}
}
```

## Error Handling

The CLI will exit with status code 1 and write error messages to stderr if:

- The input JSON is malformed
- The Numscript has parsing errors
- The execution fails due to insufficient balances or other runtime errors

Example error output:
```
Exception: Not enough funds. Needed [USD/2 100] (only [USD/2 99] available)
```

## License

This project maintains the same license as the original [formancehq/numscript](https://github.com/formancehq/numscript) repository.

## Contributing

This is a fork focused on WASM compatibility. For core Numscript language contributions, please consider contributing to the upstream [formancehq/numscript](https://github.com/formancehq/numscript) repository. 