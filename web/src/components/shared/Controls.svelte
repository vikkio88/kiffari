<script>
  export let background = false;
</script>

<div class="controls" class:background>
  <slot>Here be controls</slot>
</div>

<style>
  .controls {
    position: fixed;
    bottom: 0;
    right: 0;
  }

  .background {
    padding: 0.5rem 1rem;
    border-radius: 8px;
    background-color: var(--main-bg-color);
  }
</style>
