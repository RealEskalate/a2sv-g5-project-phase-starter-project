"use client";

import React, { useRef, useEffect } from 'react';
import { Chart, CategoryScale, BarController, BarElement, LinearScale, Tooltip, Legend } from 'chart.js';

// Register the components
Chart.register(CategoryScale, BarController, BarElement, LinearScale, Tooltip, Legend);

// const expenses = [11200, 16200, 12400, 6400, 15000, 12000, 12100]
// const months = ['Aug', 'Sep', 'Oct', 'Nov', 'Dec', 'Jan']


interface MyExpenseBarChartProps {
    Expenses: number[],
    months: string[]
}

const MyExpenseBarChart = ({Expenses, months} : MyExpenseBarChartProps) => {

    const chartRef = useRef<HTMLCanvasElement>(null);
    const chartInstanceRef = useRef<Chart | null>(null);

    useEffect(() => {
        if (chartRef.current && !chartInstanceRef.current) {
            const context = chartRef.current.getContext('2d');

            if (context) {
                chartInstanceRef.current = new Chart(context, {
                    type: 'bar',
                    data: {
                        labels: months,
                        datasets: [
                        {
                            label: 'months',
                            data: Expenses,
                            backgroundColor: '#EDF0F7',
                            borderRadius: 10,
                            borderSkipped: false,
                            barPercentage: 0.8,
                            categoryPercentage: 0.8,
                            hoverBackgroundColor: '#16DBCC',
                        }]
                    },
                    options: {
                        scales: {
                           y: {
                                
                                beginAtZero: true,
                                grid: {
                                    display: false,
                                },
                                ticks: {
                                    display: false,
                                    color: '#718EBF' // Font color for y-axis labels
                                },
                                border: {
                                    display: false, // Remove the y-axis line
                                }
                            },
                            x: {
                                grid: {
                                    display: false // Remove vertical grid lines
                                },
                                ticks: {
                                    color: '#718EBF' // Font color for x-axis labels
                                },
                                border: {
                                    display: false, // Remove the y-axis line
                                }
                                
                            }
                            
                        },
                        layout: {
                            padding: {
                                left: 0,
                                right: 20,
                                top: 10,
                                bottom: 10
                            }
                        },
                        plugins: {
                            legend: {
                                display: false,
                            },
                            tooltip: {
                                enabled: true, 
                                usePointStyle: false,
                                displayColors: false,
                                callbacks: {
                                    //remove circle from tooltip

                                    label: function (context) {
                                        const value = context.raw;
                                        return `$${value}`;
                                    },
                                    title: function () {
                                        return ''; // Remove the default title
                                    }
                                },
                                backgroundColor: '#ffffff',
                                titleColor: '#000000',
                                bodyColor: '#000000',
                                borderColor: '#16DBCC',
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
            <div className=''>
                <div className="bg-white rounded-3xl">
                    <div className='relative h-[200px] text-sm text-barTextGray' >
                        <canvas ref={chartRef} />
                    </div>
                </div>
            </div>
    );
};


export default MyExpenseBarChart