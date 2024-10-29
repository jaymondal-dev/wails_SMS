<script>
    import { onMount } from "svelte";
    import {
        ListClasses,
        CreateClass,
        DeleteClass,
        ListStudents,
        AddStudent,
        DeleteStudent,
    } from "../wailsjs/go/main/App";
    let showCreateClass = false;
    function toggleCreateClass() {
        showCreateClass = !showCreateClass;
    }
    let showSettingClass = false;
    function toggleSettingsClass() {
        showSettingClass = !showSettingClass;
    }
    let showcreateStudent = false;
    function toggleCreateStudent() {
        showcreateStudent = !showcreateStudent;
    }
    let confirmDeleteClass = false;
    function toggleconfirmDeleteClass() {
        confirmDeleteClass = !confirmDeleteClass;
    }
    let classes = [];
    let classespre = null;
    let selectedClass = null;
    let prestudents = null;
    let students = [];
    let newClassName = "";
    let newClassYear = new Date().getFullYear();
    let newStudentName = "";
    let selectedStudent = null;
    let selectedTable = null;
    let editStudentName = "";

    onMount(async () => {
        await loadClasses();
    });

    async function loadClasses() {
        classespre = await ListClasses();
        if (classespre == null) {
            classes = [];
        } else {
            classes = classespre;
        }
    }

    async function handleCreateClass() {
        await CreateClass(newClassName, newClassYear);
        newClassName = "";
        newClassYear = new Date().getFullYear();
        await loadClasses();
        toggleCreateClass();
    }

    async function handleDeleteClass(id) {
        await DeleteClass(id);
        selectedClass = null;
        toggleconfirmDeleteClass();
        toggleSettingsClass();
        await loadClasses();

        // if (selectedClass && selectedClass.id === id) {
        //   selectedClass = null;
        //   students = [];
        // }
    }
    async function loadStudents(cls) {
        prestudents = await ListStudents(cls.TableName);
        if (prestudents != null) {
            students = prestudents;
        } else {
            students = [];
        }
    }

    async function handleSelectClass(cls) {
        selectedClass = cls;
        // students = await ListStudents(cls.TableName);
        await loadStudents(cls);
    }

    async function handleAddStudent() {
        if (selectedClass) {
            await AddStudent(selectedClass.TableName, newStudentName);
            newStudentName = "";
            prestudents = await ListStudents(selectedClass.TableName);
            if (prestudents != null) {
                students = prestudents;
            } else {
                students = [];
            }
        }
    }
    let viewStudent = false;
    async function toogleViewStudent() {
        viewStudent = !viewStudent;
    }
    async function handleSelectStudent(stu, tname) {
        // viewStudent=!viewStudent
        if (viewStudent == false) {
            toogleViewStudent();
        }
        selectedStudent = stu;
        selectedTable = tname;
        // await loadSingleStudent(stu)
    }
    async function handleDeleteStudent(id) {
        if (selectedStudent) {
            await DeleteStudent(selectedTable, id);
            loadStudents(selectedClass);
        }
        viewStudent = false;
    }
</script>

<header class="flex flex-row justify-between items-center m-4">
    <h1 class="text-3xl font-bold">School Management System</h1>

    <button
        on:click={toggleCreateClass}
        class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
        >create class</button
    >
