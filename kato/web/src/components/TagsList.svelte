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
    border: solid 2px #333;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    margin: 0 0.5rem;
  }

  .big {
    font-size: 18px;
    padding: 1em;
  }

  .tag:hover {
    color: #646cff;
    background-color: #333;
    border-color: white;
    cursor: pointer;
  }
</style>
