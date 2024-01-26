<script>
  import { navigate } from "svelte-routing";
  export let onTagClick = (tag) => () => navigate(`/tags/${tag.id}`);
  export let tags = [];
  export let big = false;
</script>

{#if tags.length < 1}
  <strong>No tags...</strong>
{:else}
  {#each tags as tag}
    <button class:big class="tag" on:click={onTagClick(tag)}>
      {tag.label}
    </button>
  {/each}
{/if}

<style>
  .tag {
    border: 2px white solid;
    border-radius: 10px;
    padding: 0.5em;
    font-size: small;
    font-weight: bold;
  }

  .big {
    font-size: 18px;
    padding: 1em;
  }

  .tag:hover {
    border-color: #646cff;
    cursor: pointer;
  }
</style>
