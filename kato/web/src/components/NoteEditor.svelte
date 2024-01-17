<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";

  let showPreview = false;
  let text = `a new note`;

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
        return { label: `${t.id}` };
      }
      return t;
    });
  }

  export let onSave = (body, tags) => {
    console.log({ body, tags });
  };

  function onSaveInternal() {
    const tags = formatTags();
    onSave(text, tags);
  }
</script>

<div class="editor">
  <TagSearch on:added_tag={onTagSelected} />
  <button on:click={() => (showPreview = !showPreview)}>Toggle Preview</button>
  <div class="note">
    {#if !showPreview}
      <div>
        <textarea
          bind:value={text}
          rows="10"
          cols="50"
          on:keydown={handleKeydown}
        />
      </div>
    {:else}
      <div class="preview">
        <SvelteMarkdown source={text} />
      </div>
    {/if}
  </div>
  <div class="controls">
    <button on:click={onSaveInternal}>Save</button>
  </div>
</div>

<style>
  .editor {
    padding: 2em;
  }
  .note textarea {
    font-size: 16px;
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
