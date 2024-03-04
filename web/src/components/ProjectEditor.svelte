<script>
  import { onMount } from "svelte";

  let input;
  export let id = null;
  export let name = null;
  export let description = null;
  export let config = { wip_limit: 0 };
  export let links = [];

  export let onSave = (projectDetails) => {
    console.log(projectDetails);
  };

  function onSaveInternal() {
    onSave({ id, name, description, links, config });
  }

  onMount(() => {
    input.focus();
  });
</script>

<form class="editor" on:submit|stopPropagation|preventDefault={onSaveInternal}>
  <input
    bind:this={input}
    type="text"
    required
    bind:value={name}
    placeholder="New Project title..."
  />
  <textarea cols="20" rows="4" required bind:value={description} />

  {#if Array.isArray(links) && links.length > 0}
    <div class="links">
      <strong>🔗 Links</strong>
      {#each links as link}
        <div class="link">
          <a href={link.href} target="_blank">{link.label}</a>
        </div>
      {/each}
    </div>
  {/if}
  <button>Save</button>
</form>

<style>
  .editor {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
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
</style>
