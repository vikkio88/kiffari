<script>
  import { navigate } from "svelte-routing";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();
</script>

<div class="frc mh80">
  <div class="wrapper">
    <button on:click={() => navigate("/kato")}>Kato 🪣</button>
    <button on:click={() => navigate("/project")}>Kiffari 🎫</button>
  </div>
</div>

<style>
  button {
    width: 70%;
    font-size: 3rem;
    /* flex: 0 1 0; */
  }

  button + button {
    margin-top: 1rem;
  }
</style>
