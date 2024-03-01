<script>
  import TaskEditor from "../components/task/TaskEditor.svelte";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();

  export let projectId = "";
</script>

<h2>Create Task</h2>
<TaskEditor {projectId} onSave={(task) => console.log(task)} />
