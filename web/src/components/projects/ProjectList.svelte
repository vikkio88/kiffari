<script>
  import { navigate } from "svelte-routing";

  export let projects = [];
</script>

{#if projects.length < 1}
  <h3>No projects... 🤷</h3>
{:else}
  <div class="projects">
    {#each projects as project}
      <div class="project">
        <h2>{project.name}</h2>
        <h3>{project.description}</h3>
        <button on:click={() => navigate(`/projects/${project.id}`)}>➡️</button>
      </div>
    {/each}
  </div>
{/if}

<style>
  .projects {
    margin: 0 auto;
    display: grid;
    grid-template-columns: auto auto;
    gap: 1rem;
    padding: 0 1rem;
  }

  .project {
    border: var(--default-borders);
    border-radius: var(--default-borders-radius);
    padding: 1rem 2rem;
  }
</style>
