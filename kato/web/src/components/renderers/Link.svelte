<script>
  import { parseLinks } from "../../libs/renderers/extractors";

  export let body = "";

  const links = parseLinks(body);
</script>

{#if links.length < 1}
  <h3>No links...</h3>
{/if}

<div class="urls">
  {#each links as link}
    <a class="link" href={link.href} target="_blank">{link.label}</a>
  {/each}
</div>

<style>
  h3 {
    text-align: center;
  }
  a.link {
    font-size: 2rem;
  }

  .urls {
    display: flex;
    flex-direction: column;
  }
</style>
