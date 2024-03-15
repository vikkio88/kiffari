<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "./TagSearch.svelte";
  import Controls from "./shared/Controls.svelte";
  import { formatDTL, nowString } from "../libs/dates";
  import { addDays, } from "date-fns";
  import { removeComments } from "../libs/renderers/cleanup";
  import {
    PLUGIN_SETUP_STRING,
    generatePlugin,
  } from "../libs/renderers/helpers";

  export const note = null;

  export let title = nowString();
  export let text = "";
  export let tags = [];
  export let dueDate = null;
  let dueDateProxy = Boolean(dueDate) ? formatDTL(dueDate) : null;
  function clearDueDate() {
    dueDateProxy = null;
  }

  let showPreview = false;

  async function handleKeydown(event) {
    /*
    this is to intercept words in case we want to match trending tags
    if (event.key === " ") {
      console.log(text.split(" ").pop());
      return;
    }
    */
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
  export function setupPlugin(event) {
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
      <button
        on:click|stopPropagation|preventDefault={() => {
          text = `${generatePlugin("Link")}${removeComments(text)}`;
        }}
      >
        Plugin:Link 🔗
      </button>
      <button
        on:click|stopPropagation|preventDefault={() => {
          text = `${generatePlugin("Todo")}${removeComments(text)}`;
        }}
      >
        Plugin:Todo ✅
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
          ◀️
        {/if}
      </button>
      {#if showAdditionalInfo}
        <div class="dueDateWrapper">
          <h3>Due Date</h3>
          <div>
            <input
              name="dueDate"
              type="datetime-local"
              bind:value={dueDateProxy}
              min={new Date().toISOString()}
            />
            {#if Boolean(dueDateProxy)}
              <button
                on:click|stopPropagation|preventDefault={clearDueDate}
                class="removeDue"
              >
                ❌
              </button>
            {/if}
          </div>
          <div class="presets">
            <button
              on:click|stopPropagation|preventDefault={() => {
                dueDateProxy = formatDTL(addDays(new Date(), 1));
              }}
            >
              Tomorrow
            </button>
            <button
              on:click|stopPropagation|preventDefault={() => {
                dueDateProxy = formatDTL(addDays(new Date(), 7));
              }}
            >
              Next Week
            </button>
          </div>
        </div>
      {/if}
    </div>
    <TagSearch
      on:updatedSelection={onTagsSelection}
      selectedTags={tags}
      suggestTrending
    />
    <Controls background>
      <button type="submit">Save 💾</button>
    </Controls>
  </form>
</div>

<style>
  .form {
    display: flex;
    flex-direction: column;
  }
  .title {
    padding: var(--input-default-pad);
    font-size: var(--input-medium-font-size);
    border-radius: var(--input-border-radius);
  }
  .editor {
    padding: 2em;
  }
  .note textarea {
    font-size: var(--input-font-size);
    padding: var(--input-default-pad);
    border-radius: var(--input-border-radius);
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
    padding: 0.5rem;
    font-size: 1rem;
  }

  .dueDateWrapper {
    display: flex;
    flex-direction: column;
  }

  .dueDateWrapper > div {
    display: flex;
    flex-direction: row;
  }

  .dueDateWrapper h3 {
    margin: 0;
    font-size: 1rem;
  }

  .presets {
    display: flex;
    flex-direction: row;
  }

  .removeDue {
    font-size: 0.5rem;
    padding: 0.5rem;
  }

  .controls {
    display: flex;
    flex-direction: row;
  }
</style>
