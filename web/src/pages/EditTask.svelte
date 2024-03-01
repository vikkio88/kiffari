<script>
  import Spinner from "../components/shared/Spinner.svelte";
  import TaskEditor from "../components/task/TaskEditor.svelte";
  import { getTask } from "../libs/api";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();

  export let id = "";
  let taskPromise = getTask(id);
</script>

<h2>Edit Task</h2>
{#await taskPromise}
  <Spinner />
{:then task}
  <TaskEditor
    projectId={task.projectId}
    title={task.title}
    text={task.description}
    tags={task.tags}
    flag={task.flag}
    status={task.status}
    onSave={(task) => console.log(task)}
  />
{/await}
