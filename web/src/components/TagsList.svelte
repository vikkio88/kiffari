<script>
  import { navigate } from "svelte-routing";
  export let onTagClick = (tag) => () => navigate(`/tags/${tag.id}`);
  export let tags = [];
  export let big = false;
</script>

{#if tags.length < 1}
  <h3>No tags... 🤷</h3>
{:else}
  {#each tags as tag}
    <button class:big class="tag" on:click={onTagClick(tag)}>
      <span class="hash">#</span>{tag.label}
    </button>
  {/each}
{/if}

<style>
  .tag {
    border: solid 2px var(--light-gray);
    border-radius: var(--default-borders-radius);
    padding: 0.5rem 1rem;
    margin: 0 0.5rem;
  }

  .hash {
    font-size: 1.2rem;
    color: var(--link-color);
    margin-right: 0.3rem;
  }

  .big {
    font-size: var(--input-medium-font-size);
    padding: var(--input-default-pad);
  }

  .tag:hover {
    color: var(--link-color);
    background-color: var(--light-gray);
    border-color: var(--border-color);
    cursor: pointer;
  }
</style>
