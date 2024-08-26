'use client';
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_KEY = 'a20eaf91948813369f6b3fac';  
const API_BASE_URL = `https://v6.exchangerate-api.com/v6/${API_KEY}/latest/ETB`;

const CurrencyConverter: React.FC = () => {
  const [currencies, setCurrencies] = useState<string[]>([]);
  const [fromCurrency, setFromCurrency] = useState<string>('ETB');
  const [toCurrency, setToCurrency] = useState<string>('USD');
  const [amount, setAmount] = useState<string>("");
  const [convertedAmount, setConvertedAmount] = useState<number | null>(null);

  useEffect(() => {
    const fetchCurrencies = async () => {
      try {
        const response = await axios.get(API_BASE_URL);
        const currencyCodes = Object.keys(response.data.conversion_rates);
        setCurrencies(currencyCodes);
      } catch (error) {
        console.error('Error fetching currencies:', error);
      }
    };

    fetchCurrencies();
  }, []);

  const handleConvert = async () => {
    try {
      const response = await axios.get(API_BASE_URL);
      const conversionRate = response.data.conversion_rates[toCurrency];
      setConvertedAmount(Number(amount) * conversionRate);
    } catch (error) {
      console.error('Error converting currency:', error);
    }
  };

  return (
    <div className='w-full  flex items-center justify-center py-5'>
      <div className='bg-white p-6 rounded-lg shadow-lg w-full max-w-md'>
        <h1 className='text-3xl font-bold text-center mb-6 text-blue-600'>Currency Converter</h1>
        <div className='mb-4'>
          <label className='block text-gray-700 mb-2'>From:</label>
          <select 
            value={fromCurrency} 
            onChange={(e) => setFromCurrency(e.target.value)}
            className='w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring focus:border-blue-500'
          >
            {currencies.map(currency => (
              <option key={currency} value={currency}>{currency}</option>
            ))}
          </select>
        </div>
        <div className='mb-4'>
          <label className='block text-gray-700 mb-2'>To:</label>
          <select 
            value={toCurrency} 
            onChange={(e) => setToCurrency(e.target.value)}
            className='w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring focus:border-blue-500'
          >
            {currencies.map(currency => (
              <option key={currency} value={currency}>{currency}</option>
            ))}
          </select>
        </div>
        <div className='mb-4'>
          <label className='block text-gray-700 mb-2'>Amount:</label>
          <input 
            type="number" 
            value={amount} 
            onChange={(e) => setAmount(e.target.value)} 
            className='w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring focus:border-blue-500'
          />
        </div>
        <div className='flex items-center justify-center'>
        <button 
          onClick={handleConvert} 
          className='bg-blue-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline'
        //   ' bg-blue-600 text-white py-3 rounded-md hover:bg-blue-700 transition duration-300'
        >
          Convert
        </button>
        </div>
        {convertedAmount !== null && (
          <div className='mt-6 p-4 bg-green-100 rounded-md text-green-800 text-center'>
            <h2 className='text-2xl font-semibold'>Converted Amount: {convertedAmount.toFixed(2)}</h2>
          </div>
        )}
      </div>
    </div>
  );
};

export default CurrencyConverter;
