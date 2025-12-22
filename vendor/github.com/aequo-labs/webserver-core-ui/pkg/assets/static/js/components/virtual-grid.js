/**
 * VirtualDataGrid - Clean windowed virtual scrolling implementation
 * Uses backend API for data windows and maintains only a sliding window in RAM
 */
class VirtualDataGrid {
    constructor(options) {
        this.options = {
            container: options.container,
            columns: options.columns || [],
            rowHeight: options.rowHeight || 40,
            headerHeight: options.headerHeight || 45,
            pageSize: options.pageSize || 50,
            fetchData: options.fetchData || (() => Promise.resolve({ data: [], total: 0 })),
            onCellEdit: options.onCellEdit || (() => {}),
            onSelectionChange: options.onSelectionChange || (() => {}),
            multiSelect: options.multiSelect !== false,
            editable: options.editable || false,
            resizable: options.resizable || false,
            sortable: options.sortable || false,
            filterable: options.filterable || false
        };

        // Windowed data management
        this.windowSize = 5; // Keep 5 pages in memory (250 rows by default)
        this.windowData = []; // Linear array of loaded rows in current window
        this.windowStartRow = 0; // First row index in current window
        this.totalRows = 0;
        this.loadingPages = new Set();
        
        // UI state
        this.selectedRows = new Set();
        this.sortColumn = null;
        this.sortDirection = 'asc';
        this.filters = {};
        this.searchText = '';
        this.scrollTop = 0;
        this.containerHeight = 0;
        this.containerWidth = 0;
        this.visibleStart = 0;
        this.visibleEnd = 0;
        this.renderTime = 0;
        this.editingCell = null;
        
        // Column management
        this.columnWidths = new Map();
        this.visibleColumns = new Set();
        this.options.columns.forEach(col => {
            this.columnWidths.set(col.id, col.width || 100);
            if (col.visible !== false) {
                this.visibleColumns.add(col.id);
            }
        });

        this.viewport = null;
        this.header = null;
        this.body = null;
        this.initialized = false;
    }

    async initialize() {
        this.setupDOM();
        this.attachEventListeners();
        await this.loadWindow(0);
        this.render();
        this.initialized = true;
    }

    setupDOM() {
        const container = this.options.container;
        container.innerHTML = `
            <div class="grid-header"></div>
            <div class="grid-viewport">
                <div class="grid-body"></div>
            </div>
        `;
        
        this.header = container.querySelector('.grid-header');
        this.viewport = container.querySelector('.grid-viewport');
        this.body = container.querySelector('.grid-body');
        
        this.containerHeight = this.options.container.offsetHeight - this.options.headerHeight;
        this.containerWidth = this.options.container.offsetWidth;
        
        this.header.style.cssText = `
            height: ${this.options.headerHeight}px;
            overflow: hidden;
            border-bottom: 2px solid var(--bulma-border);
            background: var(--bulma-scheme-main-bis);
            position: relative;
        `;
        
        this.viewport.style.cssText = `
            height: ${this.containerHeight}px;
            overflow: auto;
            position: relative;
        `;
        
        this.body.style.cssText = `
            position: relative;
            min-height: 100%;
        `;
    }

    attachEventListeners() {
        let scrollTimeout;
        
        this.viewport.addEventListener('scroll', async () => {
            this.scrollTop = this.viewport.scrollTop;
            this.render();
            
            clearTimeout(scrollTimeout);
            scrollTimeout = setTimeout(() => {
                this.checkWindowBounds();
            }, 100);
        });
    }

    async loadWindow(startRow) {
        if (this.loadingPages.has(startRow)) return;
        
        this.loadingPages.add(startRow);
        const windowSize = this.windowSize * this.options.pageSize;
        
        try {
            const result = await this.options.fetchData({
                offset: startRow,
                limit: windowSize,
                sortColumn: this.sortColumn,
                sortDirection: this.sortDirection,
                filters: this.filters,
                search: this.searchText
            });
            
            this.windowData = result.data;
            this.windowStartRow = startRow;
            this.totalRows = result.total;
            
            // Update body height for scrolling
            const totalHeight = this.totalRows * this.options.rowHeight;
            this.body.style.height = `${totalHeight}px`;
            
        } finally {
            this.loadingPages.delete(startRow);
        }
    }

