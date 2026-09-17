<script lang="ts">
    import Icon from 'astro-iconset/svelte'

    const prefix = "sculk"
    const cli = new Map<string, string>();
    cli.set("init", `${prefix} init --dp namespace 26.2`)
    cli.set("install", `${prefix} install`)
    cli.set("add", `${prefix} add id-system ...`)
    cli.set("remove", `${prefix} remove id-system ...`)
    cli.set("uninstall", `${prefix} uninstall id-system ...`)
    let cliString = $state(cli.get("init"))

    function setString(command: string) {
        cliString = cli.get(command)!
    }

    function copyToClipboard() {
        navigator.clipboard.writeText(cliString!)
    }
    
</script>

<!-- command-line examples -->
<div class=" bg-black my-10 w-full md:w-100 rounded-xl min-h-10 border border-white/10">
    <!-- buttons -->
    <div class="flex text-white px-3">
        <button onclick={() => {setString("init")}} class="p-2 border-b border-transparent hover:border-white/30 cursor-pointer">init</button>
        <button onclick={() => {setString("add")}} class="p-2 border-b border-transparent hover:border-white/30 cursor-pointer">add</button>
        <button onclick={() => {setString("install")}} class="p-2 border-b border-transparent hover:border-white/30 cursor-pointer">install</button>
        <button onclick={() => {setString("uninstall")}} class="p-2 border-b border-transparent hover:border-white/30 cursor-pointer">uninstall</button>
        <button onclick={() => {setString("remove")}} class="p-2 border-b border-transparent hover:border-white/30 cursor-pointer">remove</button>
    </div>
    
    <div class="flex items-center gap-2 px-3 py-1.5">
        <p class="text-white font-ui px-2 py-3 gap-2 ">
            {cliString}
        </p>
        <button onclick={() => {copyToClipboard()}} class="ml-auto text-white text-[18px] hover:cursor-pointer hover:bg-white/15 duration-150 p-2 rounded-lg">
            <Icon name="mdi:content-copy" />
        </button>
    </div>
</div>