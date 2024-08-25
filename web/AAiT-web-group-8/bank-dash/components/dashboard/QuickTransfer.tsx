"use client";
import React, { useState } from "react";
import { FaPaperPlane, FaChevronRight, FaChevronLeft } from "react-icons/fa";
import TransferMembers from "./Members";

const people = [
  { name: "Livia Bator", role: "CEO", img: "/images/livia.svg" },
  { name: "Randy Press", role: "Director", img: "/images/workman.svg" },
  { name: "Workman", role: "Designer", img: "/images/randy.svg" },
];

const QuickTransfer = () => {
  const [scrollIndex, setScrollIndex] = useState(0);
  const [selectedPerson, setSelectedPerson] = useState<string | null>(null);
  const [amount, setAmount] = useState<string>("");

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
              className={`cursor-pointer p-2 rounded-lg transition-all duration-300 transform ${
                selectedPerson === person.name
                  ? "bg-blue-100 scale-105"
                  : "bg-transparent"
              }`}
            >
              <TransferMembers
                person={person}
                selectedPerson={selectedPerson}
              />
            </div>
          ))}
        </div>
        <button onClick={handleNext} className="p-2 rounded-full bg-gray-200">
          <FaChevronRight />
        </button>
      </div>
      <div className="flex items-center gap-4">
        <p className="flex-2 flex-shrink-0 text-[#718EBF]">Write Amount</p>
        <div className="relative flex-3">
          <input
            className="w-full py-2 pl-4 pr-20 text-lg border rounded-full focus:outline-none"
            placeholder="Write Amount"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            disabled={!selectedPerson}
          />
          <button
            className="absolute right-0 top-1/2 transform -translate-y-1/2 px-4 py-2 bg-[#1814F3] text-white rounded-full shadow-lg flex items-center space-x-2"
            disabled={!selectedPerson || !amount}
          >
            <span>send</span>
            <FaPaperPlane size={20} />
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
