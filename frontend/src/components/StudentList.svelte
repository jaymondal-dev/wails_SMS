<script>
    export let filteredStudents = [];
    export let handleSelectStudent;
    export let selectedClass;
    export let confirmQuickDelete;
</script>

<div class="h-[85%] overflow-y-auto">
    {#if filteredStudents.length > 0}
        <ul class="w-full">
            <li class="flex justify-between items-center bg-[#282828] sticky top-0 px-2 py-1 text-sm">
                <span class="w-12">Sl no.</span>
                <span class="flex-grow">Name</span>
                <span>Actions</span>
            </li>
            {#each filteredStudents as student, index}
                <li class="flex justify-between items-center border-b border-gray-600 px-2 py-1 text-sm">
                    <span class="w-12">{index + 1}</span>
                    <span class="flex-grow">{student.name}</span>
                    <span>
                        <button
                            on:click={() => handleSelectStudent(student, selectedClass.TableName)}
                            class="p-1 text-white rounded-md mr-1"
                        >
                            <svg><!-- Info icon SVG --></svg>
                        </button>
                        <button class="p-1 text-white rounded-md mr-1">
                            <svg><!-- Download icon SVG --></svg>
                        </button>
                        <button
                            on:click={() => confirmQuickDelete(student)}
                            class="p-1 text-white rounded-md"
                        >
                            <svg><!-- Delete icon SVG --></svg>
                        </button>
                    </span>
                </li>
            {/each}
        </ul>
    {:else}
        <div class="text-center text-gray-400 mt-4">
            No matching students found
        </div>
    {/if}
</div>