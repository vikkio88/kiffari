<script>
  let clicked = false;
  export let onConfirmed = () => console.log("Confirmed");

  function handleClick() {
    if (!clicked) {
      clicked = true;
      return;
    }
  }
</script>

<button on:click={handleClick}>
  {#if !clicked}
    <slot />
  {:else}
    <div class="internal">
      Confirm?
      <button on:click|stopPropagation={() => (clicked = false)}>❌</button>
      <button on:click|stopPropagation={() => onConfirmed()}>👍</button>
    </div>
  {/if}
</button>

<style>
  .internal {
    display: flex;
    place-items: center;
    gap: 0.5em;
  }
  .internal > button {
    font-size: xx-small;
  }
</style>
