
# Lemon CLI Tool

Lemon is a lightweight command-line tool designed for cooking measurement conversions. It's a simple yet powerful utility to help you quickly convert between various units for cooking.

## Features
- Convert between volume units (cups, milliliters, liters, tablespoons, teaspoons, etc.)
- Convert between weight units (grams, ounces, pounds, kilograms, etc.)
- Easy to use with intuitive commands
- Portable: Can be run without installation or installed to your system with PATH integration

## Installation Options

### 1. Portable Use
You can simply download the `lemon.exe` file and use it portably, without installing it. Just run the executable from any directory and start converting!

#### Example:
```bash
lemon --value=2 --from=cups --to=milliliters
```

### 2. Installer with PATH Integration
Alternatively, you can use the provided installer `LemonSetup.exe` to automatically add `lemon` to your system's PATH. This allows you to use `lemon` from any terminal window without specifying the full path to the executable.

#### Installation Steps:
1. Download the `LemonSetup.exe` installer.
2. Run the installer, and it will place `lemon.exe` in the appropriate directory and add it to your PATH automatically.
3. Once installed, you can use the `lemon` command from any terminal.a

#### Example:
```bash
lemon --value=500 --from=grams --to=ounces
```

## Requirements
- **Windows 10 or later**
- No additional dependencies are required for portable use.

## Usage
```bash
lemon --help
```
This command will show a list of available options and conversions.

## License
Lemon is distributed under the MIT License.
