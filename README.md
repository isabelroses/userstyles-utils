<p align="center">
    <img alt="Sylus Logo" src="assets/stylus.webp" width="100">
    <h2 align="center">Userstyles Utils</h3>
</p>

<p align="center">
    A small collection of utilities for development on <a href="https://github.com/catppuccin/userstyles">catppuccin/userstyles</a>.
</p>

### Table of contents

<!--toc:start-->
- [Local utilities](#local-utilities)
  - [Usage](#usage)
  - [Dependencies](#dependencies)
- [External utilities](#external-utilities)
<!--toc:end-->

### Local utilities

- [count](./src/count.sh) - A fetch style script to collect data on userstyles as a whole
- [missing](./src/missing.sh) - Find missing elements in userstyles metadata
- [umeta](./src/umeta.sh) - Quickly userstyles metadata
- [who](./src/who.sh) - Find out quickly who is the maintainer is

#### Usage

I advise against installing these scripts golbally, and highly recommend you download them on a need-to-use basis.

Or you can run it via nix:
```sh
nix run github:isabelroses/userstyles-utils#<script>
```

#### Dependencies

- [yq](https://github.com/mikefarah/yq)
- [curl](https://curl.se/)
- [gum](https://github.com/charmbracelet/gum)
- [coreutils](https://www.gnu.org/software/coreutils/)


### External utilities

- [Userstyles importer](https://github.com/uncenter/catppuccin-all-userstyles-import) - Quickly import userstyles
- [The userstyles helper](https://github.com/uncenter/ctp-userstyles-helper) - Make reviewing new userstyles easier
