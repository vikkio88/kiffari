<script>
  import { Link } from "svelte-routing";
  import { getProjectWithArchivedTasks } from "../../libs/api";
  import { protectedRoute } from "../../libs/routes";
  import Spinner from "../../components/shared/Spinner.svelte";
  import Task from "../../components/task/Task.svelte";

  protectedRoute();
  export let projectId = "";

  let projectpromise = getProjectWithArchivedTasks(projectId);
</script>

{#await projectpromise}
  <Spinner />
{:then project}
  <Link to={`/projects/${project.id}`}><h2>{project.name}</h2></Link>
  <h3>Archived Tasks</h3>
  {#if !Array.isArray(project.tasks)}
    <span>No Tasks...</span>
  {:else}
    {#each project.tasks as task}
      <Task {task} />
    {/each}
  {/if}
{/await}
