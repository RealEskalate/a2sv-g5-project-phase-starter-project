"use client"
import React, { useState } from 'react';
import InputMoney from "../inputMoney/InputMoney";
import PersonCard from "../personCard/PersonCard";

const SendMoney = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const cards = [
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Abebe Bekele",
      jobTitle: "Dentist",
    },
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Ayele Erede",
      jobTitle: "Scientist",
    },
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Metana Heder",
      jobTitle: "Compiler",
    },
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Metana Heder",
      jobTitle: "Wayer",
    },
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Metana Heder",
      jobTitle: "CEO",
    },
    {
      imageLink: "https://avatarfiles.alphacoders.com/347/347546.png",
      fullName: "Metana Heder",
      jobTitle: "Firecracker",
    },
  ];

  const visibleCards = cards.slice(currentIndex, currentIndex + 3);

  const handleNext = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex + 1 < cards.length - 2 ? prevIndex + 1 : 0
    );
  };

  return (
    <div className="flex flex-col gap-4 rounded-3xl bg-white w-fit px-5 pt-7 pb-9">
      <div className="flex gap-4 items-center">
        {visibleCards.map((card, index) => (
          <PersonCard
            key={index}
            imageLink={card.imageLink}
            fullName={card.fullName}
            jobTitle={card.jobTitle}
          />
        ))}
        <button
          onClick={handleNext}
          className="w-[50px] h-[50px] rounded-full shadow-custom-shadow text-center text-[#718EBF]"
        >
          {`>`}
        </button>
      </div>
      <div>
        <InputMoney />
      </div>
    </div>
  );
};

export default SendMoney;
