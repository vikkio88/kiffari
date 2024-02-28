<script>
  import ProjectList from "../components/ProjectList.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { getProjects } from "../libs/api";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();
  let projectsPromise = getProjects();
</script>

{#await projectsPromise}
  <Spinner />
{:then projects}
  <h2>Projects</h2>
  <ProjectList {projects} />
{/await}
