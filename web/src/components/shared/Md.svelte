<script>
  import SvelteMarkdown from "svelte-markdown";

  export let source = "";
</script>

<div class="mdBody">
  <SvelteMarkdown {source} />
</div>

<style>
  .mdBody {
    min-height: 30vh;
  }
</style>
