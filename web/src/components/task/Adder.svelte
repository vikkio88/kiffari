<script>
  import { tick, createEventDispatcher } from "svelte";

  const d = createEventDispatcher();

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

    d("taskSubmitted", { title: value });
    value = "";
    clicked = false;
  }
</script>

<form
  class="addWrapper"
  class:centered={clicked}
  on:submit|preventDefault|stopPropagation={submit}
>
  {#if !clicked}
    <button class="smaller" on:click|preventDefault|stopPropagation={onAdd}>
      ➕
    </button>
  {:else}
    <input
      bind:this={input}
      bind:value
      type="text"
      placeholder="Task title..."
    />
    <div>
      <button class="smaller">✅</button>
      <button
        class="smaller"
        on:click|preventDefault|stopPropagation={() => (clicked = false)}
        >❌</button
      >
    </div>
  {/if}
</form>

<style>
  .addWrapper {
    min-width: 80%;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
    margin-top: 1rem;
  }

  .centered {
    justify-content: center;
    gap: 1rem;
  }

  input {
    padding: 1em;
    border-radius: 10px;
    width: 60%;
  }
</style>
