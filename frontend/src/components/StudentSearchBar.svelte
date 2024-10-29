<script>
    export let searchTerm = "";
    export let handleSearch;
</script>

<input
    type="text"
    placeholder="Search students..."
    class="w-64 p-2 rounded bg-[#282828] text-white h-8"
    on:input={handleSearch}
    bind:value={searchTerm}
/>