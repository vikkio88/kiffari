<script>
  export let title = "";
  export let cancellingTitle = "Cancel";
  export let classes = ""
  export let confirmLabel = "Confirm?";
  export let onConfirmed = () => console.log("Confirmed");
  let clicked = false;

  function handleClick() {
    if (!clicked) {
      clicked = true;
      return;
    }
  }
</script>

<button {title} class={classes} on:click={handleClick}>
  {#if !clicked}
    <slot />
  {:else}
    <div class="internal">
      {confirmLabel}
      <button title={cancellingTitle} on:click|stopPropagation={() => (clicked = false)}>❌</button>
      <button {title} on:click|stopPropagation={() => onConfirmed()}>👍</button>
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
