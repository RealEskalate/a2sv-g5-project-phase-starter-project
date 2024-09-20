"use client";

import React, { useRef, useEffect } from 'react';
import { Chart, CategoryScale, BarController, BarElement, LinearScale, Tooltip, Legend } from 'chart.js';

// Register the components
Chart.register(CategoryScale, BarController, BarElement, LinearScale, Tooltip, Legend);

// const weekdays = ['Sat', 'Sun','Mon', 'Tue', 'Wed', 'Thu', 'Fri' ]
// const deposits = [680, 300,575, 200, 425, 350, 605 ]
// const withdraws = [200, 260, 300, 400, 300, 200, 100]

interface WeeklyActivityBarChartProps {
    weekdays: string[],
    deposits: number[],
    withdraws: number[],
}

const WeeklyActivityBarChart = ({weekdays, deposits, withdraws} : WeeklyActivityBarChartProps) => {
    const chartRef = useRef<HTMLCanvasElement>(null);
    const chartInstanceRef = useRef<Chart | null>(null);

    useEffect(() => {
        if (chartRef.current && !chartInstanceRef.current) {
            const context = chartRef.current.getContext('2d');

            if (context) {
                chartInstanceRef.current = new Chart(context, {
                    type: 'bar',
                    data: {
                        labels: weekdays,
                        datasets: [{
                            label: 'Deposit',
                            data: deposits,
                            backgroundColor: '#1814F3',
                            borderRadius: 100,
                            borderSkipped: false,
                            barPercentage: 0.6,
                            categoryPercentage: 0.5,

                        },
                        {
                            label: 'Withdraw',
                            data: withdraws,
                            backgroundColor: '#16DBCC',
                            borderRadius: 100,
                            borderSkipped: false,
                            barPercentage: 0.6,
                            categoryPercentage: 0.5,
                        }]
                    },
                    options: {
                        scales: {
                           y: {
                                beginAtZero: true,
                                ticks: {
                                    color: '#718EBF' // Font color for y-axis labels
                                },
                                border :{
                                    display: false
                                }
                            },
                            x: {
                                grid: {
                                    display: false // Remove vertical grid lines
                                },
                                ticks: {
                                    color: '#718EBF' // Font color for x-axis labels
                                },
                                
                            }
                            
                        },
                        layout: {
                            padding: {
                                left: 20,
                                right: 20,
                                top: 10,
                                bottom: 10
                            }
                        },
                        plugins: {
                            legend: {
                                display: true,
                                position: 'top',
                                align: 'end',
                                labels: {
                                    usePointStyle: true,
                                    pointStyle: 'circle',
                                    color: '#718EBF',
                                    font: {
                                        size: 16, // Font size
                                    },
                                    padding: 25, 
                                    boxHeight: 12,
                                },
                            },
                            tooltip: {
                                enabled: true, 
                                usePointStyle: true,
                                callbacks: {
                                    label: function (context) {
                                        const label = context.dataset.label || '';
                                        const value = context.raw;
                                        return `${label}: $${value}`;
                                    },
                                    title: function () {
                                        return ''; // Remove the default title
                                    }
                                },
                                backgroundColor: '#ffffff',
                                titleColor: '#000000',
                                bodyColor: '#000000',
                                borderColor: '#e5e5e5',
                                borderWidth: 1,
                                padding: 10,
                                cornerRadius: 8,
                                titleFont: {
                                    size: 14
                                },
                                bodyFont: {
                                    size: 12
                                },
                                bodyAlign: 'center',
                                titleAlign: 'center'
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
            <h3 className="text-[22px] font-semibold text-[#343C6A] mb-4">Weekly Activity</h3>
            <div className="bg-white rounded-3xl h-[350px] ">
                <div className='relative ml-4 h-full text-sm text-[#718EBF] ' >
                    <canvas ref={chartRef} />
                </div>
            </div>
        </div>
    );
};

export default WeeklyActivityBarChart;