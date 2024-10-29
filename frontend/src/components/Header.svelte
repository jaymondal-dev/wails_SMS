<script>
    export let toggleCreateClass;
</script>

<header class="flex flex-row justify-between items-center m-4">
    <h1 class="text-3xl font-bold">School Management System</h1>
    <button
        on:click={toggleCreateClass}
        class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
    >
        Create Class
    </button>
</header>