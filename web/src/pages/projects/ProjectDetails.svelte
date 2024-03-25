<script>
  import Project from "../../components/projects/Project.svelte";
  import Spinner from "../../components/shared/Spinner.svelte";
  import { getProject } from "../../libs/api";
  import { protectedRoute } from "../../libs/routes";
  protectedRoute();

  export let id = "";

  let projectPromise = getProject(id);
</script>

{#await projectPromise}
  <Spinner />
{:then project}
  <Project {project} />
{/await}
