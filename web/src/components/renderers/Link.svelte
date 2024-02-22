<script>
  import { parseLinks } from "../../libs/renderers/extractors";

  export let body = "";

  const links = parseLinks(body);
</script>

<div class="wrapper">
  <h3>🔗</h3>
  {#if links.length < 1}
    <h3>No links...</h3>
  {/if}

  <div class="urls">
    {#each links as link}
      <a class="link" href={link.href} target="_blank">{link.label}</a>
    {/each}
  </div>
</div>

<style>
  h3 {
    text-align: center;
  }
  a.link {
    font-size: 2rem;
  }

  .wrapper {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .urls {
    display: flex;
    flex-direction: column;
  }
</style>
