<script>
  import { navigate } from "svelte-routing";
  import ProjectList from "../components/projects/ProjectList.svelte";
  import Controls from "../components/shared/Controls.svelte";
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

<Controls>
  <button class="big-control" on:click={() => navigate("/create-project")}>
    Add Project 🎫
  </button>
</Controls>
