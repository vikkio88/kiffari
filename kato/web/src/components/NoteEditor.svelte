<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";
  import Controls from "./shared/Controls.svelte";

  export const note = null;

  let showPreview = false;
  export let title = "new note title";
  export let text = "a new note body";
  export let tags = []

  async function handleKeydown(event) {
    if (event.key !== "Tab") return;

    event.preventDefault();

    const { selectionStart: start, selectionEnd: end, value } = this;
    text = value.substring(0, start) + "  " + value.substring(end);
    await tick();
    this.selectionStart = this.selectionEnd = start + 2;
  }

  function onTagSelected(e) {
    tags = e.detail.tags;
  }

  function formatTags() {
    return tags.map((t) => {
      if (t["$created"]) {
        return { label: `${t.value ?? t.id}` };
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
  <TagSearch on:added_tag={onTagSelected} initialTags={tags} />
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
  <Controls>
    <button on:click={onSaveInternal}>Save</button>
  </Controls>
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
</style>
