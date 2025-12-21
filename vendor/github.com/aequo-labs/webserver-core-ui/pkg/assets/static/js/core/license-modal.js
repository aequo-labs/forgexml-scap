// license-modal.js - License Management Modal functionality

// License state
let licenseState = {
    licensed: false,
    activated: false,
    licenseId: '',
    customerName: '',
    productName: '',
    licenseType: '',
    expiryDate: '',
    machineId: '',
    features: [],
    message: '',
    expired: false,
    graceperiod: false
};

// Open the license modal
function openLicenseModal() {
    const modal = document.getElementById('license-modal');
    if (modal) {
        modal.classList.add('is-active');
        loadLicenseStatus();
    }
}

// Close the license modal
function closeLicenseModal() {
    const modal = document.getElementById('license-modal');
    if (modal) {
        modal.classList.remove('is-active');
    }
}

// Load license status from server
async function loadLicenseStatus() {
    showLicenseLoading(true);
    hideLicenseError();
    hideLicenseSuccess();
    
    try {
        const response = await fetch('/api/license/status', {
            method: 'GET',
            credentials: 'same-origin'
        });
        
        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }
        
        const data = await response.json();
        licenseState = {
            licensed: data.licensed || false,
            activated: data.activated || false,
            licenseId: data.licenseId || '',
            customerName: data.customerName || '',
            productName: data.productName || '',
            licenseType: data.licenseType || '',
            expiryDate: data.expiryDate || '',
            machineId: data.machineId || '',
            features: data.features || [],
            message: data.message || '',
            expired: data.expired || false,
            graceperiod: data.graceperiod || false
        };
        
        updateLicenseUI();
        updateNavbarLicenseIcon();
    } catch (error) {
        console.error('Failed to load license status:', error);
        showLicenseError('Failed to load license status: ' + error.message);
        // Show unlicensed state on error
        licenseState.licensed = false;
        updateLicenseUI();
    } finally {
        showLicenseLoading(false);
    }
}

// Update the UI based on license state
function updateLicenseUI() {
    const contentEl = document.getElementById('license-content');
    const loadingEl = document.getElementById('license-loading');
    
    // Hide loading, show content
    hideElement(loadingEl);
    showElement(contentEl);
    
    // Update status section
    const licensedEl = document.getElementById('license-status-licensed');
    const unlicensedEl = document.getElementById('license-status-unlicensed');
    const expiredEl = document.getElementById('license-status-expired');
    const detailsEl = document.getElementById('license-details-section');
    const actionBtns = document.getElementById('license-action-buttons');
    const installTitle = document.getElementById('install-section-title');
    
    // Hide all status notifications first
    hideElement(licensedEl);
    hideElement(unlicensedEl);
    hideElement(expiredEl);
    
    if (licenseState.licensed) {
        if (licenseState.expired) {
            // Expired license
            if (expiredEl) {
                showElement(expiredEl);
                const expiredMsg = document.getElementById('license-expired-message');
                if (expiredMsg) {
                    expiredMsg.textContent = `License expired on ${formatDate(licenseState.expiryDate)}. Please renew your license.`;
                }
            }
        } else {
            // Valid license
            if (licensedEl) {
                showElement(licensedEl);
                const statusMsg = document.getElementById('license-status-message');
                if (statusMsg) {
                    statusMsg.textContent = licenseState.message || `Valid until ${formatDate(licenseState.expiryDate)}`;
                }
            }
        }
        
        // Show license details
        showElement(detailsEl);
        updateLicenseDetails();
        
        // Show action buttons
        showElement(actionBtns);
        updateActionButtons();
        
        // Change install section title
        if (installTitle) installTitle.textContent = 'Replace License';
    } else {
        // Unlicensed
        showElement(unlicensedEl);
        hideElement(detailsEl);
        hideElement(actionBtns);
        if (installTitle) installTitle.textContent = 'Install License';
        
        // Show machine ID even when unlicensed
        if (licenseState.machineId) {
            if (detailsEl) {
                showElement(detailsEl);
                // Hide most details, just show machine ID
                const rows = detailsEl.querySelectorAll('tr');
                rows.forEach(row => {
                    const label = row.querySelector('td:first-child');
                    if (label && label.textContent !== 'Machine ID') {
                        hideElement(row);
                    } else {
                        showElement(row);
                    }
                });
            }
            const machineIdEl = document.getElementById('license-machine-id');
            if (machineIdEl) machineIdEl.textContent = licenseState.machineId;
        }
    }
}

