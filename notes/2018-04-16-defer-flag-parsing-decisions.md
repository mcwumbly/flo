# Defer decision about flag parsing

This is still very experimental, and sorting out
the exact UX for the CLI feels like a distraction.

Rather than choose a fully-featured flag parsing
library like [cobra](https://github.com/spf13/cobra)
or something more light-weight like [jhanda](https://github.com/pivotal-cf/jhanda),
for now we are simply deferring this decision and
living with the built-in flag parsing library, and
intentionally not fussing over details like command
or flag names.

That is to say, this is very open to being revisited
when other things are further along.
