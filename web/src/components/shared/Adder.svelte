<script>
  import { tick, createEventDispatcher } from "svelte";

  const d = createEventDispatcher();

  export let title = "Add";
  export let placeholder = "Name/Title...";
  export let rightAligned = false;
  export let centered = false;
  let input;
  let clicked = false;

  let value = "";
  async function onAdd() {
    clicked = true;
    await tick();
    input.focus();
  }

  function submit() {
    if (!clicked) {
      return;
    }

    d("added", value);
    value = "";
    clicked = false;
  }
</script>

<form
  class="addWrapper"
  class:centered={clicked || centered}
  class:rightAligned
  on:submit|preventDefault|stopPropagation={submit}
>
  {#if !clicked}
    <button
      class="smaller"
      on:click|preventDefault|stopPropagation={onAdd}
      {title}
    >
      ➕
    </button>
  {:else}
    <input bind:this={input} bind:value type="text" {placeholder} />
    <div>
      <button class="smaller">✅</button>
      <button
        class="smaller"
        on:click|preventDefault|stopPropagation={() => (clicked = false)}
      >
        ❌
      </button>
    </div>
  {/if}
</form>

<style>
  .addWrapper {
    min-width: 80%;
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 1rem;
  }

  .rightAligned {
    align-items: flex-end;
  }

  .centered {
    justify-content: center;
    gap: 1rem;
  }

  input {
    padding: 1em;
    border-radius: 10px;
    width: 60%;
    font-size: 1.3rem;
  }
</style>
