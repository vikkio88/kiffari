<script>
  import { navigate } from "svelte-routing";
  export let tags = [];
</script>

{#if tags.length < 1}
  <strong>No tags...</strong>
{:else}
  <div class="list">
    {#each tags as tag}
      <button class="tag" on:click={() => navigate(`/tags/${tag.id}`)}>
        {tag.label}
      </button>
    {/each}
  </div>
{/if}

<style>
  .list {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5em;
    padding: 2em;
  }
  .tag {
    border: 2px white solid;
    border-radius: 10px;
    padding: 0.5em;
    font-size: small;
    font-weight: bold;
  }

  .tag:hover {
    border-color: #646cff;
    cursor: pointer;
  }
</style>
