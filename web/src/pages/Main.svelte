<script>
  import { navigate } from "svelte-routing";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();
</script>

<div class="frc mh80">
  <div class="wrapper">
    <button on:click={() => navigate("/kato")}>Kato 🪣</button>
    <button on:click={() => navigate("/kiffari")}>Kiffari 🎫</button>
  </div>
</div>

<style>
  button {
    width: 70%;
    font-size: 3rem;
  }

  button + button {
    margin-top: 1rem;
  }
</style>
