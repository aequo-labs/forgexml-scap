/**
 * VirtualTreeView - High-performance tree view with virtual scrolling
 */
class VirtualTreeView {
    constructor(options) {
        this.options = {
            container: options.container,
            itemHeight: options.itemHeight || 28,
            bufferSize: options.bufferSize || 10,
            fetchData: options.fetchData || (() => Promise.resolve({ nodes: [], total: 0 })),
            onSelect: options.onSelect || (() => {}),
            onExpand: options.onExpand || (() => {}),
            onCollapse: options.onCollapse || (() => {}),
            multiSelect: options.multiSelect !== false,
            dragDrop: options.dragDrop !== false,
            lazyLoad: options.lazyLoad !== false
        };

        this.nodes = [];
        this.flatNodes = [];
        this.selectedNodes = new Set();
        this.expandedNodes = new Set();
        this.loadingNodes = new Set();
        this.filterText = '';
        this.scrollTop = 0;
        this.containerHeight = 0;
        this.totalHeight = 0;
        this.visibleStart = 0;
        this.visibleEnd = 0;
        this.renderTime = 0;

        this.viewport = null;
        this.content = null;
        this.initialized = false;
    }

    async initialize() {
        this.setupDOM();
        this.attachEventListeners();
        await this.loadRootNodes();
        this.updateFlatNodes();
        this.render();
        this.initialized = true;
    }

    setupDOM() {
        const container = this.options.container;
        container.innerHTML = `
            <div class="tree-viewport">
                <div class="tree-content"></div>
            </div>
        `;
        this.viewport = container.querySelector('.tree-viewport');
        this.content = container.querySelector('.tree-content');
        
        // Set initial dimensions
        this.containerHeight = this.viewport.offsetHeight;
    }

    attachEventListeners() {
        // Scroll handling
        this.viewport.addEventListener('scroll', this.handleScroll.bind(this));
        
        // Click handling
        this.content.addEventListener('click', this.handleClick.bind(this));
        
        // Keyboard navigation
        this.viewport.addEventListener('keydown', this.handleKeydown.bind(this));
        
        // Drag and drop
        if (this.options.dragDrop) {
            this.content.addEventListener('dragstart', this.handleDragStart.bind(this));
            this.content.addEventListener('dragover', this.handleDragOver.bind(this));
            this.content.addEventListener('drop', this.handleDrop.bind(this));
            this.content.addEventListener('dragend', this.handleDragEnd.bind(this));
        }
    }

    async loadRootNodes() {
        const result = await this.options.fetchData(null, 0, 50);
        this.nodes = result.nodes || [];
        this.totalNodesCount = result.total || this.nodes.length;
    }

    updateFlatNodes() {
        const startTime = performance.now();
        this.flatNodes = [];
        
        const processNode = (node, level = 0) => {
            if (this.filterText && !this.matchesFilter(node)) {
                return;
            }
            
            node.level = level;
            this.flatNodes.push(node);
            
            if (this.expandedNodes.has(node.id) && node.children) {
                node.children.forEach(child => processNode(child, level + 1));
            }
        };
        
        this.nodes.forEach(node => processNode(node, 0));
        
        this.totalHeight = this.flatNodes.length * this.options.itemHeight;
        this.content.style.height = `${this.totalHeight}px`;
        
        this.renderTime = Math.round(performance.now() - startTime);
    }

    matchesFilter(node) {
        if (!this.filterText) return true;
        const searchLower = this.filterText.toLowerCase();
        return node.label.toLowerCase().includes(searchLower) ||
               (node.id && node.id.toLowerCase().includes(searchLower));
    }

    render() {
        const scrollTop = this.viewport.scrollTop;
        this.visibleStart = Math.floor(scrollTop / this.options.itemHeight);
        this.visibleEnd = Math.ceil((scrollTop + this.containerHeight) / this.options.itemHeight);
        
        // Add buffer
        this.visibleStart = Math.max(0, this.visibleStart - this.options.bufferSize);
        this.visibleEnd = Math.min(this.flatNodes.length, this.visibleEnd + this.options.bufferSize);
        
        // Clear content
        this.content.innerHTML = '';
        
        // Render visible nodes
        for (let i = this.visibleStart; i < this.visibleEnd; i++) {
            const node = this.flatNodes[i];
            if (node) {
                const element = this.createNodeElement(node, i);
                this.content.appendChild(element);
            }
        }
    }

