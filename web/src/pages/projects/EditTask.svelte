<script>
  import { navigate } from "svelte-routing";
  import Spinner from "../../components/shared/Spinner.svelte";
  import TaskEditor from "../../components/task/TaskEditor.svelte";
  import ErrorToast from "../../components/shared/ErrorToast.svelte";
  import { catchLogout, getTask, updateTask } from "../../libs/api";
  import { protectedRoute } from "../../libs/routes";
  protectedRoute();

  export let id = "";
  let taskPromise = getTask(id);
  let error = null;

  async function onUpdate(projectId, task) {
    taskPromise = updateTask(projectId, { id, ...task }).then((resp) => {
      if (resp.ok) {
        navigate(`/tasks/${id}`);
        return;
      }

      if (resp.status === 401) {
        catchLogout();
        return;
      }

      if (resp.status === 400) {
        error = "Could not update task...";
        taskPromise = null;
        return;
      }
    });
  }
</script>

<h2>Edit Task</h2>
{#if taskPromise != null}
  {#await taskPromise}
    <Spinner />
  {:then task}
    <TaskEditor
      projectId={task.project_id}
      text={task.description}
      title={task.title}
      tags={task.tags}
      flag={task.flag}
      status={task.status}
      priority={task.priority}
      category={task.category}
      onSave={onUpdate}
    />
  {/await}
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
