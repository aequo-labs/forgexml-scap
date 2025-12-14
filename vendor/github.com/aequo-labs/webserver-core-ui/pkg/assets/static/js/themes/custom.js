/*
* Custom theme for Go-Algo that uses CSS variables
* This theme will adapt to light/dark mode automatically
*/
(function (root, factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as an anonymous module.
        define(['exports', 'echarts'], factory);
    } else if (
        typeof exports === 'object' &&
        typeof exports.nodeName !== 'string'
    ) {
        // CommonJS
        factory(exports, require('echarts/lib/echarts'));
    } else {
        // Browser globals
        factory({}, root.echarts);
    }
})(this, function (exports, echarts) {
    var log = function (msg) {
        if (typeof console !== 'undefined') {
            console && console.error && console.error(msg);
        }
    };
    if (!echarts) {
        log('ECharts is not Loaded');
        return;
    }

    // Function to get CSS variable value with fallback
    function getCssVar(varName, fallback) {
        if (typeof document === 'undefined') return fallback;

        // Try to get the CSS variable
        var value = getComputedStyle(document.documentElement).getPropertyValue(varName).trim();

        // If the value is empty, try to get it from the parent document (if in an iframe)
        if (!value && window.parent && window.parent.document) {
            try {
                value = getComputedStyle(window.parent.document.documentElement).getPropertyValue(varName).trim();
            } catch (e) {
                // Ignore cross-origin errors
                console.warn('Could not access parent document CSS variables:', e);
            }
        }

        return value || fallback;
    }

    // Detect if we're in dark mode
    function isDarkMode() {
        try {
            // First check if the document has dark-mode class
            if (document.documentElement.classList.contains('dark-mode')) {
                return true;
            }

            // Then check if the parent document has dark-mode class (if in an iframe)
            if (window.parent && window.parent.document) {
                try {
                    return window.parent.document.documentElement.classList.contains('dark-mode');
                } catch (e) {
                    // Ignore cross-origin errors
                }
            }

            return false;
        } catch (e) {
            return false; // Default to light mode
        }
    }

    // Determine if we're in dark mode
    var darkMode = isDarkMode();
    console.log('Dark mode detected:', darkMode);

    // Define colors based on CSS variables with fallbacks
    var colors = {
        // Light mode fallbacks (from light-mode.css)
        light: {
            primaryColor: '#3273dc',
            secondaryColor: '#4a4a4a',
            successColor: '#23d160',
            warningColor: '#ffdd57',
            dangerColor: '#ff3860',
            infoColor: '#209cee',
            textColor: '#333333',
            textMuted: '#666666',
            bgColor: '#ffffff',
            axisTick: '#cccccc',
            axisLine: '#cccccc',
            axisLabelColor: '#333333',
            splitLine: 'rgba(0, 0, 0, 0.1)',
            borderColor: '#e1e4e8',
            tooltipBg: 'rgba(255, 255, 255, 0.95)',
            tooltipBorder: '#ddd',
            inactiveColor: '#aaaaaa'
        },
        // Dark mode fallbacks (from dark-mode.css)
        dark: {
            primaryColor: '#4a9eff',
            secondaryColor: '#b5b5b5',
            successColor: '#48c774',
            warningColor: '#ffe08a',
            dangerColor: '#f14668',
            infoColor: '#3298dc',
            textColor: '#f5f5f5',
            textMuted: '#aaaaaa',
            bgColor: '#1a1a1a',
            axisTick: '#555555',
            axisLine: '#555555',
            axisLabelColor: '#f5f5f5',
            splitLine: 'rgba(255, 255, 255, 0.1)',
            borderColor: '#3e3e3e',
            tooltipBg: 'rgba(40, 40, 40, 0.95)',
            tooltipBorder: '#3e3e3e',
            inactiveColor: '#777777'
        }
    };

    // Get theme colors based on mode
    var theme = darkMode ? colors.dark : colors.light;

    // Set up color palette - these colors work well in both modes
    /*
    var colorPalette = [
        theme.primaryColor,
        theme.successColor,
        theme.infoColor,
        theme.warningColor,
        theme.dangerColor,
        theme.secondaryColor
    ];
    */
    var colorPalette = [
        '#1E90FF',                 // Bright Blue (Principal)
        '#32CD32',                 // Lime Green (Realized Gains)
        '#FF4500',                 // Orange-Red (Unrealized Gains)
        '#FFD700',                 // Gold (Market Value)
        '#FF1493'                  // Deep Pink (Cash)
    ];

    // Base theme configuration that's the same for both modes
    var baseThemeConfig = {
        color: colorPalette,
        backgroundColor: 'transparent',
        textStyle: {
            color: theme.textColor
        },
        title: {
            textStyle: {
                color: theme.textColor
            },
            subtextStyle: {
                color: theme.textMuted
            }
        },
        legend: {
            textStyle: {
                color: theme.textColor
            },
            inactiveColor: theme.inactiveColor,
            pageIconColor: theme.primaryColor,
            pageIconInactiveColor: theme.textMuted,
            pageTextStyle: {
                color: theme.textColor
            }
        },
        tooltip: {
            backgroundColor: theme.tooltipBg,
            borderColor: theme.tooltipBorder,
            textStyle: {
                color: theme.textColor
            }
        },
        xAxis: {
            axisLine: {
                lineStyle: {
                    color: theme.axisLine
                }
            },
            axisTick: {
                lineStyle: {
                    color: theme.axisTick
                }
            },
            axisLabel: {
                color: theme.axisLabelColor
            },
            splitLine: {
                lineStyle: {
                    color: theme.splitLine
                }
            }
        },
        yAxis: {
            axisLine: {
                lineStyle: {
                    color: theme.axisLine
                }
            },
            axisTick: {
                lineStyle: {
                    color: theme.axisTick
                }
            },
            axisLabel: {
                color: theme.axisLabelColor
            },
            splitLine: {
                lineStyle: {
                    color: theme.splitLine
                }
            }
        },
        line: {
            symbol: 'circle'
        }
    };

    // Determine final theme config based on mode
    var themeConfig;
    if (darkMode) {
        // For dark mode, clone the base config and add axis pointer styling
        themeConfig = JSON.parse(JSON.stringify(baseThemeConfig));
        themeConfig.tooltip.axisPointer = {
            label: {
                color: '#333333',
                backgroundColor: '#ffffff'
            }
        };
    } else {
        // For light mode, use base config as is (no axis pointer styling)
        themeConfig = baseThemeConfig;
    }

    // Register the theme
    echarts.registerTheme('custom', themeConfig);

    // Fix for the tooltip axis pointer labels
    var originalInit = echarts.init;
    echarts.init = function () {
        var chart = originalInit.apply(this, arguments);

        // Now let's add specific overrides for problematic elements
        var tooltipStyleId = 'echarts-theme-fixes';
        if (!document.getElementById(tooltipStyleId)) {
            var style = document.createElement('style');
            style.id = tooltipStyleId;

            style.textContent = `
                /* Fix tooltip axis pointers - use CSS variables for theme awareness */
                .ec-tooltip-axisPointer label {
                    color: var(--button-text-inverse) !important;
                    background-color: rgba(var(--primary-color-rgb), 0.95) !important;
                    border: 1px solid var(--border-color) !important;
                    border-radius: 4px !important;
                    padding: 4px 8px !important;
                    font-size: 12px !important;
                    box-shadow: 0 2px 8px rgba(0,0,0,0.2) !important;
                }
                
                /* Alternative fallback for better contrast in both modes */
                html.dark-mode .ec-tooltip-axisPointer label {
                    color: #ffffff !important;
                    background-color: rgba(40, 40, 40, 0.95) !important;
                    border: 1px solid #3e3e3e !important;
                }
                
                html:not(.dark-mode) .ec-tooltip-axisPointer label {
                    color: #333333 !important;
                    background-color: rgba(255, 255, 255, 0.95) !important;
                    border: 1px solid #ddd !important;
                }
                
                /* Fix inactive legend items - force specific colors */
                .legend .inactive text,
                g.legend .inactive text {
                    fill: ${theme.inactiveColor} !important;
                    opacity: 0.7 !important;
                }
            `;

            document.head.appendChild(style);
        }

        return chart;
    };

    console.log('ECharts theme registered with mode-specific fixes');
});