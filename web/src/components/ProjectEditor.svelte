<script>
  import { onMount } from "svelte";
  import LinksEditor from "./shared/LinksEditor.svelte";

  let input;
  export let id = null;
  export let name = null;
  export let description = null;
  export let config = { wip_limit: 0 };
  export let links = [];

  export let onSave = (projectDetails) => {
    console.log(projectDetails);
  };
  export let onClose = null;

  function onSaveInternal() {
    onSave({ id, name, description, links, config });
  }

  onMount(() => {
    input.focus();
  });

  // TODO: add links editor and config editor
</script>

<form on:submit|stopPropagation|preventDefault={onSaveInternal}>
  <div class="editor">
    <input
      bind:this={input}
      type="text"
      required
      bind:value={name}
      placeholder="New Project title..."
    />
    <textarea
      placeholder="A small project description..."
      cols="20"
      rows="4"
      required
      bind:value={description}
    />

    <LinksEditor bind:links />
  </div>
  <div class="ctrls">
    {#if Boolean(onClose)}
      <button
        class="smaller"
        title="Edit"
        on:click|preventDefault|stopPropagation={onClose}
      >
        ❌
      </button>
    {/if}
    <button>Save 💾</button>
  </div>
</form>

<style>
  .editor {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 0.5rem;
  }

  .editor input {
    padding: 1em;
    font-size: 16px;
    border-radius: 10px;
  }

  .editor textarea {
    font-size: 16px;
    padding: 1em;
    border-radius: 10px;
  }

  .ctrls {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    padding: 0 1rem 1rem 0;
  }
</style>