    createNodeElement(node, index) {
        const div = document.createElement('div');
        div.className = 'tree-node';
        div.dataset.nodeId = node.id;
        div.dataset.index = index;
        div.style.top = `${index * this.options.itemHeight}px`;
        div.style.paddingLeft = `${node.level * 20}px`;
        
        if (this.selectedNodes.has(node.id)) {
            div.classList.add('selected');
        }
        
        if (this.loadingNodes.has(node.id)) {
            div.classList.add('loading');
        }
        
        // Arrow for expandable nodes
        const arrow = document.createElement('span');
        arrow.className = 'tree-node-arrow';
        if (node.hasChildren) {
            arrow.innerHTML = '<i class="fas fa-chevron-right"></i>';
            if (this.expandedNodes.has(node.id)) {
                arrow.classList.add('expanded');
            }
        } else {
            arrow.classList.add('no-children');
        }
        
        // Icon - ensure Font Awesome style prefix is present
        const icon = document.createElement('span');
        icon.className = 'tree-node-icon';
        let iconClass = node.icon || 'fas fa-file';
        // Add 'fas' prefix if icon doesn't already have a FA style prefix (fas, far, fab, fal, fad)
        if (iconClass && !iconClass.match(/^fa[srldb]\s/)) {
            iconClass = 'fas ' + iconClass;
        }
        icon.innerHTML = `<i class="${iconClass}"></i>`;
        
        // Label
        const label = document.createElement('span');
        label.className = 'tree-node-label';
        label.textContent = node.label;
        
        // Badge (optional)
        if (node.badge) {
            const badge = document.createElement('span');
            badge.className = 'tree-node-badge';
            badge.textContent = node.badge;
            div.appendChild(badge);
        }
        
        div.appendChild(arrow);
        div.appendChild(icon);
        div.appendChild(label);
        
        // Make draggable if enabled
        if (this.options.dragDrop) {
            div.draggable = true;
        }
        
        return div;
    }

    handleScroll() {
        this.scrollTop = this.viewport.scrollTop;
        this.render();
    }

    async handleClick(event) {
        const nodeElement = event.target.closest('.tree-node');
        if (!nodeElement) return;
        
        const nodeId = nodeElement.dataset.nodeId;
        const node = this.findNode(nodeId);
        if (!node) return;
        
        // Check if arrow was clicked
        const arrow = event.target.closest('.tree-node-arrow');
        if (arrow && node.hasChildren) {
            await this.toggleNode(node);
            return;
        }
        
        // Handle selection
        if (event.ctrlKey || event.metaKey) {
            // Multi-select
            if (this.selectedNodes.has(nodeId)) {
                this.selectedNodes.delete(nodeId);
            } else {
                this.selectedNodes.add(nodeId);
            }
        } else {
            // Single select
            this.selectedNodes.clear();
            this.selectedNodes.add(nodeId);
        }
        
        this.render();
        this.options.onSelect(node, event);
    }

    async toggleNode(node) {
        console.log('toggleNode called:', node.id, 'hasChildren:', node.hasChildren, 'children:', node.children);
        if (this.expandedNodes.has(node.id)) {
            this.expandedNodes.delete(node.id);
            this.options.onCollapse(node);
        } else {
            this.expandedNodes.add(node.id);
            console.log('Expanding node:', node.id);
            
            // Load children if needed
            if (!node.children && node.hasChildren && this.options.lazyLoad) {
                console.log('Loading children for:', node.id);
                this.loadingNodes.add(node.id);
                this.render();
                
                try {
                    const result = await this.options.fetchData(node.id, 0, 50);
                    console.log('Fetched children:', result);
                    node.children = result.nodes || [];
                    node.childrenLoaded = true;
                } catch (err) {
                    console.error('Error fetching children:', err);
                }
                
                this.loadingNodes.delete(node.id);
            }
            
            await this.options.onExpand(node);
        }
        
        this.updateFlatNodes();
        console.log('After updateFlatNodes, flatNodes count:', this.flatNodes.length);
        this.render();
    }

