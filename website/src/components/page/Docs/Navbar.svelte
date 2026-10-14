<script>
    import { isSidebarToggled } from "../../../store.svelte.ts";
    import Icon from 'astro-iconset/svelte'
</script>

<div class="p-2 w-full h-15 border-b border-dim font-ui flex flex-row items-center">

    <!-- Sidebar Button (Hidden on PC) -->
    <button onclick={() => {
        isSidebarToggled.set(true)
    }} id="docs-lsidebar-toggle" class="block md:hidden text-3xl p-2">
        <Icon name="akar-icons:settings-horizontal"/>
    </button>


</div>
