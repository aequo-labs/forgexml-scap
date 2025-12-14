/*
* Licensed to the Apache Software Foundation (ASF) under one
* or more contributor license agreements.  See the NOTICE file
* distributed with this work for additional information
* regarding copyright ownership.  The ASF licenses this file
* to you under the Apache License, Version 2.0 (the
* "License"); you may not use this file except in compliance
* with the License.  You may obtain a copy of the License at
*
*   http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
*/
(function(root, factory) {
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
})(this, function(exports, echarts) {
    var log = function(msg) {
        if (typeof console !== 'undefined') {
            console && console.error && console.error(msg);
        }
    };
    if (!echarts) {
        log('ECharts is not Loaded');
        return;
    }

    var colorPalette = [
        '#fc97af',
        '#87f7cf',
        '#f7f494',
        '#72ccff',
        '#f7c5a0',
        '#d4a4eb',
        '#d2f5a6',
        '#76f2f2'
    ];

    var theme = {
        color: colorPalette,

        backgroundColor: '#293441',

        title: {
            textStyle: {
                fontWeight: 'normal',
                color: '#ffffff'
            }
        },

        visualMap: {
            color: ['#fc97af', '#87f7cf']
        },

        toolbox: {
            iconStyle: {
                borderColor: '#999999'
            }
        },

        tooltip: {
            backgroundColor: 'rgba(50,50,50,0.8)',
            textStyle: {
                color: '#ffffff'
            }
        },

        dataZoom: {
            dataBackgroundColor: '#dedede',
            fillerColor: 'rgba(154,217,247,0.2)',
            handleColor: '#dddddd'
        },

        timeline: {
            lineStyle: {
                color: '#ffffff'
            },
            controlStyle: {
                color: '#ffffff',
                borderColor: '#ffffff'
            }
        },

        candlestick: {
            itemStyle: {
                color: '#fc97af',
                color0: '#87f7cf'
            },
            lineStyle: {
                width: 1,
                color: '#fc97af',
                color0: '#87f7cf'
            },
            areaStyle: {
                color: '#fc97af',
                color0: '#87f7cf'
            }
        },

        graph: {
            itemStyle: {
                color: '#fc97af'
            },
            linkStyle: {
                color: '#ffffff'
            }
        },

        map: {
            itemStyle: {
                color: '#fc97af',
                borderColor: '#444444',
                areaColor: '#323c48'
            },
            areaStyle: {
                color: '#323c48'
            },
            label: {
                color: '#ffffff'
            }
        },

        gauge: {
            axisLine: {
                show: true,
                lineStyle: {
                    color: [
                        [0.2, '#87f7cf'],
                        [0.8, '#fc97af'],
                        [1, '#f7f494']
                    ],
                    width: 5
                }
            },
            axisTick: {
                splitNumber: 10,
                length: 8,
                lineStyle: {
                    color: 'auto'
                }
            },
            axisLabel: {
                color: 'auto'
            },
            splitLine: {
                length: 12,
                lineStyle: {
                    color: 'auto'
                }
            },
            pointer: {
                length: '90%',
                width: 3,
                color: 'auto'
            },
            title: {
                color: '#ffffff'
            },
            detail: {
                color: 'auto'
            }
        },
        
        textStyle: {
            color: '#ffffff'
        },
        
        legend: {
            textStyle: {
                color: '#ffffff'
            }
        },
        
        xAxis: {
            axisLine: {
                lineStyle: {
                    color: '#ffffff'
                }
            },
            axisLabel: {
                textStyle: {
                    color: '#ffffff'
                }
            },
            splitLine: {
                lineStyle: {
                    color: '#484753'
                }
            }
        },
        
        yAxis: {
            axisLine: {
                lineStyle: {
                    color: '#ffffff'
                }
            },
            axisLabel: {
                textStyle: {
                    color: '#ffffff'
                }
            },
            splitLine: {
                lineStyle: {
                    color: '#484753'
                }
            }
        }
    };
    echarts.registerTheme('chalk', theme);
});