    handleKeydown(event) {
        const selectedId = Array.from(this.selectedNodes)[0];
        if (!selectedId) return;
        
        const currentIndex = this.flatNodes.findIndex(n => n.id === selectedId);
        if (currentIndex === -1) return;
        
        let newIndex = currentIndex;
        
        switch (event.key) {
            case 'ArrowUp':
                event.preventDefault();
                newIndex = Math.max(0, currentIndex - 1);
                break;
            case 'ArrowDown':
                event.preventDefault();
                newIndex = Math.min(this.flatNodes.length - 1, currentIndex + 1);
                break;
            case 'ArrowLeft':
                event.preventDefault();
                const currentNode = this.flatNodes[currentIndex];
                if (this.expandedNodes.has(currentNode.id)) {
                    this.toggleNode(currentNode);
                } else if (currentNode.parent) {
                    // Navigate to parent
                    const parentIndex = this.flatNodes.findIndex(n => n.id === currentNode.parent);
                    if (parentIndex !== -1) {
                        newIndex = parentIndex;
                    }
                }
                break;
            case 'ArrowRight':
                event.preventDefault();
                const node = this.flatNodes[currentIndex];
                if (node.hasChildren && !this.expandedNodes.has(node.id)) {
                    this.toggleNode(node);
                } else if (this.expandedNodes.has(node.id) && node.children && node.children.length > 0) {
                    // Navigate to first child
                    newIndex = currentIndex + 1;
                }
                break;
            case 'Enter':
            case ' ':
                event.preventDefault();
                this.toggleNode(this.flatNodes[currentIndex]);
                break;
            default:
                return;
        }
        
        if (newIndex !== currentIndex) {
            const newNode = this.flatNodes[newIndex];
            this.selectedNodes.clear();
            this.selectedNodes.add(newNode.id);
            this.ensureNodeVisible(newIndex);
            this.render();
            this.options.onSelect(newNode, event);
        }
    }

    ensureNodeVisible(index) {
        const nodeTop = index * this.options.itemHeight;
        const nodeBottom = nodeTop + this.options.itemHeight;
        const scrollTop = this.viewport.scrollTop;
        const scrollBottom = scrollTop + this.containerHeight;
        
        if (nodeTop < scrollTop) {
            this.viewport.scrollTop = nodeTop;
        } else if (nodeBottom > scrollBottom) {
            this.viewport.scrollTop = nodeBottom - this.containerHeight;
        }
    }

    handleDragStart(event) {
        const nodeElement = event.target.closest('.tree-node');
        if (!nodeElement) return;
        
        const nodeId = nodeElement.dataset.nodeId;
        event.dataTransfer.effectAllowed = 'move';
        event.dataTransfer.setData('text/plain', nodeId);
        nodeElement.classList.add('dragging');
    }

    handleDragOver(event) {
        event.preventDefault();
        event.dataTransfer.dropEffect = 'move';
        
        const nodeElement = event.target.closest('.tree-node');
        if (nodeElement) {
            nodeElement.classList.add('drag-over');
        }
    }

    handleDrop(event) {
        event.preventDefault();
        
        const draggedId = event.dataTransfer.getData('text/plain');
        const targetElement = event.target.closest('.tree-node');
        
        if (targetElement) {
            const targetId = targetElement.dataset.nodeId;
            this.moveNode(draggedId, targetId);
        }
        
        // Clean up
        document.querySelectorAll('.drag-over').forEach(el => {
            el.classList.remove('drag-over');
        });
    }

    handleDragEnd(event) {
        document.querySelectorAll('.dragging').forEach(el => {
            el.classList.remove('dragging');
        });
        document.querySelectorAll('.drag-over').forEach(el => {
            el.classList.remove('drag-over');
        });
    }

    moveNode(sourceId, targetId) {
        // Implementation would depend on tree structure requirements
        console.log(`Moving node ${sourceId} to ${targetId}`);
        // Update tree structure and re-render
        this.updateFlatNodes();
        this.render();
    }

    findNode(nodeId, nodes = this.nodes) {
        for (const node of nodes) {
            if (node.id === nodeId) return node;
            if (node.children) {
                const found = this.findNode(nodeId, node.children);
                if (found) return found;
            }
        }
        return null;
    }

    filter(text) {
        this.filterText = text;
        if (text) {
            // Expand all nodes that match or have matching children
            this.expandMatchingNodes();
        }
        this.updateFlatNodes();
        this.render();
    }

    expandMatchingNodes() {
        const expandIfMatches = (node) => {
            let hasMatch = this.matchesFilter(node);
            
            if (node.children) {
                for (const child of node.children) {
                    if (expandIfMatches(child)) {
                        hasMatch = true;
                        this.expandedNodes.add(node.id);
                    }
                }
            }
            
            return hasMatch;
        };
        
        this.nodes.forEach(node => expandIfMatches(node));
    }

