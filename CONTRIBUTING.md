# Contributing to The Cause #

Welcome to ScamMateo, a product of BS workshops! Don't bother looking us up, because we don't exist!

All works of ScamMateo are fully owned by ScamMateo. All contributors to ScamMateo agree to grant ScamMateo a non-exclusive,
perpetual license to reproduce, display, modify, and profit from their contributions. What did you expect when you submitted
something to a place called ScamMateo?

## Formatting ##

Please make a pull request submitting new scams to the scam folder. You can submit a scam to scams/<my scam>.md filename.
You MUST include a title block with at least the following fields, with the date being to date to publish the article:
 - title
 - keyword
 - author
 - date

A> See scams/MrPizzaManNowOfferingNewServices.md for an example

Markdown is rendered to HTML through the [golang markdown package] (https://pkg.go.dev/github.com/gomarkdown/markdown "golang markdown documentation")

This means that you will have access to a number of extensions in addition to the standard markdown syntax.
 1. [commonmark markdown spec] (https://spec.commonmark.org/0.30/ "commonmark markdown specification")
 2. [mmark markdown extensions] (https://mmark.miek.nl/post/syntax/ "mmark markdown extensions")
 3. [golang markdown extra extensions] (https://pkg.go.dev/github.com/gomarkdown/markdown#readme-extensions "extra extensions in golang markdown package")

## Checking your work ##

If you set up your repo as described in README.md, then there are two ways you can check that your markdown will display
as you expect.

1. devbox installs the `glow` cli tool for rendering markdown in the terminal. Just open a terminal (vscode terminal or
   similar should be fine), navigate to this directory, and run `glow scams/<my scam>.md`. Note that this isn't perfect,
   but it should give you a reasonable idea of how it will look and it's really quick.
2. Start the dev server (by running `air`), open a web browser, and navigate to your new document. Dev mode servers should
   allow viewing all documents, not just those which have publish dates in the past.
