<script>
  export let open = true;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
  class="header crs-pointer"
  on:click={() => (open = !open)}
  aria-pressed="false"
  role="button"
  tabindex="-1"
>
  <slot name="header">Header</slot>
  <div>
    {#if open}
      🔽
    {:else}
      ◀️
    {/if}
  </div>
</div>
{#if open}
  <slot name="content">Content</slot>
{/if}

<style>
  .header {
    border-bottom: 1px white solid;
    padding: 2rem;
    display: grid;
    grid-template-columns: 90% 10%;
  }

  .header:hover {
    background-color: #3a3a3a;
  }
</style>