    async checkWindowBounds() {
        const currentRow = Math.floor(this.scrollTop / this.options.rowHeight);
        const windowEndRow = this.windowStartRow + this.windowData.length;
        
        // Check if we need to load a new window
        const bufferRows = this.options.pageSize; // One page buffer
        
        if (currentRow < this.windowStartRow + bufferRows || 
            currentRow > windowEndRow - bufferRows) {
            
            // Calculate new window start (centered around current position)
            const newWindowStart = Math.max(0, 
                Math.floor((currentRow - (this.windowSize * this.options.pageSize) / 2) 
                / this.options.pageSize) * this.options.pageSize
            );
            
            if (newWindowStart !== this.windowStartRow) {
                await this.loadWindow(newWindowStart);
                this.render();
            }
        }
    }

    render() {
        const startTime = performance.now();
        
        // Calculate visible range
        this.visibleStart = Math.floor(this.scrollTop / this.options.rowHeight);
        this.visibleEnd = Math.ceil((this.scrollTop + this.containerHeight) / this.options.rowHeight);
        
        // Add buffer
        const bufferSize = 5;
        this.visibleStart = Math.max(0, this.visibleStart - bufferSize);
        this.visibleEnd = Math.min(this.totalRows, this.visibleEnd + bufferSize);
        
        this.renderHeader();
        this.renderRows();
        
        this.renderTime = Math.round(performance.now() - startTime);
    }

    renderHeader() {
        this.header.innerHTML = '';
        
        const headerRow = document.createElement('div');
        headerRow.className = 'grid-header-row';
        headerRow.style.cssText = `
            display: flex;
            height: 100%;
            align-items: center;
            position: relative;
        `;
        
        this.getVisibleColumns().forEach(col => {
            const width = this.columnWidths.get(col.id);
            const headerCell = document.createElement('div');
            headerCell.className = 'grid-header-cell';
            headerCell.dataset.columnId = col.id;
            headerCell.style.cssText = `
                width: ${width}px;
                padding: 0.75em;
                border-right: 1px solid var(--bulma-table-cell-border-color, #dbdbdb);
                display: flex;
                align-items: center;
                justify-content: space-between;
                cursor: ${this.options.sortable ? 'pointer' : 'default'};
                position: relative;
                user-select: none;
                font-weight: 700;
                background-color: var(--bulma-table-head-background-color, transparent);
                color: var(--bulma-table-head-color, #363636);
            `;
            
            headerCell.textContent = col.label;
            
            if (this.options.sortable && this.sortColumn === col.id) {
                const sortIcon = document.createElement('i');
                sortIcon.className = `fas fa-sort-${this.sortDirection === 'asc' ? 'up' : 'down'}`;
                sortIcon.style.marginLeft = '5px';
                headerCell.appendChild(sortIcon);
            }
            
            headerRow.appendChild(headerCell);
        });
        
        this.header.appendChild(headerRow);
    }

    renderRows() {
        this.body.innerHTML = '';
        
        for (let i = this.visibleStart; i < this.visibleEnd && i < this.totalRows; i++) {
            const windowIndex = i - this.windowStartRow;
            const row = this.windowData[windowIndex];
            
            if (row) {
                const rowElement = this.createRowElement(row, i);
                this.body.appendChild(rowElement);
            } else {
                const placeholderRow = this.createPlaceholderRow(i);
                this.body.appendChild(placeholderRow);
            }
        }
    }

