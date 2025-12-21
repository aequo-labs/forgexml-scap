/**
 * TabSystem - Flexible and reusable tab component
 */
class TabSystem {
    constructor(container, options = {}) {
        this.container = typeof container === 'string' ? 
            document.querySelector(container) : container;
        
        if (!this.container) {
            throw new Error('Tab container not found');
        }

        this.options = {
            animation: options.animation || 'none', // 'fade', 'slide', 'none'
            history: options.history || false,
            keyboard: options.keyboard || false,
            lazyLoad: options.lazyLoad || false,
            closable: options.closable || false,
            orientation: options.orientation || 'horizontal', // 'horizontal', 'vertical'
            onBeforeChange: options.onBeforeChange || null,
            onAfterChange: options.onAfterChange || null,
            loadContent: options.loadContent || null
        };

        this.tabsList = null;
        this.tabContent = null;
        this.tabs = [];
        this.activeTab = null;
        this.tabCounter = 0;
        this.eventHandlers = {};
        this.loadedTabs = new Set();

        this.init();
    }

    init() {
        this.tabsList = this.container.querySelector('.tabs ul');
        this.tabContent = this.container.querySelector('.tab-content');
        
        if (!this.tabsList || !this.tabContent) {
            throw new Error('Tab structure is invalid. Missing .tabs ul or .tab-content');
        }

        this.setupTabs();
        this.attachEventListeners();
        this.setActiveTab(this.getActiveTabIndex());
        
        if (this.options.history) {
            this.initHistory();
        }
    }

    setupTabs() {
        const tabItems = Array.from(this.tabsList.children);
        const tabPanes = Array.from(this.tabContent.children);

        this.tabs = tabItems.map((item, index) => {
            const link = item.querySelector('a');
            const pane = tabPanes[index];
            
            if (!link || !pane) {
                throw new Error(`Tab ${index} is missing link or pane`);
            }

            const tab = {
                id: this.generateTabId(link.textContent.trim()),
                title: link.textContent.trim(),
                element: item,
                link: link,
                pane: pane,
                index: index,
                active: item.classList.contains('is-active'),
                closable: this.options.closable || item.querySelector('.delete'),
                loaded: !this.options.lazyLoad
            };

            // Add data attributes
            item.dataset.tabId = tab.id;
            pane.dataset.tabId = tab.id;

            return tab;
        });

        // Mark initially loaded tabs
        if (this.options.lazyLoad) {
            this.tabs.forEach(tab => {
                if (tab.active) {
                    this.loadedTabs.add(tab.id);
                }
            });
        }
    }

    attachEventListeners() {
        // Tab click handling
        this.tabsList.addEventListener('click', this.handleTabClick.bind(this));
        
        // Close button handling
        if (this.options.closable) {
            this.tabsList.addEventListener('click', this.handleCloseClick.bind(this));
        }

        // Keyboard navigation
        if (this.options.keyboard) {
            this.tabsList.addEventListener('keydown', this.handleKeydown.bind(this));
            // Make tabs focusable
            this.tabs.forEach(tab => {
                tab.link.tabIndex = 0;
            });
        }

        // Window resize for responsive behavior
        window.addEventListener('resize', this.handleResize.bind(this));
    }

    handleTabClick(event) {
        event.preventDefault();
        const link = event.target.closest('a');
        if (!link) return;

        const tabItem = link.closest('li');
        const tabId = tabItem.dataset.tabId;
        const tab = this.tabs.find(t => t.id === tabId);
        
        if (tab && !tab.active) {
            this.activate(tab.id);
        }
    }

    handleCloseClick(event) {
        const closeButton = event.target.closest('.delete');
        if (!closeButton) return;

        event.preventDefault();
        event.stopPropagation();

        const tabItem = closeButton.closest('li');
        const tabId = tabItem.dataset.tabId;
        
        this.close(tabId);
    }

    handleKeydown(event) {
        const focusedLink = document.activeElement;
        const currentTab = this.tabs.find(tab => tab.link === focusedLink);
        if (!currentTab) return;

        let targetTab = null;

        switch (event.key) {
            case 'ArrowLeft':
            case 'ArrowUp':
                event.preventDefault();
                targetTab = this.getPreviousTab(currentTab.index);
                break;
            case 'ArrowRight':
            case 'ArrowDown':
                event.preventDefault();
                targetTab = this.getNextTab(currentTab.index);
                break;
            case 'Enter':
            case ' ':
                event.preventDefault();
                this.activate(currentTab.id);
                return;
            case 'Home':
                event.preventDefault();
                targetTab = this.tabs[0];
                break;
            case 'End':
                event.preventDefault();
                targetTab = this.tabs[this.tabs.length - 1];
                break;
            default:
                return;
        }

        if (targetTab) {
            targetTab.link.focus();
        }
    }

