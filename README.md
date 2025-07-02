# Numscript WASM CLI

This project was originally a fork of the original [formancehq/numscript](https://github.com/formancehq/numscript) repository, adapted for WASM compatibility while maintaining the core CLI functionality.

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

Clone the repository and build the binary:

```bash
git clone git@github.com:PagoPlus/numscript-wasm.git

cd numscript-wasm

go build -o numscript-wasm main.go
```

This will create a `numscript-wasm` executable in the current directory.

## Usage

The CLI provides two commands: `version` and `run`.

### Version Command

Display the current version of the CLI:

```bash
./numscript-wasm version
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

| Field | Type | Description |
|-------|------|-------------|
| `script` | string | The Numscript code to execute |
| `variables` | object | [variables](https://docs.formance.com/numscript/reference/variables) used inside the script. |
| `metadata` | object | [metada variables](https://docs.formance.com/numscript/reference/metadata) |
| `balances` | object | Current account balances (nested object: account -> asset -> amount) |
| `featureFlags` | object | Experimental features to enable (empty object) |

#### Complete Example

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
./numscript-wasm run < example.json
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

#### Allowing experimental features

If you want to use experimental features, you need to enable them by adding their respective feature flag:

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