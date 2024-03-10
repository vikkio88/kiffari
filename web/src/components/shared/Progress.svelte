<script>
    export let total = 1;
    export let value = 0;

    let perc = (value / total) * 100;
</script>

<div class="wrapper">
    <div class="value" style={`width:${perc}%`}>
        {`${value}/${total}`}
    </div>
</div>

<style>
    .wrapper {
        text-align: unset;
        display: block;
        background-color: var(--x-light-gray);
    }

    .value {
        background-color: var(--success-color);
        padding: 0.2rem;
        font-weight: bold;
        color: var(--main-bg-color);
    }
</style>
