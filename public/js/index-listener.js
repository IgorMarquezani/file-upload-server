const previewFile = () => {
  const fileInput = document.getElementById('fileInput');
  const file = fileInput.files[0];

  if (file) {
    const reader = new FileReader();
    reader.onload = function(e) {
      const preview = document.getElementById("preview");
      preview.src = e.target.result; // Set the preview image source
      preview.style.display = "block"; // Show the image
    };
    reader.readAsDataURL(file); // Convert to base64
  }
}
