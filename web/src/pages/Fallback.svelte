<script>
    import { navigate } from "svelte-routing";
</script>

<div class="wrapper">
    <h1>Ooops!</h1>
    <h3>you found an empty Kato...</h3>
    <img src="/favicon.svg" alt="A Kato rolling around" />
    <button on:click={() => navigate("/")}>Back Home 🏠</button>
</div>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 80vh;
        gap: 5rem;
    }
    .wrapper img {
        width: 150px;
        animation: tip 1s ease-out infinite;
    }

    @keyframes tip {
        0% {
            transform: rotate(-10deg);
        }

        60% {
            transform: rotate(180deg);
        }

        70% {
            transform: translateY(-10%) rotate(180deg);
        }
        
        80% {
            transform: translateY(15%) rotate(180deg);
        }


        90% {
            transform: translateY(10%) rotate(180deg);
        }

        100% {
            transform: translateY(15%) rotate(180deg);
        }
    }
</style>
