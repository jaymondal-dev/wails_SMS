<script>
    export let toggleSettingsClass;
    export let toggleconfirmDeleteClass;
    export let confirmDeleteClass;
    export let handleDeleteClass;
    export let selectedClass;
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative">
        <button
            class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
            on:click={toggleSettingsClass}
        >
            <svg><!-- Close icon SVG --></svg>
        </button>

        <h1 class="text-2xl font-bold">Settings</h1>

        <div class="flex flex-col gap-2 mt-4">
            <button
                class="px-4 py-2 bg-red-600 text-white font-semibold rounded-md hover:bg-red-700 transition-colors"
                on:click={toggleconfirmDeleteClass}
            >
                Delete Class
            </button>
            
            {#if confirmDeleteClass}
                <p class="text-red-500">
                    This will delete the selected class. Action cannot be reversed.
                </p>
                <button
                    class="px-4 py-2 bg-red-600 text-white font-semibold rounded-md hover:bg-red-700 transition-colors"
                    on:click={() => handleDeleteClass(selectedClass.id)}
                >
                    Confirm Delete
                </button>
            {/if}
        </div>
    </div>
</div>