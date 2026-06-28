# yze-go-errconst

A [`yze`](https://github.com/gomatic/yze) analyzer (group `go`, category `errors`) enforcing the gomatic Go standard that errors are sentinel constants (`errs.Const`): it flags `errors.New(...)` and `fmt.Errorf(...)` calls that do not wrap a cause with `%w`.

- **Rule:** `yze/go/errconst`
- **Library:** exports `Analyzer` and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-go-errconst` runs it standalone (`text`/`-json`/`-fix`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
