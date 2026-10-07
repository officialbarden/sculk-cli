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
