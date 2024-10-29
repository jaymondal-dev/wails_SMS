<script>
    export let toggleCreateClass;
    export let handleCreateClass;
    export let newClassName;
    export let newClassYear;
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative">
        <button
            class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
            on:click={toggleCreateClass}
        >
            <svg><!-- Close icon SVG --></svg>
        </button>

        <h2 class="text-xl font-bold mb-2">Create New Class</h2>
        <select class="p-2 rounded text-black" bind:value={newClassName}>
            <option value="">Select Class</option>
            {#each Array(5) as _, i}
                <option value="Class {i + 1}">Class {i + 1}</option>
            {/each}
        </select>

        <select class="p-2 rounded text-black" bind:value={newClassYear}>
            <option value="">Select Year</option>
            {#each Array.from({ length: 11 }, (_, i) => 2020 + i) as year}
                <option value={year}>{year}</option>
            {/each}
        </select>

        <div class="flex justify-between mt-4">
            <button
                class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
                on:click={handleCreateClass}
            >
                Add Class
            </button>
        </div>
    </div>
</div>