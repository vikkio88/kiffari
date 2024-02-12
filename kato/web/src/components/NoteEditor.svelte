<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";
  import Controls from "./shared/Controls.svelte";
  import { formatDTL } from "../libs/dates";

  export const note = null;

  export let title = "new note title";
  export let text = "";
  export let tags = [];
  export let dueDate = null;
  let dueDateProxy = Boolean(dueDate)
    ? formatDTL(dueDate)
    : null;
  function clearDueDate() {
    dueDateProxy = null;
  }

  let showPreview = false;

  const PLUGIN_SETUP_STRING = "<!--\nPlugin: \n-->\n";

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

  export let onSave = (note) => {
    console.log(note);
  };

  function onSaveInternal() {
    onSave({
      title,
      body: text,
      tags,
      due_date: Boolean(dueDateProxy) ? new Date(dueDateProxy) : null,
    });
  }

  function setupPlugin(event) {
    showPreview = false;
    if (!text.includes("<!--")) {
      text = `${PLUGIN_SETUP_STRING}${text}`;
    }

    event.currentTarget.selectionStart = 11; //this is the end of `Plugin: ` string
    // TODO: check whether tick() here would work
    // event.currentTarget.focus();
  }

  let showAdditionalInfo = Boolean(dueDate);
  function toggleAdditionalInfo() {
    showAdditionalInfo = !showAdditionalInfo;
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
      <button on:click|stopPropagation|preventDefault={setupPlugin}>
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
    <div class="additionalInfo">
      <button
        class="toggler"
        on:click|stopPropagation|preventDefault={toggleAdditionalInfo}
      >
        {#if !showAdditionalInfo}
          More ➕
        {:else}
          ◀️ Less
        {/if}
      </button>
      {#if showAdditionalInfo}
        <div>
          <div>
            <label for="dueDate">Due Date:</label>
            <input
              name="dueDate"
              type="datetime-local"
              bind:value={dueDateProxy}
              min={new Date().toISOString()}
            />
            <button on:click|stopPropagation|preventDefault={clearDueDate}
              >❌</button
            >
          </div>
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

  .additionalInfo {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-start;
  }

  .additionalInfo input {
    padding: 0.6rem;
    font-size: 18px;
  }

  .controls {
    display: flex;
    flex-direction: row;
  }
</style>
