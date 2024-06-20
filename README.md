# PortPatrol

PortPatrol integrates Go's concurrent network scanning with Rust's robust packet analysis. It offers efficient detection of active hosts and open ports, coupled with deep inspection features for superior network security and performance.

## Overview

PortPatrol consists of two main components:

- **Go Application**: Implements a network scanner using Go's concurrency features.
- **Rust Library**: Provides packet analysis and deep inspection capabilities.

## Features

- **Network Scanner**: Utilizes Go's concurrency model to scan ports and detect open services.
- **Packet Analysis**: Rust component for deep inspection of network packets, ensuring robust security analysis.
- **Cross-Language Integration**: Demonstrates seamless integration of Go and Rust for enhanced performance and functionality.

## Installation

### Prerequisites

- Go (v1.16 or higher)
- Rust (v1.55 or higher)

### Build Instructions

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd PortPatrol
   ```

2. Build the Rust library:

   ```bash
   cd rust-lib
   cargo build --release
   ```

3. Build the Go executable:

   ```bash
   cd ../go-app
   go build -o portpatrol
   ```

## Usage

Run the built executable portpatrol to initiate network scanning and packet analysis. Example usage:

    ./portpatrol --target 192.168.1.1

## License

This project is licensed under the MIT License - see the LICENSE file for details.
