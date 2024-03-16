<script>
  import { mdConfigParser } from "../libs/renderers/mdConfigParser";
  import Markdown from "./renderers/TaskMarkdown.svelte";
  import Todos from "./renderers/TaskTodos.svelte";

  export let task = {};
  let config = mdConfigParser(task?.description ?? {});

  const pluginMapper = {
    todo: Todos,
    default: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.plugin] ?? pluginMapper.default}
  {task}
/>