    async expandAll() {
        const expandNode = async (node) => {
            if (node.hasChildren) {
                this.expandedNodes.add(node.id);
                if (!node.children && this.options.lazyLoad) {
                    const result = await this.options.fetchData(node.id, 0, 50);
                    node.children = result.nodes || [];
                }
                if (node.children) {
                    for (const child of node.children) {
                        await expandNode(child);
                    }
                }
            }
        };
        
        for (const node of this.nodes) {
            await expandNode(node);
        }
        
        this.updateFlatNodes();
        this.render();
    }

    collapseAll() {
        this.expandedNodes.clear();
        this.updateFlatNodes();
        this.render();
    }

    async refresh() {
        this.nodes = [];
        this.flatNodes = [];
        this.selectedNodes.clear();
        this.expandedNodes.clear();
        await this.loadRootNodes();
        this.updateFlatNodes();
        this.render();
    }

    updateNode(node) {
        this.updateFlatNodes();
        this.render();
    }

    getSelectedNodes() {
        return Array.from(this.selectedNodes).map(id => this.findNode(id)).filter(Boolean);
    }

    getStats() {
        return {
            visibleNodes: this.visibleEnd - this.visibleStart,
            totalNodes: this.flatNodes.length,
            renderTime: this.renderTime
        };
    }

    handleResize() {
        this.containerHeight = this.viewport.offsetHeight;
        this.render();
    }
}

/**
 * VirtualTree - Simplified wrapper for VirtualTreeView
 * Provides easier initialization and common use cases
 */
class VirtualTree {
    constructor(containerId, options = {}) {
        // Support both string ID and DOM element
        const container = typeof containerId === 'string' 
            ? document.getElementById(containerId)
            : containerId;
        
        if (!container) {
            console.error(`VirtualTree: Container not found: ${containerId}`);
            this.container = document.createElement('div');
            this.container.innerHTML = '<div class="notification is-danger">Tree container not found</div>';
            return;
        }
        
        this.container = container;
        this.data = [];
        
        // Merge options with defaults
        this.options = {
            itemHeight: options.itemHeight || 80,
            bufferSize: options.bufferSize || 10,
            expandable: options.expandable !== false,
            selectable: options.selectable !== false,
            ...options
        };
        
        // Create fetchData function based on options
        const fetchData = options.fetchData || this.createDefaultFetchData();
        
        // Initialize the underlying VirtualTreeView
        this.treeView = new VirtualTreeView({
            container: container,
            itemHeight: this.options.itemHeight,
            bufferSize: this.options.bufferSize,
            fetchData: fetchData,
            onSelect: options.onSelect || (() => {}),
            onExpand: options.onExpand || (() => {}),
            onCollapse: options.onCollapse || (() => {})
        });
        
        // Auto-initialize
        this.treeView.initialize().catch(err => {
            console.error('VirtualTree initialization failed:', err);
            this.container.innerHTML = `
                <div class="notification is-warning">
                    <p><strong>Tree initialization failed</strong></p>
                    <p class="is-size-7">${err.message}</p>
                </div>
            `;
        });
    }
    
    createDefaultFetchData() {
        // Return a function that uses the local data
        return async (parentId, offset, limit) => {
            // If we have static data, use it
            if (this.data && this.data.length > 0) {
                return {
                    nodes: this.data,
                    total: this.data.length,
                    hasMore: false
                };
            }
            
            // Otherwise return empty
            return {
                nodes: [],
                total: 0,
                hasMore: false
            };
        };
    }
    
    // Delegate common methods to underlying tree view
    setData(data) {
        this.data = data;
        if (this.treeView && this.treeView.initialized) {
            this.treeView.nodes = data;
            this.treeView.updateFlatNodes();
            this.treeView.render();
        }
    }
    
    expandAll() {
        if (this.treeView) {
            return this.treeView.expandAll();
        }
    }
    
    collapseAll() {
        if (this.treeView) {
            this.treeView.collapseAll();
        }
    }
    
    refresh() {
        if (this.treeView) {
            return this.treeView.refresh();
        }
    }
    
    filter(text) {
        if (this.treeView) {
            this.treeView.filter(text);
        }
    }
    
    getSelectedNodes() {
        return this.treeView ? this.treeView.getSelectedNodes() : [];
    }
    
    getStats() {
        return this.treeView ? this.treeView.getStats() : { visibleNodes: 0, totalNodes: 0, renderTime: 0 };
    }
}