    handleResize() {
        // Handle responsive behavior
        this.adjustTabLayout();
    }

    async activate(tabId) {
        const tab = this.tabs.find(t => t.id === tabId);
        if (!tab) return false;

        // Call before change callback
        if (this.options.onBeforeChange) {
            const result = await this.options.onBeforeChange(tab, this.activeTab);
            if (result === false) return false;
        }

        const previousTab = this.activeTab;

        // Load content if lazy loading is enabled
        if (this.options.lazyLoad && !this.loadedTabs.has(tabId)) {
            await this.loadTabContent(tab);
        }

        // Deactivate current tab
        if (this.activeTab) {
            this.activeTab.element.classList.remove('is-active');
            this.activeTab.active = false; // Fix: properly clear active state
            this.hidePane(this.activeTab.pane);
        }

        // Activate new tab
        tab.element.classList.add('is-active');
        tab.active = true;
        this.activeTab = tab;

        this.showPane(tab.pane);

        // Update history
        if (this.options.history) {
            this.updateHistory(tabId);
        }

        // Emit change event
        this.emit('change', tab, previousTab);

        // Call after change callback
        if (this.options.onAfterChange) {
            this.options.onAfterChange(tab, previousTab);
        }

        return true;
    }

    async loadTabContent(tab) {
        if (this.options.loadContent) {
            try {
                const content = await this.options.loadContent(tab.id);
                if (content) {
                    tab.pane.innerHTML = content;
                }
            } catch (error) {
                console.error('Failed to load tab content:', error);
                tab.pane.innerHTML = `<div class="notification is-danger">Failed to load content</div>`;
            }
        }
        this.loadedTabs.add(tab.id);
    }

    hidePane(pane) {
        switch (this.options.animation) {
            case 'fade':
                pane.classList.remove('is-visible');
                setTimeout(() => {
                    pane.classList.remove('is-active');
                    pane.classList.add('tab-pane-hidden');
                }, 150);
                break;
            case 'slide':
                pane.classList.remove('is-visible');
                setTimeout(() => {
                    pane.classList.remove('is-active');
                    pane.classList.add('tab-pane-hidden');
                }, 300);
                break;
            default:
                pane.classList.remove('is-active');
                break;
        }
    }

    showPane(pane) {
        switch (this.options.animation) {
            case 'fade':
                pane.classList.remove('tab-pane-hidden');
                pane.classList.add('is-active', 'tab-pane-fade');
                // Force reflow
                pane.offsetHeight;
                pane.classList.add('is-visible');
                break;
            case 'slide':
                pane.classList.remove('tab-pane-hidden');
                pane.classList.add('is-active', 'tab-pane-slide', 'tab-pane-slide-enter');
                // Force reflow
                pane.offsetHeight;
                pane.classList.remove('tab-pane-slide-enter');
                pane.classList.add('is-visible');
                break;
            default:
                pane.classList.add('is-active');
                break;
        }
    }

    addTab(options) {
        const tab = {
            id: options.id || this.generateTabId(options.title),
            title: options.title,
            content: options.content || '',
            closable: options.closable !== false,
            loaded: true,
            index: this.tabs.length
        };

        // Create tab element
        const tabItem = document.createElement('li');
        tabItem.dataset.tabId = tab.id;
        
        const tabLink = document.createElement('a');
        tabLink.textContent = tab.title;
        tabLink.href = '#';
        
        tabItem.appendChild(tabLink);

        // Add close button if closable
        if (tab.closable) {
            const closeBtn = document.createElement('button');
            closeBtn.className = 'delete is-small';
            closeBtn.title = 'Close tab';
            tabItem.appendChild(closeBtn);
        }

        // Create pane element
        const pane = document.createElement('div');
        pane.className = 'tab-pane';
        pane.dataset.tabId = tab.id;
        pane.innerHTML = tab.content;

        // Add to DOM
        this.tabsList.appendChild(tabItem);
        this.tabContent.appendChild(pane);

        // Update tab object with DOM elements
        tab.element = tabItem;
        tab.link = tabLink;
        tab.pane = pane;

        // Add to tabs array
        this.tabs.push(tab);

        // Activate if it's the first tab or if specified
        if (options.activate || this.tabs.length === 1) {
            this.activate(tab.id);
        }

        this.emit('add', tab);
        return tab;
    }

