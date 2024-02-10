<script>
    import SvelteMarkdown from "svelte-markdown";

    export let body = "";
</script>

<div class="mdBody">
    <SvelteMarkdown source={body} />
</div>

<style>
    .mdBody {
        min-height: 40vh;
    }
</style>
