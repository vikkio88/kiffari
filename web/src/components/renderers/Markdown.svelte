<script>
  import SvelteMarkdown from "svelte-markdown";
  import CopyCode from "./shared/CopyCode.svelte";
  export let note = {};
  const { body } = note;
</script>

<div class="mdBody">
  <SvelteMarkdown source={body} />
</div>
<div class="tools">
  <CopyCode {body} />
</div>

<style>
  .mdBody {
    min-height: 30vh;
  }

  .tools {
    font-size: smaller;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
  }
</style>
