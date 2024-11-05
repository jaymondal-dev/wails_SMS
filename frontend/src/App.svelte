<script>
    const SchoolName = "Name of the School";

    const atcmSubE = "Ability to communicate";
    const atcoSubE = "ability to correlate";
    const aipsSubE = "Ability in Problem Solving";
    const aimpcSubE = "Ability in physical and mental coordination";
    const atcmSubB = "সংযোগ স্থাপনে সক্ষমতা";
    const atcoSubB = "সমন্বয় সাধনে সক্ষমতা";
    const aipsSubB = "সমস্যা সমাধনে সক্ষমতা";
    const aimpcSubB = "মানসিক ও শারীরিক সমন্বয় সাধনে সক্ষমতা";

    const flSubE = "First Language";
    const slSubE = "Second Language";
    const mSubE = "Mathematics";
    const eSubE = "Our Environment";
    const hpethSubE = "Health and Physical Education Theory";
    const hpeprSubE = "Health and Physical Education Practical";
    const awethSubE = "Art and Work Education Theory";
    const aweprSubE = "Art and Work Education Practical";
    const flSubB = "প্রথম ভাষা";
    const slSubB = "দ্বিতীয় ভাষা";
    const mSubB = "গণিত";
    const eSubB = "আমাদের পরিবেশ";
    const hpethSubB = "স্বাস্থ্য ও শরীর শিক্ষা তাত্বিক";
    const hpeprSubB = "স্বাস্থ্য ও শরীর শিক্ষা ব্যবহারিক";
    const awethSubB = "কলা ও কর্ম শিক্ষা তাত্বিক";
    const aweprSubB = "কলা ও কর্ম শিক্ষা ব্যবহারিক";

    // Constants for Class 1 & 2
    const Term1atcm = 10;
    const Term1atco = 10;
    const Term1aips = 10;
    const Term1aimpc = 10;
    const Term2atcm = 10;
    const Term2atco = 10;
    const Term2aips = 10;
    const Term2aimpc = 10;
    const Term3atcm = 30;
    const Term3atco = 30;
    const Term3aips = 30;
    const Term3aimpc = 30;

    // Constants for Class 3, 4, 5
    const Term1fl = 10;
    const Term1sl = 10;
    const Term1M = 10;
    const Term1E = 10;
    const Term1hpeth = 0;
    const Term1hpepr = 5;
    const Term1aweth = 0;
    const Term1awepr = 5;
    const Term2fl = 20;
    const Term2sl = 20;
    const Term2M = 20;
    const Term2E = 20;
    const Term2hpeth = 0;
    const Term2hpepr = 10;
    const Term2aweth = 0;
    const Term2awepr = 10;
    const Term3fl = 70;
    const Term3sl = 70;
    const Term3M = 70;
    const Term3E = 70;
    const Term3hpeth = 15;
    const Term3hpepr = 10;
    const Term3aweth = 15;
    const Term3awepr = 10;
    import { onMount } from "svelte";
    import {
        ListClasses,
        CreateClass,
        DeleteClass,
        ListStudents12,
        ListStudents345,
        AddStudent,
        DeleteStudent,
        UpdateStudent12,
        UpdateStudent345,
        GenerateAndDownloadPDF,
    } from "../wailsjs/go/main/App";
    import PDF from "./PDF.svelte";
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
    let selectedStudent = null;
    let selectedTable = null;
    let editStudentName = "";
    let editStudentRoll = 0;
    let editStudentId = "";
    let editStudentMotherName = "";
    let editStudentFatherName = "";
    let editStudentGuardianName = "";

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
    }

    async function loadStudents(cls) {
        try {
            if (
                cls.name === "Class 3" ||
                cls.name === "Class 4" ||
                cls.name === "Class 5"
            ) {
                prestudents = await ListStudents345(cls.TableName);
            } else {
                prestudents = await ListStudents12(cls.TableName);
            }
            if (prestudents != null) {
                students = prestudents;
            } else {
                students = [];
            }
        } catch (error) {
            console.error("Error loading students:", error);
            students = [];
        }
    }

    async function handleSelectClass(cls) {
        selectedClass = cls;
        searchTerm = "";
        await loadStudents(cls);
    }

    let newStudentName = "";
    let newStudentRoll = 0;
    let newStudentId = "";
    let newStudentMotherName = "";
    let newStudentFatherName = "";
    let newStudentGuardianName = "";

    async function handleAddStudent() {
        if (selectedClass) {
            await AddStudent(
                selectedClass.TableName,
                newStudentName,
                newStudentRoll,
                newStudentId,
                newStudentMotherName,
                newStudentFatherName,
                newStudentGuardianName,
            );
            newStudentName = "";
            newStudentRoll = 0;
            newStudentId = "";
            newStudentMotherName = "";
            newStudentFatherName = "";
            newStudentGuardianName = "";
            await loadStudents(selectedClass);
            toggleCreateStudent();
        }
    }

    let viewStudent = false;
    async function toogleViewStudent() {
        viewStudent = !viewStudent;
    }

    async function handleSelectStudent(stu, tname) {
        if (viewStudent == false) {
            toogleViewStudent();
        }
        selectedStudent = stu;
        selectedTable = tname;
        if (editMode) {
            editStudentName = selectedStudent.name;
            editStudentRoll = selectedStudent.roll;
            editStudentId = selectedStudent.studentId;
            editStudentMotherName = selectedStudent.motherName;
            editStudentFatherName = selectedStudent.fatherName;
            editStudentGuardianName = selectedStudent.guardianName;
        }
    }

    let confirmDeleteStudent = false;
    let editMode = false;

    function toggleConfirmDeleteStudent() {
        confirmDeleteStudent = !confirmDeleteStudent;
    }

    function toggleEditMode() {
        editMode = !editMode;
        if (editMode) {
            editStudentName = selectedStudent.name;
            editStudentRoll = selectedStudent.roll;
            editStudentId = selectedStudent.studentId;
            editStudentMotherName = selectedStudent.motherName;
            editStudentFatherName = selectedStudent.fatherName;
            editStudentGuardianName = selectedStudent.guardianName;
        }
    }

    async function handleEditStudent() {
        if (selectedStudent && selectedTable) {
            if (
                selectedClass.name === "Class 1" ||
                selectedClass.name === "Class 2" ||
                selectedClass.name === "PP"
            ) {
                await UpdateStudent12(
                    selectedTable,
                    selectedStudent.id,
                    editStudentName,
                    editStudentRoll,
                    editStudentId,
                    editStudentMotherName,
                    editStudentFatherName,
                    editStudentGuardianName,
                    selectedStudent.term1atcm,
                    selectedStudent.term1atco,
                    selectedStudent.term1aips,
                    selectedStudent.term1aimpc,
                    selectedStudent.term2atcm,
                    selectedStudent.term2atco,
                    selectedStudent.term2aips,
                    selectedStudent.term2aimpc,
                    selectedStudent.term3atcm,
                    selectedStudent.term3atco,
                    selectedStudent.term3aips,
                    selectedStudent.term3aimpc,
                );
            } else {
                await UpdateStudent345(
                    selectedTable,
                    selectedStudent.id,
                    editStudentName,
                    editStudentRoll,
                    editStudentId,
                    editStudentMotherName,
                    editStudentFatherName,
                    editStudentGuardianName,
                    selectedStudent.term1fl,
                    selectedStudent.term1sl,
                    selectedStudent.term1m,
                    selectedStudent.term1e,
                    selectedStudent.term1hpeth,
                    selectedStudent.term1hpepr,
                    selectedStudent.term1aweth,
                    selectedStudent.term1awepr,
                    selectedStudent.term2fl,
                    selectedStudent.term2sl,
                    selectedStudent.term2m,
                    selectedStudent.term2e,
                    selectedStudent.term2hpeth,
                    selectedStudent.term2hpepr,
                    selectedStudent.term2aweth,
                    selectedStudent.term2awepr,
                    selectedStudent.term3fl,
                    selectedStudent.term3sl,
                    selectedStudent.term3m,
                    selectedStudent.term3e,
                    selectedStudent.term3hpeth,
                    selectedStudent.term3hpepr,
                    selectedStudent.term3aweth,
                    selectedStudent.term3awepr,
                );
            }
            selectedStudent.name = editStudentName;
            selectedStudent.roll = editStudentRoll;
            selectedStudent.studentId = editStudentId;
            selectedStudent.motherName = editStudentMotherName;
            selectedStudent.fatherName = editStudentFatherName;
            selectedStudent.guardianName = editStudentGuardianName;
            toggleEditMode();
            await loadStudents(selectedClass);
        }
    }

    async function handleConfirmDeleteStudent() {
        if (selectedStudent) {
            await DeleteStudent(selectedTable, selectedStudent.id);
            loadStudents(selectedClass);
        }
        viewStudent = false;
        confirmDeleteStudent = false;
    }

    let searchTerm = "";
    let filteredStudents = [];

    $: {
        if (students.length > 0 && searchTerm) {
            filteredStudents = students.filter((student) =>
                student.name.toLowerCase().includes(searchTerm.toLowerCase()),
            );
        } else {
            filteredStudents = students;
        }
    }

    function handleSearch(event) {
        searchTerm = event.target.value;
    }

    let showQuickDeleteConfirm = false;
    let studentToDelete = null;

    function confirmQuickDelete(student) {
        studentToDelete = student;
        showQuickDeleteConfirm = true;
    }

    function cancelQuickDelete() {
        studentToDelete = null;
        showQuickDeleteConfirm = false;
    }

    async function handleQuickDelete() {
        if (studentToDelete && selectedClass) {
            await DeleteStudent(selectedClass.TableName, studentToDelete.id);
            await loadStudents(selectedClass);
            showQuickDeleteConfirm = false;
            studentToDelete = null;
        }
    }
    function validateNumberInput(value, max) {
        const num = parseInt(value);
        if (num > max) return max;
        if (num < 0) return 0;
        return num;
    }

    function handleMarkInput(event, field, maxValue) {
        const inputEl = event.target;
        const validatedValue = validateNumberInput(inputEl.value, maxValue);
        inputEl.value = validatedValue.toString();
        selectedStudent[field] = validatedValue;
    }

    function downloadPDF(c_name, s_name, c_id, s_id) {
        // Call the backend function
        GenerateAndDownloadPDF(c_name, s_name, c_id, s_id)
            .then((result) => {
                // Convert base64 back to bytes
                const binaryString = window.atob(result.data);
                const bytes = new Uint8Array(binaryString.length);
                for (let i = 0; i < binaryString.length; i++) {
                    bytes[i] = binaryString.charCodeAt(i);
                }

                // Create blob and download
                const blob = new Blob([bytes], { type: "application/pdf" });
                const url = window.URL.createObjectURL(blob);
                const link = document.createElement("a");
                link.href = url;
                link.setAttribute("download", result.filename);
                document.body.appendChild(link);
                link.click();
                document.body.removeChild(link);
                window.URL.revokeObjectURL(url);
            })
            .catch((error) => {
                console.error("Error generating PDF:", error);
            });
        alert("Pdf exported successfully.");
    }
    // let urlOpen = "";
    // async function downloadPDF() {
    //     try {
    //         // Call backend function to get the URL
    //         const url = await GeneratePDFURL();
    //         if (url!=null){
    //           urlOpen=url
    //         }
    //         // Redirect the user to the URL for download
    //     } catch (error) {
    //         console.error("Error generating PDF:", error);
    //     }
    // }