    close(tabId) {
        const tabIndex = this.tabs.findIndex(t => t.id === tabId);
        if (tabIndex === -1) return false;

        const tab = this.tabs[tabIndex];
        
        // Prevent closing if it's the last tab
        if (this.tabs.length === 1) {
            return false;
        }

        // If closing active tab, activate another one
        if (tab.active) {
            const nextTab = this.getNextTab(tabIndex) || this.getPreviousTab(tabIndex);
            if (nextTab) {
                this.activate(nextTab.id);
            }
        }

        // Remove from DOM
        tab.element.remove();
        tab.pane.remove();

        // Remove from arrays
        this.tabs.splice(tabIndex, 1);
        this.loadedTabs.delete(tabId);

        // Update indices
        this.tabs.forEach((t, i) => {
            t.index = i;
        });

        this.emit('close', tab);
        return true;
    }

    getNextTab(currentIndex) {
        return this.tabs[currentIndex + 1] || this.tabs[0];
    }

    getPreviousTab(currentIndex) {
        return this.tabs[currentIndex - 1] || this.tabs[this.tabs.length - 1];
    }

    getActiveTabIndex() {
        return this.tabs.findIndex(tab => tab.active);
    }

    setActiveTab(index) {
        if (index >= 0 && index < this.tabs.length) {
            this.activate(this.tabs[index].id);
        }
    }

    generateTabId(title) {
        return title.toLowerCase()
            .replace(/[^a-z0-9]/g, '-')
            .replace(/-+/g, '-')
            .replace(/^-|-$/g, '') + '-' + (++this.tabCounter);
    }

    initHistory() {
        // Handle browser back/forward
        window.addEventListener('popstate', (event) => {
            const hash = window.location.hash.slice(1);
            if (hash) {
                const tab = this.tabs.find(t => t.id === hash);
                if (tab) {
                    this.activate(tab.id);
                }
            }
        });

        // Set initial hash if active tab exists
        if (this.activeTab && !window.location.hash) {
            window.location.hash = this.activeTab.id;
        }
    }

    updateHistory(tabId) {
        if (this.options.history) {
            window.location.hash = tabId;
        }
    }

    adjustTabLayout() {
        // Handle responsive tab layout
        if (window.innerWidth < 768) {
            this.container.classList.add('is-mobile');
        } else {
            this.container.classList.remove('is-mobile');
        }
    }

    // Configuration methods
    setAnimation(animation) {
        this.options.animation = animation;
    }

    setLazyLoad(enabled) {
        this.options.lazyLoad = enabled;
    }

    setClosable(enabled) {
        this.options.closable = enabled;
        this.tabs.forEach(tab => {
            const closeBtn = tab.element.querySelector('.delete');
            if (enabled && !closeBtn) {
                const btn = document.createElement('button');
                btn.className = 'delete is-small';
                btn.title = 'Close tab';
                tab.element.appendChild(btn);
            } else if (!enabled && closeBtn) {
                closeBtn.remove();
            }
        });
    }

    // Event system
    on(event, handler) {
        if (!this.eventHandlers[event]) {
            this.eventHandlers[event] = [];
        }
        this.eventHandlers[event].push(handler);
    }

    off(event, handler) {
        if (this.eventHandlers[event]) {
            const index = this.eventHandlers[event].indexOf(handler);
            if (index > -1) {
                this.eventHandlers[event].splice(index, 1);
            }
        }
    }

    emit(event, ...args) {
        if (this.eventHandlers[event]) {
            this.eventHandlers[event].forEach(handler => {
                handler.apply(this, args);
            });
        }
    }

    // Public API methods
    getActiveTab() {
        return this.activeTab;
    }

    getTabs() {
        return [...this.tabs];
    }

    getTab(tabId) {
        return this.tabs.find(t => t.id === tabId);
    }

    refresh() {
        this.loadedTabs.clear();
        if (this.activeTab && this.options.lazyLoad) {
            this.loadTabContent(this.activeTab);
        }
    }

    destroy() {
        // Remove event listeners
        this.tabsList.removeEventListener('click', this.handleTabClick);
        if (this.options.closable) {
            this.tabsList.removeEventListener('click', this.handleCloseClick);
        }
        if (this.options.keyboard) {
            this.tabsList.removeEventListener('keydown', this.handleKeydown);
        }
        window.removeEventListener('resize', this.handleResize);

        // Clear references
        this.tabs = [];
        this.activeTab = null;
        this.eventHandlers = {};
        this.loadedTabs.clear();
    }
}