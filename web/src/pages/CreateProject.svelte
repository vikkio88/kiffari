<script>
  import { navigate } from "svelte-routing";
  import ProjectEditor from "../components/ProjectEditor.svelte";
  import ErrorToast from "../components/shared/ErrorToast.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { catchLogout, createProject } from "../libs/api";

  let error = null;
  let postPromise = null;

  async function onSave(project) {
    const newProject = {
      name: project.name,
      description: project.description,
      links: project.links,
      config: project.config,
    };

    const resp = await createProject(newProject);
    if (resp.status == 401) {
      catchLogout();
      return;
    }

    if (resp.status == 400) {
      postPromise = null;
      error = `Could not create.`;
    }

    if (resp.status == 401) {
      catchLogout();
    }

    if (resp.status == 200) {
      const { id } = await resp.json();
      navigate(`/projects/${id}`, { replace: true });
    }
  }
</script>

<h2>Create New Project</h2>

{#if postPromise == null}
  <ProjectEditor {onSave} />
{:else}
  <h2>Creating...</h2>
  <Spinner />
{/if}

{#if Boolean(error)}
  <ErrorToast {error} onDismiss={() => (error = null)} />
{/if}
