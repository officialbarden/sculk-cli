<script>
    import { isSidebarToggled } from "../../../store.svelte.ts";
    let { children } = $props();
    
    let displaySidebar = isSidebarToggled
</script>

{#if displaySidebar.get()}
    <div class="sticky">
        <div class="bg-main-dark/75 backdrop-blur-xl absolute top-0 md:relative md:flex flex-col px-4 min-w-65 md:max-w-200 h-full md:min-h-screen font-ui border border-dim">
            {@render children()}
        </div>
    </div>
{/if}