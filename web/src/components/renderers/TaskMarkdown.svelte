<script>
  import CopyCode from "./shared/CopyCode.svelte";
  import Md from "../shared/Md.svelte";
  export let task = {};
  const { description } = task;
</script>

<Md source={description} />
<div class="tools">
  <CopyCode body={description} />
</div>

<style>
  .tools {
    font-size: smaller;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
  }
</style>
