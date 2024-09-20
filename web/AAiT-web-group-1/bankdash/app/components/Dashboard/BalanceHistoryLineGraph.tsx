"use client";

import React, { useRef, useEffect } from 'react';
import { Chart, LineController, ArcElement, CategoryScale, LinearScale, PointElement, LineElement, Filler, Tooltip, Legend, ScriptableContext, } from "chart.js";
Chart.register(ArcElement, Tooltip, Filler, Legend, CategoryScale, LinearScale, PointElement, LineController, LineElement, PointElement, LinearScale, CategoryScale );

interface GRAPHDATA {
    balanceHistory: number[];
}


const BalanceHistoryLineGraph = ({balanceHistory}:GRAPHDATA) => {
    const chartRef = useRef<HTMLCanvasElement>(null);
    const chartInstanceRef = useRef<Chart | null>(null);

    useEffect(() => {
        if (chartRef.current && !chartInstanceRef.current) {
            const context = chartRef.current.getContext('2d');

            if (context) {
                // Create a linear gradient
                chartInstanceRef.current = new Chart(context, {
                    type: 'line',
                    data: {
                        labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'], // Months
                        datasets: [{
                            label: '',
                            data: balanceHistory,// Sample data
                            borderWidth: 2,
                            pointBackgroundColor: '#1D8CF8',
                            pointBorderColor: '#ffffff',
                            pointBorderWidth: 2,
                            pointRadius: 0, // Remove the points
                            tension: 0.4, // Smooth curve
                            
                            backgroundColor: (context: ScriptableContext<"line">) => {
                                const ctx = context.chart.ctx;
                                const gradient = ctx.createLinearGradient(0, 0, 0, 250);
                                gradient.addColorStop(0, "rgba(91,56,237,0.45)");
                                gradient.addColorStop(1, "rgba(91,56,237,0.0)");
                                return gradient;
                            }, //background gradient color
                            borderColor: '#1814F3', // Line color
                            fill: true,
                        }]
                    },
                    options: {
                        scales: {
                            x: {
                                title: {
                                    display: false,
                                },
                                grid: {
                                    display: true,
                                    tickBorderDash: [5]
                                },
                                border: {
                                    dash: [5,10],
                                },  
                                ticks: {
                                    color: '#718EBF'
                                }
                            },
                            y: {
                                title: {
                                    display: false,
                                },
                                beginAtZero: true,
                                border: {
                                    dash: [5,10],
                                },  
                                ticks: {
                                    color: '#718EBF'
                                }
                            }
                        },
                        plugins: {
                            legend: {
                                display: false,
                            },
                            tooltip: {
                            },
                            datalabels: {
                                display: false // Disable data labels
                            }
                        },
                        maintainAspectRatio: false,
                    }
                });
            }
        }

        // Cleanup: Destroy the chart instance when the component unmounts
        return () => {
            if (chartInstanceRef.current) {
                chartInstanceRef.current.destroy();
                chartInstanceRef.current = null;
            }
        };
    }, []);

    return (
        <div className='h-full'>
            <h3 className="text-[22px] font-semibold text-[#343C6A] mb-5">Monthly Activity</h3>
            <div className="bg-white rounded-3xl p-6 h-72">
                <div className='relative h-full'>
                    <canvas ref={chartRef} />
                </div>
            </div>
        </div>
    );
};

export default BalanceHistoryLineGraph;
