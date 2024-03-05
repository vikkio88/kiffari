<script>
  import TagsList from "../components/TagsList.svelte";
  import Markdown from "../components/renderers/Markdown.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import Header from "../components/shared/Header.svelte";
  import { navigate } from "svelte-routing";
  import { archive, getTask, unarchive, del, catchLogout } from "../libs/api";
  import { getDate } from "../libs/dates";
  import { protectedRoute } from "../libs/routes";
  import {
    D_TASK_STATUS_LABELS as STATUS_LABELS,
    D_TASK_CATEGORY_LABELS as CATEGORY_LABELS,
    D_TASK_STATUS_EMOJIS as STATUS_EMOJIS,
  } from "../const";
  import Controls from "../components/shared/Controls.svelte";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";
  import DashedHead from "../components/shared/DashedHead.svelte";
  protectedRoute();

  export let id = "";

  let taskPromise = getTask(id);

  async function onArchiveToggle(isArchived) {
    const resp = await (isArchived
      ? unarchive("tasks", id)
      : archive("tasks", id));
    if (resp.ok) {
      window.location.reload();
    }

    if (resp.status === 401) {
      catchLogout();
    }
  }
  async function onDelete(projectId) {
    const resp = await del("tasks", id);
    if (resp.status === 401) {
      catchLogout();
    }

    if (resp.ok) {
      navigate(`/projects/${projectId}`, { replace: true });
    }
  }
</script>

{#await taskPromise}
  <Spinner />
{:then task}
  {#if task.archived}
    <DashedHead>Archived</DashedHead>
  {/if}
  <div class="topBar">
    <button on:click={() => navigate(`/projects/${task.project_id}`)}>
      🔙
    </button>
    <div class="status">
      <span class="chip" title={STATUS_LABELS[task.status]}
        >{STATUS_EMOJIS[task.status]}</span
      >
      <span class="chip" title={task.category}
        >{CATEGORY_LABELS[task.category]}</span
      >
      {#if Boolean(task.flag)}
        <span class="chip">{task.flag} 🚩</span>
      {/if}
    </div>
  </div>
  <h2>{task.title}</h2>
  <div class="date">
    {getDate(task)}
  </div>
  <div class="description">
    <Markdown body={task.description} />
  </div>
  {#if Array.isArray(task.tags) && task.tags.length > 0}
    <div class="tags">
      <Header text="Tags" />
      <div class="tagList">
        <TagsList tags={task.tags} />
      </div>
    </div>
  {/if}
  <Controls background>
    {#if !task.archived}
      <button title="Edit" on:click={() => navigate(`/edit-task/${id}`)}>
        📝
      </button>
    {/if}
    <ConfirmButton
      title={task.archived ? "Un-Archive" : "Archive"}
      confirmLabel={task.archived ? "Un-Archive?" : "Archive?"}
      onConfirmed={() => onArchiveToggle(task.archived)}
    >
      {#if !task.archived}
        🗄️
      {:else}
        🔄
      {/if}
    </ConfirmButton>
    <ConfirmButton
      title="Delete"
      confirmLabel="Delete?"
      onConfirmed={() => onDelete(task.project_id)}
    >
      🗑️
    </ConfirmButton>
  </Controls>
{/await}

<style>
  .description {
    border-radius: 10px;
    text-align: left;
    padding: 1.5rem;
    margin: 0.5rem 0;
  }

  .date {
    font-size: smaller;
    color: #a3a3a3;
  }

  .topBar {
    padding: 1rem;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .chip {
    border: solid 2px #333;
    border-radius: 10px;
    padding: 0.5rem 2rem;
    margin: 0 0.5rem;
    cursor: help;
  }
</style>
