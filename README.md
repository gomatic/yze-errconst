# yze-errconst

A [`yze`](https://github.com/gomatic/yze) analyzer (category `errors`) enforcing the gomatic Go standard that errors are sentinel constants (`errs.Const`): it flags `errors.New(...)` and `fmt.Errorf(...)` calls that do not wrap a cause with `%w`.

- **Rule:** `yze/errconst`
- **Library:** exports `Analyzer` and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-errconst` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

## The `%w` proxy (v1 behavior)

An `fmt.Errorf` call is allowed when its **literal** format string contains `%w`; every other `fmt.Errorf` is flagged. The `%w` token is the proxy for "this wraps an existing error rather than creating a new one". The analyzer deliberately does **not** trace the wrapped argument back to a constant sentinel: the convention permits adding context to *either* a sentinel *or* an injected cause, and an injected cause is by definition a non-sentinel error value, so requiring a sentinel argument would wrongly reject the permitted injected-cause case. Two consequences follow, both intentional and pinned by fixtures:

- `fmt.Errorf("%w: ...", cause)` is allowed even when `cause` is a non-sentinel error value.
- A const-held format string (rather than a string literal) is not inspected, so it escapes the check.

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
