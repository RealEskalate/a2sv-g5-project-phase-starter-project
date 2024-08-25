import React from "react";
import Image from "next/image";
import { PersonProps } from "@/types/index.";


const TransferMembers: React.FC<PersonProps> = ({ person, selectedPerson }) => {
  return (
    <div
      className={`text-center ${
        selectedPerson === person.name ? "border-2 border-blue-600 p-2" : ""
      }`}
    >
      <Image
        src={person.img}
        alt={person.name}
        className="rounded-full mx-auto mb-2"
        width={48}
        height={48}
      />
      <p className="text-sm font-semibold">{person.name}</p>
      <p className="text-xs text-gray-500">{person.role}</p>
    </div>
  );
};

export default TransferMembers;
