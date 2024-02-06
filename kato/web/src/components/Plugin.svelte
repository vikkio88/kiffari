<script>
    export let config = "";
    let plugins = parseConfig(config);

    /**
     * @param {String} config
     */
    function parseConfig(config) {
        const results = config.match(/<\!--(.+?)-->/s);
        if (results.length < 2) {
            return [];
        }
        console.log(results[1].trim());
        return [];
    }
</script>

<ul>
    {#each plugins as plugin}
        <li>{plugin.name}</li>
    {/each}
</ul>
