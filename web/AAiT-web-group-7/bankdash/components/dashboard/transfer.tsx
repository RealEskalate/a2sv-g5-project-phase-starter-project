'use client';
import React, { useState } from 'react';
import { FaPaperPlane, FaChevronRight, FaChevronLeft } from 'react-icons/fa';
import TransferMembers from './transferMembers';

const people = [
  { name: 'Livia Bator', role: 'CEO', img: '/images/livia.jpg' },
  { name: 'Randy Press', role: 'Director', img: '/images/randy.jpg' },
  { name: 'Workman', role: 'Designer', img: '/images/workman.jpg' },
  { name: 'miki simon', role: 'Manager', img: '/images/additional.jpg' },
];

const SendMoneyComponent = () => {
  const [scrollIndex, setScrollIndex] = useState(0);
  const [selectedPerson, setSelectedPerson] = useState<string | null>(null);

  const handleNext = () => {
    setScrollIndex((prev) => (prev + 1) % people.length);
  };

  const handlePrev = () => {
    setScrollIndex((prev) => (prev - 1 + people.length) % people.length);
  };

  const handlePersonClick = (name: string) => {
    setSelectedPerson(name);
  };

  return (
    <div className="p-7 bg-white rounded-xl shadow-lg w-full md:w-full max-w-lg mx-auto">
      <div className="flex items-center justify-between mb-4">
        <button onClick={handlePrev} className="p-2 rounded-full bg-gray-200">
          <FaChevronLeft />
        </button>
        <div className="flex items-center justify-center space-x-4 overflow-hidden w-full">
          {people.slice(scrollIndex, scrollIndex + 3).map((person, index) => (
            <div
              key={index}
              onClick={() => handlePersonClick(person.name)}
              className={`cursor-pointer p-2 rounded-lg`}
            >
              <TransferMembers person={person} selectedPerson={selectedPerson} />
            </div>
          ))}
        </div>
        <button onClick={handleNext} className="p-2 rounded-full bg-gray-200">
          <FaChevronRight />
        </button>
      </div>
      <div className='flex items-center gap-4'>
        <p className='flex-2'>Write amount</p>
        <div className="relative flex-3">
          <input
            className="w-full py-2 pl-3 pr-20 text-lg border rounded-full focus:outline-none"
            placeholder="Write Amount"
            disabled={!selectedPerson}
          />
          <button
            className="absolute right-2 top-2 px-4 py-2 bg-blue-600 text-white rounded-full"
            disabled={!selectedPerson}
          >
            <FaPaperPlane />
          </button>
        </div>
      </div>
    </div>
  );
};

export default SendMoneyComponent;