</script>

<header class="flex flex-row justify-between items-center m-4">
    <h1 class="text-3xl font-bold">{SchoolName}</h1>

    <button
        on:click={toggleCreateClass}
        class="px-4 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
        >Create Class</button
    >
</header>
<main class="m-4 h-[85vh] flex flex-row items-start justify-start gap-2">
    <aside class="w-48 bg-[#1E1E1E] p-2 h-[85vh] rounded-md overflow-auto">
        <h2 class="text-2xl font-bold sticky">Class Lists</h2>
        {#if classes.length > 0}
            <ul class="flex flex-col gap-2 mt-6 overflow-auto">
                {#each classes as cls}
                    <li
                        class="flex flex-row justify-between items-center text-sm"
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
            <p class="mt-4 text-gray-400">No classes found</p>
        {/if}
    </aside>

    {#if selectedClass}
        <div class="relative bg-[#1E1E1E] h-full w-full p-4 rounded-md">
            <div
                class="absolute top-4 right-2 flex-end flex flex-row gap-2 items-center justify-center"
            >
                <button
                    class=" px-2 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
                    on:click={toggleSettingsClass}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="20"
                        height="20"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="lucide lucide-bolt"
                        ><path
                            d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"
                        /><circle cx="12" cy="12" r="4" /></svg
                    ></button
                >
                <button
                    class=" px-2 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
                    on:click={toggleCreateStudent}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="20"
                        height="20"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="lucide lucide-user-round-plus"
                        ><path d="M2 21a8 8 0 0 1 13.292-6" /><circle
                            cx="10"
                            cy="8"
                            r="5"
                        /><path d="M19 16v6" /><path d="M22 19h-6" /></svg
                    ></button
                >
                <button
                    class=" px-2 py-2 bg-[#151515] text-white font-semibold rounded-md hover:bg-[#282828] transition-colors"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="20"
                        height="20"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="lucide lucide-download"
                        ><path
                            d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"
                        /><polyline points="7 10 12 15 17 10" /><line
                            x1="12"
                            x2="12"
                            y1="15"
                            y2="3"
                        /></svg
                    >
                </button>
                <input
                    type="text"
                    placeholder="Search students..."
                    class="w-64 p-2 rounded bg-[#282828] text-white h-8"
                    on:input={handleSearch}
                    bind:value={searchTerm}
                />
            </div>

            <div class="mb-4">
                <h2 class="text-2xl font-bold">
                    Class Name : {selectedClass.name}
                    <span class="text-neutral-600"
                        >Year: {selectedClass.year}</span
                    >
                </h2>
                <p>Table name: {selectedClass.TableName}</p>
                <p>Total number of students: {students.length}</p>
            </div>
            <div class="h-[85%] overflow-y-auto">
                {#if filteredStudents.length > 0}
                    <ul class="w-full">
                        <li
                            class="flex justify-between items-center bg-[#282828] sticky top-0 px-2 py-1 text-sm"
                        >
                            <span class="w-12">Sl no.</span>
                            <span class="flex-grow">Name</span>
                            <span>Actions</span>
                        </li>
                        {#each filteredStudents as student, index}
                            <li
                                class="flex justify-between items-center border-b border-gray-600 px-2 py-1 text-sm"
                            >
                                <span class="w-12">{index + 1}</span>
                                <span class="flex-grow">{student.name}</span>
                                <span>
                                    <button
                                        on:click={() =>
                                            handleSelectStudent(
                                                student,
                                                selectedClass.TableName,
                                            )}
                                        class="p-1 text-white rounded-md mr-1"
                                    >
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="20"
                                            height="20"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="1"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                        >
                                            <circle cx="12" cy="12" r="10" />
                                            <line
                                                x1="12"
                                                y1="16"
                                                x2="12"
                                                y2="12"
                                            />
                                            <line
                                                x1="12"
                                                y1="8"
                                                x2="12"
                                                y2="8"
                                            />
                                        </svg>
                                    </button>
                                    <button
                                        class="p-1 text-white rounded-md mr-1"
                                        on:click={() =>
                                            downloadPDF(
                                                selectedClass.name,
                                                student.name,
                                                selectedClass.id,
                                                student.id,
                                            )}
                                    >
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="20"
                                            height="20"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="1"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                        >
                                            <path
                                                d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4"
                                            />
                                            <polyline
                                                points="7 10 12 15 17 10"
                                            />
                                            <line
                                                x1="12"
                                                y1="15"
                                                x2="12"
                                                y2="3"
                                            />
                                        </svg>
                                    </button>
                                    <button
                                        on:click={() =>
                                            confirmQuickDelete(student)}
                                        class="p-1 text-white rounded-md"
                                    >
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="20"
                                            height="20"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="1"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                        >
                                            <polyline points="3 6 5 6 21 6" />
                                            <path
                                                d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"
                                            />
                                            <line
                                                x1="10"
                                                y1="11"
                                                x2="10"
                                                y2="17"
                                            />
                                            <line
                                                x1="14"
                                                y1="11"
                                                x2="14"
                                                y2="17"
                                            />
                                        </svg>
                                    </button>
                                </span>
                            </li>
                        {/each}
                    </ul>
                {:else}
                    <div class="text-center text-gray-400 mt-4">
                        {students.length > 0
                            ? "No matching students found"
                            : "No students found"}
                    </div>
                {/if}
            </div>
        </div>
    {:else}
        <div class="h-full w-full flex justify-center items-center">
            <h2 class="text-4xl font-semibold text-gray-400">
                Select a class to view
            </h2>
        </div>
    {/if}
</main>
<!-- Modal for Create Class -->
{#if showCreateClass}
    <div
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
        <div
            class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative"
        >
            <button
                class="absolute top-4 right-4 p-2"
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
                >
                    <path d="M18 6 6 18" />
                    <path d="m6 6 12 12" />
                </svg>
            </button>

            <h2 class="text-xl font-bold mb-4">Create New Class</h2>
            <select
                class="w-full p-2 rounded bg-[#282828] text-black mb-2"
                bind:value={newClassName}
            >
                <option value="">Select Class</option>
                <option value="PP">PP</option>

                <option value="Class 1">Class 1</option>
                <option value="Class 2">Class 2</option>
                <option value="Class 3">Class 3</option>
                <option value="Class 4">Class 4</option>
                <option value="Class 5">Class 5</option>
            </select>

            <select
                class="w-full p-2 rounded bg-[#282828] text-black mb-2"
                bind:value={newClassYear}
            >
                <option value="">Select Year</option>
                {#each Array.from({ length: 11 }, (_, i) => 2020 + i) as year}
                    <option value={year}>{year}</option>
                {/each}
            </select>

            <button
                class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                on:click={handleCreateClass}
            >
                Create Class
            </button>
        </div>
    </div>
{/if}

<!-- Modal for Class Settings -->
{#if showSettingClass}
    <div
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
        <div
            class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative"
        >
            <button
                class="absolute top-4 right-4 p-2"
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
                >
                    <path d="M18 6 6 18" />
                    <path d="m6 6 12 12" />
                </svg>
            </button>

            <h2 class="text-xl font-bold mb-4">Class Settings</h2>

            <button
                class="w-full px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
                on:click={toggleconfirmDeleteClass}
            >
                Delete Class
            </button>

            {#if confirmDeleteClass}
                <div class="mt-4 p-4 bg-red-900 rounded-md">
                    <p class="text-white mb-2">
                        Are you sure you want to delete this class? This action
                        cannot be undone.
                    </p>
                    <button
                        class="w-full px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
                        on:click={() => handleDeleteClass(selectedClass.id)}
                    >
                        Confirm Delete
                    </button>
                </div>
            {/if}
        </div>
    </div>
{/if}

<!-- Modal for Create Student -->
{#if showcreateStudent}
    <div
        class="fixed inset-y-0 right-0 bg-black bg-opacity-50 flex items-start justify-end z-50 h-full w-[50%]"
    >
        <div
            class="bg-[#1E1E1E] border-l-2 border-l-gray-700 w-full h-full p-4 flex flex-col gap-2 relative"
        >
            <button
                class="absolute top-4 right-4 p-2"
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
                >
                    <path d="M18 6 6 18" />
                    <path d="m6 6 12 12" />
                </svg>
            </button>

            <h2 class="text-xl font-bold mb-4">Add New Student</h2>

            <input
                bind:value={newStudentName}
                placeholder="Student Name"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />
            <input
                bind:value={newStudentRoll}
                placeholder="Roll Number"
                type="number"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />
            <input
                bind:value={newStudentId}
                placeholder="Student ID"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />
            <input
                bind:value={newStudentMotherName}
                placeholder="Mother's Name"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />
            <input
                bind:value={newStudentFatherName}
                placeholder="Father's Name"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />
            <input
                bind:value={newStudentGuardianName}
                placeholder="Guardian's Name"
                class="w-full p-2 rounded bg-[#282828] text-white mb-2"
            />

            <button
                class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                on:click={handleAddStudent}
            >
                Add Student
            </button>
        </div>
    </div>
{/if}
{#if viewStudent}
    <div
        class="fixed inset-y-0 right-0 bg-black bg-opacity-50 flex items-start justify-end z-50 h-full w-[70%]"
    >
        <div
            class="bg-[#1E1E1E] border-l-2 border-l-gray-700 w-full h-full p-4 flex flex-col gap-2 relative overflow-y-auto"
        >
            <button
                class="absolute top-4 right-4 p-2"
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
                >
                    <path d="M18 6 6 18" />
                    <path d="m6 6 12 12" />
                </svg>
            </button>

            <h2 class="text-xl font-bold mb-4">
                {editMode ? "Edit Student" : "Student Details"}
            </h2>

            {#if editMode}
                <div class="space-y-2">
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Student Name</p>
                        <input
                            bind:value={editStudentName}
                            type="text"
                            placeholder="Student Name"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Roll Number</p>
                        <input
                            bind:value={editStudentRoll}
                            type="number"
                            placeholder="Roll Number"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Student ID</p>
                        <input
                            bind:value={editStudentId}
                            type="text"
                            placeholder="Student ID"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Mother's Name</p>
                        <input
                            bind:value={editStudentMotherName}
                            type="text"
                            placeholder="Mother's Name"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Father's Name</p>
                        <input
                            bind:value={editStudentFatherName}
                            type="text"
                            placeholder="Father's Name"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                    <div class="mb-2">
                        <p class="block text-sm mb-1">Guardian's Name</p>
                        <input
                            bind:value={editStudentGuardianName}
                            type="text"
                            placeholder="Guardian's Name"
                            class="w-full p-2 rounded bg-[#282828] text-white"
                        />
                    </div>
                </div>

                <!-- Class-specific mark inputs -->
                {#if selectedClass.name === "Class 1" || selectedClass.name === "Class 2"|| selectedClass.name === "PP"}
                    {#each ["Term 1", "Term 2", "Term 3"] as term}
                        <div class="mt-4">
                            <h3 class="font-bold mb-2">{term} Marks</h3>
                            <div class="grid grid-cols-2 gap-2">
                                <div>
                                    <p class="block text-sm mb-1">
                                        {atcmSubE}/{atcmSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 30 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}atcm`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {atcoSubE}/{atcoSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 30 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}atco`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {aipsSubE}/{aipsSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 30 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}aips`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {aimpcSubE}/{aimpcSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 30 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}aimpc`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                            </div>
                        </div>
                    {/each}
                {:else}
                    {#each ["Term 1", "Term 2", "Term 3"] as term}
                        <div class="mt-4">
                            <h3 class="font-bold mb-2">{term} Marks</h3>
                            <div class="grid grid-cols-2 gap-2">
                                <div>
                                    <p class="block text-sm mb-1">
                                        {flSubE}/{flSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 70 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}fl`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {slSubE}/{slSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 70 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}sl`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {mSubE}/{mSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 70 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}m`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {eSubE}/{eSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 70 : 10}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}e`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {hpethSubE}/{hpethSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 15 : 0}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}hpeth`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {hpeprSubE}/{hpeprSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3"
                                            ? 10
                                            : term === "Term 2"
                                              ? 10
                                              : 5}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}hpepr`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {awethSubE}/{awethSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3" ? 15 : 0}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}aweth`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                                <div>
                                    <p class="block text-sm mb-1">
                                        {aweprSubE}/{aweprSubB}
                                    </p>
                                    <input
                                        type="number"
                                        min="0"
                                        max={term === "Term 3"
                                            ? 10
                                            : term === "Term 2"
                                              ? 10
                                              : 5}
                                        bind:value={selectedStudent[
                                            `${term.toLowerCase().replace(" ", "")}awepr`
                                        ]}
                                        class="w-full p-2 rounded bg-[#282828] text-white"
                                    />
                                </div>
                            </div>
                        </div>
                    {/each}
                {/if}

                <div class="flex space-x-2 mt-4">
                    <button
                        on:click={handleEditStudent}
                        class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors"
                    >
                        Save Changes
                    </button>
                    <button
                        on:click={toggleEditMode}
                        class="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 transition-colors"
                    >
                        Cancel
                    </button>
                </div>
            {:else}
                <div class="space-y-4">
                    <!-- Basic Info Display -->
                    <p class="text-sm text-gray-400">
                        ID: {selectedStudent.id}
                    </p>
                    <p class="text-lg">Name: {selectedStudent.name}</p>
                    <p class="text-lg">Roll: {selectedStudent.roll}</p>
                    <p class="text-lg">
                        Student ID: {selectedStudent.studentId}
                    </p>
                    <p class="text-lg">
                        Mother's Name: {selectedStudent.motherName}
                    </p>
                    <p class="text-lg">
                        Father's Name: {selectedStudent.fatherName}
                    </p>
                    <p class="text-lg">
                        Guardian's Name: {selectedStudent.guardianName}
                    </p>

                    <!-- Marks Display -->
                    {#if selectedClass.name === "Class 1" || selectedClass.name === "Class 2"|| selectedClass.name === "PP"}
                        {#each ["Term 1", "Term 2", "Term 3"] as term}
                            <div class="mt-4">
                                <h3 class="font-bold mb-2">{term} Marks</h3>
                                <p>
                                    {atcmSubE}/{atcmSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}atcm`
                                    ]}
                                </p>
                                <p>
                                    {atcoSubE}/{atcoSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}atco`
                                    ]}
                                </p>
                                <p>
                                    {aipsSubE}/{aipsSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}aips`
                                    ]}
                                </p>
                                <p>
                                    {aimpcSubE}/{aimpcSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}aimpc`
                                    ]}
                                </p>
                            </div>
                        {/each}
                    {:else}
                        {#each ["Term 1", "Term 2", "Term 3"] as term}
                            <div class="mt-4">
                                <h3 class="font-bold mb-2">{term} Marks</h3>
                                <p>
                                    {flSubE}/{flSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}fl`
                                    ]}
                                </p>
                                <p>
                                    {slSubE}/{slSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}sl`
                                    ]}
                                </p>
                                <p>
                                    {mSubE}/{mSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}m`
                                    ]}
                                </p>
                                <p>
                                    {eSubE}/{eSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}e`
                                    ]}
                                </p>
                                <p>
                                    {hpethSubE}/{hpethSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}hpeth`
                                    ]}
                                </p>
                                <p>
                                    {hpeprSubE}/{hpeprSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}hpepr`
                                    ]}
                                </p>
                                <p>
                                    {awethSubE}/{awethSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}aweth`
                                    ]}
                                </p>
                                <p>
                                    {aweprSubE}/{aweprSubB}: {selectedStudent[
                                        `${term.toLowerCase().replace(" ", "")}awepr`
                                    ]}
                                </p>
                            </div>
                        {/each}
                    {/if}

                    <p class="text-sm text-gray-400">
                        Created: {new Date(
                            selectedStudent.createdAt,
                        ).toLocaleString()}
                    </p>

                    <!-- Action Buttons -->
                    <div class="flex space-x-2">
                        <button
                            on:click={toggleEditMode}
                            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                        >
                            Edit
                        </button>
                        <button
                            on:click={toggleConfirmDeleteStudent}
                            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
                        >
                            Delete
                        </button>
                        <button
                            class="px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
                            on:click={() =>
                                downloadPDF(
                                    selectedClass.name,
                                    selectedStudent.name,
                                    selectedClass.id,
                                    selectedStudent.id,
                                )}
                        >
                            Export to PDF
                        </button>
                    </div>

                    <!-- Delete Confirmation -->
                    {#if confirmDeleteStudent}
                        <div
                            class="mt-4 p-4 bg-red-900 border border-red-700 text-white rounded-md"
                        >
                            <p class="mb-2">
                                Are you sure you want to delete this student?
                                This action cannot be undone.
                            </p>
                            <div class="flex space-x-2">
                                <button
                                    on:click={handleConfirmDeleteStudent}
                                    class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
                                >
                                    Confirm Delete
                                </button>
                                <button
                                    on:click={toggleConfirmDeleteStudent}
                                    class="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 transition-colors"
                                >
                                    Cancel
                                </button>
                            </div>
                        </div>
                    {/if}
                </div>
            {/if}
        </div>
    </div>
{/if}
{#if showQuickDeleteConfirm}
    <div
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
        <div
            class="bg-[#1E1E1E] w-96 p-4 rounded-md flex flex-col gap-2 relative"
        >
            <h2 class="text-xl font-bold mb-2">Confirm Delete</h2>
            <p>Are you sure you want to delete {studentToDelete?.name}?</p>
            <div class="flex justify-end space-x-2 mt-4">
                <button
                    on:click={cancelQuickDelete}
                    class="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 transition-colors"
                >
                    Cancel
                </button>
                <button
                    on:click={handleQuickDelete}
                    class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
                >
                    Delete
                </button>
            </div>
        </div>
    </div>
{/if}
