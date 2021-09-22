# destinycli

So... I wanted to learn Go!

## Purpose

This package is not meant to be taken seriously. I am using it as a way to teach myself [https://golang.org/](Go) and the many new _(to me)_ concepts that come along with a strictly typed language.

I've always found that when teaching myself a new concept that working with strings has helped me ease into it. And it just so happens that a command line application is just that. Input a few strings, do some process with said strings, output more strings.

## Proposed Outcome

I want to create a simple command line application that will query the [Destiny API](https://bungie-net.github.io/multi/index.html) for item information. If possible, I would like to expand on this, but that might not happen.

There is no real need for this type of application within the greater Destiny community. It's just something I thought of while reading through a few Go tutorials.

## Proposed API

Fetch an item and display information about it:

**Input**

```shell
$~ destinycli -i Lorentz Driver
```

**Output**

```shell
Fetching information on Lorentz Driver

Lorentz Driver

Charge Time     - 533
Impact          - ######### (12)
Range           - ########### (18)
Stability       - ############ (24)
Handling        - ###### (8)
Reload Speed    - ####### (10)
Magazine        - 6
```

English will be the default return language, but the API has 12 other supported languages. To support these, we just need to add a flag so the user can select their language.

```shell
$~ destinycli -i Lorentz Driver --lang=es-mx
```

We might even be able to add support for a local config so that the client would not have to do this. Maybe read a local OS level variable that could give use this info. Again, part of the reason I'm doing this is to work through these sorta problems and solve them. That command could look something like:

```shell
$~ destinycli --config-set lang=de
```

On that note, we could let the user set their bungie id, or preferred platform. That sorta stuff. But that's getting ahead of ourselves.
