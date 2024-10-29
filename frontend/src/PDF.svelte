<!-- <script>
    import PdfViewer from "svelte-pdf";
    // import { GenerateAndDownloadPDF } from "wailsjs/go/main/App";
    export let binaryString;
    // function downloadPDF() {
    //     // Call the backend function
    //      GenerateAndDownloadPDF()
    //         .then((result) => {
    //             // Convert base64 back to bytes
    //             binaryString = window.atob(result.data);
    //             const bytes = new Uint8Array(binaryString.length);
    //             for (let i = 0; i < binaryString.length; i++) {
    //                 bytes[i] = binaryString.charCodeAt(i);
    //             }

    //             // Create blob and download
    //             const blob = new Blob([bytes], { type: 'application/pdf' });
    //             const url = window.URL.createObjectURL(blob);
    //             const link = document.createElement('a');
    //             link.href = url;
    //             link.setAttribute('download', result.filename);
    //             document.body.appendChild(link);
    //             link.click();
    //             document.body.removeChild(link);
    //             window.URL.revokeObjectURL(url);
    //         })
    //         .catch((error) => {
    //             console.error('Error generating PDF:', error);
    //         });
    // }
</script>

<PdfViewer data={atob(binaryString)} /> -->
