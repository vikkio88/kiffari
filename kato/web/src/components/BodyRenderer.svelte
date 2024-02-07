<script>
  import SvelteMarkdown from "svelte-markdown";
  import { mdConfigParser } from "../libs/mdConfigParser";
  import Code from "./renderers/Code.svelte";
  //TODO: add MD force view

  export let body = "";
  let config = mdConfigParser(body);

  const pluginMapper = {
    Code: Code,
  };
</script>

{#if pluginMapper[config.Plugin]}
  <svelte:component this={pluginMapper[config.Plugin]} {body} />
{:else}
  <SvelteMarkdown source={body} />
{/if}
