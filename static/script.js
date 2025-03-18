document.addEventListener('DOMContentLoaded', function() {
    // Form validation enhancement
    const addForm = document.querySelector('form[action="/add"]');
    if (addForm) {
        initializeAddForm(addForm);
    }

    // Table sorting
    const tables = document.querySelectorAll('table');
    tables.forEach(initializeTableSort);

    // Search form enhancement
    const searchForm = document.querySelector('form[action="/search"]');
    if (searchForm) {
        initializeSearchForm(searchForm);
    }
});

function initializeAddForm(form) {
    // Tags input enhancement
    const tagsInput = form.querySelector('#tags');
    if (tagsInput) {
        // Convert comma-separated input into tag spans
        tagsInput.addEventListener('blur', function() {
            const tags = this.value.split(',').map(t => t.trim()).filter(t => t);
            this.value = tags.join(', ');
        });

        // Basic destination validation
        const destInput = form.querySelector('#destination');
        destInput.addEventListener('blur', function() {
            const val = this.value.trim();
            if (val.length < 516) {
                this.setCustomValidity('Destination appears too short for a valid I2P address');
            } else {
                this.setCustomValidity('');
            }
        });
    }
}

function initializeSearchForm(form) {
    // Add "clear search" functionality
    const queryInput = form.querySelector('#q');
    if (queryInput && queryInput.value) {
        const clearButton = document.createElement('button');
        clearButton.type = 'button';
        clearButton.textContent = 'Clear';
        clearButton.className = 'clear-search';
        clearButton.onclick = () => {
            queryInput.value = '';
            form.submit();
        };
        queryInput.parentNode.appendChild(clearButton);
    }
}

function initializeTableSort(table) {
    const headers = table.querySelectorAll('th');
    headers.forEach((header, index) => {
        if (!header.classList.contains('no-sort')) {
            header.style.cursor = 'pointer';
            header.addEventListener('click', () => sortTable(table, index));
            header.title = 'Click to sort';
        }
    });
}

function sortTable(table, column) {
    const tbody = table.querySelector('tbody');
    const rows = Array.from(tbody.querySelectorAll('tr'));
    const isAsc = table.querySelector('th').classList.contains('sort-asc');

    // Sort the rows
    rows.sort((a, b) => {
        let aVal = a.cells[column].textContent.trim();
        let bVal = b.cells[column].textContent.trim();

        // Handle date sorting
        if (isValidDate(aVal) && isValidDate(bVal)) {
            return compareStandardDates(aVal, bVal);
        }

        // Default string comparison
        return isAsc ? bVal.localeCompare(aVal) : aVal.localeCompare(bVal);
    });

    // Update sort indicators
    table.querySelectorAll('th').forEach(th => {
        th.classList.remove('sort-asc', 'sort-desc');
    });
    table.querySelector(`th:nth-child(${column + 1})`).classList.add(
        isAsc ? 'sort-desc' : 'sort-asc'
    );

    // Rebuild the table body
    rows.forEach(row => tbody.appendChild(row));
}

function isValidDate(dateStr) {
    const date = new Date(dateStr);
    return date instanceof Date && !isNaN(date);
}

function compareStandardDates(a, b) {
    return new Date(a) - new Date(b);
}

// Optional: Add keyboard navigation for accessibility
document.addEventListener('keydown', function(e) {
    if (e.key === 'Enter' && e.target.tagName === 'TH') {
        e.target.click();
    }
});