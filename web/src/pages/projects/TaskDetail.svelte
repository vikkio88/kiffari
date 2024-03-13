<script>
  import TagsList from "../../components/TagsList.svelte";
  import Spinner from "../../components/shared/Spinner.svelte";
  import Header from "../../components/shared/Header.svelte";
  import { Link, navigate } from "svelte-routing";
  import {
    archive,
    getTask,
    unarchive,
    del,
    catchLogout,
  } from "../../libs/api";
  import { getDate } from "../../libs/dates";
  import { protectedRoute } from "../../libs/routes";
  import {
    D_TASK_STATUS_LABELS as STATUS_LABELS,
    D_TASK_CATEGORY_LABELS as CATEGORY_LABELS,
    D_TASK_STATUS_EMOJIS as STATUS_EMOJIS,
  } from "../../const";
  import Controls from "../../components/shared/Controls.svelte";
  import ConfirmButton from "../../components/shared/ConfirmButton.svelte";
  import DashedHead from "../../components/shared/DashedHead.svelte";
  import Md from "../../components/shared/Md.svelte";
  import TaskMdRenderer from "../../components/TaskMdRenderer.svelte";
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
    <button on:click={() => history.back()}> 🔙 </button>
    <div class="status">
      <span class="chip" title={STATUS_LABELS[task.status]}>
        {STATUS_EMOJIS[task.status]}
      </span>
      <span class="chip" title={task.category.toUpperCase()}>
        {CATEGORY_LABELS[task.category]}
      </span>
      {#if Boolean(task.priority)}
        <span class="chip" title={`Priority: ${task.priority}`}>🔥</span>
      {/if}
      {#if Boolean(task.flag)}
        <span class="chip">{task.flag} 🚩</span>
      {/if}
    </div>
  </div>
  <h2>{task.title}</h2>
  <Link to={`/projects/${task.project_id}`}>
    <strong title="Project">({task.project.name})</strong>
  </Link>
  <div class="date">
    {getDate(task)}
  </div>
  <div class="description">
    <TaskMdRenderer {task} />
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
    border-radius: var(--default-border-radius);
    text-align: left;
    padding: 1.5rem;
    margin: 0.5rem 0;
  }

  .date {
    font-size: smaller;
    color: var(--x-light-gray);
  }

  .topBar {
    padding: var(--input-default-pad);
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .chip {
    border: solid 2px var(--light-gray);
    border-radius: var(--default-border-radius);
    padding: 0.5rem 2rem;
    margin: 0 0.5rem;
    cursor: help;
  }
</style>
