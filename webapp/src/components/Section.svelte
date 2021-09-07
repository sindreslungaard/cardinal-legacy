<script>
    import {Link, navigate} from "svelte-navigator";

    export let tab;
</script>

<main class="relative flex px-4 mb-2 gap-4 h-full">
    <section class="flex-1 tabs flex-1 w-full h-full flex flex-col">
        <menu class="flex-none" role="tablist">
            <button on:click={() => navigate("/")} role="tab" aria-selected="{tab == "index"}">Overview</button>
            <button on:click={() => navigate("/containers")} role="tab" aria-selected="{tab == "containers"}">Containers</button>
            <button on:click={() => navigate("/servers")} role="tab" aria-selected="{tab == "servers"}">Servers</button>
        </menu>
        
        <article role="tabpanel" class="flex-1">
            <slot></slot>
        </article>
    
    </section>

    <aside class="logs hidden lg:block flex-none h-full">
        <div class="w-full h-full backdrop">
            
        </div>
    </aside>
</main>

<style>
    button:focus {
        outline: none !important;
    }
</style>