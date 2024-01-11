<script>
    import Select from 'svelte-select';

    let filterText = '';

    let value = null;

    let items = [
      
    ];

    function handleFilter(e) {        
        if (value?.find(i => i.label.toLowerCase().includes(filterText.toLowerCase()))) return;
        if (e.detail.length === 0 && filterText.length > 0) {
            const prev = items.filter((i) => !i.created);
            items = [...prev, { value: filterText, label: filterText, created: true }];
        }
    }
    
    function handleChange(e) {
        items = items.map((i) => {
            delete i.created;
            return i;
        });
    }
</script>

<div class="tagSearch">
    <Select on:change={handleChange} multiple on:filter={handleFilter} bind:filterText bind:value {items}>
        <div slot="item" let:item>
            {item.created ? 'Add new: ' : ''}
            {item.label}
        </div>
    </Select>
</div>

<style>
    .tagSearch {
        width: 100%;
        color: black;
    }
</style>