    createRowElement(row, index) {
        const rowDiv = document.createElement('div');
        rowDiv.className = 'grid-row';
        rowDiv.dataset.rowId = row.id || row.ID || index;
        rowDiv.dataset.index = index;
        
        // Add even/odd class for styling
        if (index % 2 === 0) {
            rowDiv.classList.add('grid-row-even');
        } else {
            rowDiv.classList.add('grid-row-odd');
        }
        
        rowDiv.style.cssText = `
            position: absolute;
            top: ${index * this.options.rowHeight}px;
            height: ${this.options.rowHeight}px;
            width: 100%;
            display: flex;
            align-items: center;
            border-bottom: 1px solid var(--bulma-border-weak);
        `;
        
        if (this.selectedRows.has(row.id)) {
            rowDiv.classList.add('selected');
        }
        
        // Add row click handler if provided
        if (this.options.onRowClick) {
            rowDiv.style.cursor = 'pointer';
            rowDiv.addEventListener('click', (e) => {
                // Don't trigger row click if clicking on a link or button
                if (e.target.closest('a, button')) return;
                this.options.onRowClick(row, index, e);
            });
        }
        
        this.getVisibleColumns().forEach(col => {
            const width = this.columnWidths.get(col.id);
            const cell = document.createElement('div');
            cell.className = 'grid-cell';
            cell.dataset.columnId = col.id;
            cell.style.cssText = `
                width: ${width}px;
                padding: 0.75em;
                border-right: 1px solid var(--bulma-table-cell-border-color, #dbdbdb);
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                vertical-align: top;
            `;
            
            const value = row[col.id];
            
            // Check if column has custom render function
            if (typeof col.render === 'function') {
                cell.innerHTML = col.render(value, row, index);
                cell.style.whiteSpace = 'normal'; // Allow wrapping for custom HTML
            } else {
                cell.textContent = this.formatCellValue(value, col);
                cell.title = cell.textContent;
            }
            
            rowDiv.appendChild(cell);
        });
        
        return rowDiv;
    }
    createPlaceholderRow(index) {
        const rowDiv = document.createElement('div');
        rowDiv.className = 'grid-row loading';
        rowDiv.dataset.index = index;
        rowDiv.style.cssText = `
            position: absolute;
            top: ${index * this.options.rowHeight}px;
            height: ${this.options.rowHeight}px;
            width: 100%;
            display: flex;
            align-items: center;
            border-bottom: 1px solid var(--bulma-border-weak);
            background-color: var(--bulma-scheme-main-ter);
            opacity: 0.7;
        `;
        
        this.getVisibleColumns().forEach(col => {
            const width = this.columnWidths.get(col.id);
            const cell = document.createElement('div');
            cell.className = 'grid-cell loading-cell';
            cell.style.cssText = `
                width: ${width}px;
                padding: 0.75em;
                border-right: 1px solid var(--bulma-table-cell-border-color, #dbdbdb);
                overflow: hidden;
                vertical-align: top;
            `;
            
            const skeleton = document.createElement('div');
            skeleton.style.cssText = `
                height: 16px;
                background: linear-gradient(90deg, transparent, rgba(0,0,0,0.1), transparent);
                background-size: 200% 100%;
                animation: shimmer 1.5s infinite;
                border-radius: 3px;
            `;
            
            cell.appendChild(skeleton);
            rowDiv.appendChild(cell);
        });
        
        return rowDiv;
    }

    formatCellValue(value, column) {
        if (value == null) return '';
        
        switch (column.format) {
            case 'currency':
                return new Intl.NumberFormat('en-US', {
                    style: 'currency',
                    currency: 'USD'
                }).format(value);
            case 'percentage':
                return (value * 100).toFixed(1) + '%';
            case 'date':
                return value instanceof Date ? 
                    value.toLocaleDateString() : 
                    new Date(value).toLocaleDateString();
            default:
                return String(value);
        }
    }

    getVisibleColumns() {
        return this.options.columns.filter(col => this.visibleColumns.has(col.id));
    }

    async refresh() {
        this.windowData = [];
        this.windowStartRow = 0;
        this.selectedRows.clear();
        await this.loadWindow(0);
        this.render();
    }

