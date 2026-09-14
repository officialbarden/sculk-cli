# sculk-cli
a CLI-app for project initialisation, inspired by Nodejs' Package Manager (NPM) & Rust's Cargo.

## Todo

### Goals
- [ ] Make a website - Sculk users can find available libraries to download on this website.
- [ ] Conflict-less merging of libraries
- [ ] storing all installed libraries into a `libraries.json` (similar to package.json)

if `sculk` is initialised *inside* a datapack/resourcepack, then merge libraries with the datapack/resourcepack
if `sculk` is initialised *outside* a datapack, install all libraries under a 'sculk-libraries' directory.

### Commands
- [ ] an 'install' command to install all libraries mentioned in `libraries.json`.