</header>
<main class="m-4 h-[85vh] flex flex-row items-start justify-start gap-2">
    <aside class="w-48 bg-[#1E1E1E] p-2 h-[85vh] rounded-md">
        <h2 class="text-2xl font-bold">Class Lists</h2>
        {#if classes.length > 0}
            <ul class="flex flex-col gap-2 mt-4">
                {#each classes as cls}
                    <li
                        class="flex flex-row justify-between items-start text-sm"
                    >
                        <span>{cls.name} - ({cls.year})</span>
                        <button on:click={() => handleSelectClass(cls)}
                            ><svg
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
                                ><path d="m9 18 6-6-6-6" /></svg
                            ></button
                        >
                    </li>
                {/each}
            </ul>
        {:else}
            no classes
        {/if}
    </aside>

    {#if selectedClass}
        <div class="relative bg-[#1E1E1E] h-full w-full p-4 rounded-md">
            <button
                class="absolute z-30 top-2 right-2 px-2 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
                on:click={toggleSettingsClass}
            >
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
                    class="lucide lucide-settings"
                    ><path
                        d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"
                    /><circle cx="12" cy="12" r="3" /></svg
                ></button
            >
            <button
                class="absolute z-30 top-2 right-14 px-2 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
                on:click={toggleCreateStudent}
            >
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
                    class="lucide lucide-list-plus"
                    ><path d="M11 12H3" /><path d="M16 6H3" /><path
                        d="M16 18H3"
                    /><path d="M18 9v6" /><path d="M21 12h-6" /></svg
                ></button
            >
            <div class="mb-4">
                <h2 class="text-2xl font-bold">
                    Class Name : {selectedClass.name}
                </h2>
                <h2 class="text-2xl">Year: {selectedClass.year}</h2>
                <p>table name: {selectedClass.TableName}</p>
                <p>total number of students:{students.length}</p>
            </div>
            <div class="h-[70%] overflow-y-auto">
                {#if students.length > 0}
                    <ul class="w-full">
                        <li
                            class="flex justify-between items-center bg-[#282828] sticky top-0 px-2 py-1 text-sm"
                        >
                            <span class="w-12">Sl no.</span>
                            <!-- <span class="w-8">ID</span> -->

                            <span class="flex-grow">Name</span>
                            <span>Actions</span>
                        </li>
                        {#each students as student}
                            <li
                                class="flex justify-between items-center border-b border-gray-600 px-2 py-1 text-sm"
                            >
                                <span class="w-12"
                                    >{students.indexOf(student) + 1}</span
                                >
                                <!-- <span class="w-8">{student.id}</span> -->

                                <span class="flex-grow">{student.name}</span>
                                <span>
                                    <button
                                        on:click={() =>
                                            handleSelectStudent(
                                                student,
                                                selectedClass.TableName,
                                            )}
                                        class="px-1 py-0.5 bg-blue-500 text-white rounded mr-1 text-xs"
                                        >View</button
                                    >
                                    <button
                                        class="px-1 py-0.5 bg-red-500 text-white rounded text-xs"
                                        >export</button
                                    >
                                </span>
                            </li>
                        {/each}
                    </ul>
                {:else}
                    <div>no student found</div>
                {/if}
            </div>
        </div>
    {:else}
        <div class="h-full w-full flex justify-center items-center">
            <h2 class="text-4xl font-semibold">select a class to view</h2>
        </div>
    {/if}
</main>
{#if showCreateClass == true}
    <div
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
        <div
            class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative"
        >
            <button
                class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
                on:click={toggleCreateClass}
            >
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
                    class="lucide lucide-x"
                    ><path d="M18 6 6 18" /><path d="m6 6 12 12" /></svg
                ></button
            >

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
                    class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
                    on:click={handleCreateClass}>Add Class</button
                >
            </div>
        </div>
    </div>
{/if}
{#if showSettingClass == true}
    <div
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
        <div
            class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative"
        >
            <button
                class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
                on:click={toggleSettingsClass}
            >
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
                    class="lucide lucide-x"
                    ><path d="M18 6 6 18" /><path d="m6 6 12 12" /></svg
                ></button
            >

            <h1 class="text-2xl font-bold">Settings</h1>

            <div class="flex flex-col gap-2 mt-4">
                <button
                    class=" px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
                    on:click={toggleconfirmDeleteClass}>Delete Class</button
                >
                {#if confirmDeleteClass == true}
                    <p class="text-red-500">
                        this will delete the selected class, action can't be
                        reversed.
                    </p>
                    <button
                        class="px-4 py-2 bg-red-500 text-white font-semibold rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                        on:click={() => handleDeleteClass(selectedClass.id)}
                    >
                        Confirm Delete
                    </button>
                {/if}
            </div>
        </div>
    </div>
{/if}
{#if showcreateStudent == true}
    <div
        class="fixed inset-y-0 right-0 bg-black bg-opacity-50 flex items-start justify-end z-50 h-full w-[40%] border-l-2 border-l-gray-600"
    >
        <div
            class="bg-[#1E1E1E] w-full h-full p-4 flex flex-col gap-2 relative"
        >
            <button
                class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
                on:click={toggleCreateStudent}
            >
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
                    class="lucide lucide-x"
                    ><path d="M18 6 6 18" /><path d="m6 6 12 12" /></svg
                ></button
            >

            <h2 class="text-xl font-bold mb-2">Create New Student</h2>
            <div>
                <input
                    bind:value={newStudentName}
                    placeholder="Student Name"
                    class="p-2 text-black w-full"
                />
            </div>
            <div class="flex justify-between mt-4">
                <button
                    class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:border-blue-500 hover:border-2 focus:border-blue-500 focus:border-2 border-2 border-[#282828]"
                    on:click={handleAddStudent}>Add student</button
                >
            </div>
        </div>
    </div>
{/if}
{#if viewStudent}
    <div
        class="fixed inset-y-0 right-0 bg-black bg-opacity-50 flex items-start justify-end z-50 h-full w-[40%] border-l-2 border-l-gray-600"
    >
        <div
            class="bg-[#1E1E1E] w-full h-full p-4 flex flex-col gap-2 relative"
        >
            <button
                class="px-4 py-4 text-white rounded-md h-8 w-8 absolute top-0 right-4"
                on:click={toogleViewStudent}
            >
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
                    class="lucide lucide-x"
                    ><path d="M18 6 6 18" /><path d="m6 6 12 12" /></svg
                ></button
            >

            <h2 class="text-xl font-bold mb-2">Student details</h2>
            <div>
                <p>{selectedStudent.id}</p>
                <p>{selectedStudent.name}</p>
                <p>{selectedStudent.created_at}</p>
                <p>{selectedTable}</p>
                <button
                    on:click={() => {
                        handleDeleteStudent(selectedStudent.id);
                    }}
                    class="px-2 py-1 bg-red-500 text-white font-semibold rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                >
                    <p>Delete</p>
                </button>
                <button
                    class="px-2 py-1 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
                >
                    <p>Edit</p>
                </button>
                <button
                    class="px-2 py-1 bg-green-500 text-white font-semibold rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
                >
                    <p>Export to PDF</p>
                </button>
               
            </div>
        </div>
    </div>
{/if}

<style>
</style>
