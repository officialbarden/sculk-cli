## Todo

### Goals
- [ ] Make a website - Sculk users can find available libraries to download on this website.
- [X] ~~Conflict-less merging of libraries~~
- [X] ~~storing all installed libraries into a `libraries.json` (similar to package.json)~~

### Commands
- [X] ~~Handle library versioning, updating.~~
- [X] ~~An 'uninstall' command to remove files without causing conflicts.~~
- [X] ~~an 'install' command to install all libraries mentioned in `libraries.json`.~~

 <br>

# Welcome to Sculk-CLI
Sculk CLI is a CLI-app for project initialisation, datapack library* handling, sharing and merging, inspired by Nodejs' Package Manager (NPM), Python's PIP & Rust's Cargo.

## Why should I use sculk-cli?

Similar to npm, pip and cargo, sculk-cli (or just 'sculk') is meant to be a CLI tool that allows datapack developers to seemlessly integrate libraries created by other datapack developers and verified by sculk devs. 

Generally, the process of installing a library* from an external source, like Github, Smithed or Modrinth, takes a while as the developer has to navigate to the library*'s project page, choose the version that is satisfactory to them, click install and drag-drop-extract the zip file in their working directory. Sculk-cli makes it all possible in ~3 command-line tools!

Similarly, because there is no direct Mojang support for what a default datapack can be (except the '/datapack create' command), it gets annoying to create the same #minecraft:load or #minecraft:tick files everytime you want to setup a new datapack project. Sculk fixes that by creating the directories with just one command!


# Initializing a 'sculk project'

Open the shell *inside* your datapack folder (e.g. /saves/WORLD_NAME/datapacks/<test_datapack>)

```shell
sculk init --dp/--rp <namespace> 26.3
```

This will *initialize* a new 'sculk project' inside the datapack.
```
/test_datapack
    ├── data
    │     ├── minecraft/tags/function
    │     │     ├── tick.json
    │     │     └── load.json
    │     └── <namespace>/function/global
    │     │     └── load.mcfunction
    │     │     └── tick.mcfunction
    ├── pack.mcmeta
    └── libraries.json
```

libraries.json is a json file used by **sculk** to find/track/log libraries. During initialization, it'll look something like this:
```json
{
   "author": "USER",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": []
}
```
The fields in libraries.json are pretty self explanatory. 

The 'version' field is a semantic way for a sculk developer to versonify their projects / libraries. It also allows sculk-cli to find updated libraries for the same game version, in a sculk-project*.

The 'game_version' field is important, as it tells the sculk-cli which game version is your sculk-project* at the current moment is compatible with.

Notes:
- sculk-project: A sculk project can be both - a **datapack** or a **resourcepack**. Hence, instead of specifying a datapack/resourcepack as is, the docs use the term 'sculk-project'.

# Installing a sculk library*

Just like how a developer may download a particular package, say for example: the pandas library* of python, the command for it in command-line would be: `pip install pandas`

Similarly, sculk-cli uses acronyms/identifiers to identify certain libraries/projects. Sculk doesn't have its own package storage and hosting platform, so it sources the packages from Git Hosting platforms like GitHub and Codeberg. Internally, each unique "identifier" is mapped to a source repository, which is cloned and seemlessly merged with the current project.

Run the following command:
```
sculk add id-system
```

Sculk will then reach in its internal hashmap, get the source link and get the contents of the library*, and merge them with your sculk-project as is. Your libraries.json will change, and a new library* will be added in the 'libraries[]' field.

```json
{
   "author": "USER",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": [
      {
         "identifier": "id-system",
         "name": "ID System",
         "source": "http://github.com/officialbarden/id-system",
         "version": "1.0.0",
         "game_version": "26.2"
      }
   ]
}
```

Here, you can see that the library* information is stored inside your `libraries.json` file. You can share this .json file with other sculk developers, and ask them to run the command `sculk install` to install all packages specified in the 'libraries[]' field.

If you're working on an unstable version, like a snapshot: you can use the `--ignore` flag to ignore Game Version Mismatch Checking and install packages directly meant for other Game Versions into your sculk project. 

*Note: By using the --ignore flag, you are to take full responsibility of how the library interacts with your sculk project, as some additional stuff, like tags, advancements, predicates and more may get installed and merged into your project without a warning!*

## Why have a seperate libraries.json if the library* contents are merged in the final datapack, anyways?

Great Question! Sculk isn't just a 'library* installer & merger', it's a full blown package manager! (well, a barebones version of a traditional package manager). By having a seperate 'libraries.json' file in the sculk-project's root directory, sculk developers can keep track of what libraries exist in their codebase! At the same time, say a sculk-library* developer pushed a new version of the library* onto their git repository, by running one command in the shell (`sculk update`), sculk developers can pull the changes into their files without the hassle of merging new updates by hand.


Notes:
- library: in npm & pip, they're called packages. In cargo, they're called crates and in sculk, they're called libraries. Though, we don't object to interchangability of the two terms: library and package.

# Creating/Publishing a sculk library
You can create and submit your sculk libraries in our [sculk-discord](
https://discord.gg/JWZkAgsyry) for approval and direct integration into the cli. As long as your datapack outputs into a standard datapack, you may use whatever pre-compiler you desire however, for seemless integration with sculk, here are some rules you will have to follow while writing your sculk library.

1. Use a unique namespace<br>
Generally, we'll encourage library publishers to use the same as the identifier they wish to associate their sculk library with. However, incase there are issues like the identifier slug being pre-used, we ask developers to change their namespaces to a much more customized and unique slug. You can look at the [wiki-guide](https://minecraft.wiki/w/Identifier#Legal_characters) to learn what characters you can add to your namespace to make it more unique. <br><br>
As best practice, we would encourage developers to follow [Smithed's Conventions](https://docs.smithed.dev/conventions/index.html).

2. Do Not Use Extremely Minimal Namespaces/Filenames <br>
Same as the first rule, to prevent potential merge conflicting with other libraries, we hope that your naming convention, especially if you're writing to namespaces that are extremely common (namespaces like 'minecraft', 'code', 'namespace' etc.) for whatever reason, your file name must be unique.

Once you've followed the above recommendations/rules and created a library, make a `libraries.json` and paste the following code inside it:
```json
{
   "author": "AUTHOR-NAME",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": []
}
```

Fields like author are currently only semantic. The 'version' field and 'game_version' are what you should pay attention to, as sculk will look at these values while making decisions during the merging of your library.
