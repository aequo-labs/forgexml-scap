// confirm-modal.js - Reusable Confirmation Modal functionality

let confirmModalCallback = null;
let confirmModalData = null;

/**
 * Show a confirmation modal
 * @param {Object} options - Configuration options
 * @param {string} options.title - Modal title (default: "Confirm")
 * @param {string} options.message - Message to display
 * @param {string} options.confirmText - Text for confirm button (default: "Confirm")
 * @param {string} options.confirmClass - CSS class for confirm button (default: "is-primary")
 * @param {string} options.icon - FontAwesome icon class (default: "fa-question-circle")
 * @param {string} options.confirmIcon - Icon for confirm button (default: "fa-check")
 * @param {Function} options.onConfirm - Callback when confirmed
 * @param {any} options.data - Optional data to pass to callback
 */
function showConfirmModal(options) {
    const modal = document.getElementById('confirm-modal');
    if (!modal) {
        console.error('Confirm modal not found. Make sure to include the confirm-modal template.');
        // Fallback to native confirm
        if (confirm(options.message || 'Are you sure?')) {
            if (options.onConfirm) options.onConfirm(options.data);
        }
        return;
    }

    // Set title
    const titleEl = document.getElementById('confirm-modal-title-text');
    if (titleEl) titleEl.textContent = options.title || 'Confirm';

    // Set icon
    const iconEl = document.getElementById('confirm-modal-icon');
    if (iconEl) {
        iconEl.className = 'fas ' + (options.icon || 'fa-question-circle');
    }

    // Set message
    const messageEl = document.getElementById('confirm-modal-message');
    if (messageEl) messageEl.textContent = options.message || 'Are you sure?';

    // Set confirm button
    const confirmBtn = document.getElementById('confirm-modal-confirm-btn');
    if (confirmBtn) {
        // Remove old classes and add new ones
        confirmBtn.className = 'button ' + (options.confirmClass || 'is-primary');
    }

    const btnTextEl = document.getElementById('confirm-modal-btn-text');
    if (btnTextEl) btnTextEl.textContent = options.confirmText || 'Confirm';

    const btnIconEl = document.getElementById('confirm-modal-btn-icon');
    if (btnIconEl) {
        btnIconEl.className = 'fas ' + (options.confirmIcon || 'fa-check');
    }

    // Store callback and data
    confirmModalCallback = options.onConfirm || null;
    confirmModalData = options.data || null;

    // Show modal
    modal.classList.add('is-active');
}

/**
 * Close the confirmation modal
 */
function closeConfirmModal() {
    const modal = document.getElementById('confirm-modal');
    if (modal) {
        modal.classList.remove('is-active');
    }
    confirmModalCallback = null;
    confirmModalData = null;
}

/**
 * Execute the confirm action
 */
function confirmModalAction() {
    const callback = confirmModalCallback;
    const data = confirmModalData;
    
    closeConfirmModal();
    
    if (callback) {
        callback(data);
    }
}

/**
 * Convenience function for delete confirmations
 */
function showDeleteConfirm(message, onConfirm, data) {
    showConfirmModal({
        title: 'Delete',
        message: message,
        confirmText: 'Delete',
        confirmClass: 'is-danger',
        icon: 'fa-trash',
        confirmIcon: 'fa-trash',
        onConfirm: onConfirm,
        data: data
    });
}

/**
 * Convenience function for warning confirmations
 */
function showWarningConfirm(title, message, confirmText, onConfirm, data) {
    showConfirmModal({
        title: title,
        message: message,
        confirmText: confirmText || 'Continue',
        confirmClass: 'is-warning',
        icon: 'fa-exclamation-triangle',
        confirmIcon: 'fa-check',
        onConfirm: onConfirm,
        data: data
    });
}

/**
 * Show an alert modal (information only, no callback)
 */
function showAlertModal(title, message, icon) {
    const modal = document.getElementById('confirm-modal');
    if (!modal) {
        alert(message);
        return;
    }

    // Set title
    const titleEl = document.getElementById('confirm-modal-title-text');
    if (titleEl) titleEl.textContent = title || 'Notice';

    // Set icon
    const iconEl = document.getElementById('confirm-modal-icon');
    if (iconEl) {
        iconEl.className = 'fas ' + (icon || 'fa-info-circle');
    }

    // Set message
    const messageEl = document.getElementById('confirm-modal-message');
    if (messageEl) messageEl.textContent = message;

    // Hide confirm button, only show close
    const confirmBtn = document.getElementById('confirm-modal-confirm-btn');
    if (confirmBtn) confirmBtn.classList.add('d-none');

    // Show modal
    modal.classList.add('is-active');

    // Reset button visibility on close
    const originalClose = closeConfirmModal;
    closeConfirmModal = function() {
        if (confirmBtn) confirmBtn.classList.remove('d-none');
        closeConfirmModal = originalClose;
        originalClose();
    };
}

// Close modal on escape key
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        const modal = document.getElementById('confirm-modal');
        if (modal && modal.classList.contains('is-active')) {
            closeConfirmModal();
        }
    }
});

// Close modal when clicking background
document.addEventListener('click', function(event) {
    const modal = document.getElementById('confirm-modal');
    if (modal && event.target === modal.querySelector('.modal-background')) {
        closeConfirmModal();
    }
});

// Expose functions globally
window.showConfirmModal = showConfirmModal;
window.closeConfirmModal = closeConfirmModal;
window.confirmModalAction = confirmModalAction;
window.showDeleteConfirm = showDeleteConfirm;
window.showWarningConfirm = showWarningConfirm;
window.showAlertModal = showAlertModal;
