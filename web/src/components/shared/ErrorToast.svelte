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
        background-color: var(--danger-color);
        border-radius: var(--input-border-radius);
        position: absolute;
        bottom: 0;
        left: 0;
        /* left: 50%; */
        /* right: 50%; */
        margin: 0 0 5rem 5rem;
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
