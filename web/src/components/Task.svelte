<script>
  import { createEventDispatcher } from "svelte";
  import { D_TASK_STATUS_LABELS, D_TASK_WORKFLOW as WF } from "../const";
  import { navigate } from "svelte-routing";
  const d = createEventDispatcher();
  export let task = null;

  function statusChange(newStatus) {
    d("updatedTask", {
      task: { ...task, status: newStatus },
      prevStatus: task.status,
    });
  }
</script>

<div class="wrapper">
  <div>
    <button title="Info" on:click={() => navigate(`/tasks/${task.id}`)}>
      ℹ️
    </button>
  </div>
  <div class="mainInfo">
    <h2>{task.title}</h2>
    <div class="info">
      {#if Array.isArray(task.links) && task.links.length > 0}
        <div class="crs-pointer" title="Has linked Tasks">🔗</div>
      {/if}
      {#if Boolean(task.flag)}
        <div class="crs-pointer" title={`FLAGGED: ${task.flag}`}>🚩</div>
      {/if}
      {#if Boolean(task.description)}
        <div class="crs-pointer" title="Description">📄</div>
      {/if}
    </div>
  </div>
  <div class="controls">
    {#if task.status && WF[task.status].to}
      <button
        on:click={() => statusChange(WF[task.status].to)}
        title={`Move to "${D_TASK_STATUS_LABELS[WF[task.status].to]}"`}
      >
        ⬆️
      </button>
    {/if}
    {#if task.status && WF[task.status].from}
      <button
        title={`Move to "${D_TASK_STATUS_LABELS[WF[task.status].from]}"`}
        on:click={() => statusChange(WF[task.status].from)}
      >
        ⬇️
      </button>
    {/if}
  </div>
</div>

<style>
  h2 {
    font-size: 1.3rem;
    margin: 0;
  }
  .wrapper {
    padding: 1rem 2rem;
    border: 2px white solid;
    border-radius: 10px;
    margin: 1rem 0 0 0;
    min-width: 80%;
    display: flex;
    align-items: center;
  }
  .mainInfo {
    flex: 1 1 auto;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }

  .info {
    margin-left: 1rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 0.4rem;
  }

  .controls {
    font-size: 0.8rem;
  }
</style>
