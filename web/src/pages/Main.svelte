<script>
    import { navigate } from "svelte-routing";
    import { protectedRoute } from "../libs/routes";
    protectedRoute();
</script>

<div class="wrapper">
    <button on:click={() => navigate("/kato")}>Kato 🪣</button>
    <button on:click={() => navigate("/project")}>Kiffari 🎫</button>
</div>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 80vh;
    }
    button {
        padding: 3rem;
        font-size: 3rem;
    }

    button + button {
        margin-top: 1rem;
    }
</style>
