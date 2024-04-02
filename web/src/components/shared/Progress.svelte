<script>
    export let total = 1;
    export let value = 0;
</script>

<meter min={0} max={total} {value} />

<style>
    meter {
        height: 1rem;
    }
</style>
