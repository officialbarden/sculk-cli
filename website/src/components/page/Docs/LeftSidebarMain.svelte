<script>
    import { isSidebarToggled } from "../../../store.svelte.ts";
    import Icon from 'astro-iconset/svelte'

    let { children } = $props();

    let displaySidebar = isSidebarToggled
    let open = $derived(displaySidebar.get())

    // Animate Sidebar
</script>

<div
    class="
        md:hidden fixed top-0 left-0 z-50 h-dvh w-64
        overflow-y-auto flex flex-col px-4
        bg-main-dark/75 backdrop-blur-xl font-ui border border-dim
        transition-transform duration-300 ease-out
        {open ? 'translate-x-0' : '-translate-x-full'}"
    inert={!open}
    aria-hidden={!open}
>
    {@render children()}

    <!-- Buttons -->
    <div class="mt-auto py-4">
        <button
            class="p-1 border border-dim rounded-lg text-2xl"
            onclick={() => displaySidebar.set(false)}
        >
            <Icon name="akar-icons:chevron-left-small" />
        </button>
    </div>
</div>

<!-- Desktop: sticky column -->
<div class="hidden md:block">
    <div class="
        sticky top-0 h-screen overflow-y-auto
        flex flex-col px-4
        bg-main-dark/75 backdrop-blur-xl
        font-ui border border-dim">
        {@render children()}
    </div>
</div>