// Update license details in the UI
function updateLicenseDetails() {
    const idEl = document.getElementById('license-id');
    const customerEl = document.getElementById('license-customer');
    const productEl = document.getElementById('license-product');
    const typeEl = document.getElementById('license-type');
    const expiryEl = document.getElementById('license-expiry');
    const machineIdEl = document.getElementById('license-machine-id');
    const activatedEl = document.getElementById('license-activated');
    const notActivatedEl = document.getElementById('license-not-activated');
    const featuresSection = document.getElementById('license-features-section');
    const featuresEl = document.getElementById('license-features');
    
    // Show all rows
    const detailsEl = document.getElementById('license-details-section');
    if (detailsEl) {
        const rows = detailsEl.querySelectorAll('tr');
        rows.forEach(row => showElement(row));
    }
    
    if (idEl) idEl.textContent = licenseState.licenseId;
    if (customerEl) customerEl.textContent = licenseState.customerName;
    if (productEl) productEl.textContent = licenseState.productName;
    if (typeEl) typeEl.textContent = licenseState.licenseType;
    if (expiryEl) expiryEl.textContent = formatDate(licenseState.expiryDate);
    if (machineIdEl) machineIdEl.textContent = licenseState.machineId;
    
    // Activation status
    if (licenseState.activated) {
        showElement(activatedEl);
        hideElement(notActivatedEl);
    } else {
        hideElement(activatedEl);
        showElement(notActivatedEl);
    }
    
    // Features
    if (licenseState.features && licenseState.features.length > 0) {
        showElement(featuresSection);
        if (featuresEl) {
            featuresEl.innerHTML = licenseState.features.map(feature => 
                `<span class="tag is-info is-light">${escapeHtml(feature)}</span>`
            ).join('');
        }
    } else {
        hideElement(featuresSection);
    }
}

// Update action buttons based on state
function updateActionButtons() {
    const activateBtn = document.getElementById('license-activate-btn');
    const deactivateBtn = document.getElementById('license-deactivate-btn');
    
    if (licenseState.licensed && !licenseState.expired) {
        if (licenseState.activated) {
            hideElement(activateBtn);
            showElement(deactivateBtn);
        } else {
            showElement(activateBtn);
            hideElement(deactivateBtn);
        }
    } else {
        hideElement(activateBtn);
        hideElement(deactivateBtn);
    }
}

// Update navbar license icon color based on license status
// Colors: green = licensed & valid, blue = grace period, yellow/warning = unlicensed, red = expired
function updateNavbarLicenseIcon() {
    const icon = document.getElementById('license-icon');
    if (!icon) return;
    
    // Remove all color classes
    icon.classList.remove('has-text-success', 'has-text-info', 'has-text-warning', 'has-text-danger');
    
    if (licenseState.licensed) {
        if (licenseState.expired) {
            // Expired - red
            icon.classList.add('has-text-danger');
        } else if (licenseState.graceperiod) {
            // Grace period - blue
            icon.classList.add('has-text-info');
        } else {
            // Valid license - green
            icon.classList.add('has-text-success');
        }
    } else {
        // Unlicensed - yellow/warning
        icon.classList.add('has-text-warning');
    }
}

// Fetch license status on page load to set navbar icon color
async function initLicenseStatus() {
    try {
        const response = await fetch('/api/license/status', {
            method: 'GET',
            credentials: 'same-origin'
        });
        
        if (!response.ok) {
            return;
        }
        
        const data = await response.json();
        licenseState = {
            licensed: data.licensed || false,
            activated: data.activated || false,
            licenseId: data.licenseId || '',
            customerName: data.customerName || '',
            productName: data.productName || '',
            licenseType: data.licenseType || '',
            expiryDate: data.expiryDate || '',
            machineId: data.machineId || '',
            features: data.features || [],
            message: data.message || '',
            expired: data.expired || false,
            graceperiod: data.graceperiod || false
        };
        
        updateNavbarLicenseIcon();
    } catch (error) {
        console.error('Failed to init license status:', error);
    }
}

// Install license from key
async function installLicense() {
    const keyInput = document.getElementById('license-key-input');
    const installBtn = document.getElementById('license-install-btn');
    
    if (!keyInput || !keyInput.value.trim()) {
        showLicenseError('Please enter a license key.');
        return;
    }
    
    const licenseKey = keyInput.value.trim();
    
    // Disable button and show loading
    if (installBtn) {
        installBtn.classList.add('is-loading');
        installBtn.disabled = true;
    }
    hideLicenseError();
    hideLicenseSuccess();
    
    try {
        const response = await fetch('/api/license/install', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'same-origin',
            body: JSON.stringify({ licenseKey: licenseKey })
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.error || `HTTP ${response.status}`);
        }
        
        showLicenseSuccess(data.message || 'License installed successfully!');
        keyInput.value = '';
        
        // Reload license status
        await loadLicenseStatus();
    } catch (error) {
        console.error('Failed to install license:', error);
        showLicenseError('Failed to install license: ' + error.message);
    } finally {
        if (installBtn) {
            installBtn.classList.remove('is-loading');
            installBtn.disabled = false;
        }
    }
}

