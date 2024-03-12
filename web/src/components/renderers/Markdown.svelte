<script>
  import SvelteMarkdown from "svelte-markdown";
  import CopyCode from "./shared/CopyCode.svelte";
  import Md from "../shared/Md.svelte";
  export let note = {};
  const { body } = note;
</script>

<Md source={body} />
<div class="tools">
  <CopyCode {body} />
</div>

<style>
  .tools {
    font-size: smaller;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
  }
</style>
