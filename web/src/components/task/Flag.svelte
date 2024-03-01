<script>
  let input;
  export let flag = null;
  let clicked = false;
</script>

<div>
  {#if !Boolean(flag) && !clicked}
    <button on:click|stopPropagation|preventDefault={() => (clicked = true)}>
      🚩
    </button>
  {:else}
    <input
      bind:this={input}
      bind:value={flag}
      type="text"
      placeholder="Flag Description..."
    />
    <button
    class="smaller"
      on:click|stopPropagation|preventDefault={() => {
        flag = null;
        clicked = false;
      }}
    >
      ❌
    </button>
  {/if}
</div>

<style>
  div {
    padding: 1rem 2rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    gap: 0.5rem;
    font-weight: bold;
  }
  /*TODO: Standardize the input css */
  input {
    padding: 1em;
    border-radius: 10px;
    width: 90%;
  }
</style>
