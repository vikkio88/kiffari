<script>
  export let open = true;
</script>

<div class="header">
  <slot name="header">Header</slot>
  <button on:click={() => (open = !open)}>
    {#if open}
      🔼
    {:else}
      🔽
    {/if}
  </button>
</div>
{#if open}
  <slot name="content">Content</slot>
{/if}

<style>
  .header {
    border-bottom: 1px white solid;
    padding: 2rem;
  }

  .header > button {
    position: absolute;
    right: 0;

  }
</style>
