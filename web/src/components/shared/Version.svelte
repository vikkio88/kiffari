<script>
  import { navigate } from "svelte-routing";
</script>

<button class="version" on:click={() => navigate("/about")}>
  version: {import.meta.env.VITE_VERSION}
</button>

<style>
  .version {
    color: var(--fade-gray);
  }

  .version {
    border: 0;
    background-color: transparent;
  }
</style>
