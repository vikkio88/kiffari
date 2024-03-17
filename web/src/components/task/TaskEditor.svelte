<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "../TagSearch.svelte";
  import Controls from "../shared/Controls.svelte";
  import Status from "./Status.svelte";
  import {
    D_TASK_STATUS as STATUS,
    D_TASK_CATEGORIES as CATEGORIES,
    D_TASK_CATEGORY_LABELS as CATEGORY_LABELS,
  } from "../../const";
  import Flag from "./Flag.svelte";
  import Select from "../shared/Select.svelte";
  import { generatePlugin } from "../../libs/renderers/helpers";
  import { removeComments } from "../../libs/renderers/cleanup";

  export let projectId = null;

  export let title = "";
  export let text = "";
  export let flag = null;

  export let priority = 0;
  export let status = STATUS.BACKLOG;
  export let tags = [];
  export let category = CATEGORIES.FEATURE;

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

  export let onSave = (projectId, task) => {
    console.log("default onSave", { projectId, task });
  };

  function onSaveInternal() {
    onSave(projectId, {
      title,
      description: text,
      status,
      category,
      priority,
      tags,
      flag,
    });
  }
</script>

<div class="editor">
  <div class="controls">
    <button on:click={() => (showPreview = !showPreview)}>
      {#if showPreview}
        Code 🤓
      {:else}
        Markdown 😎
      {/if}
    </button>
    <button
      on:click={() => {
        text = `${generatePlugin("Todo")}${removeComments(text)}`;
      }}
    >
      Plugin:Todo ✅
    </button>
    <button
      disabled
      title="Coming soon!"
      on:click={() => {
        console.log("Link other Task");
      }}
    >
      Link Task 🔗
    </button>
  </div>
  <div class="task">
    {#if !showPreview}
      <div class="form">
        <input
          required
          placeholder="New Task Title...."
          class="title"
          bind:value={title}
          on:focus={(e) => {
            e.currentTarget.select();
          }}
          type="text"
        />
        <textarea
          placeholder="Here goes the task description if needed..."
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
    <Select
      bind:selected={category}
      options={Object.values(CATEGORIES).map((c) => ({
        value: c,
        label: `${c.toUpperCase()} ${CATEGORY_LABELS[c]}`,
      }))}
    />
    <Status bind:status />
    <div class="frc g-1">
      <label for="priority">Priority</label>
      <input
        name="priority"
        size="2"
        type="number"
        min="0"
        step="1"
        bind:value={priority}
      />
    </div>
    <Flag bind:flag />
  </div>
  <TagSearch
    on:updatedSelection={onTagsSelection}
    selectedTags={tags}
    suggestTrending
  />
  <Controls background>
    <button on:click={onSaveInternal}>Save 💾</button>
  </Controls>
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
  .task textarea {
    font-size: var(--input-font-size);
    padding: var(--input-default-pad);
    border-radius: var(--input-border-radius);
  }

  .preview {
    display: block;
    text-align: left;
    min-height: 30vh;
  }

  .preview h2 {
    text-align: center;
  }

  .additionalInfo {
    margin-top: 1rem;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: flex-start;
  }

  .controls {
    display: flex;
    flex-direction: row;
    overflow: auto;
  }

  input[type="number"] {
    padding: var(--input-default-pad);
    font-size: var(--input-small-font-size);
    border-radius: var(--input-border-radius);
  }
</style>
