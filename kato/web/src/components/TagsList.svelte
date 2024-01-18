<script>
  export let tags = [];
</script>

{#if tags.length < 1}
  <h3>No tags...</h3>
{:else}
  {#each tags as tag}
    <h3>{tag.label}</h3>
  {/each}
{/if}
