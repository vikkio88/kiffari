<script>
  import SvelteMarkdown from "svelte-markdown";
  import { onMount, tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";

  export const note = null;

  let showPreview = false;
  let text = `a new note`;
  let title = "a title";

  let selectedTags = [];

  async function handleKeydown(event) {
    if (event.key !== "Tab") return;

    event.preventDefault();

    const { selectionStart: start, selectionEnd: end, value } = this;
    text = value.substring(0, start) + "  " + value.substring(end);
    await tick();
    this.selectionStart = this.selectionEnd = start + 2;
  }

  function onTagSelected(e) {
    selectedTags = e.detail.tags;
  }

  function formatTags() {
    return selectedTags.map((t) => {
      if (t["$created"]) {
        return { label: `${t.value}` };
      }
      return t;
    });
  }

  export let onSave = (title, body, tags) => {
    console.log({ title, body, tags });
  };

  function onSaveInternal() {
    const tags = formatTags();
    onSave(title, text, tags);
  }
</script>

<div class="editor">
  <TagSearch on:added_tag={onTagSelected} />
  <button on:click={() => (showPreview = !showPreview)}>Toggle Preview</button>
  <div class="note">
    {#if !showPreview}
      <div class="form">
        <input class="title" bind:value={title} type="text" />
        <textarea
          bind:value={text}
          rows="10"
          cols="50"
          on:keydown={handleKeydown}
        />
      </div>
    {:else}
      <div class="preview">
        <h2>{title}</h2>
        <SvelteMarkdown source={text} />
      </div>
    {/if}
  </div>
  <div class="controls">
    <button on:click={onSaveInternal}>Save</button>
  </div>
</div>

<style>
  .form {
    display: flex;
    flex-direction: column;
  }
  .title {
    padding: 1em;
    font-size: 18px;
  }
  .editor {
    padding: 2em;
  }
  .note textarea {
    font-size: 16px;
    padding: 1em;
  }

  .preview {
    display: block;
  }

  .controls {
    position: absolute;
    bottom: 0;
    right: 0;
    padding: 2em;
  }
</style>
