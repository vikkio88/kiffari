<script>
  import { navigate } from "svelte-routing";
  import { addTask, catchLogout } from "../libs/api";
  import TaskEditor from "../components/task/TaskEditor.svelte";
  import ErrorToast from "../components/shared/ErrorToast.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();

  export let projectId = "";
  let creatingPromise = null;
  let error = null;

  async function createTask(projectId, task) {
    creatingPromise = addTask(projectId, { ...task });
    const resp = await creatingPromise;
    if (resp.ok) {
      const task = await resp.json();
      navigate(`/tasks/${task.id}`);
      return;
    }

    if (resp.status === 401) {
      catchLogout();
      return;
    }

    if (resp.status === 400) {
      error = "Could not update task...";
      creatingPromise = null;
      return;
    }
  }
</script>

<h2>Create Task</h2>
{#if Boolean(creatingPromise)}
  <Spinner />
{:else}
  <TaskEditor {projectId} onSave={createTask} />
{/if}

{#if Boolean(error)}
  <ErrorToast
    {error}
    onDismiss={() => {
      error = null;
      navigate(`/projects`);
    }}
  />
{/if}
