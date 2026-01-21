# otp

`otp` is a package for Go, designed to Generate otp and Verify otp.

## Features

- **Generate Otp**: Generates otp for any length of digit.
- **Verify Otp**: It verifys the otp.
- **Send Email**: It will send the otp to the email you want to send.

## Installation 

To install `otp`, run:

```sh
go get github.com/parthvinchhi/otp
```

## Usage
Here's an example of how to use otp:

```
    package main

    import (
        "fmt"

        "github.com/parthvinchhi/otp"
    )

    func main() {
        value := otp.GenerateOtp(6)
        fmt.Println(value)

        ok := otp.VerifyOtp(value, 1234)
        fmt.Println(ok)
    }

```

## Usage
Here's an example of how to send otp:

```
    package main

    import (
        "fmt"

        "github.com/parthvinchhi/otp"
    )

    func main() {
        value := otp.GenerateOtp(6)

        cfg := otp.NewEmailConfig(
            "YOUR_EMAIL",
            "EMAIL_ID_PASSWORD",
            "SMTP_SERVER",
            "SMTP_PORT",
        )

        if err := cfg.SendEmail([]string{"SEND_TO_EMAIL"}, value); err != nil {
            log.Fatalf("Error sending otp via email: %v", err)
        }
    }

```