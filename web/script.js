document.addEventListener('DOMContentLoaded', function() {
    const fileInput = document.getElementById('svgFile');
    const fileName = document.getElementById('fileName');
    const dropArea = document.getElementById('dropArea');
    const uploadForm = document.getElementById('uploadForm');
    const convertBtn = document.getElementById('convertBtn');
    const loadingIndicator = document.getElementById('loadingIndicator');
    const errorMessage = document.getElementById('errorMessage');
  
// Dark mode toggle functionality
const themeToggle = document.getElementById('themeToggle');
const prefersDarkScheme = window.matchMedia('(prefers-color-scheme: dark)');

// Check for saved theme preference or use system preference
const currentTheme = localStorage.getItem('theme') || 
                    (prefersDarkScheme.matches ? 'dark' : 'light');

if (currentTheme === 'dark') {
  document.body.classList.add('dark-mode');
  themeToggle.textContent = '☀️'; // Sun icon for dark mode
} else {
  themeToggle.textContent = '🌙'; // Moon icon for light mode
}

themeToggle.addEventListener('click', () => {
  document.body.classList.toggle('dark-mode');
  const theme = document.body.classList.contains('dark-mode') ? 'dark' : 'light';
  localStorage.setItem('theme', theme);
  themeToggle.textContent = theme === 'dark' ? '☀️' : '🌙'; // Update emoji based on new theme
});
    
    // Handle file selection
    fileInput.addEventListener('change', function() {
      if (this.files.length > 0) {
        fileName.textContent = this.files[0].name;
      } else {
        fileName.textContent = 'No file selected';
      }
    });
    
    // Drag and drop functionality
    ['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
      dropArea.addEventListener(eventName, preventDefaults, false);
    });
    
    function preventDefaults(e) {
      e.preventDefault();
      e.stopPropagation();
    }
    
    ['dragenter', 'dragover'].forEach(eventName => {
      dropArea.addEventListener(eventName, highlight, false);
    });
    
    ['dragleave', 'drop'].forEach(eventName => {
      dropArea.addEventListener(eventName, unhighlight, false);
    });
    
    function highlight() {
      dropArea.classList.add('highlight');
    }
    
    function unhighlight() {
      dropArea.classList.remove('highlight');
    }
    
    dropArea.addEventListener('drop', handleDrop, false);
    
    function handleDrop(e) {
      const dt = e.dataTransfer;
      const files = dt.files;
      
      if (files.length > 0) {
        fileInput.files = files;
        fileName.textContent = files[0].name;
      }
    }
    
    // Form submission
    uploadForm.addEventListener('submit', async function (event) {
      event.preventDefault();
      
      const file = fileInput.files[0];
      
      if (!file) {
        showError("Please select an SVG file first.");
        convertBtn.classList.add('shake');
        setTimeout(() => {
          convertBtn.classList.remove('shake');
        }, 500);
        return;
      }
      
      // Check if file is SVG
      if (!file.name.toLowerCase().endsWith('.svg')) {
        showError("Please select a valid SVG file.");
        return;
      }
      
      // Show loading indicator
      loadingIndicator.style.display = 'block';
      convertBtn.disabled = true;
      errorMessage.style.display = 'none';
      
      const arrayBuffer = await file.arrayBuffer();
      
      try {
        const response = await fetch('http://localhost:8080/svg-to-pdf', {
          method: 'POST',
          headers: {
            'Content-Type': 'image/svg+xml',
          },
          body: arrayBuffer
        });
        
        if (!response.ok) {
          throw new Error("Failed to convert SVG. Server responded with " + response.status);
        }
        
        const blob = await response.blob();
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = file.name.replace('.svg', '') + '.pdf';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
      } catch (error) {
        showError("Error: " + error.message);
      } finally {
        loadingIndicator.style.display = 'none';
        convertBtn.disabled = false;
      }
    });
    
    function showError(message) {
      errorMessage.textContent = message;
      errorMessage.style.display = 'block';
    }
  });