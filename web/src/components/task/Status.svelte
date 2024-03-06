<script>
  import {
    D_TASK_STATUS as STATUS,
    D_TASK_STATUS_LABELS as LABEL,
    D_TASK_WORKFLOW as WF,
  } from "../../const";

  export let status = STATUS.BACKLOG;
</script>

<div>
  {#if WF[status].from}
    <button
      class="smaller"
      on:click|preventDefault|stopPropagation={() => (status = WF[status].from)}
    >
      ⬅️
    </button>
  {/if}
  {LABEL[status]}
  {#if WF[status].to}
    <button
      class="smaller"
      on:click|preventDefault|stopPropagation={() => (status = WF[status].to)}
    >
      ➡️
    </button>
  {/if}
</div>

<style>
  div {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    gap: 0.5rem;
    font-weight: bold;
  }
</style>
