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
  <div class="accordionControls" class:inverted={open}>◀️</div>
  <slot name="header">Header</slot>
</div>
{#if open}
  <slot name="content">Content</slot>
{/if}

<style>
  .header {
    padding: 1.5rem;
    display: flex;
    flex-direction: row;
    gap: 1rem;
    align-items: center;
    border-bottom: var(--thin-borders);
  }

  .header:hover {
    background-color: var(--div-item-hover-color);
  }

  .accordionControls {
    transition: transform 0.4s;
    transform: rotate(180deg);
  }

  .inverted {
    transform: rotate(270deg);
    transition: transform 0.4s;
  }
</style>
