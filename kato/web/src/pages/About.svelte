<div class="wrapper">
    <div>
        made with ♥ by <a href="https://vikkio.me" target="_blank">vikkio</a>
    </div>

    <div>
        Source available on <a
            href="https://github.com/vikkio88/kiffari/tree/main/kato"
            target="_blank"
            rel="noopener noreferrer">Github</a
        >
    </div>

    <h3>
        Version: <a
            href={`https://github.com/vikkio88/kiffari/commit/${
                import.meta.env.VITE_VERSION
            }`}
            target="_blank"
            rel="noopener noreferrer">{import.meta.env.VITE_VERSION}</a
        >
    </h3>
</div>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 80vh;
    }
</style>
