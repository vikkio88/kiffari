<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";
  import Controls from "./shared/Controls.svelte";

  export const note = null;

  export let title = "new note title";
  export let text = "";
  export let tags = [];
  let showPreview = false;

  async function handleKeydown(event) {
    if (event.key !== "Tab") return;

    event.preventDefault();

    const {
      selectionStart: start,
      selectionEnd: end,
      value,
    } = event.currentTarget;
    text = value.substring(0, start) + "  " + value.substring(end);
    await tick();
    this.selectionStart = this.selectionEnd = start + 2;
  }

  function onTagsSelection(e) {
    tags = e.detail;
  }

  export let onSave = (title, body, tags) => {
    console.log({ title, body, tags });
  };

  function onSaveInternal() {
    onSave(title, text, tags);
  }
</script>

<div class="editor">
  <form on:submit|preventDefault={onSaveInternal}>
    <div class="controls">
      <button
        on:click|stopPropagation|preventDefault={() =>
          (showPreview = !showPreview)}
      >
        {#if showPreview}
          Code 🤓
        {:else}
          Markdown 😎
        {/if}
      </button>
      <button on:click|stopPropagation|preventDefault={() => {
        showPreview = false;
        if (text.includes("<!--")) return;
        text = `<!--\nPlugin: \n-->\n${text}`;
      }}>
        Plugin ⚙️
      </button>
    </div>
    <div class="note">
      {#if !showPreview}
        <div class="form">
          <input
            required
            class="title"
            bind:value={title}
            on:focus={(e) => {
              e.currentTarget.select();
            }}
            type="text"
          />
          <textarea
            required
            placeholder="A New Note body..."
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
    <TagSearch on:updatedSelection={onTagsSelection} selectedTags={tags} />
    <Controls>
      <button type="submit">Save</button>
    </Controls>
  </form>
</div>

<style>
  .form {
    display: flex;
    flex-direction: column;
  }
  .title {
    padding: 1em;
    font-size: 18px;
    border-radius: 10px;
  }
  .editor {
    padding: 2em;
  }
  .note textarea {
    font-size: 16px;
    padding: 1em;
    border-radius: 10px;
  }

  .preview {
    display: block;
    text-align: left;
  }

  .preview h2 {
    text-align: center;
  }

  .controls {
    display: flex;
    flex-direction: row;
  }
</style>
