# `pkg`
[![Maintainability](https://api.codeclimate.com/v1/badges/2ec596b9250233491a69/maintainability)](https://codeclimate.com/github/BillotP/pkg/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/2ec596b9250233491a69/test_coverage)](https://codeclimate.com/github/BillotP/pkg/test_coverage)

This repository holds commonly used packages that provide various self-contained utilities and business logic helpers.

Here is a brief summary of a selection of packages:

### `pkg/log`

Provides a ready-to-use logger that reports to Sentry by default, wraps `"github.com/upfluence/log"` behind the scenes.

### `pkg/timeutil` & `pkg/timeutil/timetest`

Wraps commonly used time functions so that they can be easily mocked later.

### `pkg/group`

Facilitates a common concurrency pattern by spawning a groups of goroutines working on a same task.
Handles error cancellation and propagation.

### `pkg/cfg`

Fetches custom values from environment variables, or default values if they don't exist.

### `pkg/currency`

Abstracts the money type and provides an interface, a default implementation (based on the EU central bank API) and helpers to exchange amounts of money to various currencies

### `pkg/slices`

Provides common operations on slices missing from `"slices"`.

### `pkg/pointers`

Can return a pointer to a value. And other pointer stuff.
