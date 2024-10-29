<script>
    export let classes = [];
    export let handleSelectClass;
</script>

<aside class="w-48 bg-[#1E1E1E] p-2 h-[85vh] rounded-md overflow-auto">
    <h2 class="text-2xl font-bold sticky">Class Lists</h2>
    {#if classes.length > 0}
        <ul class="flex flex-col gap-2 mt-6 overflow-auto">
            {#each classes as cls}
                <li class="flex flex-row justify-between items-center text-sm">
                    <span>{cls.name} - ({cls.year})</span>
                    <button on:click={() => handleSelectClass(cls)}>
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="lucide lucide-chevron-right"
                        >
                            <path d="m9 18 6-6-6-6" />
                        </svg>
                    </button>
                </li>
            {/each}
        </ul>
    {:else}
        <p class="mt-4 text-gray-400">No classes found</p>
    {/if}
</aside>