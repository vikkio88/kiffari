<script>
  export let tags = [];
</script>

{#if tags.length < 1}
  <strong>No tags...</strong>
{:else}
  <div class="list">
    {#each tags as tag}
      <div class="tag">{tag.label}</div>
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
    color: #747bff;
    cursor: pointer;
  }
</style>
