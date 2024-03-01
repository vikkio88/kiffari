<script>
  import SvelteMarkdown from "svelte-markdown";
  import { tick } from "svelte";
  import TagSearch from "../TagSearch.svelte";
  import Controls from "../shared/Controls.svelte";
  import Status from "./Status.svelte";
  import { D_TASK_STATUS } from "../../const";
  import Flag from "./Flag.svelte";

  export let projectId = null;

  export let title = "";
  export let text = "";
  export let flag = null;

  export let status = D_TASK_STATUS.BACKLOG;
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

  export let onSave = (projectId, task) => {
    console.log("default onSave", { projectId, task });
  };

  function onSaveInternal() {
    onSave(projectId, {
      title,
      description: text,
      status: status,
      tags,
      flag,
    });
  }

  let showAdditionalInfo = false;
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
        <Status bind:status />
        <Flag bind:flag />
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
    padding: 1em;
    font-size: 18px;
    border-radius: 10px;
  }
  .editor {
    padding: 2em;
  }
  .task textarea {
    font-size: 16px;
    padding: 1em;
    border-radius: 10px;
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
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-start;
  }

  .controls {
    display: flex;
    flex-direction: row;
  }
</style>