// Activate license
async function activateLicense() {
    const activateBtn = document.getElementById('license-activate-btn');
    
    if (activateBtn) {
        activateBtn.classList.add('is-loading');
        activateBtn.disabled = true;
    }
    hideLicenseError();
    hideLicenseSuccess();
    
    try {
        const response = await fetch('/api/license/activate', {
            method: 'POST',
            credentials: 'same-origin'
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.error || `HTTP ${response.status}`);
        }
        
        showLicenseSuccess(data.message || 'License activated successfully!');
        
        // Reload license status
        await loadLicenseStatus();
    } catch (error) {
        console.error('Failed to activate license:', error);
        showLicenseError('Failed to activate license: ' + error.message);
    } finally {
        if (activateBtn) {
            activateBtn.classList.remove('is-loading');
            activateBtn.disabled = false;
        }
    }
}

// Deactivate license
async function deactivateLicense() {
    const deactivateBtn = document.getElementById('license-deactivate-btn');
    
    if (deactivateBtn) {
        deactivateBtn.classList.add('is-loading');
        deactivateBtn.disabled = true;
    }
    hideLicenseError();
    hideLicenseSuccess();
    
    try {
        const response = await fetch('/api/license/deactivate', {
            method: 'POST',
            credentials: 'same-origin'
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.error || `HTTP ${response.status}`);
        }
        
        showLicenseSuccess(data.message || 'License deactivated successfully!');
        
        // Reload license status
        await loadLicenseStatus();
    } catch (error) {
        console.error('Failed to deactivate license:', error);
        showLicenseError('Failed to deactivate license: ' + error.message);
    } finally {
        if (deactivateBtn) {
            deactivateBtn.classList.remove('is-loading');
            deactivateBtn.disabled = false;
        }
    }
}

// Helper functions
function showLicenseLoading(show) {
    const loadingEl = document.getElementById('license-loading');
    const contentEl = document.getElementById('license-content');
    
    setElementVisible(loadingEl, show);
    setElementVisible(contentEl, !show);
}

function showLicenseError(message) {
    const errorEl = document.getElementById('license-error');
    const msgEl = document.getElementById('license-error-message');
    
    showElement(errorEl);
    if (msgEl) msgEl.textContent = message;
}

function hideLicenseError() {
    const errorEl = document.getElementById('license-error');
    hideElement(errorEl);
}

function showLicenseSuccess(message) {
    const successEl = document.getElementById('license-success');
    const msgEl = document.getElementById('license-success-message');
    
    showElement(successEl);
    if (msgEl) msgEl.textContent = message;
}

function hideLicenseSuccess() {
    const successEl = document.getElementById('license-success');
    hideElement(successEl);
}

function formatDate(dateStr) {
    if (!dateStr) return 'N/A';
    try {
        const date = new Date(dateStr);
        return date.toLocaleDateString(undefined, { 
            year: 'numeric', 
            month: 'long', 
            day: 'numeric' 
        });
    } catch {
        return dateStr;
    }
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Helper functions for class-based visibility (avoids inline styles)
function showElement(el) {
    if (el) el.classList.remove('d-none');
}

function hideElement(el) {
    if (el) el.classList.add('d-none');
}

function setElementVisible(el, visible) {
    if (el) {
        if (visible) {
            el.classList.remove('d-none');
        } else {
            el.classList.add('d-none');
        }
    }
}

// Close modal when clicking outside
document.addEventListener('click', function(event) {
    const modal = document.getElementById('license-modal');
    if (modal && event.target === modal.querySelector('.modal-background')) {
        closeLicenseModal();
    }
});

// Close modal on escape key
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        closeLicenseModal();
    }
});

// Initialize license button click handler and fetch initial status
document.addEventListener('DOMContentLoaded', function() {
    const licenseBtn = document.getElementById('license-button');
    if (licenseBtn) {
        licenseBtn.addEventListener('click', openLicenseModal);
        // Fetch license status to set navbar icon color
        initLicenseStatus();
    }
});

// Expose functions globally
window.openLicenseModal = openLicenseModal;
window.closeLicenseModal = closeLicenseModal;
window.installLicense = installLicense;
window.activateLicense = activateLicense;
window.deactivateLicense = deactivateLicense;
window.hideLicenseError = hideLicenseError;
window.hideLicenseSuccess = hideLicenseSuccess;