    getStats() {
        return {
            loadedRows: this.windowData.length,
            totalRows: this.totalRows,
            visibleRows: this.visibleEnd - this.visibleStart,
            renderTime: this.renderTime,
            selectedRows: this.selectedRows.size,
            windowStart: this.windowStartRow,
            windowEnd: this.windowStartRow + this.windowData.length
        };
    }

    search(text) {
        this.searchText = text;
        this.refresh();
    }

    setPageSize(size) {
        this.options.pageSize = size;
        this.refresh();
    }

    handleResize() {
        this.containerHeight = this.viewport.offsetHeight;
        this.containerWidth = this.options.container.offsetWidth;
        this.render();
    }
}

/**
 * VirtualGrid - Simplified wrapper for VirtualDataGrid
 * Provides easier initialization and common use cases
 */
class VirtualGrid {
    constructor(containerId, options = {}) {
        // Support both string ID and DOM element
        const container = typeof containerId === 'string' 
            ? document.getElementById(containerId)
            : containerId;
        
        if (!container) {
            console.error(`VirtualGrid: Container not found: ${containerId}`);
            this.container = document.createElement('div');
            this.container.innerHTML = '<div class="notification is-danger">Grid container not found</div>';
            return;
        }
        
        this.container = container;
        this.data = [];
        
        // Merge options with defaults
        this.options = {
            columns: options.columns || [],
            rowHeight: options.rowHeight || 40,
            headerHeight: options.headerHeight || 45,
            pageSize: options.pageSize || 50,
            multiSelect: options.multiSelect !== false,
            editable: options.editable || false,
            sortable: options.sortable !== false,
            filterable: options.filterable || false,
            ...options
        };
        
        // Create fetchData function based on options
        const fetchData = options.fetchData || this.createDefaultFetchData();
        
        // Initialize the underlying VirtualDataGrid
        this.gridView = new VirtualDataGrid({
            container: container,
            columns: this.options.columns,
            rowHeight: this.options.rowHeight,
            headerHeight: this.options.headerHeight,
            pageSize: this.options.pageSize,
            fetchData: fetchData,
            onCellEdit: options.onCellEdit || (() => {}),
            onSelectionChange: options.onSelectionChange || (() => {}),
            multiSelect: this.options.multiSelect,
            editable: this.options.editable,
            sortable: this.options.sortable,
            filterable: this.options.filterable
        });
        
        // Auto-initialize
        this.gridView.initialize().catch(err => {
            console.error('VirtualGrid initialization failed:', err);
            this.container.innerHTML = `
                <div class="notification is-warning">
                    <p><strong>Grid initialization failed</strong></p>
                    <p class="is-size-7">${err.message}</p>
                </div>
            `;
        });
    }
    
    createDefaultFetchData() {
        // Return a function that uses the local data
        return async (offset, limit, sort, filters) => {
            // If we have static data, use it
            if (this.data && this.data.length > 0) {
                const start = offset || 0;
                const end = Math.min(start + (limit || 50), this.data.length);
                return {
                    data: this.data.slice(start, end),
                    total: this.data.length,
                    hasMore: end < this.data.length
                };
            }
            
            // Otherwise return empty
            return {
                data: [],
                total: 0,
                hasMore: false
            };
        };
    }
    
    // Delegate common methods to underlying grid view
    setData(data) {
        this.data = data;
        if (this.gridView && this.gridView.initialized) {
            this.gridView.refresh();
        }
    }
    
    refresh() {
        if (this.gridView) {
            return this.gridView.refresh();
        }
    }
    
    filter(filters) {
        if (this.gridView) {
            this.gridView.applyFilter(filters);
        }
    }
    
    sort(column, direction) {
        if (this.gridView) {
            this.gridView.sortBy(column, direction);
        }
    }
    
    getSelectedRows() {
        return this.gridView ? this.gridView.getSelectedRows() : [];
    }
    
    getStats() {
        return this.gridView ? this.gridView.getStats() : { 
            visibleRows: 0, 
            totalRows: 0, 
            renderTime: 0,
            loadedRows: 0
        };
    }
}
