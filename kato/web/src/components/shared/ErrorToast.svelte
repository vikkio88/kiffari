<script>
    import { fade } from "svelte/transition";

    export let error = null;
    export let onDismiss = () => {
        error = null;
    };
</script>

{#if Boolean(error)}
    <div transition:fade class="error">
        {error}
        <button on:click={onDismiss}>❌</button>
    </div>
{/if}

<style>
    .error {
        background-color: #DB3A34;
        border-radius: 10px;
        position: absolute;
        bottom: 0;
        left: 45%;
        margin-bottom: 5rem;
        padding: 1rem 2rem;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
        font-weight: bold;
    }

    .error button {
        font-size: small;
    }
</style>
