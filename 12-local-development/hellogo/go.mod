module github.com/danzBraham/hellogo

go 1.22.2

// Be aware that using the "replace" keyword isn't advised,
// but can be useful to get up and running quickly.
// The proper way to create and depend on modules is to publish them to a remote repository.
// When you do that, the "replace" keyword can be dropped from the go.mod
replace github.com/danzBraham/gostrings v0.0.0 => ../gostrings

require (
  github.com/danzBraham/gostrings v0.0.0
)